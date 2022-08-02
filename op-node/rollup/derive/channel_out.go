package derive

import (
	"bytes"
	"compress/zlib"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type ChannelOut struct {
	id ChannelID
	// Frame ID of the next frame to emit. Increment after emitting
	frame uint64
	// How much we've pulled from the reader so far
	offset uint64
	// scratch for temporary buffering
	scratch bytes.Buffer

	// Compressor stage. Write input data to it
	compress *zlib.Writer
	// post compression buffer
	buf bytes.Buffer

	closed bool
}

func (co *ChannelOut) ID() ChannelID {
	return co.id
}

func NewChannelOut(channelTime uint64) (*ChannelOut, error) {
	c := &ChannelOut{
		id: ChannelID{
			Time: channelTime,
		},
		frame:  0,
		offset: 0,
	}
	_, err := rand.Read(c.id.Data[:])
	if err != nil {
		return nil, err
	}

	compress, err := zlib.NewWriterLevel(&c.buf, zlib.BestCompression)
	if err != nil {
		return nil, err
	}
	c.compress = compress

	return c, nil
}

// TODO: reuse ChannelOut for performance
func (co *ChannelOut) Reset(channelTime uint64) error {
	co.frame = 0
	co.offset = 0
	co.buf.Reset()
	co.scratch.Reset()
	co.compress.Reset(&co.buf)
	co.closed = false
	co.id.Time = channelTime
	_, err := rand.Read(co.id.Data[:])
	if err != nil {
		return err
	}
	return nil
}

func (co *ChannelOut) AddBlock(block *types.Block) error {
	if co.closed {
		return errors.New("already closed")
	}
	return blockToBatch(block, co.compress)
}

func makeUVarint(x uint64) []byte {
	var tmp [binary.MaxVarintLen64]byte
	n := binary.PutUvarint(tmp[:], x)
	return tmp[:n]
}

// ReadyBytes returns the number of bytes that the channel out can immediately output into a frame.
// Use `Flush` or `Close` to move data from the compression buffer into the ready buffer if more bytes
// are needed. Add blocks may add to the ready buffer, but it is not guaranteed due to the compression stage.
func (co *ChannelOut) ReadyBytes() int {
	return co.buf.Len()
}

// Flush flushes the internal compression stage to the ready buffer. It enables pulling a larger & more
// complete frame. It reduces the compression efficiency.
func (co *ChannelOut) Flush() error {
	return co.compress.Flush()
}

func (co *ChannelOut) Close() error {
	if co.closed {
		return errors.New("already closed")
	}
	co.closed = true
	return co.compress.Close()
}

// OutputFrame writes a frame to w with a given max size
// Use `ReadyBytes`, `Flush`, and `Close` to modify the ready buffer.
// Returns io.EOF when the channel is closed & there are no more frames
// Returns nil if there is still more buffered data.
// Returns and error if it ran into an error during processing.
func (co *ChannelOut) OutputFrame(w *bytes.Buffer, maxSize uint64) error {
	f := Frame{
		ID:          co.id,
		FrameNumber: co.frame,
	}

	// Copy data from the local buffer into the frame data buffer
	// Don't go past the maxSize even with the max possible uvarints
	// +1 for single byte of frame content, +1 for lastFrame bool
	// +24 for maximum uvarints
	// +32 for the data ID
	maxDataSize := maxSize - 32 - 24 - 1 - 1
	if maxDataSize >= uint64(co.buf.Len()) {
		maxDataSize = uint64(co.buf.Len())
		// If we are closed & will not spill past the current frame, end it.
		if co.closed {
			f.IsLast = true
		}
	}
	f.Data = make([]byte, maxDataSize)

	if _, err := io.ReadFull(&co.buf, f.Data); err != nil {
		return err
	}

	if _, err := f.MarshalBinary(w); err != nil {
		return err
	}

	co.frame += 1
	if f.IsLast {
		return io.EOF
	} else {
		return nil
	}
}

// blockToBatch writes the raw block bytes (after batch encoding) to the writer
func blockToBatch(block *types.Block, w io.Writer) error {
	var opaqueTxs []hexutil.Bytes
	for _, tx := range block.Transactions() {
		if tx.Type() == types.DepositTxType {
			continue
		}
		otx, err := tx.MarshalBinary()
		if err != nil {
			return err // TODO: wrap err
		}
		opaqueTxs = append(opaqueTxs, otx)
	}
	l1InfoTx := block.Transactions()[0]
	l1Info, err := L1InfoDepositTxData(l1InfoTx.Data())
	if err != nil {
		return err // TODO: wrap err
	}

	batch := &BatchData{BatchV1{
		EpochNum:     rollup.Epoch(l1Info.Number),
		EpochHash:    l1Info.BlockHash,
		Timestamp:    block.Time(),
		Transactions: opaqueTxs,
	},
	}
	return rlp.Encode(w, batch)
}
