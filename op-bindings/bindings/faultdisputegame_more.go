// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const FaultDisputeGameStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"createdAt\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_userDefinedValueType(Timestamp)1017\"},{\"astId\":1001,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"resolvedAt\",\"offset\":8,\"slot\":\"0\",\"type\":\"t_userDefinedValueType(Timestamp)1017\"},{\"astId\":1002,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"status\",\"offset\":16,\"slot\":\"0\",\"type\":\"t_enum(GameStatus)1010\"},{\"astId\":1003,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"l1Head\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_userDefinedValueType(Hash)1015\"},{\"astId\":1004,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"claimData\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_struct(ClaimData)1011_storage)dyn_storage\"},{\"astId\":1005,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"credit\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1006,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"claims\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_mapping(t_userDefinedValueType(ClaimHash)1013,t_bool)\"},{\"astId\":1007,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"subgames\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_mapping(t_uint256,t_array(t_uint256)dyn_storage)\"},{\"astId\":1008,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"subgameAtRootResolved\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_bool\"},{\"astId\":1009,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"initialized\",\"offset\":1,\"slot\":\"6\",\"type\":\"t_bool\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_struct(ClaimData)1011_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct IFaultDisputeGame.ClaimData[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(ClaimData)1011_storage\"},\"t_array(t_uint256)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"uint256[]\",\"numberOfBytes\":\"32\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_enum(GameStatus)1010\":{\"encoding\":\"inplace\",\"label\":\"enum GameStatus\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_array(t_uint256)dyn_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint256[])\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_array(t_uint256)dyn_storage\"},\"t_mapping(t_userDefinedValueType(ClaimHash)1013,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(ClaimHash =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_userDefinedValueType(ClaimHash)1013\",\"value\":\"t_bool\"},\"t_struct(ClaimData)1011_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IFaultDisputeGame.ClaimData\",\"numberOfBytes\":\"160\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_userDefinedValueType(Claim)1012\":{\"encoding\":\"inplace\",\"label\":\"Claim\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(ClaimHash)1013\":{\"encoding\":\"inplace\",\"label\":\"ClaimHash\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(Clock)1014\":{\"encoding\":\"inplace\",\"label\":\"Clock\",\"numberOfBytes\":\"16\"},\"t_userDefinedValueType(Hash)1015\":{\"encoding\":\"inplace\",\"label\":\"Hash\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(Position)1016\":{\"encoding\":\"inplace\",\"label\":\"Position\",\"numberOfBytes\":\"16\"},\"t_userDefinedValueType(Timestamp)1017\":{\"encoding\":\"inplace\",\"label\":\"Timestamp\",\"numberOfBytes\":\"8\"}}}"

var FaultDisputeGameStorageLayout = new(solc.StorageLayout)

var FaultDisputeGameDeployedBin = "0x6080604052600436106101cd5760003560e01c80638d450a95116100f7578063d5d44d8011610095578063f8f43ff611610064578063f8f43ff614610697578063fa24f743146106b7578063fa315aa9146106db578063fdffbb281461070e57600080fd5b8063d5d44d80146105e4578063d8cc1a3c14610611578063e1f0c37614610631578063ec5e63081461066457600080fd5b8063c395e1ca116100d1578063c395e1ca14610505578063c55cd0c714610526578063c6f0308c14610539578063cf09e0d0146105c357600080fd5b80638d450a9514610454578063bbdc02db14610487578063bcef3b55146104c857600080fd5b8063609d33341161016f57806368800abf1161013e57806368800abf146103c45780638129fc1c146103f75780638980e0cc146103ff5780638b85902b1461041457600080fd5b8063609d33341461036657806360e274641461037b578063632247ea1461039b5780636361506d146103ae57600080fd5b80632810e1d6116101ab5780632810e1d61461029557806335fef567146102aa5780633a768463146102bf57806354fd4d501461031057600080fd5b80630356fe3a146101d257806319effeb414610214578063200d2ed21461025a575b600080fd5b3480156101de57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b6040519081526020015b60405180910390f35b34801561022057600080fd5b506000546102419068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff909116815260200161020b565b34801561026657600080fd5b5060005461028890700100000000000000000000000000000000900460ff1681565b60405161020b9190613285565b3480156102a157600080fd5b50610288610721565b6102bd6102b83660046132c6565b61091e565b005b3480156102cb57600080fd5b5060405173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815260200161020b565b34801561031c57600080fd5b506103596040518060400160405280600581526020017f302e322e3000000000000000000000000000000000000000000000000000000081525081565b60405161020b9190613353565b34801561037257600080fd5b5061035961092e565b34801561038757600080fd5b506102bd610396366004613388565b610940565b6102bd6103a93660046133c1565b6109f0565b3480156103ba57600080fd5b5061020160015481565b3480156103d057600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610201565b6102bd611186565b34801561040b57600080fd5b50600254610201565b34801561042057600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900360200135610201565b34801561046057600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610201565b34801561049357600080fd5b5060405163ffffffff7f000000000000000000000000000000000000000000000000000000000000000016815260200161020b565b3480156104d457600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900335610201565b34801561051157600080fd5b506102016105203660046133f6565b50600090565b6102bd6105343660046132c6565b6114ac565b34801561054557600080fd5b50610559610554366004613428565b6114b8565b6040805163ffffffff909816885273ffffffffffffffffffffffffffffffffffffffff968716602089015295909416948601949094526fffffffffffffffffffffffffffffffff9182166060860152608085015291821660a08401521660c082015260e00161020b565b3480156105cf57600080fd5b506000546102419067ffffffffffffffff1681565b3480156105f057600080fd5b506102016105ff366004613388565b60036020526000908152604090205481565b34801561061d57600080fd5b506102bd61062c36600461348a565b61154f565b34801561063d57600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610241565b34801561067057600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610201565b3480156106a357600080fd5b506102bd6106b2366004613514565b611b2d565b3480156106c357600080fd5b506106cc611f96565b60405161020b93929190613540565b3480156106e757600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610201565b6102bd61071c366004613428565b611ff3565b600080600054700100000000000000000000000000000000900460ff16600281111561074f5761074f613256565b14610786576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60065460ff166107c2576040517f9a07664600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660026000815481106107ee576107ee61356e565b6000918252602090912060059091020154640100000000900473ffffffffffffffffffffffffffffffffffffffff161461082957600161082c565b60025b6000805467ffffffffffffffff421668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff82168117835592935083927fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffff000000000000000000ffffffffffffffff909116177001000000000000000000000000000000008360028111156108dd576108dd613256565b0217905560028111156108f2576108f2613256565b6040517f5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da6090600090a290565b61092a828260006109f0565b5050565b606061093b602080612454565b905090565b73ffffffffffffffffffffffffffffffffffffffff8116600081815260036020526040808220805490839055905190929083908381818185875af1925050503d80600081146109ab576040519150601f19603f3d011682016040523d82523d6000602084013e6109b0565b606091505b50509050806109eb576040517f83e6cc6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b60008054700100000000000000000000000000000000900460ff166002811115610a1c57610a1c613256565b14610a53576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060028481548110610a6857610a6861356e565b600091825260208083206040805160e0810182526005909402909101805463ffffffff808216865273ffffffffffffffffffffffffffffffffffffffff6401000000009092048216948601949094526001820154169184019190915260028101546fffffffffffffffffffffffffffffffff90811660608501526003820154608085015260049091015480821660a0850181905270010000000000000000000000000000000090910490911660c0840152919350909190610b2d90839086906124eb16565b90506000610bcd826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169050861580610c0f5750610c0c7f000000000000000000000000000000000000000000000000000000000000000060026135cc565b81145b8015610c19575084155b15610c50576040517fa42637bc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000811115610caa576040517f56f57b2b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610cd57f000000000000000000000000000000000000000000000000000000000000000060016135cc565b8103610ce757610ce7868885886124f3565b835160009063ffffffff90811614610d47576002856000015163ffffffff1681548110610d1657610d1661356e565b906000526020600020906005020160040160109054906101000a90046fffffffffffffffffffffffffffffffff1690505b60c0850151600090610d6b9067ffffffffffffffff165b67ffffffffffffffff1690565b67ffffffffffffffff1642610d95610d5e856fffffffffffffffffffffffffffffffff1660401c90565b67ffffffffffffffff16610da991906135cc565b610db391906135e4565b90507f000000000000000000000000000000000000000000000000000000000000000060011c677fffffffffffffff1667ffffffffffffffff82161115610e26576040517f3381d11400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000604082901b421760008a8152608087901b6fffffffffffffffffffffffffffffffff8d1617602052604081209192509060008181526004602052604090205490915060ff1615610ea4576040517f80497e3b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016004600083815260200190815260200160002060006101000a81548160ff02191690831515021790555060026040518060e001604052808d63ffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff168152602001346fffffffffffffffffffffffffffffffff1681526020018c8152602001886fffffffffffffffffffffffffffffffff168152602001846fffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003906000526020600020906005020160009091909190915060008201518160000160006101000a81548163ffffffff021916908363ffffffff16021790555060208201518160000160046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff1602179055506080820151816003015560a08201518160040160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060c08201518160040160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff1602179055505050600560008c8152602001908152602001600020600160028054905061113a91906135e4565b8154600181018355600092835260208320015560405133918c918e917f9b3245740ec3b155098a55be84957a4da13eaf7f14a8bc6f53126c0b9350f2be91a45050505050505050505050565b600654610100900460ff16156111c8576040517f0dc149f000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003602001351161127f576040517ff40239db000000000000000000000000000000000000000000000000000000008152367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90033560048201526024015b60405180910390fd5b60463611156112965763c407e0256000526004601cfd5b6040805160e08101825263ffffffff8152600060208201523291810191909152346fffffffffffffffffffffffffffffffff16606082015260029060808101367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900335815260016020820152604001426fffffffffffffffffffffffffffffffff90811690915282546001808201855560009485526020808620855160059094020180549186015163ffffffff9094167fffffffffffffffff0000000000000000000000000000000000000000000000009092169190911764010000000073ffffffffffffffffffffffffffffffffffffffff94851602178155604085015181830180547fffffffffffffffffffffffff000000000000000000000000000000000000000016919094161790925560608401516002830180547fffffffffffffffffffffffffffffffff00000000000000000000000000000000169185169190911790556080840151600383015560a084015160c09094015193831670010000000000000000000000000000000094909316939093029190911760049091015581547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000164267ffffffffffffffff161790915561147a90436135e4565b40600155600680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16610100179055565b61092a828260016109f0565b600281815481106114c857600080fd5b60009182526020909120600590910201805460018201546002830154600384015460049094015463ffffffff8416955064010000000090930473ffffffffffffffffffffffffffffffffffffffff908116949216926fffffffffffffffffffffffffffffffff91821692918082169170010000000000000000000000000000000090041687565b60008054700100000000000000000000000000000000900460ff16600281111561157b5761157b613256565b146115b2576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600287815481106115c7576115c761356e565b6000918252602082206005919091020160048101549092506fffffffffffffffffffffffffffffffff16908715821760011b90506116267f000000000000000000000000000000000000000000000000000000000000000060016135cc565b6116c2826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff1614611703576040517f5f53dd9800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008089156117f2576117567f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006135e4565b6001901b611775846fffffffffffffffffffffffffffffffff166126b4565b67ffffffffffffffff16611789919061362a565b156117c6576117bd6117ae60016fffffffffffffffffffffffffffffffff871661363e565b865463ffffffff16600061275a565b600301546117e8565b7f00000000000000000000000000000000000000000000000000000000000000005b915084905061181c565b600385015491506118196117ae6fffffffffffffffffffffffffffffffff8616600161366f565b90505b600882901b60088a8a6040516118339291906136a3565b6040518091039020901b14611874576040517f696550ff00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061187f8c61283e565b9050600061188e836003015490565b6040517fe14ced320000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063e14ced3290611908908f908f908f908f908a906004016136fc565b6020604051808303816000875af1158015611927573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061194b9190613736565b6004850154911491506000906002906119f6906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b611a92896fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b611a9c919061374f565b611aa69190613770565b67ffffffffffffffff161590508115158103611aee576040517ffb4e40dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505085547fffffffffffffffff0000000000000000000000000000000000000000ffffffff163364010000000002179095555050505050505050505050565b60008054700100000000000000000000000000000000900460ff166002811115611b5957611b59613256565b14611b90576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080600080611b9f8661286d565b93509350935093506000611bb585858585612c9a565b905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16637dc0d1d06040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c489190613797565b905060018903611d105773ffffffffffffffffffffffffffffffffffffffff81166352f0f3ad8a846001545b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815260048101939093526024830191909152604482015260206064820152608481018a905260a4015b6020604051808303816000875af1158015611ce6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d0a9190613736565b50611f8b565b60028903611d3c5773ffffffffffffffffffffffffffffffffffffffff81166352f0f3ad8a8489611c74565b60038903611d685773ffffffffffffffffffffffffffffffffffffffff81166352f0f3ad8a8487611c74565b60048903611ee05760006fffffffffffffffffffffffffffffffff861615611e0057611dc66fffffffffffffffffffffffffffffffff87167f0000000000000000000000000000000000000000000000000000000000000000612d59565b611df0907f00000000000000000000000000000000000000000000000000000000000000006135cc565b611dfb9060016135cc565b611e22565b7f00000000000000000000000000000000000000000000000000000000000000005b905073ffffffffffffffffffffffffffffffffffffffff82166352f0f3ad8b8560405160e084901b7fffffffff000000000000000000000000000000000000000000000000000000001681526004810192909252602482015260c084901b604482015260086064820152608481018b905260a4016020604051808303816000875af1158015611eb5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ed99190613736565b5050611f8b565b60058903611f59576040517f52f0f3ad000000000000000000000000000000000000000000000000000000008152600481018a9052602481018390524660c01b6044820152600860648201526084810188905273ffffffffffffffffffffffffffffffffffffffff8216906352f0f3ad9060a401611cc7565b6040517fff137e6500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050505050505050565b7f0000000000000000000000000000000000000000000000000000000000000000367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003356060611fec61092e565b9050909192565b60008054700100000000000000000000000000000000900460ff16600281111561201f5761201f613256565b14612056576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006002828154811061206b5761206b61356e565b6000918252602082206005919091020160048101549092506120ad90700100000000000000000000000000000000900460401c67ffffffffffffffff16610d5e565b60048301549091506000906120df90700100000000000000000000000000000000900467ffffffffffffffff16610d5e565b6120e9904261374f565b9050677fffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000060011c1661212382846137b4565b67ffffffffffffffff1611612164576040517ff2440b5300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000848152600560205260409020805485158015612184575060065460ff165b156121bb576040517ff1a9458100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b801580156121c857508515155b1561222d578454640100000000900473ffffffffffffffffffffffffffffffffffffffff16600081156121fb5781612217565b600187015473ffffffffffffffffffffffffffffffffffffffff165b90506122238188612e0e565b5050505050505050565b60006fffffffffffffffffffffffffffffffff815b8381101561237357600085828154811061225e5761225e61356e565b60009182526020808320909101548083526005909152604090912054909150156122b4576040517f9a07664600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600282815481106122c9576122c961356e565b600091825260209091206005909102018054909150640100000000900473ffffffffffffffffffffffffffffffffffffffff16158015612322575060048101546fffffffffffffffffffffffffffffffff908116908516115b15612360576001810154600482015473ffffffffffffffffffffffffffffffffffffffff90911695506fffffffffffffffffffffffffffffffff1693505b50508061236c906137d7565b9050612242565b506123bb73ffffffffffffffffffffffffffffffffffffffff83161561239957826123b5565b600188015473ffffffffffffffffffffffffffffffffffffffff165b88612e0e565b86547fffffffffffffffff0000000000000000000000000000000000000000ffffffff1664010000000073ffffffffffffffffffffffffffffffffffffffff84160217875560008881526005602052604081206124179161321c565b8760000361222357600680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555050505050505050565b6060600061248b84367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036135cc565b90508267ffffffffffffffff1667ffffffffffffffff8111156124b0576124b061380f565b6040519080825280601f01601f1916602001820160405280156124da576020820181803683370190505b509150828160208401375092915050565b151760011b90565b60006125126fffffffffffffffffffffffffffffffff8416600161366f565b905060006125228286600161275a565b9050600086901a8380612615575061255b60027f000000000000000000000000000000000000000000000000000000000000000061362a565b60048301546002906125ff906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b6126099190613770565b67ffffffffffffffff16145b1561266d5760ff81166001148061262f575060ff81166002145b612668576040517ff40239db00000000000000000000000000000000000000000000000000000000815260048101889052602401611276565b6126ab565b60ff8116156126ab576040517ff40239db00000000000000000000000000000000000000000000000000000000815260048101889052602401611276565b50505050505050565b600080612741837e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b600167ffffffffffffffff919091161b90920392915050565b600080826127a35761279e6fffffffffffffffffffffffffffffffff86167f0000000000000000000000000000000000000000000000000000000000000000612ef9565b6127be565b6127be856fffffffffffffffffffffffffffffffff166130c0565b9050600284815481106127d3576127d361356e565b906000526020600020906005020191505b60048201546fffffffffffffffffffffffffffffffff82811691161461283657815460028054909163ffffffff169081106128215761282161356e565b906000526020600020906005020191506127e4565b509392505050565b600080600080600061284f8661286d565b935093509350935061286384848484612c9a565b9695505050505050565b600080600080600085905060006002828154811061288d5761288d61356e565b600091825260209091206004600590920201908101549091507f000000000000000000000000000000000000000000000000000000000000000090612964906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff16116129a5576040517fb34b5c2200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000815b60048301547f000000000000000000000000000000000000000000000000000000000000000090612a6c906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169250821115612ae857825463ffffffff16612ab27f000000000000000000000000000000000000000000000000000000000000000060016135cc565b8303612abc578391505b60028181548110612acf57612acf61356e565b90600052602060002090600502019350809450506129a9565b600481810154908401546fffffffffffffffffffffffffffffffff91821691166000816fffffffffffffffffffffffffffffffff16612b51612b3c856fffffffffffffffffffffffffffffffff1660011c90565b6fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff161490508015612c36576000612b89836fffffffffffffffffffffffffffffffff166126b4565b67ffffffffffffffff161115612bec576000612bc3612bbb60016fffffffffffffffffffffffffffffffff861661363e565b89600161275a565b6003810154600490910154909c506fffffffffffffffffffffffffffffffff169a50612c109050565b7f00000000000000000000000000000000000000000000000000000000000000009a505b600386015460048701549099506fffffffffffffffffffffffffffffffff169750612c8c565b6000612c58612bbb6fffffffffffffffffffffffffffffffff8516600161366f565b6003808901546004808b015492840154930154909e506fffffffffffffffffffffffffffffffff9182169d50919b50169850505b505050505050509193509193565b60006fffffffffffffffffffffffffffffffff84168103612d00578282604051602001612ce39291909182526fffffffffffffffffffffffffffffffff16602082015260400190565b604051602081830303815290604052805190602001209050612d51565b60408051602081018790526fffffffffffffffffffffffffffffffff8087169282019290925260608101859052908316608082015260a0016040516020818303038152906040528051906020012090505b949350505050565b600080612de6847e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff1690508083036001841b600180831b0386831b17039250505092915050565b60028101546fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff000000000000000000000000000000018101612e7e576040517ff1a9458100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002820180547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff17905573ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604081208054839290612eef9084906135cc565b9091555050505050565b600081612f98846fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff1611612fd9576040517fb34b5c2200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612fe2836130c0565b905081613081826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff16116130ba576130b761309e8360016135cc565b6fffffffffffffffffffffffffffffffff83169061316c565b90505b92915050565b60008119600183011681613154827e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169390931c8015179392505050565b6000806131f9847e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169050808303600180821b0385821b179250505092915050565b508054600082559060005260206000209081019061323a919061323d565b50565b5b80821115613252576000815560010161323e565b5090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60208101600383106132c0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b600080604083850312156132d957600080fd5b50508035926020909101359150565b6000815180845260005b8181101561330e576020818501810151868301820152016132f2565b81811115613320576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006130b760208301846132e8565b73ffffffffffffffffffffffffffffffffffffffff8116811461323a57600080fd5b60006020828403121561339a57600080fd5b81356133a581613366565b9392505050565b803580151581146133bc57600080fd5b919050565b6000806000606084860312156133d657600080fd5b83359250602084013591506133ed604085016133ac565b90509250925092565b60006020828403121561340857600080fd5b81356fffffffffffffffffffffffffffffffff811681146133a557600080fd5b60006020828403121561343a57600080fd5b5035919050565b60008083601f84011261345357600080fd5b50813567ffffffffffffffff81111561346b57600080fd5b60208301915083602082850101111561348357600080fd5b9250929050565b600080600080600080608087890312156134a357600080fd5b863595506134b3602088016133ac565b9450604087013567ffffffffffffffff808211156134d057600080fd5b6134dc8a838b01613441565b909650945060608901359150808211156134f557600080fd5b5061350289828a01613441565b979a9699509497509295939492505050565b60008060006060848603121561352957600080fd5b505081359360208301359350604090920135919050565b63ffffffff8416815282602082015260606040820152600061356560608301846132e8565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156135df576135df61359d565b500190565b6000828210156135f6576135f661359d565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082613639576136396135fb565b500690565b60006fffffffffffffffffffffffffffffffff838116908316818110156136675761366761359d565b039392505050565b60006fffffffffffffffffffffffffffffffff80831681851680830382111561369a5761369a61359d565b01949350505050565b8183823760009101908152919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6060815260006137106060830187896136b3565b82810360208401526137238186886136b3565b9150508260408301529695505050505050565b60006020828403121561374857600080fd5b5051919050565b600067ffffffffffffffff838116908316818110156136675761366761359d565b600067ffffffffffffffff8084168061378b5761378b6135fb565b92169190910692915050565b6000602082840312156137a957600080fd5b81516133a581613366565b600067ffffffffffffffff80831681851680830382111561369a5761369a61359d565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036138085761380861359d565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(FaultDisputeGameStorageLayoutJSON), FaultDisputeGameStorageLayout); err != nil {
		panic(err)
	}

	layouts["FaultDisputeGame"] = FaultDisputeGameStorageLayout
	deployedBytecodes["FaultDisputeGame"] = FaultDisputeGameDeployedBin
	immutableReferences["FaultDisputeGame"] = true
}
