package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/chaincfg"
	"github.com/ethereum-optimism/optimism/op-node/sources"
	"github.com/ethereum-optimism/optimism/op-program/host/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

// Use HexToHash(...).Hex() to ensure the strings are the correct length for a hash
var l1HeadValue = common.HexToHash("0x111111").Hex()
var l2HeadValue = common.HexToHash("0x222222").Hex()
var l2ClaimValue = common.HexToHash("0x333333").Hex()

func TestLogLevel(t *testing.T) {
	t.Run("RejectInvalid", func(t *testing.T) {
		verifyArgsInvalid(t, "unknown level: foo", addRequiredArgs("--log.level=foo"))
	})

	for _, lvl := range []string{"trace", "debug", "info", "error", "crit"} {
		lvl := lvl
		t.Run("AcceptValid_"+lvl, func(t *testing.T) {
			logger, _, err := runWithArgs(addRequiredArgs("--log.level", lvl))
			require.NoError(t, err)
			require.NotNil(t, logger)
		})
	}
}

func TestDefaultCLIOptionsMatchDefaultConfig(t *testing.T) {
	cfg := configForArgs(t, addRequiredArgs())
	defaultCfg := config.NewConfig(
		&chaincfg.Goerli,
		"genesis.json",
		common.HexToHash(l1HeadValue),
		common.HexToHash(l2HeadValue),
		common.HexToHash(l2ClaimValue))
	require.Equal(t, defaultCfg, cfg)
}

func TestNetwork(t *testing.T) {
	t.Run("Unknown", func(t *testing.T) {
		verifyArgsInvalid(t, "invalid network bar", replaceRequiredArg("--network", "bar"))
	})

	t.Run("Required", func(t *testing.T) {
		verifyArgsInvalid(t, "flag rollup.config or network is required", addRequiredArgsExcept("--network"))
	})

	t.Run("DisallowNetworkAndRollupConfig", func(t *testing.T) {
		verifyArgsInvalid(t, "cannot specify both rollup.config and network", addRequiredArgs("--rollup.config=foo"))
	})

	t.Run("RollupConfig", func(t *testing.T) {
		dir := t.TempDir()
		configJson, err := json.Marshal(chaincfg.Goerli)
		require.NoError(t, err)
		configFile := dir + "/config.json"
		err = os.WriteFile(configFile, configJson, os.ModePerm)
		require.NoError(t, err)

		cfg := configForArgs(t, addRequiredArgsExcept("--network", "--rollup.config", configFile))
		require.Equal(t, chaincfg.Goerli, *cfg.Rollup)
	})

	for name, cfg := range chaincfg.NetworksByName {
		name := name
		expected := cfg
		t.Run("Network_"+name, func(t *testing.T) {
			cfg := configForArgs(t, replaceRequiredArg("--network", name))
			require.Equal(t, expected, *cfg.Rollup)
		})
	}
}

func TestL2(t *testing.T) {
	expected := "https://example.com:8545"
	cfg := configForArgs(t, addRequiredArgs("--l2", expected))
	require.Equal(t, expected, cfg.L2URL)
}

func TestL2Genesis(t *testing.T) {
	t.Run("Required", func(t *testing.T) {
		verifyArgsInvalid(t, "flag l2.genesis is required", addRequiredArgsExcept("--l2.genesis"))
	})

	t.Run("Valid", func(t *testing.T) {
		cfg := configForArgs(t, replaceRequiredArg("--l2.genesis", "/tmp/genesis.json"))
		require.Equal(t, "/tmp/genesis.json", cfg.L2GenesisPath)
	})
}

func TestL2Head(t *testing.T) {
	t.Run("Required", func(t *testing.T) {
		verifyArgsInvalid(t, "flag l2.head is required", addRequiredArgsExcept("--l2.head"))
	})

	t.Run("Valid", func(t *testing.T) {
		cfg := configForArgs(t, replaceRequiredArg("--l2.head", l2HeadValue))
		require.Equal(t, common.HexToHash(l2HeadValue), cfg.L2Head)
	})

	t.Run("Invalid", func(t *testing.T) {
		verifyArgsInvalid(t, config.ErrInvalidL2Head.Error(), replaceRequiredArg("--l2.head", "something"))
	})
}

func TestL1Head(t *testing.T) {
	t.Run("Required", func(t *testing.T) {
		verifyArgsInvalid(t, "flag l1.head is required", addRequiredArgsExcept("--l1.head"))
	})

	t.Run("Valid", func(t *testing.T) {
		cfg := configForArgs(t, replaceRequiredArg("--l1.head", l1HeadValue))
		require.Equal(t, common.HexToHash(l1HeadValue), cfg.L1Head)
	})

	t.Run("Invalid", func(t *testing.T) {
		verifyArgsInvalid(t, config.ErrInvalidL1Head.Error(), replaceRequiredArg("--l1.head", "something"))
	})
}

func TestL1(t *testing.T) {
	expected := "https://example.com:8545"
	cfg := configForArgs(t, addRequiredArgs("--l1", expected))
	require.Equal(t, expected, cfg.L1URL)
}

func TestL1TrustRPC(t *testing.T) {
	t.Run("DefaultFalse", func(t *testing.T) {
		cfg := configForArgs(t, addRequiredArgs())
		require.False(t, cfg.L1TrustRPC)
	})
	t.Run("Enabled", func(t *testing.T) {
		cfg := configForArgs(t, addRequiredArgs("--l1.trustrpc"))
		require.True(t, cfg.L1TrustRPC)
	})
	t.Run("EnabledWithArg", func(t *testing.T) {
		cfg := configForArgs(t, addRequiredArgs("--l1.trustrpc=true"))
		require.True(t, cfg.L1TrustRPC)
	})
	t.Run("Disabled", func(t *testing.T) {
		cfg := configForArgs(t, addRequiredArgs("--l1.trustrpc=false"))
		require.False(t, cfg.L1TrustRPC)
	})
}

func TestL1RPCKind(t *testing.T) {
	t.Run("DefaultBasic", func(t *testing.T) {
		cfg := configForArgs(t, addRequiredArgs())
		require.Equal(t, sources.RPCKindBasic, cfg.L1RPCKind)
	})
	for _, kind := range sources.RPCProviderKinds {
		t.Run(kind.String(), func(t *testing.T) {
			cfg := configForArgs(t, addRequiredArgs("--l1.rpckind", kind.String()))
			require.Equal(t, kind, cfg.L1RPCKind)
		})
	}
	t.Run("RequireLowercase", func(t *testing.T) {
		verifyArgsInvalid(t, "rpc kind", addRequiredArgs("--l1.rpckind", "AlChemY"))
	})
	t.Run("UnknownKind", func(t *testing.T) {
		verifyArgsInvalid(t, "\"foo\"", addRequiredArgs("--l1.rpckind", "foo"))
	})
}

// Offline support will be added later, but for now it just bails out with an error
func TestOfflineModeNotSupported(t *testing.T) {
	logger := log.New()
	cfg := config.NewConfig(&chaincfg.Goerli, "genesis.json", common.HexToHash(l1HeadValue), common.HexToHash(l2HeadValue), common.HexToHash(l2ClaimValue))
	err := FaultProofProgram(logger, cfg)
	require.ErrorContains(t, err, "offline mode not supported")
}

func TestL2Claim(t *testing.T) {
	t.Run("Required", func(t *testing.T) {
		verifyArgsInvalid(t, "flag l2.claim is required", addRequiredArgsExcept("--l2.claim"))
	})

	t.Run("Valid", func(t *testing.T) {
		cfg := configForArgs(t, replaceRequiredArg("--l2.claim", l2ClaimValue))
		require.EqualValues(t, common.HexToHash(l2ClaimValue), cfg.L2Claim)
	})

	t.Run("Invalid", func(t *testing.T) {
		verifyArgsInvalid(t, config.ErrInvalidL2Claim.Error(), replaceRequiredArg("--l2.claim", "something"))
	})
}

func verifyArgsInvalid(t *testing.T, messageContains string, cliArgs []string) {
	_, _, err := runWithArgs(cliArgs)
	require.ErrorContains(t, err, messageContains)
}

func configForArgs(t *testing.T, cliArgs []string) *config.Config {
	_, cfg, err := runWithArgs(cliArgs)
	require.NoError(t, err)
	return cfg
}

func runWithArgs(cliArgs []string) (log.Logger, *config.Config, error) {
	var cfg *config.Config
	var logger log.Logger
	fullArgs := append([]string{"op-program"}, cliArgs...)
	err := run(fullArgs, func(log log.Logger, config *config.Config) error {
		logger = log
		cfg = config
		return nil
	})
	return logger, cfg, err
}

func addRequiredArgs(args ...string) []string {
	req := requiredArgs()
	combined := toArgList(req)
	return append(combined, args...)
}

func addRequiredArgsExcept(name string, optionalArgs ...string) []string {
	req := requiredArgs()
	delete(req, name)
	return append(toArgList(req), optionalArgs...)
}

func replaceRequiredArg(name string, value string) []string {
	req := requiredArgs()
	req[name] = value
	return toArgList(req)
}

// requiredArgs returns map of argument names to values which are the minimal arguments required
// to create a valid Config
func requiredArgs() map[string]string {
	return map[string]string{
		"--network":    "goerli",
		"--l1.head":    l1HeadValue,
		"--l2.head":    l2HeadValue,
		"--l2.claim":   l2ClaimValue,
		"--l2.genesis": "genesis.json",
	}
}

func toArgList(req map[string]string) []string {
	var combined []string
	for name, value := range req {
		combined = append(combined, name)
		combined = append(combined, value)
	}
	return combined
}
