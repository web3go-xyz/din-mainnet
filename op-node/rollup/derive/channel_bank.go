package derive

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type NextDataProvider interface {
	NextData(ctx context.Context) ([]byte, error)
	Origin() eth.L1BlockRef
}

// ChannelBank is a stateful stage that does the following:
// 1. Unmarshalls frames from L1 transaction data
// 2. Applies those frames to a channel
// 3. Attempts to read from the channel when it is ready
// 4. Prunes channels (not frames) when the channel bank is too large.
//
// Note: we prune before we ingest data.
// As we switch between ingesting data & reading, the prune step occurs at an odd point
// Specifically, the channel bank is not allowed to become too large between successive calls
// to `IngestData`. This means that we can do an ingest and then do a read while becoming too large.

// ChannelBank buffers channel frames, and emits full channel data
type ChannelBank struct {
	log log.Logger
	cfg *rollup.Config

	channels     map[ChannelID]*Channel // channels by ID
	channelQueue []ChannelID            // channels in FIFO order

	origin eth.L1BlockRef

	prev    NextDataProvider
	fetcher L1Fetcher
}

var _ PullStage = (*ChannelBank)(nil)

// NewChannelBank creates a ChannelBank, which should be Reset(origin) before use.
func NewChannelBank(log log.Logger, cfg *rollup.Config, prev NextDataProvider, fetcher L1Fetcher) *ChannelBank {
	return &ChannelBank{
		log:          log,
		cfg:          cfg,
		channels:     make(map[ChannelID]*Channel),
		channelQueue: make([]ChannelID, 0, 10),
		prev:         prev,
		fetcher:      fetcher,
	}
}

func (cb *ChannelBank) Origin() eth.L1BlockRef {
	return cb.prev.Origin()
}

func (cb *ChannelBank) prune() {
	// check total size
	totalSize := uint64(0)
	for _, ch := range cb.channels {
		totalSize += ch.size
	}
	// prune until it is reasonable again. The high-priority channel failed to be read, so we start pruning there.
	for totalSize > MaxChannelBankSize {
		id := cb.channelQueue[0]
		ch := cb.channels[id]
		cb.channelQueue = cb.channelQueue[1:]
		delete(cb.channels, id)
		totalSize -= ch.size
	}
}

// IngestData adds new L1 data to the channel bank.
// Read() should be called repeatedly first, until everything has been read, before adding new data.\
func (cb *ChannelBank) IngestData(data []byte) {
	cb.log.Debug("channel bank got new data", "origin", cb.origin, "data_len", len(data))

	// TODO: Why is the prune here?
	cb.prune()

	frames, err := ParseFrames(data)
	if err != nil {
		cb.log.Warn("malformed frame", "err", err)
		return
	}

	// Process each frame
	for _, f := range frames {
		currentCh, ok := cb.channels[f.ID]
		if !ok {
			// create new channel if it doesn't exist yet
			currentCh = NewChannel(f.ID, cb.origin)
			cb.channels[f.ID] = currentCh
			cb.channelQueue = append(cb.channelQueue, f.ID)
		}

		// check if the channel is not timed out
		if currentCh.OpenBlockNumber()+cb.cfg.ChannelTimeout < cb.origin.Number {
			cb.log.Warn("channel is timed out, ignore frame", "channel", f.ID, "frame", f.FrameNumber)
			continue
		}

		cb.log.Trace("ingesting frame", "channel", f.ID, "frame_number", f.FrameNumber, "length", len(f.Data))
		if err := currentCh.AddFrame(f, cb.origin); err != nil {
			cb.log.Warn("failed to ingest frame into channel", "channel", f.ID, "frame_number", f.FrameNumber, "err", err)
			continue
		}
	}
}

// Read the raw data of the first channel, if it's timed-out or closed.
// Read returns io.EOF if there is nothing new to read.
func (cb *ChannelBank) Read() (data []byte, err error) {
	if len(cb.channelQueue) == 0 {
		return nil, io.EOF
	}
	first := cb.channelQueue[0]
	ch := cb.channels[first]
	timedOut := ch.OpenBlockNumber()+cb.cfg.ChannelTimeout < cb.origin.Number
	if timedOut {
		cb.log.Debug("channel timed out", "channel", first, "frames", len(ch.inputs))
		delete(cb.channels, first)
		cb.channelQueue = cb.channelQueue[1:]
		return nil, io.EOF
	}
	if !ch.IsReady() {
		return nil, io.EOF
	}

	delete(cb.channels, first)
	cb.channelQueue = cb.channelQueue[1:]
	r := ch.Reader()
	// Suprress error here. io.ReadAll does return nil instead of io.EOF though.
	data, _ = io.ReadAll(r)
	return data, nil
}

// NextData pulls the next piece of data from the channel bank.
// Note that it attempts to pull data out of the channel bank prior to
// loading data in (unlike most other stages). This is to ensure maintain
// consistency around channel bank pruning which depends upon the order
// of operations.
func (cb *ChannelBank) NextData(ctx context.Context) ([]byte, error) {
	if cb.origin != cb.prev.Origin() {
		cb.origin = cb.prev.Origin()
	}

	// Do the read from the channel bank first
	data, err := cb.Read()
	if err == io.EOF {
		// continue - We will attempt to load data into the channel bank
	} else if err != nil {
		return nil, err
	} else {
		return data, nil
	}

	// Then load data into the channel bank
	if data, err := cb.prev.NextData(ctx); err == io.EOF {
		return nil, io.EOF
	} else if err != nil {
		return nil, err
	} else {
		cb.IngestData(data)
		return nil, NewTemporaryError(errors.New("not enough data"))
	}
}

// ResetStep walks back the L1 chain, starting at the origin of the next stage,
// to find the origin that the channel bank should be reset to,
// to get consistent reads starting at origin.
// Any channel data before this origin will be timed out by the time the channel bank is synced up to the origin,
// so it is not relevant to replay it into the bank.
func (cb *ChannelBank) Reset(ctx context.Context, base eth.L1BlockRef) error {
	cb.log.Debug("walking back to find reset origin for channel bank", "origin", base)
	// go back in history if we are not distant enough from the next stage
	resetBlock := base.Number - cb.cfg.ChannelTimeout
	if base.Number < cb.cfg.ChannelTimeout {
		resetBlock = 0 // don't underflow
	}
	parent, err := cb.fetcher.L1BlockRefByNumber(ctx, resetBlock)
	if err != nil {
		return NewTemporaryError(fmt.Errorf("failed to find channel bank block, failed to retrieve L1 reference: %w", err))
	}
	cb.origin = parent
	return io.EOF
}

type L1BlockRefByHashFetcher interface {
	L1BlockRefByHash(context.Context, common.Hash) (eth.L1BlockRef, error)
}
