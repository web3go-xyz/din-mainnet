package node

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/core/types"
)

// make a set of headers which chain correctly
func makeHeaders(numHeaders uint64, prevHeader *types.Header) []types.Header {
	headers := make([]types.Header, numHeaders)
	for i := range headers {
		if i == 0 {
			if prevHeader == nil {
				// genesis
				headers[i] = types.Header{Number: big.NewInt(0)}
			} else {
				// chain onto the previous header
				headers[i] = types.Header{Number: big.NewInt(prevHeader.Number.Int64() + 1)}
				headers[i].ParentHash = prevHeader.Hash()
			}
		} else {
			headers[i] = types.Header{Number: big.NewInt(headers[i-1].Number.Int64() + 1)}
			headers[i].ParentHash = headers[i-1].Hash()
		}
	}

	return headers
}

func TestHeaderTraversalNextFinalizedHeadersNoOp(t *testing.T) {
	client := new(MockEthClient)

	// start from block 10 as the latest fetched block
	lastHeader := &types.Header{Number: big.NewInt(10)}
	headerTraversal := NewHeaderTraversal(client, lastHeader, bigZero)

	// no new headers when matched with head
	client.On("BlockHeaderByNumber", (*big.Int)(nil)).Return(lastHeader, nil)
	headers, err := headerTraversal.NextFinalizedHeaders(100)
	require.NoError(t, err)
	require.Empty(t, headers)
}

func TestHeaderTraversalNextFinalizedHeadersCursored(t *testing.T) {
	client := new(MockEthClient)

	// start from genesis
	headerTraversal := NewHeaderTraversal(client, nil, bigZero)

	// blocks [0..4]
	headers := makeHeaders(5, nil)
	client.On("BlockHeaderByNumber", (*big.Int)(nil)).Return(&headers[4], nil).Times(1) // Times so that we can override next
	client.On("BlockHeadersByRange", mock.MatchedBy(BigIntMatcher(0)), mock.MatchedBy(BigIntMatcher(4))).Return(headers, nil)
	headers, err := headerTraversal.NextFinalizedHeaders(5)
	require.NoError(t, err)
	require.Len(t, headers, 5)

	// blocks [5..9]
	headers = makeHeaders(5, &headers[len(headers)-1])
	client.On("BlockHeaderByNumber", (*big.Int)(nil)).Return(&headers[4], nil)
	client.On("BlockHeadersByRange", mock.MatchedBy(BigIntMatcher(5)), mock.MatchedBy(BigIntMatcher(9))).Return(headers, nil)
	headers, err = headerTraversal.NextFinalizedHeaders(5)
	require.NoError(t, err)
	require.Len(t, headers, 5)
}

func TestHeaderTraversalNextFinalizedHeadersMaxSize(t *testing.T) {
	client := new(MockEthClient)

	// start from genesis
	headerTraversal := NewHeaderTraversal(client, nil, bigZero)

	// 100 "available" headers
	client.On("BlockHeaderByNumber", (*big.Int)(nil)).Return(&types.Header{Number: big.NewInt(100)}, nil)

	// clamped by the supplied size
	headers := makeHeaders(5, nil)
	client.On("BlockHeadersByRange", mock.MatchedBy(BigIntMatcher(0)), mock.MatchedBy(BigIntMatcher(4))).Return(headers, nil)
	headers, err := headerTraversal.NextFinalizedHeaders(5)
	require.NoError(t, err)
	require.Len(t, headers, 5)

	// clamped by the supplied size. FinalizedHeight == 100
	headers = makeHeaders(10, &headers[len(headers)-1])
	client.On("BlockHeadersByRange", mock.MatchedBy(BigIntMatcher(5)), mock.MatchedBy(BigIntMatcher(14))).Return(headers, nil)
	headers, err = headerTraversal.NextFinalizedHeaders(10)
	require.NoError(t, err)
	require.Len(t, headers, 10)
}

func TestHeaderTraversalMismatchedProviderStateError(t *testing.T) {
	client := new(MockEthClient)

	// start from genesis
	headerTraversal := NewHeaderTraversal(client, nil, bigZero)

	// blocks [0..4]
	headers := makeHeaders(5, nil)
	client.On("BlockHeaderByNumber", (*big.Int)(nil)).Return(&headers[4], nil).Times(1) // Times so that we can override next
	client.On("BlockHeadersByRange", mock.MatchedBy(BigIntMatcher(0)), mock.MatchedBy(BigIntMatcher(4))).Return(headers, nil)
	headers, err := headerTraversal.NextFinalizedHeaders(5)
	require.NoError(t, err)
	require.Len(t, headers, 5)

	// blocks [5..9]. Next batch is not chained correctly (starts again from genesis)
	headers = makeHeaders(5, nil)
	client.On("BlockHeaderByNumber", (*big.Int)(nil)).Return(&types.Header{Number: big.NewInt(9)}, nil)
	client.On("BlockHeadersByRange", mock.MatchedBy(BigIntMatcher(5)), mock.MatchedBy(BigIntMatcher(9))).Return(headers, nil)
	headers, err = headerTraversal.NextFinalizedHeaders(5)
	require.Nil(t, headers)
	require.Equal(t, ErrHeaderTraversalAndProviderMismatchedState, err)
}
