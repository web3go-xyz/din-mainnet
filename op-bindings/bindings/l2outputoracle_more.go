// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2OutputOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"startingBlockNumber\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"src/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"startingTimestamp\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1004,\"contract\":\"src/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"l2Outputs\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_struct(OutputProposal)1007_storage)dyn_storage\"},{\"astId\":1005,\"contract\":\"src/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"challenger\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_address\"},{\"astId\":1006,\"contract\":\"src/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"proposer\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_struct(OutputProposal)1007_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct Types.OutputProposal[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(OutputProposal)1007_storage\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_struct(OutputProposal)1007_storage\":{\"encoding\":\"inplace\",\"label\":\"struct Types.OutputProposal\",\"numberOfBytes\":\"64\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2OutputOracleStorageLayout = new(solc.StorageLayout)

var L2OutputOracleDeployedBin = "0x60806040526004361061018a5760003560e01c806389c44cbb116100d6578063ce5db8d61161007f578063dcec334811610059578063dcec334814610517578063e1a41bcf1461052c578063f4daa2911461055f57600080fd5b8063ce5db8d6146104a4578063cf8e5cf0146104d7578063d1de856c146104f757600080fd5b8063a25ae557116100b0578063a25ae557146103f0578063a8e4fb901461044c578063bffa7f0f1461047957600080fd5b806389c44cbb1461038a57806393991af3146103aa5780639aaab648146103dd57600080fd5b806369f16eec1161013857806370872aa51161011257806370872aa51461033e5780637f00642014610354578063887862721461037457600080fd5b806369f16eec146102e95780636abcf563146102fe5780636b4d98dd1461031357600080fd5b8063529933df11610169578063529933df1461020d578063534db0e21461024157806354fd4d501461029357600080fd5b80622134cc1461018f578063019e2729146101d65780634599c788146101f8575b600080fd5b34801561019b57600080fd5b506101c37f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b3480156101e257600080fd5b506101f66101f136600461131c565b610593565b005b34801561020457600080fd5b506101c36107f2565b34801561021957600080fd5b506101c37f000000000000000000000000000000000000000000000000000000000000000081565b34801561024d57600080fd5b5060045461026e9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101cd565b34801561029f57600080fd5b506102dc6040518060400160405280600581526020017f312e352e3100000000000000000000000000000000000000000000000000000081525081565b6040516101cd9190611362565b3480156102f557600080fd5b506101c3610865565b34801561030a57600080fd5b506003546101c3565b34801561031f57600080fd5b5060045473ffffffffffffffffffffffffffffffffffffffff1661026e565b34801561034a57600080fd5b506101c360015481565b34801561036057600080fd5b506101c361036f3660046113d5565b610877565b34801561038057600080fd5b506101c360025481565b34801561039657600080fd5b506101f66103a53660046113d5565b610a8b565b3480156103b657600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006101c3565b6101f66103eb3660046113ee565b610d43565b3480156103fc57600080fd5b5061041061040b3660046113d5565b6111a4565b60408051825181526020808401516fffffffffffffffffffffffffffffffff9081169183019190915292820151909216908201526060016101cd565b34801561045857600080fd5b5060055461026e9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561048557600080fd5b5060055473ffffffffffffffffffffffffffffffffffffffff1661026e565b3480156104b057600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006101c3565b3480156104e357600080fd5b506104106104f23660046113d5565b611238565b34801561050357600080fd5b506101c36105123660046113d5565b611270565b34801561052357600080fd5b506101c36112be565b34801561053857600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006101c3565b34801561056b57600080fd5b506101c37f000000000000000000000000000000000000000000000000000000000000000081565b600054600290610100900460ff161580156105b5575060005460ff8083169116105b610646576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff8316176101001790554284111561072e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526044602482018190527f4c324f75747075744f7261636c653a207374617274696e67204c322074696d65908201527f7374616d70206d757374206265206c657373207468616e2063757272656e742060648201527f74696d6500000000000000000000000000000000000000000000000000000000608482015260a40161063d565b600284905560018590556005805473ffffffffffffffffffffffffffffffffffffffff8581167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556004805492851692909116919091179055600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050505050565b6003546000901561085c576003805461080d9060019061144f565b8154811061081d5761081d611466565b600091825260209091206002909102016001015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16919050565b6001545b905090565b6003546000906108609060019061144f565b60006108816107f2565b821115610936576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f7420666f72206120626c6f636b207468617420686173206e6f74206265656e2060648201527f70726f706f736564000000000000000000000000000000000000000000000000608482015260a40161063d565b6003546109eb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604660248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f74206173206e6f206f7574707574732068617665206265656e2070726f706f7360648201527f6564207965740000000000000000000000000000000000000000000000000000608482015260a40161063d565b6003546000905b80821015610a845760006002610a088385611495565b610a1291906114ad565b90508460038281548110610a2857610a28611466565b600091825260209091206002909102016001015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff161015610a7a57610a73816001611495565b9250610a7e565b8091505b506109f2565b5092915050565b60045473ffffffffffffffffffffffffffffffffffffffff163314610b32576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4c324f75747075744f7261636c653a206f6e6c7920746865206368616c6c656e60448201527f67657220616464726573732063616e2064656c657465206f7574707574730000606482015260840161063d565b6003548110610be9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f4c324f75747075744f7261636c653a2063616e6e6f742064656c657465206f7560448201527f747075747320616674657220746865206c6174657374206f757470757420696e60648201527f6465780000000000000000000000000000000000000000000000000000000000608482015260a40161063d565b7f000000000000000000000000000000000000000000000000000000000000000060038281548110610c1d57610c1d611466565b6000918252602090912060016002909202010154610c4d906fffffffffffffffffffffffffffffffff164261144f565b10610d00576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604660248201527f4c324f75747075744f7261636c653a2063616e6e6f742064656c657465206f7560448201527f74707574732074686174206861766520616c7265616479206265656e2066696e60648201527f616c697a65640000000000000000000000000000000000000000000000000000608482015260a40161063d565b6000610d0b60035490565b90508160035581817f4ee37ac2c786ec85e87592d3c5c8a1dd66f8496dda3f125d9ea8ca5f657629b660405160405180910390a35050565b60055473ffffffffffffffffffffffffffffffffffffffff163314610e10576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f4c324f75747075744f7261636c653a206f6e6c79207468652070726f706f736560448201527f7220616464726573732063616e2070726f706f7365206e6577206f757470757460648201527f7300000000000000000000000000000000000000000000000000000000000000608482015260a40161063d565b610e186112be565b8314610ecc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f4c324f75747075744f7261636c653a20626c6f636b206e756d626572206d757360448201527f7420626520657175616c20746f206e65787420657870656374656420626c6f6360648201527f6b206e756d626572000000000000000000000000000000000000000000000000608482015260a40161063d565b42610ed684611270565b10610f63576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f4c324f75747075744f7261636c653a2063616e6e6f742070726f706f7365204c60448201527f32206f757470757420696e207468652066757475726500000000000000000000606482015260840161063d565b83610ff0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4c324f75747075744f7261636c653a204c32206f75747075742070726f706f7360448201527f616c2063616e6e6f7420626520746865207a65726f2068617368000000000000606482015260840161063d565b81156110ac57818140146110ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604960248201527f4c324f75747075744f7261636c653a20626c6f636b206861736820646f65732060448201527f6e6f74206d61746368207468652068617368206174207468652065787065637460648201527f6564206865696768740000000000000000000000000000000000000000000000608482015260a40161063d565b826110b660035490565b857fa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2426040516110e891815260200190565b60405180910390a45050604080516060810182529283526fffffffffffffffffffffffffffffffff4281166020850190815292811691840191825260038054600181018255600091909152935160029094027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b810194909455915190518216700100000000000000000000000000000000029116177fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c90910155565b6040805160608101825260008082526020820181905291810191909152600382815481106111d4576111d4611466565b600091825260209182902060408051606081018252600290930290910180548352600101546fffffffffffffffffffffffffffffffff8082169484019490945270010000000000000000000000000000000090049092169181019190915292915050565b6040805160608101825260008082526020820181905291810191909152600361126083610877565b815481106111d4576111d4611466565b60007f0000000000000000000000000000000000000000000000000000000000000000600154836112a1919061144f565b6112ab91906114e8565b6002546112b89190611495565b92915050565b60007f00000000000000000000000000000000000000000000000000000000000000006112e96107f2565b6108609190611495565b803573ffffffffffffffffffffffffffffffffffffffff8116811461131757600080fd5b919050565b6000806000806080858703121561133257600080fd5b8435935060208501359250611349604086016112f3565b9150611357606086016112f3565b905092959194509250565b600060208083528351808285015260005b8181101561138f57858101830151858201604001528201611373565b818111156113a1576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b6000602082840312156113e757600080fd5b5035919050565b6000806000806080858703121561140457600080fd5b5050823594602084013594506040840135936060013592509050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561146157611461611420565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082198211156114a8576114a8611420565b500190565b6000826114e3577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561152057611520611420565b50029056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2OutputOracleStorageLayoutJSON), L2OutputOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2OutputOracle"] = L2OutputOracleStorageLayout
	deployedBytecodes["L2OutputOracle"] = L2OutputOracleDeployedBin
}
