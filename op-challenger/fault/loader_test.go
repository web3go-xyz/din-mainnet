package fault

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/testlog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

var (
	mockClaimDataError = fmt.Errorf("claim data errored")
	mockClaimLenError  = fmt.Errorf("claim len errored")
)

type mockClaimFetcher struct {
	claimDataError bool
	claimLenError  bool
	currentIndex   uint64
	returnClaims   []struct {
		ParentIndex uint32
		Countered   bool
		Claim       [32]byte
		Position    *big.Int
		Clock       *big.Int
	}
}

func newMockClaimFetcher() *mockClaimFetcher {
	return &mockClaimFetcher{
		returnClaims: []struct {
			ParentIndex uint32
			Countered   bool
			Claim       [32]byte
			Position    *big.Int
			Clock       *big.Int
		}{
			{
				Claim:    [32]byte{0x00},
				Position: big.NewInt(0),
			},
			{
				Claim:    [32]byte{0x01},
				Position: big.NewInt(0),
			},
			{
				Claim:    [32]byte{0x02},
				Position: big.NewInt(0),
			},
		},
	}
}

func (m *mockClaimFetcher) ClaimData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ParentIndex uint32
	Countered   bool
	Claim       [32]byte
	Position    *big.Int
	Clock       *big.Int
}, error) {
	if m.claimDataError {
		return struct {
			ParentIndex uint32
			Countered   bool
			Claim       [32]byte
			Position    *big.Int
			Clock       *big.Int
		}{}, mockClaimDataError
	}
	returnClaim := m.returnClaims[m.currentIndex]
	m.currentIndex++
	return returnClaim, nil
}

func (m *mockClaimFetcher) ClaimDataLen(opts *bind.CallOpts) (*big.Int, error) {
	if m.claimLenError {
		return big.NewInt(0), mockClaimLenError
	}
	return big.NewInt(int64(len(m.returnClaims))), nil
}

// TestLoader_FetchClaims_Succeeds tests [loader.FetchClaims].
func TestLoader_FetchClaims_Succeeds(t *testing.T) {
	log := testlog.Logger(t, log.LvlError)
	mockClaimFetcher := newMockClaimFetcher()
	expectedClaims := mockClaimFetcher.returnClaims
	loader := NewLoader(log, mockClaimFetcher)
	claims, err := loader.FetchClaims(context.Background())
	require.NoError(t, err)
	require.ElementsMatch(t, []Claim{
		{
			ClaimData: ClaimData{
				Value:    expectedClaims[0].Claim,
				Position: NewPositionFromGIndex(expectedClaims[0].Position.Uint64()),
			},
			Parent: ClaimData{
				Value:    expectedClaims[0].Claim,
				Position: NewPositionFromGIndex(expectedClaims[0].Position.Uint64()),
			},
			ContractIndex: 0,
		},
		{
			ClaimData: ClaimData{
				Value:    expectedClaims[1].Claim,
				Position: NewPositionFromGIndex(expectedClaims[1].Position.Uint64()),
			},
			Parent: ClaimData{
				Value:    expectedClaims[0].Claim,
				Position: NewPositionFromGIndex(expectedClaims[1].Position.Uint64()),
			},
			ContractIndex: 1,
		},
		{
			ClaimData: ClaimData{
				Value:    expectedClaims[2].Claim,
				Position: NewPositionFromGIndex(expectedClaims[2].Position.Uint64()),
			},
			Parent: ClaimData{
				Value:    expectedClaims[0].Claim,
				Position: NewPositionFromGIndex(expectedClaims[2].Position.Uint64()),
			},
			ContractIndex: 2,
		},
	}, claims)
}

// TestLoader_FetchClaims_ClaimDataErrors tests [loader.FetchClaims]
// when the claim fetcher [ClaimData] function call errors.
func TestLoader_FetchClaims_ClaimDataErrors(t *testing.T) {
	log := testlog.Logger(t, log.LvlError)
	mockClaimFetcher := newMockClaimFetcher()
	mockClaimFetcher.claimDataError = true
	loader := NewLoader(log, mockClaimFetcher)
	claims, err := loader.FetchClaims(context.Background())
	require.ErrorIs(t, err, mockClaimDataError)
	require.Empty(t, claims)
}

// TestLoader_FetchClaims_ClaimLenErrors tests [loader.FetchClaims]
// when the claim fetcher [ClaimDataLen] function call errors.
func TestLoader_FetchClaims_ClaimLenErrors(t *testing.T) {
	log := testlog.Logger(t, log.LvlError)
	mockClaimFetcher := newMockClaimFetcher()
	mockClaimFetcher.claimLenError = true
	loader := NewLoader(log, mockClaimFetcher)
	claims, err := loader.FetchClaims(context.Background())
	require.ErrorIs(t, err, mockClaimLenError)
	require.Empty(t, claims)
}
