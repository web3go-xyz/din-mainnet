// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1019_storage\"},{\"astId\":1004,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1018_storage\"},{\"astId\":1006,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_paused\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1018_storage\"},{\"astId\":1008,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_status\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)1018_storage\"},{\"astId\":1010,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":1015,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_array(t_uint256)1017_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1017_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[42]\",\"numberOfBytes\":\"1344\"},\"t_array(t_uint256)1018_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1019_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L2CrossDomainMessengerDeployedBin = "0x6080604052600436106101755760003560e01c80638129fc1c116100cb578063a71198691161007f578063d764ad0b11610059578063d764ad0b14610413578063ecc7042814610426578063f2fde38b1461048b57600080fd5b8063a711986914610390578063b1b1b209146103c3578063b28ade25146103f357600080fd5b80638da5cb5b116100b05780638da5cb5b146103015780639fce812c1461032c578063a4e7f8bd1461036057600080fd5b80638129fc1c146102d75780638456cb59146102ec57600080fd5b80633f827a5a1161012d5780636e296e45116101075780636e296e4514610271578063715018a6146102ab5780637dea7cc3146102c057600080fd5b80633f827a5a1461020357806354fd4d501461022b5780635c975abb1461024d57600080fd5b80632828d7e81161015e5780632828d7e8146101c35780633dbb202b146101d95780633f4ba83a146101ee57600080fd5b8063028f85f71461017a5780630c568498146101ad575b600080fd5b34801561018657600080fd5b5061018f601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101b957600080fd5b5061018f6103e881565b3480156101cf57600080fd5b5061018f6103f881565b6101ec6101e7366004611f94565b6104ab565b005b3480156101fa57600080fd5b506101ec61070f565b34801561020f57600080fd5b50610218600181565b60405161ffff90911681526020016101a4565b34801561023757600080fd5b50610240610721565b6040516101a49190612073565b34801561025957600080fd5b5060655460ff165b60405190151581526020016101a4565b34801561027d57600080fd5b506102866107c4565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a4565b3480156102b757600080fd5b506101ec6108b0565b3480156102cc57600080fd5b5061018f62030d4081565b3480156102e357600080fd5b506101ec6108c2565b3480156102f857600080fd5b506101ec610abf565b34801561030d57600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff16610286565b34801561033857600080fd5b506102867f000000000000000000000000000000000000000000000000000000000000000081565b34801561036c57600080fd5b5061026161037b36600461208d565b60ce6020526000908152604090205460ff1681565b34801561039c57600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610286565b3480156103cf57600080fd5b506102616103de36600461208d565b60cb6020526000908152604090205460ff1681565b3480156103ff57600080fd5b5061018f61040e3660046120a6565b610acf565b6101ec6104213660046120fa565b610b1b565b34801561043257600080fd5b5061047d60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6040519081526020016101a4565b34801561049757600080fd5b506101ec6104a636600461217c565b61133e565b6105e47f00000000000000000000000000000000000000000000000000000000000000006104da858585610acf565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061054660cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c60405160240161056297969594939291906121e0565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261140e565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561066960cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b8660405161067b95949392919061223f565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b61071761149c565b61071f61151d565b565b606061074c7f000000000000000000000000000000000000000000000000000000000000000061159a565b6107757f000000000000000000000000000000000000000000000000000000000000000061159a565b61079e7f000000000000000000000000000000000000000000000000000000000000000061159a565b6040516020016107b09392919061228d565b604051602081830303815290604052905090565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff215301610893576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6108b861149c565b61071f60006116cf565b6000547501000000000000000000000000000000000000000000900460ff161580801561090d575060005460017401000000000000000000000000000000000000000090910460ff16105b8061093f5750303b15801561093f575060005474010000000000000000000000000000000000000000900460ff166001145b6109cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161088a565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790558015610a5157600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790555b610a59611746565b8015610abc57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b610ac761149c565b61071f61183d565b600062030d40610ae0601085612332565b6103e8610af56103f863ffffffff8716612332565b610aff9190612391565b610b0991906123b8565b610b1391906123b8565b949350505050565b610b5f878787878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061189892505050565b600081815263cafebabe60205260409020805415610bb2576308c379a06000526020805260006040527f1f5265656e7472616e637947756172643a207265656e7472616e742063616c6c605f526064601cfd5b60018155610bbe611966565b60f089901c6000819003610cc2576000610c1e898b88888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508e6119d3565b600081815260cb602052604090205490915060ff1615610cc0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c61796564000000000000000000606482015260840161088a565b505b6000610d088b8b8b8b8b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506119f292505050565b905073ffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330181167f000000000000000000000000000000000000000000000000000000000000000090911603610da057873414610d7c57610d7c6123e4565b600081815260ce602052604090205460ff1615610d9b57610d9b6123e4565b610ef2565b3415610e54576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a40161088a565b600081815260ce602052604090205460ff16610ef2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c6179656400000000000000000000000000000000606482015260840161088a565b610efb89611a15565b15610fae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a40161088a565b600081815260cb602052604090205460ff161561104d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c6179656400000000000000000000606482015260840161088a565b61105961afc888612413565b5a10156110e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a20696e737566666963696560448201527f6e742067617320746f2072656c6179206d657373616765000000000000000000606482015260840161088a565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8c1617905560006111848a61113c61138861afc861242b565b5a611147919061242b565b8b8a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611a6a92505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055905080151560010361121f57600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a261132c565b600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff320161132c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d65737361676500000000000000000000000000000000000000606482015260840161088a565b50505060008155505050505050505050565b61134661149c565b73ffffffffffffffffffffffffffffffffffffffff81166113e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161088a565b610abc816116cf565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6040517fc2b3e5ac0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000169063c2b3e5ac90849061146490889088908790600401612442565b6000604051808303818588803b15801561147d57600080fd5b505af1158015611491573d6000803e3d6000fd5b505050505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461071f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161088a565b611525611a84565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6060816000036115dd57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561160757806115f18161248a565b91506116009050600a836124c2565b91506115e1565b60008167ffffffffffffffff811115611622576116226124d6565b6040519080825280601f01601f19166020018201604052801561164c576020820181803683370190505b5090505b8415610b135761166160018361242b565b915061166e600a86612505565b611679906030612413565b60f81b81838151811061168e5761168e612519565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506116c8600a866124c2565b9450611650565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000547501000000000000000000000000000000000000000000900460ff166117f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088a565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055611825611af0565b61182d611b9b565b611835611c4f565b61071f611d24565b611845611966565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586115703390565b600060f087901c8082036118ba576118b28688858b6119d3565b91505061195c565b8061ffff166001036118d4576118b28888888888886119f2565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f48617368696e673a20756e6b6e6f776e2063726f737320646f6d61696e206d6560448201527f73736167652076657273696f6e00000000000000000000000000000000000000606482015260840161088a565b9695505050505050565b60655460ff161561071f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161088a565b60006119e185858585611dd6565b805190602001209050949350505050565b6000611a02878787878787611e6f565b8051906020012090509695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff8216301480611a64575073ffffffffffffffffffffffffffffffffffffffff8216734200000000000000000000000000000000000016145b92915050565b600080600080845160208601878a8af19695505050505050565b60655460ff1661071f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161088a565b6000547501000000000000000000000000000000000000000000900460ff1661071f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088a565b6000547501000000000000000000000000000000000000000000900460ff16611c46576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088a565b61071f336116cf565b6000547501000000000000000000000000000000000000000000900460ff16611cfa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088a565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6000547501000000000000000000000000000000000000000000900460ff16611dcf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088a565b6001609755565b606084848484604051602401611def9493929190612548565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b6060868686868686604051602401611e8c96959493929190612592565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611f3257600080fd5b919050565b60008083601f840112611f4957600080fd5b50813567ffffffffffffffff811115611f6157600080fd5b602083019150836020828501011115611f7957600080fd5b9250929050565b803563ffffffff81168114611f3257600080fd5b60008060008060608587031215611faa57600080fd5b611fb385611f0e565b9350602085013567ffffffffffffffff811115611fcf57600080fd5b611fdb87828801611f37565b9094509250611fee905060408601611f80565b905092959194509250565b60005b83811015612014578181015183820152602001611ffc565b83811115612023576000848401525b50505050565b60008151808452612041816020860160208601611ff9565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006120866020830184612029565b9392505050565b60006020828403121561209f57600080fd5b5035919050565b6000806000604084860312156120bb57600080fd5b833567ffffffffffffffff8111156120d257600080fd5b6120de86828701611f37565b90945092506120f1905060208501611f80565b90509250925092565b600080600080600080600060c0888a03121561211557600080fd5b8735965061212560208901611f0e565b955061213360408901611f0e565b9450606088013593506080880135925060a088013567ffffffffffffffff81111561215d57600080fd5b6121698a828b01611f37565b989b979a50959850939692959293505050565b60006020828403121561218e57600080fd5b61208682611f0e565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a083015261223260c083018486612197565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815260806020820152600061226f608083018688612197565b905083604083015263ffffffff831660608301529695505050505050565b6000845161229f818460208901611ff9565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516122db816001850160208a01611ff9565b600192019182015283516122f6816002840160208801611ff9565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff8083168185168183048111821515161561235957612359612303565b02949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff808416806123ac576123ac612362565b92169190910492915050565b600067ffffffffffffffff8083168185168083038211156123db576123db612303565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b6000821982111561242657612426612303565b500190565b60008282101561243d5761243d612303565b500390565b73ffffffffffffffffffffffffffffffffffffffff8416815267ffffffffffffffff831660208201526060604082015260006124816060830184612029565b95945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036124bb576124bb612303565b5060010190565b6000826124d1576124d1612362565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008261251457612514612362565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250608060408301526125816080830185612029565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a08301526125dd60c0830184612029565b9897505050505050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2CrossDomainMessengerStorageLayoutJSON), L2CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2CrossDomainMessenger"] = L2CrossDomainMessengerStorageLayout
	deployedBytecodes["L2CrossDomainMessenger"] = L2CrossDomainMessengerDeployedBin
}
