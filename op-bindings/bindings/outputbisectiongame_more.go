// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const OutputBisectionGameStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"createdAt\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_userDefinedValueType(Timestamp)1017\"},{\"astId\":1001,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"resolvedAt\",\"offset\":8,\"slot\":\"0\",\"type\":\"t_userDefinedValueType(Timestamp)1017\"},{\"astId\":1002,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"status\",\"offset\":16,\"slot\":\"0\",\"type\":\"t_enum(GameStatus)1010\"},{\"astId\":1003,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"bondManager\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_contract(IBondManager)1009\"},{\"astId\":1004,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"settlementHead\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_userDefinedValueType(Hash)1015\"},{\"astId\":1005,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"claimData\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_struct(ClaimData)1011_storage)dyn_storage\"},{\"astId\":1006,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"claims\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_mapping(t_userDefinedValueType(ClaimHash)1013,t_bool)\"},{\"astId\":1007,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"subgames\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_mapping(t_uint256,t_array(t_uint256)dyn_storage)\"},{\"astId\":1008,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"subgameAtRootResolved\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_bool\"}],\"types\":{\"t_array(t_struct(ClaimData)1011_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct IOutputBisectionGame.ClaimData[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(ClaimData)1011_storage\"},\"t_array(t_uint256)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"uint256[]\",\"numberOfBytes\":\"32\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(IBondManager)1009\":{\"encoding\":\"inplace\",\"label\":\"contract IBondManager\",\"numberOfBytes\":\"20\"},\"t_enum(GameStatus)1010\":{\"encoding\":\"inplace\",\"label\":\"enum GameStatus\",\"numberOfBytes\":\"1\"},\"t_mapping(t_uint256,t_array(t_uint256)dyn_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint256[])\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_array(t_uint256)dyn_storage\"},\"t_mapping(t_userDefinedValueType(ClaimHash)1013,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(ClaimHash =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_userDefinedValueType(ClaimHash)1013\",\"value\":\"t_bool\"},\"t_struct(ClaimData)1011_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IOutputBisectionGame.ClaimData\",\"numberOfBytes\":\"96\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_userDefinedValueType(Claim)1012\":{\"encoding\":\"inplace\",\"label\":\"Claim\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(ClaimHash)1013\":{\"encoding\":\"inplace\",\"label\":\"ClaimHash\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(Clock)1014\":{\"encoding\":\"inplace\",\"label\":\"Clock\",\"numberOfBytes\":\"16\"},\"t_userDefinedValueType(Hash)1015\":{\"encoding\":\"inplace\",\"label\":\"Hash\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(Position)1016\":{\"encoding\":\"inplace\",\"label\":\"Position\",\"numberOfBytes\":\"16\"},\"t_userDefinedValueType(Timestamp)1017\":{\"encoding\":\"inplace\",\"label\":\"Timestamp\",\"numberOfBytes\":\"8\"}}}"

var OutputBisectionGameStorageLayout = new(solc.StorageLayout)

var OutputBisectionGameDeployedBin = "0x6080604052600436106101b75760003560e01c80638b85902b116100ec578063c6f0308c1161008a578063d8cc1a3c11610064578063d8cc1a3c14610649578063f8f43ff614610669578063fa24f74314610689578063fdffbb28146106ad57600080fd5b8063c6f0308c14610590578063caa4ba2d146105f4578063cf09e0d01461062857600080fd5b8063bbdc02db116100c6578063bbdc02db146104ce578063bcef3b551461050c578063c31b29ce14610549578063c55cd0c71461057d57600080fd5b80638b85902b146104445780639293129814610484578063a85c3381146104b857600080fd5b80634778efe811610159578063632247ea11610133578063632247ea146103d35780636737abeb146103e65780638129fc1c1461041a5780638980e0cc1461042f57600080fd5b80634778efe81461033457806354fd4d5014610368578063609d3334146103be57600080fd5b8063266198f911610195578063266198f9146102845780632810e1d6146102b857806335fef567146102cd578063363cc427146102e257600080fd5b806319effeb4146101bc578063200d2ed21461020757806324185bc614610242575b600080fd5b3480156101c857600080fd5b506000546101e99068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b34801561021357600080fd5b5060005461023590700100000000000000000000000000000000900460ff1681565b6040516101fe9190612734565b34801561024e57600080fd5b506102767f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016101fe565b34801561029057600080fd5b506102767f000000000000000000000000000000000000000000000000000000000000000081565b3480156102c457600080fd5b506102356106c0565b6102e06102db366004612775565b610891565b005b3480156102ee57600080fd5b5060015461030f9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101fe565b34801561034057600080fd5b506102767f000000000000000000000000000000000000000000000000000000000000000081565b34801561037457600080fd5b506103b16040518060400160405280600681526020017f302e302e3133000000000000000000000000000000000000000000000000000081525081565b6040516101fe9190612802565b3480156103ca57600080fd5b506103b16108a1565b6102e06103e1366004612831565b6108b3565b3480156103f257600080fd5b506102767f000000000000000000000000000000000000000000000000000000000000000081565b34801561042657600080fd5b506102e0610fab565b34801561043b57600080fd5b50600354610276565b34801561045057600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900360200135610276565b34801561049057600080fd5b5061030f7f000000000000000000000000000000000000000000000000000000000000000081565b3480156104c457600080fd5b5061027660025481565b3480156104da57600080fd5b5060405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101fe565b34801561051857600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900335610276565b34801561055557600080fd5b506101e97f000000000000000000000000000000000000000000000000000000000000000081565b6102e061058b366004612775565b6110fc565b34801561059c57600080fd5b506105b06105ab366004612866565b611108565b6040805163ffffffff90961686529315156020860152928401919091526fffffffffffffffffffffffffffffffff908116606084015216608082015260a0016101fe565b34801561060057600080fd5b506102767f000000000000000000000000000000000000000000000000000000000000000081565b34801561063457600080fd5b506000546101e99067ffffffffffffffff1681565b34801561065557600080fd5b506102e06106643660046128c8565b611179565b34801561067557600080fd5b506102e0610684366004612952565b6116ba565b34801561069557600080fd5b5061069e611b98565b6040516101fe9392919061297e565b6102e06106bb366004612866565b611bf5565b600080600054700100000000000000000000000000000000900460ff1660028111156106ee576106ee612705565b14610725576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60065460ff16610761576040517f9a07664600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003600081548110610775576107756129a9565b6000918252602090912060039091020154640100000000900460ff1661079c57600261079f565b60015b6000805467ffffffffffffffff421668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff82168117835592935083927fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffff000000000000000000ffffffffffffffff9091161770010000000000000000000000000000000083600281111561085057610850612705565b02179055600281111561086557610865612705565b6040517f5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da6090600090a290565b61089d828260006108b3565b5050565b60606108ae602080611f2f565b905090565b60008054700100000000000000000000000000000000900460ff1660028111156108df576108df612705565b14610916576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82158015610922575080155b15610959576040517fa42637bc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006003848154811061096e5761096e6129a9565b600091825260208083206040805160a081018252600394909402909101805463ffffffff808216865264010000000090910460ff16151593850193909352600181015491840191909152600201546fffffffffffffffffffffffffffffffff80821660608501819052700100000000000000000000000000000000909204166080840152919350610a0291908590611fc616565b90507f0000000000000000000000000000000000000000000000000000000000000000610ac1826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff161115610b03576040517f56f57b2b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b2e7f00000000000000000000000000000000000000000000000000000000000000006001612a07565b610bca826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff1603610be257610be284611fce565b815160009063ffffffff90811614610c42576003836000015163ffffffff1681548110610c1157610c116129a9565b906000526020600020906003020160020160109054906101000a90046fffffffffffffffffffffffffffffffff1690505b608083015160009067ffffffffffffffff1667ffffffffffffffff1642610c7b846fffffffffffffffffffffffffffffffff1660401c90565b67ffffffffffffffff16610c8f9190612a07565b610c999190612a1f565b9050677fffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000060011c1667ffffffffffffffff82161115610d0c576040517f3381d11400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000604082901b42176000888152608086901b6fffffffffffffffffffffffffffffffff8b1617602052604081209192509060008181526004602052604090205490915060ff1615610d8a576040517f80497e3b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081815260046020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155815160a08101835263ffffffff808f1682529381018581529281018d81526fffffffffffffffffffffffffffffffff808c1660608401908152898216608085019081526003805480880182559981905294519885027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b8101805498511515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009099169a909916999099179690961790965590517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c8701559351925184167001000000000000000000000000000000000292909316919091177fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85d9093019290925580548b908110610f0057610f006129a9565b600091825260208083206003928302018054941515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff909516949094179093558b82526005909252604090209054610f6190600190612a1f565b8154600181018355600092835260208320015560405133918a918c917f9b3245740ec3b155098a55be84957a4da13eaf7f14a8bc6f53126c0b9350f2be91a4505050505050505050565b600080547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000164267ffffffffffffffff161781556040805160a08101825263ffffffff815260208101929092526003919081016110307ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe369081013560f01c90033590565b815260016020820152604001426fffffffffffffffffffffffffffffffff908116909152825460018181018555600094855260209485902084516003909302018054958501511515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000090961663ffffffff90931692909217949094178155604083015181850155606083015160809093015182167001000000000000000000000000000000000292909116919091176002909101556110f69043612a1f565b40600255565b61089d828260016108b3565b6003818154811061111857600080fd5b600091825260209091206003909102018054600182015460029092015463ffffffff8216935064010000000090910460ff1691906fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041685565b60008054700100000000000000000000000000000000900460ff1660028111156111a5576111a5612705565b146111dc576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600387815481106111f1576111f16129a9565b6000918252602082206003919091020160028101549092506fffffffffffffffffffffffffffffffff16908715821760011b90506112507f00000000000000000000000000000000000000000000000000000000000000006001612a07565b6112ec826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff161461132d576040517f5f53dd9800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008089156113b057611351836fffffffffffffffffffffffffffffffff16612022565b67ffffffffffffffff16156113845761137b61136e600186612a36565b865463ffffffff166120c8565b600101546113a6565b7f00000000000000000000000000000000000000000000000000000000000000005b91508490506113ca565b846001015491506113c784600161136e9190612a67565b90505b600882901b60088a8a6040516113e1929190612a9b565b6040518091039020901b14611422576040517f696550ff00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061142d8c612165565b9050600082600101547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e14ced328d8d8d8d886040518663ffffffff1660e01b8152600401611497959493929190612af4565b6020604051808303816000875af11580156114b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114da9190612b2e565b600285810154929091149250600091611585906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b611621896fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b61162b9190612b47565b6116359190612b68565b67ffffffffffffffff16159050811515810361167d576040517ffb4e40dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505085547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff16640100000000179095555050505050505050505050565b60008054700100000000000000000000000000000000900460ff1660028111156116e6576116e6612705565b1461171d576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008060008061172c86612194565b935093509350935060006117428585858561255f565b905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16637dc0d1d06040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117d59190612bb6565b905088600103611897576002546040517f52f0f3ad000000000000000000000000000000000000000000000000000000008152600481018b9052602481018490526044810191909152602060648201526084810188905273ffffffffffffffffffffffffffffffffffffffff8216906352f0f3ad9060a4015b6020604051808303816000875af115801561186d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118919190612b2e565b50611b8d565b8860020361190e576040517f52f0f3ad000000000000000000000000000000000000000000000000000000008152600481018a90526024810183905260448101879052602060648201526084810188905273ffffffffffffffffffffffffffffffffffffffff8216906352f0f3ad9060a40161184e565b88600303611985576040517f52f0f3ad000000000000000000000000000000000000000000000000000000008152600481018a90526024810183905260448101859052602060648201526084810188905273ffffffffffffffffffffffffffffffffffffffff8216906352f0f3ad9060a40161184e565b88600403611ae25760006fffffffffffffffffffffffffffffffff861615611a06576119c2866fffffffffffffffffffffffffffffffff16612022565b6119f69067ffffffffffffffff167f0000000000000000000000000000000000000000000000000000000000000000612a07565b611a01906001612a07565b611a28565b7f00000000000000000000000000000000000000000000000000000000000000005b6040517f52f0f3ad000000000000000000000000000000000000000000000000000000008152600481018c90526024810185905260c082901b604482015260086064820152608481018a905290915073ffffffffffffffffffffffffffffffffffffffff8316906352f0f3ad9060a4016020604051808303816000875af1158015611ab7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611adb9190612b2e565b5050611b8d565b88600503611b5b576040517f52f0f3ad000000000000000000000000000000000000000000000000000000008152600481018a9052602481018390524660c01b6044820152600860648201526084810188905273ffffffffffffffffffffffffffffffffffffffff8216906352f0f3ad9060a40161184e565b6040517fff137e6500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050505050505050565b7f0000000000000000000000000000000000000000000000000000000000000000367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003356060611bee6108a1565b9050909192565b60008054700100000000000000000000000000000000900460ff166002811115611c2157611c21612705565b14611c58576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060038281548110611c6d57611c6d6129a9565b60009182526020909120600260039092020190810154909150677fffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000060011c1690611cdd90700100000000000000000000000000000000900467ffffffffffffffff1642612a1f565b6002830154611d0d9190700100000000000000000000000000000000900460401c67ffffffffffffffff16612a07565b11611d44576040517ff2440b5300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082815260056020526040902082158015611d62575060065460ff165b15611d99576040517ff1a9458100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8054158015611da757508215155b15611dde576040517ff1a9458100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000805b8254811015611eac576000838281548110611dff57611dff6129a9565b6000918252602080832090910154808352600590915260409091205490915015611e55576040517f9a07664600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060038281548110611e6a57611e6a6129a9565b600091825260209091206003909102018054909150640100000000900460ff16611e9957600193505050611eac565b505080611ea590612bec565b9050611de2565b5082547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff16640100000000821515021783556000848152600560205260408120611ef5916126cb565b83600003611f2957600680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555b50505050565b60606000611f6684367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003612a07565b90508267ffffffffffffffff1667ffffffffffffffff811115611f8b57611f8b612c24565b6040519080825280601f01601f191660200182016040528015611fb5576020820181803683370190505b509150828160208401375092915050565b151760011b90565b600081901a6001811480611fe5575060ff81166002145b61089d576040517ff40239db0000000000000000000000000000000000000000000000000000000081526004810183905260240160405180910390fd5b6000806120af837e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b600167ffffffffffffffff919091161b90920392915050565b6000806120e6846fffffffffffffffffffffffffffffffff1661261f565b9050600383815481106120fb576120fb6129a9565b906000526020600020906003020191505b60028201546fffffffffffffffffffffffffffffffff82811691161461215e57815460038054909163ffffffff16908110612149576121496129a9565b9060005260206000209060030201915061210c565b5092915050565b600080600080600061217686612194565b935093509350935061218a8484848461255f565b9695505050505050565b60008060008060008590506000600382815481106121b4576121b46129a9565b600091825260209091206002600390920201908101549091507f00000000000000000000000000000000000000000000000000000000000000009061228b906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff16116122cc576040517fb34b5c2200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000815b60028301547f000000000000000000000000000000000000000000000000000000000000000090612393906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff16925082111561240f57825463ffffffff166123d97f00000000000000000000000000000000000000000000000000000000000000006001612a07565b83036123e3578391505b600381815481106123f6576123f66129a9565b90600052602060002090600302019350809450506122d0565b600280820154908401546fffffffffffffffffffffffffffffffff918216911660008161243c8460011c90565b6fffffffffffffffffffffffffffffffff16149050801561250d576000612474836fffffffffffffffffffffffffffffffff16612022565b67ffffffffffffffff1611156124c357600061249a612494600185612a36565b896120c8565b6001810154600290910154909c506fffffffffffffffffffffffffffffffff169a506124e79050565b7f00000000000000000000000000000000000000000000000000000000000000009a505b600186015460028701549099506fffffffffffffffffffffffffffffffff169750612551565b600061251d612494846001612a67565b6001808901546002808b015492840154930154909e506fffffffffffffffffffffffffffffffff9182169d50919b50169850505b505050505050509193509193565b6000836fffffffffffffffffffffffffffffffff166000036125c65782826040516020016125a99291909182526fffffffffffffffffffffffffffffffff16602082015260400190565b604051602081830303815290604052805190602001209050612617565b60408051602081018790526fffffffffffffffffffffffffffffffff8087169282019290925260608101859052908316608082015260a0016040516020818303038152906040528051906020012090505b949350505050565b600081196001830116816126b3827e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169390931c8015179392505050565b50805460008255906000526020600020908101906126e991906126ec565b50565b5b8082111561270157600081556001016126ed565b5090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b602081016003831061276f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b6000806040838503121561278857600080fd5b50508035926020909101359150565b6000815180845260005b818110156127bd576020818501810151868301820152016127a1565b818111156127cf576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006128156020830184612797565b9392505050565b8035801515811461282c57600080fd5b919050565b60008060006060848603121561284657600080fd5b833592506020840135915061285d6040850161281c565b90509250925092565b60006020828403121561287857600080fd5b5035919050565b60008083601f84011261289157600080fd5b50813567ffffffffffffffff8111156128a957600080fd5b6020830191508360208285010111156128c157600080fd5b9250929050565b600080600080600080608087890312156128e157600080fd5b863595506128f16020880161281c565b9450604087013567ffffffffffffffff8082111561290e57600080fd5b61291a8a838b0161287f565b9096509450606089013591508082111561293357600080fd5b5061294089828a0161287f565b979a9699509497509295939492505050565b60008060006060848603121561296757600080fd5b505081359360208301359350604090920135919050565b60ff841681528260208201526060604082015260006129a06060830184612797565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115612a1a57612a1a6129d8565b500190565b600082821015612a3157612a316129d8565b500390565b60006fffffffffffffffffffffffffffffffff83811690831681811015612a5f57612a5f6129d8565b039392505050565b60006fffffffffffffffffffffffffffffffff808316818516808303821115612a9257612a926129d8565b01949350505050565b8183823760009101908152919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b606081526000612b08606083018789612aab565b8281036020840152612b1b818688612aab565b9150508260408301529695505050505050565b600060208284031215612b4057600080fd5b5051919050565b600067ffffffffffffffff83811690831681811015612a5f57612a5f6129d8565b600067ffffffffffffffff80841680612baa577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910692915050565b600060208284031215612bc857600080fd5b815173ffffffffffffffffffffffffffffffffffffffff8116811461281557600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612c1d57612c1d6129d8565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(OutputBisectionGameStorageLayoutJSON), OutputBisectionGameStorageLayout); err != nil {
		panic(err)
	}

	layouts["OutputBisectionGame"] = OutputBisectionGameStorageLayout
	deployedBytecodes["OutputBisectionGame"] = OutputBisectionGameDeployedBin
	immutableReferences["OutputBisectionGame"] = true
}
