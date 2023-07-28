package alphabet

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum-optimism/optimism/op-challenger/fault/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrIndexTooLarge = errors.New("index is larger than the maximum index")
)

// AlphabetProvider is a [TraceProvider] that provides claims for specific
// indices in the given trace.
type AlphabetProvider struct {
	state  []string
	maxLen uint64
}

// NewAlphabetProvider returns a new [AlphabetProvider].
func NewAlphabetProvider(state string, depth uint64) *AlphabetProvider {
	return &AlphabetProvider{
		state:  strings.Split(state, ""),
		maxLen: uint64(1 << depth),
	}
}

// GetOracleData should not return any preimage oracle data for the alphabet provider.
func (p *AlphabetProvider) GetOracleData(ctx context.Context, i uint64) (*types.PreimageOracleData, error) {
	return &types.PreimageOracleData{}, nil
}

// GetPreimage returns the preimage for the given hash.
func (ap *AlphabetProvider) GetPreimage(ctx context.Context, i uint64) ([]byte, []byte, error) {
	// The index cannot be larger than the maximum index as computed by the depth.
	if i >= ap.maxLen {
		return nil, nil, ErrIndexTooLarge
	}
	// We extend the deepest hash to the maximum depth if the trace is not expansive.
	if i >= uint64(len(ap.state)) {
		return ap.GetPreimage(ctx, uint64(len(ap.state))-1)
	}
	return BuildAlphabetPreimage(i, ap.state[i]), []byte{}, nil
}

// Get returns the claim value at the given index in the trace.
func (ap *AlphabetProvider) Get(ctx context.Context, i uint64) (common.Hash, error) {
	claimBytes, _, err := ap.GetPreimage(ctx, i)
	if err != nil {
		return common.Hash{}, err
	}
	return crypto.Keccak256Hash(claimBytes), nil
}

func (ap *AlphabetProvider) AbsolutePreState(ctx context.Context) []byte {
	return common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000060")
}

// BuildAlphabetPreimage constructs the claim bytes for the index and state item.
func BuildAlphabetPreimage(i uint64, letter string) []byte {
	return append(IndexToBytes(i), LetterToBytes(letter)...)
}

// IndexToBytes converts an index to a byte slice big endian
func IndexToBytes(i uint64) []byte {
	big := new(big.Int)
	big.SetUint64(i)
	out := make([]byte, 32)
	return big.FillBytes(out)
}

// LetterToBytes converts a letter to a 32 byte array
func LetterToBytes(letter string) []byte {
	out := make([]byte, 32)
	out[31] = byte(letter[0])
	return out
}
