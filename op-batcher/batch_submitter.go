package op_batcher

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"math/big"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	opmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	oppprof "github.com/ethereum-optimism/optimism/op-service/pprof"
	oprpc "github.com/ethereum-optimism/optimism/op-service/rpc"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum-optimism/optimism/op-batcher/sequencer"
	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-proposer/rollupclient"
	"github.com/ethereum-optimism/optimism/op-proposer/txmgr"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/urfave/cli"
)

const (
	// defaultDialTimeout is default duration the service will wait on
	// startup to make a connection to either the L1 or L2 backends.
	defaultDialTimeout = 5 * time.Second
)

// Main is the entrypoint into the Batch Submitter. This method returns a
// closure that executes the service and blocks until the service exits. The use
// of a closure allows the parameters bound to the top-level main package, e.g.
// GitVersion, to be captured and used once the function is executed.
func Main(version string) func(cliCtx *cli.Context) error {
	return func(cliCtx *cli.Context) error {
		cfg := NewConfig(cliCtx)
		if err := cfg.Check(); err != nil {
			return fmt.Errorf("invalid CLI flags: %w", err)
		}

		l := oplog.NewLogger(cfg.LogConfig)
		l.Info("Initializing Batch Submitter")

		batchSubmitter, err := NewBatchSubmitter(cfg, l)
		if err != nil {
			l.Error("Unable to create Batch Submitter", "error", err)
			return err
		}

		l.Info("Starting Batch Submitter")

		if err := batchSubmitter.Start(); err != nil {
			l.Error("Unable to start Batch Submitter", "error", err)
			return err
		}
		defer batchSubmitter.Stop()

		ctx, cancel := context.WithCancel(context.Background())

		l.Info("Batch Submitter started")
		pprofConfig := cfg.PprofConfig
		if pprofConfig.Enabled {
			l.Info("starting pprof", "addr", pprofConfig.ListenAddr, "port", pprofConfig.ListenPort)
			go func() {
				if err := oppprof.ListenAndServe(ctx, pprofConfig.ListenAddr, pprofConfig.ListenPort); err != nil {
					l.Error("error starting pprof", "err", err)
				}
			}()
		}

		registry := opmetrics.NewRegistry()
		metricsCfg := cfg.MetricsConfig
		if metricsCfg.Enabled {
			l.Info("starting metrics server", "addr", metricsCfg.ListenAddr, "port", metricsCfg.ListenPort)
			go func() {
				if err := opmetrics.ListenAndServe(ctx, registry, metricsCfg.ListenAddr, metricsCfg.ListenPort); err != nil {
					l.Error("error starting metrics server", err)
				}
			}()
		}

		rpcCfg := cfg.RPCConfig
		server := oprpc.NewServer(
			rpcCfg.ListenAddr,
			rpcCfg.ListenPort,
			version,
		)
		if err := server.Start(); err != nil {
			cancel()
			return fmt.Errorf("error starting RPC server: %w", err)
		}

		interruptChannel := make(chan os.Signal, 1)
		signal.Notify(interruptChannel, []os.Signal{
			os.Interrupt,
			os.Kill,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		}...)
		<-interruptChannel
		cancel()
		_ = server.Stop()
		return nil
	}
}

// BatchSubmitter encapsulates a service responsible for submitting L2 tx
// batches to L1 for availability.
type BatchSubmitter struct {
	txMgr txmgr.TxManager
	cfg   sequencer.Config
	wg    sync.WaitGroup
	done  chan struct{}
	log   log.Logger

	ctx    context.Context
	cancel context.CancelFunc

	lastSubmittedBlock eth.BlockID

	ch *derive.ChannelOut
}

// NewBatchSubmitter initializes the BatchSubmitter, gathering any resources
// that will be needed during operation.
func NewBatchSubmitter(cfg Config, l log.Logger) (*BatchSubmitter, error) {
	ctx := context.Background()

	var err error
	var sequencerPrivKey *ecdsa.PrivateKey
	var addr common.Address

	if cfg.PrivateKey != "" && cfg.Mnemonic != "" {
		return nil, errors.New("cannot specify both a private key and a mnemonic")
	}

	if cfg.PrivateKey == "" {
		// Parse wallet private key that will be used to submit L2 txs to the batch
		// inbox address.
		wallet, err := hdwallet.NewFromMnemonic(cfg.Mnemonic)
		if err != nil {
			return nil, err
		}

		acc := accounts.Account{
			URL: accounts.URL{
				Path: cfg.SequencerHDPath,
			},
		}
		addr, err = wallet.Address(acc)
		if err != nil {
			return nil, err
		}

		sequencerPrivKey, err = wallet.PrivateKey(acc)
		if err != nil {
			return nil, err
		}
	} else {
		sequencerPrivKey, err = crypto.HexToECDSA(strings.TrimPrefix(cfg.PrivateKey, "0x"))
		if err != nil {
			return nil, err
		}

		addr = crypto.PubkeyToAddress(sequencerPrivKey.PublicKey)
	}

	batchInboxAddress, err := parseAddress(cfg.SequencerBatchInboxAddress)
	if err != nil {
		return nil, err
	}

	// Connect to L1 and L2 providers. Perform these last since they are the
	// most expensive.
	l1Client, err := dialEthClientWithTimeout(ctx, cfg.L1EthRpc)
	if err != nil {
		return nil, err
	}

	l2Client, err := dialEthClientWithTimeout(ctx, cfg.L2EthRpc)
	if err != nil {
		return nil, err
	}

	rollupClient, err := dialRollupClientWithTimeout(ctx, cfg.RollupRpc)
	if err != nil {
		return nil, err
	}

	chainID, err := l1Client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	sequencerBalance, err := l1Client.BalanceAt(ctx, addr, nil)
	if err != nil {
		return nil, err
	}

	log.Info("starting batch submitter", "submitter_addr", addr, "submitter_bal", sequencerBalance)

	txManagerConfig := txmgr.Config{
		Log:                       l,
		Name:                      "Batch Submitter",
		ResubmissionTimeout:       cfg.ResubmissionTimeout,
		ReceiptQueryInterval:      time.Second,
		NumConfirmations:          cfg.NumConfirmations,
		SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount,
	}

	batcherCfg := sequencer.Config{
		Log:               l,
		Name:              "Batch Submitter",
		L1Client:          l1Client,
		L2Client:          l2Client,
		RollupNode:        rollupClient,
		MinL1TxSize:       cfg.MinL1TxSize,
		MaxL1TxSize:       cfg.MaxL1TxSize,
		BatchInboxAddress: batchInboxAddress,
		ChannelTimeout:    cfg.ChannelTimeout,
		ChainID:           chainID,
		PrivKey:           sequencerPrivKey,
		PollInterval:      cfg.PollInterval,
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &BatchSubmitter{
		cfg:   batcherCfg,
		txMgr: txmgr.NewSimpleTxManager("batcher", txManagerConfig, l1Client),
		done:  make(chan struct{}),
		log:   l,
		// TODO: this context only exists because the even loop doesn't reach done
		// if the tx manager is blocking forever due to e.g. insufficient balance.
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

func (l *BatchSubmitter) Start() error {
	l.wg.Add(1)
	go l.loop()
	return nil
}

func (l *BatchSubmitter) Stop() {
	l.cancel()
	close(l.done)
	l.wg.Wait()
}

func (l *BatchSubmitter) loop() {
	defer l.wg.Done()

	ticker := time.NewTicker(l.cfg.PollInterval)
	defer ticker.Stop()
mainLoop:
	for {
		select {
		case <-ticker.C:
			// Do the simplest thing of one channel per range of blocks since the iteration of this loop.
			// The channel is closed at the end of this loop (to avoid lifecycle management of the channel).
			ctx, cancel := context.WithTimeout(l.ctx, time.Second*10)
			syncStatus, err := l.cfg.RollupNode.SyncStatus(ctx)
			cancel()
			if err != nil {
				l.log.Warn("issue fetching L2 head", "err", err)
				continue
			}
			if syncStatus.HeadL1 == (eth.L1BlockRef{}) {
				l.log.Info("Rollup node has no L1 head info yet")
				continue
			}
			l.log.Info("Got new L2 sync status", "safe_head", syncStatus.SafeL2, "unsafe_head", syncStatus.UnsafeL2, "last_submitted", l.lastSubmittedBlock, "l1_head", syncStatus.HeadL1)
			if syncStatus.SafeL2.Number >= syncStatus.UnsafeL2.Number {
				l.log.Trace("No unsubmitted blocks from sequencer")
				continue
			}
			// If we just started, start at safe-head
			if l.lastSubmittedBlock == (eth.BlockID{}) {
				l.log.Info("Starting batch-submitter work at safe-head", "safe", syncStatus.SafeL2)
				l.lastSubmittedBlock = syncStatus.SafeL2.ID()
			}
			// If it's lagging behind, catch it up.
			if l.lastSubmittedBlock.Number < syncStatus.SafeL2.Number {
				l.log.Warn("last submitted block lagged behind L2 safe head: batch submission will continue from the safe head now", "last", l.lastSubmittedBlock, "safe", syncStatus.SafeL2)
				l.lastSubmittedBlock = syncStatus.SafeL2.ID()
			}
			if ch, err := derive.NewChannelOut(syncStatus.HeadL1.Time); err != nil {
				l.log.Error("Error creating channel", "err", err)
				continue
			} else {
				l.ch = ch
			}
			prevID := l.lastSubmittedBlock
			for i := l.lastSubmittedBlock.Number + 1; i <= syncStatus.UnsafeL2.Number; i++ {
				ctx, cancel := context.WithTimeout(l.ctx, time.Second*10)
				block, err := l.cfg.L2Client.BlockByNumber(ctx, new(big.Int).SetUint64(i))
				cancel()
				if err != nil {
					l.log.Error("issue fetching L2 block", "err", err)
					continue mainLoop
				}
				if block.ParentHash() != prevID.Hash {
					l.log.Error("detected a reorg in L2 chain vs previous submitted information, resetting to safe head now", "safe_head", syncStatus.SafeL2)
					l.lastSubmittedBlock = syncStatus.SafeL2.ID()
					continue mainLoop
				}
				if err := l.ch.AddBlock(block); err != nil {
					l.log.Error("issue adding L2 Block to the channel", "err", err, "channel_id", l.ch.ID())
					continue mainLoop
				}
				prevID = eth.BlockID{Hash: block.Hash(), Number: block.NumberU64()}
				l.log.Info("added L2 block to channel", "block", prevID, "channel_id", l.ch.ID(), "tx_count", len(block.Transactions()), "time", block.Time())
			}
			if err := l.ch.Close(); err != nil {
				l.log.Error("issue getting adding L2 Block", "err", err)
				continue
			}
			// Hand role do-while loop to fully pull all frames out of the channel
			for {
				// Collect the output frame
				data := new(bytes.Buffer)
				data.WriteByte(derive.DerivationVersion0)
				done := false
				// subtract one, to account for the version byte
				if err := l.ch.OutputFrame(data, l.cfg.MaxL1TxSize-1); err == io.EOF {
					done = true
				} else if err != nil {
					l.log.Error("error outputting frame", "err", err)
					continue mainLoop
				}

				// Query for the submitter's current nonce.
				walletAddr := crypto.PubkeyToAddress(l.cfg.PrivKey.PublicKey)
				ctx, cancel = context.WithTimeout(l.ctx, time.Second*10)
				nonce, err := l.cfg.L1Client.NonceAt(ctx, walletAddr, nil)
				cancel()
				if err != nil {
					l.log.Error("unable to get current nonce", "err", err)
					continue mainLoop
				}

				// Create the transaction
				ctx, cancel = context.WithTimeout(l.ctx, time.Second*10)
				tx, err := l.CraftTx(ctx, data.Bytes(), nonce)
				cancel()
				if err != nil {
					l.log.Error("unable to craft tx", "err", err)
					continue mainLoop
				}

				// Construct the a closure that will update the txn with the current gas prices.
				updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
					l.log.Debug("updating batch tx gas price")
					return l.UpdateGasPrice(ctx, tx)
				}

				// Wait until one of our submitted transactions confirms. If no
				// receipt is received it's likely our gas price was too low.
				// TODO: does the tx manager nicely replace the tx?
				//  (submit a new one, that's within the channel timeout, but higher fee than previously submitted tx? Or use a cheap cancel tx?)
				ctx, cancel = context.WithTimeout(l.ctx, time.Second*time.Duration(l.cfg.ChannelTimeout))
				receipt, err := l.txMgr.Send(ctx, updateGasPrice, l.cfg.L1Client.SendTransaction)
				cancel()
				if err != nil {
					l.log.Warn("unable to publish tx", "err", err)
					continue mainLoop
				}

				// The transaction was successfully submitted.
				l.log.Info("tx successfully published", "tx_hash", receipt.TxHash, "channel_id", l.ch.ID())

				// If `ch.OutputFrame` returned io.EOF we don't need to submit any more frames for this channel.
				if done {
					break // local do-while loop
				}
			}
			// TODO: if we exit to the mainLoop early on an error,
			// it would be nice if we can determine which blocks are still readable from the partially submitted data.
			// We can open a channel-in-reader, parse the data up to which we managed to submit it,
			// and then take the block hash (if we remember which blocks we put in the channel)
			//
			// Now we just continue batch submission from the end of the channel.
			l.lastSubmittedBlock = prevID

		case <-l.done:
			return
		}
	}
}

// NOTE: This method SHOULD NOT publish the resulting transaction.
func (l *BatchSubmitter) CraftTx(ctx context.Context, data []byte, nonce uint64) (*types.Transaction, error) {
	gasTipCap, err := l.cfg.L1Client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, err
	}

	head, err := l.cfg.L1Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	gasFeeCap := txmgr.CalcGasFeeCap(head.BaseFee, gasTipCap)

	rawTx := &types.DynamicFeeTx{
		ChainID:   l.cfg.ChainID,
		Nonce:     nonce,
		To:        &l.cfg.BatchInboxAddress,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Data:      data,
	}
	l.log.Debug("creating tx", "to", rawTx.To, "from", crypto.PubkeyToAddress(l.cfg.PrivKey.PublicKey))

	gas, err := core.IntrinsicGas(rawTx.Data, nil, false, true, true)
	if err != nil {
		return nil, err
	}
	rawTx.Gas = gas

	return types.SignNewTx(l.cfg.PrivKey, types.LatestSignerForChainID(l.cfg.ChainID), rawTx)
}

// UpdateGasPrice signs an otherwise identical txn to the one provided but with
// updated gas prices sampled from the existing network conditions.
//
// NOTE: Thie method SHOULD NOT publish the resulting transaction.
func (l *BatchSubmitter) UpdateGasPrice(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	gasTipCap, err := l.cfg.L1Client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, err
	}

	head, err := l.cfg.L1Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	gasFeeCap := txmgr.CalcGasFeeCap(head.BaseFee, gasTipCap)

	rawTx := &types.DynamicFeeTx{
		ChainID:   l.cfg.ChainID,
		Nonce:     tx.Nonce(),
		To:        tx.To(),
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       tx.Gas(),
		Data:      tx.Data(),
	}

	return types.SignNewTx(l.cfg.PrivKey, types.LatestSignerForChainID(l.cfg.ChainID), rawTx)
}

// SendTransaction injects a signed transaction into the pending pool for
// execution.
func (l *BatchSubmitter) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return l.cfg.L1Client.SendTransaction(ctx, tx)
}

// dialEthClientWithTimeout attempts to dial the L1 provider using the provided
// URL. If the dial doesn't complete within defaultDialTimeout seconds, this
// method will return an error.
func dialEthClientWithTimeout(ctx context.Context, url string) (
	*ethclient.Client, error) {

	ctxt, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancel()

	return ethclient.DialContext(ctxt, url)
}

// dialRollupClientWithTimeout attempts to dial the RPC provider using the provided
// URL. If the dial doesn't complete within defaultDialTimeout seconds, this
// method will return an error.
func dialRollupClientWithTimeout(ctx context.Context, url string) (*rollupclient.RollupClient, error) {
	ctxt, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancel()

	client, err := rpc.DialContext(ctxt, url)
	if err != nil {
		return nil, err
	}

	return rollupclient.NewRollupClient(client), nil
}

// parseAddress parses an ETH address from a hex string. This method will fail if
// the address is not a valid hexadecimal address.
func parseAddress(address string) (common.Address, error) {
	if common.IsHexAddress(address) {
		return common.HexToAddress(address), nil
	}
	return common.Address{}, fmt.Errorf("invalid address: %v", address)
}
