package solver

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrStepNonLeafNode = errors.New("cannot step on non-leaf claims")
	ErrStepAgreedClaim = errors.New("cannot step on claims we agree with")
)

// Solver uses a [TraceProvider] to determine the moves to make in a dispute game.
type Solver struct {
	trace     types.TraceProvider
	gameDepth int
}

// NewSolver creates a new [Solver] using the provided [TraceProvider].
func NewSolver(gameDepth int, traceProvider types.TraceProvider) *Solver {
	return &Solver{
		traceProvider,
		gameDepth,
	}
}

// NextMove returns the next move to make given the current state of the game.
func (s *Solver) NextMove(ctx context.Context, claim types.Claim, agreeWithClaimLevel bool) (*types.Claim, error) {
	if agreeWithClaimLevel {
		return nil, nil
	}
	if claim.Depth() == s.gameDepth {
		return nil, types.ErrGameDepthReached
	}
	agree, err := s.agreeWithClaim(ctx, claim.ClaimData)
	if err != nil {
		return nil, err
	}
	if agree {
		return s.defend(ctx, claim)
	} else {
		return s.attack(ctx, claim)
	}
}

type StepData struct {
	LeafClaim  types.Claim
	IsAttack   bool
	PreState   []byte
	ProofData  []byte
	OracleData *types.PreimageOracleData
}

// AttemptStep determines what step should occur for a given leaf claim.
// An error will be returned if the claim is not at the max depth.
func (s *Solver) AttemptStep(ctx context.Context, claim types.Claim, agreeWithClaimLevel bool) (StepData, error) {
	if claim.Depth() != s.gameDepth {
		return StepData{}, ErrStepNonLeafNode
	}
	if agreeWithClaimLevel {
		return StepData{}, ErrStepAgreedClaim
	}
	claimCorrect, err := s.agreeWithClaim(ctx, claim.ClaimData)
	if err != nil {
		return StepData{}, err
	}
	index := claim.TraceIndex(s.gameDepth)
	var preState []byte
	var proofData []byte
	var oracleData *types.PreimageOracleData

	if !claimCorrect {
		// Attack the claim by executing step index, so we need to get the pre-state of that index
		preState, proofData, oracleData, err = s.trace.GetStepData(ctx, index)
		if err != nil {
			return StepData{}, err
		}
	} else {
		// We agree with the claim so Defend and use this claim as the starting point to execute the step after
		// Thus we need the pre-state of the next step
		// Note: This makes our maximum depth 63 because we need to add 1 without overflowing.
		preState, proofData, oracleData, err = s.trace.GetStepData(ctx, index+1)
		if err != nil {
			return StepData{}, err
		}
	}

	return StepData{
		LeafClaim:  claim,
		IsAttack:   !claimCorrect,
		PreState:   preState,
		ProofData:  proofData,
		OracleData: oracleData,
	}, nil
}

// attack returns a response that attacks the claim.
func (s *Solver) attack(ctx context.Context, claim types.Claim) (*types.Claim, error) {
	position := claim.Attack()
	value, err := s.traceAtPosition(ctx, position)
	if err != nil {
		return nil, fmt.Errorf("attack claim: %w", err)
	}
	return &types.Claim{
		ClaimData:           types.ClaimData{Value: value, Position: position},
		Parent:              claim.ClaimData,
		ParentContractIndex: claim.ContractIndex,
	}, nil
}

// defend returns a response that defends the claim.
func (s *Solver) defend(ctx context.Context, claim types.Claim) (*types.Claim, error) {
	if claim.IsRoot() {
		return nil, nil
	}
	position := claim.Defend()
	value, err := s.traceAtPosition(ctx, position)
	if err != nil {
		return nil, fmt.Errorf("defend claim: %w", err)
	}
	return &types.Claim{
		ClaimData:           types.ClaimData{Value: value, Position: position},
		Parent:              claim.ClaimData,
		ParentContractIndex: claim.ContractIndex,
	}, nil
}

// agreeWithClaim returns true if the claim is correct according to the internal [TraceProvider].
func (s *Solver) agreeWithClaim(ctx context.Context, claim types.ClaimData) (bool, error) {
	ourValue, err := s.traceAtPosition(ctx, claim.Position)
	return bytes.Equal(ourValue[:], claim.Value[:]), err
}

// traceAtPosition returns the [common.Hash] from internal [TraceProvider] at the given [Position].
func (s *Solver) traceAtPosition(ctx context.Context, p types.Position) (common.Hash, error) {
	index := p.TraceIndex(s.gameDepth)
	hash, err := s.trace.Get(ctx, index)
	return hash, err
}
