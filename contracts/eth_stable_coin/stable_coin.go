// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_stable_coin

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

// EthStableCoinMetaData contains all meta data concerning the EthStableCoin contract.
var EthStableCoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"increase\",\"type\":\"bool\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610c19380380610c1983398101604081905261002f91610035565b5061004e565b60006020828403121561004757600080fd5b5051919050565b610bbc8061005d6000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806354fd4d501161008c5780639dc29fac116100665780639dc29fac1461023e578063a9059cbb14610251578063ce50dfc914610264578063dd62ed3e1461027757600080fd5b806354fd4d50146101db57806370a08231146101fb57806395d89b411461021b57600080fd5b806323b872dd116100c857806323b872dd1461017257806330adf81f14610185578063313ce567146101ac57806340c10f19146101c657600080fd5b806306fdde03146100ef578063095ea7b31461013857806318160ddd1461015b575b600080fd5b6101226040518060400160405280601481526020017326bab63a3496a230b49029ba30b13632b1b7b4b760611b81525081565b60405161012f91906109d8565b60405180910390f35b61014b610146366004610a42565b6102a2565b604051901515815260200161012f565b61016460005481565b60405190815260200161012f565b61014b610180366004610a6c565b61030f565b6101647fea2aa0a1be11a07ed86d755c93467f4f82362b452371d1ba94d1715123511acb81565b6101b4601281565b60405160ff909116815260200161012f565b6101d96101d4366004610a42565b610542565b005b610122604051806040016040528060018152602001603160f81b81525081565b610164610209366004610aa8565b60016020526000908152604090205481565b610122604051806040016040528060048152602001634d44414960e01b81525081565b6101d961024c366004610a42565b6105d5565b61014b61025f366004610a42565b6107d5565b6101d9610272366004610ac3565b6107e9565b610164610285366004610b17565b600260209081526000928352604080842090915290825290205481565b3360008181526002602090815260408083206001600160a01b038716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906102fd9086815260200190565b60405180910390a35060015b92915050565b6001600160a01b03831660009081526001602052604081205482111561037c5760405162461bcd60e51b815260206004820152601e60248201527f4d756c74692d4461692f696e73756666696369656e742d62616c616e6365000060448201526064015b60405180910390fd5b6001600160a01b03841633148015906103bf57506001600160a01b03841660009081526002602090815260408083203384529091529020546001600160801b0314155b1561048a576001600160a01b03841660009081526002602090815260408083203384529091529020548211156104375760405162461bcd60e51b815260206004820181905260248201527f4d756c74692d4461692f696e73756666696369656e742d616c6c6f77616e63656044820152606401610373565b6001600160a01b038416600090815260026020908152604080832033845290915290205461046590836109a2565b6001600160a01b03851660009081526002602090815260408083203384529091529020555b6001600160a01b0384166000908152600160205260409020546104ad90836109a2565b6001600160a01b0380861660009081526001602052604080822093909355908516815220546104dc90836109bd565b6001600160a01b0380851660008181526001602052604090819020939093559151908616907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906105309086815260200190565b60405180910390a35060019392505050565b6001600160a01b03821660009081526001602052604090205461056590826109bd565b6001600160a01b0383166000908152600160205260408120919091555461058c90826109bd565b60009081556040518281526001600160a01b03841691907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020015b60405180910390a35050565b6001600160a01b03821660009081526001602052604090205481111561063d5760405162461bcd60e51b815260206004820152601e60248201527f4d756c74692d4461692f696e73756666696369656e742d62616c616e636500006044820152606401610373565b6001600160a01b038216331480159061068057506001600160a01b03821660009081526002602090815260408083203384529091529020546001600160801b0314155b1561074b576001600160a01b03821660009081526002602090815260408083203384529091529020548111156106f85760405162461bcd60e51b815260206004820181905260248201527f4d756c74692d4461692f696e73756666696369656e742d616c6c6f77616e63656044820152606401610373565b6001600160a01b038216600090815260026020908152604080832033845290915290205461072690826109a2565b6001600160a01b03831660009081526002602090815260408083203384529091529020555b6001600160a01b03821660009081526001602052604090205461076e90826109a2565b6001600160a01b0383166000908152600160205260408120919091555461079590826109a2565b60009081556040518281526001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016105c9565b60006107e233848461030f565b9392505050565b6001600160a01b03841661083f5760405162461bcd60e51b815260206004820152601b60248201527f4d756c74692d4461692f696e76616c69642d616464726573732d3000000000006044820152606401610373565b6001600160a01b03841633146108975760405162461bcd60e51b815260206004820152601860248201527f4d756c74692d4461692f696e76616c69642d7065726d697400000000000000006044820152606401610373565b80156108f8576001600160a01b038085166000908152600260209081526040808320938716835292905220546108cd90836109bd565b6001600160a01b0380861660009081526002602090815260408083209388168352929052205561094f565b6001600160a01b0380851660009081526002602090815260408083209387168352929052205461092890836109a2565b6001600160a01b038086166000908152600260209081526040808320938816835292905220555b826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161099491815260200190565b60405180910390a350505050565b6000826109af8382610b60565b915081111561030957600080fd5b6000826109ca8382610b73565b915081101561030957600080fd5b600060208083528351808285015260005b81811015610a05578581018301518582016040015282016109e9565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610a3d57600080fd5b919050565b60008060408385031215610a5557600080fd5b610a5e83610a26565b946020939093013593505050565b600080600060608486031215610a8157600080fd5b610a8a84610a26565b9250610a9860208501610a26565b9150604084013590509250925092565b600060208284031215610aba57600080fd5b6107e282610a26565b60008060008060808587031215610ad957600080fd5b610ae285610a26565b9350610af060208601610a26565b92506040850135915060608501358015158114610b0c57600080fd5b939692955090935050565b60008060408385031215610b2a57600080fd5b610b3383610a26565b9150610b4160208401610a26565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b8181038181111561030957610309610b4a565b8082018082111561030957610309610b4a56fea26469706673582212200b38b0ef4312fae3112ccef0a685f7e5abb6a419a6312db9569e4115651e691664736f6c63430008120033",
}

// EthStableCoinABI is the input ABI used to generate the binding from.
// Deprecated: Use EthStableCoinMetaData.ABI instead.
var EthStableCoinABI = EthStableCoinMetaData.ABI

// EthStableCoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthStableCoinMetaData.Bin instead.
var EthStableCoinBin = EthStableCoinMetaData.Bin

// DeployEthStableCoin deploys a new Ethereum contract, binding an instance of EthStableCoin to it.
func DeployEthStableCoin(auth *bind.TransactOpts, backend bind.ContractBackend, chainId_ *big.Int) (common.Address, *types.Transaction, *EthStableCoin, error) {
	parsed, err := EthStableCoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthStableCoinBin), backend, chainId_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthStableCoin{EthStableCoinCaller: EthStableCoinCaller{contract: contract}, EthStableCoinTransactor: EthStableCoinTransactor{contract: contract}, EthStableCoinFilterer: EthStableCoinFilterer{contract: contract}}, nil
}

// EthStableCoin is an auto generated Go binding around an Ethereum contract.
type EthStableCoin struct {
	EthStableCoinCaller     // Read-only binding to the contract
	EthStableCoinTransactor // Write-only binding to the contract
	EthStableCoinFilterer   // Log filterer for contract events
}

// EthStableCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthStableCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthStableCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthStableCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthStableCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthStableCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthStableCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthStableCoinSession struct {
	Contract     *EthStableCoin    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthStableCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthStableCoinCallerSession struct {
	Contract *EthStableCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EthStableCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthStableCoinTransactorSession struct {
	Contract     *EthStableCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EthStableCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthStableCoinRaw struct {
	Contract *EthStableCoin // Generic contract binding to access the raw methods on
}

// EthStableCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthStableCoinCallerRaw struct {
	Contract *EthStableCoinCaller // Generic read-only contract binding to access the raw methods on
}

// EthStableCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthStableCoinTransactorRaw struct {
	Contract *EthStableCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthStableCoin creates a new instance of EthStableCoin, bound to a specific deployed contract.
func NewEthStableCoin(address common.Address, backend bind.ContractBackend) (*EthStableCoin, error) {
	contract, err := bindEthStableCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthStableCoin{EthStableCoinCaller: EthStableCoinCaller{contract: contract}, EthStableCoinTransactor: EthStableCoinTransactor{contract: contract}, EthStableCoinFilterer: EthStableCoinFilterer{contract: contract}}, nil
}

// NewEthStableCoinCaller creates a new read-only instance of EthStableCoin, bound to a specific deployed contract.
func NewEthStableCoinCaller(address common.Address, caller bind.ContractCaller) (*EthStableCoinCaller, error) {
	contract, err := bindEthStableCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthStableCoinCaller{contract: contract}, nil
}

// NewEthStableCoinTransactor creates a new write-only instance of EthStableCoin, bound to a specific deployed contract.
func NewEthStableCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*EthStableCoinTransactor, error) {
	contract, err := bindEthStableCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthStableCoinTransactor{contract: contract}, nil
}

// NewEthStableCoinFilterer creates a new log filterer instance of EthStableCoin, bound to a specific deployed contract.
func NewEthStableCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*EthStableCoinFilterer, error) {
	contract, err := bindEthStableCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthStableCoinFilterer{contract: contract}, nil
}

// bindEthStableCoin binds a generic wrapper to an already deployed contract.
func bindEthStableCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthStableCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthStableCoin *EthStableCoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthStableCoin.Contract.EthStableCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthStableCoin *EthStableCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthStableCoin.Contract.EthStableCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthStableCoin *EthStableCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthStableCoin.Contract.EthStableCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthStableCoin *EthStableCoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthStableCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthStableCoin *EthStableCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthStableCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthStableCoin *EthStableCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthStableCoin.Contract.contract.Transact(opts, method, params...)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_EthStableCoin *EthStableCoinCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthStableCoin.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_EthStableCoin *EthStableCoinSession) PERMITTYPEHASH() ([32]byte, error) {
	return _EthStableCoin.Contract.PERMITTYPEHASH(&_EthStableCoin.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_EthStableCoin *EthStableCoinCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _EthStableCoin.Contract.PERMITTYPEHASH(&_EthStableCoin.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_EthStableCoin *EthStableCoinCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EthStableCoin.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_EthStableCoin *EthStableCoinSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EthStableCoin.Contract.Allowance(&_EthStableCoin.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_EthStableCoin *EthStableCoinCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EthStableCoin.Contract.Allowance(&_EthStableCoin.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_EthStableCoin *EthStableCoinCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EthStableCoin.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_EthStableCoin *EthStableCoinSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _EthStableCoin.Contract.BalanceOf(&_EthStableCoin.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_EthStableCoin *EthStableCoinCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _EthStableCoin.Contract.BalanceOf(&_EthStableCoin.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EthStableCoin *EthStableCoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EthStableCoin.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EthStableCoin *EthStableCoinSession) Decimals() (uint8, error) {
	return _EthStableCoin.Contract.Decimals(&_EthStableCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EthStableCoin *EthStableCoinCallerSession) Decimals() (uint8, error) {
	return _EthStableCoin.Contract.Decimals(&_EthStableCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EthStableCoin *EthStableCoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EthStableCoin.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EthStableCoin *EthStableCoinSession) Name() (string, error) {
	return _EthStableCoin.Contract.Name(&_EthStableCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EthStableCoin *EthStableCoinCallerSession) Name() (string, error) {
	return _EthStableCoin.Contract.Name(&_EthStableCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EthStableCoin *EthStableCoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EthStableCoin.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EthStableCoin *EthStableCoinSession) Symbol() (string, error) {
	return _EthStableCoin.Contract.Symbol(&_EthStableCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EthStableCoin *EthStableCoinCallerSession) Symbol() (string, error) {
	return _EthStableCoin.Contract.Symbol(&_EthStableCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EthStableCoin *EthStableCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthStableCoin.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EthStableCoin *EthStableCoinSession) TotalSupply() (*big.Int, error) {
	return _EthStableCoin.Contract.TotalSupply(&_EthStableCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EthStableCoin *EthStableCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _EthStableCoin.Contract.TotalSupply(&_EthStableCoin.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EthStableCoin *EthStableCoinCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EthStableCoin.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EthStableCoin *EthStableCoinSession) Version() (string, error) {
	return _EthStableCoin.Contract.Version(&_EthStableCoin.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EthStableCoin *EthStableCoinCallerSession) Version() (string, error) {
	return _EthStableCoin.Contract.Version(&_EthStableCoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_EthStableCoin *EthStableCoinTransactor) Approve(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.contract.Transact(opts, "approve", usr, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_EthStableCoin *EthStableCoinSession) Approve(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.Contract.Approve(&_EthStableCoin.TransactOpts, usr, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_EthStableCoin *EthStableCoinTransactorSession) Approve(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.Contract.Approve(&_EthStableCoin.TransactOpts, usr, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address usr, uint256 wad) returns()
func (_EthStableCoin *EthStableCoinTransactor) Burn(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.contract.Transact(opts, "burn", usr, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address usr, uint256 wad) returns()
func (_EthStableCoin *EthStableCoinSession) Burn(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.Contract.Burn(&_EthStableCoin.TransactOpts, usr, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address usr, uint256 wad) returns()
func (_EthStableCoin *EthStableCoinTransactorSession) Burn(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.Contract.Burn(&_EthStableCoin.TransactOpts, usr, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address usr, uint256 wad) returns()
func (_EthStableCoin *EthStableCoinTransactor) Mint(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.contract.Transact(opts, "mint", usr, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address usr, uint256 wad) returns()
func (_EthStableCoin *EthStableCoinSession) Mint(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.Contract.Mint(&_EthStableCoin.TransactOpts, usr, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address usr, uint256 wad) returns()
func (_EthStableCoin *EthStableCoinTransactorSession) Mint(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.Contract.Mint(&_EthStableCoin.TransactOpts, usr, wad)
}

// Permit is a paid mutator transaction binding the contract method 0xce50dfc9.
//
// Solidity: function permit(address holder, address spender, uint256 wad, bool increase) returns()
func (_EthStableCoin *EthStableCoinTransactor) Permit(opts *bind.TransactOpts, holder common.Address, spender common.Address, wad *big.Int, increase bool) (*types.Transaction, error) {
	return _EthStableCoin.contract.Transact(opts, "permit", holder, spender, wad, increase)
}

// Permit is a paid mutator transaction binding the contract method 0xce50dfc9.
//
// Solidity: function permit(address holder, address spender, uint256 wad, bool increase) returns()
func (_EthStableCoin *EthStableCoinSession) Permit(holder common.Address, spender common.Address, wad *big.Int, increase bool) (*types.Transaction, error) {
	return _EthStableCoin.Contract.Permit(&_EthStableCoin.TransactOpts, holder, spender, wad, increase)
}

// Permit is a paid mutator transaction binding the contract method 0xce50dfc9.
//
// Solidity: function permit(address holder, address spender, uint256 wad, bool increase) returns()
func (_EthStableCoin *EthStableCoinTransactorSession) Permit(holder common.Address, spender common.Address, wad *big.Int, increase bool) (*types.Transaction, error) {
	return _EthStableCoin.Contract.Permit(&_EthStableCoin.TransactOpts, holder, spender, wad, increase)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_EthStableCoin *EthStableCoinTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_EthStableCoin *EthStableCoinSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.Contract.Transfer(&_EthStableCoin.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_EthStableCoin *EthStableCoinTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.Contract.Transfer(&_EthStableCoin.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_EthStableCoin *EthStableCoinTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_EthStableCoin *EthStableCoinSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.Contract.TransferFrom(&_EthStableCoin.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_EthStableCoin *EthStableCoinTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _EthStableCoin.Contract.TransferFrom(&_EthStableCoin.TransactOpts, src, dst, wad)
}

// EthStableCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EthStableCoin contract.
type EthStableCoinApprovalIterator struct {
	Event *EthStableCoinApproval // Event containing the contract specifics and raw log

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
func (it *EthStableCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthStableCoinApproval)
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
		it.Event = new(EthStableCoinApproval)
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
func (it *EthStableCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthStableCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthStableCoinApproval represents a Approval event raised by the EthStableCoin contract.
type EthStableCoinApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_EthStableCoin *EthStableCoinFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*EthStableCoinApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _EthStableCoin.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &EthStableCoinApprovalIterator{contract: _EthStableCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_EthStableCoin *EthStableCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EthStableCoinApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _EthStableCoin.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthStableCoinApproval)
				if err := _EthStableCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_EthStableCoin *EthStableCoinFilterer) ParseApproval(log types.Log) (*EthStableCoinApproval, error) {
	event := new(EthStableCoinApproval)
	if err := _EthStableCoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthStableCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EthStableCoin contract.
type EthStableCoinTransferIterator struct {
	Event *EthStableCoinTransfer // Event containing the contract specifics and raw log

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
func (it *EthStableCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthStableCoinTransfer)
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
		it.Event = new(EthStableCoinTransfer)
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
func (it *EthStableCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthStableCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthStableCoinTransfer represents a Transfer event raised by the EthStableCoin contract.
type EthStableCoinTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_EthStableCoin *EthStableCoinFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*EthStableCoinTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _EthStableCoin.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &EthStableCoinTransferIterator{contract: _EthStableCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_EthStableCoin *EthStableCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EthStableCoinTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _EthStableCoin.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthStableCoinTransfer)
				if err := _EthStableCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_EthStableCoin *EthStableCoinFilterer) ParseTransfer(log types.Log) (*EthStableCoinTransfer, error) {
	event := new(EthStableCoinTransfer)
	if err := _EthStableCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
