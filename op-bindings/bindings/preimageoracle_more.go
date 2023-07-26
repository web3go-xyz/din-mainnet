// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const PreimageOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageLengths\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1001,\"contract\":\"contracts/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageParts\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\"},{\"astId\":1002,\"contract\":\"contracts/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimagePartOk\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\"}],\"types\":{\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bytes32))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bytes32)\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_mapping(t_uint256,t_bytes32)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bytes32)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bytes32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var PreimageOracleStorageLayout = new(solc.StorageLayout)

var PreimageOracleDeployedBin = "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c8063e03110e11161005b578063e03110e114610111578063e159261114610139578063fe4ac08e1461014e578063fef2b4ed146101c357600080fd5b806361238bde146100825780638542cf50146100c0578063a57c202c146100fe575b600080fd5b6100ad610090366004610433565b600160209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b6100ee6100ce366004610433565b600260209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016100b7565b6100ad61010c36600461049e565b6101e3565b61012461011f366004610433565b610242565b604080519283526020830191909152016100b7565b61014c6101473660046104e0565b610333565b005b61014c61015c36600461052c565b6000838152600260209081526040808320878452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558684528252808320968352958152858220939093559283529082905291902055565b6100ad6101d136600461055e565b60006020819052908152604090205481565b60243560c081901b608052600090608881858237207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f0200000000000000000000000000000000000000000000000000000000000000179392505050565b6000828152600260209081526040808320848452909152812054819060ff166102cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7072652d696d616765206d757374206578697374000000000000000000000000604482015260640160405180910390fd5b50600083815260208181526040909120546102e78160086105a6565b6102f28560206105a6565b1061031057836103038260086105a6565b61030d91906105bf565b91505b506000938452600160209081526040808620948652939052919092205492909150565b6044356000806008830186111561034957600080fd5b60c083901b6080526088838682378087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80151908490207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02000000000000000000000000000000000000000000000000000000000000001760008181526002602090815260408083208b8452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209a83529981528982209390935590815290819052959095209190915550505050565b6000806040838503121561044657600080fd5b50508035926020909101359150565b60008083601f84011261046757600080fd5b50813567ffffffffffffffff81111561047f57600080fd5b60208301915083602082850101111561049757600080fd5b9250929050565b600080602083850312156104b157600080fd5b823567ffffffffffffffff8111156104c857600080fd5b6104d485828601610455565b90969095509350505050565b6000806000604084860312156104f557600080fd5b83359250602084013567ffffffffffffffff81111561051357600080fd5b61051f86828701610455565b9497909650939450505050565b6000806000806080858703121561054257600080fd5b5050823594602084013594506040840135936060013592509050565b60006020828403121561057057600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156105b9576105b9610577565b92915050565b818103818111156105b9576105b961057756fea164736f6c6343000813000a"

var PreimageOracleDeployedSourceMap = "144:4615:9:-:0;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;357:68;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;413:25:81;;;401:2;386:18;357:68:9;;;;;;;;501:66;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;614:14:81;;607:22;589:41;;577:2;562:18;501:66:9;449:187:81;2154:850:9;;;;;;:::i;:::-;;:::i;838:564::-;;;;;;:::i;:::-;;:::i;:::-;;;;1581:25:81;;;1637:2;1622:18;;1615:34;;;;1554:18;838:564:9;1407:248:81;3295:1462:9;;;;;;:::i;:::-;;:::i;:::-;;1744:262;;;;;;:::i;:::-;1877:19;;;;:14;:19;;;;;;;;:31;;;;;;;;:38;;;;1911:4;1877:38;;;;;;1925:18;;;;;;;;:30;;;;;;;;;:37;;;;1972:20;;;;;;;;;;:27;1744:262;238:50;;;;;;:::i;:::-;;;;;;;;;;;;;;;2154:850;2321:4;2308:18;2567:3;2563:14;;;2458:4;2551:27;2231:12;;2598:11;2308:18;2718:16;2598:11;2700:41;2840:20;2954:19;2947:27;2976:11;2944:44;;2154:850;-1:-1:-1;;;2154:850:9:o;838:564::-;938:12;991:20;;;:14;:20;;;;;;;;:29;;;;;;;;;938:12;;991:29;;983:62;;;;;;;3101:2:81;983:62:9;;;3083:21:81;3140:2;3120:18;;;3113:30;3179:22;3159:18;;;3152:50;3219:18;;983:62:9;;;;;;;;-1:-1:-1;1176:14:9;1193:21;;;1164:2;1193:21;;;;;;;;1244:10;1193:21;1253:1;1244:10;:::i;:::-;1228:12;:7;1238:2;1228:12;:::i;:::-;:26;1224:87;;1293:7;1280:10;:6;1289:1;1280:10;:::i;:::-;:20;;;;:::i;:::-;1270:30;;1224:87;-1:-1:-1;1367:19:9;;;;:13;:19;;;;;;;;:28;;;;;;;;;;;;838:564;;-1:-1:-1;838:564:9:o;3295:1462::-;3591:4;3578:18;3396:12;;3720:1;3710:12;;3694:29;;3691:77;;;3752:1;3749;3742:12;3691:77;4011:3;4007:14;;;3911:4;3995:27;4042:11;4016:4;4161:16;4042:11;4143:41;4374:29;;;4378:11;4374:29;4368:36;4426:20;;;;4573:19;4566:27;4595:11;4563:44;4626:19;;;;4604:1;4626:19;;;;;;;;:32;;;;;;;;:39;;;;4661:4;4626:39;;;;;;4675:18;;;;;;;;:31;;;;;;;;;:38;;;;4723:20;;;;;;;;;;;:27;;;;-1:-1:-1;;;;3295:1462:9:o;14:248:81:-;82:6;90;143:2;131:9;122:7;118:23;114:32;111:52;;;159:1;156;149:12;111:52;-1:-1:-1;;182:23:81;;;252:2;237:18;;;224:32;;-1:-1:-1;14:248:81:o;641:347::-;692:8;702:6;756:3;749:4;741:6;737:17;733:27;723:55;;774:1;771;764:12;723:55;-1:-1:-1;797:20:81;;840:18;829:30;;826:50;;;872:1;869;862:12;826:50;909:4;901:6;897:17;885:29;;961:3;954:4;945:6;937;933:19;929:30;926:39;923:59;;;978:1;975;968:12;923:59;641:347;;;;;:::o;993:409::-;1063:6;1071;1124:2;1112:9;1103:7;1099:23;1095:32;1092:52;;;1140:1;1137;1130:12;1092:52;1180:9;1167:23;1213:18;1205:6;1202:30;1199:50;;;1245:1;1242;1235:12;1199:50;1284:58;1334:7;1325:6;1314:9;1310:22;1284:58;:::i;:::-;1361:8;;1258:84;;-1:-1:-1;993:409:81;-1:-1:-1;;;;993:409:81:o;1660:477::-;1739:6;1747;1755;1808:2;1796:9;1787:7;1783:23;1779:32;1776:52;;;1824:1;1821;1814:12;1776:52;1860:9;1847:23;1837:33;;1921:2;1910:9;1906:18;1893:32;1948:18;1940:6;1937:30;1934:50;;;1980:1;1977;1970:12;1934:50;2019:58;2069:7;2060:6;2049:9;2045:22;2019:58;:::i;:::-;1660:477;;2096:8;;-1:-1:-1;1993:84:81;;-1:-1:-1;;;;1660:477:81:o;2142:385::-;2228:6;2236;2244;2252;2305:3;2293:9;2284:7;2280:23;2276:33;2273:53;;;2322:1;2319;2312:12;2273:53;-1:-1:-1;;2345:23:81;;;2415:2;2400:18;;2387:32;;-1:-1:-1;2466:2:81;2451:18;;2438:32;;2517:2;2502:18;2489:32;;-1:-1:-1;2142:385:81;-1:-1:-1;2142:385:81:o;2532:180::-;2591:6;2644:2;2632:9;2623:7;2619:23;2615:32;2612:52;;;2660:1;2657;2650:12;2612:52;-1:-1:-1;2683:23:81;;2532:180;-1:-1:-1;2532:180:81:o;3248:184::-;3300:77;3297:1;3290:88;3397:4;3394:1;3387:15;3421:4;3418:1;3411:15;3437:125;3502:9;;;3523:10;;;3520:36;;;3536:18;;:::i;:::-;3437:125;;;;:::o;3567:128::-;3634:9;;;3655:11;;;3652:37;;;3669:18;;:::i"

func init() {
	if err := json.Unmarshal([]byte(PreimageOracleStorageLayoutJSON), PreimageOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["PreimageOracle"] = PreimageOracleStorageLayout
	deployedBytecodes["PreimageOracle"] = PreimageOracleDeployedBin
}
