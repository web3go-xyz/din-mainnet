package batchsubmitter

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum-optimism/optimism/go/batch-submitter/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

// Driver is an interface for creating and submitting batch transactions for a
// specific contract.
type Driver interface {
	// Name is an identifier used to prefix logs for a particular service.
	Name() string

	// WalletAddr is the wallet address used to pay for batch transaction
	// fees.
	WalletAddr() common.Address

	// GetBatchBlockRange returns the start and end L2 block heights that
	// need to be processed. Note that the end value is *exclusive*,
	// therefore if the returned values are identical nothing needs to be
	// processed.
	GetBatchBlockRange(ctx context.Context) (*big.Int, *big.Int, error)

	// SubmitBatchTx transforms the L2 blocks between start and end into a
	// batch transaction using the given nonce and gasPrice. The final
	// transaction is published and returned to the call.
	SubmitBatchTx(
		ctx context.Context,
		start, end, nonce, gasPrice *big.Int,
	) (*types.Transaction, error)
}

type ServiceConfig struct {
	Context         context.Context
	Driver          Driver
	PollInterval    time.Duration
	L1Client        *ethclient.Client
	TxManagerConfig txmgr.Config
}

type Service struct {
	cfg    ServiceConfig
	ctx    context.Context
	cancel func()

	txMgr txmgr.TxManager

	wg sync.WaitGroup
}

func NewService(cfg ServiceConfig) *Service {
	ctx, cancel := context.WithCancel(cfg.Context)

	txMgr := txmgr.NewSimpleTxManager(
		cfg.Driver.Name(), cfg.TxManagerConfig, cfg.L1Client,
	)

	return &Service{
		cfg:    cfg,
		ctx:    ctx,
		cancel: cancel,
		txMgr:  txMgr,
	}
}

func (s *Service) Start() error {
	s.wg.Add(1)
	go s.eventLoop()
	return nil
}

func (s *Service) Stop() error {
	s.cancel()
	s.wg.Wait()
	return nil
}

func (s *Service) eventLoop() {
	defer s.wg.Done()

	name := s.cfg.Driver.Name()

	for {
		select {
		case <-time.After(s.cfg.PollInterval):
			log.Info(name + " fetching current block range")

			start, end, err := s.cfg.Driver.GetBatchBlockRange(s.ctx)
			if err != nil {
				log.Error(name+" unable to get block range", "err", err)
				continue
			}

			log.Info(name+" block range", "start", start, "end", end)

			// No new updates.
			if start.Cmp(end) == 0 {
				continue
			}

			nonce64, err := s.cfg.L1Client.NonceAt(
				s.ctx, s.cfg.Driver.WalletAddr(), nil,
			)
			if err != nil {
				log.Error(name+" unable to get current nonce",
					"err", err)
				continue
			}
			nonce := new(big.Int).SetUint64(nonce64)

			sendTx := func(
				ctx context.Context,
				gasPrice *big.Int,
			) (*types.Transaction, error) {
				log.Info(name+" attempting batch tx", "start", start,
					"end", end, "nonce", nonce,
					"gasPrice", gasPrice)
				return s.cfg.Driver.SubmitBatchTx(
					ctx, start, end, nonce, gasPrice,
				)
			}

			receipt, err := s.txMgr.Send(s.ctx, sendTx)
			if err != nil {
				log.Error(name+" unable to publish batch tx",
					"err", err)
				continue
			}

			log.Info(name+" batch tx successfully published",
				"tx_hash", receipt.TxHash)

		case err := <-s.ctx.Done():
			log.Error(name+" service shutting down", "err", err)
			return
		}
	}
}
