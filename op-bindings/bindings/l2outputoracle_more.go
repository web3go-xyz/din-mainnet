// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2OutputOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"startingBlockNumber\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"startingTimestamp\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1004,\"contract\":\"contracts/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"l2Outputs\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_struct(OutputProposal)9692_storage)dyn_storage\"}],\"types\":{\"t_array(t_struct(OutputProposal)9692_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct Types.OutputProposal[]\",\"numberOfBytes\":\"32\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_struct(OutputProposal)9692_storage\":{\"encoding\":\"inplace\",\"label\":\"struct Types.OutputProposal\",\"numberOfBytes\":\"64\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2OutputOracleStorageLayout = new(solc.StorageLayout)

var L2OutputOracleDeployedBin = "0x6080604052600436106101435760003560e01c806388786272116100c0578063cf8e5cf011610074578063dcec334811610059578063dcec3348146103ce578063e4a30116146103e3578063f4daa2911461040357600080fd5b8063cf8e5cf01461038e578063d1de856c146103ae57600080fd5b80639aaab648116100a55780639aaab648146102eb578063a25ae557146102fe578063bffa7f0f1461035a57600080fd5b806388786272146102b357806389c44cbb146102c957600080fd5b806369f16eec116101175780636b4d98dd116100fc5780636b4d98dd1461022457806370872aa51461027d5780637f0064201461029357600080fd5b806369f16eec146101fa5780636abcf5631461020f57600080fd5b80622134cc146101485780634599c7881461018f578063529933df146101a457806354fd4d50146101d8575b600080fd5b34801561015457600080fd5b5061017c7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b34801561019b57600080fd5b5061017c610437565b3480156101b057600080fd5b5061017c7f000000000000000000000000000000000000000000000000000000000000000081565b3480156101e457600080fd5b506101ed6104aa565b60405161018691906113f2565b34801561020657600080fd5b5061017c61054d565b34801561021b57600080fd5b5060035461017c565b34801561023057600080fd5b506102587f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610186565b34801561028957600080fd5b5061017c60015481565b34801561029f57600080fd5b5061017c6102ae366004611443565b61055f565b3480156102bf57600080fd5b5061017c60025481565b3480156102d557600080fd5b506102e96102e4366004611443565b610778565b005b6102e96102f936600461145c565b610a4e565b34801561030a57600080fd5b5061031e610319366004611443565b610ecd565b60408051825181526020808401516fffffffffffffffffffffffffffffffff908116918301919091529282015190921690820152606001610186565b34801561036657600080fd5b506102587f000000000000000000000000000000000000000000000000000000000000000081565b34801561039a57600080fd5b5061031e6103a9366004611443565b610f61565b3480156103ba57600080fd5b5061017c6103c9366004611443565b610f99565b3480156103da57600080fd5b5061017c610fe7565b3480156103ef57600080fd5b506102e96103fe36600461148e565b61101c565b34801561040f57600080fd5b5061017c7f000000000000000000000000000000000000000000000000000000000000000081565b600354600090156104a15760038054610452906001906114df565b81548110610462576104626114f6565b600091825260209091206002909102016001015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16919050565b6001545b905090565b60606104d57f0000000000000000000000000000000000000000000000000000000000000000611285565b6104fe7f0000000000000000000000000000000000000000000000000000000000000000611285565b6105277f0000000000000000000000000000000000000000000000000000000000000000611285565b60405160200161053993929190611525565b604051602081830303815290604052905090565b6003546000906104a5906001906114df565b6000610569610437565b821115610623576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f7420666f72206120626c6f636b207468617420686173206e6f74206265656e2060648201527f70726f706f736564000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b6003546106d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604660248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f74206173206e6f206f7574707574732068617665206265656e2070726f706f7360648201527f6564207965740000000000000000000000000000000000000000000000000000608482015260a40161061a565b6003546000905b8082101561077157600060026106f5838561159b565b6106ff91906115e2565b90508460038281548110610715576107156114f6565b600091825260209091206002909102016001015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1610156107675761076081600161159b565b925061076b565b8091505b506106df565b5092915050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461083d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4c324f75747075744f7261636c653a206f6e6c7920746865206368616c6c656e60448201527f67657220616464726573732063616e2064656c657465206f7574707574730000606482015260840161061a565b60035481106108f4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f4c324f75747075744f7261636c653a2063616e6e6f742064656c657465206f7560448201527f747075747320616674657220746865206c6174657374206f757470757420696e60648201527f6465780000000000000000000000000000000000000000000000000000000000608482015260a40161061a565b7f000000000000000000000000000000000000000000000000000000000000000060038281548110610928576109286114f6565b6000918252602090912060016002909202010154610958906fffffffffffffffffffffffffffffffff16426114df565b10610a0b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604660248201527f4c324f75747075744f7261636c653a2063616e6e6f742064656c657465206f7560448201527f74707574732074686174206861766520616c7265616479206265656e2066696e60648201527f616c697a65640000000000000000000000000000000000000000000000000000608482015260a40161061a565b6000610a1660035490565b90508160035581817f4ee37ac2c786ec85e87592d3c5c8a1dd66f8496dda3f125d9ea8ca5f657629b660405160405180910390a35050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610b39576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f4c324f75747075744f7261636c653a206f6e6c79207468652070726f706f736560448201527f7220616464726573732063616e2070726f706f7365206e6577206f757470757460648201527f7300000000000000000000000000000000000000000000000000000000000000608482015260a40161061a565b610b41610fe7565b8314610bf5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f4c324f75747075744f7261636c653a20626c6f636b206e756d626572206d757360448201527f7420626520657175616c20746f206e65787420657870656374656420626c6f6360648201527f6b206e756d626572000000000000000000000000000000000000000000000000608482015260a40161061a565b42610bff84610f99565b10610c8c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f4c324f75747075744f7261636c653a2063616e6e6f742070726f706f7365204c60448201527f32206f757470757420696e207468652066757475726500000000000000000000606482015260840161061a565b83610d19576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4c324f75747075744f7261636c653a204c32206f75747075742070726f706f7360448201527f616c2063616e6e6f7420626520746865207a65726f2068617368000000000000606482015260840161061a565b8115610dd55781814014610dd5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604960248201527f4c324f75747075744f7261636c653a20626c6f636b206861736820646f65732060448201527f6e6f74206d61746368207468652068617368206174207468652065787065637460648201527f6564206865696768740000000000000000000000000000000000000000000000608482015260a40161061a565b82610ddf60035490565b857fa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e242604051610e1191815260200190565b60405180910390a45050604080516060810182529283526fffffffffffffffffffffffffffffffff4281166020850190815292811691840191825260038054600181018255600091909152935160029094027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b810194909455915190518216700100000000000000000000000000000000029116177fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c90910155565b604080516060810182526000808252602082018190529181019190915260038281548110610efd57610efd6114f6565b600091825260209182902060408051606081018252600290930290910180548352600101546fffffffffffffffffffffffffffffffff8082169484019490945270010000000000000000000000000000000090049092169181019190915292915050565b60408051606081018252600080825260208201819052918101919091526003610f898361055f565b81548110610efd57610efd6114f6565b60007f000000000000000000000000000000000000000000000000000000000000000060015483610fca91906114df565b610fd491906115f6565b600254610fe1919061159b565b92915050565b60007f0000000000000000000000000000000000000000000000000000000000000000611012610437565b6104a5919061159b565b600054610100900460ff161580801561103c5750600054600160ff909116105b806110565750303b158015611056575060005460ff166001145b6110e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161061a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561114057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b428211156111f7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526044602482018190527f4c324f75747075744f7261636c653a207374617274696e67204c322074696d65908201527f7374616d70206d757374206265206c657373207468616e2063757272656e742060648201527f74696d6500000000000000000000000000000000000000000000000000000000608482015260a40161061a565b60028290556001839055801561126457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6060816000036112c857505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156112f257806112dc81611633565b91506112eb9050600a836115e2565b91506112cc565b60008167ffffffffffffffff81111561130d5761130d61166b565b6040519080825280601f01601f191660200182016040528015611337576020820181803683370190505b5090505b84156113ba5761134c6001836114df565b9150611359600a8661169a565b61136490603061159b565b60f81b818381518110611379576113796114f6565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506113b3600a866115e2565b945061133b565b949350505050565b60005b838110156113dd5781810151838201526020016113c5565b838111156113ec576000848401525b50505050565b60208152600082518060208401526114118160408501602087016113c2565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60006020828403121561145557600080fd5b5035919050565b6000806000806080858703121561147257600080fd5b5050823594602084013594506040840135936060013592509050565b600080604083850312156114a157600080fd5b50508035926020909101359150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156114f1576114f16114b0565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600084516115378184602089016113c2565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611573816001850160208a016113c2565b6001920191820152835161158e8160028401602088016113c2565b0160020195945050505050565b600082198211156115ae576115ae6114b0565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826115f1576115f16115b3565b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561162e5761162e6114b0565b500290565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611664576116646114b0565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000826116a9576116a96115b3565b50069056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2OutputOracleStorageLayoutJSON), L2OutputOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2OutputOracle"] = L2OutputOracleStorageLayout
	deployedBytecodes["L2OutputOracle"] = L2OutputOracleDeployedBin
}
