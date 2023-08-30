package etl

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum-optimism/optimism/indexer/config"
	"github.com/ethereum-optimism/optimism/indexer/database"
	"github.com/ethereum-optimism/optimism/indexer/node"
	"github.com/ethereum-optimism/optimism/op-service/retry"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type L1ETL struct {
	ETL

	db *database.DB
}

// NewL1ETL creates a new L1ETL instance that will start indexing from different starting points
// depending on the state of the database and the supplied start height.
func NewL1ETL(cfg Config, log log.Logger, db *database.DB, metrics Metricer, client node.EthClient, contracts config.L1Contracts) (*L1ETL, error) {
	log = log.New("etl", "l1")

	latestHeader, err := db.Blocks.L1LatestBlockHeader()
	if err != nil {
		return nil, err
	}

	cSlice, err := contracts.AsSlice()
	if err != nil {
		return nil, err
	}

	// Determine the starting height for traversal
	var fromHeader *types.Header
	if latestHeader != nil {
		log.Info("detected last indexed block", "number", latestHeader.Number, "hash", latestHeader.Hash)
		fromHeader = latestHeader.RLPHeader.Header()

	} else if cfg.StartHeight.BitLen() > 0 {
		log.Info("no indexed state in storage, starting from supplied L1 height", "height", cfg.StartHeight.String())
		header, err := client.BlockHeaderByNumber(cfg.StartHeight)
		if err != nil {
			return nil, fmt.Errorf("could not fetch starting block header: %w", err)
		}

		fromHeader = header

	} else {
		log.Info("no indexed state in storage, starting from L1 genesis")
	}

	// NOTE - The use of un-buffered channel here assumes that downstream consumers
	// will be able to keep up with the rate of incoming batches
	etlBatches := make(chan ETLBatch)
	etl := ETL{
		loopInterval:     time.Duration(cfg.LoopIntervalMsec) * time.Millisecond,
		headerBufferSize: uint64(cfg.HeaderBufferSize),

		log:             log,
		metrics:         metrics,
		headerTraversal: node.NewHeaderTraversal(client, fromHeader, cfg.ConfirmationDepth),
		ethClient:       client,
		contracts:       cSlice,
		etlBatches:      etlBatches,
	}

	return &L1ETL{ETL: etl, db: db}, nil
}

func (l1Etl *L1ETL) Start(ctx context.Context) error {
	errCh := make(chan error, 1)
	go func() {
		errCh <- l1Etl.ETL.Start(ctx)
	}()

	for {
		select {
		case err := <-errCh:
			return err

		// Index incoming batches (only L1 blocks that have an emitted log)
		case batch := <-l1Etl.etlBatches:
			l1BlockHeaders := make([]database.L1BlockHeader, 0, len(batch.Headers))
			for i := range batch.Headers {
				if _, ok := batch.HeadersWithLog[batch.Headers[i].Hash()]; ok {
					l1BlockHeaders = append(l1BlockHeaders, database.L1BlockHeader{BlockHeader: database.BlockHeaderFromHeader(&batch.Headers[i])})
				}
			}
			if len(l1BlockHeaders) == 0 {
				batch.Logger.Info("no l1 blocks with logs in batch")
				continue
			}

			l1ContractEvents := make([]database.L1ContractEvent, len(batch.Logs))
			for i := range batch.Logs {
				timestamp := batch.HeaderMap[batch.Logs[i].BlockHash].Time
				l1ContractEvents[i] = database.L1ContractEvent{ContractEvent: database.ContractEventFromLog(&batch.Logs[i], timestamp)}
			}

			// Continually try to persist this batch. If it fails after 10 attempts, we simply error out
			retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
			if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
				if err := l1Etl.db.Transaction(func(tx *database.DB) error {
					if err := tx.Blocks.StoreL1BlockHeaders(l1BlockHeaders); err != nil {
						return err
					}
					// we must have logs if we have l1 blocks
					if err := tx.ContractEvents.StoreL1ContractEvents(l1ContractEvents); err != nil {
						return err
					}
					return nil
				}); err != nil {
					batch.Logger.Error("unable to persist batch", "err", err)
					return nil, err
				}

				l1Etl.ETL.metrics.RecordIndexedHeaders(len(l1BlockHeaders))
				l1Etl.ETL.metrics.RecordIndexedLatestHeight(l1BlockHeaders[len(l1BlockHeaders)-1].Number)
				l1Etl.ETL.metrics.RecordIndexedLogs(len(l1ContractEvents))

				// a-ok!
				return nil, nil
			}); err != nil {
				return err
			}

			batch.Logger.Info("indexed batch")
		}
	}
}
