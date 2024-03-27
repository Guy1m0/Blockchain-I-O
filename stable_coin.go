// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stable_coin

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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"grantAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"increase\",\"type\":\"bool\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"revokeAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610e8f380380610e8f83398101604081905261002f91610134565b3360009081526020818152604091829020805460ff1916600190811790915582518084018452601481527f4d756c74692d44616920537461626c65636f696e0000000000000000000000009083015282518084018452908152603160f81b9082015281517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818301527f2d9ad88db3c39aaff1ff69973cba1fb0aa876e4f26308fb2735f327ab76a29b2818401527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6606082015260808101939093523060a0808501919091528251808503909101815260c0909301909152815191012060045561014d565b60006020828403121561014657600080fd5b5051919050565b610d338061015c6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c806354fd4d50116100a25780639dc29fac116100715780639dc29fac146102a9578063a9059cbb146102bc578063c634b78e146102cf578063ce50dfc9146102e2578063dd62ed3e146102f557600080fd5b806354fd4d501461023357806370a082311461025357806395d89b41146102735780639a19c7b01461029657600080fd5b806330adf81f116100e957806330adf81f146101b1578063313ce567146101d85780633644e515146101f257806340c10f19146101fb578063429b62e51461021057600080fd5b806306fdde031461011b578063095ea7b31461016457806318160ddd1461018757806323b872dd1461019e575b600080fd5b61014e6040518060400160405280601481526020017326bab63a3496a230b49029ba30b13632b1b7b4b760611b81525081565b60405161015b9190610b27565b60405180910390f35b610177610172366004610b91565b610320565b604051901515815260200161015b565b61019060015481565b60405190815260200161015b565b6101776101ac366004610bbb565b61038d565b6101907fea2aa0a1be11a07ed86d755c93467f4f82362b452371d1ba94d1715123511acb81565b6101e0601281565b60405160ff909116815260200161015b565b61019060045481565b61020e610209366004610b91565b6105c0565b005b61017761021e366004610bf7565b60006020819052908152604090205460ff1681565b61014e604051806040016040528060018152602001603160f81b81525081565b610190610261366004610bf7565b60026020526000908152604090205481565b61014e604051806040016040528060048152602001634d44414960e01b81525081565b61020e6102a4366004610bf7565b610681565b61020e6102b7366004610b91565b6106d1565b6101776102ca366004610b91565b6108d1565b61020e6102dd366004610bf7565b6108e5565b61020e6102f0366004610c12565b610938565b610190610303366004610c66565b600360209081526000928352604080842090915290825290205481565b3360008181526003602090815260408083206001600160a01b038716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259061037b9086815260200190565b60405180910390a35060015b92915050565b6001600160a01b0383166000908152600260205260408120548211156103fa5760405162461bcd60e51b815260206004820152601e60248201527f4d756c74692d4461692f696e73756666696369656e742d62616c616e6365000060448201526064015b60405180910390fd5b6001600160a01b038416331480159061043d57506001600160a01b03841660009081526003602090815260408083203384529091529020546001600160801b0314155b15610508576001600160a01b03841660009081526003602090815260408083203384529091529020548211156104b55760405162461bcd60e51b815260206004820181905260248201527f4d756c74692d4461692f696e73756666696369656e742d616c6c6f77616e636560448201526064016103f1565b6001600160a01b03841660009081526003602090815260408083203384529091529020546104e39083610af1565b6001600160a01b03851660009081526003602090815260408083203384529091529020555b6001600160a01b03841660009081526002602052604090205461052b9083610af1565b6001600160a01b03808616600090815260026020526040808220939093559085168152205461055a9083610b0c565b6001600160a01b0380851660008181526002602052604090819020939093559151908616907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906105ae9086815260200190565b60405180910390a35060019392505050565b3360009081526020819052604090205460ff166105ef5760405162461bcd60e51b81526004016103f190610c99565b6001600160a01b0382166000908152600260205260409020546106129082610b0c565b6001600160a01b0383166000908152600260205260409020556001546106389082610b0c565b6001556040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020015b60405180910390a35050565b3360009081526020819052604090205460ff166106b05760405162461bcd60e51b81526004016103f190610c99565b6001600160a01b03166000908152602081905260409020805460ff19169055565b6001600160a01b0382166000908152600260205260409020548111156107395760405162461bcd60e51b815260206004820152601e60248201527f4d756c74692d4461692f696e73756666696369656e742d62616c616e6365000060448201526064016103f1565b6001600160a01b038216331480159061077c57506001600160a01b03821660009081526003602090815260408083203384529091529020546001600160801b0314155b15610847576001600160a01b03821660009081526003602090815260408083203384529091529020548111156107f45760405162461bcd60e51b815260206004820181905260248201527f4d756c74692d4461692f696e73756666696369656e742d616c6c6f77616e636560448201526064016103f1565b6001600160a01b03821660009081526003602090815260408083203384529091529020546108229082610af1565b6001600160a01b03831660009081526003602090815260408083203384529091529020555b6001600160a01b03821660009081526002602052604090205461086a9082610af1565b6001600160a01b0383166000908152600260205260409020556001546108909082610af1565b6001556040518181526000906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610675565b60006108de33848461038d565b9392505050565b3360009081526020819052604090205460ff166109145760405162461bcd60e51b81526004016103f190610c99565b6001600160a01b03166000908152602081905260409020805460ff19166001179055565b6001600160a01b03841661098e5760405162461bcd60e51b815260206004820152601b60248201527f4d756c74692d4461692f696e76616c69642d616464726573732d30000000000060448201526064016103f1565b6001600160a01b03841633146109e65760405162461bcd60e51b815260206004820152601860248201527f4d756c74692d4461692f696e76616c69642d7065726d6974000000000000000060448201526064016103f1565b8015610a47576001600160a01b03808516600090815260036020908152604080832093871683529290522054610a1c9083610b0c565b6001600160a01b03808616600090815260036020908152604080832093881683529290522055610a9e565b6001600160a01b03808516600090815260036020908152604080832093871683529290522054610a779083610af1565b6001600160a01b038086166000908152600360209081526040808320938816835292905220555b826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610ae391815260200190565b60405180910390a350505050565b600082610afe8382610cd7565b915081111561038757600080fd5b600082610b198382610cea565b915081101561038757600080fd5b600060208083528351808285015260005b81811015610b5457858101830151858201604001528201610b38565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610b8c57600080fd5b919050565b60008060408385031215610ba457600080fd5b610bad83610b75565b946020939093013593505050565b600080600060608486031215610bd057600080fd5b610bd984610b75565b9250610be760208501610b75565b9150604084013590509250925092565b600060208284031215610c0957600080fd5b6108de82610b75565b60008060008060808587031215610c2857600080fd5b610c3185610b75565b9350610c3f60208601610b75565b92506040850135915060608501358015158114610c5b57600080fd5b939692955090935050565b60008060408385031215610c7957600080fd5b610c8283610b75565b9150610c9060208401610b75565b90509250929050565b6020808252600e908201526d139bdd08185d5d1a1bdc9a5e995960921b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b8181038181111561038757610387610cc1565b8082018082111561038757610387610cc156fea2646970667358221220a053f9da94b87fdc7a5d9f6975066a004050219238989fc7fc15adfe7f6ce18264736f6c63430008120033",
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

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_EthStableCoin *EthStableCoinCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthStableCoin.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_EthStableCoin *EthStableCoinSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _EthStableCoin.Contract.DOMAINSEPARATOR(&_EthStableCoin.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_EthStableCoin *EthStableCoinCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _EthStableCoin.Contract.DOMAINSEPARATOR(&_EthStableCoin.CallOpts)
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

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EthStableCoin *EthStableCoinCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EthStableCoin.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EthStableCoin *EthStableCoinSession) Admins(arg0 common.Address) (bool, error) {
	return _EthStableCoin.Contract.Admins(&_EthStableCoin.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EthStableCoin *EthStableCoinCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _EthStableCoin.Contract.Admins(&_EthStableCoin.CallOpts, arg0)
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

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address addr) returns()
func (_EthStableCoin *EthStableCoinTransactor) GrantAdminRole(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _EthStableCoin.contract.Transact(opts, "grantAdminRole", addr)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address addr) returns()
func (_EthStableCoin *EthStableCoinSession) GrantAdminRole(addr common.Address) (*types.Transaction, error) {
	return _EthStableCoin.Contract.GrantAdminRole(&_EthStableCoin.TransactOpts, addr)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address addr) returns()
func (_EthStableCoin *EthStableCoinTransactorSession) GrantAdminRole(addr common.Address) (*types.Transaction, error) {
	return _EthStableCoin.Contract.GrantAdminRole(&_EthStableCoin.TransactOpts, addr)
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

// RevokeAdminRole is a paid mutator transaction binding the contract method 0x9a19c7b0.
//
// Solidity: function revokeAdminRole(address addr) returns()
func (_EthStableCoin *EthStableCoinTransactor) RevokeAdminRole(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _EthStableCoin.contract.Transact(opts, "revokeAdminRole", addr)
}

// RevokeAdminRole is a paid mutator transaction binding the contract method 0x9a19c7b0.
//
// Solidity: function revokeAdminRole(address addr) returns()
func (_EthStableCoin *EthStableCoinSession) RevokeAdminRole(addr common.Address) (*types.Transaction, error) {
	return _EthStableCoin.Contract.RevokeAdminRole(&_EthStableCoin.TransactOpts, addr)
}

// RevokeAdminRole is a paid mutator transaction binding the contract method 0x9a19c7b0.
//
// Solidity: function revokeAdminRole(address addr) returns()
func (_EthStableCoin *EthStableCoinTransactorSession) RevokeAdminRole(addr common.Address) (*types.Transaction, error) {
	return _EthStableCoin.Contract.RevokeAdminRole(&_EthStableCoin.TransactOpts, addr)
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
