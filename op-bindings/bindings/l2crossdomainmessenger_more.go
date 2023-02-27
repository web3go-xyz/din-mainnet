// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1020_storage\"},{\"astId\":1004,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1019_storage\"},{\"astId\":1006,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1019_storage\"},{\"astId\":1008,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap_reentrancy_guard\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)1019_storage\"},{\"astId\":1010,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":1015,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"reentrancyLocks\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1017,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"208\",\"type\":\"t_array(t_uint256)1018_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1018_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[41]\",\"numberOfBytes\":\"1312\"},\"t_array(t_uint256)1019_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1020_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L2CrossDomainMessengerDeployedBin = "0x6080604052600436106100f35760003560e01c80638129fc1c1161008a578063b1b1b20911610059578063b1b1b209146102c3578063b28ade25146102f3578063d764ad0b14610313578063ecc704281461032657600080fd5b80638129fc1c146102075780639fce812c1461021c578063a4e7f8bd14610250578063a71198691461029057600080fd5b80633f827a5a116100c65780633f827a5a1461016c57806354fd4d50146101945780636e296e45146101b65780637dea7cc3146101f057600080fd5b8063028f85f7146100f85780630c5684981461012b5780632828d7e8146101415780633dbb202b14610157575b600080fd5b34801561010457600080fd5b5061010d601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b34801561013757600080fd5b5061010d6103e881565b34801561014d57600080fd5b5061010d6103f881565b61016a610165366004611815565b61038b565b005b34801561017857600080fd5b50610181600181565b60405161ffff9091168152602001610122565b3480156101a057600080fd5b506101a96105ef565b60405161012291906118f4565b3480156101c257600080fd5b506101cb610692565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610122565b3480156101fc57600080fd5b5061010d62030d4081565b34801561021357600080fd5b5061016a61077e565b34801561022857600080fd5b506101cb7f000000000000000000000000000000000000000000000000000000000000000081565b34801561025c57600080fd5b5061028061026b36600461190e565b60ce6020526000908152604090205460ff1681565b6040519015158152602001610122565b34801561029c57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006101cb565b3480156102cf57600080fd5b506102806102de36600461190e565b60cb6020526000908152604090205460ff1681565b3480156102ff57600080fd5b5061010d61030e366004611927565b61097b565b61016a61032136600461197b565b6109c7565b34801561033257600080fd5b5061037d60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b604051908152602001610122565b6104c47f00000000000000000000000000000000000000000000000000000000000000006103ba85858561097b565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061042660cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c6040516024016104429796959493929190611a46565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261130a565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561054960cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b8660405161055b959493929190611aa5565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b606061061a7f0000000000000000000000000000000000000000000000000000000000000000611398565b6106437f0000000000000000000000000000000000000000000000000000000000000000611398565b61066c7f0000000000000000000000000000000000000000000000000000000000000000611398565b60405160200161067e93929190611af3565b604051602081830303815290604052905090565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff215301610761576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6000547501000000000000000000000000000000000000000000900460ff16158080156107c9575060005460017401000000000000000000000000000000000000000090910460ff16105b806107fb5750303b1580156107fb575060005474010000000000000000000000000000000000000000900460ff166001145b610887576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610758565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000179055801561090d57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790555b6109156114cd565b801561097857600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b600062030d4061098c601085611b98565b6103e86109a16103f863ffffffff8716611b98565b6109ab9190611bf7565b6109b59190611c1e565b6109bf9190611c1e565b949350505050565b60f087901c60028110610a82576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a401610758565b8061ffff16600003610b77576000610ad3878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f92506115a6915050565b600081815260cb602052604090205490915060ff1615610b75576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c617965640000000000000000006064820152608401610758565b505b6000610bbd898989898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506115c592505050565b600081815260cf602052604090205490915060ff1615610c39576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610758565b600081815260cf6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610ceb600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330173ffffffffffffffffffffffffffffffffffffffff1614905090565b15610d2357853414610cff57610cff611c4a565b600081815260ce602052604090205460ff1615610d1e57610d1e611c4a565b610e75565b3415610dd7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a401610758565b600081815260ce602052604090205460ff16610e75576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c61796564000000000000000000000000000000006064820152608401610758565b610e7e876115e8565b15610f31576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a401610758565b600081815260cb602052604090205460ff1615610fd0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c61796564000000000000000000006064820152608401610758565b610fdc61afc886611c79565b5a101561106b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a20696e737566666963696560448201527f6e742067617320746f2072656c6179206d6573736167650000000000000000006064820152608401610758565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a161790556000611107886110bf61138861afc8611c91565b5a6110ca9190611c91565b8988888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061163d92505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905590508015156001036111a257600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a26112af565b600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016112af576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d657373616765000000000000000000000000000000000000006064820152608401610758565b50600090815260cf6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555050505050505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6040517fc2b3e5ac0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000169063c2b3e5ac90849061136090889088908790600401611ca8565b6000604051808303818588803b15801561137957600080fd5b505af115801561138d573d6000803e3d6000fd5b505050505050505050565b6060816000036113db57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561140557806113ef81611cf0565b91506113fe9050600a83611d28565b91506113df565b60008167ffffffffffffffff81111561142057611420611d3c565b6040519080825280601f01601f19166020018201604052801561144a576020820181803683370190505b5090505b84156109bf5761145f600183611c91565b915061146c600a86611d6b565b611477906030611c79565b60f81b81838151811061148c5761148c611d7f565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506114c6600a86611d28565b945061144e565b6000547501000000000000000000000000000000000000000000900460ff16611578576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610758565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b60006115b485858585611657565b805190602001209050949350505050565b60006115d58787878787876116f0565b8051906020012090509695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff8216301480611637575073ffffffffffffffffffffffffffffffffffffffff8216734200000000000000000000000000000000000016145b92915050565b600080600080845160208601878a8af19695505050505050565b6060848484846040516024016116709493929190611dae565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b606086868686868660405160240161170d96959493929190611df8565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146117b357600080fd5b919050565b60008083601f8401126117ca57600080fd5b50813567ffffffffffffffff8111156117e257600080fd5b6020830191508360208285010111156117fa57600080fd5b9250929050565b803563ffffffff811681146117b357600080fd5b6000806000806060858703121561182b57600080fd5b6118348561178f565b9350602085013567ffffffffffffffff81111561185057600080fd5b61185c878288016117b8565b909450925061186f905060408601611801565b905092959194509250565b60005b8381101561189557818101518382015260200161187d565b838111156118a4576000848401525b50505050565b600081518084526118c281602086016020860161187a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061190760208301846118aa565b9392505050565b60006020828403121561192057600080fd5b5035919050565b60008060006040848603121561193c57600080fd5b833567ffffffffffffffff81111561195357600080fd5b61195f868287016117b8565b9094509250611972905060208501611801565b90509250925092565b600080600080600080600060c0888a03121561199657600080fd5b873596506119a66020890161178f565b95506119b46040890161178f565b9450606088013593506080880135925060a088013567ffffffffffffffff8111156119de57600080fd5b6119ea8a828b016117b8565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a0830152611a9860c0830184866119fd565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201526000611ad56080830186886119fd565b905083604083015263ffffffff831660608301529695505050505050565b60008451611b0581846020890161187a565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611b41816001850160208a0161187a565b60019201918201528351611b5c81600284016020880161187a565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615611bbf57611bbf611b69565b02949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff80841680611c1257611c12611bc8565b92169190910492915050565b600067ffffffffffffffff808316818516808303821115611c4157611c41611b69565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60008219821115611c8c57611c8c611b69565b500190565b600082821015611ca357611ca3611b69565b500390565b73ffffffffffffffffffffffffffffffffffffffff8416815267ffffffffffffffff83166020820152606060408201526000611ce760608301846118aa565b95945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611d2157611d21611b69565b5060010190565b600082611d3757611d37611bc8565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082611d7a57611d7a611bc8565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525060806040830152611de760808301856118aa565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a0830152611e4360c08301846118aa565b9897505050505050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2CrossDomainMessengerStorageLayoutJSON), L2CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2CrossDomainMessenger"] = L2CrossDomainMessengerStorageLayout
	deployedBytecodes["L2CrossDomainMessenger"] = L2CrossDomainMessengerDeployedBin
}
