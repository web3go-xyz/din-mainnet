// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2StandardBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer_1_0_20\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1002,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1003,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_uint256)47_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)47_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[47]\",\"numberOfBytes\":\"1504\",\"base\":\"t_uint256\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L2StandardBridgeStorageLayout = new(solc.StorageLayout)

var L2StandardBridgeDeployedBin = "0x6080604052600436106100f75760003560e01c8063662a633a1161008a578063927ede2d11610059578063927ede2d146103d3578063a3a7954814610407578063c89701a214610224578063e11013dd1461041a57600080fd5b8063662a633a146103265780637f46ddb214610339578063870876231461036d5780638f601f661461038d57600080fd5b806336c717c1116100c657806336c717c1146102245780633cb747bf1461027d578063540abf73146102b057806354fd4d50146102d057600080fd5b80630166a07a146101cb57806309fc8843146101eb5780631635f5fd146101fe57806332b7006d1461021157600080fd5b366101c657333b15610190576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b6101c473deaddeaddeaddeaddeaddeaddeaddeaddead000033333462030d406040518060200160405280600081525061042d565b005b600080fd5b3480156101d757600080fd5b506101c46101e636600461226f565b610508565b6101c46101f9366004612320565b6108f5565b6101c461020c366004612373565b6109cc565b6101c461021f3660046123e6565b610e99565b34801561023057600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561028957600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610253565b3480156102bc57600080fd5b506101c46102cb36600461243a565b610f73565b3480156102dc57600080fd5b506103196040518060400160405280600581526020017f312e352e3000000000000000000000000000000000000000000000000000000081525081565b6040516102749190612527565b6101c461033436600461226f565b610fb8565b34801561034557600080fd5b506102537f000000000000000000000000000000000000000000000000000000000000000081565b34801561037957600080fd5b506101c461038836600461253a565b61102b565b34801561039957600080fd5b506103c56103a83660046125bd565b600260209081526000928352604080842090915290825290205481565b604051908152602001610274565b3480156103df57600080fd5b506102537f000000000000000000000000000000000000000000000000000000000000000081565b6101c461041536600461253a565b6110ff565b6101c46104283660046125f6565b611143565b7fffffffffffffffffffffffff215221522152215221522152215221522153000073ffffffffffffffffffffffffffffffffffffffff87160161047c57610477858585858561118c565b610500565b60008673ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ed9190612659565b90506104fe87828888888888611370565b505b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561062657507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060e9190612659565b73ffffffffffffffffffffffffffffffffffffffff16145b6106d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a401610187565b6106e1876116b7565b1561082f576106f08787611719565b6107a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a401610187565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b15801561081257600080fd5b505af1158015610826573d6000803e3d6000fd5b505050506108b1565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a168352929052205461086d9084906126a5565b73ffffffffffffffffffffffffffffffffffffffff8089166000818152600260209081526040808320948c16835293905291909120919091556108b1908585611839565b6104fe878787878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061190d92505050565b333b15610984576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f410000000000000000006064820152608401610187565b6109c73333348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061118c92505050565b505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148015610aea57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610aae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad29190612659565b73ffffffffffffffffffffffffffffffffffffffff16145b610b9c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a401610187565b823414610c2b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e742072657175697265640000000000006064820152608401610187565b3073ffffffffffffffffffffffffffffffffffffffff851603610cd0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c6600000000000000000000000000000000000000000000000000000000006064820152608401610187565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610dab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201527f657373656e6765720000000000000000000000000000000000000000000000006064820152608401610187565b610ded85858585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061199b92505050565b6000610e0a855a8660405180602001604052806000815250611a3c565b905080610500576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c656400000000000000000000000000000000000000000000000000000000006064820152608401610187565b333b15610f28576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f410000000000000000006064820152608401610187565b610f6c853333878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061042d92505050565b5050505050565b6104fe87873388888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061137092505050565b73ffffffffffffffffffffffffffffffffffffffff8716158015611005575073ffffffffffffffffffffffffffffffffffffffff861673deaddeaddeaddeaddeaddeaddeaddeaddead0000145b1561101c5761101785858585856109cc565b6104fe565b6104fe86888787878787610508565b333b156110ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f410000000000000000006064820152608401610187565b61050086863333888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061137092505050565b610500863387878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061042d92505050565b6111863385348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061118c92505050565b50505050565b82341461121b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5374616e646172644272696467653a206272696467696e6720455448206d757360448201527f7420696e636c7564652073756666696369656e74204554482076616c756500006064820152608401610187565b61122785858584611a56565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633dbb202b847f0000000000000000000000000000000000000000000000000000000000000000631635f5fd60e01b898989886040516024016112a494939291906126bc565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b909216825261133792918890600401612705565b6000604051808303818588803b15801561135057600080fd5b505af1158015611364573d6000803e3d6000fd5b50505050505050505050565b611379876116b7565b156114c7576113888787611719565b61143a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a401610187565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201859052881690639dc29fac90604401600060405180830381600087803b1580156114aa57600080fd5b505af11580156114be573d6000803e3d6000fd5b5050505061155b565b6114e973ffffffffffffffffffffffffffffffffffffffff8816863086611af7565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a168352929052205461152790849061274a565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600260209081526040808320938b16835292905220555b611569878787878786611b55565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633dbb202b7f0000000000000000000000000000000000000000000000000000000000000000630166a07a60e01b898b8a8a8a896040516024016115e996959493929190612762565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b909216825261167c92918790600401612705565b600060405180830381600087803b15801561169657600080fd5b505af11580156116aa573d6000803e3d6000fd5b5050505050505050505050565b60006116e3827f1d1d8b6300000000000000000000000000000000000000000000000000000000611be3565b806117135750611713827fec4fc8e300000000000000000000000000000000000000000000000000000000611be3565b92915050565b6000611745837f1d1d8b6300000000000000000000000000000000000000000000000000000000611be3565b156117ee578273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa158015611795573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117b99190612659565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050611713565b8273ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015611795573d6000803e3d6000fd5b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526109c79084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611c06565b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89868686604051611985939291906127bd565b60405180910390a4610500868686868686611d12565b8373ffffffffffffffffffffffffffffffffffffffff1673deaddeaddeaddeaddeaddeaddeaddeaddead000073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89868686604051611a28939291906127bd565b60405180910390a461118684848484611d9a565b600080600080845160208601878a8af19695505050505050565b8373ffffffffffffffffffffffffffffffffffffffff1673deaddeaddeaddeaddeaddeaddeaddeaddead000073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e868686604051611ae3939291906127bd565b60405180910390a461118684848484611e07565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526111869085907f23b872dd000000000000000000000000000000000000000000000000000000009060840161188b565b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e868686604051611bcd939291906127bd565b60405180910390a4610500868686868686611e66565b6000611bee83611ede565b8015611bff5750611bff8383611f42565b9392505050565b6000611c68826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166120119092919063ffffffff16565b8051909150156109c75780806020019051810190611c8691906127fb565b6109c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610187565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd868686604051611d8a939291906127bd565b60405180910390a4505050505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d8484604051611df992919061281d565b60405180910390a350505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af58484604051611df992919061281d565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf868686604051611d8a939291906127bd565b6000611f0a827f01ffc9a700000000000000000000000000000000000000000000000000000000611f42565b80156117135750611f3b827fffffffff00000000000000000000000000000000000000000000000000000000611f42565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015611ffa575060208210155b80156120065750600081115b979650505050505050565b60606120208484600085612028565b949350505050565b6060824710156120ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610187565b73ffffffffffffffffffffffffffffffffffffffff85163b612138576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610187565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516121619190612836565b60006040518083038185875af1925050503d806000811461219e576040519150601f19603f3d011682016040523d82523d6000602084013e6121a3565b606091505b5091509150612006828286606083156121bd575081611bff565b8251156121cd5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101879190612527565b73ffffffffffffffffffffffffffffffffffffffff8116811461222357600080fd5b50565b60008083601f84011261223857600080fd5b50813567ffffffffffffffff81111561225057600080fd5b60208301915083602082850101111561226857600080fd5b9250929050565b600080600080600080600060c0888a03121561228a57600080fd5b873561229581612201565b965060208801356122a581612201565b955060408801356122b581612201565b945060608801356122c581612201565b93506080880135925060a088013567ffffffffffffffff8111156122e857600080fd5b6122f48a828b01612226565b989b979a50959850939692959293505050565b803563ffffffff8116811461231b57600080fd5b919050565b60008060006040848603121561233557600080fd5b61233e84612307565b9250602084013567ffffffffffffffff81111561235a57600080fd5b61236686828701612226565b9497909650939450505050565b60008060008060006080868803121561238b57600080fd5b853561239681612201565b945060208601356123a681612201565b935060408601359250606086013567ffffffffffffffff8111156123c957600080fd5b6123d588828901612226565b969995985093965092949392505050565b6000806000806000608086880312156123fe57600080fd5b853561240981612201565b94506020860135935061241e60408701612307565b9250606086013567ffffffffffffffff8111156123c957600080fd5b600080600080600080600060c0888a03121561245557600080fd5b873561246081612201565b9650602088013561247081612201565b9550604088013561248081612201565b94506060880135935061249560808901612307565b925060a088013567ffffffffffffffff8111156122e857600080fd5b60005b838110156124cc5781810151838201526020016124b4565b838111156111865750506000910152565b600081518084526124f58160208601602086016124b1565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611bff60208301846124dd565b60008060008060008060a0878903121561255357600080fd5b863561255e81612201565b9550602087013561256e81612201565b94506040870135935061258360608801612307565b9250608087013567ffffffffffffffff81111561259f57600080fd5b6125ab89828a01612226565b979a9699509497509295939492505050565b600080604083850312156125d057600080fd5b82356125db81612201565b915060208301356125eb81612201565b809150509250929050565b6000806000806060858703121561260c57600080fd5b843561261781612201565b935061262560208601612307565b9250604085013567ffffffffffffffff81111561264157600080fd5b61264d87828801612226565b95989497509550505050565b60006020828403121561266b57600080fd5b8151611bff81612201565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156126b7576126b7612676565b500390565b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250836040830152608060608301526126fb60808301846124dd565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815260606020820152600061273460608301856124dd565b905063ffffffff83166040830152949350505050565b6000821982111561275d5761275d612676565b500190565b600073ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a08301526127b160c08301846124dd565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681528260208201526060604082015260006127f260608301846124dd565b95945050505050565b60006020828403121561280d57600080fd5b81518015158114611bff57600080fd5b82815260406020820152600061202060408301846124dd565b600082516128488184602087016124b1565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2StandardBridgeStorageLayoutJSON), L2StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2StandardBridge"] = L2StandardBridgeStorageLayout
	deployedBytecodes["L2StandardBridge"] = L2StandardBridgeDeployedBin
}
