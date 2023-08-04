package cannon

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/ethereum-optimism/optimism/op-challenger/config"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum/go-ethereum/log"
)

const (
	snapsDir     = "snapshots"
	preimagesDir = "snapshots"
)

var snapshotNameRegexp = regexp.MustCompile(`^[0-9]+\.json$`)

type snapshotSelect func(logger log.Logger, dir string, absolutePreState string, i uint64) (string, error)
type cmdExecutor func(ctx context.Context, l log.Logger, binary string, args ...string) error

type Executor struct {
	logger           log.Logger
	l1               string
	l2               string
	inputs           localGameInputs
	cannon           string
	server           string
	absolutePreState string
	dataDir          string
	snapshotFreq     uint
	selectSnapshot   snapshotSelect
	cmdExecutor      cmdExecutor
}

func NewExecutor(logger log.Logger, cfg *config.Config, inputs localGameInputs) *Executor {
	return &Executor{
		logger:           logger,
		l1:               cfg.L1EthRpc,
		l2:               cfg.CannonL2,
		inputs:           inputs,
		cannon:           cfg.CannonBin,
		server:           cfg.CannonServer,
		absolutePreState: cfg.CannonAbsolutePreState,
		dataDir:          cfg.CannonDatadir,
		snapshotFreq:     cfg.CannonSnapshotFreq,
		selectSnapshot:   findStartingSnapshot,
		cmdExecutor:      runCmd,
	}
}

func (e *Executor) GenerateProof(ctx context.Context, dir string, i uint64) error {
	start, err := e.selectSnapshot(e.logger, filepath.Join(e.dataDir, snapsDir), e.absolutePreState, i)
	if err != nil {
		return fmt.Errorf("find starting snapshot: %w", err)
	}
	args := []string{
		"run",
		"--input", start,
		"--proof-at", "=" + strconv.FormatUint(i, 10),
		"--stop-at", "=" + strconv.FormatUint(i+1, 10),
		"--proof-fmt", filepath.Join(dir, proofsDir, "%d.json"),
		"--snapshot-at", "%" + strconv.FormatUint(uint64(e.snapshotFreq), 10),
		"--snapshot-fmt", filepath.Join(e.dataDir, snapsDir, "%d.json"),
		"--",
		e.server,
		"--l1", e.l1,
		"--l2", e.l2,
		"--datadir", filepath.Join(e.dataDir, preimagesDir),
		"--l1.head", e.inputs.l1Head.Hex(),
		"--l2.head", e.inputs.l2Head.Hex(),
		"--l2.outputroot", e.inputs.l2OutputRoot.Hex(),
		"--l2.claim", e.inputs.l2Claim.Hex(),
		"--l2.blocknumber", e.inputs.l2BlockNumber.Text(10),
		"--l2.chainid", e.inputs.l2ChainId.Text(10),
	}

	e.logger.Info("Generating trace", "proof", i, "cmd", e.cannon, "args", args)
	return e.cmdExecutor(ctx, e.logger.New("proof", i), e.cannon, args...)
}

func runCmd(ctx context.Context, l log.Logger, binary string, args ...string) error {
	cmd := exec.CommandContext(ctx, binary, args...)
	stdOut := oplog.NewWriter(l, log.LvlInfo)
	defer stdOut.Close()
	stdErr := oplog.NewWriter(l, log.LvlError)
	defer stdErr.Close()
	cmd.Stdout = stdOut
	cmd.Stderr = stdErr
	return cmd.Run()
}

// findStartingSnapshot finds the closest snapshot before the specified traceIndex in snapDir.
// If no suitable snapshot can be found it returns absolutePreState.
func findStartingSnapshot(logger log.Logger, snapDir string, absolutePreState string, traceIndex uint64) (string, error) {
	// Find the closest snapshot to start from
	entries, err := os.ReadDir(snapDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return absolutePreState, nil
		}
		return "", fmt.Errorf("list snapshots in %v: %w", snapDir, err)
	}
	bestSnap := uint64(0)
	for _, entry := range entries {
		if entry.IsDir() {
			logger.Warn("Unexpected directory in snapshots dir: %v/%v", snapDir, entry.Name())
			continue
		}
		name := entry.Name()
		if !snapshotNameRegexp.MatchString(name) {
			logger.Warn("Unexpected file in snapshots dir: %v/%v", snapDir, entry.Name())
			continue
		}
		index, err := strconv.ParseUint(name[0:len(name)-len(".json")], 10, 64)
		if err != nil {
			logger.Error("Unable to parse trace index of snapshot file: %v/%v", snapDir, entry.Name())
			continue
		}
		if index > bestSnap && index < traceIndex {
			bestSnap = index
		}
	}
	if bestSnap == 0 {
		return absolutePreState, nil
	}
	startFrom := fmt.Sprintf("%v/%v.json", snapDir, bestSnap)

	return startFrom, nil
}
