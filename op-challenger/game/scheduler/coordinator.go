package scheduler

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-challenger/game/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/exp/slices"
)

var errUnknownGame = errors.New("unknown game")

type PlayerCreator func(address common.Address, dir string) (GamePlayer, error)

type gameState struct {
	player   GamePlayer
	inflight bool
	status   types.GameStatus
}

// coordinator manages the set of current games, queues games to be played (on separate worker threads) and
// cleans up data files once a game is resolved.
// All function calls must be made on the same thread.
type coordinator struct {
	// jobQueue is the outgoing queue for jobs being sent to workers for progression
	jobQueue chan<- job

	// resultQueue is the incoming queue of jobs that have been completed by workers
	resultQueue <-chan job

	logger       log.Logger
	m            SchedulerMetricer
	createPlayer PlayerCreator
	states       map[common.Address]*gameState
	disk         DiskManager
}

// schedule takes the current list of games to attempt to progress, filters out games that have previous
// progressions already in-flight and schedules jobs to progress on the outbound jobQueue.
// To avoid deadlock, it may process results from the inbound resultQueue while adding jobs to the outbound jobQueue.
// Returns an error if a game couldn't be scheduled because of an error. It will continue attempting to progress
// all games even if an error occurs with one game.
func (c *coordinator) schedule(ctx context.Context, games []common.Address) error {
	// First remove any game states we no longer require
	for addr, state := range c.states {
		if !state.inflight && !slices.Contains(games, addr) {
			delete(c.states, addr)
		}
	}

	var gamesInProgress int
	var gamesChallengerWon int
	var gamesDefenderWon int
	var errs []error
	var jobs []job
	// Next collect all the jobs to schedule and ensure all games are recorded in the states map.
	// Otherwise, results may start being processed before all games are recorded, resulting in existing
	// data directories potentially being deleted for games that are required.
	for _, addr := range games {
		if j, err := c.createJob(addr); err != nil {
			errs = append(errs, err)
		} else if j != nil {
			jobs = append(jobs, *j)
			c.m.RecordGameUpdateScheduled()
		}
		state, ok := c.states[addr]
		if ok {
			switch state.status {
			case types.GameStatusInProgress:
				gamesInProgress++
			case types.GameStatusDefenderWon:
				gamesDefenderWon++
			case types.GameStatusChallengerWon:
				gamesChallengerWon++
			}
		} else {
			c.logger.Warn("Game not found in states map", "game", addr)
		}
	}
	c.m.RecordGamesStatus(gamesInProgress, gamesChallengerWon, gamesDefenderWon)

	// Finally, enqueue the jobs
	for _, j := range jobs {
		errs = append(errs, c.enqueueJob(ctx, j))
	}
	return errors.Join(errs...)
}

// createJob updates the state for the specified game and returns the job to enqueue for it, if any
// Returns (nil, nil) when there is no error and no job to enqueue
func (c *coordinator) createJob(game common.Address) (*job, error) {
	state, ok := c.states[game]
	if !ok {
		state = &gameState{}
		c.states[game] = state
	}
	if state.inflight {
		c.logger.Debug("Not rescheduling already in-flight game", "game", game)
		return nil, nil
	}
	// Create the player separately to the state so we retry creating it if it fails on the first attempt.
	if state.player == nil {
		player, err := c.createPlayer(game, c.disk.DirForGame(game))
		if err != nil {
			return nil, fmt.Errorf("failed to create game player: %w", err)
		}
		state.player = player
	}
	state.inflight = true
	return &job{addr: game, player: state.player}, nil
}

func (c *coordinator) enqueueJob(ctx context.Context, j job) error {
	for {
		select {
		case c.jobQueue <- j:
			return nil
		case result := <-c.resultQueue:
			if err := c.processResult(result); err != nil {
				c.logger.Error("Failed to process result", "err", err)
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (c *coordinator) processResult(j job) error {
	state, ok := c.states[j.addr]
	if !ok {
		return fmt.Errorf("game %v received unexpected result: %w", j.addr, errUnknownGame)
	}
	state.inflight = false
	state.status = j.status
	c.deleteResolvedGameFiles()
	c.m.RecordGameUpdateCompleted()
	return nil
}

func (c *coordinator) deleteResolvedGameFiles() {
	var keepGames []common.Address
	for addr, state := range c.states {
		if state.status == types.GameStatusInProgress || state.inflight {
			keepGames = append(keepGames, addr)
		}
	}
	if err := c.disk.RemoveAllExcept(keepGames); err != nil {
		c.logger.Error("Unable to cleanup game data", "err", err)
	}
}

func newCoordinator(logger log.Logger, m SchedulerMetricer, jobQueue chan<- job, resultQueue <-chan job, createPlayer PlayerCreator, disk DiskManager) *coordinator {
	return &coordinator{
		logger:       logger,
		m:            m,
		jobQueue:     jobQueue,
		resultQueue:  resultQueue,
		createPlayer: createPlayer,
		disk:         disk,
		states:       make(map[common.Address]*gameState),
	}
}
