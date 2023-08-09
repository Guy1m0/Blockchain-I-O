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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_asset_id\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DecisionMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"WaitResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdarwBid\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"not_winner_platform\",\"type\":\"bool\"}],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proceed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620011cc380380620011cc8339810160408190526200003491620000bb565b600780546001600160a01b0319166001600160a01b038416178155604080518082019091529081526672756e6e696e6760c81b60208201526004906200007b908262000240565b50600580546001600160a01b0319163317905560066200009c828262000240565b5050506200030c565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215620000cf57600080fd5b82516001600160a01b0381168114620000e757600080fd5b602084810151919350906001600160401b03808211156200010757600080fd5b818601915086601f8301126200011c57600080fd5b815181811115620001315762000131620000a5565b604051601f8201601f19908116603f011681019083821181831017156200015c576200015c620000a5565b8160405282815289868487010111156200017557600080fd5b600093505b828410156200019957848401860151818501870152928501926200017a565b60008684830101528096505050505050509250929050565b600181811c90821680620001c657607f821691505b602082108103620001e757634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200023b57600081815260208120601f850160051c81016020861015620002165750805b601f850160051c820191505b81811015620002375782815560010162000222565b5050505b505050565b81516001600160401b038111156200025c576200025c620000a5565b62000274816200026d8454620001b1565b84620001ed565b602080601f831160018114620002ac5760008415620002935750858301515b600019600386901b1c1916600185901b17855562000237565b600085815260208120601f198616915b82811015620002dd57888601518255948401946001909101908401620002bc565b5085821015620002fc5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610eb0806200031c6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063454a2ab311610071578063454a2ab31461011257806378ed134f146101255780638da5cb5b1461013857806391f9015714610163578063d57bde7914610176578063fc0c546a1461018d57600080fd5b806312fa6feb146100b9578063200d2ed2146100db5780632694c611146100f05780632a33fec6146100f857806335a063b4146101025780633ccfd60b1461010a575b600080fd5b6003546100c69060ff1681565b60405190151581526020015b60405180910390f35b6100e36101a0565b6040516100d29190610b26565b6100e361022e565b61010061023b565b005b610100610412565b6100c66105a8565b610100610120366004610b74565b61074b565b610100610133366004610b9b565b610991565b60055461014b906001600160a01b031681565b6040516001600160a01b0390911681526020016100d2565b60005461014b906001600160a01b031681565b61017f60015481565b6040519081526020016100d2565b60075461014b906001600160a01b031681565b600480546101ad90610bbf565b80601f01602080910402602001604051908101604052809291908181526020018280546101d990610bbf565b80156102265780601f106101fb57610100808354040283529160200191610226565b820191906000526020600020905b81548152906001019060200180831161020957829003601f168201915b505050505081565b600680546101ad90610bbf565b60405165656e64696e6760d01b60208201526026016040516020818303038152906040528051906020012060046040516020016102789190610bf9565b60405160208183030381529060405280519060200120146102e05760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064015b60405180910390fd5b6000546001600160a01b031633146103335760405162461bcd60e51b81526020600482015260166024820152754e6f7420617574686f72697a6564206163636573732160501b60448201526064016102d7565b604080518082019091526005815264195b99195960da1b602082015260049061035c9082610cd4565b50600754600154604051632770a7eb60e21b815230600482015260248101919091526001600160a01b0390911690639dc29fac90604401600060405180830381600087803b1580156103ad57600080fd5b505af11580156103c1573d6000803e3d6000fd5b50506000546001546040517f6ad8c9e6ea2bba271edecfd782216e47cf0d6a4b4d7328a94c0b33b3861c411d945061040893506001600160a01b0390921691600690610d94565b60405180910390a1565b60405165656e64696e6760d01b602082015260260160405160208183030381529060405280519060200120600460405160200161044f9190610bf9565b60405160208183030381529060405280519060200120146104b25760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064016102d7565b6005546001600160a01b03163314806104d557506000546001600160a01b031633145b61051a5760405162461bcd60e51b81526020600482015260166024820152754e6f7420617574686f72697a6564206163636573732160501b60448201526064016102d7565b604080518082019091526005815264195b99195960da1b60208201526004906105439082610cd4565b5060018054600080546001600160a01b03168152600260205260408082209290925580546001600160a01b031916815591829055517f6ad8c9e6ea2bba271edecfd782216e47cf0d6a4b4d7328a94c0b33b3861c411d91610408918190600690610d94565b60405164195b99195960da1b60208201526000906025016040516020818303038152906040528051906020012060046040516020016105e79190610bf9565b604051602081830303815290604052805190602001201461064a5760405162461bcd60e51b815260206004820152601c60248201527f436f6e7472616374206e6f7420696e20454e444544207374617475730000000060448201526064016102d7565b33600090815260026020526040902054801561070a57336000818152600260205260408082209190915560075490516323b872dd60e01b81523060048201526024810192909252604482018390526001600160a01b0316906323b872dd906064016020604051808303816000875af11580156106ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ee9190610e36565b61070a5733600090815260026020526040812091909155919050565b60408051338152602081018390527fe47ad0d09522305576ea36f7f51118044f2d4e81be386a155f0e2cac8e1220b5910160405180910390a1600191505090565b6040516672756e6e696e6760c81b60208201526027016040516020818303038152906040528051906020012060046040516020016107899190610bf9565b60405160208183030381529060405280519060200120146107ec5760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e2052554e4e494e4720737461747573000060448201526064016102d7565b600154811161083d5760405162461bcd60e51b815260206004820152601e60248201527f546865726520616c7265616479206973206120686967686572206269642e000060448201526064016102d7565b6007546040516323b872dd60e01b8152336004820152306024820152604481018390526000916001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610895573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b99190610e36565b9050806109015760405162461bcd60e51b81526020600482015260166024820152752a37b5b2b7103a3930b739b332b9103330b4b632b21760511b60448201526064016102d7565b6001541561093a57600154600080546001600160a01b031681526002602052604081208054909190610934908490610e53565b90915550505b600080546001600160a01b03191633908117909155600183905560408051918252602082018490527ff4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300910160405180910390a15050565b6005546001600160a01b031633146109fb5760405162461bcd60e51b815260206004820152602760248201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736044820152662073746174757360c81b60648201526084016102d7565b6040516672756e6e696e6760c81b6020820152602701604051602081830303815290604052805190602001206004604051602001610a399190610bf9565b6040516020818303038152906040528051906020012014610a9c5760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e2052554e4e494e4720737461747573000060448201526064016102d7565b604080518082019091526006815265656e64696e6760d01b6020820152600490610ac69082610cd4565b508080610ad35750600154155b15610ae357610ae0610412565b50565b6000546040516001600160a01b0390911681527ff50e28b7cc4028b8226b711d1857dce1bb3f3325b879f1fe4653c4631c6bd28c9060200160405180910390a150565b600060208083528351808285015260005b81811015610b5357858101830151858201604001528201610b37565b506000604082860101526040601f19601f8301168501019250505092915050565b600060208284031215610b8657600080fd5b5035919050565b8015158114610ae057600080fd5b600060208284031215610bad57600080fd5b8135610bb881610b8d565b9392505050565b600181811c90821680610bd357607f821691505b602082108103610bf357634e487b7160e01b600052602260045260246000fd5b50919050565b6000808354610c0781610bbf565b60018281168015610c1f5760018114610c3457610c63565b60ff1984168752821515830287019450610c63565b8760005260208060002060005b85811015610c5a5781548a820152908401908201610c41565b50505082870194505b50929695505050505050565b634e487b7160e01b600052604160045260246000fd5b601f821115610ccf57600081815260208120601f850160051c81016020861015610cac5750805b601f850160051c820191505b81811015610ccb57828155600101610cb8565b5050505b505050565b815167ffffffffffffffff811115610cee57610cee610c6f565b610d0281610cfc8454610bbf565b84610c85565b602080601f831160018114610d375760008415610d1f5750858301515b600019600386901b1c1916600185901b178555610ccb565b600085815260208120601f198616915b82811015610d6657888601518255948401946001909101908401610d47565b5085821015610d845787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60018060a01b03841681526000602084818401526060604084015260008454610dbc81610bbf565b8060608701526080600180841660008114610dde5760018114610df857610e26565b60ff1985168984015283151560051b890183019550610e26565b896000528660002060005b85811015610e1e5781548b8201860152908301908801610e03565b8a0184019650505b50939a9950505050505050505050565b600060208284031215610e4857600080fd5b8151610bb881610b8d565b80820180821115610e7457634e487b7160e01b600052601160045260246000fd5b9291505056fea2646970667358221220692ee4d8809090e7af9584df3033dcad836dff9081d66ca3f3c2e7288d21a1d964736f6c63430008120033",
}

// EthAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use EthAuctionMetaData.ABI instead.
var EthAuctionABI = EthAuctionMetaData.ABI

// EthAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthAuctionMetaData.Bin instead.
var EthAuctionBin = EthAuctionMetaData.Bin

// DeployEthAuction deploys a new Ethereum contract, binding an instance of EthAuction to it.
func DeployEthAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _asset_id string) (common.Address, *types.Transaction, *EthAuction, error) {
	parsed, err := EthAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthAuctionBin), backend, _token, _asset_id)
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

// AssetId is a free data retrieval call binding the contract method 0x2694c611.
//
// Solidity: function asset_id() view returns(string)
func (_EthAuction *EthAuctionCaller) AssetId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EthAuction.contract.Call(opts, &out, "asset_id")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0x2694c611.
//
// Solidity: function asset_id() view returns(string)
func (_EthAuction *EthAuctionSession) AssetId() (string, error) {
	return _EthAuction.Contract.AssetId(&_EthAuction.CallOpts)
}

// AssetId is a free data retrieval call binding the contract method 0x2694c611.
//
// Solidity: function asset_id() view returns(string)
func (_EthAuction *EthAuctionCallerSession) AssetId() (string, error) {
	return _EthAuction.Contract.AssetId(&_EthAuction.CallOpts)
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

// EndAuction is a paid mutator transaction binding the contract method 0x78ed134f.
//
// Solidity: function endAuction(bool not_winner_platform) returns()
func (_EthAuction *EthAuctionTransactor) EndAuction(opts *bind.TransactOpts, not_winner_platform bool) (*types.Transaction, error) {
	return _EthAuction.contract.Transact(opts, "endAuction", not_winner_platform)
}

// EndAuction is a paid mutator transaction binding the contract method 0x78ed134f.
//
// Solidity: function endAuction(bool not_winner_platform) returns()
func (_EthAuction *EthAuctionSession) EndAuction(not_winner_platform bool) (*types.Transaction, error) {
	return _EthAuction.Contract.EndAuction(&_EthAuction.TransactOpts, not_winner_platform)
}

// EndAuction is a paid mutator transaction binding the contract method 0x78ed134f.
//
// Solidity: function endAuction(bool not_winner_platform) returns()
func (_EthAuction *EthAuctionTransactorSession) EndAuction(not_winner_platform bool) (*types.Transaction, error) {
	return _EthAuction.Contract.EndAuction(&_EthAuction.TransactOpts, not_winner_platform)
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

// EthAuctionDecisionMadeIterator is returned from FilterDecisionMade and is used to iterate over the raw logs and unpacked data for DecisionMade events raised by the EthAuction contract.
type EthAuctionDecisionMadeIterator struct {
	Event *EthAuctionDecisionMade // Event containing the contract specifics and raw log

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
func (it *EthAuctionDecisionMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthAuctionDecisionMade)
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
		it.Event = new(EthAuctionDecisionMade)
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
func (it *EthAuctionDecisionMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthAuctionDecisionMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthAuctionDecisionMade represents a DecisionMade event raised by the EthAuction contract.
type EthAuctionDecisionMade struct {
	Winner common.Address
	Amount *big.Int
	Id     string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecisionMade is a free log retrieval operation binding the contract event 0x6ad8c9e6ea2bba271edecfd782216e47cf0d6a4b4d7328a94c0b33b3861c411d.
//
// Solidity: event DecisionMade(address winner, uint256 amount, string id)
func (_EthAuction *EthAuctionFilterer) FilterDecisionMade(opts *bind.FilterOpts) (*EthAuctionDecisionMadeIterator, error) {

	logs, sub, err := _EthAuction.contract.FilterLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return &EthAuctionDecisionMadeIterator{contract: _EthAuction.contract, event: "DecisionMade", logs: logs, sub: sub}, nil
}

// WatchDecisionMade is a free log subscription operation binding the contract event 0x6ad8c9e6ea2bba271edecfd782216e47cf0d6a4b4d7328a94c0b33b3861c411d.
//
// Solidity: event DecisionMade(address winner, uint256 amount, string id)
func (_EthAuction *EthAuctionFilterer) WatchDecisionMade(opts *bind.WatchOpts, sink chan<- *EthAuctionDecisionMade) (event.Subscription, error) {

	logs, sub, err := _EthAuction.contract.WatchLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthAuctionDecisionMade)
				if err := _EthAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
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

// ParseDecisionMade is a log parse operation binding the contract event 0x6ad8c9e6ea2bba271edecfd782216e47cf0d6a4b4d7328a94c0b33b3861c411d.
//
// Solidity: event DecisionMade(address winner, uint256 amount, string id)
func (_EthAuction *EthAuctionFilterer) ParseDecisionMade(log types.Log) (*EthAuctionDecisionMade, error) {
	event := new(EthAuctionDecisionMade)
	if err := _EthAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
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
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWaitResponse is a free log retrieval operation binding the contract event 0xf50e28b7cc4028b8226b711d1857dce1bb3f3325b879f1fe4653c4631c6bd28c.
//
// Solidity: event WaitResponse(address winner)
func (_EthAuction *EthAuctionFilterer) FilterWaitResponse(opts *bind.FilterOpts) (*EthAuctionWaitResponseIterator, error) {

	logs, sub, err := _EthAuction.contract.FilterLogs(opts, "WaitResponse")
	if err != nil {
		return nil, err
	}
	return &EthAuctionWaitResponseIterator{contract: _EthAuction.contract, event: "WaitResponse", logs: logs, sub: sub}, nil
}

// WatchWaitResponse is a free log subscription operation binding the contract event 0xf50e28b7cc4028b8226b711d1857dce1bb3f3325b879f1fe4653c4631c6bd28c.
//
// Solidity: event WaitResponse(address winner)
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

// ParseWaitResponse is a log parse operation binding the contract event 0xf50e28b7cc4028b8226b711d1857dce1bb3f3325b879f1fe4653c4631c6bd28c.
//
// Solidity: event WaitResponse(address winner)
func (_EthAuction *EthAuctionFilterer) ParseWaitResponse(log types.Log) (*EthAuctionWaitResponse, error) {
	event := new(EthAuctionWaitResponse)
	if err := _EthAuction.contract.UnpackLog(event, "WaitResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthAuctionWithdarwBidIterator is returned from FilterWithdarwBid and is used to iterate over the raw logs and unpacked data for WithdarwBid events raised by the EthAuction contract.
type EthAuctionWithdarwBidIterator struct {
	Event *EthAuctionWithdarwBid // Event containing the contract specifics and raw log

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
func (it *EthAuctionWithdarwBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthAuctionWithdarwBid)
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
		it.Event = new(EthAuctionWithdarwBid)
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
func (it *EthAuctionWithdarwBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthAuctionWithdarwBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthAuctionWithdarwBid represents a WithdarwBid event raised by the EthAuction contract.
type EthAuctionWithdarwBid struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdarwBid is a free log retrieval operation binding the contract event 0xe47ad0d09522305576ea36f7f51118044f2d4e81be386a155f0e2cac8e1220b5.
//
// Solidity: event WithdarwBid(address bidder, uint256 amount)
func (_EthAuction *EthAuctionFilterer) FilterWithdarwBid(opts *bind.FilterOpts) (*EthAuctionWithdarwBidIterator, error) {

	logs, sub, err := _EthAuction.contract.FilterLogs(opts, "WithdarwBid")
	if err != nil {
		return nil, err
	}
	return &EthAuctionWithdarwBidIterator{contract: _EthAuction.contract, event: "WithdarwBid", logs: logs, sub: sub}, nil
}

// WatchWithdarwBid is a free log subscription operation binding the contract event 0xe47ad0d09522305576ea36f7f51118044f2d4e81be386a155f0e2cac8e1220b5.
//
// Solidity: event WithdarwBid(address bidder, uint256 amount)
func (_EthAuction *EthAuctionFilterer) WatchWithdarwBid(opts *bind.WatchOpts, sink chan<- *EthAuctionWithdarwBid) (event.Subscription, error) {

	logs, sub, err := _EthAuction.contract.WatchLogs(opts, "WithdarwBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthAuctionWithdarwBid)
				if err := _EthAuction.contract.UnpackLog(event, "WithdarwBid", log); err != nil {
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

// ParseWithdarwBid is a log parse operation binding the contract event 0xe47ad0d09522305576ea36f7f51118044f2d4e81be386a155f0e2cac8e1220b5.
//
// Solidity: event WithdarwBid(address bidder, uint256 amount)
func (_EthAuction *EthAuctionFilterer) ParseWithdarwBid(log types.Log) (*EthAuctionWithdarwBid, error) {
	event := new(EthAuctionWithdarwBid)
	if err := _EthAuction.contract.UnpackLog(event, "WithdarwBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
