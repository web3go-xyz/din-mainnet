package l1

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/sources"
	cll1 "github.com/ethereum-optimism/optimism/op-program/client/l1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

// Needs to implement the Oracle interface
var _ cll1.Oracle = (*FetchingL1Oracle)(nil)

// Want to be able to use an L1Client as the data source
var _ Source = (*sources.L1Client)(nil)

func TestHeaderByHash(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expected := &sources.HeaderInfo{}
		source := &stubSource{nextInfo: expected}
		oracle := newFetchingOracle(source)

		actual := oracle.HeaderByHash(expected.Hash())
		require.Equal(t, expected, actual)
	})

	t.Run("UnknownBlock", func(t *testing.T) {
		oracle := newFetchingOracle(&stubSource{})
		hash := common.HexToHash("0x4455")
		require.PanicsWithError(t, fmt.Errorf("unknown block: %s", hash).Error(), func() {
			oracle.HeaderByHash(hash)
		})
	})

	t.Run("Error", func(t *testing.T) {
		err := errors.New("kaboom")
		source := &stubSource{nextErr: err}
		oracle := newFetchingOracle(source)

		hash := common.HexToHash("0x8888")
		require.PanicsWithError(t, fmt.Errorf("retrieve block %s: %w", hash, err).Error(), func() {
			oracle.HeaderByHash(hash)
		})
	})
}

func TestTransactionsByHash(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expectedInfo := &sources.HeaderInfo{}
		expectedTxs := types.Transactions{
			&types.Transaction{},
		}
		source := &stubSource{nextInfo: expectedInfo, nextTxs: expectedTxs}
		oracle := newFetchingOracle(source)

		info, txs := oracle.TransactionsByHash(expectedInfo.Hash())
		require.Equal(t, expectedInfo, info)
		require.Equal(t, expectedTxs, txs)
	})

	t.Run("UnknownBlock_NoInfo", func(t *testing.T) {
		oracle := newFetchingOracle(&stubSource{})
		hash := common.HexToHash("0x4455")
		require.PanicsWithError(t, fmt.Errorf("unknown block: %s", hash).Error(), func() {
			oracle.TransactionsByHash(hash)
		})
	})

	t.Run("UnknownBlock_NoTxs", func(t *testing.T) {
		oracle := newFetchingOracle(&stubSource{nextInfo: &sources.HeaderInfo{}})
		hash := common.HexToHash("0x4455")
		require.PanicsWithError(t, fmt.Errorf("unknown block: %s", hash).Error(), func() {
			oracle.TransactionsByHash(hash)
		})
	})

	t.Run("Error", func(t *testing.T) {
		err := errors.New("kaboom")
		source := &stubSource{nextErr: err}
		oracle := newFetchingOracle(source)

		hash := common.HexToHash("0x8888")
		require.PanicsWithError(t, fmt.Errorf("retrieve transactions for block %s: %w", hash, err).Error(), func() {
			oracle.TransactionsByHash(hash)
		})
	})
}

func TestReceiptsByHash(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expectedInfo := &sources.HeaderInfo{}
		expectedRcpts := types.Receipts{
			&types.Receipt{},
		}
		source := &stubSource{nextInfo: expectedInfo, nextRcpts: expectedRcpts}
		oracle := newFetchingOracle(source)

		info, rcpts := oracle.ReceiptsByHash(expectedInfo.Hash())
		require.Equal(t, expectedInfo, info)
		require.Equal(t, expectedRcpts, rcpts)
	})

	t.Run("UnknownBlock_NoInfo", func(t *testing.T) {
		oracle := newFetchingOracle(&stubSource{})
		hash := common.HexToHash("0x4455")
		require.PanicsWithError(t, fmt.Errorf("unknown block: %s", hash).Error(), func() {
			oracle.ReceiptsByHash(hash)
		})
	})

	t.Run("UnknownBlock_NoTxs", func(t *testing.T) {
		oracle := newFetchingOracle(&stubSource{nextInfo: &sources.HeaderInfo{}})
		hash := common.HexToHash("0x4455")
		require.PanicsWithError(t, fmt.Errorf("unknown block: %s", hash).Error(), func() {
			oracle.ReceiptsByHash(hash)
		})
	})

	t.Run("Error", func(t *testing.T) {
		err := errors.New("kaboom")
		source := &stubSource{nextErr: err}
		oracle := newFetchingOracle(source)

		hash := common.HexToHash("0x8888")
		require.PanicsWithError(t, fmt.Errorf("retrieve receipts for block %s: %w", hash, err).Error(), func() {
			oracle.ReceiptsByHash(hash)
		})
	})
}

func newFetchingOracle(source Source) *FetchingL1Oracle {
	return &FetchingL1Oracle{
		ctx:    context.Background(),
		source: source,
	}
}

type stubSource struct {
	nextInfo  eth.BlockInfo
	nextTxs   types.Transactions
	nextRcpts types.Receipts
	nextErr   error
}

func (s stubSource) InfoByHash(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, error) {
	return s.nextInfo, s.nextErr
}

func (s stubSource) InfoAndTxsByHash(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, types.Transactions, error) {
	return s.nextInfo, s.nextTxs, s.nextErr
}

func (s stubSource) FetchReceipts(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, types.Receipts, error) {
	return s.nextInfo, s.nextRcpts, s.nextErr
}
