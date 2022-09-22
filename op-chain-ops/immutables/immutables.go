package immutables

import (
	"fmt"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-chain-ops/deployer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ImmutableValues represents the values to be set in immutable code.
// The key is the name of the variable and the value is the value to set in
// immutable code.
type ImmutableValues map[string]any

// ImmutableConfig represents the immutable configuration for the L2 predeploy
// contracts.
type ImmutableConfig map[string]ImmutableValues

// DeploymentResults represents the output of deploying each of the
// contracts so that the immutables can be set properly in the bytecode.
type DeploymentResults map[string]hexutil.Bytes

// BuildOptimism will deploy the L2 predeploys so that their immutables are set
// correctly.
func BuildOptimism(immutable ImmutableConfig) (DeploymentResults, error) {
	deployments := []deployer.Constructor{
		{
			Name: "GasPriceOracle",
		},
		{
			Name: "L1Block",
		},
		{
			Name: "L2CrossDomainMessenger",
			Args: []interface{}{
				immutable["L2CrossDomainMessenger"]["otherMessenger"],
			},
		},
		{
			Name: "L2StandardBridge",
			Args: []interface{}{
				immutable["L2StandardBridge"]["otherBridge"],
			},
		},
		{
			Name: "L2ToL1MessagePasser",
		},
		{
			Name: "SequencerFeeVault",
		},
		{
			Name: "OptimismMintableERC20Factory",
		},
		{
			Name: "DeployerWhitelist",
		},
		{
			Name: "LegacyMessagePasser",
		},
		{
			Name: "L1BlockNumber",
		},
	}
	return BuildL2(deployments)
}

// BuildL2 will deploy contracts to a simulated backend so that their immutables
// can be properly set. The bytecode returned in the results is suitable to be
// inserted into the state via state surgery.
func BuildL2(constructors []deployer.Constructor) (DeploymentResults, error) {
	deployments, err := deployer.Deploy(deployer.NewBackend(), constructors, l2Deployer)
	if err != nil {
		return nil, err
	}
	results := make(DeploymentResults)
	for _, dep := range deployments {
		results[dep.Name] = dep.Bytecode
	}
	return results, nil
}

func l2Deployer(backend *backends.SimulatedBackend, opts *bind.TransactOpts, deployment deployer.Constructor) (common.Address, error) {
	var addr common.Address
	var err error
	switch deployment.Name {
	case "GasPriceOracle":
		// The owner of the gas price oracle is not immutable, not required
		// to be set here. It cannot be `address(0)`
		owner := common.Address{1}
		addr, _, _, err = bindings.DeployGasPriceOracle(opts, backend, owner)
	case "L1Block":
		// No arguments required for the L1Block contract
		addr, _, _, err = bindings.DeployL1Block(opts, backend)
	case "L2CrossDomainMessenger":
		otherMessenger, ok := deployment.Args[0].(common.Address)
		if !ok {
			return common.Address{}, fmt.Errorf("invalid type for otherMessenger")
		}
		addr, _, _, err = bindings.DeployL2CrossDomainMessenger(opts, backend, otherMessenger)
	case "L2StandardBridge":
		otherBridge, ok := deployment.Args[0].(common.Address)
		if !ok {
			return common.Address{}, fmt.Errorf("invalid type for otherBridge")
		}
		addr, _, _, err = bindings.DeployL2StandardBridge(opts, backend, otherBridge)
	case "L2ToL1MessagePasser":
		// No arguments required for L2ToL1MessagePasser
		addr, _, _, err = bindings.DeployL2ToL1MessagePasser(opts, backend)
	case "SequencerFeeVault":
		// No arguments to SequencerFeeVault
		addr, _, _, err = bindings.DeploySequencerFeeVault(opts, backend)
	case "OptimismMintableERC20Factory":
		addr, _, _, err = bindings.DeployOptimismMintableERC20Factory(opts, backend, predeploys.L2StandardBridgeAddr)
	case "DeployerWhitelist":
		addr, _, _, err = bindings.DeployDeployerWhitelist(opts, backend)
	case "LegacyMessagePasser":
		addr, _, _, err = bindings.DeployLegacyMessagePasser(opts, backend)
	case "L1BlockNumber":
		addr, _, _, err = bindings.DeployL1BlockNumber(opts, backend)
	default:
		return addr, fmt.Errorf("unknown contract: %s", deployment.Name)
	}

	return addr, err
}
