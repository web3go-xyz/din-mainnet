// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PreimageOracleMetaData contains all meta data concerning the PreimageOracle contract.
var PreimageOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"partOffset\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"part\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"cheat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_partOffset\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_preimage\",\"type\":\"bytes\"}],\"name\":\"loadKeccak256PreimagePart\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ident\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_word\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_partOffset\",\"type\":\"uint256\"}],\"name\":\"loadLocalData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"preimageLengths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"preimagePartOk\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"preimageParts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"}],\"name\":\"readPreimage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"dat_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"datLen_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610666806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063e03110e11161005b578063e03110e11461011d578063e159261114610159578063fe4ac08e146101d2578063fef2b4ed146102015761007d565b806361238bde146100825780638542cf50146100b75780639a1f5e7f146100ee575b600080fd5b6100a56004803603604081101561009857600080fd5b508035906020013561021e565b60408051918252519081900360200190f35b6100da600480360360408110156100cd57600080fd5b508035906020013561023b565b604080519115158252519081900360200190f35b6100a56004803603608081101561010457600080fd5b508035906020810135906040810135906060013561025b565b6101406004803603604081101561013357600080fd5b5080359060200135610346565b6040805192835260208301919091528051918290030190f35b6101d06004803603604081101561016f57600080fd5b8135919081019060408101602082013564010000000081111561019157600080fd5b8201836020820111156101a357600080fd5b803590602001918460018302840111640100000000831117156101c557600080fd5b509092509050610419565b005b6101d0600480360360808110156101e857600080fd5b5080359060208101359060408101359060600135610519565b6100a56004803603602081101561021757600080fd5b5035610580565b600160209081526000928352604080842090915290825290205481565b600260209081526000928352604080842090915290825290205460ff1681565b600061026685610592565b9050826008018211806102795750602083115b156102cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806106396021913960400191505060405180910390fd5b6000602081815260c085901b82526008959095528251828252600286526040808320858452875280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845287528083209483529386528382205581815293849052922055919050565b6000828152600260209081526040808320848452909152812054819060ff166103d057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7072652d696d616765206d757374206578697374000000000000000000000000604482015290519081900360640190fd5b506000838152602081815260409091205460088101848301106103f65783816008010391505b506000938452600160209081526040808620948652939052919092205492909150565b6044356000806008830186111561042f57600080fd5b60c083901b6080526088838682378087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80151908490207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02000000000000000000000000000000000000000000000000000000000000001760008181526002602090815260408083208b8452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209a83529981528982209390935590815290819052959095209190915550505050565b6000838152600260209081526040808320878452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558684528252808320968352958152858220939093559283529082905291902055565b60006020819052908152604090205481565b7f01000000000000000000000000000000000000000000000000000000000000007effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8216176105df816105e5565b92915050565b600090815233602052604090207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f0100000000000000000000000000000000000000000000000000000000000000179056fe70617274206f6666736574206d7573742062652077697468696e20626f756e6473a164736f6c6343000706000a",
}

// PreimageOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use PreimageOracleMetaData.ABI instead.
var PreimageOracleABI = PreimageOracleMetaData.ABI

// PreimageOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PreimageOracleMetaData.Bin instead.
var PreimageOracleBin = PreimageOracleMetaData.Bin

// DeployPreimageOracle deploys a new Ethereum contract, binding an instance of PreimageOracle to it.
func DeployPreimageOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PreimageOracle, error) {
	parsed, err := PreimageOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PreimageOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PreimageOracle{PreimageOracleCaller: PreimageOracleCaller{contract: contract}, PreimageOracleTransactor: PreimageOracleTransactor{contract: contract}, PreimageOracleFilterer: PreimageOracleFilterer{contract: contract}}, nil
}

// PreimageOracle is an auto generated Go binding around an Ethereum contract.
type PreimageOracle struct {
	PreimageOracleCaller     // Read-only binding to the contract
	PreimageOracleTransactor // Write-only binding to the contract
	PreimageOracleFilterer   // Log filterer for contract events
}

// PreimageOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PreimageOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreimageOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PreimageOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreimageOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PreimageOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreimageOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PreimageOracleSession struct {
	Contract     *PreimageOracle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PreimageOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PreimageOracleCallerSession struct {
	Contract *PreimageOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PreimageOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PreimageOracleTransactorSession struct {
	Contract     *PreimageOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PreimageOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PreimageOracleRaw struct {
	Contract *PreimageOracle // Generic contract binding to access the raw methods on
}

// PreimageOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PreimageOracleCallerRaw struct {
	Contract *PreimageOracleCaller // Generic read-only contract binding to access the raw methods on
}

// PreimageOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PreimageOracleTransactorRaw struct {
	Contract *PreimageOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPreimageOracle creates a new instance of PreimageOracle, bound to a specific deployed contract.
func NewPreimageOracle(address common.Address, backend bind.ContractBackend) (*PreimageOracle, error) {
	contract, err := bindPreimageOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PreimageOracle{PreimageOracleCaller: PreimageOracleCaller{contract: contract}, PreimageOracleTransactor: PreimageOracleTransactor{contract: contract}, PreimageOracleFilterer: PreimageOracleFilterer{contract: contract}}, nil
}

// NewPreimageOracleCaller creates a new read-only instance of PreimageOracle, bound to a specific deployed contract.
func NewPreimageOracleCaller(address common.Address, caller bind.ContractCaller) (*PreimageOracleCaller, error) {
	contract, err := bindPreimageOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PreimageOracleCaller{contract: contract}, nil
}

// NewPreimageOracleTransactor creates a new write-only instance of PreimageOracle, bound to a specific deployed contract.
func NewPreimageOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*PreimageOracleTransactor, error) {
	contract, err := bindPreimageOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PreimageOracleTransactor{contract: contract}, nil
}

// NewPreimageOracleFilterer creates a new log filterer instance of PreimageOracle, bound to a specific deployed contract.
func NewPreimageOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*PreimageOracleFilterer, error) {
	contract, err := bindPreimageOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PreimageOracleFilterer{contract: contract}, nil
}

// bindPreimageOracle binds a generic wrapper to an already deployed contract.
func bindPreimageOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PreimageOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreimageOracle *PreimageOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PreimageOracle.Contract.PreimageOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreimageOracle *PreimageOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreimageOracle.Contract.PreimageOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreimageOracle *PreimageOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreimageOracle.Contract.PreimageOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreimageOracle *PreimageOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PreimageOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreimageOracle *PreimageOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreimageOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreimageOracle *PreimageOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreimageOracle.Contract.contract.Transact(opts, method, params...)
}

// PreimageLengths is a free data retrieval call binding the contract method 0xfef2b4ed.
//
// Solidity: function preimageLengths(bytes32 ) view returns(uint256)
func (_PreimageOracle *PreimageOracleCaller) PreimageLengths(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PreimageOracle.contract.Call(opts, &out, "preimageLengths", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreimageLengths is a free data retrieval call binding the contract method 0xfef2b4ed.
//
// Solidity: function preimageLengths(bytes32 ) view returns(uint256)
func (_PreimageOracle *PreimageOracleSession) PreimageLengths(arg0 [32]byte) (*big.Int, error) {
	return _PreimageOracle.Contract.PreimageLengths(&_PreimageOracle.CallOpts, arg0)
}

// PreimageLengths is a free data retrieval call binding the contract method 0xfef2b4ed.
//
// Solidity: function preimageLengths(bytes32 ) view returns(uint256)
func (_PreimageOracle *PreimageOracleCallerSession) PreimageLengths(arg0 [32]byte) (*big.Int, error) {
	return _PreimageOracle.Contract.PreimageLengths(&_PreimageOracle.CallOpts, arg0)
}

// PreimagePartOk is a free data retrieval call binding the contract method 0x8542cf50.
//
// Solidity: function preimagePartOk(bytes32 , uint256 ) view returns(bool)
func (_PreimageOracle *PreimageOracleCaller) PreimagePartOk(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _PreimageOracle.contract.Call(opts, &out, "preimagePartOk", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PreimagePartOk is a free data retrieval call binding the contract method 0x8542cf50.
//
// Solidity: function preimagePartOk(bytes32 , uint256 ) view returns(bool)
func (_PreimageOracle *PreimageOracleSession) PreimagePartOk(arg0 [32]byte, arg1 *big.Int) (bool, error) {
	return _PreimageOracle.Contract.PreimagePartOk(&_PreimageOracle.CallOpts, arg0, arg1)
}

// PreimagePartOk is a free data retrieval call binding the contract method 0x8542cf50.
//
// Solidity: function preimagePartOk(bytes32 , uint256 ) view returns(bool)
func (_PreimageOracle *PreimageOracleCallerSession) PreimagePartOk(arg0 [32]byte, arg1 *big.Int) (bool, error) {
	return _PreimageOracle.Contract.PreimagePartOk(&_PreimageOracle.CallOpts, arg0, arg1)
}

// PreimageParts is a free data retrieval call binding the contract method 0x61238bde.
//
// Solidity: function preimageParts(bytes32 , uint256 ) view returns(bytes32)
func (_PreimageOracle *PreimageOracleCaller) PreimageParts(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PreimageOracle.contract.Call(opts, &out, "preimageParts", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PreimageParts is a free data retrieval call binding the contract method 0x61238bde.
//
// Solidity: function preimageParts(bytes32 , uint256 ) view returns(bytes32)
func (_PreimageOracle *PreimageOracleSession) PreimageParts(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _PreimageOracle.Contract.PreimageParts(&_PreimageOracle.CallOpts, arg0, arg1)
}

// PreimageParts is a free data retrieval call binding the contract method 0x61238bde.
//
// Solidity: function preimageParts(bytes32 , uint256 ) view returns(bytes32)
func (_PreimageOracle *PreimageOracleCallerSession) PreimageParts(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _PreimageOracle.Contract.PreimageParts(&_PreimageOracle.CallOpts, arg0, arg1)
}

// ReadPreimage is a free data retrieval call binding the contract method 0xe03110e1.
//
// Solidity: function readPreimage(bytes32 _key, uint256 _offset) view returns(bytes32 dat_, uint256 datLen_)
func (_PreimageOracle *PreimageOracleCaller) ReadPreimage(opts *bind.CallOpts, _key [32]byte, _offset *big.Int) (struct {
	Dat    [32]byte
	DatLen *big.Int
}, error) {
	var out []interface{}
	err := _PreimageOracle.contract.Call(opts, &out, "readPreimage", _key, _offset)

	outstruct := new(struct {
		Dat    [32]byte
		DatLen *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dat = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.DatLen = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReadPreimage is a free data retrieval call binding the contract method 0xe03110e1.
//
// Solidity: function readPreimage(bytes32 _key, uint256 _offset) view returns(bytes32 dat_, uint256 datLen_)
func (_PreimageOracle *PreimageOracleSession) ReadPreimage(_key [32]byte, _offset *big.Int) (struct {
	Dat    [32]byte
	DatLen *big.Int
}, error) {
	return _PreimageOracle.Contract.ReadPreimage(&_PreimageOracle.CallOpts, _key, _offset)
}

// ReadPreimage is a free data retrieval call binding the contract method 0xe03110e1.
//
// Solidity: function readPreimage(bytes32 _key, uint256 _offset) view returns(bytes32 dat_, uint256 datLen_)
func (_PreimageOracle *PreimageOracleCallerSession) ReadPreimage(_key [32]byte, _offset *big.Int) (struct {
	Dat    [32]byte
	DatLen *big.Int
}, error) {
	return _PreimageOracle.Contract.ReadPreimage(&_PreimageOracle.CallOpts, _key, _offset)
}

// Cheat is a paid mutator transaction binding the contract method 0xfe4ac08e.
//
// Solidity: function cheat(uint256 partOffset, bytes32 key, bytes32 part, uint256 size) returns()
func (_PreimageOracle *PreimageOracleTransactor) Cheat(opts *bind.TransactOpts, partOffset *big.Int, key [32]byte, part [32]byte, size *big.Int) (*types.Transaction, error) {
	return _PreimageOracle.contract.Transact(opts, "cheat", partOffset, key, part, size)
}

// Cheat is a paid mutator transaction binding the contract method 0xfe4ac08e.
//
// Solidity: function cheat(uint256 partOffset, bytes32 key, bytes32 part, uint256 size) returns()
func (_PreimageOracle *PreimageOracleSession) Cheat(partOffset *big.Int, key [32]byte, part [32]byte, size *big.Int) (*types.Transaction, error) {
	return _PreimageOracle.Contract.Cheat(&_PreimageOracle.TransactOpts, partOffset, key, part, size)
}

// Cheat is a paid mutator transaction binding the contract method 0xfe4ac08e.
//
// Solidity: function cheat(uint256 partOffset, bytes32 key, bytes32 part, uint256 size) returns()
func (_PreimageOracle *PreimageOracleTransactorSession) Cheat(partOffset *big.Int, key [32]byte, part [32]byte, size *big.Int) (*types.Transaction, error) {
	return _PreimageOracle.Contract.Cheat(&_PreimageOracle.TransactOpts, partOffset, key, part, size)
}

// LoadKeccak256PreimagePart is a paid mutator transaction binding the contract method 0xe1592611.
//
// Solidity: function loadKeccak256PreimagePart(uint256 _partOffset, bytes _preimage) returns()
func (_PreimageOracle *PreimageOracleTransactor) LoadKeccak256PreimagePart(opts *bind.TransactOpts, _partOffset *big.Int, _preimage []byte) (*types.Transaction, error) {
	return _PreimageOracle.contract.Transact(opts, "loadKeccak256PreimagePart", _partOffset, _preimage)
}

// LoadKeccak256PreimagePart is a paid mutator transaction binding the contract method 0xe1592611.
//
// Solidity: function loadKeccak256PreimagePart(uint256 _partOffset, bytes _preimage) returns()
func (_PreimageOracle *PreimageOracleSession) LoadKeccak256PreimagePart(_partOffset *big.Int, _preimage []byte) (*types.Transaction, error) {
	return _PreimageOracle.Contract.LoadKeccak256PreimagePart(&_PreimageOracle.TransactOpts, _partOffset, _preimage)
}

// LoadKeccak256PreimagePart is a paid mutator transaction binding the contract method 0xe1592611.
//
// Solidity: function loadKeccak256PreimagePart(uint256 _partOffset, bytes _preimage) returns()
func (_PreimageOracle *PreimageOracleTransactorSession) LoadKeccak256PreimagePart(_partOffset *big.Int, _preimage []byte) (*types.Transaction, error) {
	return _PreimageOracle.Contract.LoadKeccak256PreimagePart(&_PreimageOracle.TransactOpts, _partOffset, _preimage)
}

// LoadLocalData is a paid mutator transaction binding the contract method 0x9a1f5e7f.
//
// Solidity: function loadLocalData(uint256 _ident, bytes32 _word, uint256 _size, uint256 _partOffset) returns(bytes32 key_)
func (_PreimageOracle *PreimageOracleTransactor) LoadLocalData(opts *bind.TransactOpts, _ident *big.Int, _word [32]byte, _size *big.Int, _partOffset *big.Int) (*types.Transaction, error) {
	return _PreimageOracle.contract.Transact(opts, "loadLocalData", _ident, _word, _size, _partOffset)
}

// LoadLocalData is a paid mutator transaction binding the contract method 0x9a1f5e7f.
//
// Solidity: function loadLocalData(uint256 _ident, bytes32 _word, uint256 _size, uint256 _partOffset) returns(bytes32 key_)
func (_PreimageOracle *PreimageOracleSession) LoadLocalData(_ident *big.Int, _word [32]byte, _size *big.Int, _partOffset *big.Int) (*types.Transaction, error) {
	return _PreimageOracle.Contract.LoadLocalData(&_PreimageOracle.TransactOpts, _ident, _word, _size, _partOffset)
}

// LoadLocalData is a paid mutator transaction binding the contract method 0x9a1f5e7f.
//
// Solidity: function loadLocalData(uint256 _ident, bytes32 _word, uint256 _size, uint256 _partOffset) returns(bytes32 key_)
func (_PreimageOracle *PreimageOracleTransactorSession) LoadLocalData(_ident *big.Int, _word [32]byte, _size *big.Int, _partOffset *big.Int) (*types.Transaction, error) {
	return _PreimageOracle.Contract.LoadLocalData(&_PreimageOracle.TransactOpts, _ident, _word, _size, _partOffset)
}
