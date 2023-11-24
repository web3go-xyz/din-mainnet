// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const PreimageOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageLengths\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1001,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageParts\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\"},{\"astId\":1002,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimagePartOk\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\"}],\"types\":{\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bytes32))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bytes32)\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_mapping(t_uint256,t_bytes32)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bytes32)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bytes32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var PreimageOracleStorageLayout = new(solc.StorageLayout)

var PreimageOracleDeployedBin = "0x608060405234801561001057600080fd5b50600436106100725760003560e01c8063e03110e111610050578063e03110e114610106578063e15926111461012e578063fef2b4ed1461014357600080fd5b806352f0f3ad1461007757806361238bde1461009d5780638542cf50146100c8575b600080fd5b61008a6100853660046104df565b610163565b6040519081526020015b60405180910390f35b61008a6100ab36600461051a565b600160209081526000928352604080842090915290825290205481565b6100f66100d636600461051a565b600260209081526000928352604080842090915290825290205460ff1681565b6040519015158152602001610094565b61011961011436600461051a565b610238565b60408051928352602083019190915201610094565b61014161013c36600461053c565b610329565b005b61008a6101513660046105b8565b60006020819052908152604090205481565b600061016f8686610432565b905061017c836008610600565b8211806101895750602083115b156101c0576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602081815260c085901b82526008959095528251828252600286526040808320858452875280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558484528752808320948352938652838220558181529384905292205592915050565b6000828152600260209081526040808320848452909152812054819060ff166102c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7072652d696d616765206d757374206578697374000000000000000000000000604482015260640160405180910390fd5b50600083815260208181526040909120546102dd816008610600565b6102e8856020610600565b1061030657836102f9826008610600565b6103039190610618565b91505b506000938452600160209081526040808620948652939052919092205492909150565b604435600080600883018611156103485763fe2549876000526004601cfd5b60c083901b6080526088838682378087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80151908490207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02000000000000000000000000000000000000000000000000000000000000001760008181526002602090815260408083208b8452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209a83529981528982209390935590815290819052959095209190915550505050565b7f01000000000000000000000000000000000000000000000000000000000000007effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8316176104d8818360408051600093845233602052918152606090922091527effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01000000000000000000000000000000000000000000000000000000000000001790565b9392505050565b600080600080600060a086880312156104f757600080fd5b505083359560208501359550604085013594606081013594506080013592509050565b6000806040838503121561052d57600080fd5b50508035926020909101359150565b60008060006040848603121561055157600080fd5b83359250602084013567ffffffffffffffff8082111561057057600080fd5b818601915086601f83011261058457600080fd5b81358181111561059357600080fd5b8760208285010111156105a557600080fd5b6020830194508093505050509250925092565b6000602082840312156105ca57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115610613576106136105d1565b500190565b60008282101561062a5761062a6105d1565b50039056fea164736f6c634300080f000a"

var PreimageOracleDeployedSourceMap = "306:3911:142:-:0;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1367:1211;;;;;;:::i;:::-;;:::i;:::-;;;619:25:313;;;607:2;592:18;1367:1211:142;;;;;;;;537:68;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;680:66;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1073:14:313;;1066:22;1048:41;;1036:2;1021:18;680:66:142;908:187:313;789:536:142;;;;;;:::i;:::-;;:::i;:::-;;;;1274:25:313;;;1330:2;1315:18;;1308:34;;;;1247:18;789:536:142;1100:248:313;2620:1595:142;;;;;;:::i;:::-;;:::i;:::-;;419:50;;;;;;:::i;:::-;;;;;;;;;;;;;;;1367:1211;1560:12;1665:51;1694:6;1702:13;1665:28;:51::i;:::-;1658:58;-1:-1:-1;1810:9:142;:5;1818:1;1810:9;:::i;:::-;1796:11;:23;:37;;;;1831:2;1823:5;:10;1796:37;1792:90;;;1856:15;;;;;;;;;;;;;;1792:90;1951:12;2051:4;2044:18;;;2152:3;2148:15;;;2135:29;;2184:4;2177:19;;;;2286:18;;2376:20;;;:14;:20;;;;;;:33;;;;;;;;:40;;;;2412:4;2376:40;;;;;;2426:19;;;;;;;;:32;;;;;;;;;:39;2542:21;;;;;;;;;:29;2391:4;1367:1211;-1:-1:-1;;1367:1211:142:o;789:536::-;865:12;914:20;;;:14;:20;;;;;;;;:29;;;;;;;;;865:12;;914:29;;906:62;;;;;;;2908:2:313;906:62:142;;;2890:21:313;2947:2;2927:18;;;2920:30;2986:22;2966:18;;;2959:50;3026:18;;906:62:142;;;;;;;;-1:-1:-1;1099:14:142;1116:21;;;1087:2;1116:21;;;;;;;;1167:10;1116:21;1176:1;1167:10;:::i;:::-;1151:12;:7;1161:2;1151:12;:::i;:::-;:26;1147:87;;1216:7;1203:10;:6;1212:1;1203:10;:::i;:::-;:20;;;;:::i;:::-;1193:30;;1147:87;-1:-1:-1;1290:19:142;;;;:13;:19;;;;;;;;:28;;;;;;;;;;;;789:536;;-1:-1:-1;789:536:142:o;2620:1595::-;2916:4;2903:18;2721:12;;3045:1;3035:12;;3019:29;;3016:210;;;3120:10;3117:1;3110:21;3210:1;3204:4;3197:15;3016:210;3469:3;3465:14;;;3369:4;3453:27;3500:11;3474:4;3619:16;3500:11;3601:41;3832:29;;;3836:11;3832:29;3826:36;3884:20;;;;4031:19;4024:27;4053:11;4021:44;4084:19;;;;4062:1;4084:19;;;;;;;;:32;;;;;;;;:39;;;;4119:4;4084:39;;;;;;4133:18;;;;;;;;:31;;;;;;;;;:38;;;;4181:20;;;;;;;;;;;:27;;;;-1:-1:-1;;;;2620:1595:142:o;552:449:141:-;835:11;860:19;848:32;;832:49;965:29;832:49;980:13;1676:4;1670:11;;1533:21;1787:15;;;1828:8;1822:4;1815:22;1850:27;;;1996:4;1983:18;;;2098:17;;2003:19;1979:44;2025:11;1976:61;;1455:676;965:29;958:36;552:449;-1:-1:-1;;;552:449:141:o;14:454:313:-;109:6;117;125;133;141;194:3;182:9;173:7;169:23;165:33;162:53;;;211:1;208;201:12;162:53;-1:-1:-1;;234:23:313;;;304:2;289:18;;276:32;;-1:-1:-1;355:2:313;340:18;;327:32;;406:2;391:18;;378:32;;-1:-1:-1;457:3:313;442:19;429:33;;-1:-1:-1;14:454:313;-1:-1:-1;14:454:313:o;655:248::-;723:6;731;784:2;772:9;763:7;759:23;755:32;752:52;;;800:1;797;790:12;752:52;-1:-1:-1;;823:23:313;;;893:2;878:18;;;865:32;;-1:-1:-1;655:248:313:o;1353:659::-;1432:6;1440;1448;1501:2;1489:9;1480:7;1476:23;1472:32;1469:52;;;1517:1;1514;1507:12;1469:52;1553:9;1540:23;1530:33;;1614:2;1603:9;1599:18;1586:32;1637:18;1678:2;1670:6;1667:14;1664:34;;;1694:1;1691;1684:12;1664:34;1732:6;1721:9;1717:22;1707:32;;1777:7;1770:4;1766:2;1762:13;1758:27;1748:55;;1799:1;1796;1789:12;1748:55;1839:2;1826:16;1865:2;1857:6;1854:14;1851:34;;;1881:1;1878;1871:12;1851:34;1926:7;1921:2;1912:6;1908:2;1904:15;1900:24;1897:37;1894:57;;;1947:1;1944;1937:12;1894:57;1978:2;1974;1970:11;1960:21;;2000:6;1990:16;;;;;1353:659;;;;;:::o;2017:180::-;2076:6;2129:2;2117:9;2108:7;2104:23;2100:32;2097:52;;;2145:1;2142;2135:12;2097:52;-1:-1:-1;2168:23:313;;2017:180;-1:-1:-1;2017:180:313:o;2384:184::-;2436:77;2433:1;2426:88;2533:4;2530:1;2523:15;2557:4;2554:1;2547:15;2573:128;2613:3;2644:1;2640:6;2637:1;2634:13;2631:39;;;2650:18;;:::i;:::-;-1:-1:-1;2686:9:313;;2573:128::o;3055:125::-;3095:4;3123:1;3120;3117:8;3114:34;;;3128:18;;:::i;:::-;-1:-1:-1;3165:9:313;;3055:125::o"


var PreimageOracleImmutableReferencesJSON = ""

func init() {
	if err := json.Unmarshal([]byte(PreimageOracleStorageLayoutJSON), PreimageOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["PreimageOracle"] = PreimageOracleStorageLayout
	deployedBytecodes["PreimageOracle"] = PreimageOracleDeployedBin
	immutableReferences["PreimageOracle"] = PreimageOracleImmutableReferencesJSON
}
