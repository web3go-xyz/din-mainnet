// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1ERC721BridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"messenger\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_contract(CrossDomainMessenger)1005\"},{\"astId\":1003,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)48_storage\"},{\"astId\":1004,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"49\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_mapping(t_uint256,t_bool)))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)48_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[48]\",\"numberOfBytes\":\"1536\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(CrossDomainMessenger)1005\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_mapping(t_uint256,t_bool)))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e mapping(uint256 =\u003e bool)))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_mapping(t_uint256,t_bool))\"},\"t_mapping(t_address,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1ERC721BridgeStorageLayout = new(solc.StorageLayout)

var L1ERC721BridgeDeployedBin = "0x608060405234801561001057600080fd5b50600436106100be5760003560e01c80637f46ddb211610076578063aa5574521161005b578063aa557452146101df578063c4d66de8146101f2578063c89701a21461020557600080fd5b80637f46ddb214610194578063927ede2d146101bb57600080fd5b806354fd4d50116100a757806354fd4d50146101285780635d93a3fc1461013d578063761f44931461018157600080fd5b80633687011a146100c35780633cb747bf146100d8575b600080fd5b6100d66100d1366004610ff3565b61022b565b005b6000546100fe9062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101306102d7565b60405161011f91906110f0565b61017161014b36600461110a565b603160209081526000938452604080852082529284528284209052825290205460ff1681565b604051901515815260200161011f565b6100d661018f36600461114b565b61037a565b6100fe7f000000000000000000000000000000000000000000000000000000000000000081565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff166100fe565b6100d66101ed3660046111e3565b6107e5565b6100d661020036600461125a565b6108a1565b7f00000000000000000000000000000000000000000000000000000000000000006100fe565b333b156102bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732314272696467653a206163636f756e74206973206e6f742065787460448201527f65726e616c6c79206f776e65640000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102cf86863333888888886109eb565b505050505050565b60606103027f0000000000000000000000000000000000000000000000000000000000000000610d4b565b61032b7f0000000000000000000000000000000000000000000000000000000000000000610d4b565b6103547f0000000000000000000000000000000000000000000000000000000000000000610d4b565b60405160200161036693929190611277565b604051602081830303815290604052905090565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314801561048257507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610446573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061046a91906112ed565b73ffffffffffffffffffffffffffffffffffffffff16145b61050e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4552433732314272696467653a2066756e6374696f6e2063616e206f6e6c792060448201527f62652063616c6c65642066726f6d20746865206f74686572206272696467650060648201526084016102b6565b3073ffffffffffffffffffffffffffffffffffffffff8816036105b3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4c314552433732314272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c660000000000000000000000000000000000000000000060648201526084016102b6565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152603160209081526040808320938a1683529281528282208683529052205460ff161515600114610682576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f4c314552433732314272696467653a20546f6b656e204944206973206e6f742060448201527f657363726f77656420696e20746865204c31204272696467650000000000000060648201526084016102b6565b73ffffffffffffffffffffffffffffffffffffffff87811660008181526031602090815260408083208b8616845282528083208884529091529081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517f42842e0e000000000000000000000000000000000000000000000000000000008152306004820152918616602483015260448201859052906342842e0e90606401600060405180830381600087803b15801561074257600080fd5b505af1158015610756573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac878787876040516107d49493929190611353565b60405180910390a450505050505050565b73ffffffffffffffffffffffffffffffffffffffff8516610888576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433732314272696467653a206e667420726563697069656e742063616e6e60448201527f6f7420626520616464726573732830290000000000000000000000000000000060648201526084016102b6565b61089887873388888888886109eb565b50505050505050565b600054600290610100900460ff161580156108c3575060005460ff8083169116105b61094f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016102b6565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff83161761010017905561098982610e88565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b73ffffffffffffffffffffffffffffffffffffffff8716610a8e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4c314552433732314272696467653a2072656d6f746520746f6b656e2063616e60448201527f6e6f74206265206164647265737328302900000000000000000000000000000060648201526084016102b6565b600063761f449360e01b888a8989898888604051602401610ab59796959493929190611393565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000959095169490941790935273ffffffffffffffffffffffffffffffffffffffff8c81166000818152603186528381208e8416825286528381208b82529095529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590517f23b872dd000000000000000000000000000000000000000000000000000000008152908a166004820152306024820152604481018890529092506323b872dd90606401600060405180830381600087803b158015610bf557600080fd5b505af1158015610c09573d6000803e3d6000fd5b50506000546040517f3dbb202b0000000000000000000000000000000000000000000000000000000081526201000090910473ffffffffffffffffffffffffffffffffffffffff169250633dbb202b9150610c8c907f000000000000000000000000000000000000000000000000000000000000000090859089906004016113f0565b600060405180830381600087803b158015610ca657600080fd5b505af1158015610cba573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff167fb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a589898888604051610d389493929190611353565b60405180910390a4505050505050505050565b606081600003610d8e57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610db85780610da281611464565b9150610db19050600a836114cb565b9150610d92565b60008167ffffffffffffffff811115610dd357610dd36114df565b6040519080825280601f01601f191660200182016040528015610dfd576020820181803683370190505b5090505b8415610e8057610e1260018361150e565b9150610e1f600a86611525565b610e2a906030611539565b60f81b818381518110610e3f57610e3f611551565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610e79600a866114cb565b9450610e01565b949350505050565b600054610100900460ff16610f1f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102b6565b6000805473ffffffffffffffffffffffffffffffffffffffff90921662010000027fffffffffffffffffffff0000000000000000000000000000000000000000ffff909216919091179055565b73ffffffffffffffffffffffffffffffffffffffff81168114610f8e57600080fd5b50565b803563ffffffff81168114610fa557600080fd5b919050565b60008083601f840112610fbc57600080fd5b50813567ffffffffffffffff811115610fd457600080fd5b602083019150836020828501011115610fec57600080fd5b9250929050565b60008060008060008060a0878903121561100c57600080fd5b863561101781610f6c565b9550602087013561102781610f6c565b94506040870135935061103c60608801610f91565b9250608087013567ffffffffffffffff81111561105857600080fd5b61106489828a01610faa565b979a9699509497509295939492505050565b60005b83811015611091578181015183820152602001611079565b838111156110a0576000848401525b50505050565b600081518084526110be816020860160208601611076565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061110360208301846110a6565b9392505050565b60008060006060848603121561111f57600080fd5b833561112a81610f6c565b9250602084013561113a81610f6c565b929592945050506040919091013590565b600080600080600080600060c0888a03121561116657600080fd5b873561117181610f6c565b9650602088013561118181610f6c565b9550604088013561119181610f6c565b945060608801356111a181610f6c565b93506080880135925060a088013567ffffffffffffffff8111156111c457600080fd5b6111d08a828b01610faa565b989b979a50959850939692959293505050565b600080600080600080600060c0888a0312156111fe57600080fd5b873561120981610f6c565b9650602088013561121981610f6c565b9550604088013561122981610f6c565b94506060880135935061123e60808901610f91565b925060a088013567ffffffffffffffff8111156111c457600080fd5b60006020828403121561126c57600080fd5b813561110381610f6c565b60008451611289818460208901611076565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516112c5816001850160208a01611076565b600192019182015283516112e0816002840160208801611076565b0160020195945050505050565b6000602082840312156112ff57600080fd5b815161110381610f6c565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8516815283602082015260606040820152600061138960608301848661130a565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a08301526113e360c08301848661130a565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815260606020820152600061141f60608301856110a6565b905063ffffffff83166040830152949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361149557611495611435565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826114da576114da61149c565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008282101561152057611520611435565b500390565b6000826115345761153461149c565b500690565b6000821982111561154c5761154c611435565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1ERC721BridgeStorageLayoutJSON), L1ERC721BridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1ERC721Bridge"] = L1ERC721BridgeStorageLayout
	deployedBytecodes["L1ERC721Bridge"] = L1ERC721BridgeDeployedBin
}
