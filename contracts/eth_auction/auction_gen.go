// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_auction

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

// EthAuctionMetaData contains all meta data concerning the EthAuction contract.
var EthAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WaitResponse\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proceed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000fcc38038062000fcc833981016040819052620000349162000096565b600680546001600160a01b0319166001600160a01b03831617905560408051808201909152600781526672756e6e696e6760c81b60208201526004906200007c90826200016d565b5050600580546001600160a01b0319163317905562000239565b600060208284031215620000a957600080fd5b81516001600160a01b0381168114620000c157600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620000f357607f821691505b6020821081036200011457634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200016857600081815260208120601f850160051c81016020861015620001435750805b601f850160051c820191505b8181101562000164578281556001016200014f565b5050505b505050565b81516001600160401b03811115620001895762000189620000c8565b620001a1816200019a8454620000de565b846200011a565b602080601f831160018114620001d95760008415620001c05750858301515b600019600386901b1c1916600185901b17855562000164565b600085815260208120601f198616915b828110156200020a57888601518255948401946001909101908401620001e9565b5085821015620002295787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610d8380620002496000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063454a2ab311610071578063454a2ab3146100ff5780638da5cb5b1461011257806391f901571461013d578063d57bde7914610150578063fc0c546a14610167578063fe67a54b1461017a57600080fd5b806312fa6feb146100ae578063200d2ed2146100d05780632a33fec6146100e557806335a063b4146100ef5780633ccfd60b146100f7575b600080fd5b6003546100bb9060ff1681565b60405190151581526020015b60405180910390f35b6100d8610182565b6040516100c79190610ac1565b6100ed610210565b005b6100ed6103fc565b6100bb610594565b6100ed61010d366004610b0f565b6106fe565b600554610125906001600160a01b031681565b6040516001600160a01b0390911681526020016100c7565b600054610125906001600160a01b031681565b61015960015481565b6040519081526020016100c7565b600654610125906001600160a01b031681565b6100ed610944565b6004805461018f90610b28565b80601f01602080910402602001604051908101604052809291908181526020018280546101bb90610b28565b80156102085780601f106101dd57610100808354040283529160200191610208565b820191906000526020600020905b8154815290600101906020018083116101eb57829003601f168201915b505050505081565b60405165656e64696e6760d01b602082015260260160405160208183030381529060405280519060200120600460405160200161024d9190610b62565b60405160208183030381529060405280519060200120146102b55760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064015b60405180910390fd5b6005546001600160a01b03163314806102d857506000546001600160a01b031633145b61031d5760405162461bcd60e51b81526020600482015260166024820152754e6f7420617574686f72697a6564206163636573732160501b60448201526064016102ac565b604080518082019091526005815264195b99195960da1b60208201526004906103469082610c3d565b50600654600154604051630852cd8d60e31b81526001600160a01b03909216916342966c689161037c9160040190815260200190565b600060405180830381600087803b15801561039657600080fd5b505af11580156103aa573d6000803e3d6000fd5b5050600054600154604080516001600160a01b03909316835260208301919091527fdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda93500190505b60405180910390a1565b60405165656e64696e6760d01b60208201526026016040516020818303038152906040528051906020012060046040516020016104399190610b62565b604051602081830303815290604052805190602001201461049c5760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064016102ac565b6005546001600160a01b03163314806104bf57506000546001600160a01b031633145b6105045760405162461bcd60e51b81526020600482015260166024820152754e6f7420617574686f72697a6564206163636573732160501b60448201526064016102ac565b604080518082019091526005815264195b99195960da1b602082015260049061052d9082610c3d565b5060018054600080546001600160a01b031681526002602090815260408083209390935581546001600160a01b0319168255928190558151818152928301527fdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda91016103f2565b60405164195b99195960da1b60208201526000906025016040516020818303038152906040528051906020012060046040516020016105d39190610b62565b60405160208183030381529060405280519060200120146106365760405162461bcd60e51b815260206004820152601c60248201527f436f6e7472616374206e6f7420696e20454e444544207374617475730000000060448201526064016102ac565b3360009081526002602052604090205480156106f657336000818152600260205260408082209190915560065490516323b872dd60e01b81523060048201526024810192909252604482018390526001600160a01b0316906323b872dd906064016020604051808303816000875af11580156106b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106da9190610cfd565b6106f65733600090815260026020526040812091909155919050565b600191505090565b6040516672756e6e696e6760c81b602082015260270160405160208183030381529060405280519060200120600460405160200161073c9190610b62565b604051602081830303815290604052805190602001201461079f5760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e2052554e4e494e4720737461747573000060448201526064016102ac565b60015481116107f05760405162461bcd60e51b815260206004820152601e60248201527f546865726520616c7265616479206973206120686967686572206269642e000060448201526064016102ac565b6006546040516323b872dd60e01b8152336004820152306024820152604481018390526000916001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610848573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061086c9190610cfd565b9050806108b45760405162461bcd60e51b81526020600482015260166024820152752a37b5b2b7103a3930b739b332b9103330b4b632b21760511b60448201526064016102ac565b600154156108ed57600154600080546001600160a01b0316815260026020526040812080549091906108e7908490610d26565b90915550505b600080546001600160a01b03191633908117909155600183905560408051918252602082018490527ff4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300910160405180910390a15050565b6005546001600160a01b031633146109ae5760405162461bcd60e51b815260206004820152602760248201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736044820152662073746174757360c81b60648201526084016102ac565b6040516672756e6e696e6760c81b60208201526027016040516020818303038152906040528051906020012060046040516020016109ec9190610b62565b6040516020818303038152906040528051906020012014610a4f5760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e2052554e4e494e4720737461747573000060448201526064016102ac565b604080518082019091526006815265656e64696e6760d01b6020820152600490610a799082610c3d565b50600054600154604080516001600160a01b03909316835260208301919091527f0472d43abcf056976d4a0f799dccd47a3fadcc3f385763eed302bdcc802bc20c91016103f2565b600060208083528351808285015260005b81811015610aee57858101830151858201604001528201610ad2565b506000604082860101526040601f19601f8301168501019250505092915050565b600060208284031215610b2157600080fd5b5035919050565b600181811c90821680610b3c57607f821691505b602082108103610b5c57634e487b7160e01b600052602260045260246000fd5b50919050565b6000808354610b7081610b28565b60018281168015610b885760018114610b9d57610bcc565b60ff1984168752821515830287019450610bcc565b8760005260208060002060005b85811015610bc35781548a820152908401908201610baa565b50505082870194505b50929695505050505050565b634e487b7160e01b600052604160045260246000fd5b601f821115610c3857600081815260208120601f850160051c81016020861015610c155750805b601f850160051c820191505b81811015610c3457828155600101610c21565b5050505b505050565b815167ffffffffffffffff811115610c5757610c57610bd8565b610c6b81610c658454610b28565b84610bee565b602080601f831160018114610ca05760008415610c885750858301515b600019600386901b1c1916600185901b178555610c34565b600085815260208120601f198616915b82811015610ccf57888601518255948401946001909101908401610cb0565b5085821015610ced5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600060208284031215610d0f57600080fd5b81518015158114610d1f57600080fd5b9392505050565b80820180821115610d4757634e487b7160e01b600052601160045260246000fd5b9291505056fea26469706673582212207a59d4aa6eb6aacc534f626d0b54601150e21237403f6441ded47387151f643a64736f6c63430008120033",
}

// EthAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use EthAuctionMetaData.ABI instead.
var EthAuctionABI = EthAuctionMetaData.ABI

// EthAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthAuctionMetaData.Bin instead.
var EthAuctionBin = EthAuctionMetaData.Bin

// DeployEthAuction deploys a new Ethereum contract, binding an instance of EthAuction to it.
func DeployEthAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *EthAuction, error) {
	parsed, err := EthAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthAuctionBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthAuction{EthAuctionCaller: EthAuctionCaller{contract: contract}, EthAuctionTransactor: EthAuctionTransactor{contract: contract}, EthAuctionFilterer: EthAuctionFilterer{contract: contract}}, nil
}

// EthAuction is an auto generated Go binding around an Ethereum contract.
type EthAuction struct {
	EthAuctionCaller     // Read-only binding to the contract
	EthAuctionTransactor // Write-only binding to the contract
	EthAuctionFilterer   // Log filterer for contract events
}

// EthAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthAuctionSession struct {
	Contract     *EthAuction       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthAuctionCallerSession struct {
	Contract *EthAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EthAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthAuctionTransactorSession struct {
	Contract     *EthAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EthAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthAuctionRaw struct {
	Contract *EthAuction // Generic contract binding to access the raw methods on
}

// EthAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthAuctionCallerRaw struct {
	Contract *EthAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// EthAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthAuctionTransactorRaw struct {
	Contract *EthAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthAuction creates a new instance of EthAuction, bound to a specific deployed contract.
func NewEthAuction(address common.Address, backend bind.ContractBackend) (*EthAuction, error) {
	contract, err := bindEthAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthAuction{EthAuctionCaller: EthAuctionCaller{contract: contract}, EthAuctionTransactor: EthAuctionTransactor{contract: contract}, EthAuctionFilterer: EthAuctionFilterer{contract: contract}}, nil
}

// NewEthAuctionCaller creates a new read-only instance of EthAuction, bound to a specific deployed contract.
func NewEthAuctionCaller(address common.Address, caller bind.ContractCaller) (*EthAuctionCaller, error) {
	contract, err := bindEthAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthAuctionCaller{contract: contract}, nil
}

// NewEthAuctionTransactor creates a new write-only instance of EthAuction, bound to a specific deployed contract.
func NewEthAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*EthAuctionTransactor, error) {
	contract, err := bindEthAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthAuctionTransactor{contract: contract}, nil
}

// NewEthAuctionFilterer creates a new log filterer instance of EthAuction, bound to a specific deployed contract.
func NewEthAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*EthAuctionFilterer, error) {
	contract, err := bindEthAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthAuctionFilterer{contract: contract}, nil
}

// bindEthAuction binds a generic wrapper to an already deployed contract.
func bindEthAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthAuction *EthAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthAuction.Contract.EthAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthAuction *EthAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthAuction.Contract.EthAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthAuction *EthAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthAuction.Contract.EthAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthAuction *EthAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthAuction *EthAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthAuction *EthAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthAuction.Contract.contract.Transact(opts, method, params...)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_EthAuction *EthAuctionCaller) Ended(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EthAuction.contract.Call(opts, &out, "ended")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_EthAuction *EthAuctionSession) Ended() (bool, error) {
	return _EthAuction.Contract.Ended(&_EthAuction.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_EthAuction *EthAuctionCallerSession) Ended() (bool, error) {
	return _EthAuction.Contract.Ended(&_EthAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_EthAuction *EthAuctionCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthAuction.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_EthAuction *EthAuctionSession) HighestBid() (*big.Int, error) {
	return _EthAuction.Contract.HighestBid(&_EthAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_EthAuction *EthAuctionCallerSession) HighestBid() (*big.Int, error) {
	return _EthAuction.Contract.HighestBid(&_EthAuction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_EthAuction *EthAuctionCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthAuction.contract.Call(opts, &out, "highestBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_EthAuction *EthAuctionSession) HighestBidder() (common.Address, error) {
	return _EthAuction.Contract.HighestBidder(&_EthAuction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_EthAuction *EthAuctionCallerSession) HighestBidder() (common.Address, error) {
	return _EthAuction.Contract.HighestBidder(&_EthAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthAuction *EthAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthAuction *EthAuctionSession) Owner() (common.Address, error) {
	return _EthAuction.Contract.Owner(&_EthAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthAuction *EthAuctionCallerSession) Owner() (common.Address, error) {
	return _EthAuction.Contract.Owner(&_EthAuction.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(string)
func (_EthAuction *EthAuctionCaller) Status(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EthAuction.contract.Call(opts, &out, "status")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(string)
func (_EthAuction *EthAuctionSession) Status() (string, error) {
	return _EthAuction.Contract.Status(&_EthAuction.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(string)
func (_EthAuction *EthAuctionCallerSession) Status() (string, error) {
	return _EthAuction.Contract.Status(&_EthAuction.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthAuction *EthAuctionCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthAuction.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthAuction *EthAuctionSession) Token() (common.Address, error) {
	return _EthAuction.Contract.Token(&_EthAuction.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthAuction *EthAuctionCallerSession) Token() (common.Address, error) {
	return _EthAuction.Contract.Token(&_EthAuction.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_EthAuction *EthAuctionTransactor) Abort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthAuction.contract.Transact(opts, "abort")
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_EthAuction *EthAuctionSession) Abort() (*types.Transaction, error) {
	return _EthAuction.Contract.Abort(&_EthAuction.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_EthAuction *EthAuctionTransactorSession) Abort() (*types.Transaction, error) {
	return _EthAuction.Contract.Abort(&_EthAuction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 bidAmount) returns()
func (_EthAuction *EthAuctionTransactor) Bid(opts *bind.TransactOpts, bidAmount *big.Int) (*types.Transaction, error) {
	return _EthAuction.contract.Transact(opts, "bid", bidAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 bidAmount) returns()
func (_EthAuction *EthAuctionSession) Bid(bidAmount *big.Int) (*types.Transaction, error) {
	return _EthAuction.Contract.Bid(&_EthAuction.TransactOpts, bidAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 bidAmount) returns()
func (_EthAuction *EthAuctionTransactorSession) Bid(bidAmount *big.Int) (*types.Transaction, error) {
	return _EthAuction.Contract.Bid(&_EthAuction.TransactOpts, bidAmount)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_EthAuction *EthAuctionTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthAuction.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_EthAuction *EthAuctionSession) EndAuction() (*types.Transaction, error) {
	return _EthAuction.Contract.EndAuction(&_EthAuction.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_EthAuction *EthAuctionTransactorSession) EndAuction() (*types.Transaction, error) {
	return _EthAuction.Contract.EndAuction(&_EthAuction.TransactOpts)
}

// Proceed is a paid mutator transaction binding the contract method 0x2a33fec6.
//
// Solidity: function proceed() returns()
func (_EthAuction *EthAuctionTransactor) Proceed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthAuction.contract.Transact(opts, "proceed")
}

// Proceed is a paid mutator transaction binding the contract method 0x2a33fec6.
//
// Solidity: function proceed() returns()
func (_EthAuction *EthAuctionSession) Proceed() (*types.Transaction, error) {
	return _EthAuction.Contract.Proceed(&_EthAuction.TransactOpts)
}

// Proceed is a paid mutator transaction binding the contract method 0x2a33fec6.
//
// Solidity: function proceed() returns()
func (_EthAuction *EthAuctionTransactorSession) Proceed() (*types.Transaction, error) {
	return _EthAuction.Contract.Proceed(&_EthAuction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_EthAuction *EthAuctionTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthAuction.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_EthAuction *EthAuctionSession) Withdraw() (*types.Transaction, error) {
	return _EthAuction.Contract.Withdraw(&_EthAuction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_EthAuction *EthAuctionTransactorSession) Withdraw() (*types.Transaction, error) {
	return _EthAuction.Contract.Withdraw(&_EthAuction.TransactOpts)
}

// EthAuctionAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the EthAuction contract.
type EthAuctionAuctionEndedIterator struct {
	Event *EthAuctionAuctionEnded // Event containing the contract specifics and raw log

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
func (it *EthAuctionAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthAuctionAuctionEnded)
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
		it.Event = new(EthAuctionAuctionEnded)
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
func (it *EthAuctionAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthAuctionAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthAuctionAuctionEnded represents a AuctionEnded event raised by the EthAuction contract.
type EthAuctionAuctionEnded struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_EthAuction *EthAuctionFilterer) FilterAuctionEnded(opts *bind.FilterOpts) (*EthAuctionAuctionEndedIterator, error) {

	logs, sub, err := _EthAuction.contract.FilterLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return &EthAuctionAuctionEndedIterator{contract: _EthAuction.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_EthAuction *EthAuctionFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *EthAuctionAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _EthAuction.contract.WatchLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthAuctionAuctionEnded)
				if err := _EthAuction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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

// ParseAuctionEnded is a log parse operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_EthAuction *EthAuctionFilterer) ParseAuctionEnded(log types.Log) (*EthAuctionAuctionEnded, error) {
	event := new(EthAuctionAuctionEnded)
	if err := _EthAuction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthAuctionHighestBidIncreasedIterator is returned from FilterHighestBidIncreased and is used to iterate over the raw logs and unpacked data for HighestBidIncreased events raised by the EthAuction contract.
type EthAuctionHighestBidIncreasedIterator struct {
	Event *EthAuctionHighestBidIncreased // Event containing the contract specifics and raw log

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
func (it *EthAuctionHighestBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthAuctionHighestBidIncreased)
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
		it.Event = new(EthAuctionHighestBidIncreased)
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
func (it *EthAuctionHighestBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthAuctionHighestBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthAuctionHighestBidIncreased represents a HighestBidIncreased event raised by the EthAuction contract.
type EthAuctionHighestBidIncreased struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0xf4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300.
//
// Solidity: event HighestBidIncreased(address bidder, uint256 amount)
func (_EthAuction *EthAuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*EthAuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _EthAuction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &EthAuctionHighestBidIncreasedIterator{contract: _EthAuction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0xf4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300.
//
// Solidity: event HighestBidIncreased(address bidder, uint256 amount)
func (_EthAuction *EthAuctionFilterer) WatchHighestBidIncreased(opts *bind.WatchOpts, sink chan<- *EthAuctionHighestBidIncreased) (event.Subscription, error) {

	logs, sub, err := _EthAuction.contract.WatchLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthAuctionHighestBidIncreased)
				if err := _EthAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
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

// ParseHighestBidIncreased is a log parse operation binding the contract event 0xf4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300.
//
// Solidity: event HighestBidIncreased(address bidder, uint256 amount)
func (_EthAuction *EthAuctionFilterer) ParseHighestBidIncreased(log types.Log) (*EthAuctionHighestBidIncreased, error) {
	event := new(EthAuctionHighestBidIncreased)
	if err := _EthAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthAuctionWaitResponseIterator is returned from FilterWaitResponse and is used to iterate over the raw logs and unpacked data for WaitResponse events raised by the EthAuction contract.
type EthAuctionWaitResponseIterator struct {
	Event *EthAuctionWaitResponse // Event containing the contract specifics and raw log

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
func (it *EthAuctionWaitResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthAuctionWaitResponse)
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
		it.Event = new(EthAuctionWaitResponse)
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
func (it *EthAuctionWaitResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthAuctionWaitResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthAuctionWaitResponse represents a WaitResponse event raised by the EthAuction contract.
type EthAuctionWaitResponse struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWaitResponse is a free log retrieval operation binding the contract event 0x0472d43abcf056976d4a0f799dccd47a3fadcc3f385763eed302bdcc802bc20c.
//
// Solidity: event WaitResponse(address winner, uint256 amount)
func (_EthAuction *EthAuctionFilterer) FilterWaitResponse(opts *bind.FilterOpts) (*EthAuctionWaitResponseIterator, error) {

	logs, sub, err := _EthAuction.contract.FilterLogs(opts, "WaitResponse")
	if err != nil {
		return nil, err
	}
	return &EthAuctionWaitResponseIterator{contract: _EthAuction.contract, event: "WaitResponse", logs: logs, sub: sub}, nil
}

// WatchWaitResponse is a free log subscription operation binding the contract event 0x0472d43abcf056976d4a0f799dccd47a3fadcc3f385763eed302bdcc802bc20c.
//
// Solidity: event WaitResponse(address winner, uint256 amount)
func (_EthAuction *EthAuctionFilterer) WatchWaitResponse(opts *bind.WatchOpts, sink chan<- *EthAuctionWaitResponse) (event.Subscription, error) {

	logs, sub, err := _EthAuction.contract.WatchLogs(opts, "WaitResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthAuctionWaitResponse)
				if err := _EthAuction.contract.UnpackLog(event, "WaitResponse", log); err != nil {
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

// ParseWaitResponse is a log parse operation binding the contract event 0x0472d43abcf056976d4a0f799dccd47a3fadcc3f385763eed302bdcc802bc20c.
//
// Solidity: event WaitResponse(address winner, uint256 amount)
func (_EthAuction *EthAuctionFilterer) ParseWaitResponse(log types.Log) (*EthAuctionWaitResponse, error) {
	event := new(EthAuctionWaitResponse)
	if err := _EthAuction.contract.UnpackLog(event, "WaitResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
