// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2StandardBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":25104,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer0\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint256\"},{\"astId\":25107,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer1\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":25114,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L2StandardBridgeStorageLayout = new(solc.StorageLayout)

var L2StandardBridgeDeployedBin = "0x6080604052600436106100e15760003560e01c8063662a633a1161007f578063a3a7954811610059578063a3a7954814610312578063af565a1314610325578063c89701a214610345578063e11013dd1461037957600080fd5b8063662a633a1461029957806387087623146102ac5780638f601f66146102cc57600080fd5b806332b7006d116100bb57806332b7006d146101e65780633cb747bf146101f9578063540abf731461025757806354fd4d501461027757600080fd5b80630166a07a146101a057806309fc8843146101c05780631635f5fd146101d357600080fd5b3661019b57333b1561017a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b61019933333462030d406040518060200160405280600081525061038c565b005b600080fd5b3480156101ac57600080fd5b506101996101bb36600461223c565b61053c565b6101996101ce3660046122ed565b6108c3565b6101996101e1366004612340565b61099a565b6101996101f43660046123b3565b610dbb565b34801561020557600080fd5b5061022d7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561026357600080fd5b50610199610272366004612407565b610e60565b34801561028357600080fd5b5061028c610e70565b60405161024e91906124f4565b6101996102a736600461223c565b610f13565b3480156102b857600080fd5b506101996102c7366004612507565b611000565b3480156102d857600080fd5b506103046102e736600461258a565b600260209081526000928352604080842090915290825290205481565b60405190815260200161024e565b610199610320366004612507565b61109f565b34801561033157600080fd5b506101996103403660046125c3565b6110ae565b34801561035157600080fd5b5061022d7f000000000000000000000000000000000000000000000000000000000000000081565b610199610387366004612614565b6113c1565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af585846040516103eb929190612677565b60405180910390a37f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633dbb202b847f0000000000000000000000000000000000000000000000000000000000000000631635f5fd60e01b898989886040516024016104709493929190612690565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b9092168252610503929188906004016126d9565b6000604051808303818588803b15801561051c57600080fd5b505af1158015610530573d6000803e3d6000fd5b50505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561065a57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561061e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610642919061271e565b73ffffffffffffffffffffffffffffffffffffffff16145b61070c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a401610171565b6040517faf565a1300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808916600483015280881660248301528516604482015260648101849052309063af565a1390608401600060405180830381600087803b15801561078a57600080fd5b505af192505050801561079b575060015b610837576107b0878786888760008888611404565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f2755817676249910615f0a6a240ad225abe5343df8d527f7294c4af36a92009a8787878760405161082a9493929190612784565b60405180910390a46108ba565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd878787876040516108b19493929190612784565b60405180910390a45b50505050505050565b333b15610952576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f410000000000000000006064820152608401610171565b6109953333348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061038c92505050565b505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148015610ab857507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a7c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa0919061271e565b73ffffffffffffffffffffffffffffffffffffffff16145b610b6a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a401610171565b823414610bf9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e742072657175697265640000000000006064820152608401610171565b3073ffffffffffffffffffffffffffffffffffffffff851603610c9e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c6600000000000000000000000000000000000000000000000000000000006064820152608401610171565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d858585604051610cff939291906127ba565b60405180910390a36000610d24855a86604051806020016040528060008152506115d7565b905080610db3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c656400000000000000000000000000000000000000000000000000000000006064820152608401610171565b505050505050565b333b15610e4a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f410000000000000000006064820152608401610171565b610e59853333878787876115f1565b5050505050565b6108ba8787338888888888611828565b6060610e9b7f0000000000000000000000000000000000000000000000000000000000000000611ad2565b610ec47f0000000000000000000000000000000000000000000000000000000000000000611ad2565b610eed7f0000000000000000000000000000000000000000000000000000000000000000611ad2565b604051602001610eff939291906127dd565b604051602081830303815290604052905090565b73ffffffffffffffffffffffffffffffffffffffff8716158015610f60575073ffffffffffffffffffffffffffffffffffffffff861673deaddeaddeaddeaddeaddeaddeaddeaddead0000145b15610f7757610f72858585858561099a565b610f86565b610f868688878787878761053c565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89878787876040516108b19493929190612784565b333b1561108f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f410000000000000000006064820152608401610171565b610db38686333388888888611828565b610db3863387878787876115f1565b33301461113d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642062792073656c66000000000000000000000000006064820152608401610171565b3073ffffffffffffffffffffffffffffffffffffffff8516036111e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5374616e646172644272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c66000000000000000000000000000000000000000000006064820152608401610171565b6111eb84611c0f565b15611339576111fa8484611c41565b6112ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a401610171565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152602482018390528516906340c10f1990604401600060405180830381600087803b15801561131c57600080fd5b505af1158015611330573d6000803e3d6000fd5b505050506113bb565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260026020908152604080832093871683529290522054611377908290612882565b73ffffffffffffffffffffffffffffffffffffffff8086166000818152600260209081526040808320948916835293905291909120919091556113bb908383611ce8565b50505050565b6113bb3385348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061038c92505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633dbb202b7f0000000000000000000000000000000000000000000000000000000000000000630166a07a60e01b8a8c8b8b8b8a8a6040516024016114869796959493929190612899565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b9092168252611519929188906004016126d9565b600060405180830381600087803b15801561153357600080fd5b505af1158015611547573d6000803e3d6000fd5b505050508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf888887876040516115c59493929190612784565b60405180910390a45050505050505050565b600080600080845160208601878a8af19695505050505050565b60008773ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa15801561163e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611662919061271e565b90507fffffffffffffffffffffffff215221522152215221522152215221522153000073ffffffffffffffffffffffffffffffffffffffff89160161179e57843414611756576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f4c325374616e646172644272696467653a20455448207769746864726177616c60448201527f73206d75737420696e636c7564652073756666696369656e742045544820766160648201527f6c75650000000000000000000000000000000000000000000000000000000000608482015260a401610171565b6117998787878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061038c92505050565b6117ae565b6117ae8882898989898989611828565b8673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e898988886040516115c59493929190612784565b3073ffffffffffffffffffffffffffffffffffffffff8916036118cd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5374616e646172644272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c66000000000000000000000000000000000000000000006064820152608401610171565b6118d688611c0f565b15611a24576118e58888611c41565b611997576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a401610171565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff878116600483015260248201869052891690639dc29fac90604401600060405180830381600087803b158015611a0757600080fd5b505af1158015611a1b573d6000803e3d6000fd5b50505050611ab8565b611a4673ffffffffffffffffffffffffffffffffffffffff8916873087611dbc565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600260209081526040808320938b1683529290522054611a849085906128f6565b73ffffffffffffffffffffffffffffffffffffffff808a166000908152600260209081526040808320938c16835292905220555b611ac88888888888888888611404565b5050505050505050565b606081600003611b1557505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611b3f5780611b298161290e565b9150611b389050600a83612975565b9150611b19565b60008167ffffffffffffffff811115611b5a57611b5a612989565b6040519080825280601f01601f191660200182016040528015611b84576020820181803683370190505b5090505b8415611c0757611b99600183612882565b9150611ba6600a866129b8565b611bb19060306128f6565b60f81b818381518110611bc657611bc66129cc565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611c00600a86612975565b9450611b88565b949350505050565b6000611c3b827f1d1d8b6300000000000000000000000000000000000000000000000000000000611e1a565b92915050565b60008273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c8e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cb2919061271e565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614905092915050565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526109959084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611e3d565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526113bb9085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611d3a565b6000611e2583611f49565b8015611e365750611e368383611fad565b9392505050565b6000611e9f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661207c9092919063ffffffff16565b8051909150156109955780806020019051810190611ebd91906129fb565b610995576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610171565b6000611f75827f01ffc9a700000000000000000000000000000000000000000000000000000000611fad565b8015611c3b5750611fa6827fffffffff00000000000000000000000000000000000000000000000000000000611fad565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015612065575060208210155b80156120715750600081115b979650505050505050565b6060611c0784846000858573ffffffffffffffffffffffffffffffffffffffff85163b612105576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610171565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161212e9190612a1d565b60006040518083038185875af1925050503d806000811461216b576040519150601f19603f3d011682016040523d82523d6000602084013e612170565b606091505b50915091506120718282866060831561218a575081611e36565b82511561219a5782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017191906124f4565b73ffffffffffffffffffffffffffffffffffffffff811681146121f057600080fd5b50565b60008083601f84011261220557600080fd5b50813567ffffffffffffffff81111561221d57600080fd5b60208301915083602082850101111561223557600080fd5b9250929050565b600080600080600080600060c0888a03121561225757600080fd5b8735612262816121ce565b96506020880135612272816121ce565b95506040880135612282816121ce565b94506060880135612292816121ce565b93506080880135925060a088013567ffffffffffffffff8111156122b557600080fd5b6122c18a828b016121f3565b989b979a50959850939692959293505050565b803563ffffffff811681146122e857600080fd5b919050565b60008060006040848603121561230257600080fd5b61230b846122d4565b9250602084013567ffffffffffffffff81111561232757600080fd5b612333868287016121f3565b9497909650939450505050565b60008060008060006080868803121561235857600080fd5b8535612363816121ce565b94506020860135612373816121ce565b935060408601359250606086013567ffffffffffffffff81111561239657600080fd5b6123a2888289016121f3565b969995985093965092949392505050565b6000806000806000608086880312156123cb57600080fd5b85356123d6816121ce565b9450602086013593506123eb604087016122d4565b9250606086013567ffffffffffffffff81111561239657600080fd5b600080600080600080600060c0888a03121561242257600080fd5b873561242d816121ce565b9650602088013561243d816121ce565b9550604088013561244d816121ce565b945060608801359350612462608089016122d4565b925060a088013567ffffffffffffffff8111156122b557600080fd5b60005b83811015612499578181015183820152602001612481565b838111156113bb5750506000910152565b600081518084526124c281602086016020860161247e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611e3660208301846124aa565b60008060008060008060a0878903121561252057600080fd5b863561252b816121ce565b9550602087013561253b816121ce565b945060408701359350612550606088016122d4565b9250608087013567ffffffffffffffff81111561256c57600080fd5b61257889828a016121f3565b979a9699509497509295939492505050565b6000806040838503121561259d57600080fd5b82356125a8816121ce565b915060208301356125b8816121ce565b809150509250929050565b600080600080608085870312156125d957600080fd5b84356125e4816121ce565b935060208501356125f4816121ce565b92506040850135612604816121ce565b9396929550929360600135925050565b6000806000806060858703121561262a57600080fd5b8435612635816121ce565b9350612643602086016122d4565b9250604085013567ffffffffffffffff81111561265f57600080fd5b61266b878288016121f3565b95989497509550505050565b828152604060208201526000611c0760408301846124aa565b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250836040830152608060608301526126cf60808301846124aa565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815260606020820152600061270860608301856124aa565b905063ffffffff83166040830152949350505050565b60006020828403121561273057600080fd5b8151611e36816121ce565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff851681528360208201526060604082015260006126cf60608301848661273b565b8381526040602082015260006127d460408301848661273b565b95945050505050565b600084516127ef81846020890161247e565b80830190507f2e00000000000000000000000000000000000000000000000000000000000000808252855161282b816001850160208a0161247e565b6001920191820152835161284681600284016020880161247e565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561289457612894612853565b500390565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a08301526128e960c08301848661273b565b9998505050505050505050565b6000821982111561290957612909612853565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361293f5761293f612853565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261298457612984612946565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000826129c7576129c7612946565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215612a0d57600080fd5b81518015158114611e3657600080fd5b60008251612a2f81846020870161247e565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2StandardBridgeStorageLayoutJSON), L2StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2StandardBridge"] = L2StandardBridgeStorageLayout
	deployedBytecodes["L2StandardBridge"] = L2StandardBridgeDeployedBin
}
