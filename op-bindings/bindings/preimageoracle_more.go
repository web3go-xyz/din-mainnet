// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const PreimageOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageLengths\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1001,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageParts\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\"},{\"astId\":1002,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimagePartOk\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\"},{\"astId\":1003,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"zeroHashes\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_bytes32)16_storage\"},{\"astId\":1004,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposals\",\"offset\":0,\"slot\":\"19\",\"type\":\"t_array(t_struct(LargePreimageProposalKeys)1009_storage)dyn_storage\"},{\"astId\":1005,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalBranches\",\"offset\":0,\"slot\":\"20\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_bytes32)16_storage))\"},{\"astId\":1006,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalMetadata\",\"offset\":0,\"slot\":\"21\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1010))\"},{\"astId\":1007,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalParts\",\"offset\":0,\"slot\":\"22\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_bytes32))\"},{\"astId\":1008,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalBlocks\",\"offset\":0,\"slot\":\"23\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_uint64)dyn_storage))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_bytes32)16_storage\":{\"encoding\":\"inplace\",\"label\":\"bytes32[16]\",\"numberOfBytes\":\"512\",\"base\":\"t_bytes32\"},\"t_array(t_struct(LargePreimageProposalKeys)1009_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct PreimageOracle.LargePreimageProposalKeys[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(LargePreimageProposalKeys)1009_storage\"},\"t_array(t_uint64)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"uint64[]\",\"numberOfBytes\":\"32\",\"base\":\"t_uint64\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_bytes32)16_storage))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e bytes32[16]))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_array(t_bytes32)16_storage)\"},\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_uint64)dyn_storage))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e uint64[]))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_array(t_uint64)dyn_storage)\"},\"t_mapping(t_address,t_mapping(t_uint256,t_bytes32))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e bytes32))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_bytes32)\"},\"t_mapping(t_address,t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1010))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e LPPMetaData))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1010)\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bytes32))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bytes32)\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_array(t_bytes32)16_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bytes32[16])\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_array(t_bytes32)16_storage\"},\"t_mapping(t_uint256,t_array(t_uint64)dyn_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint64[])\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_array(t_uint64)dyn_storage\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_mapping(t_uint256,t_bytes32)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bytes32)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bytes32\"},\"t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1010)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e LPPMetaData)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_userDefinedValueType(LPPMetaData)1010\"},\"t_struct(LargePreimageProposalKeys)1009_storage\":{\"encoding\":\"inplace\",\"label\":\"struct PreimageOracle.LargePreimageProposalKeys\",\"numberOfBytes\":\"64\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_userDefinedValueType(LPPMetaData)1010\":{\"encoding\":\"inplace\",\"label\":\"LPPMetaData\",\"numberOfBytes\":\"32\"}}}"

var PreimageOracleStorageLayout = new(solc.StorageLayout)

var PreimageOracleDeployedBin = "0x608060405234801561001057600080fd5b506004361061018d5760003560e01c80639d53a648116100e3578063da35c6641161008c578063ec5efcbc11610066578063ec5efcbc146103f5578063faf37bc714610408578063fef2b4ed1461041b57600080fd5b8063da35c664146103b2578063e03110e1146103ba578063e1592611146103e257600080fd5b8063b4801e61116100bd578063b4801e6114610382578063c3a079ed14610395578063d18534b51461039f57600080fd5b80639d53a648146103025780639f99ef8214610344578063b2e67ba81461035757600080fd5b806352f0f3ad116101455780637ac547671161011f5780637ac54767146102855780638542cf5014610298578063882856ef146102d657600080fd5b806352f0f3ad1461021c57806361238bde1461022f5780636551927b1461025a57600080fd5b80632055b36b116101765780632055b36b146101f75780633909af5c146101ff5780634d52b4c91461021457600080fd5b8063013cf08b146101925780630359a563146101d6575b600080fd5b6101a56101a036600461252c565b61043b565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152015b60405180910390f35b6101e96101e436600461256e565b610480565b6040519081526020016101cd565b6101e9601081565b61021261020d366004612769565b6105b8565b005b6101e9610806565b6101e961022a366004612855565b610821565b6101e961023d366004612890565b600160209081526000928352604080842090915290825290205481565b6101e961026836600461256e565b601560209081526000928352604080842090915290825290205481565b6101e961029336600461252c565b6108f6565b6102c66102a6366004612890565b600260209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016101cd565b6102e96102e43660046128b2565b61090d565b60405167ffffffffffffffff90911681526020016101cd565b6101e961031036600461256e565b73ffffffffffffffffffffffffffffffffffffffff9091166000908152601760209081526040808320938352929052205490565b610212610352366004612927565b610967565b6101e961036536600461256e565b601660209081526000928352604080842090915290825290205481565b6101e96103903660046128b2565b610ee2565b6101e96201518081565b6102126103ad366004612769565b610f14565b6013546101e9565b6103cd6103c8366004612890565b6112c2565b604080519283526020830191909152016101cd565b6102126103f03660046129b8565b6113b3565b610212610403366004612a04565b6114bc565b610212610416366004612a9d565b611636565b6101e961042936600461252c565b60006020819052908152604090205481565b6013818154811061044b57600080fd5b60009182526020909120600290910201805460019091015473ffffffffffffffffffffffffffffffffffffffff909116915082565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260156020908152604080832084845290915281205481906104c39060601c63ffffffff1690565b63ffffffff16905060005b60108110156105b057816001166001036105565773ffffffffffffffffffffffffffffffffffffffff851660009081526014602090815260408083208784529091529020816010811061052357610523612ad9565b01546040805160208101929092528101849052606001604051602081830303815290604052805190602001209250610597565b826003826010811061056a5761056a612ad9565b01546040805160208101939093528201526060016040516020818303038152906040528051906020012092505b60019190911c90806105a881612b37565b9150506104ce565b505092915050565b60006105c48a8a610480565b90506105e786868360208b01356105e26105dd8d612b6f565b6117e9565b611829565b8015610605575061060583838360208801356105e26105dd8a612b6f565b61063b576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8660400135886040516020016106519190612c3e565b604051602081830303815290604052805190602001201461069e576040517f1968a90200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8360200135876020013560016106b49190612c7c565b146106eb576040517f9a3b119900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610733886106f98680612c94565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061188a92505050565b61073c886119e5565b8360400135886040516020016107529190612c3e565b604051602081830303815290604052805190602001200361079f576040517f9843145b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505073ffffffffffffffffffffffffffffffffffffffff9590951660009081526015602090815260408083209683529590529390932080547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117905550505050565b600161081460106002612e1b565b61081e9190612e27565b81565b600061082d8686612281565b905061083a836008612c7c565b8211806108475750602083115b1561087e576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602081815260c085901b82526008959095528251828252600286526040808320858452875280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558484528752808320948352938652838220558181529384905292205592915050565b6003816010811061090657600080fd5b0154905081565b6017602052826000526040600020602052816000526040600020818154811061093557600080fd5b906000526020600020906004918282040191900660080292509250509054906101000a900467ffffffffffffffff1681565b3332146109a0576040517fba092d1600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606081156109b9576109b2868661232e565b90506109f3565b85858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293505050505b3360009081526014602090815260408083208a845290915280822081516102008101928390529160109082845b815481526020019060010190808311610a205750503360009081526015602090815260408083208f8452909152902054939450610a6292508391506123b79050565b63ffffffff16600003610aa1576040517f87138d5c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610aab8160c01c90565b67ffffffffffffffff1615610aec576040517f475a253500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610afe8260a01c63ffffffff1690565b67ffffffffffffffff1690506000610b1c8360401c63ffffffff1690565b63ffffffff169050600882108015610b32575080155b15610bb9576000610b498460801c63ffffffff1690565b905060008160c01b6000528b356008528351905080601660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008f8152602001908152602001600020819055505050610c6e565b60088210158015610bd7575080610bd1600884612e27565b92508210155b8015610beb5750610be88982612c7c565b82105b15610c6e576000610bfc8284612e27565b905089610c0a826020612c7c565b10158015610c16575086155b15610c4d576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526016602090815260408083208f84529091529020908b013590555b6000610c808460601c63ffffffff1690565b63ffffffff169050855160208701608882048a1415608883061715610cad576307b1daf16000526004601cfd5b60405160c8810160405260005b83811015610d5d578083018051835260208101516020840152604081015160408401526060810151606084015260808101516080840152508460888301526088810460051b8d013560a883015260c882206001860195508560005b610200811015610d52576001821615610d325782818d0152610d52565b8b81015160009081526020938452604090209260019290921c9101610d15565b505050608801610cba565b50505050600160106002610d719190612e1b565b610d7b9190612e27565b811115610db4576040517f6229572300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526014602090815260408083208f84529091529020610dda908660106124a2565b503360009081526017602090815260408083208f845282528220805460018101825590835291206004820401805460039092166008026101000a67ffffffffffffffff818102199093164390931602919091179055610e8e610e3c838c612c7c565b60401b7fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff606084901b167fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff8716171790565b93508615610eb95777ffffffffffffffffffffffffffffffffffffffffffffffff84164260c01b1793505b50503360009081526015602090815260408083209c83529b905299909920555050505050505050565b60146020528260005260406000206020528160005260406000208160108110610f0a57600080fd5b0154925083915050565b73ffffffffffffffffffffffffffffffffffffffff891660009081526015602090815260408083208b845290915290205467ffffffffffffffff811615610f87576040517fc334f06900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62015180610f958260c01c90565b610fa99067ffffffffffffffff1642612e27565b11610fe0576040517f55d4cbf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610fec8b8b610480565b905061100587878360208c01356105e26105dd8e612b6f565b8015611023575061102384848360208901356105e26105dd8b612b6f565b611059576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b87604001358960405160200161106f9190612c3e565b60405160208183030381529060405280519060200120146110bc576040517f1968a90200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8460200135886020013560016110d29190612c7c565b141580611104575060016110ec8360601c63ffffffff1690565b6110f69190612e3e565b63ffffffff16856020013514155b1561113b576040517f9a3b119900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061114d8360801c63ffffffff1690565b63ffffffff169050806111668460401c63ffffffff1690565b63ffffffff16146111a3576040517f7b1dafd100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111b18a6106f98880612c94565b6111ba8a6119e5565b60006111c58b6123c3565b905060006111d98560a01c63ffffffff1690565b67ffffffffffffffff169050600160026000848152602001908152602001600020600083815260200190815260200160002060006101000a81548160ff021916908315150217905550601660008f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008e8152602001908152602001600020546001600084815260200190815260200160002060008381526020019081526020016000208190555082600080848152602001908152602001600020819055505050505050505050505050505050565b6000828152600260209081526040808320848452909152812054819060ff1661134b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7072652d696d616765206d757374206578697374000000000000000000000000604482015260640160405180910390fd5b5060008381526020818152604090912054611367816008612c7c565b611372856020612c7c565b106113905783611383826008612c7c565b61138d9190612e27565b91505b506000938452600160209081526040808620948652939052919092205492909150565b604435600080600883018611156113d25763fe2549876000526004601cfd5b60c083901b6080526088838682378087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80151908490207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02000000000000000000000000000000000000000000000000000000000000001760008181526002602090815260408083208b8452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209a83529981528982209390935590815290819052959095209190915550505050565b60006114c88686610480565b90506114e183838360208801356105e26105dd8a612b6f565b611517576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602084013515611553576040517f9a3b119900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61155b6124e0565b611569816106f98780612c94565b611572816119e5565b8460400135816040516020016115889190612c3e565b60405160208183030381529060405280519060200120036115d5576040517f9843145b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505073ffffffffffffffffffffffffffffffffffffffff9290921660009081526015602090815260408083209383529290522080547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117905550565b33321461166f576040517fba092d1600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61167a816008612e63565b63ffffffff168263ffffffff16106116be576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336000818152601560209081526040808320878452825280832080547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1660a09790971b7fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff169690961760809590951b9490941790945582518084019093529082529181019283526013805460018101825592525160029091027f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a0908101805473ffffffffffffffffffffffffffffffffffffffff9093167fffffffffffffffffffffffff00000000000000000000000000000000000000009093169290921790915590517f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a09190910155565b600081600001518260200151836040015160405160200161180c93929190612e8b565b604051602081830303815290604052805190602001209050919050565b60008160005b601081101561187d578060051b880135600186831c16600181146118625760008481526020839052604090209350611873565b600082815260208590526040902093505b505060010161182f565b5090931495945050505050565b608881511461189857600080fd5b6020810160208301611919565b8260031b8201518060001a8160011a60081b178160021a60101b8260031a60181b17178160041a60201b8260051a60281b178260061a60301b8360071a60381b1717179050611913816118fe868560059190911b015190565b1867ffffffffffffffff16600586901b840152565b50505050565b611925600083836118a5565b611931600183836118a5565b61193d600283836118a5565b611949600383836118a5565b611955600483836118a5565b611961600583836118a5565b61196d600683836118a5565b611979600783836118a5565b611985600883836118a5565b611991600983836118a5565b61199d600a83836118a5565b6119a9600b83836118a5565b6119b5600c83836118a5565b6119c1600d83836118a5565b6119cd600e83836118a5565b6119d9600f83836118a5565b611913601083836118a5565b6040805178010000000000008082800000000000808a8000000080008000602082015279808b00000000800000018000000080008081800000000000800991810191909152788a00000000000000880000000080008009000000008000000a60608201527b8000808b800000000000008b8000000000008089800000000000800360808201527f80000000000080028000000000000080000000000000800a800000008000000a60a08201527f800000008000808180000000000080800000000080000001800000008000800860c082015260009060e00160405160208183030381529060405290506020820160208201612161565b6102808101516101e082015161014083015160a0840151845118189118186102a082015161020083015161016084015160c0850151602086015118189118186102c083015161022084015161018085015160e0860151604087015118189118186102e08401516102408501516101a0860151610100870151606088015118189118186103008501516102608601516101c0870151610120880151608089015118189118188084603f1c611b988660011b67ffffffffffffffff1690565b18188584603f1c611bb38660011b67ffffffffffffffff1690565b18188584603f1c611bce8660011b67ffffffffffffffff1690565b181895508483603f1c611beb8560011b67ffffffffffffffff1690565b181894508387603f1c611c088960011b67ffffffffffffffff1690565b60208b01518b51861867ffffffffffffffff168c5291189190911897508118600181901b603f9190911c18935060c08801518118601481901c602c9190911b1867ffffffffffffffff1660208901526101208801518718602c81901c60149190911b1867ffffffffffffffff1660c08901526102c08801518618600381901c603d9190911b1867ffffffffffffffff166101208901526101c08801518718601981901c60279190911b1867ffffffffffffffff166102c08901526102808801518218602e81901c60129190911b1867ffffffffffffffff166101c089015260408801518618600281901c603e9190911b1867ffffffffffffffff166102808901526101808801518618601581901c602b9190911b1867ffffffffffffffff1660408901526101a08801518518602781901c60199190911b1867ffffffffffffffff166101808901526102608801518718603881901c60089190911b1867ffffffffffffffff166101a08901526102e08801518518600881901c60389190911b1867ffffffffffffffff166102608901526101e08801518218601781901c60299190911b1867ffffffffffffffff166102e089015260808801518718602581901c601b9190911b1867ffffffffffffffff166101e08901526103008801518718603281901c600e9190911b1867ffffffffffffffff1660808901526102a08801518118603e81901c60029190911b1867ffffffffffffffff166103008901526101008801518518600981901c60379190911b1867ffffffffffffffff166102a08901526102008801518118601381901c602d9190911b1867ffffffffffffffff1661010089015260a08801518218601c81901c60249190911b1867ffffffffffffffff1661020089015260608801518518602481901c601c9190911b1867ffffffffffffffff1660a08901526102408801518518602b81901c60159190911b1867ffffffffffffffff1660608901526102208801518618603181901c600f9190911b1867ffffffffffffffff166102408901526101608801518118603681901c600a9190911b1867ffffffffffffffff166102208901525060e08701518518603a81901c60069190911b1867ffffffffffffffff166101608801526101408701518118603d81901c60039190911b1867ffffffffffffffff1660e0880152505067ffffffffffffffff81166101408601525050505050565b611f8881611adb565b805160208201805160408401805160608601805160808801805167ffffffffffffffff871986168a188116808c528619851689188216909952831982169095188516909552841988169091188316909152941990921618811690925260a08301805160c0808601805160e0880180516101008a0180516101208c018051861985168a188d16909a528319821686188c16909652801989169092188a169092528619861618881690529219909216909218841690526101408401805161016086018051610180880180516101a08a0180516101c08c0180518619851689188d169099528319821686188c16909652801988169092188a169092528519851618881690529119909116909118841690526101e08401805161020086018051610220880180516102408a0180516102608c0180518619851689188d169099528319821686188c16909652801988169092188a16909252851985161888169052911990911690911884169052610280840180516102a0860180516102c0880180516102e08a0180516103008c0180518619851689188d169099528319821686188c16909652801988169092188a16909252851985161888169052911990911690911884169052600386901b850151901c9081189091168252611913565b61216d60008284611f7f565b61217960018284611f7f565b61218560028284611f7f565b61219160038284611f7f565b61219d60048284611f7f565b6121a960058284611f7f565b6121b560068284611f7f565b6121c160078284611f7f565b6121cd60088284611f7f565b6121d960098284611f7f565b6121e5600a8284611f7f565b6121f1600b8284611f7f565b6121fd600c8284611f7f565b612209600d8284611f7f565b612215600e8284611f7f565b612221600f8284611f7f565b61222d60108284611f7f565b61223960118284611f7f565b61224560128284611f7f565b61225160138284611f7f565b61225d60148284611f7f565b61226960158284611f7f565b61227560168284611f7f565b61191360178284611f7f565b7f01000000000000000000000000000000000000000000000000000000000000007effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831617612327818360408051600093845233602052918152606090922091527effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01000000000000000000000000000000000000000000000000000000000000001790565b9392505050565b606060405190508160208201818101828683376088830680801561238c576088829003850160808582017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01536001845160001a178453865261239e565b60018353608060878401536088850186525b5050505050601f19603f82510116810160405292915050565b60801c63ffffffff1690565b6000612446565b66ff00ff00ff00ff8160081c1667ff00ff00ff00ff006123f48360081b67ffffffffffffffff1690565b1617905065ffff0000ffff8160101c1667ffff0000ffff00006124218360101b67ffffffffffffffff1690565b1617905060008160201c61243f8360201b67ffffffffffffffff1690565b1792915050565b6080820151602083019061245e906123ca565b6123ca565b604082015161246c906123ca565b60401b1761248461245960018460059190911b015190565b825160809190911b90612496906123ca565b60c01b17179392505050565b82601081019282156124d0579160200282015b828111156124d05782518255916020019190600101906124b5565b506124dc9291506124f8565b5090565b60405180602001604052806124f361250d565b905290565b5b808211156124dc57600081556001016124f9565b6040518061032001604052806019906020820280368337509192915050565b60006020828403121561253e57600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461256957600080fd5b919050565b6000806040838503121561258157600080fd5b61258a83612545565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610320810167ffffffffffffffff811182821017156125eb576125eb612598565b60405290565b6040516060810167ffffffffffffffff811182821017156125eb576125eb612598565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561265b5761265b612598565b604052919050565b600061032080838503121561267757600080fd5b604051602080820167ffffffffffffffff838210818311171561269c5761269c612598565b8160405283955087601f8801126126b257600080fd5b6126ba6125c7565b94870194915081888611156126ce57600080fd5b875b868110156126f657803583811681146126e95760008081fd5b84529284019284016126d0565b50909352509295945050505050565b60006060828403121561271757600080fd5b50919050565b60008083601f84011261272f57600080fd5b50813567ffffffffffffffff81111561274757600080fd5b6020830191508360208260051b850101111561276257600080fd5b9250929050565b60008060008060008060008060006103e08a8c03121561278857600080fd5b6127918a612545565b985060208a013597506127a78b60408c01612663565b96506103608a013567ffffffffffffffff808211156127c557600080fd5b6127d18d838e01612705565b97506103808c01359150808211156127e857600080fd5b6127f48d838e0161271d565b90975095506103a08c013591508082111561280e57600080fd5b61281a8d838e01612705565b94506103c08c013591508082111561283157600080fd5b5061283e8c828d0161271d565b915080935050809150509295985092959850929598565b600080600080600060a0868803121561286d57600080fd5b505083359560208501359550604085013594606081013594506080013592509050565b600080604083850312156128a357600080fd5b50508035926020909101359150565b6000806000606084860312156128c757600080fd5b6128d084612545565b95602085013595506040909401359392505050565b60008083601f8401126128f757600080fd5b50813567ffffffffffffffff81111561290f57600080fd5b60208301915083602082850101111561276257600080fd5b6000806000806000806080878903121561294057600080fd5b86359550602087013567ffffffffffffffff8082111561295f57600080fd5b61296b8a838b016128e5565b9097509550604089013591508082111561298457600080fd5b5061299189828a0161271d565b909450925050606087013580151581146129aa57600080fd5b809150509295509295509295565b6000806000604084860312156129cd57600080fd5b83359250602084013567ffffffffffffffff8111156129eb57600080fd5b6129f7868287016128e5565b9497909650939450505050565b600080600080600060808688031215612a1c57600080fd5b612a2586612545565b945060208601359350604086013567ffffffffffffffff80821115612a4957600080fd5b612a5589838a01612705565b94506060880135915080821115612a6b57600080fd5b50612a788882890161271d565b969995985093965092949392505050565b803563ffffffff8116811461256957600080fd5b600080600060608486031215612ab257600080fd5b83359250612ac260208501612a89565b9150612ad060408501612a89565b90509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612b6857612b68612b08565b5060010190565b600060608236031215612b8157600080fd5b612b896125f1565b823567ffffffffffffffff80821115612ba157600080fd5b9084019036601f830112612bb457600080fd5b8135602082821115612bc857612bc8612598565b612bf8817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85011601612614565b92508183523681838601011115612c0e57600080fd5b81818501828501376000918301810191909152908352848101359083015250604092830135928101929092525090565b81516103208201908260005b6019811015612c7357825167ffffffffffffffff16825260209283019290910190600101612c4a565b50505092915050565b60008219821115612c8f57612c8f612b08565b500190565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112612cc957600080fd5b83018035915067ffffffffffffffff821115612ce457600080fd5b60200191503681900382131561276257600080fd5b600181815b80851115612d5257817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115612d3857612d38612b08565b80851615612d4557918102915b93841c9390800290612cfe565b509250929050565b600082612d6957506001612e15565b81612d7657506000612e15565b8160018114612d8c5760028114612d9657612db2565b6001915050612e15565b60ff841115612da757612da7612b08565b50506001821b612e15565b5060208310610133831016604e8410600b8410161715612dd5575081810a612e15565b612ddf8383612cf9565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115612e1157612e11612b08565b0290505b92915050565b60006123278383612d5a565b600082821015612e3957612e39612b08565b500390565b600063ffffffff83811690831681811015612e5b57612e5b612b08565b039392505050565b600063ffffffff808316818516808303821115612e8257612e82612b08565b01949350505050565b6000845160005b81811015612eac5760208188018101518583015201612e92565b81811115612ebb576000828501525b509190910192835250602082015260400191905056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(PreimageOracleStorageLayoutJSON), PreimageOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["PreimageOracle"] = PreimageOracleStorageLayout
	deployedBytecodes["PreimageOracle"] = PreimageOracleDeployedBin
	immutableReferences["PreimageOracle"] = false
}
