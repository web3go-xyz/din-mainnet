// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const FaultDisputeGameStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"createdAt\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_userDefinedValueType(Timestamp)1017\"},{\"astId\":1001,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"status\",\"offset\":8,\"slot\":\"0\",\"type\":\"t_enum(GameStatus)1008\"},{\"astId\":1002,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"bondManager\",\"offset\":9,\"slot\":\"0\",\"type\":\"t_contract(IBondManager)1007\"},{\"astId\":1003,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"l1Head\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_userDefinedValueType(Hash)1015\"},{\"astId\":1004,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"claimData\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_struct(ClaimData)1009_storage)dyn_storage\"},{\"astId\":1005,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"proposals\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_struct(OutputProposals)1011_storage\"},{\"astId\":1006,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"claims\",\"offset\":0,\"slot\":\"7\",\"type\":\"t_mapping(t_userDefinedValueType(ClaimHash)1013,t_bool)\"}],\"types\":{\"t_array(t_struct(ClaimData)1009_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct IFaultDisputeGame.ClaimData[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(ClaimData)1009_storage\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(IBondManager)1007\":{\"encoding\":\"inplace\",\"label\":\"contract IBondManager\",\"numberOfBytes\":\"20\"},\"t_enum(GameStatus)1008\":{\"encoding\":\"inplace\",\"label\":\"enum GameStatus\",\"numberOfBytes\":\"1\"},\"t_mapping(t_userDefinedValueType(ClaimHash)1013,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(ClaimHash =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_userDefinedValueType(ClaimHash)1013\",\"value\":\"t_bool\"},\"t_struct(ClaimData)1009_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IFaultDisputeGame.ClaimData\",\"numberOfBytes\":\"96\"},\"t_struct(OutputProposal)1010_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IFaultDisputeGame.OutputProposal\",\"numberOfBytes\":\"64\"},\"t_struct(OutputProposals)1011_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IFaultDisputeGame.OutputProposals\",\"numberOfBytes\":\"128\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_userDefinedValueType(Claim)1012\":{\"encoding\":\"inplace\",\"label\":\"Claim\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(ClaimHash)1013\":{\"encoding\":\"inplace\",\"label\":\"ClaimHash\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(Clock)1014\":{\"encoding\":\"inplace\",\"label\":\"Clock\",\"numberOfBytes\":\"16\"},\"t_userDefinedValueType(Hash)1015\":{\"encoding\":\"inplace\",\"label\":\"Hash\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(Position)1016\":{\"encoding\":\"inplace\",\"label\":\"Position\",\"numberOfBytes\":\"16\"},\"t_userDefinedValueType(Timestamp)1017\":{\"encoding\":\"inplace\",\"label\":\"Timestamp\",\"numberOfBytes\":\"8\"}}}"

var FaultDisputeGameStorageLayout = new(solc.StorageLayout)

var FaultDisputeGameDeployedBin = "0x6080604052600436106101ac5760003560e01c80636361506d116100ec578063c0c3a0921161008a578063c6f0308c11610064578063c6f0308c1461061d578063cf09e0d014610681578063d8cc1a3c146106a2578063fa24f743146106c257600080fd5b8063c0c3a09214610589578063c31b29ce146105bd578063c55cd0c71461060a57600080fd5b80638b85902b116100c65780638b85902b1461049a57806392931298146104da578063bbdc02db1461050e578063bcef3b551461054c57600080fd5b80636361506d1461045a5780638129fc1c146104705780638980e0cc1461048557600080fd5b8063363cc4271161015957806354fd4d501161013357806354fd4d501461038057806355ef20e6146103a2578063609d333414610432578063632247ea1461044757600080fd5b8063363cc427146102b95780634778efe814610318578063529184c91461034c57600080fd5b80632810e1d61161018a5780632810e1d614610251578063298c90051461026657806335fef567146102a657600080fd5b80631e27052a146101b1578063200d2ed2146101d3578063266198f91461020f575b600080fd5b3480156101bd57600080fd5b506101d16101cc366004612417565b6106e6565b005b3480156101df57600080fd5b506000546101f99068010000000000000000900460ff1681565b6040516102069190612468565b60405180910390f35b34801561021b57600080fd5b506102437f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610206565b34801561025d57600080fd5b506101f96108a5565b34801561027257600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900360400135610243565b6101d16102b4366004612417565b610ccb565b3480156102c557600080fd5b506000546102f3906901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610206565b34801561032457600080fd5b506102437f000000000000000000000000000000000000000000000000000000000000000081565b34801561035857600080fd5b506102f37f000000000000000000000000000000000000000000000000000000000000000081565b34801561038c57600080fd5b50610395610cdb565b604051610206919061251f565b3480156103ae57600080fd5b5060408051606080820183526003546fffffffffffffffffffffffffffffffff808216845270010000000000000000000000000000000091829004811660208086019190915260045485870152855193840186526005548083168552929092041690820152600654928101929092526104249182565b604051610206929190612539565b34801561043e57600080fd5b50610395610d7e565b6101d16104553660046125a2565b610d8c565b34801561046657600080fd5b5061024360015481565b34801561047c57600080fd5b506101d1611360565b34801561049157600080fd5b50600254610243565b3480156104a657600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900360200135610243565b3480156104e657600080fd5b506102f37f000000000000000000000000000000000000000000000000000000000000000081565b34801561051a57600080fd5b5060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610206565b34801561055857600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900335610243565b34801561059557600080fd5b506102f37f000000000000000000000000000000000000000000000000000000000000000081565b3480156105c957600080fd5b506105f17f000000000000000000000000000000000000000000000000000000000000000081565b60405167ffffffffffffffff9091168152602001610206565b6101d1610618366004612417565b611954565b34801561062957600080fd5b5061063d6106383660046125d7565b611960565b6040805163ffffffff90961686529315156020860152928401919091526fffffffffffffffffffffffffffffffff908116606084015216608082015260a001610206565b34801561068d57600080fd5b506000546105f19067ffffffffffffffff1681565b3480156106ae57600080fd5b506101d16106bd366004612639565b6119d1565b3480156106ce57600080fd5b506106d7611f3a565b604051610206939291906126c3565b6000805468010000000000000000900460ff16600281111561070a5761070a612439565b14610741576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16637dc0d1d06040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d291906126ee565b7f9a1f5e7f00000000000000000000000000000000000000000000000000000000601c8190526020859052909150600084600181146108395760028114610843576003811461084d576004811461085757600581146108675763ff137e656000526004601cfd5b600154915061086e565b600454915061086e565b600654915061086e565b60035460801c60c01b915061086e565b4660c01b91505b50604052600160038511811b6005031b60605260808390526000806084601c82865af161089f573d6000803e3d6000fd5b50505050565b60008060005468010000000000000000900460ff1660028111156108cb576108cb612439565b14610902576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025460009061091490600190612753565b90506fffffffffffffffffffffffffffffffff815b67ffffffffffffffff8110156109fe5760006002828154811061094e5761094e61276a565b6000918252602090912060039091020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9093019290915060ff640100000000909104161561099f5750610929565b60028101546000906109e3906fffffffffffffffffffffffffffffffff167f0000000000000000000000000000000000000000000000000000000000000000611f97565b9050838110156109f7578093508260010194505b5050610929565b50600060028381548110610a1457610a1461276a565b600091825260208220600390910201805490925063ffffffff90811691908214610a7e5760028281548110610a4b57610a4b61276a565b906000526020600020906003020160020160109054906101000a90046fffffffffffffffffffffffffffffffff16610aaa565b600283015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff165b9050677fffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000060011c16610aee67ffffffffffffffff831642612753565b610b0a836fffffffffffffffffffffffffffffffff1660401c90565b67ffffffffffffffff16610b1e9190612799565b11610b55576040517ff2440b5300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600283810154610bf7906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b610c0191906127e0565b67ffffffffffffffff16158015610c2857506fffffffffffffffffffffffffffffffff8414155b15610c365760029550610c3b565b600195505b600080548791907fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff1668010000000000000000836002811115610c8057610c80612439565b021790556002811115610c9557610c95612439565b6040517f5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da6090600090a2505050505090565b905090565b610cd782826000610d8c565b5050565b6060610d067f000000000000000000000000000000000000000000000000000000000000000061204c565b610d2f7f000000000000000000000000000000000000000000000000000000000000000061204c565b610d587f000000000000000000000000000000000000000000000000000000000000000061204c565b604051602001610d6a93929190612807565b604051602081830303815290604052905090565b6060610cc660206040612189565b6000805468010000000000000000900460ff166002811115610db057610db0612439565b14610de7576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82158015610df3575080155b15610e2a576040517fa42637bc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060028481548110610e3f57610e3f61276a565b600091825260208083206040805160a081018252600394909402909101805463ffffffff808216865264010000000090910460ff16151593850193909352600181015491840191909152600201546fffffffffffffffffffffffffffffffff80821660608501819052700100000000000000000000000000000000909204166080840152919350610ed39190859061222016565b90507f0000000000000000000000000000000000000000000000000000000000000000610f92826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff161115610fd4576040517f56f57b2b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160009063ffffffff90811614611034576002836000015163ffffffff16815481106110035761100361276a565b906000526020600020906003020160020160109054906101000a90046fffffffffffffffffffffffffffffffff1690505b608083015160009067ffffffffffffffff1667ffffffffffffffff164261106d846fffffffffffffffffffffffffffffffff1660401c90565b67ffffffffffffffff166110819190612799565b61108b9190612753565b9050677fffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000060011c1667ffffffffffffffff821611156110fe576040517f3381d11400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000604082901b42179050600061111f888660009182526020526040902090565b60008181526007602052604090205490915060ff161561116b576040517f80497e3b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081815260076020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155815160a08101835263ffffffff808f1682529381018581529281018d81526fffffffffffffffffffffffffffffffff808c16606084019081528982166080850190815260028054808801825599819052945160039099027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace8101805498511515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009099169a909916999099179690961790965590517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf8701559351925184167001000000000000000000000000000000000292909316919091177f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad09093019290925580548b9081106112e3576112e361276a565b6000918252602082206003909102018054921515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff9093169290921790915560405133918a918c917f9b3245740ec3b155098a55be84957a4da13eaf7f14a8bc6f53126c0b9350f2be91a4505050505050505050565b367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90033560001a6001146113f7576040517ff40239db000000000000000000000000000000000000000000000000000000008152367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900335600482015260240160405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000164267ffffffffffffffff161781556040805160a08101825263ffffffff8152602081019290925260029190810161147c7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe369081013560f01c90033590565b815260016020820152604001426fffffffffffffffffffffffffffffffff9081169091528254600181810185556000948552602080862085516003909402018054918601511515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000090921663ffffffff909416939093171782556040840151908201556060830151608090930151821670010000000000000000000000000000000002929091169190911760029091015573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016637f0064206115a560207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe369081013560f01c9003013590565b6040518263ffffffff1660e01b81526004016115c391815260200190565b602060405180830381865afa1580156115e0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611604919061287d565b9050600073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663a25ae55761164f600185612753565b6040518263ffffffff1660e01b815260040161166d91815260200190565b606060405180830381865afa15801561168a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116ae91906128e5565b6040517fa25ae5570000000000000000000000000000000000000000000000000000000081526004810184905290915060009073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063a25ae55790602401606060405180830381865afa15801561173f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061176391906128e5565b9050600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166399d548aa367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003604001356040518263ffffffff1660e01b81526004016117ef91815260200190565b6040805180830381865afa15801561180b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061182f9190612971565b905081602001516fffffffffffffffffffffffffffffffff16816020015167ffffffffffffffff161161188e576040517f13809ba500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160a0810182529081908101806118a9600189612753565b6fffffffffffffffffffffffffffffffff9081168252604088810151821660208085019190915298519281019290925291835280516060810182529782168852858101518216888801529451878601529085019590955280518051818601519087167001000000000000000000000000000000009188168202176003559084015160045590840151805194810151948616949095160292909217600555919091015160065551600155565b610cd782826001610d8c565b6002818154811061197057600080fd5b600091825260209091206003909102018054600182015460029092015463ffffffff8216935064010000000090910460ff1691906fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041685565b6000805468010000000000000000900460ff1660028111156119f5576119f5612439565b14611a2c576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060028781548110611a4157611a4161276a565b6000918252602082206003919091020160028101549092506fffffffffffffffffffffffffffffffff16908715821760011b9050611aa07f00000000000000000000000000000000000000000000000000000000000000006001612799565b611b3c826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff1614611b7d576040517f5f53dd9800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000808915611c0057611ba1836fffffffffffffffffffffffffffffffff16612228565b67ffffffffffffffff1615611bd457611bcb611bbe6001866129f8565b865463ffffffff166122ce565b60010154611bf6565b7f00000000000000000000000000000000000000000000000000000000000000005b9150849050611c1a565b84600101549150611c17846001611bbe9190612a29565b90505b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8216611c7b8a8a604051611c4f929190612a5d565b60405180910390207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b14611cb2576040517f696550ff00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081600101547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f8e0cb968c8c8c8c6040518563ffffffff1660e01b8152600401611d189493929190612ab6565b6020604051808303816000875af1158015611d37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d5b919061287d565b600284810154929091149250600091611e06906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b611ea2886fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b611eac9190612ae8565b611eb691906127e0565b67ffffffffffffffff161590508115158103611efe576040517ffb4e40dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505084547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff166401000000001790945550505050505050505050565b7f0000000000000000000000000000000000000000000000000000000000000000367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003356060611f90610d7e565b9050909192565b600080612024847e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff1690508083036001841b600180831b0386831b17039250505092915050565b60608160000361208f57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156120b957806120a381612b09565b91506120b29050600a83612b41565b9150612093565b60008167ffffffffffffffff8111156120d4576120d4612896565b6040519080825280601f01601f1916602001820160405280156120fe576020820181803683370190505b5090505b841561218157612113600183612753565b9150612120600a86612b55565b61212b906030612799565b60f81b8183815181106121405761214061276a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535061217a600a86612b41565b9450612102565b949350505050565b606060006121c084367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003612799565b90508267ffffffffffffffff1667ffffffffffffffff8111156121e5576121e5612896565b6040519080825280601f01601f19166020018201604052801561220f576020820181803683370190505b509150828160208401375092915050565b151760011b90565b6000806122b5837e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b600167ffffffffffffffff919091161b90920392915050565b6000806122ec846fffffffffffffffffffffffffffffffff1661236b565b9050600283815481106123015761230161276a565b906000526020600020906003020191505b60028201546fffffffffffffffffffffffffffffffff82811691161461236457815460028054909163ffffffff1690811061234f5761234f61276a565b90600052602060002090600302019150612312565b5092915050565b600081196001830116816123ff827e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169390931c8015179392505050565b6000806040838503121561242a57600080fd5b50508035926020909101359150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60208101600383106124a3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b838110156124c45781810151838201526020016124ac565b8381111561089f5750506000910152565b600081518084526124ed8160208601602086016124a9565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061253260208301846124d5565b9392505050565b82516fffffffffffffffffffffffffffffffff90811682526020808501518216818401526040808601518185015284518316606085015290840151909116608083015282015160a082015260c08101612532565b8035801515811461259d57600080fd5b919050565b6000806000606084860312156125b757600080fd5b83359250602084013591506125ce6040850161258d565b90509250925092565b6000602082840312156125e957600080fd5b5035919050565b60008083601f84011261260257600080fd5b50813567ffffffffffffffff81111561261a57600080fd5b60208301915083602082850101111561263257600080fd5b9250929050565b6000806000806000806080878903121561265257600080fd5b863595506126626020880161258d565b9450604087013567ffffffffffffffff8082111561267f57600080fd5b61268b8a838b016125f0565b909650945060608901359150808211156126a457600080fd5b506126b189828a016125f0565b979a9699509497509295939492505050565b60ff841681528260208201526060604082015260006126e560608301846124d5565b95945050505050565b60006020828403121561270057600080fd5b815173ffffffffffffffffffffffffffffffffffffffff8116811461253257600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561276557612765612724565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082198211156127ac576127ac612724565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff808416806127fb576127fb6127b1565b92169190910692915050565b600084516128198184602089016124a9565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551612855816001850160208a016124a9565b600192019182015283516128708160028401602088016124a9565b0160020195945050505050565b60006020828403121561288f57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b80516fffffffffffffffffffffffffffffffff8116811461259d57600080fd5b6000606082840312156128f757600080fd5b6040516060810181811067ffffffffffffffff82111715612941577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405282518152612954602084016128c5565b6020820152612965604084016128c5565b60408201529392505050565b60006040828403121561298357600080fd5b6040516040810167ffffffffffffffff82821081831117156129ce577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b81604052845183526020850151915080821682146129eb57600080fd5b5060208201529392505050565b60006fffffffffffffffffffffffffffffffff83811690831681811015612a2157612a21612724565b039392505050565b60006fffffffffffffffffffffffffffffffff808316818516808303821115612a5457612a54612724565b01949350505050565b8183823760009101908152919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000612aca604083018688612a6d565b8281036020840152612add818587612a6d565b979650505050505050565b600067ffffffffffffffff83811690831681811015612a2157612a21612724565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612b3a57612b3a612724565b5060010190565b600082612b5057612b506127b1565b500490565b600082612b6457612b646127b1565b50069056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(FaultDisputeGameStorageLayoutJSON), FaultDisputeGameStorageLayout); err != nil {
		panic(err)
	}

	layouts["FaultDisputeGame"] = FaultDisputeGameStorageLayout
	deployedBytecodes["FaultDisputeGame"] = FaultDisputeGameDeployedBin
}
