// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// S2citizenMetaData contains all meta data concerning the S2citizen contract.
var S2citizenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_citizenCreationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"boughtIdentityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bytesContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"changeGender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"changeGenderCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"changeSpecialMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"changeSpecialMessageCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"identityId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemCacheId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"landDeedId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"genderFemale\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"specialMessage\",\"type\":\"string\"}],\"name\":\"createCitizen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"citizenId\",\"type\":\"uint256\"}],\"name\":\"disassembleCitizen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"citizenId\",\"type\":\"uint256\"}],\"name\":\"getGenderOfTokenId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"citizenId\",\"type\":\"uint256\"}],\"name\":\"getIdentityIdOfTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"citizenId\",\"type\":\"uint256\"}],\"name\":\"getItemCacheIdOfTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"citizenId\",\"type\":\"uint256\"}],\"name\":\"getLandDeedIdOfTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"citizenId\",\"type\":\"uint256\"}],\"name\":\"getSpecialMessageOfTokenId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"identityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"itemContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"landContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintedIdentityCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newestCitizen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setBoughtIdentitiesActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setBoughtIdentityContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBytesAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"setChangeGenderCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"setChangeMessageCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setCitizenAlternateMintContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setCitizenMintActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setCitizenMintContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setFemaleActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setIdentityAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setItemContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setLandContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"setMintedIdentityCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// S2citizenABI is the input ABI used to generate the binding from.
// Deprecated: Use S2citizenMetaData.ABI instead.
var S2citizenABI = S2citizenMetaData.ABI

// S2citizen is an auto generated Go binding around an Ethereum contract.
type S2citizen struct {
	S2citizenCaller     // Read-only binding to the contract
	S2citizenTransactor // Write-only binding to the contract
	S2citizenFilterer   // Log filterer for contract events
}

// S2citizenCaller is an auto generated read-only Go binding around an Ethereum contract.
type S2citizenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// S2citizenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type S2citizenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// S2citizenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type S2citizenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// S2citizenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type S2citizenSession struct {
	Contract     *S2citizen        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// S2citizenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type S2citizenCallerSession struct {
	Contract *S2citizenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// S2citizenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type S2citizenTransactorSession struct {
	Contract     *S2citizenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// S2citizenRaw is an auto generated low-level Go binding around an Ethereum contract.
type S2citizenRaw struct {
	Contract *S2citizen // Generic contract binding to access the raw methods on
}

// S2citizenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type S2citizenCallerRaw struct {
	Contract *S2citizenCaller // Generic read-only contract binding to access the raw methods on
}

// S2citizenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type S2citizenTransactorRaw struct {
	Contract *S2citizenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewS2citizen creates a new instance of S2citizen, bound to a specific deployed contract.
func NewS2citizen(address common.Address, backend bind.ContractBackend) (*S2citizen, error) {
	contract, err := bindS2citizen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &S2citizen{S2citizenCaller: S2citizenCaller{contract: contract}, S2citizenTransactor: S2citizenTransactor{contract: contract}, S2citizenFilterer: S2citizenFilterer{contract: contract}}, nil
}

// NewS2citizenCaller creates a new read-only instance of S2citizen, bound to a specific deployed contract.
func NewS2citizenCaller(address common.Address, caller bind.ContractCaller) (*S2citizenCaller, error) {
	contract, err := bindS2citizen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &S2citizenCaller{contract: contract}, nil
}

// NewS2citizenTransactor creates a new write-only instance of S2citizen, bound to a specific deployed contract.
func NewS2citizenTransactor(address common.Address, transactor bind.ContractTransactor) (*S2citizenTransactor, error) {
	contract, err := bindS2citizen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &S2citizenTransactor{contract: contract}, nil
}

// NewS2citizenFilterer creates a new log filterer instance of S2citizen, bound to a specific deployed contract.
func NewS2citizenFilterer(address common.Address, filterer bind.ContractFilterer) (*S2citizenFilterer, error) {
	contract, err := bindS2citizen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &S2citizenFilterer{contract: contract}, nil
}

// bindS2citizen binds a generic wrapper to an already deployed contract.
func bindS2citizen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(S2citizenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_S2citizen *S2citizenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _S2citizen.Contract.S2citizenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_S2citizen *S2citizenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _S2citizen.Contract.S2citizenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_S2citizen *S2citizenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _S2citizen.Contract.S2citizenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_S2citizen *S2citizenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _S2citizen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_S2citizen *S2citizenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _S2citizen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_S2citizen *S2citizenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _S2citizen.Contract.contract.Transact(opts, method, params...)
}

// CitizenCreationTime is a free data retrieval call binding the contract method 0x8bb192ff.
//
// Solidity: function _citizenCreationTime(uint256 ) view returns(uint256)
func (_S2citizen *S2citizenCaller) CitizenCreationTime(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "_citizenCreationTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CitizenCreationTime is a free data retrieval call binding the contract method 0x8bb192ff.
//
// Solidity: function _citizenCreationTime(uint256 ) view returns(uint256)
func (_S2citizen *S2citizenSession) CitizenCreationTime(arg0 *big.Int) (*big.Int, error) {
	return _S2citizen.Contract.CitizenCreationTime(&_S2citizen.CallOpts, arg0)
}

// CitizenCreationTime is a free data retrieval call binding the contract method 0x8bb192ff.
//
// Solidity: function _citizenCreationTime(uint256 ) view returns(uint256)
func (_S2citizen *S2citizenCallerSession) CitizenCreationTime(arg0 *big.Int) (*big.Int, error) {
	return _S2citizen.Contract.CitizenCreationTime(&_S2citizen.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_S2citizen *S2citizenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_S2citizen *S2citizenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _S2citizen.Contract.BalanceOf(&_S2citizen.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_S2citizen *S2citizenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _S2citizen.Contract.BalanceOf(&_S2citizen.CallOpts, owner)
}

// BoughtIdentityContract is a free data retrieval call binding the contract method 0xeead4b16.
//
// Solidity: function boughtIdentityContract() view returns(address)
func (_S2citizen *S2citizenCaller) BoughtIdentityContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "boughtIdentityContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BoughtIdentityContract is a free data retrieval call binding the contract method 0xeead4b16.
//
// Solidity: function boughtIdentityContract() view returns(address)
func (_S2citizen *S2citizenSession) BoughtIdentityContract() (common.Address, error) {
	return _S2citizen.Contract.BoughtIdentityContract(&_S2citizen.CallOpts)
}

// BoughtIdentityContract is a free data retrieval call binding the contract method 0xeead4b16.
//
// Solidity: function boughtIdentityContract() view returns(address)
func (_S2citizen *S2citizenCallerSession) BoughtIdentityContract() (common.Address, error) {
	return _S2citizen.Contract.BoughtIdentityContract(&_S2citizen.CallOpts)
}

// BytesContract is a free data retrieval call binding the contract method 0xd1506be4.
//
// Solidity: function bytesContract() view returns(address)
func (_S2citizen *S2citizenCaller) BytesContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "bytesContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BytesContract is a free data retrieval call binding the contract method 0xd1506be4.
//
// Solidity: function bytesContract() view returns(address)
func (_S2citizen *S2citizenSession) BytesContract() (common.Address, error) {
	return _S2citizen.Contract.BytesContract(&_S2citizen.CallOpts)
}

// BytesContract is a free data retrieval call binding the contract method 0xd1506be4.
//
// Solidity: function bytesContract() view returns(address)
func (_S2citizen *S2citizenCallerSession) BytesContract() (common.Address, error) {
	return _S2citizen.Contract.BytesContract(&_S2citizen.CallOpts)
}

// ChangeGenderCost is a free data retrieval call binding the contract method 0x552cbf33.
//
// Solidity: function changeGenderCost() view returns(uint256)
func (_S2citizen *S2citizenCaller) ChangeGenderCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "changeGenderCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChangeGenderCost is a free data retrieval call binding the contract method 0x552cbf33.
//
// Solidity: function changeGenderCost() view returns(uint256)
func (_S2citizen *S2citizenSession) ChangeGenderCost() (*big.Int, error) {
	return _S2citizen.Contract.ChangeGenderCost(&_S2citizen.CallOpts)
}

// ChangeGenderCost is a free data retrieval call binding the contract method 0x552cbf33.
//
// Solidity: function changeGenderCost() view returns(uint256)
func (_S2citizen *S2citizenCallerSession) ChangeGenderCost() (*big.Int, error) {
	return _S2citizen.Contract.ChangeGenderCost(&_S2citizen.CallOpts)
}

// ChangeSpecialMessageCost is a free data retrieval call binding the contract method 0x1682c3c9.
//
// Solidity: function changeSpecialMessageCost() view returns(uint256)
func (_S2citizen *S2citizenCaller) ChangeSpecialMessageCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "changeSpecialMessageCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChangeSpecialMessageCost is a free data retrieval call binding the contract method 0x1682c3c9.
//
// Solidity: function changeSpecialMessageCost() view returns(uint256)
func (_S2citizen *S2citizenSession) ChangeSpecialMessageCost() (*big.Int, error) {
	return _S2citizen.Contract.ChangeSpecialMessageCost(&_S2citizen.CallOpts)
}

// ChangeSpecialMessageCost is a free data retrieval call binding the contract method 0x1682c3c9.
//
// Solidity: function changeSpecialMessageCost() view returns(uint256)
func (_S2citizen *S2citizenCallerSession) ChangeSpecialMessageCost() (*big.Int, error) {
	return _S2citizen.Contract.ChangeSpecialMessageCost(&_S2citizen.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_S2citizen *S2citizenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_S2citizen *S2citizenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _S2citizen.Contract.GetApproved(&_S2citizen.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_S2citizen *S2citizenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _S2citizen.Contract.GetApproved(&_S2citizen.CallOpts, tokenId)
}

// GetGenderOfTokenId is a free data retrieval call binding the contract method 0x6d58cb30.
//
// Solidity: function getGenderOfTokenId(uint256 citizenId) view returns(bool)
func (_S2citizen *S2citizenCaller) GetGenderOfTokenId(opts *bind.CallOpts, citizenId *big.Int) (bool, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "getGenderOfTokenId", citizenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetGenderOfTokenId is a free data retrieval call binding the contract method 0x6d58cb30.
//
// Solidity: function getGenderOfTokenId(uint256 citizenId) view returns(bool)
func (_S2citizen *S2citizenSession) GetGenderOfTokenId(citizenId *big.Int) (bool, error) {
	return _S2citizen.Contract.GetGenderOfTokenId(&_S2citizen.CallOpts, citizenId)
}

// GetGenderOfTokenId is a free data retrieval call binding the contract method 0x6d58cb30.
//
// Solidity: function getGenderOfTokenId(uint256 citizenId) view returns(bool)
func (_S2citizen *S2citizenCallerSession) GetGenderOfTokenId(citizenId *big.Int) (bool, error) {
	return _S2citizen.Contract.GetGenderOfTokenId(&_S2citizen.CallOpts, citizenId)
}

// GetIdentityIdOfTokenId is a free data retrieval call binding the contract method 0xe3df4296.
//
// Solidity: function getIdentityIdOfTokenId(uint256 citizenId) view returns(uint256)
func (_S2citizen *S2citizenCaller) GetIdentityIdOfTokenId(opts *bind.CallOpts, citizenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "getIdentityIdOfTokenId", citizenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIdentityIdOfTokenId is a free data retrieval call binding the contract method 0xe3df4296.
//
// Solidity: function getIdentityIdOfTokenId(uint256 citizenId) view returns(uint256)
func (_S2citizen *S2citizenSession) GetIdentityIdOfTokenId(citizenId *big.Int) (*big.Int, error) {
	return _S2citizen.Contract.GetIdentityIdOfTokenId(&_S2citizen.CallOpts, citizenId)
}

// GetIdentityIdOfTokenId is a free data retrieval call binding the contract method 0xe3df4296.
//
// Solidity: function getIdentityIdOfTokenId(uint256 citizenId) view returns(uint256)
func (_S2citizen *S2citizenCallerSession) GetIdentityIdOfTokenId(citizenId *big.Int) (*big.Int, error) {
	return _S2citizen.Contract.GetIdentityIdOfTokenId(&_S2citizen.CallOpts, citizenId)
}

// GetItemCacheIdOfTokenId is a free data retrieval call binding the contract method 0x66ad0752.
//
// Solidity: function getItemCacheIdOfTokenId(uint256 citizenId) view returns(uint256)
func (_S2citizen *S2citizenCaller) GetItemCacheIdOfTokenId(opts *bind.CallOpts, citizenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "getItemCacheIdOfTokenId", citizenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetItemCacheIdOfTokenId is a free data retrieval call binding the contract method 0x66ad0752.
//
// Solidity: function getItemCacheIdOfTokenId(uint256 citizenId) view returns(uint256)
func (_S2citizen *S2citizenSession) GetItemCacheIdOfTokenId(citizenId *big.Int) (*big.Int, error) {
	return _S2citizen.Contract.GetItemCacheIdOfTokenId(&_S2citizen.CallOpts, citizenId)
}

// GetItemCacheIdOfTokenId is a free data retrieval call binding the contract method 0x66ad0752.
//
// Solidity: function getItemCacheIdOfTokenId(uint256 citizenId) view returns(uint256)
func (_S2citizen *S2citizenCallerSession) GetItemCacheIdOfTokenId(citizenId *big.Int) (*big.Int, error) {
	return _S2citizen.Contract.GetItemCacheIdOfTokenId(&_S2citizen.CallOpts, citizenId)
}

// GetLandDeedIdOfTokenId is a free data retrieval call binding the contract method 0x0b633c84.
//
// Solidity: function getLandDeedIdOfTokenId(uint256 citizenId) view returns(uint256)
func (_S2citizen *S2citizenCaller) GetLandDeedIdOfTokenId(opts *bind.CallOpts, citizenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "getLandDeedIdOfTokenId", citizenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLandDeedIdOfTokenId is a free data retrieval call binding the contract method 0x0b633c84.
//
// Solidity: function getLandDeedIdOfTokenId(uint256 citizenId) view returns(uint256)
func (_S2citizen *S2citizenSession) GetLandDeedIdOfTokenId(citizenId *big.Int) (*big.Int, error) {
	return _S2citizen.Contract.GetLandDeedIdOfTokenId(&_S2citizen.CallOpts, citizenId)
}

// GetLandDeedIdOfTokenId is a free data retrieval call binding the contract method 0x0b633c84.
//
// Solidity: function getLandDeedIdOfTokenId(uint256 citizenId) view returns(uint256)
func (_S2citizen *S2citizenCallerSession) GetLandDeedIdOfTokenId(citizenId *big.Int) (*big.Int, error) {
	return _S2citizen.Contract.GetLandDeedIdOfTokenId(&_S2citizen.CallOpts, citizenId)
}

// GetSpecialMessageOfTokenId is a free data retrieval call binding the contract method 0x55ed458e.
//
// Solidity: function getSpecialMessageOfTokenId(uint256 citizenId) view returns(string)
func (_S2citizen *S2citizenCaller) GetSpecialMessageOfTokenId(opts *bind.CallOpts, citizenId *big.Int) (string, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "getSpecialMessageOfTokenId", citizenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSpecialMessageOfTokenId is a free data retrieval call binding the contract method 0x55ed458e.
//
// Solidity: function getSpecialMessageOfTokenId(uint256 citizenId) view returns(string)
func (_S2citizen *S2citizenSession) GetSpecialMessageOfTokenId(citizenId *big.Int) (string, error) {
	return _S2citizen.Contract.GetSpecialMessageOfTokenId(&_S2citizen.CallOpts, citizenId)
}

// GetSpecialMessageOfTokenId is a free data retrieval call binding the contract method 0x55ed458e.
//
// Solidity: function getSpecialMessageOfTokenId(uint256 citizenId) view returns(string)
func (_S2citizen *S2citizenCallerSession) GetSpecialMessageOfTokenId(citizenId *big.Int) (string, error) {
	return _S2citizen.Contract.GetSpecialMessageOfTokenId(&_S2citizen.CallOpts, citizenId)
}

// IdentityContract is a free data retrieval call binding the contract method 0x67031bae.
//
// Solidity: function identityContract() view returns(address)
func (_S2citizen *S2citizenCaller) IdentityContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "identityContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityContract is a free data retrieval call binding the contract method 0x67031bae.
//
// Solidity: function identityContract() view returns(address)
func (_S2citizen *S2citizenSession) IdentityContract() (common.Address, error) {
	return _S2citizen.Contract.IdentityContract(&_S2citizen.CallOpts)
}

// IdentityContract is a free data retrieval call binding the contract method 0x67031bae.
//
// Solidity: function identityContract() view returns(address)
func (_S2citizen *S2citizenCallerSession) IdentityContract() (common.Address, error) {
	return _S2citizen.Contract.IdentityContract(&_S2citizen.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address operator) view returns(bool)
func (_S2citizen *S2citizenCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "isApprovedForAll", _owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address operator) view returns(bool)
func (_S2citizen *S2citizenSession) IsApprovedForAll(_owner common.Address, operator common.Address) (bool, error) {
	return _S2citizen.Contract.IsApprovedForAll(&_S2citizen.CallOpts, _owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address operator) view returns(bool)
func (_S2citizen *S2citizenCallerSession) IsApprovedForAll(_owner common.Address, operator common.Address) (bool, error) {
	return _S2citizen.Contract.IsApprovedForAll(&_S2citizen.CallOpts, _owner, operator)
}

// ItemContract is a free data retrieval call binding the contract method 0x8d76f940.
//
// Solidity: function itemContract() view returns(address)
func (_S2citizen *S2citizenCaller) ItemContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "itemContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ItemContract is a free data retrieval call binding the contract method 0x8d76f940.
//
// Solidity: function itemContract() view returns(address)
func (_S2citizen *S2citizenSession) ItemContract() (common.Address, error) {
	return _S2citizen.Contract.ItemContract(&_S2citizen.CallOpts)
}

// ItemContract is a free data retrieval call binding the contract method 0x8d76f940.
//
// Solidity: function itemContract() view returns(address)
func (_S2citizen *S2citizenCallerSession) ItemContract() (common.Address, error) {
	return _S2citizen.Contract.ItemContract(&_S2citizen.CallOpts)
}

// LandContract is a free data retrieval call binding the contract method 0x03d07340.
//
// Solidity: function landContract() view returns(address)
func (_S2citizen *S2citizenCaller) LandContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "landContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LandContract is a free data retrieval call binding the contract method 0x03d07340.
//
// Solidity: function landContract() view returns(address)
func (_S2citizen *S2citizenSession) LandContract() (common.Address, error) {
	return _S2citizen.Contract.LandContract(&_S2citizen.CallOpts)
}

// LandContract is a free data retrieval call binding the contract method 0x03d07340.
//
// Solidity: function landContract() view returns(address)
func (_S2citizen *S2citizenCallerSession) LandContract() (common.Address, error) {
	return _S2citizen.Contract.LandContract(&_S2citizen.CallOpts)
}

// MintedIdentityCost is a free data retrieval call binding the contract method 0x92aa8eec.
//
// Solidity: function mintedIdentityCost() view returns(uint256)
func (_S2citizen *S2citizenCaller) MintedIdentityCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "mintedIdentityCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintedIdentityCost is a free data retrieval call binding the contract method 0x92aa8eec.
//
// Solidity: function mintedIdentityCost() view returns(uint256)
func (_S2citizen *S2citizenSession) MintedIdentityCost() (*big.Int, error) {
	return _S2citizen.Contract.MintedIdentityCost(&_S2citizen.CallOpts)
}

// MintedIdentityCost is a free data retrieval call binding the contract method 0x92aa8eec.
//
// Solidity: function mintedIdentityCost() view returns(uint256)
func (_S2citizen *S2citizenCallerSession) MintedIdentityCost() (*big.Int, error) {
	return _S2citizen.Contract.MintedIdentityCost(&_S2citizen.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_S2citizen *S2citizenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_S2citizen *S2citizenSession) Name() (string, error) {
	return _S2citizen.Contract.Name(&_S2citizen.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_S2citizen *S2citizenCallerSession) Name() (string, error) {
	return _S2citizen.Contract.Name(&_S2citizen.CallOpts)
}

// NewestCitizen is a free data retrieval call binding the contract method 0xcad8a327.
//
// Solidity: function newestCitizen() view returns(uint256)
func (_S2citizen *S2citizenCaller) NewestCitizen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "newestCitizen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewestCitizen is a free data retrieval call binding the contract method 0xcad8a327.
//
// Solidity: function newestCitizen() view returns(uint256)
func (_S2citizen *S2citizenSession) NewestCitizen() (*big.Int, error) {
	return _S2citizen.Contract.NewestCitizen(&_S2citizen.CallOpts)
}

// NewestCitizen is a free data retrieval call binding the contract method 0xcad8a327.
//
// Solidity: function newestCitizen() view returns(uint256)
func (_S2citizen *S2citizenCallerSession) NewestCitizen() (*big.Int, error) {
	return _S2citizen.Contract.NewestCitizen(&_S2citizen.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_S2citizen *S2citizenCaller) OnERC721Received(opts *bind.CallOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "onERC721Received", operator, from, tokenId, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_S2citizen *S2citizenSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _S2citizen.Contract.OnERC721Received(&_S2citizen.CallOpts, operator, from, tokenId, data)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_S2citizen *S2citizenCallerSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _S2citizen.Contract.OnERC721Received(&_S2citizen.CallOpts, operator, from, tokenId, data)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_S2citizen *S2citizenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_S2citizen *S2citizenSession) Owner() (common.Address, error) {
	return _S2citizen.Contract.Owner(&_S2citizen.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_S2citizen *S2citizenCallerSession) Owner() (common.Address, error) {
	return _S2citizen.Contract.Owner(&_S2citizen.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_S2citizen *S2citizenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_S2citizen *S2citizenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _S2citizen.Contract.OwnerOf(&_S2citizen.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_S2citizen *S2citizenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _S2citizen.Contract.OwnerOf(&_S2citizen.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_S2citizen *S2citizenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_S2citizen *S2citizenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _S2citizen.Contract.SupportsInterface(&_S2citizen.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_S2citizen *S2citizenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _S2citizen.Contract.SupportsInterface(&_S2citizen.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_S2citizen *S2citizenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_S2citizen *S2citizenSession) Symbol() (string, error) {
	return _S2citizen.Contract.Symbol(&_S2citizen.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_S2citizen *S2citizenCallerSession) Symbol() (string, error) {
	return _S2citizen.Contract.Symbol(&_S2citizen.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_S2citizen *S2citizenCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_S2citizen *S2citizenSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _S2citizen.Contract.TokenByIndex(&_S2citizen.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_S2citizen *S2citizenCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _S2citizen.Contract.TokenByIndex(&_S2citizen.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_S2citizen *S2citizenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_S2citizen *S2citizenSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _S2citizen.Contract.TokenOfOwnerByIndex(&_S2citizen.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_S2citizen *S2citizenCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _S2citizen.Contract.TokenOfOwnerByIndex(&_S2citizen.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_S2citizen *S2citizenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_S2citizen *S2citizenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _S2citizen.Contract.TokenURI(&_S2citizen.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_S2citizen *S2citizenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _S2citizen.Contract.TokenURI(&_S2citizen.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_S2citizen *S2citizenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _S2citizen.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_S2citizen *S2citizenSession) TotalSupply() (*big.Int, error) {
	return _S2citizen.Contract.TotalSupply(&_S2citizen.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_S2citizen *S2citizenCallerSession) TotalSupply() (*big.Int, error) {
	return _S2citizen.Contract.TotalSupply(&_S2citizen.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_S2citizen *S2citizenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_S2citizen *S2citizenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.Approve(&_S2citizen.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_S2citizen *S2citizenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.Approve(&_S2citizen.TransactOpts, to, tokenId)
}

// ChangeGender is a paid mutator transaction binding the contract method 0x652296f1.
//
// Solidity: function changeGender(uint256 tokenId) returns()
func (_S2citizen *S2citizenTransactor) ChangeGender(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "changeGender", tokenId)
}

// ChangeGender is a paid mutator transaction binding the contract method 0x652296f1.
//
// Solidity: function changeGender(uint256 tokenId) returns()
func (_S2citizen *S2citizenSession) ChangeGender(tokenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.ChangeGender(&_S2citizen.TransactOpts, tokenId)
}

// ChangeGender is a paid mutator transaction binding the contract method 0x652296f1.
//
// Solidity: function changeGender(uint256 tokenId) returns()
func (_S2citizen *S2citizenTransactorSession) ChangeGender(tokenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.ChangeGender(&_S2citizen.TransactOpts, tokenId)
}

// ChangeSpecialMessage is a paid mutator transaction binding the contract method 0xeb4807dc.
//
// Solidity: function changeSpecialMessage(uint256 tokenId, string _message) returns()
func (_S2citizen *S2citizenTransactor) ChangeSpecialMessage(opts *bind.TransactOpts, tokenId *big.Int, _message string) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "changeSpecialMessage", tokenId, _message)
}

// ChangeSpecialMessage is a paid mutator transaction binding the contract method 0xeb4807dc.
//
// Solidity: function changeSpecialMessage(uint256 tokenId, string _message) returns()
func (_S2citizen *S2citizenSession) ChangeSpecialMessage(tokenId *big.Int, _message string) (*types.Transaction, error) {
	return _S2citizen.Contract.ChangeSpecialMessage(&_S2citizen.TransactOpts, tokenId, _message)
}

// ChangeSpecialMessage is a paid mutator transaction binding the contract method 0xeb4807dc.
//
// Solidity: function changeSpecialMessage(uint256 tokenId, string _message) returns()
func (_S2citizen *S2citizenTransactorSession) ChangeSpecialMessage(tokenId *big.Int, _message string) (*types.Transaction, error) {
	return _S2citizen.Contract.ChangeSpecialMessage(&_S2citizen.TransactOpts, tokenId, _message)
}

// CreateCitizen is a paid mutator transaction binding the contract method 0x8860b661.
//
// Solidity: function createCitizen(uint256 identityId, uint256 itemCacheId, uint256 landDeedId, bool genderFemale, string specialMessage) returns()
func (_S2citizen *S2citizenTransactor) CreateCitizen(opts *bind.TransactOpts, identityId *big.Int, itemCacheId *big.Int, landDeedId *big.Int, genderFemale bool, specialMessage string) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "createCitizen", identityId, itemCacheId, landDeedId, genderFemale, specialMessage)
}

// CreateCitizen is a paid mutator transaction binding the contract method 0x8860b661.
//
// Solidity: function createCitizen(uint256 identityId, uint256 itemCacheId, uint256 landDeedId, bool genderFemale, string specialMessage) returns()
func (_S2citizen *S2citizenSession) CreateCitizen(identityId *big.Int, itemCacheId *big.Int, landDeedId *big.Int, genderFemale bool, specialMessage string) (*types.Transaction, error) {
	return _S2citizen.Contract.CreateCitizen(&_S2citizen.TransactOpts, identityId, itemCacheId, landDeedId, genderFemale, specialMessage)
}

// CreateCitizen is a paid mutator transaction binding the contract method 0x8860b661.
//
// Solidity: function createCitizen(uint256 identityId, uint256 itemCacheId, uint256 landDeedId, bool genderFemale, string specialMessage) returns()
func (_S2citizen *S2citizenTransactorSession) CreateCitizen(identityId *big.Int, itemCacheId *big.Int, landDeedId *big.Int, genderFemale bool, specialMessage string) (*types.Transaction, error) {
	return _S2citizen.Contract.CreateCitizen(&_S2citizen.TransactOpts, identityId, itemCacheId, landDeedId, genderFemale, specialMessage)
}

// DisassembleCitizen is a paid mutator transaction binding the contract method 0x695349b8.
//
// Solidity: function disassembleCitizen(uint256 citizenId) returns()
func (_S2citizen *S2citizenTransactor) DisassembleCitizen(opts *bind.TransactOpts, citizenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "disassembleCitizen", citizenId)
}

// DisassembleCitizen is a paid mutator transaction binding the contract method 0x695349b8.
//
// Solidity: function disassembleCitizen(uint256 citizenId) returns()
func (_S2citizen *S2citizenSession) DisassembleCitizen(citizenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.DisassembleCitizen(&_S2citizen.TransactOpts, citizenId)
}

// DisassembleCitizen is a paid mutator transaction binding the contract method 0x695349b8.
//
// Solidity: function disassembleCitizen(uint256 citizenId) returns()
func (_S2citizen *S2citizenTransactorSession) DisassembleCitizen(citizenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.DisassembleCitizen(&_S2citizen.TransactOpts, citizenId)
}

// MintIdentity is a paid mutator transaction binding the contract method 0xa68bfa6e.
//
// Solidity: function mintIdentity() returns()
func (_S2citizen *S2citizenTransactor) MintIdentity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "mintIdentity")
}

// MintIdentity is a paid mutator transaction binding the contract method 0xa68bfa6e.
//
// Solidity: function mintIdentity() returns()
func (_S2citizen *S2citizenSession) MintIdentity() (*types.Transaction, error) {
	return _S2citizen.Contract.MintIdentity(&_S2citizen.TransactOpts)
}

// MintIdentity is a paid mutator transaction binding the contract method 0xa68bfa6e.
//
// Solidity: function mintIdentity() returns()
func (_S2citizen *S2citizenTransactorSession) MintIdentity() (*types.Transaction, error) {
	return _S2citizen.Contract.MintIdentity(&_S2citizen.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_S2citizen *S2citizenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_S2citizen *S2citizenSession) RenounceOwnership() (*types.Transaction, error) {
	return _S2citizen.Contract.RenounceOwnership(&_S2citizen.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_S2citizen *S2citizenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _S2citizen.Contract.RenounceOwnership(&_S2citizen.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_S2citizen *S2citizenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_S2citizen *S2citizenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.SafeTransferFrom(&_S2citizen.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_S2citizen *S2citizenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.SafeTransferFrom(&_S2citizen.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_S2citizen *S2citizenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_S2citizen *S2citizenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _S2citizen.Contract.SafeTransferFrom0(&_S2citizen.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_S2citizen *S2citizenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _S2citizen.Contract.SafeTransferFrom0(&_S2citizen.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_S2citizen *S2citizenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_S2citizen *S2citizenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _S2citizen.Contract.SetApprovalForAll(&_S2citizen.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_S2citizen *S2citizenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _S2citizen.Contract.SetApprovalForAll(&_S2citizen.TransactOpts, operator, approved)
}

// SetBoughtIdentitiesActive is a paid mutator transaction binding the contract method 0x5f7159bd.
//
// Solidity: function setBoughtIdentitiesActive() returns()
func (_S2citizen *S2citizenTransactor) SetBoughtIdentitiesActive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "setBoughtIdentitiesActive")
}

// SetBoughtIdentitiesActive is a paid mutator transaction binding the contract method 0x5f7159bd.
//
// Solidity: function setBoughtIdentitiesActive() returns()
func (_S2citizen *S2citizenSession) SetBoughtIdentitiesActive() (*types.Transaction, error) {
	return _S2citizen.Contract.SetBoughtIdentitiesActive(&_S2citizen.TransactOpts)
}

// SetBoughtIdentitiesActive is a paid mutator transaction binding the contract method 0x5f7159bd.
//
// Solidity: function setBoughtIdentitiesActive() returns()
func (_S2citizen *S2citizenTransactorSession) SetBoughtIdentitiesActive() (*types.Transaction, error) {
	return _S2citizen.Contract.SetBoughtIdentitiesActive(&_S2citizen.TransactOpts)
}

// SetBoughtIdentityContract is a paid mutator transaction binding the contract method 0xa1e4aff2.
//
// Solidity: function setBoughtIdentityContract(address _address) returns()
func (_S2citizen *S2citizenTransactor) SetBoughtIdentityContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "setBoughtIdentityContract", _address)
}

// SetBoughtIdentityContract is a paid mutator transaction binding the contract method 0xa1e4aff2.
//
// Solidity: function setBoughtIdentityContract(address _address) returns()
func (_S2citizen *S2citizenSession) SetBoughtIdentityContract(_address common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.SetBoughtIdentityContract(&_S2citizen.TransactOpts, _address)
}

// SetBoughtIdentityContract is a paid mutator transaction binding the contract method 0xa1e4aff2.
//
// Solidity: function setBoughtIdentityContract(address _address) returns()
func (_S2citizen *S2citizenTransactorSession) SetBoughtIdentityContract(_address common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.SetBoughtIdentityContract(&_S2citizen.TransactOpts, _address)
}

// SetBytesAddress is a paid mutator transaction binding the contract method 0x140c08a5.
//
// Solidity: function setBytesAddress(address contractAddress) returns()
func (_S2citizen *S2citizenTransactor) SetBytesAddress(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "setBytesAddress", contractAddress)
}

// SetBytesAddress is a paid mutator transaction binding the contract method 0x140c08a5.
//
// Solidity: function setBytesAddress(address contractAddress) returns()
func (_S2citizen *S2citizenSession) SetBytesAddress(contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.SetBytesAddress(&_S2citizen.TransactOpts, contractAddress)
}

// SetBytesAddress is a paid mutator transaction binding the contract method 0x140c08a5.
//
// Solidity: function setBytesAddress(address contractAddress) returns()
func (_S2citizen *S2citizenTransactorSession) SetBytesAddress(contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.SetBytesAddress(&_S2citizen.TransactOpts, contractAddress)
}

// SetChangeGenderCost is a paid mutator transaction binding the contract method 0xb79c3024.
//
// Solidity: function setChangeGenderCost(uint256 _cost) returns()
func (_S2citizen *S2citizenTransactor) SetChangeGenderCost(opts *bind.TransactOpts, _cost *big.Int) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "setChangeGenderCost", _cost)
}

// SetChangeGenderCost is a paid mutator transaction binding the contract method 0xb79c3024.
//
// Solidity: function setChangeGenderCost(uint256 _cost) returns()
func (_S2citizen *S2citizenSession) SetChangeGenderCost(_cost *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.SetChangeGenderCost(&_S2citizen.TransactOpts, _cost)
}

// SetChangeGenderCost is a paid mutator transaction binding the contract method 0xb79c3024.
//
// Solidity: function setChangeGenderCost(uint256 _cost) returns()
func (_S2citizen *S2citizenTransactorSession) SetChangeGenderCost(_cost *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.SetChangeGenderCost(&_S2citizen.TransactOpts, _cost)
}

// SetChangeMessageCost is a paid mutator transaction binding the contract method 0x9eb936b9.
//
// Solidity: function setChangeMessageCost(uint256 _cost) returns()
func (_S2citizen *S2citizenTransactor) SetChangeMessageCost(opts *bind.TransactOpts, _cost *big.Int) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "setChangeMessageCost", _cost)
}

// SetChangeMessageCost is a paid mutator transaction binding the contract method 0x9eb936b9.
//
// Solidity: function setChangeMessageCost(uint256 _cost) returns()
func (_S2citizen *S2citizenSession) SetChangeMessageCost(_cost *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.SetChangeMessageCost(&_S2citizen.TransactOpts, _cost)
}

// SetChangeMessageCost is a paid mutator transaction binding the contract method 0x9eb936b9.
//
// Solidity: function setChangeMessageCost(uint256 _cost) returns()
func (_S2citizen *S2citizenTransactorSession) SetChangeMessageCost(_cost *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.SetChangeMessageCost(&_S2citizen.TransactOpts, _cost)
}

// SetCitizenAlternateMintContract is a paid mutator transaction binding the contract method 0xbfd58ea6.
//
// Solidity: function setCitizenAlternateMintContract(address contractAddress) returns()
func (_S2citizen *S2citizenTransactor) SetCitizenAlternateMintContract(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "setCitizenAlternateMintContract", contractAddress)
}

// SetCitizenAlternateMintContract is a paid mutator transaction binding the contract method 0xbfd58ea6.
//
// Solidity: function setCitizenAlternateMintContract(address contractAddress) returns()
func (_S2citizen *S2citizenSession) SetCitizenAlternateMintContract(contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.SetCitizenAlternateMintContract(&_S2citizen.TransactOpts, contractAddress)
}

// SetCitizenAlternateMintContract is a paid mutator transaction binding the contract method 0xbfd58ea6.
//
// Solidity: function setCitizenAlternateMintContract(address contractAddress) returns()
func (_S2citizen *S2citizenTransactorSession) SetCitizenAlternateMintContract(contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.SetCitizenAlternateMintContract(&_S2citizen.TransactOpts, contractAddress)
}

// SetCitizenMintActive is a paid mutator transaction binding the contract method 0x0d68ad5d.
//
// Solidity: function setCitizenMintActive() returns()
func (_S2citizen *S2citizenTransactor) SetCitizenMintActive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "setCitizenMintActive")
}

// SetCitizenMintActive is a paid mutator transaction binding the contract method 0x0d68ad5d.
//
// Solidity: function setCitizenMintActive() returns()
func (_S2citizen *S2citizenSession) SetCitizenMintActive() (*types.Transaction, error) {
	return _S2citizen.Contract.SetCitizenMintActive(&_S2citizen.TransactOpts)
}

// SetCitizenMintActive is a paid mutator transaction binding the contract method 0x0d68ad5d.
//
// Solidity: function setCitizenMintActive() returns()
func (_S2citizen *S2citizenTransactorSession) SetCitizenMintActive() (*types.Transaction, error) {
	return _S2citizen.Contract.SetCitizenMintActive(&_S2citizen.TransactOpts)
}

// SetCitizenMintContract is a paid mutator transaction binding the contract method 0xffef6c07.
//
// Solidity: function setCitizenMintContract(address contractAddress) returns()
func (_S2citizen *S2citizenTransactor) SetCitizenMintContract(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "setCitizenMintContract", contractAddress)
}

// SetCitizenMintContract is a paid mutator transaction binding the contract method 0xffef6c07.
//
// Solidity: function setCitizenMintContract(address contractAddress) returns()
func (_S2citizen *S2citizenSession) SetCitizenMintContract(contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.SetCitizenMintContract(&_S2citizen.TransactOpts, contractAddress)
}

// SetCitizenMintContract is a paid mutator transaction binding the contract method 0xffef6c07.
//
// Solidity: function setCitizenMintContract(address contractAddress) returns()
func (_S2citizen *S2citizenTransactorSession) SetCitizenMintContract(contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.SetCitizenMintContract(&_S2citizen.TransactOpts, contractAddress)
}

// SetFemaleActive is a paid mutator transaction binding the contract method 0xdddec1ba.
//
// Solidity: function setFemaleActive() returns()
func (_S2citizen *S2citizenTransactor) SetFemaleActive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "setFemaleActive")
}

// SetFemaleActive is a paid mutator transaction binding the contract method 0xdddec1ba.
//
// Solidity: function setFemaleActive() returns()
func (_S2citizen *S2citizenSession) SetFemaleActive() (*types.Transaction, error) {
	return _S2citizen.Contract.SetFemaleActive(&_S2citizen.TransactOpts)
}

// SetFemaleActive is a paid mutator transaction binding the contract method 0xdddec1ba.
//
// Solidity: function setFemaleActive() returns()
func (_S2citizen *S2citizenTransactorSession) SetFemaleActive() (*types.Transaction, error) {
	return _S2citizen.Contract.SetFemaleActive(&_S2citizen.TransactOpts)
}

// SetIdentityAddress is a paid mutator transaction binding the contract method 0x68241af4.
//
// Solidity: function setIdentityAddress(address contractAddress) returns()
func (_S2citizen *S2citizenTransactor) SetIdentityAddress(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "setIdentityAddress", contractAddress)
}

// SetIdentityAddress is a paid mutator transaction binding the contract method 0x68241af4.
//
// Solidity: function setIdentityAddress(address contractAddress) returns()
func (_S2citizen *S2citizenSession) SetIdentityAddress(contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.SetIdentityAddress(&_S2citizen.TransactOpts, contractAddress)
}

// SetIdentityAddress is a paid mutator transaction binding the contract method 0x68241af4.
//
// Solidity: function setIdentityAddress(address contractAddress) returns()
func (_S2citizen *S2citizenTransactorSession) SetIdentityAddress(contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.SetIdentityAddress(&_S2citizen.TransactOpts, contractAddress)
}

// SetItemContract is a paid mutator transaction binding the contract method 0xa7120433.
//
// Solidity: function setItemContract(address contractAddress) returns()
func (_S2citizen *S2citizenTransactor) SetItemContract(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "setItemContract", contractAddress)
}

// SetItemContract is a paid mutator transaction binding the contract method 0xa7120433.
//
// Solidity: function setItemContract(address contractAddress) returns()
func (_S2citizen *S2citizenSession) SetItemContract(contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.SetItemContract(&_S2citizen.TransactOpts, contractAddress)
}

// SetItemContract is a paid mutator transaction binding the contract method 0xa7120433.
//
// Solidity: function setItemContract(address contractAddress) returns()
func (_S2citizen *S2citizenTransactorSession) SetItemContract(contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.SetItemContract(&_S2citizen.TransactOpts, contractAddress)
}

// SetLandContract is a paid mutator transaction binding the contract method 0xe86dd092.
//
// Solidity: function setLandContract(address contractAddress) returns()
func (_S2citizen *S2citizenTransactor) SetLandContract(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "setLandContract", contractAddress)
}

// SetLandContract is a paid mutator transaction binding the contract method 0xe86dd092.
//
// Solidity: function setLandContract(address contractAddress) returns()
func (_S2citizen *S2citizenSession) SetLandContract(contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.SetLandContract(&_S2citizen.TransactOpts, contractAddress)
}

// SetLandContract is a paid mutator transaction binding the contract method 0xe86dd092.
//
// Solidity: function setLandContract(address contractAddress) returns()
func (_S2citizen *S2citizenTransactorSession) SetLandContract(contractAddress common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.SetLandContract(&_S2citizen.TransactOpts, contractAddress)
}

// SetMintedIdentityCost is a paid mutator transaction binding the contract method 0xcbaaf86b.
//
// Solidity: function setMintedIdentityCost(uint256 _cost) returns()
func (_S2citizen *S2citizenTransactor) SetMintedIdentityCost(opts *bind.TransactOpts, _cost *big.Int) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "setMintedIdentityCost", _cost)
}

// SetMintedIdentityCost is a paid mutator transaction binding the contract method 0xcbaaf86b.
//
// Solidity: function setMintedIdentityCost(uint256 _cost) returns()
func (_S2citizen *S2citizenSession) SetMintedIdentityCost(_cost *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.SetMintedIdentityCost(&_S2citizen.TransactOpts, _cost)
}

// SetMintedIdentityCost is a paid mutator transaction binding the contract method 0xcbaaf86b.
//
// Solidity: function setMintedIdentityCost(uint256 _cost) returns()
func (_S2citizen *S2citizenTransactorSession) SetMintedIdentityCost(_cost *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.SetMintedIdentityCost(&_S2citizen.TransactOpts, _cost)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_S2citizen *S2citizenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_S2citizen *S2citizenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.TransferFrom(&_S2citizen.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_S2citizen *S2citizenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _S2citizen.Contract.TransferFrom(&_S2citizen.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_S2citizen *S2citizenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _S2citizen.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_S2citizen *S2citizenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.TransferOwnership(&_S2citizen.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_S2citizen *S2citizenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _S2citizen.Contract.TransferOwnership(&_S2citizen.TransactOpts, newOwner)
}

// S2citizenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the S2citizen contract.
type S2citizenApprovalIterator struct {
	Event *S2citizenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *S2citizenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(S2citizenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(S2citizenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *S2citizenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *S2citizenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// S2citizenApproval represents a Approval event raised by the S2citizen contract.
type S2citizenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_S2citizen *S2citizenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*S2citizenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _S2citizen.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &S2citizenApprovalIterator{contract: _S2citizen.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_S2citizen *S2citizenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *S2citizenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _S2citizen.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(S2citizenApproval)
				if err := _S2citizen.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_S2citizen *S2citizenFilterer) ParseApproval(log types.Log) (*S2citizenApproval, error) {
	event := new(S2citizenApproval)
	if err := _S2citizen.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// S2citizenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the S2citizen contract.
type S2citizenApprovalForAllIterator struct {
	Event *S2citizenApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *S2citizenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(S2citizenApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(S2citizenApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *S2citizenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *S2citizenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// S2citizenApprovalForAll represents a ApprovalForAll event raised by the S2citizen contract.
type S2citizenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_S2citizen *S2citizenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*S2citizenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _S2citizen.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &S2citizenApprovalForAllIterator{contract: _S2citizen.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_S2citizen *S2citizenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *S2citizenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _S2citizen.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(S2citizenApprovalForAll)
				if err := _S2citizen.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_S2citizen *S2citizenFilterer) ParseApprovalForAll(log types.Log) (*S2citizenApprovalForAll, error) {
	event := new(S2citizenApprovalForAll)
	if err := _S2citizen.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// S2citizenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the S2citizen contract.
type S2citizenOwnershipTransferredIterator struct {
	Event *S2citizenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *S2citizenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(S2citizenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(S2citizenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *S2citizenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *S2citizenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// S2citizenOwnershipTransferred represents a OwnershipTransferred event raised by the S2citizen contract.
type S2citizenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_S2citizen *S2citizenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*S2citizenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _S2citizen.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &S2citizenOwnershipTransferredIterator{contract: _S2citizen.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_S2citizen *S2citizenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *S2citizenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _S2citizen.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(S2citizenOwnershipTransferred)
				if err := _S2citizen.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_S2citizen *S2citizenFilterer) ParseOwnershipTransferred(log types.Log) (*S2citizenOwnershipTransferred, error) {
	event := new(S2citizenOwnershipTransferred)
	if err := _S2citizen.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// S2citizenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the S2citizen contract.
type S2citizenTransferIterator struct {
	Event *S2citizenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *S2citizenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(S2citizenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(S2citizenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *S2citizenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *S2citizenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// S2citizenTransfer represents a Transfer event raised by the S2citizen contract.
type S2citizenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_S2citizen *S2citizenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*S2citizenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _S2citizen.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &S2citizenTransferIterator{contract: _S2citizen.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_S2citizen *S2citizenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *S2citizenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _S2citizen.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(S2citizenTransfer)
				if err := _S2citizen.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_S2citizen *S2citizenFilterer) ParseTransfer(log types.Log) (*S2citizenTransfer, error) {
	event := new(S2citizenTransfer)
	if err := _S2citizen.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
