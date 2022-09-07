package genesis

import (
	"github.com/ethereum-optimism/optimism/op-chain-ops/hardhat"
	"github.com/ethereum-optimism/optimism/op-chain-ops/state"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/core"
)

type L2Addresses struct {
	ProxyAdmin                  common.Address
	L1StandardBridgeProxy       common.Address
	L1CrossDomainMessengerProxy common.Address
}

// BuildL2DeveloperGenesis will build the developer Optimism Genesis
// Block. Suitable for devnets.
func BuildL2DeveloperGenesis(hh *hardhat.Hardhat, config *DeployConfig, l1StartBlock *types.Block, l2Addrs *L2Addresses) (*core.Genesis, error) {
	genspec, err := NewL2Genesis(config, l1StartBlock)
	if err != nil {
		return nil, err
	}

	db := state.NewMemoryStateDB(genspec)

	FundDevAccounts(db)
	SetPrecompileBalances(db)

	return BuildL2Genesis(db, hh, config, l1StartBlock, l2Addrs)
}

// BuildL2Genesis will build the L2 Optimism Genesis Block
func BuildL2Genesis(db *state.MemoryStateDB, hh *hardhat.Hardhat, config *DeployConfig, l1Block *types.Block, l2Addrs *L2Addresses) (*core.Genesis, error) {
	// TODO(tynes): need a function for clearing old, unused storage slots.
	// Each deployed contract on L2 needs to have its existing storage
	// inspected and then cleared if they are no longer used.

	if err := SetL2Proxies(hh, db, l2Addrs.ProxyAdmin); err != nil {
		return nil, err
	}

	storage, err := NewL2StorageConfig(
		config,
		l1Block,
		l2Addrs.L1StandardBridgeProxy,
		l2Addrs.L1CrossDomainMessengerProxy,
	)
	if err != nil {
		return nil, err
	}

	if err := SetImplementations(hh, db, storage); err != nil {
		return nil, err
	}

	if err := MigrateDepositHashes(hh, db); err != nil {
		return nil, err
	}

	return db.Genesis(), nil
}
