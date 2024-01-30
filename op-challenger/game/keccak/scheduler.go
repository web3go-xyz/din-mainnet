package keccak

import (
	"context"
	"sync"

	keccakTypes "github.com/ethereum-optimism/optimism/op-challenger/game/keccak/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type Challenger interface {
	Challenge(ctx context.Context, blockHash common.Hash, oracle Oracle, preimages []keccakTypes.LargePreimageMetaData) error
}

type LargePreimageScheduler struct {
	log        log.Logger
	ch         chan common.Hash
	oracles    []keccakTypes.LargePreimageOracle
	challenger Challenger
	cancel     func()
	wg         sync.WaitGroup
}

func NewLargePreimageScheduler(logger log.Logger, oracles []keccakTypes.LargePreimageOracle, challenger Challenger) *LargePreimageScheduler {
	return &LargePreimageScheduler{
		log:        logger,
		ch:         make(chan common.Hash, 1),
		oracles:    oracles,
		challenger: challenger,
	}
}

func (s *LargePreimageScheduler) Start(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	s.cancel = cancel
	s.wg.Add(1)
	go s.run(ctx)
}

func (s *LargePreimageScheduler) Close() error {
	s.cancel()
	s.wg.Wait()
	return nil
}

func (s *LargePreimageScheduler) run(ctx context.Context) {
	defer s.wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case blockHash := <-s.ch:
			if err := s.verifyPreimages(ctx, blockHash); err != nil {
				s.log.Error("Failed to verify large preimages", "err", err)
			}
		}
	}
}

func (s *LargePreimageScheduler) Schedule(blockHash common.Hash, _ uint64) error {
	select {
	case s.ch <- blockHash:
	default:
		s.log.Trace("Skipping preimage check while already processing")
		// Already busy processing, skip this update
	}
	return nil
}

func (s *LargePreimageScheduler) verifyPreimages(ctx context.Context, blockHash common.Hash) error {
	for _, oracle := range s.oracles {
		if err := s.verifyOraclePreimages(ctx, oracle, blockHash); err != nil {
			s.log.Error("Failed to verify preimages in oracle %v: %w", oracle.Addr(), err)
		}
	}
	return nil
}

func (s *LargePreimageScheduler) verifyOraclePreimages(ctx context.Context, oracle keccakTypes.LargePreimageOracle, blockHash common.Hash) error {
	preimages, err := oracle.GetActivePreimages(ctx, blockHash)
	if err != nil {
		return err
	}
	toVerify := make([]keccakTypes.LargePreimageMetaData, 0, len(preimages))
	for _, preimage := range preimages {
		if preimage.ShouldVerify() {
			toVerify = append(toVerify, preimage)
		}
	}
	return s.challenger.Challenge(ctx, blockHash, oracle, toVerify)
}
