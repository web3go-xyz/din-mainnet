package prefetcher

import (
	"context"
	"math"

	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-service/backoff"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

const maxAttempts = math.MaxInt // Succeed or die trying

type RetryingL1Source struct {
	logger   log.Logger
	source   L1Source
	strategy backoff.Strategy
}

func NewRetryingL1Source(logger log.Logger, source L1Source) *RetryingL1Source {
	return &RetryingL1Source{
		logger:   logger,
		source:   source,
		strategy: backoff.Exponential(),
	}
}

func (s *RetryingL1Source) InfoByHash(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, error) {
	return backoff.Do(ctx, maxAttempts, s.strategy, func() (eth.BlockInfo, error) {
		res, err := s.source.InfoByHash(ctx, blockHash)
		if err != nil {
			s.logger.Warn("Failed to retrieve info", "hash", blockHash, "err", err)
		}
		return res, err
	})
}

func (s *RetryingL1Source) InfoAndTxsByHash(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, types.Transactions, error) {
	return backoff.Do2(ctx, maxAttempts, s.strategy, func() (eth.BlockInfo, types.Transactions, error) {
		i, t, err := s.source.InfoAndTxsByHash(ctx, blockHash)
		if err != nil {
			s.logger.Warn("Failed to retrieve l1 info and txs", "hash", blockHash, "err", err)
		}
		return i, t, err
	})
}

func (s *RetryingL1Source) FetchReceipts(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, types.Receipts, error) {
	return backoff.Do2(ctx, maxAttempts, s.strategy, func() (eth.BlockInfo, types.Receipts, error) {
		i, r, err := s.source.FetchReceipts(ctx, blockHash)
		if err != nil {
			s.logger.Warn("Failed to fetch receipts", "hash", blockHash, "err", err)
		}
		return i, r, err
	})
}

var _ L1Source = (*RetryingL1Source)(nil)

type RetryingL2Source struct {
	logger   log.Logger
	source   L2Source
	strategy backoff.Strategy
}

func (s *RetryingL2Source) InfoAndTxsByHash(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, types.Transactions, error) {
	return backoff.Do2(ctx, maxAttempts, s.strategy, func() (eth.BlockInfo, types.Transactions, error) {
		i, t, err := s.source.InfoAndTxsByHash(ctx, blockHash)
		if err != nil {
			s.logger.Warn("Failed to retrieve l2 info and txs", "hash", blockHash, "err", err)
		}
		return i, t, err
	})
}

func (s *RetryingL2Source) NodeByHash(ctx context.Context, hash common.Hash) ([]byte, error) {
	return backoff.Do(ctx, maxAttempts, s.strategy, func() ([]byte, error) {
		n, err := s.source.NodeByHash(ctx, hash)
		if err != nil {
			s.logger.Warn("Failed to retrieve node", "hash", hash, "err", err)
		}
		return n, err
	})
}

func (s *RetryingL2Source) CodeByHash(ctx context.Context, hash common.Hash) ([]byte, error) {
	return backoff.Do(ctx, maxAttempts, s.strategy, func() ([]byte, error) {
		c, err := s.source.CodeByHash(ctx, hash)
		if err != nil {
			s.logger.Warn("Failed to retrieve code", "hash", hash, "err", err)
		}
		return c, err
	})
}

func (s *RetryingL2Source) OutputByRoot(ctx context.Context, root common.Hash) (eth.Output, error) {
	var output eth.Output
	err := backoff.DoCtx(ctx, maxAttempts, s.strategy, func() error {
		o, err := s.source.OutputByRoot(ctx, root)
		if err != nil {
			s.logger.Warn("Failed to fetch l2 output", "root", root, "err", err)
			return err
		}
		output = o
		return nil
	})
	return output, err
}

func NewRetryingL2Source(logger log.Logger, source L2Source) *RetryingL2Source {
	return &RetryingL2Source{
		logger:   logger,
		source:   source,
		strategy: backoff.Exponential(),
	}
}

var _ L2Source = (*RetryingL2Source)(nil)
