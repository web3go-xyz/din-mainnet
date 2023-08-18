package fault

import (
	"context"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-challenger/config"
	"github.com/ethereum-optimism/optimism/op-challenger/fault/alphabet"
	"github.com/ethereum-optimism/optimism/op-challenger/fault/cannon"
	"github.com/ethereum-optimism/optimism/op-challenger/fault/types"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type Actor interface {
	Act(ctx context.Context) error
}

type GameInfo interface {
	GetGameStatus(context.Context) (types.GameStatus, error)
	LogGameInfo(ctx context.Context)
}

type Game struct {
	agent                   Actor
	agreeWithProposedOutput bool
	caller                  GameInfo
	logger                  log.Logger
}

func NewGame(
	ctx context.Context,
	logger log.Logger,
	cfg *config.Config,
	addr common.Address,
	txMgr txmgr.TxManager,
	client *ethclient.Client,
) (*Game, error) {
	contract, err := bindings.NewFaultDisputeGameCaller(addr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind the fault dispute game contract: %w", err)
	}

	loader := NewLoader(contract)

	gameDepth, err := loader.FetchGameDepth(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the game depth: %w", err)
	}

	var provider types.TraceProvider
	var updater types.OracleUpdater
	switch cfg.TraceType {
	case config.TraceTypeCannon:
		provider, err = cannon.NewTraceProvider(ctx, logger, cfg, client, addr)
		if err != nil {
			return nil, fmt.Errorf("create cannon trace provider: %w", err)
		}
		updater, err = cannon.NewOracleUpdater(ctx, logger, txMgr, addr, client)
		if err != nil {
			return nil, fmt.Errorf("failed to create the cannon updater: %w", err)
		}
	case config.TraceTypeAlphabet:
		provider = alphabet.NewTraceProvider(cfg.AlphabetTrace, gameDepth)
		updater = alphabet.NewOracleUpdater(logger)
	default:
		return nil, fmt.Errorf("unsupported trace type: %v", cfg.TraceType)
	}

	if err := ValidateAbsolutePrestate(ctx, provider, loader); err != nil {
		return nil, fmt.Errorf("failed to validate absolute prestate: %w", err)
	}

	gameLogger := logger.New("game", addr)
	responder, err := NewFaultResponder(gameLogger, txMgr, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create the responder: %w", err)
	}

	caller, err := NewFaultCallerFromBindings(addr, client, gameLogger)
	if err != nil {
		return nil, fmt.Errorf("failed to bind the fault contract: %w", err)
	}

	return &Game{
		agent:                   NewAgent(loader, int(gameDepth), provider, responder, updater, cfg.AgreeWithProposedOutput, gameLogger),
		agreeWithProposedOutput: cfg.AgreeWithProposedOutput,
		caller:                  caller,
		logger:                  gameLogger,
	}, nil
}

func (g *Game) ProgressGame(ctx context.Context) bool {
	g.logger.Trace("Checking if actions are required")
	if err := g.agent.Act(ctx); err != nil {
		g.logger.Error("Error when acting on game", "err", err)
	}
	if status, err := g.caller.GetGameStatus(ctx); err != nil {
		g.logger.Warn("Unable to retrieve game status", "err", err)
	} else if status != 0 {
		var expectedStatus types.GameStatus
		if g.agreeWithProposedOutput {
			expectedStatus = types.GameStatusChallengerWon
		} else {
			expectedStatus = types.GameStatusDefenderWon
		}
		if expectedStatus == status {
			g.logger.Info("Game won", "status", status)
		} else {
			g.logger.Error("Game lost", "status", status)
		}
		return true
	} else {
		g.caller.LogGameInfo(ctx)
	}
	return false
}
