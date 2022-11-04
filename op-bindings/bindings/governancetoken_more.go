// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const GovernanceTokenStorageLayoutJSON = "{\"storage\":[{\"astId\":30233,\"contract\":\"contracts/L2/GovernanceToken.sol:GovernanceToken\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":30239,\"contract\":\"contracts/L2/GovernanceToken.sol:GovernanceToken\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":30241,\"contract\":\"contracts/L2/GovernanceToken.sol:GovernanceToken\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":30243,\"contract\":\"contracts/L2/GovernanceToken.sol:GovernanceToken\",\"label\":\"_name\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_string_storage\"},{\"astId\":30245,\"contract\":\"contracts/L2/GovernanceToken.sol:GovernanceToken\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_string_storage\"},{\"astId\":31610,\"contract\":\"contracts/L2/GovernanceToken.sol:GovernanceToken\",\"label\":\"_nonces\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_mapping(t_address,t_struct(Counter)33796_storage)\"},{\"astId\":31618,\"contract\":\"contracts/L2/GovernanceToken.sol:GovernanceToken\",\"label\":\"_PERMIT_TYPEHASH_DEPRECATED_SLOT\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_bytes32\"},{\"astId\":30951,\"contract\":\"contracts/L2/GovernanceToken.sol:GovernanceToken\",\"label\":\"_delegates\",\"offset\":0,\"slot\":\"7\",\"type\":\"t_mapping(t_address,t_address)\"},{\"astId\":30957,\"contract\":\"contracts/L2/GovernanceToken.sol:GovernanceToken\",\"label\":\"_checkpoints\",\"offset\":0,\"slot\":\"8\",\"type\":\"t_mapping(t_address,t_array(t_struct(Checkpoint)30942_storage)dyn_storage)\"},{\"astId\":30961,\"contract\":\"contracts/L2/GovernanceToken.sol:GovernanceToken\",\"label\":\"_totalSupplyCheckpoints\",\"offset\":0,\"slot\":\"9\",\"type\":\"t_array(t_struct(Checkpoint)30942_storage)dyn_storage\"},{\"astId\":29883,\"contract\":\"contracts/L2/GovernanceToken.sol:GovernanceToken\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"10\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_struct(Checkpoint)30942_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct ERC20Votes.Checkpoint[]\",\"numberOfBytes\":\"32\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_address,t_address)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e address)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_address\"},\"t_mapping(t_address,t_array(t_struct(Checkpoint)30942_storage)dyn_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e struct ERC20Votes.Checkpoint[])\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_array(t_struct(Checkpoint)30942_storage)dyn_storage\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_struct(Counter)33796_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e struct Counters.Counter)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_struct(Counter)33796_storage\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_struct(Checkpoint)30942_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ERC20Votes.Checkpoint\",\"numberOfBytes\":\"32\"},\"t_struct(Counter)33796_storage\":{\"encoding\":\"inplace\",\"label\":\"struct Counters.Counter\",\"numberOfBytes\":\"32\"},\"t_uint224\":{\"encoding\":\"inplace\",\"label\":\"uint224\",\"numberOfBytes\":\"28\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"}}}"

var GovernanceTokenStorageLayout = new(solc.StorageLayout)

var GovernanceTokenDeployedBin = "0x608060405234801561001057600080fd5b50600436106101c45760003560e01c8063715018a6116100f9578063a457c2d711610097578063d505accf11610071578063d505accf14610416578063dd62ed3e14610429578063f1127ed81461046f578063f2fde38b146104c157600080fd5b8063a457c2d7146103dd578063a9059cbb146103f0578063c3cda5201461040357600080fd5b80638da5cb5b116100d35780638da5cb5b146103915780638e539e8c146103af57806395d89b41146103c25780639ab24eb0146103ca57600080fd5b8063715018a61461036357806379cc67901461036b5780637ecebe001461037e57600080fd5b80633a46b1a811610166578063587cde1e11610140578063587cde1e146102945780635c19a95c146102f25780636fcfff451461030557806370a082311461032d57600080fd5b80633a46b1a81461025957806340c10f191461026c57806342966c681461028157600080fd5b806323b872dd116101a257806323b872dd1461021c578063313ce5671461022f5780633644e5151461023e578063395093511461024657600080fd5b806306fdde03146101c9578063095ea7b3146101e757806318160ddd1461020a575b600080fd5b6101d16104d4565b6040516101de919061249d565b60405180910390f35b6101fa6101f5366004612539565b610566565b60405190151581526020016101de565b6002545b6040519081526020016101de565b6101fa61022a366004612563565b61057e565b604051601281526020016101de565b61020e6105a2565b6101fa610254366004612539565b6105b1565b61020e610267366004612539565b6105fd565b61027f61027a366004612539565b6106a3565b005b61027f61028f36600461259f565b6106b9565b6102cd6102a23660046125b8565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101de565b61027f6103003660046125b8565b6106c6565b6103186103133660046125b8565b6106d0565b60405163ffffffff90911681526020016101de565b61020e61033b3660046125b8565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b61027f610705565b61027f610379366004612539565b610719565b61020e61038c3660046125b8565b61072e565b600a5473ffffffffffffffffffffffffffffffffffffffff166102cd565b61020e6103bd36600461259f565b610759565b6101d16107cf565b61020e6103d83660046125b8565b6107de565b6101fa6103eb366004612539565b6108a9565b6101fa6103fe366004612539565b61097a565b61027f6104113660046125e4565b610988565b61027f61042436600461263c565b610aff565b61020e6104373660046126a6565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b61048261047d3660046126d9565b610cbe565b60408051825163ffffffff1681526020928301517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1692810192909252016101de565b61027f6104cf3660046125b8565b610d64565b6060600380546104e390612719565b80601f016020809104026020016040519081016040528092919081815260200182805461050f90612719565b801561055c5780601f106105315761010080835404028352916020019161055c565b820191906000526020600020905b81548152906001019060200180831161053f57829003601f168201915b5050505050905090565b600033610574818585610e18565b5060019392505050565b60003361058c858285610fcb565b6105978585856110a2565b506001949350505050565b60006105ac61135b565b905090565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061057490829086906105f8908790612795565b610e18565b600043821061066d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4552433230566f7465733a20626c6f636b206e6f7420796574206d696e65640060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260086020526040902061069c908361148f565b9392505050565b6106ab611576565b6106b582826115f7565b5050565b6106c33382611601565b50565b6106c3338261160b565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600860205260408120546106ff906116a9565b92915050565b61070d611576565b6107176000611743565b565b610724823383610fcb565b6106b58282611601565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600560205260408120546106ff565b60004382106107c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4552433230566f7465733a20626c6f636b206e6f7420796574206d696e6564006044820152606401610664565b6106ff60098361148f565b6060600480546104e390612719565b73ffffffffffffffffffffffffffffffffffffffff811660009081526008602052604081205480156108815773ffffffffffffffffffffffffffffffffffffffff8316600090815260086020526040902061083a6001836127ad565b8154811061084a5761084a6127c4565b60009182526020909120015464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16610884565b60005b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169392505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091908381101561096d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610664565b6105978286868403610e18565b6000336105748185856110a2565b834211156109f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4552433230566f7465733a207369676e617475726520657870697265640000006044820152606401610664565b604080517fe48329057bfd03d55e49b547132e39cffd9c1820ad7b9d4c5307691425d15adf602082015273ffffffffffffffffffffffffffffffffffffffff8816918101919091526060810186905260808101859052600090610a7990610a719060a001604051602081830303815290604052805190602001206117ba565b858585611823565b9050610a848161184b565b8614610aec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4552433230566f7465733a20696e76616c6964206e6f6e6365000000000000006044820152606401610664565b610af6818861160b565b50505050505050565b83421115610b69576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e650000006044820152606401610664565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610b988c61184b565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610c00826117ba565b90506000610c1082878787611823565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610ca7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e617475726500006044820152606401610664565b610cb28a8a8a610e18565b50505050505050505050565b604080518082019091526000808252602082015273ffffffffffffffffffffffffffffffffffffffff83166000908152600860205260409020805463ffffffff8416908110610d0f57610d0f6127c4565b60009182526020918290206040805180820190915291015463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16918101919091529392505050565b610d6c611576565b73ffffffffffffffffffffffffffffffffffffffff8116610e0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610664565b6106c381611743565b73ffffffffffffffffffffffffffffffffffffffff8316610eba576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610664565b73ffffffffffffffffffffffffffffffffffffffff8216610f5d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610664565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461109c578181101561108f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610664565b61109c8484848403610e18565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316611145576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610664565b73ffffffffffffffffffffffffffffffffffffffff82166111e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610664565b73ffffffffffffffffffffffffffffffffffffffff83166000908152602081905260409020548181101561129e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610664565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152602081905260408082208585039055918516815290812080548492906112e2908490612795565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161134891815260200190565b60405180910390a361109c848484611885565b60003073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156113c157507f000000000000000000000000000000000000000000000000000000000000000046145b156113eb57507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b8154600090815b818110156114f35760006114aa8284611890565b9050848682815481106114bf576114bf6127c4565b60009182526020909120015463ffffffff1611156114df578092506114ed565b6114ea816001612795565b91505b50611496565b811561154c57846115056001846127ad565b81548110611515576115156127c4565b60009182526020909120015464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1661154f565b60005b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1695945050505050565b600a5473ffffffffffffffffffffffffffffffffffffffff163314610717576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610664565b6106b582826118ab565b6106b58282611971565b73ffffffffffffffffffffffffffffffffffffffff8281166000818152600760208181526040808420805485845282862054949093528787167fffffffffffffffffffffffff00000000000000000000000000000000000000008416811790915590519190951694919391928592917f3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f9190a461109c828483611989565b600063ffffffff82111561173f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201527f32206269747300000000000000000000000000000000000000000000000000006064820152608401610664565b5090565b600a805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006106ff6117c761135b565b836040517f19010000000000000000000000000000000000000000000000000000000000006020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b600080600061183487878787611b2e565b9150915061184181611c46565b5095945050505050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526005602052604090208054600181018255905b50919050565b505050565b611880838383611e9a565b600061189f60028484186127f3565b61069c90848416612795565b6118b58282611ed9565b6002547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1015611963576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433230566f7465733a20746f74616c20737570706c79207269736b73206f60448201527f766572666c6f77696e6720766f746573000000000000000000000000000000006064820152608401610664565b61109c60096120018361200d565b61197b82826121ef565b61109c60096123e38361200d565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156119c55750600081115b156118805773ffffffffffffffffffffffffffffffffffffffff831615611a7a5773ffffffffffffffffffffffffffffffffffffffff831660009081526008602052604081208190611a1a906123e38561200d565b915091508473ffffffffffffffffffffffffffffffffffffffff167fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a7248383604051611a6f929190918252602082015260400190565b60405180910390a250505b73ffffffffffffffffffffffffffffffffffffffff8216156118805773ffffffffffffffffffffffffffffffffffffffff821660009081526008602052604081208190611aca906120018561200d565b915091508373ffffffffffffffffffffffffffffffffffffffff167fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a7248383604051611b1f929190918252602082015260400190565b60405180910390a25050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611b655750600090506003611c3d565b8460ff16601b14158015611b7d57508460ff16601c14155b15611b8e5750600090506004611c3d565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611be2573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611c3657600060019250925050611c3d565b9150600090505b94509492505050565b6000816004811115611c5a57611c5a61282e565b03611c625750565b6001816004811115611c7657611c7661282e565b03611cdd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610664565b6002816004811115611cf157611cf161282e565b03611d58576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610664565b6003816004811115611d6c57611d6c61282e565b03611df9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610664565b6004816004811115611e0d57611e0d61282e565b036106c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610664565b73ffffffffffffffffffffffffffffffffffffffff83811660009081526007602052604080822054858416835291205461188092918216911683611989565b73ffffffffffffffffffffffffffffffffffffffff8216611f56576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610664565b8060026000828254611f689190612795565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604081208054839290611fa2908490612795565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a36106b560008383611885565b600061069c8284612795565b82546000908190801561206d57856120266001836127ad565b81548110612036576120366127c4565b60009182526020909120015464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16612070565b60005b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16925061209e83858763ffffffff16565b91506000811180156120dc575043866120b86001846127ad565b815481106120c8576120c86127c4565b60009182526020909120015463ffffffff16145b15612166576120ea826123ef565b866120f66001846127ad565b81548110612106576121066127c4565b9060005260206000200160000160046101000a8154817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff02191690837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1602179055506121e6565b85604051806040016040528061217b436116a9565b63ffffffff16815260200161218f856123ef565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90811690915282546001810184556000938452602093849020835194909301519091166401000000000263ffffffff909316929092179101555b50935093915050565b73ffffffffffffffffffffffffffffffffffffffff8216612292576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610664565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015612348576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610664565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604081208383039055600280548492906123849084906127ad565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a361188083600084611885565b600061069c82846127ad565b60007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff82111561173f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203260448201527f32342062697473000000000000000000000000000000000000000000000000006064820152608401610664565b600060208083528351808285015260005b818110156124ca578581018301518582016040015282016124ae565b818111156124dc576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461253457600080fd5b919050565b6000806040838503121561254c57600080fd5b61255583612510565b946020939093013593505050565b60008060006060848603121561257857600080fd5b61258184612510565b925061258f60208501612510565b9150604084013590509250925092565b6000602082840312156125b157600080fd5b5035919050565b6000602082840312156125ca57600080fd5b61069c82612510565b803560ff8116811461253457600080fd5b60008060008060008060c087890312156125fd57600080fd5b61260687612510565b95506020870135945060408701359350612622606088016125d3565b92506080870135915060a087013590509295509295509295565b600080600080600080600060e0888a03121561265757600080fd5b61266088612510565b965061266e60208901612510565b9550604088013594506060880135935061268a608089016125d3565b925060a0880135915060c0880135905092959891949750929550565b600080604083850312156126b957600080fd5b6126c283612510565b91506126d060208401612510565b90509250929050565b600080604083850312156126ec57600080fd5b6126f583612510565b9150602083013563ffffffff8116811461270e57600080fd5b809150509250929050565b600181811c9082168061272d57607f821691505b60208210810361187a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156127a8576127a8612766565b500190565b6000828210156127bf576127bf612766565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082612829577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(GovernanceTokenStorageLayoutJSON), GovernanceTokenStorageLayout); err != nil {
		panic(err)
	}

	layouts["GovernanceToken"] = GovernanceTokenStorageLayout
	deployedBytecodes["GovernanceToken"] = GovernanceTokenDeployedBin
}
