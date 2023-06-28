// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"gasLimit\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint64\"},{\"astId\":1009,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_resourceConfig\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_struct(ResourceConfig)1945_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_struct(ResourceConfig)1945_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceConfig\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b50600436106101515760003560e01c8063b40a817c116100cd578063f2fde38b11610081578063f68016b711610066578063f68016b7146103f7578063f975e9251461040b578063ffa1ad741461041e57600080fd5b8063f2fde38b146103db578063f45e65d8146103ee57600080fd5b8063c9b26f61116100b2578063c9b26f611461028b578063cc731b021461029e578063e81b2c6d146103d257600080fd5b8063b40a817c14610265578063c71973f61461027857600080fd5b80634f16540b11610124578063715018a611610109578063715018a61461022c5780638da5cb5b14610234578063935f029e1461025257600080fd5b80634f16540b146101f057806354fd4d501461021757600080fd5b80630c18c1621461015657806318d13918146101725780631fd19ee1146101875780634add321d146101cf575b600080fd5b61015f60655481565b6040519081526020015b60405180910390f35b610185610180366004611307565b610426565b005b7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08545b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610169565b6101d76104ea565b60405167ffffffffffffffff9091168152602001610169565b61015f7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0881565b61021f610515565b60405161016991906113a3565b6101856105b8565b60335473ffffffffffffffffffffffffffffffffffffffff166101aa565b6101856102603660046113b6565b6105cc565b6101856102733660046113f0565b610665565b610185610286366004611548565b610750565b610185610299366004611564565b610764565b6103626040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260695463ffffffff8082168352640100000000820460ff9081166020850152650100000000008304169383019390935266010000000000008104831660608301526a0100000000000000000000810490921660808201526e0100000000000000000000000000009091046fffffffffffffffffffffffffffffffff1660a082015290565b6040516101699190600060c08201905063ffffffff80845116835260ff602085015116602084015260ff6040850151166040840152806060850151166060840152806080850151166080840152506fffffffffffffffffffffffffffffffff60a08401511660a083015292915050565b61015f60675481565b6101856103e9366004611307565b610794565b61015f60665481565b6068546101d79067ffffffffffffffff1681565b61018561041936600461157d565b610848565b61015f600081565b61042e610afb565b610456817f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0855565b6040805173ffffffffffffffffffffffffffffffffffffffff8316602082015260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060035b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516104de91906113a3565b60405180910390a35050565b6069546000906105109063ffffffff6a010000000000000000000082048116911661161f565b905090565b60606105407f0000000000000000000000000000000000000000000000000000000000000000610b7c565b6105697f0000000000000000000000000000000000000000000000000000000000000000610b7c565b6105927f0000000000000000000000000000000000000000000000000000000000000000610b7c565b6040516020016105a49392919061164b565b604051602081830303815290604052905090565b6105c0610afb565b6105ca6000610cb9565b565b6105d4610afb565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be8360405161065891906113a3565b60405180910390a3505050565b61066d610afb565b6106756104ea565b67ffffffffffffffff168167ffffffffffffffff1610156106f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064015b60405180910390fd5b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff831690811790915560408051602080820193909352815180820390930183528101905260026104ad565b610758610afb565b61076181610d30565b50565b61076c610afb565b60678190556040805160208082018490528251808303909101815290820190915260006104ad565b61079c610afb565b73ffffffffffffffffffffffffffffffffffffffff811661083f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016106ee565b61076181610cb9565b600054610100900460ff16158080156108685750600054600160ff909116105b806108825750303b158015610882575060005460ff166001145b61090e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016106ee565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561096c57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6109746111a4565b61097d88610794565b606587905560668690556067859055606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff86161790557f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c088390556109ed82610d30565b6109f56104ea565b67ffffffffffffffff168467ffffffffffffffff161015610a72576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016106ee565b8015610ad557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60335473ffffffffffffffffffffffffffffffffffffffff1633146105ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106ee565b606081600003610bbf57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610be95780610bd3816116c1565b9150610be29050600a83611728565b9150610bc3565b60008167ffffffffffffffff811115610c0457610c0461140b565b6040519080825280601f01601f191660200182016040528015610c2e576020820181803683370190505b5090505b8415610cb157610c4360018361173c565b9150610c50600a86611753565b610c5b906030611767565b60f81b818381518110610c7057610c7061177f565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610caa600a86611728565b9450610c32565b949350505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b8060a001516fffffffffffffffffffffffffffffffff16816060015163ffffffff161115610de0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f53797374656d436f6e6669673a206d696e206261736520666565206d7573742060448201527f6265206c657373207468616e206d61782062617365000000000000000000000060648201526084016106ee565b6001816040015160ff1611610e77576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a2064656e6f6d696e61746f72206d757374206260448201527f65206c6172676572207468616e2031000000000000000000000000000000000060648201526084016106ee565b6068546080820151825167ffffffffffffffff90921691610e9891906117ae565b63ffffffff161115610f06576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016106ee565b6000816020015160ff1611610f9d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a20656c6173746963697479206d756c7469706c60448201527f6965722063616e6e6f742062652030000000000000000000000000000000000060648201526084016106ee565b8051602082015163ffffffff82169160ff90911690610fbd9082906117cd565b610fc791906117f0565b63ffffffff161461105a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f53797374656d436f6e6669673a20707265636973696f6e206c6f73732077697460448201527f6820746172676574207265736f75726365206c696d697400000000000000000060648201526084016106ee565b805160698054602084015160408501516060860151608087015160a09097015163ffffffff9687167fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009095169490941764010000000060ff94851602177fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff166501000000000093909216929092027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1617660100000000000091851691909102177fffff0000000000000000000000000000000000000000ffffffffffffffffffff166a010000000000000000000093909416929092027fffff00000000000000000000000000000000ffffffffffffffffffffffffffff16929092176e0100000000000000000000000000006fffffffffffffffffffffffffffffffff90921691909102179055565b600054610100900460ff1661123b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106ee565b6105ca600054610100900460ff166112d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106ee565b6105ca33610cb9565b803573ffffffffffffffffffffffffffffffffffffffff8116811461130257600080fd5b919050565b60006020828403121561131957600080fd5b611322826112de565b9392505050565b60005b8381101561134457818101518382015260200161132c565b83811115611353576000848401525b50505050565b60008151808452611371816020860160208601611329565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006113226020830184611359565b600080604083850312156113c957600080fd5b50508035926020909101359150565b803567ffffffffffffffff8116811461130257600080fd5b60006020828403121561140257600080fd5b611322826113d8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b803563ffffffff8116811461130257600080fd5b803560ff8116811461130257600080fd5b80356fffffffffffffffffffffffffffffffff8116811461130257600080fd5b600060c0828403121561149157600080fd5b60405160c0810181811067ffffffffffffffff821117156114db577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040529050806114ea8361143a565b81526114f86020840161144e565b60208201526115096040840161144e565b604082015261151a6060840161143a565b606082015261152b6080840161143a565b608082015261153c60a0840161145f565b60a08201525092915050565b600060c0828403121561155a57600080fd5b611322838361147f565b60006020828403121561157657600080fd5b5035919050565b6000806000806000806000610180888a03121561159957600080fd5b6115a2886112de565b96506020880135955060408801359450606088013593506115c5608089016113d8565b92506115d360a089016112de565b91506115e28960c08a0161147f565b905092959891949750929550565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff808316818516808303821115611642576116426115f0565b01949350505050565b6000845161165d818460208901611329565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611699816001850160208a01611329565b600192019182015283516116b4816002840160208801611329565b0160020195945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036116f2576116f26115f0565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082611737576117376116f9565b500490565b60008282101561174e5761174e6115f0565b500390565b600082611762576117626116f9565b500690565b6000821982111561177a5761177a6115f0565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600063ffffffff808316818516808303821115611642576116426115f0565b600063ffffffff808416806117e4576117e46116f9565b92169190910492915050565b600063ffffffff80831681851681830481118215151615611813576118136115f0565b0294935050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
}
