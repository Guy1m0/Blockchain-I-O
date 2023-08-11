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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_asset_id\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prcd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"DecisionMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"rating\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"review\",\"type\":\"string\"}],\"name\":\"RateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"WaitResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdarwBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkFeedback\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"not_winner_platform\",\"type\":\"bool\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_score\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_feedback\",\"type\":\"string\"}],\"name\":\"provide_feedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200166e3803806200166e8339810160408190526200003491620000b9565b600880546001600160a01b0319166001600160a01b03841617905560408051808201909152600481526337b832b760e11b60208201526003906200007990826200023e565b50600480546001600160a01b0319163317905560056200009a82826200023e565b5050506200030a565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215620000cd57600080fd5b82516001600160a01b0381168114620000e557600080fd5b602084810151919350906001600160401b03808211156200010557600080fd5b818601915086601f8301126200011a57600080fd5b8151818111156200012f576200012f620000a3565b604051601f8201601f19908116603f011681019083821181831017156200015a576200015a620000a3565b8160405282815289868487010111156200017357600080fd5b600093505b8284101562000197578484018601518185018701529285019262000178565b60008684830101528096505050505050509250929050565b600181811c90821680620001c457607f821691505b602082108103620001e557634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200023957600081815260208120601f850160051c81016020861015620002145750805b601f850160051c820191505b81811015620002355782815560010162000220565b5050505b505050565b81516001600160401b038111156200025a576200025a620000a3565b62000272816200026b8454620001af565b84620001eb565b602080601f831160018114620002aa5760008415620002915750858301515b600019600386901b1c1916600185901b17855562000235565b600085815260208120601f198616915b82811015620002db57888601518255948401946001909101908401620002ba565b5085821015620002fa5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b611354806200031a6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063588579f91161008c5780639867db74116100665780639867db74146101a1578063cc5284e7146101b4578063d57bde79146101c7578063fc0c546a146101de57600080fd5b8063588579f91461014d5780638da5cb5b1461016357806391f901571461018e57600080fd5b8063200d2ed2146100d45780632694c611146100f257806335acac4e146100fa5780633ccfd60b1461010f578063454a2ab31461012757806345dad86c1461013a575b600080fd5b6100dc6101f1565b6040516100e99190610e4c565b60405180910390f35b6100dc61027f565b61010d610108366004610f09565b61028c565b005b610117610408565b60405190151581526020016100e9565b61010d610135366004610f46565b6105aa565b61010d610148366004610f70565b6107ed565b6101556109d3565b6040516100e9929190610f8d565b600454610176906001600160a01b031681565b6040516001600160a01b0390911681526020016100e9565b600054610176906001600160a01b031681565b61010d6101af366004610f09565b610b17565b61010d6101c2366004610fa6565b610cbd565b6101d060015481565b6040519081526020016100e9565b600854610176906001600160a01b031681565b600380546101fe90610fed565b80601f016020809104026020016040519081016040528092919081815260200182805461022a90610fed565b80156102775780601f1061024c57610100808354040283529160200191610277565b820191906000526020600020905b81548152906001019060200180831161025a57829003601f168201915b505050505081565b600580546101fe90610fed565b60405165656e64696e6760d01b60208201526026016040516020818303038152906040528051906020012060036040516020016102c99190611027565b60405160208183030381529060405280519060200120146103315760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064015b60405180910390fd5b6000546001600160a01b0316331461035b5760405162461bcd60e51b81526004016103289061109d565b604080518082019091526007815266636c6f73696e6760c81b6020820152600390610386908261111b565b50600080546001546040517f9a51e3deb4a767f1ba793da152f8b62dde0121eb58730d9a86124e9b59afe978936103cd936001600160a01b03169291600591908790611258565b60405180910390a15060018054600080546001600160a01b031681526002602052604081209190915580546001600160a01b03191681559055565b6040516337b832b760e11b60208201526000906024016040516020818303038152906040528051906020012060036040516020016104469190611027565b60405160208183030381529060405280519060200120036104a95760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420696e204f50454e207374617475730000000000000000006044820152606401610328565b33600090815260026020526040902054801561056957336000818152600260205260408082209190915560085490516323b872dd60e01b81523060048201526024810192909252604482018390526001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610529573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061054d91906112a5565b6105695733600090815260026020526040812091909155919050565b60408051338152602081018390527fe47ad0d09522305576ea36f7f51118044f2d4e81be386a155f0e2cac8e1220b5910160405180910390a1600191505090565b6040516337b832b760e11b60208201526024016040516020818303038152906040528051906020012060036040516020016105e59190611027565b60405160208183030381529060405280519060200120146106485760405162461bcd60e51b815260206004820152601b60248201527f436f6e7472616374206e6f7420696e204f50454e2073746174757300000000006044820152606401610328565b60015481116106995760405162461bcd60e51b815260206004820152601e60248201527f546865726520616c7265616479206973206120686967686572206269642e00006044820152606401610328565b6008546040516323b872dd60e01b8152336004820152306024820152604481018390526000916001600160a01b0316906323b872dd906064016020604051808303816000875af11580156106f1573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071591906112a5565b90508061075d5760405162461bcd60e51b81526020600482015260166024820152752a37b5b2b7103a3930b739b332b9103330b4b632b21760511b6044820152606401610328565b6001541561079657600154600080546001600160a01b0316815260026020526040812080549091906107909084906112c2565b90915550505b600080546001600160a01b03191633908117909155600183905560408051918252602082018490527ff4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300910160405180910390a15050565b6004546001600160a01b031633146108575760405162461bcd60e51b815260206004820152602760248201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736044820152662073746174757360c81b6064820152608401610328565b6040516337b832b760e11b60208201526024016040516020818303038152906040528051906020012060036040516020016108929190611027565b60405160208183030381529060405280519060200120146108f55760405162461bcd60e51b815260206004820152601b60248201527f436f6e7472616374206e6f7420696e204f50454e2073746174757300000000006044820152606401610328565b604080518082019091526006815265656e64696e6760d01b602082015260039061091f908261111b565b50808061092c5750600154155b1561098f5760408051808201909152600681526518db1bdcd95960d21b602082015260039061095b908261111b565b505060018054600080546001600160a01b031681526002602052604081209190915580546001600160a01b03191681559055565b6000546040516001600160a01b0390911681527ff50e28b7cc4028b8226b711d1857dce1bb3f3325b879f1fe4653c4631c6bd28c906020015b60405180910390a150565b600060606040516020016109f3906518db1bdcd95960d21b815260060190565b604051602081830303815290604052805190602001206003604051602001610a1b9190611027565b6040516020818303038152906040528051906020012014610a7e5760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20434c4f534544207374617475730000006044820152606401610328565b6007546006808054610a8f90610fed565b80601f0160208091040260200160405190810160405280929190818152602001828054610abb90610fed565b8015610b085780601f10610add57610100808354040283529160200191610b08565b820191906000526020600020905b815481529060010190602001808311610aeb57829003601f168201915b50505050509050915091509091565b60405165656e64696e6760d01b6020820152602601604051602081830303815290604052805190602001206003604051602001610b549190611027565b6040516020818303038152906040528051906020012014610bb75760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e47207374617475730000006044820152606401610328565b6000546001600160a01b03163314610be15760405162461bcd60e51b81526004016103289061109d565b604080518082019091526007815266636c6f73696e6760c81b6020820152600390610c0c908261111b565b50600854600154604051632770a7eb60e21b815230600482015260248101919091526001600160a01b0390911690639dc29fac90604401600060405180830381600087803b158015610c5d57600080fd5b505af1158015610c71573d6000803e3d6000fd5b5050600054600180546040517f9a51e3deb4a767f1ba793da152f8b62dde0121eb58730d9a86124e9b59afe97895506109c894506001600160a01b039093169290916005918790611258565b60405166636c6f73696e6760c81b6020820152602701604051602081830303815290604052805190602001206003604051602001610cfb9190611027565b6040516020818303038152906040528051906020012014610d5e5760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e20434c4f53494e472073746174757300006044820152606401610328565b6000546001600160a01b03163314610d885760405162461bcd60e51b81526004016103289061109d565b6006610d94828261111b565b5060078290556040517f41ba38777fb1a414ddf3be769b0df73ab388f5f58a0edc435ce9736eaa9d758e90610dcf90600590859085906112e9565b60405180910390a160408051808201909152600681526518db1bdcd95960d21b6020820152600390610e01908261111b565b505050565b6000815180845260005b81811015610e2c57602081850181015186830182015201610e10565b506000602082860101526020601f19601f83011685010191505092915050565b602081526000610e5f6020830184610e06565b9392505050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610e8d57600080fd5b813567ffffffffffffffff80821115610ea857610ea8610e66565b604051601f8301601f19908116603f01168101908282118183101715610ed057610ed0610e66565b81604052838152866020858801011115610ee957600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060208284031215610f1b57600080fd5b813567ffffffffffffffff811115610f3257600080fd5b610f3e84828501610e7c565b949350505050565b600060208284031215610f5857600080fd5b5035919050565b8015158114610f6d57600080fd5b50565b600060208284031215610f8257600080fd5b8135610e5f81610f5f565b828152604060208201526000610f3e6040830184610e06565b60008060408385031215610fb957600080fd5b82359150602083013567ffffffffffffffff811115610fd757600080fd5b610fe385828601610e7c565b9150509250929050565b600181811c9082168061100157607f821691505b60208210810361102157634e487b7160e01b600052602260045260246000fd5b50919050565b600080835461103581610fed565b6001828116801561104d576001811461106257611091565b60ff1984168752821515830287019450611091565b8760005260208060002060005b858110156110885781548a82015290840190820161106f565b50505082870194505b50929695505050505050565b6020808252601690820152754e6f7420617574686f72697a6564206163636573732160501b604082015260600190565b601f821115610e0157600081815260208120601f850160051c810160208610156110f45750805b601f850160051c820191505b8181101561111357828155600101611100565b505050505050565b815167ffffffffffffffff81111561113557611135610e66565b611149816111438454610fed565b846110cd565b602080601f83116001811461117e57600084156111665750858301515b600019600386901b1c1916600185901b178555611113565b600085815260208120601f198616915b828110156111ad5788860151825594840194600190910190840161118e565b50858210156111cb5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600081546111e881610fed565b808552602060018381168015611205576001811461121f5761124d565b60ff1985168884015283151560051b88018301955061124d565b866000528260002060005b858110156112455781548a820186015290830190840161122a565b890184019650505b505050505092915050565b60018060a01b038616815284602082015260a06040820152600061127f60a08301866111db565b841515606084015282810360808401526112998185610e06565b98975050505050505050565b6000602082840312156112b757600080fd5b8151610e5f81610f5f565b808201808211156112e357634e487b7160e01b600052601160045260246000fd5b92915050565b6060815260006112fc60608301866111db565b84602084015282810360408401526113148185610e06565b969550505050505056fea2646970667358221220af880c83ffb0885677c5f9109b9a5014493ee7146658327e0a2d8c209be06e9464736f6c63430008120033",
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

// CheckFeedback is a free data retrieval call binding the contract method 0x588579f9.
//
// Solidity: function checkFeedback() view returns(int256, string)
func (_EthAuction *EthAuctionCaller) CheckFeedback(opts *bind.CallOpts) (*big.Int, string, error) {
	var out []interface{}
	err := _EthAuction.contract.Call(opts, &out, "checkFeedback")

	if err != nil {
		return *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// CheckFeedback is a free data retrieval call binding the contract method 0x588579f9.
//
// Solidity: function checkFeedback() view returns(int256, string)
func (_EthAuction *EthAuctionSession) CheckFeedback() (*big.Int, string, error) {
	return _EthAuction.Contract.CheckFeedback(&_EthAuction.CallOpts)
}

// CheckFeedback is a free data retrieval call binding the contract method 0x588579f9.
//
// Solidity: function checkFeedback() view returns(int256, string)
func (_EthAuction *EthAuctionCallerSession) CheckFeedback() (*big.Int, string, error) {
	return _EthAuction.Contract.CheckFeedback(&_EthAuction.CallOpts)
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

// Abort is a paid mutator transaction binding the contract method 0x35acac4e.
//
// Solidity: function abort(string jsonString) returns()
func (_EthAuction *EthAuctionTransactor) Abort(opts *bind.TransactOpts, jsonString string) (*types.Transaction, error) {
	return _EthAuction.contract.Transact(opts, "abort", jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x35acac4e.
//
// Solidity: function abort(string jsonString) returns()
func (_EthAuction *EthAuctionSession) Abort(jsonString string) (*types.Transaction, error) {
	return _EthAuction.Contract.Abort(&_EthAuction.TransactOpts, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x35acac4e.
//
// Solidity: function abort(string jsonString) returns()
func (_EthAuction *EthAuctionTransactorSession) Abort(jsonString string) (*types.Transaction, error) {
	return _EthAuction.Contract.Abort(&_EthAuction.TransactOpts, jsonString)
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

// CloseAuction is a paid mutator transaction binding the contract method 0x45dad86c.
//
// Solidity: function closeAuction(bool not_winner_platform) returns()
func (_EthAuction *EthAuctionTransactor) CloseAuction(opts *bind.TransactOpts, not_winner_platform bool) (*types.Transaction, error) {
	return _EthAuction.contract.Transact(opts, "closeAuction", not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x45dad86c.
//
// Solidity: function closeAuction(bool not_winner_platform) returns()
func (_EthAuction *EthAuctionSession) CloseAuction(not_winner_platform bool) (*types.Transaction, error) {
	return _EthAuction.Contract.CloseAuction(&_EthAuction.TransactOpts, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x45dad86c.
//
// Solidity: function closeAuction(bool not_winner_platform) returns()
func (_EthAuction *EthAuctionTransactorSession) CloseAuction(not_winner_platform bool) (*types.Transaction, error) {
	return _EthAuction.Contract.CloseAuction(&_EthAuction.TransactOpts, not_winner_platform)
}

// Commit is a paid mutator transaction binding the contract method 0x9867db74.
//
// Solidity: function commit(string jsonString) returns()
func (_EthAuction *EthAuctionTransactor) Commit(opts *bind.TransactOpts, jsonString string) (*types.Transaction, error) {
	return _EthAuction.contract.Transact(opts, "commit", jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0x9867db74.
//
// Solidity: function commit(string jsonString) returns()
func (_EthAuction *EthAuctionSession) Commit(jsonString string) (*types.Transaction, error) {
	return _EthAuction.Contract.Commit(&_EthAuction.TransactOpts, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0x9867db74.
//
// Solidity: function commit(string jsonString) returns()
func (_EthAuction *EthAuctionTransactorSession) Commit(jsonString string) (*types.Transaction, error) {
	return _EthAuction.Contract.Commit(&_EthAuction.TransactOpts, jsonString)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0xcc5284e7.
//
// Solidity: function provide_feedback(int256 _score, string _feedback) returns()
func (_EthAuction *EthAuctionTransactor) ProvideFeedback(opts *bind.TransactOpts, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _EthAuction.contract.Transact(opts, "provide_feedback", _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0xcc5284e7.
//
// Solidity: function provide_feedback(int256 _score, string _feedback) returns()
func (_EthAuction *EthAuctionSession) ProvideFeedback(_score *big.Int, _feedback string) (*types.Transaction, error) {
	return _EthAuction.Contract.ProvideFeedback(&_EthAuction.TransactOpts, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0xcc5284e7.
//
// Solidity: function provide_feedback(int256 _score, string _feedback) returns()
func (_EthAuction *EthAuctionTransactorSession) ProvideFeedback(_score *big.Int, _feedback string) (*types.Transaction, error) {
	return _EthAuction.Contract.ProvideFeedback(&_EthAuction.TransactOpts, _score, _feedback)
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
	Winner     common.Address
	Amount     *big.Int
	Id         string
	Prcd       bool
	JsonString string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDecisionMade is a free log retrieval operation binding the contract event 0x9a51e3deb4a767f1ba793da152f8b62dde0121eb58730d9a86124e9b59afe978.
//
// Solidity: event DecisionMade(address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_EthAuction *EthAuctionFilterer) FilterDecisionMade(opts *bind.FilterOpts) (*EthAuctionDecisionMadeIterator, error) {

	logs, sub, err := _EthAuction.contract.FilterLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return &EthAuctionDecisionMadeIterator{contract: _EthAuction.contract, event: "DecisionMade", logs: logs, sub: sub}, nil
}

// WatchDecisionMade is a free log subscription operation binding the contract event 0x9a51e3deb4a767f1ba793da152f8b62dde0121eb58730d9a86124e9b59afe978.
//
// Solidity: event DecisionMade(address winner, uint256 amount, string id, bool prcd, string jsonString)
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

// ParseDecisionMade is a log parse operation binding the contract event 0x9a51e3deb4a767f1ba793da152f8b62dde0121eb58730d9a86124e9b59afe978.
//
// Solidity: event DecisionMade(address winner, uint256 amount, string id, bool prcd, string jsonString)
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

// EthAuctionRateAuctionIterator is returned from FilterRateAuction and is used to iterate over the raw logs and unpacked data for RateAuction events raised by the EthAuction contract.
type EthAuctionRateAuctionIterator struct {
	Event *EthAuctionRateAuction // Event containing the contract specifics and raw log

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
func (it *EthAuctionRateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthAuctionRateAuction)
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
		it.Event = new(EthAuctionRateAuction)
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
func (it *EthAuctionRateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthAuctionRateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthAuctionRateAuction represents a RateAuction event raised by the EthAuction contract.
type EthAuctionRateAuction struct {
	Id     string
	Rating *big.Int
	Review string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRateAuction is a free log retrieval operation binding the contract event 0x41ba38777fb1a414ddf3be769b0df73ab388f5f58a0edc435ce9736eaa9d758e.
//
// Solidity: event RateAuction(string id, int256 rating, string review)
func (_EthAuction *EthAuctionFilterer) FilterRateAuction(opts *bind.FilterOpts) (*EthAuctionRateAuctionIterator, error) {

	logs, sub, err := _EthAuction.contract.FilterLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return &EthAuctionRateAuctionIterator{contract: _EthAuction.contract, event: "RateAuction", logs: logs, sub: sub}, nil
}

// WatchRateAuction is a free log subscription operation binding the contract event 0x41ba38777fb1a414ddf3be769b0df73ab388f5f58a0edc435ce9736eaa9d758e.
//
// Solidity: event RateAuction(string id, int256 rating, string review)
func (_EthAuction *EthAuctionFilterer) WatchRateAuction(opts *bind.WatchOpts, sink chan<- *EthAuctionRateAuction) (event.Subscription, error) {

	logs, sub, err := _EthAuction.contract.WatchLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthAuctionRateAuction)
				if err := _EthAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
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

// ParseRateAuction is a log parse operation binding the contract event 0x41ba38777fb1a414ddf3be769b0df73ab388f5f58a0edc435ce9736eaa9d758e.
//
// Solidity: event RateAuction(string id, int256 rating, string review)
func (_EthAuction *EthAuctionFilterer) ParseRateAuction(log types.Log) (*EthAuctionRateAuction, error) {
	event := new(EthAuctionRateAuction)
	if err := _EthAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
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
