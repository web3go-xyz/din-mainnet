package alphabet

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	preimage "github.com/ethereum-optimism/optimism/op-preimage"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func alphabetClaim(index *big.Int, claim *big.Int) common.Hash {
	return alphabetStateHash(BuildAlphabetPreimage(index, claim))
}

func TestAlphabetProvider_Prestate(t *testing.T) {
	depth := types.Depth(4)
	startingL2BlockNumber := big.NewInt(2)

	// Actual preimage values generated by the solidity AlphabetVM at each step.
	expectedPrestates := []string{
		"0000000000000000000000000000000000000000000000000000000000000060",
		"00000000000000000000000000000000000000000000000000000000000000210000000000000000000000000000000000000000000000000000000000000081",
		"00000000000000000000000000000000000000000000000000000000000000220000000000000000000000000000000000000000000000000000000000000082",
		"00000000000000000000000000000000000000000000000000000000000000230000000000000000000000000000000000000000000000000000000000000083",
		"00000000000000000000000000000000000000000000000000000000000000240000000000000000000000000000000000000000000000000000000000000084",
		"00000000000000000000000000000000000000000000000000000000000000250000000000000000000000000000000000000000000000000000000000000085",
	}

	ap := NewTraceProvider(startingL2BlockNumber, depth)

	for i, expected := range expectedPrestates {
		i, expected := i, expected
		t.Run(fmt.Sprintf("Step_%v", i), func(t *testing.T) {
			result, _, _, err := ap.GetStepData(context.Background(), types.NewPosition(4, big.NewInt(int64(i))))
			require.NoError(t, err)
			require.Equalf(t, expected, common.Bytes2Hex(result), "Incorrect prestate at trace index %v", i)
		})
	}
}

// TestAlphabetProvider_Get_ClaimsByTraceIndex tests the [fault.AlphabetProvider] Get function.
func TestAlphabetProvider_Get_ClaimsByTraceIndex(t *testing.T) {
	// Create a new alphabet provider.
	depth := types.Depth(3)
	startingL2BlockNumber := big.NewInt(1)
	sbn := new(big.Int).Lsh(startingL2BlockNumber, 4)
	startingTraceIndex := new(big.Int).Add(absolutePrestateInt, sbn)
	canonicalProvider := NewTraceProvider(startingL2BlockNumber, depth)

	// Build a list of traces.
	traces := []struct {
		traceIndex   types.Position
		expectedHash common.Hash
	}{
		{
			types.NewPosition(depth, big.NewInt(7)),
			alphabetClaim(new(big.Int).Add(sbn, big.NewInt(8)), new(big.Int).Add(startingTraceIndex, big.NewInt(8))),
		},
		{
			types.NewPosition(depth, big.NewInt(3)),
			alphabetClaim(new(big.Int).Add(sbn, big.NewInt(4)), new(big.Int).Add(startingTraceIndex, big.NewInt(4))),
		},
		{
			types.NewPosition(depth, big.NewInt(5)),
			alphabetClaim(new(big.Int).Add(sbn, big.NewInt(6)), new(big.Int).Add(startingTraceIndex, big.NewInt(6))),
		},
	}

	// Execute each trace and check the alphabet provider returns the expected hash.
	for i, trace := range traces {
		expectedHash, err := canonicalProvider.Get(context.Background(), trace.traceIndex)
		require.NoError(t, err)
		require.Equalf(t, trace.expectedHash, expectedHash, "Trace %v", i)
	}
}

// TestGetPreimage_Succeeds tests the GetPreimage function
// returns the correct pre-image for a index.
func TestGetStepData_Succeeds(t *testing.T) {
	depth := types.Depth(2)
	startingL2BlockNumber := big.NewInt(1)
	ap := NewTraceProvider(startingL2BlockNumber, depth)
	expected := absolutePrestate
	pos := types.NewPosition(depth, big.NewInt(0))
	retrieved, proof, data, err := ap.GetStepData(context.Background(), pos)
	require.NoError(t, err)
	require.Equal(t, expected, retrieved)
	require.Empty(t, proof)
	key := preimage.LocalIndexKey(L2ClaimBlockNumberLocalIndex).PreimageKey()
	expectedLocalContextData := types.NewPreimageOracleData(key[:], startingL2BlockNumber.Bytes(), 0)
	require.Equal(t, expectedLocalContextData, data)
}

// TestGetPreimage_TooLargeIndex_Fails tests the GetPreimage
// function errors if the index is too large.
func TestGetStepData_TooLargeIndex_Fails(t *testing.T) {
	depth := types.Depth(2)
	startingL2BlockNumber := big.NewInt(1)
	ap := NewTraceProvider(startingL2BlockNumber, depth)
	pos := types.NewPosition(depth, big.NewInt(5))
	_, _, _, err := ap.GetStepData(context.Background(), pos)
	require.ErrorIs(t, err, ErrIndexTooLarge)
}

// TestGet_Succeeds tests the Get function.
func TestGet_Succeeds(t *testing.T) {
	depth := types.Depth(2)
	startingL2BlockNumber := big.NewInt(1)
	ap := NewTraceProvider(startingL2BlockNumber, depth)
	pos := types.NewPosition(depth, big.NewInt(0))
	claim, err := ap.Get(context.Background(), pos)
	require.NoError(t, err)
	expected := alphabetClaim(big.NewInt(17), big.NewInt(113))
	require.Equal(t, expected, claim)
}

// TestGet_IndexTooLarge tests the Get function with an index
// greater than the number of indices: 2^depth - 1.
func TestGet_IndexTooLarge(t *testing.T) {
	depth := types.Depth(2)
	startingL2BlockNumber := big.NewInt(1)
	ap := NewTraceProvider(startingL2BlockNumber, depth)
	pos := types.NewPosition(depth, big.NewInt(4))
	_, err := ap.Get(context.Background(), pos)
	require.ErrorIs(t, err, ErrIndexTooLarge)
}

func TestGet_DepthTooLarge(t *testing.T) {
	depth := types.Depth(2)
	startingL2BlockNumber := big.NewInt(1)
	ap := NewTraceProvider(startingL2BlockNumber, depth)
	pos := types.NewPosition(depth+1, big.NewInt(0))
	_, err := ap.Get(context.Background(), pos)
	require.ErrorIs(t, err, ErrIndexTooLarge)
}
