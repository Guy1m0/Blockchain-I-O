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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AwaitResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prcd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"DecisionMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"auctionType\",\"type\":\"string\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"rating\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"review\",\"type\":\"string\"}],\"name\":\"RateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_owner\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auction_type\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"checkAverageScore\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"not_winner_platform\",\"type\":\"bool\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auction_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_asset_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_asset_owner\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_score\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_feedback\",\"type\":\"string\"}],\"name\":\"provide_feedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200196438038062001964833981016040819052620000349162000090565b6001600160a01b038116608052600080546001600160a01b0319163317905560408051808201909152600f81526e22b733b634b9b41020bab1ba34b7b760891b602082015260019062000088908262000167565b505062000233565b600060208284031215620000a357600080fd5b81516001600160a01b0381168114620000bb57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620000ed57607f821691505b6020821081036200010e57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200016257600081815260208120601f850160051c810160208610156200013d5750805b601f850160051c820191505b818110156200015e5782815560010162000149565b5050505b505050565b81516001600160401b03811115620001835762000183620000c2565b6200019b81620001948454620000d8565b8462000114565b602080601f831160018114620001d35760008415620001ba5750858301515b600019600386901b1c1916600185901b1785556200015e565b600085815260208120601f198616915b828110156200020457888601518255948401946001909101908401620001e3565b5085821015620002235787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805161170e620002566000396000818161028901526105a7015261170e6000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638d10b0a611610097578063d0a1414a11610066578063d0a1414a14610256578063d960d57314610269578063e1081bb31461027c578063fc0c546a1461028457600080fd5b80638d10b0a6146101ef5780638da5cb5b14610202578063b14c63c514610215578063c84d2f6a1461024357600080fd5b8063451df52e116100d3578063451df52e14610175578063598647f8146101b65780635a4463a8146101c95780637efbf8ac146101dc57600080fd5b8063176321e9146101055780632e1a7d4d1461011a57806342d21ef71461014257806344a770bf14610162575b600080fd5b610118610113366004611075565b6102ab565b005b61012d6101283660046110bc565b6104a8565b60405190151581526020015b60405180910390f35b6101556101503660046110bc565b610686565b604051610139919061111b565b6101556101703660046110bc565b610720565b61019e6101833660046110bc565b6003602052600090815260409020546001600160a01b031681565b6040516001600160a01b039091168152602001610139565b6101186101c4366004611135565b610739565b6101186101d7366004611157565b610807565b6101186101ea3660046111c4565b6108f9565b6101556101fd3660046110bc565b610aea565b60005461019e906001600160a01b031681565b6102356102233660046110bc565b60046020526000908152604090205481565b604051908152602001610139565b61011861025136600461121b565b610b03565b610118610264366004611075565b610d3f565b6102356102773660046110bc565b610f01565b610155610fc5565b61019e7f000000000000000000000000000000000000000000000000000000000000000081565b60405165656e64696e6760d01b602082015260260160408051601f1981840301815282825280516020918201206000868152600583529290922091926102f2929101611285565b604051602081830303815290604052805190602001201461035a5760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064015b60405180910390fd5b6000828152600360205260409020546001600160a01b031633146103905760405162461bcd60e51b8152600401610351906112fb565b60405180604001604052806007815260200166636c6f73696e6760c81b8152506005600084815260200190815260200160002090816103cf919061137a565b506000828152600360209081526040808320546004835281842054600690935281842091517f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7946104349488946001600160a01b0390941693909290919088906114b7565b60405180910390a160008281526004602090815260408083205460038352818420546001600160a01b0316845260029092528220805491929091610479908490611521565b909155505050600090815260036020908152604080832080546001600160a01b03191690556004909152812055565b6040516337b832b760e11b602082015260009060240160408051601f1981840301815282825280516020918201206000868152600583529290922091926104f0929101611285565b60405160208183030381529060405280519060200120036105535760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420696e204f50454e207374617475730000000000000000006044820152606401610351565b336000908152600260205260409020548015610631573360008181526002602052604080822091909155516323b872dd60e01b81523060048201526024810191909152604481018290526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303816000875af11580156105f0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610614919061153a565b610631573360009081526002602052604081209190915592915050565b6000838152600660205260409081902090517f9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c916106759186919033908690611557565b60405180910390a150600192915050565b6005602052600090815260409020805461069f9061124b565b80601f01602080910402602001604051908101604052809291908181526020018280546106cb9061124b565b80156107185780601f106106ed57610100808354040283529160200191610718565b820191906000526020600020905b8154815290600101906020018083116106fb57829003601f168201915b505050505081565b6007602052600090815260409020805461069f9061124b565b600082815260046020526040902054156107905760008281526004602090815260408083205460038352818420546001600160a01b031684526002909252822080549192909161078a908490611521565b90915550505b600082815260036020908152604080832080546001600160a01b031916339081179091556004835281842085905560069092529182902091517fb95c1199e4385d3e33ad9cffdc96a2f61491b3798965fd93633c02ff7ade77f0926107fb928692869060019061158c565b60405180910390a15050565b6000546001600160a01b0316331461086b5760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79206f776e65722063616e20637265617465206e65772061756374696f6044820152603760f91b6064820152608401610351565b600083815260036020908152604080832080546001600160a01b03191690556004808352818420849055815180830183529081526337b832b760e11b818401528684526005909252909120906108c1908261137a565b5060008381526006602052604090206108da838261137a565b5060008381526007602052604090206108f3828261137a565b50505050565b60405166636c6f73696e6760c81b602082015260270160408051601f198184030181528282528051602091820120600087815260058352929092209192610941929101611285565b60405160208183030381529060405280519060200120146109a45760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e20434c4f53494e472073746174757300006044820152606401610351565b6000838152600360205260409020546001600160a01b031633146109da5760405162461bcd60e51b8152600401610351906112fb565b6000838152600760205260409081902090516009916109f891611285565b9081526040805191829003602090810183208054600181018255600091825282822001869055868152600790915220600891610a349190611285565b908152604051602091819003820190208054600181018255600091825291902001610a5f828261137a565b506000838152600660205260409081902090517fcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa791610aa491869190869086906115d9565b60405180910390a16040518060400160405280600681526020016518db1bdcd95960d21b8152506005600085815260200190815260200160002090816108f3919061137a565b6006602052600090815260409020805461069f9061124b565b6000546001600160a01b03163314610b6d5760405162461bcd60e51b815260206004820152602760248201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736044820152662073746174757360c81b6064820152608401610351565b6040516337b832b760e11b602082015260240160408051601f198184030181528282528051602091820120600086815260058352929092209192610bb2929101611285565b6040516020818303038152906040528051906020012014610c155760405162461bcd60e51b815260206004820152601b60248201527f436f6e7472616374206e6f7420696e204f50454e2073746174757300000000006044820152606401610351565b60405180604001604052806006815260200165656e64696e6760d01b815250600560008481526020019081526020016000209081610c53919061137a565b508080610c6c5750600082815260046020526040902054155b15610ced576040518060400160405280600681526020016518db1bdcd95960d21b815250600560008481526020019081526020016000209081610caf919061137a565b5060008281526004602090815260408083205460038352818420546001600160a01b0316845260029092528220805491929091610479908490611521565b6000828152600360209081526040918290205482518581526001600160a01b03909116918101919091527fa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b484746850027391016107fb565b60405165656e64696e6760d01b602082015260260160408051601f198184030181528282528051602091820120600086815260058352929092209192610d86929101611285565b6040516020818303038152906040528051906020012014610de95760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e47207374617475730000006044820152606401610351565b6000828152600360205260409020546001600160a01b03163314610e1f5760405162461bcd60e51b8152600401610351906112fb565b60405180604001604052806007815260200166636c6f73696e6760c81b815250600560008481526020019081526020016000209081610e5e919061137a565b5060008281526004602090815260408083205483546001600160a01b0316845260029092528220805491929091610e96908490611521565b9091555050600082815260036020908152604080832054600483528184205460069093529281902090517f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7936107fb9387936001600160a01b039092169290919060019088906114b7565b600081815260076020526040808220905182918291600991610f2291611285565b90815260405190819003602001902054905060005b81811015610fa657600085815260076020526040908190209051600991610f5d91611285565b90815260200160405180910390208181548110610f7c57610f7c611615565b906000526020600020015483610f92919061162b565b925080610f9e81611653565b915050610f37565b5080610fb383606461166c565b610fbd919061169c565b949350505050565b6001805461069f9061124b565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610ff957600080fd5b813567ffffffffffffffff8082111561101457611014610fd2565b604051601f8301601f19908116603f0116810190828211818310171561103c5761103c610fd2565b8160405283815286602085880101111561105557600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561108857600080fd5b82359150602083013567ffffffffffffffff8111156110a657600080fd5b6110b285828601610fe8565b9150509250929050565b6000602082840312156110ce57600080fd5b5035919050565b6000815180845260005b818110156110fb576020818501810151868301820152016110df565b506000602082860101526020601f19601f83011685010191505092915050565b60208152600061112e60208301846110d5565b9392505050565b6000806040838503121561114857600080fd5b50508035926020909101359150565b60008060006060848603121561116c57600080fd5b83359250602084013567ffffffffffffffff8082111561118b57600080fd5b61119787838801610fe8565b935060408601359150808211156111ad57600080fd5b506111ba86828701610fe8565b9150509250925092565b6000806000606084860312156111d957600080fd5b8335925060208401359150604084013567ffffffffffffffff8111156111fe57600080fd5b6111ba86828701610fe8565b801515811461121857600080fd5b50565b6000806040838503121561122e57600080fd5b8235915060208301356112408161120a565b809150509250929050565b600181811c9082168061125f57607f821691505b60208210810361127f57634e487b7160e01b600052602260045260246000fd5b50919050565b60008083546112938161124b565b600182811680156112ab57600181146112c0576112ef565b60ff19841687528215158302870194506112ef565b8760005260208060002060005b858110156112e65781548a8201529084019082016112cd565b50505082870194505b50929695505050505050565b6020808252601690820152754e6f7420617574686f72697a6564206163636573732160501b604082015260600190565b601f82111561137557600081815260208120601f850160051c810160208610156113525750805b601f850160051c820191505b818110156113715782815560010161135e565b5050505b505050565b815167ffffffffffffffff81111561139457611394610fd2565b6113a8816113a2845461124b565b8461132b565b602080601f8311600181146113dd57600084156113c55750858301515b600019600386901b1c1916600185901b178555611371565b600085815260208120601f198616915b8281101561140c578886015182559484019460019091019084016113ed565b508582101561142a5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600081546114478161124b565b808552602060018381168015611464576001811461147e576114ac565b60ff1985168884015283151560051b8801830195506114ac565b866000528260002060005b858110156114a45781548a8201860152908301908401611489565b890184019650505b505050505092915050565b86815260018060a01b038616602082015284604082015260c0606082015260006114e460c083018661143a565b841515608084015282810360a08401526114fe81856110d5565b9998505050505050505050565b634e487b7160e01b600052601160045260246000fd5b808201808211156115345761153461150b565b92915050565b60006020828403121561154c57600080fd5b815161112e8161120a565b848152608060208201526000611570608083018661143a565b6001600160a01b03949094166040830152506060015292915050565b85815260a0602082015260006115a560a083018761143a565b6001600160a01b03861660408401526060830185905282810360808401526115cd818561143a565b98975050505050505050565b8481526080602082015260006115f2608083018661143a565b846040840152828103606084015261160a81856110d5565b979650505050505050565b634e487b7160e01b600052603260045260246000fd5b808201828112600083128015821682158216171561164b5761164b61150b565b505092915050565b6000600182016116655761166561150b565b5060010190565b80820260008212600160ff1b841416156116885761168861150b565b81810583148215176115345761153461150b565b6000826116b957634e487b7160e01b600052601260045260246000fd5b600160ff1b8214600019841416156116d3576116d361150b565b50059056fea26469706673582212201488ce42735d1b4c3658195f7ef96d6f291c6c0c5ab9aa904650153b4e40862264736f6c63430008120033",
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

// AssetOwner is a free data retrieval call binding the contract method 0x44a770bf.
//
// Solidity: function asset_owner(uint256 ) view returns(string)
func (_EnglishAuction *EnglishAuctionCaller) AssetOwner(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _EnglishAuction.contract.Call(opts, &out, "asset_owner", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetOwner is a free data retrieval call binding the contract method 0x44a770bf.
//
// Solidity: function asset_owner(uint256 ) view returns(string)
func (_EnglishAuction *EnglishAuctionSession) AssetOwner(arg0 *big.Int) (string, error) {
	return _EnglishAuction.Contract.AssetOwner(&_EnglishAuction.CallOpts, arg0)
}

// AssetOwner is a free data retrieval call binding the contract method 0x44a770bf.
//
// Solidity: function asset_owner(uint256 ) view returns(string)
func (_EnglishAuction *EnglishAuctionCallerSession) AssetOwner(arg0 *big.Int) (string, error) {
	return _EnglishAuction.Contract.AssetOwner(&_EnglishAuction.CallOpts, arg0)
}

// AuctionType is a free data retrieval call binding the contract method 0xe1081bb3.
//
// Solidity: function auction_type() view returns(string)
func (_EnglishAuction *EnglishAuctionCaller) AuctionType(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EnglishAuction.contract.Call(opts, &out, "auction_type")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AuctionType is a free data retrieval call binding the contract method 0xe1081bb3.
//
// Solidity: function auction_type() view returns(string)
func (_EnglishAuction *EnglishAuctionSession) AuctionType() (string, error) {
	return _EnglishAuction.Contract.AuctionType(&_EnglishAuction.CallOpts)
}

// AuctionType is a free data retrieval call binding the contract method 0xe1081bb3.
//
// Solidity: function auction_type() view returns(string)
func (_EnglishAuction *EnglishAuctionCallerSession) AuctionType() (string, error) {
	return _EnglishAuction.Contract.AuctionType(&_EnglishAuction.CallOpts)
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

// Create is a paid mutator transaction binding the contract method 0x5a4463a8.
//
// Solidity: function create(uint256 _auction_id, string _asset_id, string _asset_owner) returns()
func (_EnglishAuction *EnglishAuctionTransactor) Create(opts *bind.TransactOpts, _auction_id *big.Int, _asset_id string, _asset_owner string) (*types.Transaction, error) {
	return _EnglishAuction.contract.Transact(opts, "create", _auction_id, _asset_id, _asset_owner)
}

// Create is a paid mutator transaction binding the contract method 0x5a4463a8.
//
// Solidity: function create(uint256 _auction_id, string _asset_id, string _asset_owner) returns()
func (_EnglishAuction *EnglishAuctionSession) Create(_auction_id *big.Int, _asset_id string, _asset_owner string) (*types.Transaction, error) {
	return _EnglishAuction.Contract.Create(&_EnglishAuction.TransactOpts, _auction_id, _asset_id, _asset_owner)
}

// Create is a paid mutator transaction binding the contract method 0x5a4463a8.
//
// Solidity: function create(uint256 _auction_id, string _asset_id, string _asset_owner) returns()
func (_EnglishAuction *EnglishAuctionTransactorSession) Create(_auction_id *big.Int, _asset_id string, _asset_owner string) (*types.Transaction, error) {
	return _EnglishAuction.Contract.Create(&_EnglishAuction.TransactOpts, _auction_id, _asset_id, _asset_owner)
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
	AuctionId *big.Int
	Winner    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAwaitResponse is a free log retrieval operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auctionId, address winner)
func (_EnglishAuction *EnglishAuctionFilterer) FilterAwaitResponse(opts *bind.FilterOpts) (*EnglishAuctionAwaitResponseIterator, error) {

	logs, sub, err := _EnglishAuction.contract.FilterLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return &EnglishAuctionAwaitResponseIterator{contract: _EnglishAuction.contract, event: "AwaitResponse", logs: logs, sub: sub}, nil
}

// WatchAwaitResponse is a free log subscription operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auctionId, address winner)
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
// Solidity: event AwaitResponse(uint256 auctionId, address winner)
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
	AuctionId  *big.Int
	Winner     common.Address
	Amount     *big.Int
	Id         string
	Prcd       bool
	JsonString string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDecisionMade is a free log retrieval operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auctionId, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_EnglishAuction *EnglishAuctionFilterer) FilterDecisionMade(opts *bind.FilterOpts) (*EnglishAuctionDecisionMadeIterator, error) {

	logs, sub, err := _EnglishAuction.contract.FilterLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return &EnglishAuctionDecisionMadeIterator{contract: _EnglishAuction.contract, event: "DecisionMade", logs: logs, sub: sub}, nil
}

// WatchDecisionMade is a free log subscription operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auctionId, address winner, uint256 amount, string id, bool prcd, string jsonString)
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
// Solidity: event DecisionMade(uint256 auctionId, address winner, uint256 amount, string id, bool prcd, string jsonString)
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
	AuctionId   *big.Int
	Id          string
	Bidder      common.Address
	BidAmount   *big.Int
	AuctionType string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0xb95c1199e4385d3e33ad9cffdc96a2f61491b3798965fd93633c02ff7ade77f0.
//
// Solidity: event HighestBidIncreased(uint256 auctionId, string id, address bidder, uint256 bidAmount, string auctionType)
func (_EnglishAuction *EnglishAuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*EnglishAuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _EnglishAuction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &EnglishAuctionHighestBidIncreasedIterator{contract: _EnglishAuction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0xb95c1199e4385d3e33ad9cffdc96a2f61491b3798965fd93633c02ff7ade77f0.
//
// Solidity: event HighestBidIncreased(uint256 auctionId, string id, address bidder, uint256 bidAmount, string auctionType)
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

// ParseHighestBidIncreased is a log parse operation binding the contract event 0xb95c1199e4385d3e33ad9cffdc96a2f61491b3798965fd93633c02ff7ade77f0.
//
// Solidity: event HighestBidIncreased(uint256 auctionId, string id, address bidder, uint256 bidAmount, string auctionType)
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
	AuctionId *big.Int
	Id        string
	Rating    *big.Int
	Review    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRateAuction is a free log retrieval operation binding the contract event 0xcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa7.
//
// Solidity: event RateAuction(uint256 auctionId, string id, int256 rating, string review)
func (_EnglishAuction *EnglishAuctionFilterer) FilterRateAuction(opts *bind.FilterOpts) (*EnglishAuctionRateAuctionIterator, error) {

	logs, sub, err := _EnglishAuction.contract.FilterLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return &EnglishAuctionRateAuctionIterator{contract: _EnglishAuction.contract, event: "RateAuction", logs: logs, sub: sub}, nil
}

// WatchRateAuction is a free log subscription operation binding the contract event 0xcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa7.
//
// Solidity: event RateAuction(uint256 auctionId, string id, int256 rating, string review)
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
// Solidity: event RateAuction(uint256 auctionId, string id, int256 rating, string review)
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
	AuctionId *big.Int
	Id        string
	Bidder    common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBid is a free log retrieval operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auctionId, string id, address bidder, uint256 amount)
func (_EnglishAuction *EnglishAuctionFilterer) FilterWithdrawBid(opts *bind.FilterOpts) (*EnglishAuctionWithdrawBidIterator, error) {

	logs, sub, err := _EnglishAuction.contract.FilterLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return &EnglishAuctionWithdrawBidIterator{contract: _EnglishAuction.contract, event: "WithdrawBid", logs: logs, sub: sub}, nil
}

// WatchWithdrawBid is a free log subscription operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auctionId, string id, address bidder, uint256 amount)
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
// Solidity: event WithdrawBid(uint256 auctionId, string id, address bidder, uint256 amount)
func (_EnglishAuction *EnglishAuctionFilterer) ParseWithdrawBid(log types.Log) (*EnglishAuctionWithdrawBid, error) {
	event := new(EnglishAuctionWithdrawBid)
	if err := _EnglishAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
