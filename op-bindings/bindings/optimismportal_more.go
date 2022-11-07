// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const OptimismPortalStorageLayoutJSON = "{\"storage\":[{\"astId\":30083,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":30086,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1618,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"params\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_struct(ResourceParams)1588_storage\"},{\"astId\":1623,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_uint256)48_storage\"},{\"astId\":1189,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"l2Sender\",\"offset\":0,\"slot\":\"50\",\"type\":\"t_address\"},{\"astId\":1202,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"finalizedWithdrawals\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_mapping(t_bytes32,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)48_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[48]\",\"numberOfBytes\":\"1536\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_struct(ResourceParams)1588_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceParams\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var OptimismPortalStorageLayout = new(solc.StorageLayout)

var OptimismPortalDeployedBin = "0x6080604052600436106101115760003560e01c80638b4c40b0116100a5578063ca3e99ba11610074578063cff0ab9611610059578063cff0ab9614610332578063e9e05c42146103d3578063f4daa291146103e657600080fd5b8063ca3e99ba14610308578063cd7c97891461031d57600080fd5b80638b4c40b0146101365780639bf62d821461027b578063a14238e7146102a8578063c4fc4798146102e857600080fd5b8063683d95c8116100e1578063683d95c81461021b5780636bb0291e1461023b5780638129fc1c14610250578063867ead131461026557600080fd5b80621c2ff61461013d57806313620abd1461019b57806354fd4d50146101d457806364b79208146101f657600080fd5b36610138576101363334620186a060006040518060200160405280600081525061041a565b005b600080fd5b34801561014957600080fd5b506101717f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101a757600080fd5b506101b3633b9aca0081565b6040516fffffffffffffffffffffffffffffffff9091168152602001610192565b3480156101e057600080fd5b506101e96108e7565b6040516101929190613b25565b34801561020257600080fd5b5061020d627a120081565b604051908152602001610192565b34801561022757600080fd5b50610136610236366004613cfa565b61098a565b34801561024757600080fd5b5061020d600481565b34801561025c57600080fd5b50610136610fe2565b34801561027157600080fd5b5061020d61271081565b34801561028757600080fd5b506032546101719073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102b457600080fd5b506102d86102c3366004613dee565b60336020526000908152604090205460ff1681565b6040519015158152602001610192565b3480156102f457600080fd5b506102d8610303366004613dee565b6111a0565b34801561031457600080fd5b5061020d611265565b34801561032957600080fd5b5061020d600881565b34801561033e57600080fd5b5060015461039a906fffffffffffffffffffffffffffffffff81169067ffffffffffffffff7001000000000000000000000000000000008204811691780100000000000000000000000000000000000000000000000090041683565b604080516fffffffffffffffffffffffffffffffff909416845267ffffffffffffffff9283166020850152911690820152606001610192565b6101366103e1366004613e07565b61041a565b3480156103f257600080fd5b5061020d7f000000000000000000000000000000000000000000000000000000000000000081565b8260005a905083156104d15773ffffffffffffffffffffffffffffffffffffffff8716156104d157604080517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260248101919091527f4f7074696d69736d506f7274616c3a206d7573742073656e6420746f2061646460448201527f72657373283029207768656e206372656174696e67206120636f6e747261637460648201526084015b60405180910390fd5b333281146104f2575033731111000000000000000000000000000000001111015b6000348888888860405160200161050d959493929190613e94565b604051602081830303815290604052905060008973ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c328460405161057d9190613b25565b60405180910390a450506001546000906105bd907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1643613f28565b905080156107465760006105d56004627a1200613f6e565b6001546106009190700100000000000000000000000000000000900467ffffffffffffffff16613fd6565b9050600060086106146004627a1200613f6e565b6001546106349085906fffffffffffffffffffffffffffffffff1661404a565b61063e9190613f6e565b6106489190613f6e565b6001549091506000906106949061067e906106769085906fffffffffffffffffffffffffffffffff16614106565b612710611292565b6fffffffffffffffffffffffffffffffff6112ad565b905060018411156107075761070461067e670de0b6b3a76400006106f06106bc600883613f6e565b6106ce90670de0b6b3a7640000613fd6565b6106d960018a613f28565b6106eb90670de0b6b3a764000061417a565b6112bc565b6106fa908561404a565b6106769190613f6e565b90505b6fffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff4316021760015550505b60018054849190601090610779908490700100000000000000000000000000000000900467ffffffffffffffff166141b7565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550627a1200600160000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff161315610855576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5265736f757263654d65746572696e673a2063616e6e6f7420627579206d6f7260448201527f6520676173207468616e20617661696c61626c6520676173206c696d6974000060648201526084016104c8565b600154600090610881906fffffffffffffffffffffffffffffffff1667ffffffffffffffff86166141e3565b6fffffffffffffffffffffffffffffffff16905060006108a548633b9aca006112ed565b6108af908361421b565b905060005a6108be9086613f28565b9050808211156108da576108da6108d58284613f28565b6112fd565b5050505050505050505050565b60606109127f000000000000000000000000000000000000000000000000000000000000000061132b565b61093b7f000000000000000000000000000000000000000000000000000000000000000061132b565b6109647f000000000000000000000000000000000000000000000000000000000000000061132b565b6040516020016109769392919061422f565b604051602081830303815290604052905090565b60325473ffffffffffffffffffffffffffffffffffffffff1661dead14610a33576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4f7074696d69736d506f7274616c3a2063616e206f6e6c79207472696767657260448201527f206f6e65207769746864726177616c20706572207472616e73616374696f6e0060648201526084016104c8565b3073ffffffffffffffffffffffffffffffffffffffff16856040015173ffffffffffffffffffffffffffffffffffffffff1603610af2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4f7074696d69736d506f7274616c3a20796f752063616e6e6f742073656e642060448201527f6d6573736167657320746f2074686520706f7274616c20636f6e74726163740060648201526084016104c8565b6040517fa25ae557000000000000000000000000000000000000000000000000000000008152600481018590526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a25ae557906024016040805180830381865afa158015610b7f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba391906142a5565b9050610bae81611468565b610c3a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4f7074696d69736d506f7274616c3a2070726f706f73616c206973206e6f742060448201527f7965742066696e616c697a65640000000000000000000000000000000000000060648201526084016104c8565b610c51610c4c368690038601866142f4565b6114a2565b815114610ce0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4f7074696d69736d506f7274616c3a20696e76616c6964206f7574707574207260448201527f6f6f742070726f6f66000000000000000000000000000000000000000000000060648201526084016104c8565b6000610ceb876114fe565b9050610d05816040870135610d00868861435a565b61152e565b610d91576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4f7074696d69736d506f7274616c3a20696e76616c696420776974686472617760448201527f616c20696e636c7573696f6e2070726f6f66000000000000000000000000000060648201526084016104c8565b60008181526033602052604090205460ff1615610e30576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f4f7074696d69736d506f7274616c3a207769746864726177616c20686173206160448201527f6c7265616479206265656e2066696e616c697a6564000000000000000000000060648201526084016104c8565b600081815260336020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556080870151610e7990614e20906143de565b5a1015610f08576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f4f7074696d69736d506f7274616c3a20696e73756666696369656e742067617360448201527f20746f2066696e616c697a65207769746864726177616c00000000000000000060648201526084016104c8565b8660200151603260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000610f6b886040015189608001518a606001518b60a001516115f5565b603280547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905560405190915082907fdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b90610fd090841515815260200190565b60405180910390a25050505050505050565b600054610100900460ff16158080156110025750600054600160ff909116105b8061101c5750303b15801561101c575060005460ff166001145b6110a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104c8565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561110657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b603280547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905561113a61160f565b801561119d57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6040517fa25ae55700000000000000000000000000000000000000000000000000000000815260048101829052600090819073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063a25ae557906024016040805180830381865afa15801561122f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061125391906142a5565b905061125e81611468565b9392505050565b6112736004627a1200613f6e565b81565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6000818312156112a257816112a4565b825b90505b92915050565b60008183126112a257816112a4565b60006112a4670de0b6b3a7640000836112d4866116f2565b6112de919061404a565b6112e89190613f6e565b611936565b6000818310156112a257816112a4565b6000805a90505b825a6113109083613f28565b10156113265761131f826143f6565b9150611304565b505050565b60608160000361136e57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156113985780611382816143f6565b91506113919050600a8361421b565b9150611372565b60008167ffffffffffffffff8111156113b3576113b3613b38565b6040519080825280601f01601f1916602001820160405280156113dd576020820181803683370190505b5090505b8415611460576113f2600183613f28565b91506113ff600a8661442e565b61140a9060306143de565b60f81b81838151811061141f5761141f614442565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611459600a8661421b565b94506113e1565b949350505050565b60007f0000000000000000000000000000000000000000000000000000000000000000826020015161149a91906143de565b421192915050565b600081600001518260200151836040015184606001516040516020016114e1949392919093845260208401929092526040830152606082015260800190565b604051602081830303815290604052805190602001209050919050565b80516020808301516040808501516060860151608087015160a088015193516000976114e1979096959101614471565b604080516020810185905260009181018290528190606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152828252805160209182012090830181905292506115ec9101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152828201909152600182527f01000000000000000000000000000000000000000000000000000000000000006020830152908587611b75565b95945050505050565b600080600080845160208601878a8af19695505050505050565b600054610100900460ff166116a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104c8565b60408051606081018252633b9aca00808252600060208301524367ffffffffffffffff169190920181905278010000000000000000000000000000000000000000000000000217600155565b600080821361175d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f554e444546494e4544000000000000000000000000000000000000000000000060448201526064016104c8565b6000606061176a84611b99565b03609f8181039490941b90931c6c465772b2bbbb5f824b15207a3081018102606090811d6d0388eaa27412d5aca026815d636e018202811d6d0df99ac502031bf953eff472fdcc018202811d6d13cdffb29d51d99322bdff5f2211018202811d6d0a0f742023def783a307a986912e018202811d6d01920d8043ca89b5239253284e42018202811d6c0b7a86d7375468fac667a0a527016c29508e458543d8aa4df2abee7883018302821d6d0139601a2efabe717e604cbb4894018302821d6d02247f7a7b6594320649aa03aba1018302821d7fffffffffffffffffffffffffffffffffffffff73c0c716a594e00d54e3c4cbc9018302821d7ffffffffffffffffffffffffffffffffffffffdc7b88c420e53a9890533129f6f01830290911d7fffffffffffffffffffffffffffffffffffffff465fda27eb4d63ded474e5f832019091027ffffffffffffffff5f6af8f7b3396644f18e157960000000000000000000000000105711340daa0d5f769dba1915cef59f0815a5506027d0267a36c0c95b3975ab3ee5b203a7614a3f75373f047d803ae7b6687f2b393909302929092017d57115e47018c7177eebf7cd370a3356a1b7863008a5ae8028c72b88642840160ae1d92915050565b60007ffffffffffffffffffffffffffffffffffffffffffffffffdb731c958f34d94c1821361196757506000919050565b680755bf798b4a1bf1e582126119d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4558505f4f564552464c4f57000000000000000000000000000000000000000060448201526064016104c8565b6503782dace9d9604e83901b059150600060606bb17217f7d1cf79abc9e3b39884821b056b80000000000000000000000001901d6bb17217f7d1cf79abc9e3b39881029093037fffffffffffffffffffffffffffffffffffffffdbf3ccf1604d263450f02a550481018102606090811d6d0277594991cfc85f6e2461837cd9018202811d7fffffffffffffffffffffffffffffffffffffe5adedaa1cb095af9e4da10e363c018202811d6db1bbb201f443cf962f1a1d3db4a5018202811d7ffffffffffffffffffffffffffffffffffffd38dc772608b0ae56cce01296c0eb018202811d6e05180bb14799ab47a8a8cb2a527d57016d02d16720577bd19bf614176fe9ea6c10fe68e7fd37d0007b713f765084018402831d9081019084017ffffffffffffffffffffffffffffffffffffffe2c69812cf03b0763fd454a8f7e010290911d6e0587f503bb6ea29d25fcb7401964500190910279d835ebba824c98fb31b83b2ca45c000000000000000000000000010574029d9dc38563c32e5c2f6dc192ee70ef65f9978af30260c3939093039290921c92915050565b600080611b8186611c6f565b9050611b8f81868686611ca1565b9695505050505050565b6000808211611c04576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f554e444546494e4544000000000000000000000000000000000000000000000060448201526064016104c8565b5060016fffffffffffffffffffffffffffffffff821160071b82811c67ffffffffffffffff1060061b1782811c63ffffffff1060051b1782811c61ffff1060041b1782811c60ff10600390811b90911783811c600f1060021b1783811c909110821b1791821c111790565b60608180519060200120604051602001611c8b91815260200190565b6040516020818303038152906040529050919050565b6000806000611cb1878686611cde565b91509150818015611cd357508051602080830191909120875191880191909120145b979650505050505050565b600060606000611ced85611dfc565b90506000806000611cff848a89611eeb565b81519295509093509150158080611d135750815b611d9f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4d65726b6c65547269653a2070726f76696465642070726f6f6620697320696e60448201527f76616c696400000000000000000000000000000000000000000000000000000060648201526084016104c8565b600081611dbb5760405180602001604052806000815250611de7565b611de786611dca600188613f28565b81518110611dda57611dda614442565b60200260200101516125ef565b805115159c909b509950505050505050505050565b805160609060008167ffffffffffffffff811115611e1c57611e1c613b38565b604051908082528060200260200182016040528015611e6157816020015b6040805180820190915260608082526020820152815260200190600190039081611e3a5790505b50905060005b82811015611ee3576040518060400160405280868381518110611e8c57611e8c614442565b60200260200101518152602001611ebb878481518110611eae57611eae614442565b6020026020010151612626565b815250828281518110611ed057611ed0614442565b6020908102919091010152600101611e67565b509392505050565b60006060818080611efb87612639565b9050600086604051602001611f1291815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152828201909152606080835260208301529150600090819060005b8c51811015612571578c8181518110611f7257611f72614442565b602002602001015191508284611f8891906143de565b9350611f956001886143de565b9650836000036120505781518051602091820120604051611fe592611fbf92910190815260200190565b604051602081830303815290604052868051602091820120825192909101919091201490565b61204b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4d65726b6c65547269653a20696e76616c696420726f6f74206861736800000060448201526064016104c8565b6121a7565b815151602011612106578151805160209182012060405161207a92611fbf92910190815260200190565b61204b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f4d65726b6c65547269653a20696e76616c6964206c6172676520696e7465726e60448201527f616c20686173680000000000000000000000000000000000000000000000000060648201526084016104c8565b8151855160208088019190912082519190920120146121a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4d65726b6c65547269653a20696e76616c696420696e7465726e616c206e6f6460448201527f652068617368000000000000000000000000000000000000000000000000000060648201526084016104c8565b6121b3601060016143de565b8260200151510361222557855184146125715760008685815181106121da576121da614442565b602001015160f81c60f81b60f81c9050600083602001518260ff168151811061220557612205614442565b60200260200101519050612218816127d4565b965060019450505061255f565b6002826020015151036124d757600061223d836127f9565b905060008160008151811061225457612254614442565b016020015160f81c9050600061226b6002836144c8565b6122769060026144ea565b90506000612287848360ff1661281d565b905060006122958b8a61281d565b905060006122a38383612853565b9050825182511015612337576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f4d65726b6c65547269653a20696e76616c6964206b6579206c656e677468206660448201527f6f72206c656166206f7220657874656e73696f6e206e6f64650000000000000060648201526084016104c8565b60ff85166002148061234c575060ff85166003145b156123b6578083511480156123615750808251145b1561237357612370818b6143de565b99505b6040518060400160405280600181526020017f80000000000000000000000000000000000000000000000000000000000000008152509a50505050505050612571565b60ff851615806123c9575060ff85166001145b1561244f5782518114612419576040518060400160405280600181526020017f80000000000000000000000000000000000000000000000000000000000000008152509a50505050505050612571565b612440886020015160018151811061243357612433614442565b60200260200101516127d4565b9a50975061255f945050505050565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4d65726b6c65547269653a2072656365697665642061206e6f6465207769746860448201527f20616e20756e6b6e6f776e20707265666978000000000000000000000000000060648201526084016104c8565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f4d65726b6c65547269653a20726563656976656420616e20756e70617273656160448201527f626c65206e6f646500000000000000000000000000000000000000000000000060648201526084016104c8565b80612569816143f6565b915050611f57565b508561257d868561281d565b60408051808201909152600181527f80000000000000000000000000000000000000000000000000000000000000006020918201528651908701207f56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b4211498509850985050505050505093509350939050565b602081015180516060916112a79161260990600190613f28565b8151811061261957612619614442565b6020026020010151612902565b60606112a761263483612a62565b612b4b565b8051606090600061264b82600261417a565b67ffffffffffffffff81111561266357612663613b38565b6040519080825280601f01601f19166020018201604052801561268d576020820181803683370190505b5090506000805b838110156127ca578581815181106126ae576126ae614442565b6020910101517fff000000000000000000000000000000000000000000000000000000000000008116925060041c7f0ff0000000000000000000000000000000000000000000000000000000000000168361270a83600261417a565b8151811061271a5761271a614442565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f0f0000000000000000000000000000000000000000000000000000000000000082168361277883600261417a565b6127839060016143de565b8151811061279357612793614442565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600101612694565b5090949350505050565b606060208260000151106127f0576127eb82612902565b6112a7565b6112a782612daf565b60606112a7612818836020015160008151811061261957612619614442565b612639565b60608251821061283c57506040805160208101909152600081526112a7565b6112a4838384865161284e9190613f28565b612dc5565b6000806000835185511061286857835161286b565b84515b90505b80821080156128f2575083828151811061288a5761288a614442565b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168583815181106128c9576128c9614442565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016145b15611ee35781600101915061286e565b6060600080600061291285612f9d565b91945092509050600081600181111561292d5761292d61450d565b146129ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f524c505265616465723a206465636f646564206974656d207479706520666f7260448201527f206279746573206973206e6f7420612064617461206974656d0000000000000060648201526084016104c8565b6129c482846143de565b855114612a53576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f524c505265616465723a2062797465732076616c756520636f6e7461696e732060448201527f616e20696e76616c69642072656d61696e64657200000000000000000000000060648201526084016104c8565b6115ec85602001518484613a0a565b60408051808201909152600080825260208201526000825111612b2d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f524c505265616465723a206c656e677468206f6620616e20524c50206974656d60448201527f206d7573742062652067726561746572207468616e207a65726f20746f20626560648201527f206465636f6461626c6500000000000000000000000000000000000000000000608482015260a4016104c8565b50604080518082019091528151815260209182019181019190915290565b60606000806000612b5b85612f9d565b919450925090506001816001811115612b7657612b7661450d565b14612c03576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603860248201527f524c505265616465723a206465636f646564206974656d207479706520666f7260448201527f206c697374206973206e6f742061206c697374206974656d000000000000000060648201526084016104c8565b8451612c0f83856143de565b14612c9c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f524c505265616465723a206c697374206974656d2068617320616e20696e766160448201527f6c696420646174612072656d61696e646572000000000000000000000000000060648201526084016104c8565b6040805160208082526104208201909252600091816020015b6040805180820190915260008082526020820152815260200190600190039081612cb55790505090506000845b8751811015612da357600080612d286040518060400160405280858d60000151612d0c9190613f28565b8152602001858d60200151612d2191906143de565b9052612f9d565b509150915060405180604001604052808383612d4491906143de565b8152602001848c60200151612d5991906143de565b815250858581518110612d6e57612d6e614442565b6020908102919091010152612d846001856143de565b9350612d9081836143de565b612d9a90846143de565b92505050612ce2565b50815295945050505050565b60606112a7826020015160008460000151613a0a565b60608182601f011015612e34576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f736c6963655f6f766572666c6f7700000000000000000000000000000000000060448201526064016104c8565b828284011015612ea0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f736c6963655f6f766572666c6f7700000000000000000000000000000000000060448201526064016104c8565b81830184511015612f0d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f736c6963655f6f75744f66426f756e647300000000000000000000000000000060448201526064016104c8565b606082158015612f2c5760405191506000825260208201604052612f94565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015612f65578051835260209283019201612f4d565b5050858452601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016604052505b50949350505050565b60008060008084600001511161305b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f524c505265616465723a206c656e677468206f6620616e20524c50206974656d60448201527f206d7573742062652067726561746572207468616e207a65726f20746f20626560648201527f206465636f6461626c6500000000000000000000000000000000000000000000608482015260a4016104c8565b6020840151805160001a607f8111613080576000600160009450945094505050613a03565b60b7811161328e576000613095608083613f28565b905080876000015111613150576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604e60248201527f524c505265616465723a206c656e677468206f6620636f6e74656e74206d757360448201527f742062652067726561746572207468616e20737472696e67206c656e6774682060648201527f2873686f727420737472696e6729000000000000000000000000000000000000608482015260a4016104c8565b6001838101517fff000000000000000000000000000000000000000000000000000000000000001690821415806131c957507f80000000000000000000000000000000000000000000000000000000000000007fff00000000000000000000000000000000000000000000000000000000000000821610155b61327b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f524c505265616465723a20696e76616c6964207072656669782c2073696e676c60448201527f652062797465203c203078383020617265206e6f74207072656669786564202860648201527f73686f727420737472696e672900000000000000000000000000000000000000608482015260a4016104c8565b5060019550935060009250613a03915050565b60bf81116135dc5760006132a360b783613f28565b90508087600001511161335e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605160248201527f524c505265616465723a206c656e677468206f6620636f6e74656e74206d757360448201527f74206265203e207468616e206c656e677468206f6620737472696e67206c656e60648201527f67746820286c6f6e6720737472696e6729000000000000000000000000000000608482015260a4016104c8565b60018301517fff0000000000000000000000000000000000000000000000000000000000000016600081900361343c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f524c505265616465723a206c656e677468206f6620636f6e74656e74206d757360448201527f74206e6f74206861766520616e79206c656164696e67207a65726f7320286c6f60648201527f6e6720737472696e672900000000000000000000000000000000000000000000608482015260a4016104c8565b600184015160088302610100031c60378111613500576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f524c505265616465723a206c656e677468206f6620636f6e74656e74206d757360448201527f742062652067726561746572207468616e20353520627974657320286c6f6e6760648201527f20737472696e6729000000000000000000000000000000000000000000000000608482015260a4016104c8565b61350a81846143de565b8951116135bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604c60248201527f524c505265616465723a206c656e677468206f6620636f6e74656e74206d757360448201527f742062652067726561746572207468616e20746f74616c206c656e677468202860648201527f6c6f6e6720737472696e67290000000000000000000000000000000000000000608482015260a4016104c8565b6135ca8360016143de565b9750955060009450613a039350505050565b60f781116136bd5760006135f160c083613f28565b9050808760000151116136ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f524c505265616465723a206c656e677468206f6620636f6e74656e74206d757360448201527f742062652067726561746572207468616e206c697374206c656e67746820287360648201527f686f7274206c6973742900000000000000000000000000000000000000000000608482015260a4016104c8565b600195509350849250613a03915050565b60006136ca60f783613f28565b905080876000015111613785576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f524c505265616465723a206c656e677468206f6620636f6e74656e74206d757360448201527f74206265203e207468616e206c656e677468206f66206c697374206c656e677460648201527f6820286c6f6e67206c6973742900000000000000000000000000000000000000608482015260a4016104c8565b60018301517fff00000000000000000000000000000000000000000000000000000000000000166000819003613863576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f524c505265616465723a206c656e677468206f6620636f6e74656e74206d757360448201527f74206e6f74206861766520616e79206c656164696e67207a65726f7320286c6f60648201527f6e67206c69737429000000000000000000000000000000000000000000000000608482015260a4016104c8565b600184015160088302610100031c60378111613927576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604660248201527f524c505265616465723a206c656e677468206f6620636f6e74656e74206d757360448201527f742062652067726561746572207468616e20353520627974657320286c6f6e6760648201527f206c697374290000000000000000000000000000000000000000000000000000608482015260a4016104c8565b61393181846143de565b8951116139e6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f524c505265616465723a206c656e677468206f6620636f6e74656e74206d757360448201527f742062652067726561746572207468616e20746f74616c206c656e677468202860648201527f6c6f6e67206c6973742900000000000000000000000000000000000000000000608482015260a4016104c8565b6139f18360016143de565b9750955060019450613a039350505050565b9193909250565b606060008267ffffffffffffffff811115613a2757613a27613b38565b6040519080825280601f01601f191660200182016040528015613a51576020820181803683370190505b50905082600003613a6357905061125e565b6000613a6f85876143de565b90506020820160005b85811015613a90578281015182820152602001613a78565b85811115613a9f576000868301525b50919695505050505050565b60005b83811015613ac6578181015183820152602001613aae565b83811115613ad5576000848401525b50505050565b60008151808452613af3816020860160208601613aab565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006112a46020830184613adb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160c0810167ffffffffffffffff81118282101715613b8a57613b8a613b38565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613bd757613bd7613b38565b604052919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114613c0357600080fd5b919050565b600082601f830112613c1957600080fd5b813567ffffffffffffffff811115613c3357613c33613b38565b613c6460207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601613b90565b818152846020838601011115613c7957600080fd5b816020850160208301376000918101602001919091529392505050565b600060808284031215613ca857600080fd5b50919050565b60008083601f840112613cc057600080fd5b50813567ffffffffffffffff811115613cd857600080fd5b6020830191508360208260051b8501011115613cf357600080fd5b9250929050565b600080600080600060e08688031215613d1257600080fd5b853567ffffffffffffffff80821115613d2a57600080fd5b9087019060c0828a031215613d3e57600080fd5b613d46613b67565b82358152613d5660208401613bdf565b6020820152613d6760408401613bdf565b6040820152606083013560608201526080830135608082015260a083013582811115613d9257600080fd5b613d9e8b828601613c08565b60a083015250965060208801359550613dba8960408a01613c96565b945060c0880135915080821115613dd057600080fd5b50613ddd88828901613cae565b969995985093965092949392505050565b600060208284031215613e0057600080fd5b5035919050565b600080600080600060a08688031215613e1f57600080fd5b613e2886613bdf565b945060208601359350604086013567ffffffffffffffff8082168214613e4d57600080fd5b9093506060870135908115158214613e6457600080fd5b90925060808701359080821115613e7a57600080fd5b50613e8788828901613c08565b9150509295509295909350565b8581528460208201527fffffffffffffffff0000000000000000000000000000000000000000000000008460c01b16604082015282151560f81b604882015260008251613ee8816049850160208701613aab565b919091016049019695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015613f3a57613f3a613ef9565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082613f7d57613f7d613f3f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83147f800000000000000000000000000000000000000000000000000000000000000083141615613fd157613fd1613ef9565b500590565b6000808312837f80000000000000000000000000000000000000000000000000000000000000000183128115161561401057614010613ef9565b837f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01831381161561404457614044613ef9565b50500390565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60008413600084138583048511828216161561408b5761408b613ef9565b7f800000000000000000000000000000000000000000000000000000000000000060008712868205881281841616156140c6576140c6613ef9565b600087129250878205871284841616156140e2576140e2613ef9565b878505871281841616156140f8576140f8613ef9565b505050929093029392505050565b6000808212827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0384138115161561414057614140613ef9565b827f800000000000000000000000000000000000000000000000000000000000000003841281161561417457614174613ef9565b50500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156141b2576141b2613ef9565b500290565b600067ffffffffffffffff8083168185168083038211156141da576141da613ef9565b01949350505050565b60006fffffffffffffffffffffffffffffffff8083168185168183048111821515161561421257614212613ef9565b02949350505050565b60008261422a5761422a613f3f565b500490565b60008451614241818460208901613aab565b80830190507f2e00000000000000000000000000000000000000000000000000000000000000808252855161427d816001850160208a01613aab565b60019201918201528351614298816002840160208801613aab565b0160020195945050505050565b6000604082840312156142b757600080fd5b6040516040810181811067ffffffffffffffff821117156142da576142da613b38565b604052825181526020928301519281019290925250919050565b60006080828403121561430657600080fd5b6040516080810181811067ffffffffffffffff8211171561432957614329613b38565b8060405250823581526020830135602082015260408301356040820152606083013560608201528091505092915050565b600067ffffffffffffffff8084111561437557614375613b38565b8360051b6020614386818301613b90565b86815291850191818101903684111561439e57600080fd5b865b848110156143d2578035868111156143b85760008081fd5b6143c436828b01613c08565b8452509183019183016143a0565b50979650505050505050565b600082198211156143f1576143f1613ef9565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361442757614427613ef9565b5060010190565b60008261443d5761443d613f3f565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a08301526144bc60c0830184613adb565b98975050505050505050565b600060ff8316806144db576144db613f3f565b8060ff84160691505092915050565b600060ff821660ff84168082101561450457614504613ef9565b90039392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(OptimismPortalStorageLayoutJSON), OptimismPortalStorageLayout); err != nil {
		panic(err)
	}

	layouts["OptimismPortal"] = OptimismPortalStorageLayout
	deployedBytecodes["OptimismPortal"] = OptimismPortalDeployedBin
}
