// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const WETH9StorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/vendor/WETH9.sol:WETH9\",\"label\":\"name\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_string_storage\"},{\"astId\":1001,\"contract\":\"contracts/vendor/WETH9.sol:WETH9\",\"label\":\"symbol\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_string_storage\"},{\"astId\":1002,\"contract\":\"contracts/vendor/WETH9.sol:WETH9\",\"label\":\"decimals\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint8\"},{\"astId\":1003,\"contract\":\"contracts/vendor/WETH9.sol:WETH9\",\"label\":\"balanceOf\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1004,\"contract\":\"contracts/vendor/WETH9.sol:WETH9\",\"label\":\"allowance\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var WETH9StorageLayout = new(solc.StorageLayout)

var WETH9DeployedBin = "0x6080604052600436106100bc5760003560e01c8063313ce56711610074578063a9059cbb1161004e578063a9059cbb146102cb578063d0e30db0146100bc578063dd62ed3e14610311576100bc565b8063313ce5671461024b57806370a082311461027657806395d89b41146102b6576100bc565b806318160ddd116100a557806318160ddd146101aa57806323b872dd146101d15780632e1a7d4d14610221576100bc565b806306fdde03146100c6578063095ea7b314610150575b6100c4610359565b005b3480156100d257600080fd5b506100db6103a8565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101155781810151838201526020016100fd565b50505050905090810190601f1680156101425780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015c57600080fd5b506101966004803603604081101561017357600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610454565b604080519115158252519081900360200190f35b3480156101b657600080fd5b506101bf6104c7565b60408051918252519081900360200190f35b3480156101dd57600080fd5b50610196600480360360608110156101f457600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135811691602081013590911690604001356104cb565b34801561022d57600080fd5b506100c46004803603602081101561024457600080fd5b503561066b565b34801561025757600080fd5b50610260610700565b6040805160ff9092168252519081900360200190f35b34801561028257600080fd5b506101bf6004803603602081101561029957600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610709565b3480156102c257600080fd5b506100db61071b565b3480156102d757600080fd5b50610196600480360360408110156102ee57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610793565b34801561031d57600080fd5b506101bf6004803603604081101561033457600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160200135166107a7565b33600081815260036020908152604091829020805434908101909155825190815291517fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9281900390910190a2565b6000805460408051602060026001851615610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190941693909304601f8101849004840282018401909252818152929183018282801561044c5780601f106104215761010080835404028352916020019161044c565b820191906000526020600020905b81548152906001019060200180831161042f57829003601f168201915b505050505081565b33600081815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b4790565b73ffffffffffffffffffffffffffffffffffffffff83166000908152600360205260408120548211156104fd57600080fd5b73ffffffffffffffffffffffffffffffffffffffff84163314801590610573575073ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091529020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14155b156105ed5773ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091529020548211156105b557600080fd5b73ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091529020805483900390555b73ffffffffffffffffffffffffffffffffffffffff808516600081815260036020908152604080832080548890039055938716808352918490208054870190558351868152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35060019392505050565b3360009081526003602052604090205481111561068757600080fd5b33600081815260036020526040808220805485900390555183156108fc0291849190818181858888f193505050501580156106c6573d6000803e3d6000fd5b5060408051828152905133917f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65919081900360200190a250565b60025460ff1681565b60036020526000908152604090205481565b60018054604080516020600284861615610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190941693909304601f8101849004840282018401909252818152929183018282801561044c5780601f106104215761010080835404028352916020019161044c565b60006107a03384846104cb565b9392505050565b60046020908152600092835260408084209091529082529020548156fea265627a7a72315820e496abb80c5983b030f680d0bd88f66bf44e261bc3be070d612dd72f9f1f5e9a64736f6c63430005110032"

var WETH9DeployedSourceMap = "718:1809:0:-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1289:9;:7;:9::i;:::-;718:1809;739:40;;8:9:-1;5:2;;;30:1;27;20:12;5:2;739:40:0;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;8:100:-1;33:3;30:1;27:10;8:100;;;90:11;;;84:18;71:11;;;64:39;52:2;45:10;8:100;;;12:14;739:40:0;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1755:177;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1755:177:0;;;;;;13:2:-1;8:3;5:11;2:2;;;29:1;26;19:12;2:2;-1:-1;1755:177:0;;;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;1654:95;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1654:95:0;;;:::i;:::-;;;;;;;;;;;;;;;;2065:460;;8:9:-1;5:2;;;30:1;27;20:12;5:2;2065:460:0;;;;;;13:2:-1;8:3;5:11;2:2;;;29:1;26;19:12;2:2;-1:-1;2065:460:0;;;;;;;;;;;;;;;;;;:::i;1445:203::-;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1445:203:0;;;;;;13:2:-1;8:3;5:11;2:2;;;29:1;26;19:12;2:2;-1:-1;1445:203:0;;:::i;822:27::-;;8:9:-1;5:2;;;30:1;27;20:12;5:2;822:27:0;;;:::i;:::-;;;;;;;;;;;;;;;;;;;1108:65;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1108:65:0;;;;;;13:2:-1;8:3;5:11;2:2;;;29:1;26;19:12;2:2;-1:-1;1108:65:0;;;;:::i;785:31::-;;8:9:-1;5:2;;;30:1;27;20:12;5:2;785:31:0;;;:::i;1938:121::-;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1938:121:0;;;;;;13:2:-1;8:3;5:11;2:2;;;29:1;26;19:12;2:2;-1:-1;1938:121:0;;;;;;;;;:::i;1179:65::-;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1179:65:0;;;;;;13:2:-1;8:3;5:11;2:2;;;29:1;26;19:12;2:2;-1:-1;1179:65:0;;;;;;;;;;;:::i;1310:130::-;1364:10;1354:21;;;;:9;:21;;;;;;;;;:34;;1379:9;1354:34;;;;;;1403:30;;;;;;;;;;;;;;;;;1310:130::o;739:40::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::o;1755:177::-;1837:10;1811:4;1827:21;;;:9;:21;;;;;;;;;:26;;;;;;;;;;;:32;;;1874:30;;;;;;;1811:4;;1827:26;;1837:10;;1874:30;;;;;;;;-1:-1:-1;1921:4:0;1755:177;;;;:::o;1654:95::-;1721:21;1654:95;:::o;2065:460::-;2183:14;;;2155:4;2183:14;;;:9;:14;;;;;;:21;-1:-1:-1;2183:21:0;2175:30;;;;;;2220:17;;;2227:10;2220:17;;;;:59;;-1:-1:-1;2241:14:0;;;;;;;:9;:14;;;;;;;;2256:10;2241:26;;;;;;;;2276:2;2241:38;;2220:59;2216:179;;;2303:14;;;;;;;:9;:14;;;;;;;;2318:10;2303:26;;;;;;;;:33;-1:-1:-1;2303:33:0;2295:42;;;;;;2351:14;;;;;;;:9;:14;;;;;;;;2366:10;2351:26;;;;;;;:33;;;;;;;2216:179;2405:14;;;;;;;;:9;:14;;;;;;;;:21;;;;;;;2436:14;;;;;;;;;;:21;;;;;;2473:23;;;;;;;2436:14;;2473:23;;;;;;;;;;;-1:-1:-1;2514:4:0;2065:460;;;;;:::o;1445:203::-;1508:10;1498:21;;;;:9;:21;;;;;;:28;-1:-1:-1;1498:28:0;1490:37;;;;;;1547:10;1537:21;;;;:9;:21;;;;;;:28;;;;;;;1575:24;;;;;;1562:3;;1575:24;;1537:21;1575:24;1562:3;1547:10;1575:24;;;;;;;;8:9:-1;5:2;;;45:16;42:1;39;24:38;77:16;74:1;67:27;5:2;-1:-1;1614:27:0;;;;;;;;1625:10;;1614:27;;;;;;;;;;1445:203;:::o;822:27::-;;;;;;:::o;1108:65::-;;;;;;;;;;;;;:::o;785:31::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1938:121;1995:4;2018:34;2031:10;2043:3;2048;2018:12;:34::i;:::-;2011:41;1938:121;-1:-1:-1;;;1938:121:0:o;1179:65::-;;;;;;;;;;;;;;;;;;;;;;;;:::o"

func init() {
	if err := json.Unmarshal([]byte(WETH9StorageLayoutJSON), WETH9StorageLayout); err != nil {
		panic(err)
	}

	layouts["WETH9"] = WETH9StorageLayout
	deployedBytecodes["WETH9"] = WETH9DeployedBin
}
