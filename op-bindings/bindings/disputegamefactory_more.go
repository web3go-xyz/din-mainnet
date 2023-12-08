// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const DisputeGameFactoryStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"gameImpls\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_mapping(t_userDefinedValueType(GameType)1010,t_contract(IDisputeGame)1008)\"},{\"astId\":1006,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"_disputeGames\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_mapping(t_userDefinedValueType(Hash)1011,t_userDefinedValueType(GameId)1009)\"},{\"astId\":1007,\"contract\":\"src/dispute/DisputeGameFactory.sol:DisputeGameFactory\",\"label\":\"_disputeGameList\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_array(t_userDefinedValueType(GameId)1009)dyn_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_array(t_userDefinedValueType(GameId)1009)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"GameId[]\",\"numberOfBytes\":\"32\",\"base\":\"t_userDefinedValueType(GameId)1009\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(IDisputeGame)1008\":{\"encoding\":\"inplace\",\"label\":\"contract IDisputeGame\",\"numberOfBytes\":\"20\"},\"t_mapping(t_userDefinedValueType(GameType)1010,t_contract(IDisputeGame)1008)\":{\"encoding\":\"mapping\",\"label\":\"mapping(GameType =\u003e contract IDisputeGame)\",\"numberOfBytes\":\"32\",\"key\":\"t_userDefinedValueType(GameType)1010\",\"value\":\"t_contract(IDisputeGame)1008\"},\"t_mapping(t_userDefinedValueType(Hash)1011,t_userDefinedValueType(GameId)1009)\":{\"encoding\":\"mapping\",\"label\":\"mapping(Hash =\u003e GameId)\",\"numberOfBytes\":\"32\",\"key\":\"t_userDefinedValueType(Hash)1011\",\"value\":\"t_userDefinedValueType(GameId)1009\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"},\"t_userDefinedValueType(GameId)1009\":{\"encoding\":\"inplace\",\"label\":\"GameId\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(GameType)1010\":{\"encoding\":\"inplace\",\"label\":\"GameType\",\"numberOfBytes\":\"1\"},\"t_userDefinedValueType(Hash)1011\":{\"encoding\":\"inplace\",\"label\":\"Hash\",\"numberOfBytes\":\"32\"}}}"

var DisputeGameFactoryStorageLayout = new(solc.StorageLayout)

var DisputeGameFactoryDeployedBin = "0x608060405234801561001057600080fd5b50600436106100d45760003560e01c80638da5cb5b11610081578063c4d66de81161005b578063c4d66de8146102e8578063dfa162d3146102fb578063f2fde38b1461033157600080fd5b80638da5cb5b14610231578063bb8aa1fc1461024f578063c49d5271146102a057600080fd5b80634d1975b4116100b25780634d1975b4146101d857806354fd4d50146101e0578063715018a61461022957600080fd5b806326daafbe146100d95780633142e55e1461018b57806345583b7a146101c3575b600080fd5b6101786100e7366004610d45565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0810180517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0830180517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08086018051988652968352606087529451609f0190941683209190925291905291905290565b6040519081526020015b60405180910390f35b61019e610199366004610e2e565b610344565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610182565b6101d66101d1366004610ed7565b6105a6565b005b606754610178565b61021c6040518060400160405280600581526020017f302e302e3600000000000000000000000000000000000000000000000000000081525081565b6040516101829190610f0e565b6101d661062d565b60335473ffffffffffffffffffffffffffffffffffffffff1661019e565b61026261025d366004610f81565b610641565b6040805160ff909416845267ffffffffffffffff909216602084015273ffffffffffffffffffffffffffffffffffffffff1690820152606001610182565b6102b36102ae366004610e2e565b6106a3565b6040805173ffffffffffffffffffffffffffffffffffffffff909316835267ffffffffffffffff909116602083015201610182565b6101d66102f6366004610f9a565b61072b565b61019e610309366004610fbe565b60656020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6101d661033f366004610f9a565b6108c7565b60ff841660009081526065602052604081205473ffffffffffffffffffffffffffffffffffffffff16806103ae576040517f44265d6f00000000000000000000000000000000000000000000000000000000815260ff871660048201526024015b60405180910390fd5b6104118585856040516020016103c693929190610fd9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905273ffffffffffffffffffffffffffffffffffffffff83169061099a565b91508173ffffffffffffffffffffffffffffffffffffffff16638129fc1c6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561045b57600080fd5b505af115801561046f573d6000803e3d6000fd5b5050505060006104b6878787878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506100e792505050565b60008181526066602052604090205490915015610502576040517f014f6fe5000000000000000000000000000000000000000000000000000000008152600481018290526024016103a5565b60004260b81b60f889901b178417600083815260666020526040808220839055606780546001810182559083527f9787eeb91fe3101235e4a76063c7023ecb40f923f97916639c598592fa30d6ae0183905551919250889160ff8b169173ffffffffffffffffffffffffffffffffffffffff8816917ffad0599ff449d8d9685eadecca8cb9e00924c5fd8367c1c09469824939e1ffec9190a4505050949350505050565b6105ae610ace565b60ff821660008181526065602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8616908117909155905190917f623713f72f6e427a8044bb8b3bd6834357cf285decbaa21bcc73c1d0632c4d8491a35050565b610635610ace565b61063f6000610b4f565b565b60008060006106966067858154811061065c5761065c610ff3565b906000526020600020015460f881901c9167ffffffffffffffff60b883901c169173ffffffffffffffffffffffffffffffffffffffff1690565b9196909550909350915050565b60008060006106e9878787878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506100e792505050565b60009081526066602052604090205473ffffffffffffffffffffffffffffffffffffffff81169860b89190911c67ffffffffffffffff16975095505050505050565b600054610100900460ff161580801561074b5750600054600160ff909116105b806107655750303b158015610765575060005460ff166001145b6107f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016103a5565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561084f57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610857610bc6565b61086082610b4f565b80156108c357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6108cf610ace565b73ffffffffffffffffffffffffffffffffffffffff8116610972576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016103a5565b61097b81610b4f565b50565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60006002825101603f8101600a81036040518360581b8260e81b177f6100003d81600a3d39f3363d3d373d3d3d3d610000806035363936013d7300001781528660601b601e8201527f5af43d3d93803e603357fd5bf300000000000000000000000000000000000000603282015285519150603f8101602087015b60208410610a5257805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09093019260209182019101610a15565b517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602085900360031b1b16815260f085901b9083015282816000f0945084610abf577febfef1880000000000000000000000000000000000000000000000000000000060005260206000fd5b90910160405250909392505050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461063f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103a5565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16610c5d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103a5565b61063f600054610100900460ff16610cf7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016103a5565b61063f33610b4f565b803560ff81168114610d1157600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080600060608486031215610d5a57600080fd5b610d6384610d00565b925060208401359150604084013567ffffffffffffffff80821115610d8757600080fd5b818601915086601f830112610d9b57600080fd5b813581811115610dad57610dad610d16565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610df357610df3610d16565b81604052828152896020848701011115610e0c57600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b60008060008060608587031215610e4457600080fd5b610e4d85610d00565b935060208501359250604085013567ffffffffffffffff80821115610e7157600080fd5b818701915087601f830112610e8557600080fd5b813581811115610e9457600080fd5b886020828501011115610ea657600080fd5b95989497505060200194505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461097b57600080fd5b60008060408385031215610eea57600080fd5b610ef383610d00565b91506020830135610f0381610eb5565b809150509250929050565b600060208083528351808285015260005b81811015610f3b57858101830151858201604001528201610f1f565b81811115610f4d576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b600060208284031215610f9357600080fd5b5035919050565b600060208284031215610fac57600080fd5b8135610fb781610eb5565b9392505050565b600060208284031215610fd057600080fd5b610fb782610d00565b838152818360208301376000910160200190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(DisputeGameFactoryStorageLayoutJSON), DisputeGameFactoryStorageLayout); err != nil {
		panic(err)
	}

	layouts["DisputeGameFactory"] = DisputeGameFactoryStorageLayout
	deployedBytecodes["DisputeGameFactory"] = DisputeGameFactoryDeployedBin
	immutableReferences["DisputeGameFactory"] = false
}
