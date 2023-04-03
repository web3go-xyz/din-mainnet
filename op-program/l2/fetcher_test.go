package l2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/testutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

type callContextRequest struct {
	ctx    context.Context
	method string
	args   []interface{}
}
type stubCallContext struct {
	nextResult any
	nextErr    error
	requests   []callContextRequest
}

func (c *stubCallContext) CallContext(ctx context.Context, result any, method string, args ...interface{}) error {
	if result != nil && reflect.TypeOf(result).Kind() != reflect.Ptr {
		return fmt.Errorf("call result parameter must be pointer or nil interface: %v", result)
	}
	c.requests = append(c.requests, callContextRequest{ctx: ctx, method: method, args: args})
	if c.nextErr != nil {
		return c.nextErr
	}
	res, err := json.Marshal(c.nextResult)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}
	err = json.Unmarshal(res, result)
	if err != nil {
		return fmt.Errorf("json unmarshal: %w", err)
	}
	return nil
}

func TestNodeByHash(t *testing.T) {
	rng := rand.New(rand.NewSource(1234))
	hash := testutils.RandomHash(rng)
	ctx := context.Background()

	t.Run("Error", func(t *testing.T) {
		stub := &stubCallContext{
			nextErr: errors.New("oops"),
		}
		fetcher := newFetcher(nil, stub)

		node, err := fetcher.NodeByHash(ctx, hash)
		require.ErrorIs(t, err, stub.nextErr)
		require.Nil(t, node)
	})

	t.Run("Success", func(t *testing.T) {
		expected := (hexutil.Bytes)([]byte{12, 34})
		stub := &stubCallContext{
			nextResult: expected,
		}
		fetcher := newFetcher(nil, stub)

		node, err := fetcher.NodeByHash(ctx, hash)
		require.NoError(t, err)
		require.EqualValues(t, expected, node)
	})

	t.Run("RequestArgs", func(t *testing.T) {
		stub := &stubCallContext{
			nextResult: (hexutil.Bytes)([]byte{12, 34}),
		}
		fetcher := newFetcher(nil, stub)

		_, _ = fetcher.NodeByHash(ctx, hash)
		require.Len(t, stub.requests, 1, "should make single request")
		req := stub.requests[0]
		require.Equal(t, "debug_dbGet", req.method)
		require.Equal(t, []interface{}{hash.Hex()}, req.args)
	})
}

type blockRequest struct {
	ctx       context.Context
	blockHash common.Hash
}

type stubBlockSource struct {
	requests   []blockRequest
	nextErr    error
	nextResult *types.Block
}

func (s *stubBlockSource) BlockByHash(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	s.requests = append(s.requests, blockRequest{
		ctx:       ctx,
		blockHash: blockHash,
	})
	return s.nextResult, s.nextErr
}

func TestBlockByHash(t *testing.T) {
	rng := rand.New(rand.NewSource(1234))
	hash := testutils.RandomHash(rng)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		block, _ := testutils.RandomBlock(rng, 1)
		stub := &stubBlockSource{nextResult: block}
		fetcher := newFetcher(stub, nil)

		res, err := fetcher.BlockByHash(ctx, hash)
		require.NoError(t, err)
		require.Same(t, block, res)
	})

	t.Run("Error", func(t *testing.T) {
		stub := &stubBlockSource{nextErr: errors.New("boom")}
		fetcher := newFetcher(stub, nil)

		res, err := fetcher.BlockByHash(ctx, hash)
		require.ErrorIs(t, err, stub.nextErr)
		require.Nil(t, res)
	})

	t.Run("RequestArgs", func(t *testing.T) {
		stub := &stubBlockSource{}
		fetcher := newFetcher(stub, nil)

		_, _ = fetcher.BlockByHash(ctx, hash)

		require.Len(t, stub.requests, 1, "should make single request")
		req := stub.requests[0]
		require.Equal(t, hash, req.blockHash)
	})
}

func newFetcher(blockSource BlockSource, callContext CallContext) *FetchingL2Oracle {
	return &FetchingL2Oracle{
		logger:      log.New(),
		blockSource: blockSource,
		callContext: callContext,
	}
}
