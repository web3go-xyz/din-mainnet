package solver

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/ethereum-optimism/optimism/cannon/mipsevm"
	faulttest "github.com/ethereum-optimism/optimism/op-challenger/game/fault/test"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestCalculateNextActions(t *testing.T) {
	maxDepth := 4
	claimBuilder := faulttest.NewAlphabetClaimBuilder(t, maxDepth)

	tests := []struct {
		name                string
		agreeWithOutputRoot bool
		rootClaimCorrect    bool
		setupGame           func(builder *faulttest.GameBuilder)
	}{
		{
			name:                "AttackRootClaim",
			agreeWithOutputRoot: true,
			setupGame: func(builder *faulttest.GameBuilder) {
				builder.Seq().ExpectAttack()
			},
		},
		{
			name:                "DoNotAttackRootClaimWhenDisagreeWithOutputRoot",
			agreeWithOutputRoot: false,
			setupGame:           func(builder *faulttest.GameBuilder) {},
		},
		{
			// Note: The fault dispute game contract should prevent a correct root claim from actually being posted
			// But for completeness, test we ignore it so we don't get sucked into playing an unwinnable game.
			name:                "DoNotAttackCorrectRootClaim_AgreeWithOutputRoot",
			agreeWithOutputRoot: true,
			rootClaimCorrect:    true,
			setupGame:           func(builder *faulttest.GameBuilder) {},
		},
		{
			// Note: The fault dispute game contract should prevent a correct root claim from actually being posted
			// But for completeness, test we ignore it so we don't get sucked into playing an unwinnable game.
			name:                "DoNotAttackCorrectRootClaim_DisagreeWithOutputRoot",
			agreeWithOutputRoot: false,
			rootClaimCorrect:    true,
			setupGame:           func(builder *faulttest.GameBuilder) {},
		},

		{
			name:                "DoNotPerformDuplicateMoves",
			agreeWithOutputRoot: true,
			setupGame: func(builder *faulttest.GameBuilder) {
				// Expected move has already been made.
				builder.Seq().AttackCorrect()
			},
		},

		{
			name:                "RespondToAllClaimsAtDisagreeingLevel",
			agreeWithOutputRoot: true,
			setupGame: func(builder *faulttest.GameBuilder) {
				honestClaim := builder.Seq().AttackCorrect()
				honestClaim.AttackCorrect().ExpectDefend()
				honestClaim.DefendCorrect().ExpectDefend()
				honestClaim.Attack(common.Hash{0xaa}).ExpectAttack()
				honestClaim.Attack(common.Hash{0xbb}).ExpectAttack()
				honestClaim.Defend(common.Hash{0xcc}).ExpectAttack()
				honestClaim.Defend(common.Hash{0xdd}).ExpectAttack()
			},
		},

		{
			name:                "StepAtMaxDepth",
			agreeWithOutputRoot: true,
			setupGame: func(builder *faulttest.GameBuilder) {
				lastHonestClaim := builder.Seq().
					AttackCorrect().
					AttackCorrect().
					DefendCorrect()
				lastHonestClaim.AttackCorrect().ExpectStepDefend()
				lastHonestClaim.Attack(common.Hash{0xdd}).ExpectStepAttack()
			},
		},

		{
			name:                "PoisonedPreState",
			agreeWithOutputRoot: true,
			setupGame: func(builder *faulttest.GameBuilder) {
				// Create a cannon state that is exited and resolved that the output root was invalid
				maliciousState := mipsevm.State{
					Memory:   mipsevm.NewMemory(),
					ExitCode: 1,
					Exited:   true,
				}
				maliciousStateHash, _ := maliciousState.EncodeWitness().StateHash()

				// Dishonest actor counters their own claims to set up a situation with an invalid prestate
				// The honest actor should attack all claims that support the root claim (disagree with the output root)
				builder.Seq().ExpectAttack(). // This expected action is the winning move.
								Attack(maliciousStateHash).
								Defend(maliciousStateHash).ExpectAttack().
								Attack(maliciousStateHash).
								Attack(maliciousStateHash).ExpectStepAttack()

				// The attempt to step against our malicious leaf node will fail because the pre-state won't match our
				// malicious state hash. However, it is the very first expected action, attacking the root claim with
				// the correct hash that wins the game since it will be the left-most uncountered claim.
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			builder := claimBuilder.GameBuilder(test.agreeWithOutputRoot, test.rootClaimCorrect)
			test.setupGame(builder)
			game := builder.Game
			for i, claim := range game.Claims() {
				t.Logf("Claim %v: Pos: %v ParentIdx: %v, Countered: %v, Value: %v", i, claim.Position.ToGIndex(), claim.ParentContractIndex, claim.Countered, claim.Value)
			}

			solver := NewGameSolver(maxDepth, claimBuilder.CorrectTraceProvider())
			actions, err := solver.CalculateNextActions(context.Background(), game)
			require.NoError(t, err)
			for i, action := range actions {
				t.Logf("Move %v: Type: %v, ParentIdx: %v, Attack: %v, Value: %v, PreState: %v, ProofData: %v",
					i, action.Type, action.ParentIdx, action.IsAttack, action.Value, hex.EncodeToString(action.PreState), hex.EncodeToString(action.ProofData))
				// Check that every move the solver returns meets the generic validation rules
				require.NoError(t, checkRules(game, action), "Attempting to perform invalid action")
			}
			for i, action := range builder.ExpectedActions {
				t.Logf("Expect %v: Type: %v, ParentIdx: %v, Attack: %v, Value: %v, PreState: %v, ProofData: %v",
					i, action.Type, action.ParentIdx, action.IsAttack, action.Value, hex.EncodeToString(action.PreState), hex.EncodeToString(action.ProofData))
				require.Containsf(t, actions, action, "Expected claim %v missing", i)
			}
			require.Len(t, actions, len(builder.ExpectedActions), "Incorrect number of actions")
		})
	}
}
