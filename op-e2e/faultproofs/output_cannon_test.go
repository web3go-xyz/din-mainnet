package faultproofs

import (
	"context"
	"testing"

	op_e2e "github.com/ethereum-optimism/optimism/op-e2e"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestOutputCannonGame(t *testing.T) {
	op_e2e.InitParallel(t, op_e2e.UsesCannon, op_e2e.UseExecutor(0))
	ctx := context.Background()
	sys, l1Client := startFaultDisputeSystem(t)
	t.Cleanup(sys.Close)

	rollupEndpoint := sys.RollupNodes["sequencer"].HTTPEndpoint()
	l1Endpoint := sys.NodeEndpoint("l1")
	l2Endpoint := sys.NodeEndpoint("sequencer")
	require.NotEqual(t, rollupEndpoint, l2Endpoint)

	disputeGameFactory := disputegame.NewFactoryHelper(t, ctx, sys.Cfg.L1Deployments, l1Client)
	game := disputeGameFactory.StartOutputCannonGame(ctx, rollupEndpoint, common.Hash{0x01})
	game.LogGameData(ctx)

	game.StartChallenger(ctx, sys.RollupConfig, sys.L2GenesisCfg, rollupEndpoint, l1Endpoint, l2Endpoint, "Challenger",
		challenger.WithPrivKey(sys.Cfg.Secrets.Alice),
	)

	game.LogGameData(ctx)
	// Challenger should post an output root to counter claims down to the leaf level of the top game
	splitDepth := game.SplitDepth(ctx)
	for i := int64(1); i < splitDepth; i += 2 {
		game.WaitForCorrectOutputRoot(ctx, i)
		game.Attack(ctx, i, common.Hash{0xaa})
		game.LogGameData(ctx)
	}
	game.WaitForCorrectOutputRoot(ctx, splitDepth)

	// Post the first cannon output root (with 01 status code to show the output root is invalid)
	game.Attack(ctx, splitDepth, common.Hash{0x01})

	// Challenger should counter
	game.WaitForClaimAtDepth(ctx, int(splitDepth+2))
	game.LogGameData(ctx)
}
