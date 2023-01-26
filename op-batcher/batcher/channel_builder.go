package batcher

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/core/types"
)

type (
	// channelBuilder uses a ChannelOut to create a channel with output frame
	// size approximation.
	channelBuilder struct {
		cfg ChannelConfig

		// marked as full if a) max RLP input bytes, b) max num frames or c) max
		// allowed frame index (uint16) has been reached
		fullErr error
		// current channel
		co *derive.ChannelOut
		// list of blocks in the channel. Saved in case the channel must be rebuilt
		blocks []*types.Block
		// frames data queue, to be send as txs
		frames []taggedData
	}

	ChannelConfig struct {
		// ChannelTimeout is the maximum duration, in seconds, to attempt completing
		// an opened channel. The batcher can decide to set it shorter than the
		// actual timeout, since submitting continued channel data to L1 is not
		// instantaneous. It's not worth it to work with nearly timed-out channels.
		ChannelTimeout uint64
		// The maximum byte-size a frame can have.
		MaxFrameSize uint64
		// The target number of frames to create per channel. Note that if the
		// realized compression ratio is worse than the approximate, more frames may
		// actually be created. This also depends on how close TargetFrameSize is to
		// MaxFrameSize.
		TargetFrameSize uint64
		// The target number of frames to create in this channel. If the realized
		// compression ratio is worse than approxComprRatio, additional leftover
		// frame(s) might get created.
		TargetNumFrames int
		// Approximated compression ratio to assume. Should be slightly smaller than
		// average from experiments to avoid the chances of creating a small
		// additional leftover frame.
		ApproxComprRatio float64
	}

	ChannelFullError struct {
		Err error
	}
)

func (e *ChannelFullError) Error() string {
	return "channel full: " + e.Err.Error()
}

func (e *ChannelFullError) Unwrap() error {
	return e.Err
}

var (
	ErrInputTargetReached = errors.New("target amount of input data reached")
	ErrMaxFrameIndex      = errors.New("max frame index reached (uint16)")
)

// InputThreshold calculates the input data threshold in bytes from the given
// parameters.
func (c ChannelConfig) InputThreshold() uint64 {
	return uint64(float64(c.TargetNumFrames) * float64(c.TargetFrameSize) / c.ApproxComprRatio)
}

func newChannelBuilder(cfg ChannelConfig) (*channelBuilder, error) {
	co, err := derive.NewChannelOut()
	if err != nil {
		return nil, err
	}

	return &channelBuilder{
		cfg: cfg,
		co:  co,
	}, nil
}

func (c *channelBuilder) ID() derive.ChannelID {
	return c.co.ID()
}

func (c *channelBuilder) Blocks() []*types.Block {
	return c.blocks
}

func (c *channelBuilder) Reset() error {
	c.blocks = c.blocks[:0]
	c.frames = c.frames[:0]
	return c.co.Reset()
}

// AddBlock adds a block to the channel compression pipeline. IsFull should be
// called aftewards to test whether the channel is full. If full, a new channel
// must be started.
//
// AddBlock returns a ChannelFullError if called even though the channel is
// already full. See description of FullErr for details.
//
// Call OutputFrames() afterwards to create frames.
func (c *channelBuilder) AddBlock(block *types.Block) error {
	if c.IsFull() {
		return c.FullErr()
	}

	_, err := c.co.AddBlock(block)
	if errors.Is(err, derive.ErrTooManyRLPBytes) {
		c.setFullErr(err)
		return c.FullErr()
	} else if err != nil {
		return fmt.Errorf("adding block to channel out: %w", err)
	}
	c.blocks = append(c.blocks, block)

	if c.InputTargetReached() {
		c.setFullErr(ErrInputTargetReached)
		// Adding this block still worked, so don't return error, just mark as full
	}

	return nil
}

// InputTargetReached says whether the target amount of input data has been
// reached in this channel builder. No more blocks can be added afterwards.
func (c *channelBuilder) InputTargetReached() bool {
	return uint64(c.co.InputBytes()) >= c.cfg.InputThreshold()
}

// IsFull returns whether the channel is full.
// FullErr returns the reason for the channel being full.
func (c *channelBuilder) IsFull() bool {
	return c.fullErr != nil
}

// FullErr returns the reason why the channel is full. If not full yet, it
// returns nil.
//
// It returns a ChannelFullError wrapping one of three possible reasons for the
// channel being full:
//   - ErrInputTargetReached if the target amount of input data has been reached,
//   - derive.MaxRLPBytesPerChannel if the general maximum amount of input data
//     would have been exceeded by the latest AddBlock call,
//   - ErrMaxFrameIndex if the maximum number of frames has been generated (uint16)
func (c *channelBuilder) FullErr() error {
	return c.fullErr
}

func (c *channelBuilder) setFullErr(err error) {
	c.fullErr = &ChannelFullError{Err: err}
}

// OutputFrames creates new frames with the channel out. It should be called
// after AddBlock and before iterating over available frames with HasFrame and
// NextFrame.
//
// If the input data target hasn't been reached yet, it will conservatively only
// pull readily available frames from the compression output.
// If the target has been reached, the channel is closed and all remaining
// frames will be created, possibly with a small leftover frame.
func (c *channelBuilder) OutputFrames() error {
	if c.IsFull() {
		return c.closeAndOutputAllFrames()
	}
	return c.outputReadyFrames()
}

// outputReadyFrames creates new frames as long as there's enough data ready in
// the channel out compression pipeline.
//
// This is part of an optimization to already generate frames and send them off
// as txs while still collecting blocks in the channel builder.
func (c *channelBuilder) outputReadyFrames() error {
	// TODO: Decide whether we want to fill frames to max size and use target
	// only for estimation, or use target size.
	for c.co.ReadyBytes() >= int(c.cfg.MaxFrameSize) {
		if err := c.outputFrame(); err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
	}
	return nil
}

func (c *channelBuilder) closeAndOutputAllFrames() error {
	if err := c.co.Close(); err != nil {
		return fmt.Errorf("closing channel out: %w", err)
	}

	for {
		if err := c.outputFrame(); err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
	}
}

func (c *channelBuilder) outputFrame() error {
	var buf bytes.Buffer
	fn, err := c.co.OutputFrame(&buf, c.cfg.MaxFrameSize)
	if err != io.EOF && err != nil {
		return fmt.Errorf("writing frame[%d]: %w", fn, err)
	}

	// Mark as full if max index reached
	// TODO: If there's still data in the compression pipeline of the channel out,
	// we would miss it and the whole channel would be broken because the last
	// frames would never be generated...
	// Hitting the max index is impossible with current parameters, so ignore for
	// now. Note that in order to properly catch this, we'd need to call Flush
	// after every block addition to estimate how many more frames are coming.
	if fn == math.MaxUint16 {
		c.setFullErr(ErrMaxFrameIndex)
	}

	frame := taggedData{
		id:   txID{chID: c.co.ID(), frameNumber: fn},
		data: buf.Bytes(),
	}
	c.frames = append(c.frames, frame)
	return err // possibly io.EOF (last frame)
}

// HasFrame returns whether there's any available frame. If true, it can be
// popped using NextFrame().
//
// Call OutputFrames before to create new frames from the channel out
// compression pipeline.
func (c *channelBuilder) HasFrame() bool {
	return len(c.frames) > 0
}

func (c *channelBuilder) NumFrames() int {
	return len(c.frames)
}

// NextFrame returns the next available frame.
// HasFrame must be called prior to check if there's a next frame available.
// Panics if called when there's no next frame.
func (c *channelBuilder) NextFrame() (txID, []byte) {
	if len(c.frames) == 0 {
		panic("no next frame")
	}

	f := c.frames[0]
	c.frames = c.frames[1:]
	return f.id, f.data
}

// PushFrame adds the frame back to the internal frames queue. Panics if not of
// the same channel.
func (c *channelBuilder) PushFrame(id txID, frame []byte) {
	if id.chID != c.ID() {
		panic("wrong channel")
	}
	c.frames = append(c.frames, taggedData{id: id, data: frame})
}
