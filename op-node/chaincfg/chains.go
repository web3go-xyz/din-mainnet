package chaincfg

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

var Mainnet = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0x438335a20d98863a4c0c97999eb2481921ccd28553eac6f913af7c12aec04108"),
			Number: 17422590,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0xdbf6a80fef073de06add9b0d14026d6e5a86c85f6d102c36d3d8e9cf89c2afd3"),
			Number: 105235063,
		},
		L2Time: 1686068903,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0x6887246668a3b87f54deb3b94ba47a6f63f32985"),
			Overhead:    eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000bc")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000a6fe0")),
			GasLimit:    30_000_000,
		},
	},
	BlockTime:              2,
	MaxSequencerDrift:      600,
	SeqWindowSize:          3600,
	ChannelTimeout:         300,
	L1ChainID:              big.NewInt(1),
	L2ChainID:              big.NewInt(10),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000000010"),
	DepositContractAddress: common.HexToAddress("0xbEb5Fc579115071764c7423A4f12eDde41f106Ed"),
	L1SystemConfigAddress:  common.HexToAddress("0x229047fed2591dbec1eF1118d64F7aF3dB9EB290"),
	RegolithTime:           u64Ptr(0),
}

var Goerli = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0x6ffc1bf3754c01f6bb9fe057c1578b87a8571ce2e9be5ca14bace6eccfd336c7"),
			Number: 8300214,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0x0f783549ea4313b784eadd9b8e8a69913b368b7366363ea814d7707ac505175f"),
			Number: 4061224,
		},
		L2Time: 1673550516,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0x7431310e026B69BFC676C0013E12A1A11411EEc9"),
			Overhead:    eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000f4240")),
			GasLimit:    25_000_000,
		},
	},
	BlockTime:              2,
	MaxSequencerDrift:      600,
	SeqWindowSize:          3600,
	ChannelTimeout:         300,
	L1ChainID:              big.NewInt(5),
	L2ChainID:              big.NewInt(420),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000000420"),
	DepositContractAddress: common.HexToAddress("0x5b47E1A08Ea6d985D6649300584e6722Ec4B1383"),
	L1SystemConfigAddress:  common.HexToAddress("0xAe851f927Ee40dE99aaBb7461C00f9622ab91d60"),
	RegolithTime:           u64Ptr(1679079600),
}

var Sepolia = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0x70e5634d09793b1cfaa7d0a2a5d3289a3b2308de1e82f682b4f817fc670f9797"),
			Number: 3976708,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0xfbfc64b34d705b0eb83ab8b2206c0da90a76e1ae54ae657c8cfbee0e802a9120"),
			Number: 0,
		},
		L2Time: 1690493568,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0x7431310e026b69bfc676c0013e12a1a11411eec9"),
			Overhead:    eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000bc")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000a6fe0")),
			GasLimit:    30000000,
		},
	},
	BlockTime:              2,
	MaxSequencerDrift:      600,
	SeqWindowSize:          3600,
	ChannelTimeout:         300,
	L1ChainID:              big.NewInt(11155111),
	L2ChainID:              big.NewInt(11155420),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000011155420"),
	DepositContractAddress: common.HexToAddress("0x8f6452d842438c4e22ba18baa21652ff65530df4"),
	L1SystemConfigAddress:  common.HexToAddress("0xf425ed544d2e1f1b7a8650d5897a7ccf43020791"),
	RegolithTime:           u64Ptr(0),
}

var NetworksByName = map[string]rollup.Config{
	"goerli":  Goerli,
	"mainnet": Mainnet,
	"sepolia": Sepolia,
}

var L2ChainIDToNetworkName = func() map[string]string {
	out := make(map[string]string)
	for name, netCfg := range NetworksByName {
		out[netCfg.L2ChainID.String()] = name
	}
	return out
}()

func AvailableNetworks() []string {
	var networks []string
	for name := range NetworksByName {
		networks = append(networks, name)
	}
	return networks
}

func GetRollupConfig(name string) (rollup.Config, error) {
	network, ok := NetworksByName[name]
	if !ok {
		return rollup.Config{}, fmt.Errorf("invalid network %s", name)
	}

	return network, nil
}

func u64Ptr(v uint64) *uint64 {
	return &v
}
