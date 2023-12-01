// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2OutputOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"startingBlockNumber\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"src/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"startingTimestamp\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1004,\"contract\":\"src/L1/L2OutputOracle.sol:L2OutputOracle\",\"label\":\"l2Outputs\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_struct(OutputProposal)1005_storage)dyn_storage\"}],\"types\":{\"t_array(t_struct(OutputProposal)1005_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct Types.OutputProposal[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(OutputProposal)1005_storage\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_struct(OutputProposal)1005_storage\":{\"encoding\":\"inplace\",\"label\":\"struct Types.OutputProposal\",\"numberOfBytes\":\"64\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2OutputOracleStorageLayout = new(solc.StorageLayout)

var L2OutputOracleDeployedBin = "0x60806040526004361061018a5760003560e01c806393991af3116100d6578063cf8e5cf01161007f578063e1a41bcf11610059578063e1a41bcf14610526578063e4a3011614610559578063f4daa2911461057957600080fd5b8063cf8e5cf0146104d1578063d1de856c146104f1578063dcec33481461051157600080fd5b8063a8e4fb90116100b0578063a8e4fb9014610437578063bffa7f0f1461046a578063ce5db8d61461049e57600080fd5b806393991af3146103955780639aaab648146103c8578063a25ae557146103db57600080fd5b80636abcf563116101385780637f006420116101125780637f0064201461033d578063887862721461035d57806389c44cbb1461037357600080fd5b80636abcf563146102de5780636b4d98dd146102f357806370872aa51461032757600080fd5b8063534db0e211610169578063534db0e21461021f57806354fd4d501461027357806369f16eec146102c957600080fd5b80622134cc1461018f5780634599c788146101d6578063529933df146101eb575b600080fd5b34801561019b57600080fd5b506101c37f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b3480156101e257600080fd5b506101c36105ad565b3480156101f757600080fd5b506101c37f000000000000000000000000000000000000000000000000000000000000000081565b34801561022b57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101cd565b34801561027f57600080fd5b506102bc6040518060400160405280600581526020017f312e372e3000000000000000000000000000000000000000000000000000000081525081565b6040516101cd9190611358565b3480156102d557600080fd5b506101c3610620565b3480156102ea57600080fd5b506003546101c3565b3480156102ff57600080fd5b5061024e7f000000000000000000000000000000000000000000000000000000000000000081565b34801561033357600080fd5b506101c360015481565b34801561034957600080fd5b506101c36103583660046113cb565b610632565b34801561036957600080fd5b506101c360025481565b34801561037f57600080fd5b5061039361038e3660046113cb565b61084b565b005b3480156103a157600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006101c3565b6103936103d63660046113e4565b610b21565b3480156103e757600080fd5b506103fb6103f63660046113cb565b610fa0565b60408051825181526020808401516fffffffffffffffffffffffffffffffff9081169183019190915292820151909216908201526060016101cd565b34801561044357600080fd5b507f000000000000000000000000000000000000000000000000000000000000000061024e565b34801561047657600080fd5b5061024e7f000000000000000000000000000000000000000000000000000000000000000081565b3480156104aa57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006101c3565b3480156104dd57600080fd5b506103fb6104ec3660046113cb565b611034565b3480156104fd57600080fd5b506101c361050c3660046113cb565b61106c565b34801561051d57600080fd5b506101c36110ba565b34801561053257600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006101c3565b34801561056557600080fd5b50610393610574366004611416565b6110ef565b34801561058557600080fd5b506101c37f000000000000000000000000000000000000000000000000000000000000000081565b6003546000901561061757600380546105c890600190611467565b815481106105d8576105d861147e565b600091825260209091206002909102016001015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16919050565b6001545b905090565b60035460009061061b90600190611467565b600061063c6105ad565b8211156106f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f7420666f72206120626c6f636b207468617420686173206e6f74206265656e2060648201527f70726f706f736564000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b6003546107ab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604660248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f74206173206e6f206f7574707574732068617665206265656e2070726f706f7360648201527f6564207965740000000000000000000000000000000000000000000000000000608482015260a4016106ed565b6003546000905b8082101561084457600060026107c883856114ad565b6107d291906114c5565b905084600382815481106107e8576107e861147e565b600091825260209091206002909102016001015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16101561083a576108338160016114ad565b925061083e565b8091505b506107b2565b5092915050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610910576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4c324f75747075744f7261636c653a206f6e6c7920746865206368616c6c656e60448201527f67657220616464726573732063616e2064656c657465206f757470757473000060648201526084016106ed565b60035481106109c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f4c324f75747075744f7261636c653a2063616e6e6f742064656c657465206f7560448201527f747075747320616674657220746865206c6174657374206f757470757420696e60648201527f6465780000000000000000000000000000000000000000000000000000000000608482015260a4016106ed565b7f0000000000000000000000000000000000000000000000000000000000000000600382815481106109fb576109fb61147e565b6000918252602090912060016002909202010154610a2b906fffffffffffffffffffffffffffffffff1642611467565b10610ade576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604660248201527f4c324f75747075744f7261636c653a2063616e6e6f742064656c657465206f7560448201527f74707574732074686174206861766520616c7265616479206265656e2066696e60648201527f616c697a65640000000000000000000000000000000000000000000000000000608482015260a4016106ed565b6000610ae960035490565b90508160035581817f4ee37ac2c786ec85e87592d3c5c8a1dd66f8496dda3f125d9ea8ca5f657629b660405160405180910390a35050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610c0c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f4c324f75747075744f7261636c653a206f6e6c79207468652070726f706f736560448201527f7220616464726573732063616e2070726f706f7365206e6577206f757470757460648201527f7300000000000000000000000000000000000000000000000000000000000000608482015260a4016106ed565b610c146110ba565b8314610cc8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f4c324f75747075744f7261636c653a20626c6f636b206e756d626572206d757360448201527f7420626520657175616c20746f206e65787420657870656374656420626c6f6360648201527f6b206e756d626572000000000000000000000000000000000000000000000000608482015260a4016106ed565b42610cd28461106c565b10610d5f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f4c324f75747075744f7261636c653a2063616e6e6f742070726f706f7365204c60448201527f32206f757470757420696e20746865206675747572650000000000000000000060648201526084016106ed565b83610dec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4c324f75747075744f7261636c653a204c32206f75747075742070726f706f7360448201527f616c2063616e6e6f7420626520746865207a65726f206861736800000000000060648201526084016106ed565b8115610ea85781814014610ea8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604960248201527f4c324f75747075744f7261636c653a20626c6f636b206861736820646f65732060448201527f6e6f74206d61746368207468652068617368206174207468652065787065637460648201527f6564206865696768740000000000000000000000000000000000000000000000608482015260a4016106ed565b82610eb260035490565b857fa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e242604051610ee491815260200190565b60405180910390a45050604080516060810182529283526fffffffffffffffffffffffffffffffff4281166020850190815292811691840191825260038054600181018255600091909152935160029094027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b810194909455915190518216700100000000000000000000000000000000029116177fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c90910155565b604080516060810182526000808252602082018190529181019190915260038281548110610fd057610fd061147e565b600091825260209182902060408051606081018252600290930290910180548352600101546fffffffffffffffffffffffffffffffff8082169484019490945270010000000000000000000000000000000090049092169181019190915292915050565b6040805160608101825260008082526020820181905291810191909152600361105c83610632565b81548110610fd057610fd061147e565b60007f00000000000000000000000000000000000000000000000000000000000000006001548361109d9190611467565b6110a79190611500565b6002546110b491906114ad565b92915050565b60007f00000000000000000000000000000000000000000000000000000000000000006110e56105ad565b61061b91906114ad565b600054610100900460ff161580801561110f5750600054600160ff909116105b806111295750303b158015611129575060005460ff166001145b6111b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016106ed565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561121357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b428211156112ca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526044602482018190527f4c324f75747075744f7261636c653a207374617274696e67204c322074696d65908201527f7374616d70206d757374206265206c657373207468616e2063757272656e742060648201527f74696d6500000000000000000000000000000000000000000000000000000000608482015260a4016106ed565b60028290556001839055801561133757600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b600060208083528351808285015260005b8181101561138557858101830151858201604001528201611369565b81811115611397576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b6000602082840312156113dd57600080fd5b5035919050565b600080600080608085870312156113fa57600080fd5b5050823594602084013594506040840135936060013592509050565b6000806040838503121561142957600080fd5b50508035926020909101359150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561147957611479611438565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082198211156114c0576114c0611438565b500190565b6000826114fb577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561153857611538611438565b50029056fea164736f6c634300080f000a"


var L2OutputOracleImmutableReferencesJSON = "{\"76087\":[{\"start\":509,\"length\":32},{\"start\":1333,\"length\":32},{\"start\":4286,\"length\":32}],\"76090\":[{\"start\":417,\"length\":32},{\"start\":932,\"length\":32},{\"start\":4208,\"length\":32}],\"76093\":[{\"start\":558,\"length\":32},{\"start\":773,\"length\":32},{\"start\":2147,\"length\":32}],\"76096\":[{\"start\":1094,\"length\":32},{\"start\":1148,\"length\":32},{\"start\":2873,\"length\":32}],\"76099\":[{\"start\":1197,\"length\":32},{\"start\":1419,\"length\":32},{\"start\":2505,\"length\":32}]}"

func init() {
	if err := json.Unmarshal([]byte(L2OutputOracleStorageLayoutJSON), L2OutputOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2OutputOracle"] = L2OutputOracleStorageLayout
	deployedBytecodes["L2OutputOracle"] = L2OutputOracleDeployedBin
	immutableReferences["L2OutputOracle"] = L2OutputOracleImmutableReferencesJSON
}
