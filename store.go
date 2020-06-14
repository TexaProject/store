// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"aiName\",\"type\":\"string\"}],\"name\":\"GetGlobalTexaResultURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"globalResultURL\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"aiName\",\"type\":\"string\"}],\"name\":\"GetLocalTexaResultURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"localResultURL\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"aiName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultURL\",\"type\":\"string\"}],\"name\":\"LogTexaResultURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
var StoreBin = "0x60806040523480156100115760006000fd5b50610017565b610953806100266000396000f3fe60806040523480156100115760006000fd5b50600436106100465760003560e01c80631dd155e71461004c578063621548c8146101945780637d620955146102dc57610046565b60006000fd5b61010d600480360360208110156100635760006000fd5b81019080803590602001906401000000008111156100815760006000fd5b8201836020820111156100945760006000fd5b803590602001918460018302840111640100000000831117156100b75760006000fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505090909192909091929050505061043d565b604051808060200183151515158152602001828103825284818151815260200191508051906020019080838360005b838110156101585780820151818401525b60208101905061013c565b50505050905090810190601f1680156101855780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b610255600480360360208110156101ab5760006000fd5b81019080803590602001906401000000008111156101c95760006000fd5b8201836020820111156101dc5760006000fd5b803590602001918460018302840111640100000000831117156101ff5760006000fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509090919290909192905050506105ee565b604051808060200183151515158152602001828103825284818151815260200191508051906020019080838360005b838110156102a05780820151818401525b602081019050610284565b50505050905090810190601f1680156102cd5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b61043b600480360360408110156102f35760006000fd5b81019080803590602001906401000000008111156103115760006000fd5b8201836020820111156103245760006000fd5b803590602001918460018302840111640100000000831117156103475760006000fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050909091929090919290803590602001906401000000008111156103af5760006000fd5b8201836020820111156103c25760006000fd5b803590602001918460018302840111640100000000831117156103e55760006000fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509090919290909192905050506107bf565b005b6060600060006000600050846040518082805190602001908083835b60208310151561047f57805182525b602082019150602081019050602083039250610459565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060005080546001816001161561010002031660029004905011156105ca576000600050836040518082805190602001908083835b60208310151561050c57805182525b6020820191506020810190506020830392506104e6565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060005060405160200180806108fc60229139602201828054600181600116156101000203166002900480156105ab5780601f106105895761010080835404028352918201916105ab565b820191906000526020600020905b815481529060010190602001808311610597575b50509150506040516020818303038152906040526001915091506105e9565b60006040518060200160405280600081526020015090915091506105e9565b915091565b6060600060006000600050846040518082805190602001908083835b60208310151561063057805182525b60208201915060208101905060208303925061060a565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600050805460018160011615610100020316600290049050111561079b576000600050836040518082805190602001908083835b6020831015156106bd57805182525b602082019150602081019050602083039250610697565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060005060405160200180807f687474703a2f2f6c6f63616c686f73743a333030302f232f6578706c6f72652f8152602001506020018280546001816001161561010002031660029004801561077c5780601f1061075a57610100808354040283529182019161077c565b820191906000526020600020905b815481529060010190602001808311610768575b50509150506040516020818303038152906040526001915091506107ba565b60006040518060200160405280600081526020015090915091506107ba565b915091565b806000600050836040518082805190602001908083835b6020831015156107fc57805182525b6020820191506020810190506020830392506107d6565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600050908051906020019061084592919061084b565b505b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061088c57805160ff19168380011785556108bf565b828001600101855582156108bf579182015b828111156108be578251826000509090559160200191906001019061089e565b5b5090506108cc91906108d0565b5090565b6108f891906108da565b808211156108f457600081815060009055506001016108da565b5090565b9056fe68747470733a2f2f6578706c6f72652e69706c642e696f2f232f6578706c6f72652fa2646970667358221220cc2147540b7d5938d5fc5cc8053ecc17f4363a546e24bd9bf7c8050345415d8c64736f6c634300060a0033"

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// GetGlobalTexaResultURL is a free data retrieval call binding the contract method 0x1dd155e7.
//
// Solidity: function GetGlobalTexaResultURL(string aiName) view returns(string globalResultURL, bool status)
func (_Store *StoreCaller) GetGlobalTexaResultURL(opts *bind.CallOpts, aiName string) (struct {
	GlobalResultURL string
	Status          bool
}, error) {
	ret := new(struct {
		GlobalResultURL string
		Status          bool
	})
	out := ret
	err := _Store.contract.Call(opts, out, "GetGlobalTexaResultURL", aiName)
	return *ret, err
}

// GetGlobalTexaResultURL is a free data retrieval call binding the contract method 0x1dd155e7.
//
// Solidity: function GetGlobalTexaResultURL(string aiName) view returns(string globalResultURL, bool status)
func (_Store *StoreSession) GetGlobalTexaResultURL(aiName string) (struct {
	GlobalResultURL string
	Status          bool
}, error) {
	return _Store.Contract.GetGlobalTexaResultURL(&_Store.CallOpts, aiName)
}

// GetGlobalTexaResultURL is a free data retrieval call binding the contract method 0x1dd155e7.
//
// Solidity: function GetGlobalTexaResultURL(string aiName) view returns(string globalResultURL, bool status)
func (_Store *StoreCallerSession) GetGlobalTexaResultURL(aiName string) (struct {
	GlobalResultURL string
	Status          bool
}, error) {
	return _Store.Contract.GetGlobalTexaResultURL(&_Store.CallOpts, aiName)
}

// GetLocalTexaResultURL is a free data retrieval call binding the contract method 0x621548c8.
//
// Solidity: function GetLocalTexaResultURL(string aiName) view returns(string localResultURL, bool status)
func (_Store *StoreCaller) GetLocalTexaResultURL(opts *bind.CallOpts, aiName string) (struct {
	LocalResultURL string
	Status         bool
}, error) {
	ret := new(struct {
		LocalResultURL string
		Status         bool
	})
	out := ret
	err := _Store.contract.Call(opts, out, "GetLocalTexaResultURL", aiName)
	return *ret, err
}

// GetLocalTexaResultURL is a free data retrieval call binding the contract method 0x621548c8.
//
// Solidity: function GetLocalTexaResultURL(string aiName) view returns(string localResultURL, bool status)
func (_Store *StoreSession) GetLocalTexaResultURL(aiName string) (struct {
	LocalResultURL string
	Status         bool
}, error) {
	return _Store.Contract.GetLocalTexaResultURL(&_Store.CallOpts, aiName)
}

// GetLocalTexaResultURL is a free data retrieval call binding the contract method 0x621548c8.
//
// Solidity: function GetLocalTexaResultURL(string aiName) view returns(string localResultURL, bool status)
func (_Store *StoreCallerSession) GetLocalTexaResultURL(aiName string) (struct {
	LocalResultURL string
	Status         bool
}, error) {
	return _Store.Contract.GetLocalTexaResultURL(&_Store.CallOpts, aiName)
}

// LogTexaResultURL is a paid mutator transaction binding the contract method 0x7d620955.
//
// Solidity: function LogTexaResultURL(string aiName, string resultURL) returns()
func (_Store *StoreTransactor) LogTexaResultURL(opts *bind.TransactOpts, aiName string, resultURL string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "LogTexaResultURL", aiName, resultURL)
}

// LogTexaResultURL is a paid mutator transaction binding the contract method 0x7d620955.
//
// Solidity: function LogTexaResultURL(string aiName, string resultURL) returns()
func (_Store *StoreSession) LogTexaResultURL(aiName string, resultURL string) (*types.Transaction, error) {
	return _Store.Contract.LogTexaResultURL(&_Store.TransactOpts, aiName, resultURL)
}

// LogTexaResultURL is a paid mutator transaction binding the contract method 0x7d620955.
//
// Solidity: function LogTexaResultURL(string aiName, string resultURL) returns()
func (_Store *StoreTransactorSession) LogTexaResultURL(aiName string, resultURL string) (*types.Transaction, error) {
	return _Store.Contract.LogTexaResultURL(&_Store.TransactOpts, aiName, resultURL)
}
