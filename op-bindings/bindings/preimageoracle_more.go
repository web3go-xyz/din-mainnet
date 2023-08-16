// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const PreimageOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageLengths\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1001,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageParts\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\"},{\"astId\":1002,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimagePartOk\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\"}],\"types\":{\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bytes32))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bytes32)\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_mapping(t_uint256,t_bytes32)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bytes32)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bytes32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var PreimageOracleStorageLayout = new(solc.StorageLayout)

var PreimageOracleDeployedBin = "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c8063e03110e11161005b578063e03110e114610111578063e159261114610139578063fe4ac08e1461014e578063fef2b4ed146101c357600080fd5b806361238bde146100825780638542cf50146100c05780639a1f5e7f146100fe575b600080fd5b6100ad610090366004610551565b600160209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b6100ee6100ce366004610551565b600260209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016100b7565b6100ad61010c366004610573565b6101e3565b61012461011f366004610551565b6102b6565b604080519283526020830191909152016100b7565b61014c6101473660046105a5565b6103a7565b005b61014c61015c366004610573565b6000838152600260209081526040808320878452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558684528252808320968352958152858220939093559283529082905291902055565b6100ad6101d1366004610621565b60006020819052908152604090205481565b60006101ee856104b0565b90506101fb836008610669565b8211806102085750602083115b1561023f576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602081815260c085901b82526008959095528251828252600286526040808320858452875280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845287528083209483529386528382205581815293849052922055919050565b6000828152600260209081526040808320848452909152812054819060ff1661033f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7072652d696d616765206d757374206578697374000000000000000000000000604482015260640160405180910390fd5b506000838152602081815260409091205461035b816008610669565b610366856020610669565b106103845783610377826008610669565b6103819190610681565b91505b506000938452600160209081526040808620948652939052919092205492909150565b604435600080600883018611156103c65763fe2549876000526004601cfd5b60c083901b6080526088838682378087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80151908490207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02000000000000000000000000000000000000000000000000000000000000001760008181526002602090815260408083208b8452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209a83529981528982209390935590815290819052959095209190915550505050565b7f01000000000000000000000000000000000000000000000000000000000000007effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82161761054b81600090815233602052604090207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01000000000000000000000000000000000000000000000000000000000000001790565b92915050565b6000806040838503121561056457600080fd5b50508035926020909101359150565b6000806000806080858703121561058957600080fd5b5050823594602084013594506040840135936060013592509050565b6000806000604084860312156105ba57600080fd5b83359250602084013567ffffffffffffffff808211156105d957600080fd5b818601915086601f8301126105ed57600080fd5b8135818111156105fc57600080fd5b87602082850101111561060e57600080fd5b6020830194508093505050509250925092565b60006020828403121561063357600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561067c5761067c61063a565b500190565b6000828210156106935761069361063a565b50039056fea164736f6c634300080f000a"

var PreimageOracleDeployedSourceMap = "306:4436:106:-:0;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;537:68;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;413:25:254;;;401:2;386:18;537:68:106;;;;;;;;680:66;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;614:14:254;;607:22;589:41;;577:2;562:18;680:66:106;449:187:254;1938:1165:106;;;;;;:::i;:::-;;:::i;789:536::-;;;;;;:::i;:::-;;:::i;:::-;;;;1205:25:254;;;1261:2;1246:18;;1239:34;;;;1178:18;789:536:106;1031:248:254;3145:1595:106;;;;;;:::i;:::-;;:::i;:::-;;1672:224;;;;;;:::i;:::-;1767:19;;;;:14;:19;;;;;;;;:31;;;;;;;;:38;;;;1801:4;1767:38;;;;;;1815:18;;;;;;;;:30;;;;;;;;;:37;;;;1862:20;;;;;;;;;;:27;1672:224;419:50;;;;;;:::i;:::-;;;;;;;;;;;;;;;1938:1165;2100:12;2205:36;2234:6;2205:28;:36::i;:::-;2198:43;-1:-1:-1;2335:9:106;:5;2343:1;2335:9;:::i;:::-;2321:11;:23;:37;;;;2356:2;2348:5;:10;2321:37;2317:90;;;2381:15;;;;;;;;;;;;;;2317:90;2476:12;2576:4;2569:18;;;2677:3;2673:15;;;2660:29;;2709:4;2702:19;;;;2811:18;;2901:20;;;:14;:20;;;;;;:33;;;;;;;;:40;;;;2937:4;2901:40;;;;;;2951:19;;;;;;;;:32;;;;;;;;;:39;3067:21;;;;;;;;;:29;2916:4;1938:1165;-1:-1:-1;1938:1165:106:o;789:536::-;865:12;914:20;;;:14;:20;;;;;;;;:29;;;;;;;;;865:12;;914:29;;906:62;;;;;;;3229:2:254;906:62:106;;;3211:21:254;3268:2;3248:18;;;3241:30;3307:22;3287:18;;;3280:50;3347:18;;906:62:106;;;;;;;;-1:-1:-1;1099:14:106;1116:21;;;1087:2;1116:21;;;;;;;;1167:10;1116:21;1176:1;1167:10;:::i;:::-;1151:12;:7;1161:2;1151:12;:::i;:::-;:26;1147:87;;1216:7;1203:10;:6;1212:1;1203:10;:::i;:::-;:20;;;;:::i;:::-;1193:30;;1147:87;-1:-1:-1;1290:19:106;;;;:13;:19;;;;;;;;:28;;;;;;;;;;;;789:536;;-1:-1:-1;789:536:106:o;3145:1595::-;3441:4;3428:18;3246:12;;3570:1;3560:12;;3544:29;;3541:210;;;3645:10;3642:1;3635:21;3735:1;3729:4;3722:15;3541:210;3994:3;3990:14;;;3894:4;3978:27;4025:11;3999:4;4144:16;4025:11;4126:41;4357:29;;;4361:11;4357:29;4351:36;4409:20;;;;4556:19;4549:27;4578:11;4546:44;4609:19;;;;4587:1;4609:19;;;;;;;;:32;;;;;;;;:39;;;;4644:4;4609:39;;;;;;4658:18;;;;;;;;:31;;;;;;;;;:38;;;;4706:20;;;;;;;;;;;:27;;;;-1:-1:-1;;;;3145:1595:106:o;492:353:105:-;752:11;777:19;765:32;;749:49;824:14;749:49;1277:21;1426:15;;;1467:8;1461:4;1454:22;1595:4;1582:18;;1602:19;1578:44;1624:11;1575:61;;1222:430;824:14;817:21;492:353;-1:-1:-1;;492:353:105:o;14:248:254:-;82:6;90;143:2;131:9;122:7;118:23;114:32;111:52;;;159:1;156;149:12;111:52;-1:-1:-1;;182:23:254;;;252:2;237:18;;;224:32;;-1:-1:-1;14:248:254:o;641:385::-;727:6;735;743;751;804:3;792:9;783:7;779:23;775:33;772:53;;;821:1;818;811:12;772:53;-1:-1:-1;;844:23:254;;;914:2;899:18;;886:32;;-1:-1:-1;965:2:254;950:18;;937:32;;1016:2;1001:18;988:32;;-1:-1:-1;641:385:254;-1:-1:-1;641:385:254:o;1284:659::-;1363:6;1371;1379;1432:2;1420:9;1411:7;1407:23;1403:32;1400:52;;;1448:1;1445;1438:12;1400:52;1484:9;1471:23;1461:33;;1545:2;1534:9;1530:18;1517:32;1568:18;1609:2;1601:6;1598:14;1595:34;;;1625:1;1622;1615:12;1595:34;1663:6;1652:9;1648:22;1638:32;;1708:7;1701:4;1697:2;1693:13;1689:27;1679:55;;1730:1;1727;1720:12;1679:55;1770:2;1757:16;1796:2;1788:6;1785:14;1782:34;;;1812:1;1809;1802:12;1782:34;1857:7;1852:2;1843:6;1839:2;1835:15;1831:24;1828:37;1825:57;;;1878:1;1875;1868:12;1825:57;1909:2;1905;1901:11;1891:21;;1931:6;1921:16;;;;;1284:659;;;;;:::o;2338:180::-;2397:6;2450:2;2438:9;2429:7;2425:23;2421:32;2418:52;;;2466:1;2463;2456:12;2418:52;-1:-1:-1;2489:23:254;;2338:180;-1:-1:-1;2338:180:254:o;2705:184::-;2757:77;2754:1;2747:88;2854:4;2851:1;2844:15;2878:4;2875:1;2868:15;2894:128;2934:3;2965:1;2961:6;2958:1;2955:13;2952:39;;;2971:18;;:::i;:::-;-1:-1:-1;3007:9:254;;2894:128::o;3376:125::-;3416:4;3444:1;3441;3438:8;3435:34;;;3449:18;;:::i;:::-;-1:-1:-1;3486:9:254;;3376:125::o"

func init() {
	if err := json.Unmarshal([]byte(PreimageOracleStorageLayoutJSON), PreimageOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["PreimageOracle"] = PreimageOracleStorageLayout
	deployedBytecodes["PreimageOracle"] = PreimageOracleDeployedBin
}
