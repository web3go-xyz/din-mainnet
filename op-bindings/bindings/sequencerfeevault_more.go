// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SequencerFeeVaultStorageLayoutJSON = "{\"storage\":[{\"astId\":3087,\"contract\":\"contracts/L2/SequencerFeeVault.sol:SequencerFeeVault\",\"label\":\"l1FeeWallet\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"}}}"

var SequencerFeeVaultStorageLayout = new(solc.StorageLayout)

var SequencerFeeVaultDeployedBin = "0x6080604052600436106100435760003560e01c80633ccfd60b1461004f57806354fd4d5014610066578063d3e5792b14610091578063d4ff9218146100bb57600080fd5b3661004a57005b600080fd5b34801561005b57600080fd5b5061006461010d565b005b34801561007257600080fd5b5061007b610296565b60405161008891906104f0565b60405180910390f35b34801561009d57600080fd5b506100ad67d02ab486cedc000081565b604051908152602001610088565b3480156100c757600080fd5b506000546100e89073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610088565b67d02ab486cedc00004710156101cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f53657175656e6365724665655661756c743a207769746864726177616c20616d60448201527f6f756e74206d7573742062652067726561746572207468616e206d696e696d7560648201527f6d207769746864726177616c20616d6f756e7400000000000000000000000000608482015260a40160405180910390fd5b600080546040805160208101825283815290517fa3a795480000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000109363a3a795489347936102629373deaddeaddeaddeaddeaddeaddeaddeaddead00009373ffffffffffffffffffffffffffffffffffffffff909316924792909160040161050a565b6000604051808303818588803b15801561027b57600080fd5b505af115801561028f573d6000803e3d6000fd5b5050505050565b60606102c17f0000000000000000000000000000000000000000000000000000000000000000610339565b6102ea7f0000000000000000000000000000000000000000000000000000000000000000610339565b6103137f0000000000000000000000000000000000000000000000000000000000000000610339565b60405160200161032593929190610560565b604051602081830303815290604052905090565b60608160000361037c57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156103a6578061039081610605565b915061039f9050600a8361066c565b9150610380565b60008167ffffffffffffffff8111156103c1576103c1610680565b6040519080825280601f01601f1916602001820160405280156103eb576020820181803683370190505b5090505b841561046e576104006001836106af565b915061040d600a866106c6565b6104189060306106da565b60f81b81838151811061042d5761042d6106f2565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610467600a8661066c565b94506103ef565b949350505050565b60005b83811015610491578181015183820152602001610479565b838111156104a0576000848401525b50505050565b600081518084526104be816020860160208601610476565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061050360208301846104a6565b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525084604083015263ffffffff8416606083015260a0608083015261055560a08301846104a6565b979650505050505050565b60008451610572818460208901610476565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516105ae816001850160208a01610476565b600192019182015283516105c9816002840160208801610476565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610636576106366105d6565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261067b5761067b61063d565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000828210156106c1576106c16105d6565b500390565b6000826106d5576106d561063d565b500690565b600082198211156106ed576106ed6105d6565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SequencerFeeVaultStorageLayoutJSON), SequencerFeeVaultStorageLayout); err != nil {
		panic(err)
	}

	layouts["SequencerFeeVault"] = SequencerFeeVaultStorageLayout
	deployedBytecodes["SequencerFeeVault"] = SequencerFeeVaultDeployedBin
}
