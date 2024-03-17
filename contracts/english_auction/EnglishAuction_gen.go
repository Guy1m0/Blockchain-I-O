// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package english_auction

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

// EnglishAuctionMetaData contains all meta data concerning the EnglishAuction contract.
var EnglishAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AwaitResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prcd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"DecisionMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"rating\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"review\",\"type\":\"string\"}],\"name\":\"RateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auction_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"checkAverageScore\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"not_winner_platform\",\"type\":\"bool\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auction_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_asset_id\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_score\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_feedback\",\"type\":\"string\"}],\"name\":\"provide_feedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161188038038061188083398101604081905261002f91610052565b6001600160a01b0316608052600080546001600160a01b03191633179055610082565b60006020828403121561006457600080fd5b81516001600160a01b038116811461007b57600080fd5b9392505050565b6080516117d56100ab6000396000818161028c01528181610681015261092401526117d56000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80638d10b0a611610097578063c84d2f6a11610066578063c84d2f6a1461024e578063d0a1414a14610261578063d960d57314610274578063fc0c546a1461028757600080fd5b80638d10b0a6146101d15780638da5cb5b146101e4578063b14c63c5146101f7578063b8ea6bf61461022557600080fd5b806342d21ef7116100d357806342d21ef71461014a578063451df52e1461016a578063598647f8146101ab5780637efbf8ac146101be57600080fd5b80630118fa49146100fa578063176321e91461010f5780632e1a7d4d14610122575b600080fd5b61010d6101083660046111ed565b6102ae565b005b61010d61011d3660046111ed565b61038a565b610135610130366004611234565b610582565b60405190151581526020015b60405180910390f35b61015d610158366004611234565b610760565b6040516101419190611293565b610193610178366004611234565b6002602052600090815260409020546001600160a01b031681565b6040516001600160a01b039091168152602001610141565b61010d6101b93660046112ad565b6107fa565b61010d6101cc3660046112cf565b610aad565b61015d6101df366004611234565b610c7b565b600054610193906001600160a01b031681565b610217610205366004611234565b60036020526000908152604090205481565b604051908152602001610141565b610193610233366004611234565b6006602052600090815260409020546001600160a01b031681565b61010d61025c366004611330565b610c94565b61010d61026f3660046111ed565b610ed8565b610217610282366004611234565b61109a565b6101937f000000000000000000000000000000000000000000000000000000000000000081565b6000546001600160a01b031633146103175760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79206f776e65722063616e20637265617465206e65772061756374696f6044820152603760f91b60648201526084015b60405180910390fd5b600082815260026020908152604080832080546001600160a01b0319169055600382528083208390558051808201825260048082526337b832b760e11b828501528685529092529091209061036c90826113e8565b50600082815260056020526040902061038582826113e8565b505050565b60405165656e64696e6760d01b602082015260260160408051601f1981840301815282825280516020918201206000868152600483529290922091926103d19291016114a8565b60405160208183030381529060405280519060200120146104345760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e4720737461747573000000604482015260640161030e565b6000828152600260205260409020546001600160a01b0316331461046a5760405162461bcd60e51b815260040161030e9061151e565b60405180604001604052806007815260200166636c6f73696e6760c81b8152506004600084815260200190815260200160002090816104a991906113e8565b506000828152600260209081526040808320546003835281842054600590935281842091517f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c79461050e9488946001600160a01b0390941693909290919088906115cb565b60405180910390a160008281526003602090815260408083205460028352818420546001600160a01b0316845260019092528220805491929091610553908490611635565b909155505050600090815260026020908152604080832080546001600160a01b03191690556003909152812055565b6040516337b832b760e11b602082015260009060240160408051601f1981840301815282825280516020918201206000868152600483529290922091926105ca9291016114a8565b604051602081830303815290604052805190602001200361062d5760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420696e204f50454e20737461747573000000000000000000604482015260640161030e565b33600090815260016020526040902054801561070b573360008181526001602052604080822091909155516323b872dd60e01b81523060048201526024810191909152604481018290526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303816000875af11580156106ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ee919061164e565b61070b573360009081526001602052604081209190915592915050565b6000838152600560205260409081902090517f9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c9161074f918691903390869061166b565b60405180910390a150600192915050565b6004602052600090815260409020805461077990611360565b80601f01602080910402602001604051908101604052809291908181526020018280546107a590611360565b80156107f25780601f106107c7576101008083540402835291602001916107f2565b820191906000526020600020905b8154815290600101906020018083116107d557829003601f168201915b505050505081565b6040516337b832b760e11b602082015260240160408051601f19818403018152828252805160209182012060008681526004835292909220919261083f9291016114a8565b60405160208183030381529060405280519060200120146108a25760405162461bcd60e51b815260206004820152601b60248201527f436f6e7472616374206e6f7420696e204f50454e207374617475730000000000604482015260640161030e565b60008281526003602052604090205481116108ff5760405162461bcd60e51b815260206004820152601e60248201527f546865726520616c7265616479206973206120686967686572206269642e0000604482015260640161030e565b6040516323b872dd60e01b8152336004820152306024820152604481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610975573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610999919061164e565b9050806109e15760405162461bcd60e51b81526020600482015260166024820152752a37b5b2b7103a3930b739b332b9103330b4b632b21760511b604482015260640161030e565b60008381526003602052604090205415610a385760008381526003602090815260408083205460028352818420546001600160a01b0316845260019092528220805491929091610a32908490611635565b90915550505b600083815260026020908152604080832080546001600160a01b031916339081179091556003835281842086905560059092529182902091517f463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde0892610aa0928792879061166b565b60405180910390a1505050565b60405166636c6f73696e6760c81b602082015260270160408051601f198184030181528282528051602091820120600087815260048352929092209192610af59291016114a8565b6040516020818303038152906040528051906020012014610b585760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e20434c4f53494e47207374617475730000604482015260640161030e565b6000838152600260205260409020546001600160a01b03163314610b8e5760405162461bcd60e51b815260040161030e9061151e565b600083815260066020908152604080832080546001600160a01b0390811685526008845282852080546001818101835591875285872001889055915416845260078352908320805491820181558352912001610bea82826113e8565b506000838152600560205260409081902090517fcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa791610c2f91869190869086906116a0565b60405180910390a16040518060400160405280600681526020016518db1bdcd95960d21b815250600460008581526020019081526020016000209081610c7591906113e8565b50505050565b6005602052600090815260409020805461077990611360565b6000546001600160a01b03163314610cfe5760405162461bcd60e51b815260206004820152602760248201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736044820152662073746174757360c81b606482015260840161030e565b6040516337b832b760e11b602082015260240160408051601f198184030181528282528051602091820120600086815260048352929092209192610d439291016114a8565b6040516020818303038152906040528051906020012014610da65760405162461bcd60e51b815260206004820152601b60248201527f436f6e7472616374206e6f7420696e204f50454e207374617475730000000000604482015260640161030e565b60405180604001604052806006815260200165656e64696e6760d01b815250600460008481526020019081526020016000209081610de491906113e8565b508080610dfd5750600082815260036020526040902054155b15610e7e576040518060400160405280600681526020016518db1bdcd95960d21b815250600460008481526020019081526020016000209081610e4091906113e8565b5060008281526003602090815260408083205460028352818420546001600160a01b0316845260019092528220805491929091610553908490611635565b6000828152600260209081526040918290205482518581526001600160a01b03909116918101919091527fa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b484746850027391015b60405180910390a15050565b60405165656e64696e6760d01b602082015260260160408051601f198184030181528282528051602091820120600086815260048352929092209192610f1f9291016114a8565b6040516020818303038152906040528051906020012014610f825760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e4720737461747573000000604482015260640161030e565b6000828152600260205260409020546001600160a01b03163314610fb85760405162461bcd60e51b815260040161030e9061151e565b60405180604001604052806007815260200166636c6f73696e6760c81b815250600460008481526020019081526020016000209081610ff791906113e8565b5060008281526003602090815260408083205483546001600160a01b031684526001909252822080549192909161102f908490611635565b9091555050600082815260026020908152604080832054600383528184205460059093529281902090517f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c793610ecc9387936001600160a01b039092169290919060019088906115cb565b6000818152600660209081526040808320546001600160a01b0316835260089091528120548190815b8181101561112b576000858152600660209081526040808320546001600160a01b0316835260089091529020805482908110611101576111016116dc565b90600052602060002001548361111791906116f2565b9250806111238161171a565b9150506110c3565b5080611138836064611733565b6111429190611763565b949350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261117157600080fd5b813567ffffffffffffffff8082111561118c5761118c61114a565b604051601f8301601f19908116603f011681019082821181831017156111b4576111b461114a565b816040528381528660208588010111156111cd57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561120057600080fd5b82359150602083013567ffffffffffffffff81111561121e57600080fd5b61122a85828601611160565b9150509250929050565b60006020828403121561124657600080fd5b5035919050565b6000815180845260005b8181101561127357602081850181015186830182015201611257565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006112a6602083018461124d565b9392505050565b600080604083850312156112c057600080fd5b50508035926020909101359150565b6000806000606084860312156112e457600080fd5b8335925060208401359150604084013567ffffffffffffffff81111561130957600080fd5b61131586828701611160565b9150509250925092565b801515811461132d57600080fd5b50565b6000806040838503121561134357600080fd5b8235915060208301356113558161131f565b809150509250929050565b600181811c9082168061137457607f821691505b60208210810361139457634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561038557600081815260208120601f850160051c810160208610156113c15750805b601f850160051c820191505b818110156113e0578281556001016113cd565b505050505050565b815167ffffffffffffffff8111156114025761140261114a565b611416816114108454611360565b8461139a565b602080601f83116001811461144b57600084156114335750858301515b600019600386901b1c1916600185901b1785556113e0565b600085815260208120601f198616915b8281101561147a5788860151825594840194600190910190840161145b565b50858210156114985787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008083546114b681611360565b600182811680156114ce57600181146114e357611512565b60ff1984168752821515830287019450611512565b8760005260208060002060005b858110156115095781548a8201529084019082016114f0565b50505082870194505b50929695505050505050565b6020808252601690820152754e6f7420617574686f72697a6564206163636573732160501b604082015260600190565b6000815461155b81611360565b8085526020600183811680156115785760018114611592576115c0565b60ff1985168884015283151560051b8801830195506115c0565b866000528260002060005b858110156115b85781548a820186015290830190840161159d565b890184019650505b505050505092915050565b86815260018060a01b038616602082015284604082015260c0606082015260006115f860c083018661154e565b841515608084015282810360a0840152611612818561124d565b9998505050505050505050565b634e487b7160e01b600052601160045260246000fd5b808201808211156116485761164861161f565b92915050565b60006020828403121561166057600080fd5b81516112a68161131f565b848152608060208201526000611684608083018661154e565b6001600160a01b03949094166040830152506060015292915050565b8481526080602082015260006116b9608083018661154e565b84604084015282810360608401526116d1818561124d565b979650505050505050565b634e487b7160e01b600052603260045260246000fd5b80820182811260008312801582168215821617156117125761171261161f565b505092915050565b60006001820161172c5761172c61161f565b5060010190565b80820260008212600160ff1b8414161561174f5761174f61161f565b81810583148215176116485761164861161f565b60008261178057634e487b7160e01b600052601260045260246000fd5b600160ff1b82146000198414161561179a5761179a61161f565b50059056fea26469706673582212204396192c4151be44c503068cd04a36b59ded1274a7535b239e27732c6352ff3e64736f6c63430008120033",
}

// EnglishAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use EnglishAuctionMetaData.ABI instead.
var EnglishAuctionABI = EnglishAuctionMetaData.ABI

// EnglishAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnglishAuctionMetaData.Bin instead.
var EnglishAuctionBin = EnglishAuctionMetaData.Bin

// DeployEnglishAuction deploys a new Ethereum contract, binding an instance of EnglishAuction to it.
func DeployEnglishAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *EnglishAuction, error) {
	parsed, err := EnglishAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnglishAuctionBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnglishAuction{EnglishAuctionCaller: EnglishAuctionCaller{contract: contract}, EnglishAuctionTransactor: EnglishAuctionTransactor{contract: contract}, EnglishAuctionFilterer: EnglishAuctionFilterer{contract: contract}}, nil
}

// EnglishAuction is an auto generated Go binding around an Ethereum contract.
type EnglishAuction struct {
	EnglishAuctionCaller     // Read-only binding to the contract
	EnglishAuctionTransactor // Write-only binding to the contract
	EnglishAuctionFilterer   // Log filterer for contract events
}

// EnglishAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnglishAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnglishAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnglishAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnglishAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnglishAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnglishAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnglishAuctionSession struct {
	Contract     *EnglishAuction   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnglishAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnglishAuctionCallerSession struct {
	Contract *EnglishAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EnglishAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnglishAuctionTransactorSession struct {
	Contract     *EnglishAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EnglishAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnglishAuctionRaw struct {
	Contract *EnglishAuction // Generic contract binding to access the raw methods on
}

// EnglishAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnglishAuctionCallerRaw struct {
	Contract *EnglishAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// EnglishAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnglishAuctionTransactorRaw struct {
	Contract *EnglishAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnglishAuction creates a new instance of EnglishAuction, bound to a specific deployed contract.
func NewEnglishAuction(address common.Address, backend bind.ContractBackend) (*EnglishAuction, error) {
	contract, err := bindEnglishAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnglishAuction{EnglishAuctionCaller: EnglishAuctionCaller{contract: contract}, EnglishAuctionTransactor: EnglishAuctionTransactor{contract: contract}, EnglishAuctionFilterer: EnglishAuctionFilterer{contract: contract}}, nil
}

// NewEnglishAuctionCaller creates a new read-only instance of EnglishAuction, bound to a specific deployed contract.
func NewEnglishAuctionCaller(address common.Address, caller bind.ContractCaller) (*EnglishAuctionCaller, error) {
	contract, err := bindEnglishAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnglishAuctionCaller{contract: contract}, nil
}

// NewEnglishAuctionTransactor creates a new write-only instance of EnglishAuction, bound to a specific deployed contract.
func NewEnglishAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*EnglishAuctionTransactor, error) {
	contract, err := bindEnglishAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnglishAuctionTransactor{contract: contract}, nil
}

// NewEnglishAuctionFilterer creates a new log filterer instance of EnglishAuction, bound to a specific deployed contract.
func NewEnglishAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*EnglishAuctionFilterer, error) {
	contract, err := bindEnglishAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnglishAuctionFilterer{contract: contract}, nil
}

// bindEnglishAuction binds a generic wrapper to an already deployed contract.
func bindEnglishAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnglishAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnglishAuction *EnglishAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnglishAuction.Contract.EnglishAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnglishAuction *EnglishAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnglishAuction.Contract.EnglishAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnglishAuction *EnglishAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnglishAuction.Contract.EnglishAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnglishAuction *EnglishAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnglishAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnglishAuction *EnglishAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnglishAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnglishAuction *EnglishAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnglishAuction.Contract.contract.Transact(opts, method, params...)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EnglishAuction *EnglishAuctionCaller) AssetId(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _EnglishAuction.contract.Call(opts, &out, "asset_id", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EnglishAuction *EnglishAuctionSession) AssetId(arg0 *big.Int) (string, error) {
	return _EnglishAuction.Contract.AssetId(&_EnglishAuction.CallOpts, arg0)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EnglishAuction *EnglishAuctionCallerSession) AssetId(arg0 *big.Int) (string, error) {
	return _EnglishAuction.Contract.AssetId(&_EnglishAuction.CallOpts, arg0)
}

// AuctionOwner is a free data retrieval call binding the contract method 0xb8ea6bf6.
//
// Solidity: function auction_owner(uint256 ) view returns(address)
func (_EnglishAuction *EnglishAuctionCaller) AuctionOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EnglishAuction.contract.Call(opts, &out, "auction_owner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuctionOwner is a free data retrieval call binding the contract method 0xb8ea6bf6.
//
// Solidity: function auction_owner(uint256 ) view returns(address)
func (_EnglishAuction *EnglishAuctionSession) AuctionOwner(arg0 *big.Int) (common.Address, error) {
	return _EnglishAuction.Contract.AuctionOwner(&_EnglishAuction.CallOpts, arg0)
}

// AuctionOwner is a free data retrieval call binding the contract method 0xb8ea6bf6.
//
// Solidity: function auction_owner(uint256 ) view returns(address)
func (_EnglishAuction *EnglishAuctionCallerSession) AuctionOwner(arg0 *big.Int) (common.Address, error) {
	return _EnglishAuction.Contract.AuctionOwner(&_EnglishAuction.CallOpts, arg0)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_EnglishAuction *EnglishAuctionCaller) CheckAverageScore(opts *bind.CallOpts, auctionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EnglishAuction.contract.Call(opts, &out, "checkAverageScore", auctionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_EnglishAuction *EnglishAuctionSession) CheckAverageScore(auctionId *big.Int) (*big.Int, error) {
	return _EnglishAuction.Contract.CheckAverageScore(&_EnglishAuction.CallOpts, auctionId)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_EnglishAuction *EnglishAuctionCallerSession) CheckAverageScore(auctionId *big.Int) (*big.Int, error) {
	return _EnglishAuction.Contract.CheckAverageScore(&_EnglishAuction.CallOpts, auctionId)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EnglishAuction *EnglishAuctionCaller) HighestBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EnglishAuction.contract.Call(opts, &out, "highestBid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EnglishAuction *EnglishAuctionSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _EnglishAuction.Contract.HighestBid(&_EnglishAuction.CallOpts, arg0)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EnglishAuction *EnglishAuctionCallerSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _EnglishAuction.Contract.HighestBid(&_EnglishAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EnglishAuction *EnglishAuctionCaller) HighestBidder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EnglishAuction.contract.Call(opts, &out, "highestBidder", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EnglishAuction *EnglishAuctionSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _EnglishAuction.Contract.HighestBidder(&_EnglishAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EnglishAuction *EnglishAuctionCallerSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _EnglishAuction.Contract.HighestBidder(&_EnglishAuction.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnglishAuction *EnglishAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnglishAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnglishAuction *EnglishAuctionSession) Owner() (common.Address, error) {
	return _EnglishAuction.Contract.Owner(&_EnglishAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnglishAuction *EnglishAuctionCallerSession) Owner() (common.Address, error) {
	return _EnglishAuction.Contract.Owner(&_EnglishAuction.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EnglishAuction *EnglishAuctionCaller) Status(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _EnglishAuction.contract.Call(opts, &out, "status", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EnglishAuction *EnglishAuctionSession) Status(arg0 *big.Int) (string, error) {
	return _EnglishAuction.Contract.Status(&_EnglishAuction.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EnglishAuction *EnglishAuctionCallerSession) Status(arg0 *big.Int) (string, error) {
	return _EnglishAuction.Contract.Status(&_EnglishAuction.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EnglishAuction *EnglishAuctionCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnglishAuction.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EnglishAuction *EnglishAuctionSession) Token() (common.Address, error) {
	return _EnglishAuction.Contract.Token(&_EnglishAuction.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EnglishAuction *EnglishAuctionCallerSession) Token() (common.Address, error) {
	return _EnglishAuction.Contract.Token(&_EnglishAuction.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EnglishAuction *EnglishAuctionTransactor) Abort(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EnglishAuction.contract.Transact(opts, "abort", auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EnglishAuction *EnglishAuctionSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EnglishAuction.Contract.Abort(&_EnglishAuction.TransactOpts, auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EnglishAuction *EnglishAuctionTransactorSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EnglishAuction.Contract.Abort(&_EnglishAuction.TransactOpts, auctionId, jsonString)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 auctionId, uint256 bidAmount) returns()
func (_EnglishAuction *EnglishAuctionTransactor) Bid(opts *bind.TransactOpts, auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EnglishAuction.contract.Transact(opts, "bid", auctionId, bidAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 auctionId, uint256 bidAmount) returns()
func (_EnglishAuction *EnglishAuctionSession) Bid(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EnglishAuction.Contract.Bid(&_EnglishAuction.TransactOpts, auctionId, bidAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 auctionId, uint256 bidAmount) returns()
func (_EnglishAuction *EnglishAuctionTransactorSession) Bid(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EnglishAuction.Contract.Bid(&_EnglishAuction.TransactOpts, auctionId, bidAmount)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_EnglishAuction *EnglishAuctionTransactor) CloseAuction(opts *bind.TransactOpts, auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _EnglishAuction.contract.Transact(opts, "closeAuction", auctionId, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_EnglishAuction *EnglishAuctionSession) CloseAuction(auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _EnglishAuction.Contract.CloseAuction(&_EnglishAuction.TransactOpts, auctionId, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_EnglishAuction *EnglishAuctionTransactorSession) CloseAuction(auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _EnglishAuction.Contract.CloseAuction(&_EnglishAuction.TransactOpts, auctionId, not_winner_platform)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EnglishAuction *EnglishAuctionTransactor) Commit(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EnglishAuction.contract.Transact(opts, "commit", auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EnglishAuction *EnglishAuctionSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EnglishAuction.Contract.Commit(&_EnglishAuction.TransactOpts, auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EnglishAuction *EnglishAuctionTransactorSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EnglishAuction.Contract.Commit(&_EnglishAuction.TransactOpts, auctionId, jsonString)
}

// Create is a paid mutator transaction binding the contract method 0x0118fa49.
//
// Solidity: function create(uint256 _auction_id, string _asset_id) returns()
func (_EnglishAuction *EnglishAuctionTransactor) Create(opts *bind.TransactOpts, _auction_id *big.Int, _asset_id string) (*types.Transaction, error) {
	return _EnglishAuction.contract.Transact(opts, "create", _auction_id, _asset_id)
}

// Create is a paid mutator transaction binding the contract method 0x0118fa49.
//
// Solidity: function create(uint256 _auction_id, string _asset_id) returns()
func (_EnglishAuction *EnglishAuctionSession) Create(_auction_id *big.Int, _asset_id string) (*types.Transaction, error) {
	return _EnglishAuction.Contract.Create(&_EnglishAuction.TransactOpts, _auction_id, _asset_id)
}

// Create is a paid mutator transaction binding the contract method 0x0118fa49.
//
// Solidity: function create(uint256 _auction_id, string _asset_id) returns()
func (_EnglishAuction *EnglishAuctionTransactorSession) Create(_auction_id *big.Int, _asset_id string) (*types.Transaction, error) {
	return _EnglishAuction.Contract.Create(&_EnglishAuction.TransactOpts, _auction_id, _asset_id)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_EnglishAuction *EnglishAuctionTransactor) ProvideFeedback(opts *bind.TransactOpts, auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _EnglishAuction.contract.Transact(opts, "provide_feedback", auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_EnglishAuction *EnglishAuctionSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _EnglishAuction.Contract.ProvideFeedback(&_EnglishAuction.TransactOpts, auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_EnglishAuction *EnglishAuctionTransactorSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _EnglishAuction.Contract.ProvideFeedback(&_EnglishAuction.TransactOpts, auctionId, _score, _feedback)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EnglishAuction *EnglishAuctionTransactor) Withdraw(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _EnglishAuction.contract.Transact(opts, "withdraw", auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EnglishAuction *EnglishAuctionSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _EnglishAuction.Contract.Withdraw(&_EnglishAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EnglishAuction *EnglishAuctionTransactorSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _EnglishAuction.Contract.Withdraw(&_EnglishAuction.TransactOpts, auctionId)
}

// EnglishAuctionAwaitResponseIterator is returned from FilterAwaitResponse and is used to iterate over the raw logs and unpacked data for AwaitResponse events raised by the EnglishAuction contract.
type EnglishAuctionAwaitResponseIterator struct {
	Event *EnglishAuctionAwaitResponse // Event containing the contract specifics and raw log

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
func (it *EnglishAuctionAwaitResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnglishAuctionAwaitResponse)
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
		it.Event = new(EnglishAuctionAwaitResponse)
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
func (it *EnglishAuctionAwaitResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnglishAuctionAwaitResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnglishAuctionAwaitResponse represents a AwaitResponse event raised by the EnglishAuction contract.
type EnglishAuctionAwaitResponse struct {
	Auction *big.Int
	Winner  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAwaitResponse is a free log retrieval operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_EnglishAuction *EnglishAuctionFilterer) FilterAwaitResponse(opts *bind.FilterOpts) (*EnglishAuctionAwaitResponseIterator, error) {

	logs, sub, err := _EnglishAuction.contract.FilterLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return &EnglishAuctionAwaitResponseIterator{contract: _EnglishAuction.contract, event: "AwaitResponse", logs: logs, sub: sub}, nil
}

// WatchAwaitResponse is a free log subscription operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_EnglishAuction *EnglishAuctionFilterer) WatchAwaitResponse(opts *bind.WatchOpts, sink chan<- *EnglishAuctionAwaitResponse) (event.Subscription, error) {

	logs, sub, err := _EnglishAuction.contract.WatchLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnglishAuctionAwaitResponse)
				if err := _EnglishAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
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

// ParseAwaitResponse is a log parse operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_EnglishAuction *EnglishAuctionFilterer) ParseAwaitResponse(log types.Log) (*EnglishAuctionAwaitResponse, error) {
	event := new(EnglishAuctionAwaitResponse)
	if err := _EnglishAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnglishAuctionDecisionMadeIterator is returned from FilterDecisionMade and is used to iterate over the raw logs and unpacked data for DecisionMade events raised by the EnglishAuction contract.
type EnglishAuctionDecisionMadeIterator struct {
	Event *EnglishAuctionDecisionMade // Event containing the contract specifics and raw log

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
func (it *EnglishAuctionDecisionMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnglishAuctionDecisionMade)
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
		it.Event = new(EnglishAuctionDecisionMade)
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
func (it *EnglishAuctionDecisionMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnglishAuctionDecisionMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnglishAuctionDecisionMade represents a DecisionMade event raised by the EnglishAuction contract.
type EnglishAuctionDecisionMade struct {
	Auction    *big.Int
	Winner     common.Address
	Amount     *big.Int
	Id         string
	Prcd       bool
	JsonString string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDecisionMade is a free log retrieval operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auction, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_EnglishAuction *EnglishAuctionFilterer) FilterDecisionMade(opts *bind.FilterOpts) (*EnglishAuctionDecisionMadeIterator, error) {

	logs, sub, err := _EnglishAuction.contract.FilterLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return &EnglishAuctionDecisionMadeIterator{contract: _EnglishAuction.contract, event: "DecisionMade", logs: logs, sub: sub}, nil
}

// WatchDecisionMade is a free log subscription operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auction, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_EnglishAuction *EnglishAuctionFilterer) WatchDecisionMade(opts *bind.WatchOpts, sink chan<- *EnglishAuctionDecisionMade) (event.Subscription, error) {

	logs, sub, err := _EnglishAuction.contract.WatchLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnglishAuctionDecisionMade)
				if err := _EnglishAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
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

// ParseDecisionMade is a log parse operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auction, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_EnglishAuction *EnglishAuctionFilterer) ParseDecisionMade(log types.Log) (*EnglishAuctionDecisionMade, error) {
	event := new(EnglishAuctionDecisionMade)
	if err := _EnglishAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnglishAuctionHighestBidIncreasedIterator is returned from FilterHighestBidIncreased and is used to iterate over the raw logs and unpacked data for HighestBidIncreased events raised by the EnglishAuction contract.
type EnglishAuctionHighestBidIncreasedIterator struct {
	Event *EnglishAuctionHighestBidIncreased // Event containing the contract specifics and raw log

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
func (it *EnglishAuctionHighestBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnglishAuctionHighestBidIncreased)
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
		it.Event = new(EnglishAuctionHighestBidIncreased)
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
func (it *EnglishAuctionHighestBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnglishAuctionHighestBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnglishAuctionHighestBidIncreased represents a HighestBidIncreased event raised by the EnglishAuction contract.
type EnglishAuctionHighestBidIncreased struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_EnglishAuction *EnglishAuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*EnglishAuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _EnglishAuction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &EnglishAuctionHighestBidIncreasedIterator{contract: _EnglishAuction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_EnglishAuction *EnglishAuctionFilterer) WatchHighestBidIncreased(opts *bind.WatchOpts, sink chan<- *EnglishAuctionHighestBidIncreased) (event.Subscription, error) {

	logs, sub, err := _EnglishAuction.contract.WatchLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnglishAuctionHighestBidIncreased)
				if err := _EnglishAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
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

// ParseHighestBidIncreased is a log parse operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_EnglishAuction *EnglishAuctionFilterer) ParseHighestBidIncreased(log types.Log) (*EnglishAuctionHighestBidIncreased, error) {
	event := new(EnglishAuctionHighestBidIncreased)
	if err := _EnglishAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnglishAuctionRateAuctionIterator is returned from FilterRateAuction and is used to iterate over the raw logs and unpacked data for RateAuction events raised by the EnglishAuction contract.
type EnglishAuctionRateAuctionIterator struct {
	Event *EnglishAuctionRateAuction // Event containing the contract specifics and raw log

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
func (it *EnglishAuctionRateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnglishAuctionRateAuction)
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
		it.Event = new(EnglishAuctionRateAuction)
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
func (it *EnglishAuctionRateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnglishAuctionRateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnglishAuctionRateAuction represents a RateAuction event raised by the EnglishAuction contract.
type EnglishAuctionRateAuction struct {
	Auction *big.Int
	Id      string
	Rating  *big.Int
	Review  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRateAuction is a free log retrieval operation binding the contract event 0xcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa7.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, string review)
func (_EnglishAuction *EnglishAuctionFilterer) FilterRateAuction(opts *bind.FilterOpts) (*EnglishAuctionRateAuctionIterator, error) {

	logs, sub, err := _EnglishAuction.contract.FilterLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return &EnglishAuctionRateAuctionIterator{contract: _EnglishAuction.contract, event: "RateAuction", logs: logs, sub: sub}, nil
}

// WatchRateAuction is a free log subscription operation binding the contract event 0xcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa7.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, string review)
func (_EnglishAuction *EnglishAuctionFilterer) WatchRateAuction(opts *bind.WatchOpts, sink chan<- *EnglishAuctionRateAuction) (event.Subscription, error) {

	logs, sub, err := _EnglishAuction.contract.WatchLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnglishAuctionRateAuction)
				if err := _EnglishAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
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

// ParseRateAuction is a log parse operation binding the contract event 0xcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa7.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, string review)
func (_EnglishAuction *EnglishAuctionFilterer) ParseRateAuction(log types.Log) (*EnglishAuctionRateAuction, error) {
	event := new(EnglishAuctionRateAuction)
	if err := _EnglishAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnglishAuctionWithdrawBidIterator is returned from FilterWithdrawBid and is used to iterate over the raw logs and unpacked data for WithdrawBid events raised by the EnglishAuction contract.
type EnglishAuctionWithdrawBidIterator struct {
	Event *EnglishAuctionWithdrawBid // Event containing the contract specifics and raw log

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
func (it *EnglishAuctionWithdrawBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnglishAuctionWithdrawBid)
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
		it.Event = new(EnglishAuctionWithdrawBid)
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
func (it *EnglishAuctionWithdrawBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnglishAuctionWithdrawBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnglishAuctionWithdrawBid represents a WithdrawBid event raised by the EnglishAuction contract.
type EnglishAuctionWithdrawBid struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBid is a free log retrieval operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_EnglishAuction *EnglishAuctionFilterer) FilterWithdrawBid(opts *bind.FilterOpts) (*EnglishAuctionWithdrawBidIterator, error) {

	logs, sub, err := _EnglishAuction.contract.FilterLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return &EnglishAuctionWithdrawBidIterator{contract: _EnglishAuction.contract, event: "WithdrawBid", logs: logs, sub: sub}, nil
}

// WatchWithdrawBid is a free log subscription operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_EnglishAuction *EnglishAuctionFilterer) WatchWithdrawBid(opts *bind.WatchOpts, sink chan<- *EnglishAuctionWithdrawBid) (event.Subscription, error) {

	logs, sub, err := _EnglishAuction.contract.WatchLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnglishAuctionWithdrawBid)
				if err := _EnglishAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
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

// ParseWithdrawBid is a log parse operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_EnglishAuction *EnglishAuctionFilterer) ParseWithdrawBid(log types.Log) (*EnglishAuctionWithdrawBid, error) {
	event := new(EnglishAuctionWithdrawBid)
	if err := _EnglishAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
