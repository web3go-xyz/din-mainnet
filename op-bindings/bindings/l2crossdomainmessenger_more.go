// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":24726,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":27112,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":27115,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":27726,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":26984,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":27104,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":27277,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_paused\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":27382,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":27397,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_status\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":27441,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":24778,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":24783,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":24788,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":24791,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":24794,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":24799,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"receivedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":24804,\"contract\":\"contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_array(t_uint256)42_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)42_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[42]\",\"numberOfBytes\":\"1344\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L2CrossDomainMessengerDeployedBin = "0x6080604052600436106101755760003560e01c80638129fc1c116100cb578063b28ade251161007f578063ecc7042811610059578063ecc70428146103f2578063f2fde38b14610457578063f69f81511461047757600080fd5b8063b28ade251461038b578063d764ad0b146103ab578063db505d80146103be57600080fd5b80638da5cb5b116100b05780638da5cb5b146102fd578063a711986914610328578063b1b1b2091461035b57600080fd5b80638129fc1c146102d35780638456cb59146102e857600080fd5b80633f827a5a1161012d5780636e296e45116101075780636e296e451461026d578063715018a6146102a75780637dea7cc3146102bc57600080fd5b80633f827a5a146101ff57806354fd4d50146102275780635c975abb1461024957600080fd5b80632828d7e81161015e5780632828d7e8146101bf5780633dbb202b146101d55780633f4ba83a146101ea57600080fd5b8063028f85f71461017a5780630c568498146101a9575b600080fd5b34801561018657600080fd5b5061018f601081565b60405163ffffffff90911681526020015b60405180910390f35b3480156101b557600080fd5b5061018f6103e881565b3480156101cb57600080fd5b5061018f6103f881565b6101e86101e3366004611cc9565b6104a7565b005b3480156101f657600080fd5b506101e8610711565b34801561020b57600080fd5b50610214600181565b60405161ffff90911681526020016101a0565b34801561023357600080fd5b5061023c610723565b6040516101a09190611da8565b34801561025557600080fd5b5060655460ff165b60405190151581526020016101a0565b34801561027957600080fd5b506102826107c6565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a0565b3480156102b357600080fd5b506101e86108b2565b3480156102c857600080fd5b5061018f62030d4081565b3480156102df57600080fd5b506101e86108c4565b3480156102f457600080fd5b506101e8610ac1565b34801561030957600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff16610282565b34801561033457600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610282565b34801561036757600080fd5b5061025d610376366004611dc2565b60cb6020526000908152604090205460ff1681565b34801561039757600080fd5b5061018f6103a6366004611ddb565b610ad1565b6101e86103b9366004611e2f565b610b17565b3480156103ca57600080fd5b506102827f000000000000000000000000000000000000000000000000000000000000000081565b3480156103fe57600080fd5b5061044960cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6040519081526020016101a0565b34801561046357600080fd5b506101e8610472366004611eb1565b6111f9565b34801561048357600080fd5b5061025d610492366004611dc2565b60ce6020526000908152604090205460ff1681565b6105e67f00000000000000000000000000000000000000000000000000000000000000006104d6858585610ad1565b63ffffffff16347fd764ad0b0000000000000000000000000000000000000000000000000000000061054860cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c6040516024016105649796959493929190611f15565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526112c9565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561066b60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b8660405161067d959493929190611f74565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b610719611357565b6107216113d8565b565b606061074e7f0000000000000000000000000000000000000000000000000000000000000000611455565b6107777f0000000000000000000000000000000000000000000000000000000000000000611455565b6107a07f0000000000000000000000000000000000000000000000000000000000000000611455565b6040516020016107b293929190611fc2565b604051602081830303815290604052905090565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff215301610895576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6108ba611357565b610721600061158a565b6000547501000000000000000000000000000000000000000000900460ff161580801561090f575060005460017401000000000000000000000000000000000000000090910460ff16105b806109415750303b158015610941575060005474010000000000000000000000000000000000000000900460ff166001145b6109cd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161088c565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790558015610a5357600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790555b610a5b611601565b8015610abe57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b610ac9611357565b6107216116f8565b600062030d40610ae2601085612067565b6103e8610af16103f886612067565b610afb91906120c2565b610b0591906120e5565b610b0f91906120e5565b949350505050565b600260975403610b83576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161088c565b6002609755610b90611753565b60f087901c60018114610c4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605560248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2031206d657373616765732061726520737570706f72746564206166746560648201527f722074686520426564726f636b20757067726164650000000000000000000000608482015260a40161088c565b6000610c91898989898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506117c092505050565b905073ffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330181167f000000000000000000000000000000000000000000000000000000000000000090911603610d0a57853414610d0557610d0561210d565b610e5c565b3415610dbe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a40161088c565b600081815260ce602052604090205460ff16610e5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c6179656400000000000000000000000000000000606482015260840161088c565b610e65876117e3565b15610f18576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a40161088c565b600081815260cb602052604090205460ff1615610fb7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c6179656400000000000000000000606482015260840161088c565b610fc361afc88661213c565b5a1015611052576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a20696e737566666963696560448201527f6e742067617320746f2072656c6179206d657373616765000000000000000000606482015260840161088c565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a1617905560006110ee886110a661138861afc8612154565b5a6110b19190612154565b8988888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061183892505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055905080151560010361118957600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a26111e8565b600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a25b505060016097555050505050505050565b611201611357565b73ffffffffffffffffffffffffffffffffffffffff81166112a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161088c565b610abe8161158a565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6040517fc2b3e5ac0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000169063c2b3e5ac90849061131f9088908890879060040161216b565b6000604051808303818588803b15801561133857600080fd5b505af115801561134c573d6000803e3d6000fd5b505050505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610721576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161088c565b6113e0611852565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60608160000361149857505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156114c257806114ac816121b3565b91506114bb9050600a836121eb565b915061149c565b60008167ffffffffffffffff8111156114dd576114dd6121ff565b6040519080825280601f01601f191660200182016040528015611507576020820181803683370190505b5090505b8415610b0f5761151c600183612154565b9150611529600a8661222e565b61153490603061213c565b60f81b81838151811061154957611549612242565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611583600a866121eb565b945061150b565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000547501000000000000000000000000000000000000000000900460ff166116ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088c565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead1790556116e06118be565b6116e8611969565b6116f0611a1d565b610721611af2565b611700611753565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861142b3390565b60655460ff1615610721576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161088c565b60006117d0878787878787611ba4565b8051906020012090509695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff8216301480611832575073ffffffffffffffffffffffffffffffffffffffff8216734200000000000000000000000000000000000016145b92915050565b600080600080845160208601878a8af19695505050505050565b60655460ff16610721576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161088c565b6000547501000000000000000000000000000000000000000000900460ff16610721576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088c565b6000547501000000000000000000000000000000000000000000900460ff16611a14576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088c565b6107213361158a565b6000547501000000000000000000000000000000000000000000900460ff16611ac8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088c565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6000547501000000000000000000000000000000000000000000900460ff16611b9d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161088c565b6001609755565b6060868686868686604051602401611bc196959493929190612271565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611c6757600080fd5b919050565b60008083601f840112611c7e57600080fd5b50813567ffffffffffffffff811115611c9657600080fd5b602083019150836020828501011115611cae57600080fd5b9250929050565b803563ffffffff81168114611c6757600080fd5b60008060008060608587031215611cdf57600080fd5b611ce885611c43565b9350602085013567ffffffffffffffff811115611d0457600080fd5b611d1087828801611c6c565b9094509250611d23905060408601611cb5565b905092959194509250565b60005b83811015611d49578181015183820152602001611d31565b83811115611d58576000848401525b50505050565b60008151808452611d76816020860160208601611d2e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611dbb6020830184611d5e565b9392505050565b600060208284031215611dd457600080fd5b5035919050565b600080600060408486031215611df057600080fd5b833567ffffffffffffffff811115611e0757600080fd5b611e1386828701611c6c565b9094509250611e26905060208501611cb5565b90509250925092565b600080600080600080600060c0888a031215611e4a57600080fd5b87359650611e5a60208901611c43565b9550611e6860408901611c43565b9450606088013593506080880135925060a088013567ffffffffffffffff811115611e9257600080fd5b611e9e8a828b01611c6c565b989b979a50959850939692959293505050565b600060208284031215611ec357600080fd5b611dbb82611c43565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a0830152611f6760c083018486611ecc565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201526000611fa4608083018688611ecc565b905083604083015263ffffffff831660608301529695505050505050565b60008451611fd4818460208901611d2e565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551612010816001850160208a01611d2e565b6001920191820152835161202b816002840160208801611d2e565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600063ffffffff8083168185168183048111821515161561208a5761208a612038565b02949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600063ffffffff808416806120d9576120d9612093565b92169190910492915050565b600063ffffffff80831681851680830382111561210457612104612038565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b6000821982111561214f5761214f612038565b500190565b60008282101561216657612166612038565b500390565b73ffffffffffffffffffffffffffffffffffffffff8416815267ffffffffffffffff831660208201526060604082015260006121aa6060830184611d5e565b95945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036121e4576121e4612038565b5060010190565b6000826121fa576121fa612093565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008261223d5761223d612093565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a08301526122bc60c0830184611d5e565b9897505050505050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2CrossDomainMessengerStorageLayoutJSON), L2CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2CrossDomainMessenger"] = L2CrossDomainMessengerStorageLayout
	deployedBytecodes["L2CrossDomainMessenger"] = L2CrossDomainMessengerDeployedBin
}
