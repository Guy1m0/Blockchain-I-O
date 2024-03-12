// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_closed_bid_auction

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

// EthClosedBidAuctionMetaData contains all meta data concerning the EthClosedBidAuction contract.
var EthClosedBidAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AwaitResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prcd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"DecisionMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"NewBidHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"rating\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"review\",\"type\":\"string\"}],\"name\":\"RateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"}],\"name\":\"RevealAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auction_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"checkAverageScore\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"not_winner_platform\",\"type\":\"bool\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auction_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_asset_id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_score\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_feedback\",\"type\":\"string\"}],\"name\":\"provide_feedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"revealAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051611b60380380611b6083398101604081905261002f91610052565b6001600160a01b0316608052600080546001600160a01b03191633179055610082565b60006020828403121561006457600080fd5b81516001600160a01b038116811461007b57600080fd5b9392505050565b608051611ab56100ab600039600081816102c8015281816106ea0152610ca20152611ab56000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80638da5cb5b116100a2578063c84d2f6a11610071578063c84d2f6a14610277578063d0a1414a1461028a578063d960d5731461029d578063ea1591bb146102b0578063fc0c546a146102c357600080fd5b80638da5cb5b146101fa5780639348cef71461020d578063b14c63c514610220578063b8ea6bf61461024e57600080fd5b8063451df52e116100de578063451df52e1461018057806355f78c8d146101c15780637efbf8ac146101d45780638d10b0a6146101e757600080fd5b80630d058fe314610110578063176321e9146101255780632e1a7d4d1461013857806342d21ef714610160575b600080fd5b61012361011e3660046113e7565b6102ea565b005b61012361013336600461144e565b6103f3565b61014b610146366004611495565b6105eb565b60405190151581526020015b60405180910390f35b61017361016e366004611495565b6107c9565b60405161015791906114f4565b6101a961018e366004611495565b6003602052600090815260409020546001600160a01b031681565b6040516001600160a01b039091168152602001610157565b6101236101cf366004611495565b610863565b6101236101e236600461150e565b61097a565b6101736101f5366004611495565b610b48565b6000546101a9906001600160a01b031681565b61012361021b36600461155e565b610b61565b61024061022e366004611495565b60046020526000908152604090205481565b604051908152602001610157565b6101a961025c366004611495565b6007602052600090815260409020546001600160a01b031681565b610123610285366004611591565b610e2b565b61012361029836600461144e565b610fff565b6102406102ab366004611495565b6111c1565b6101236102be36600461155e565b611271565b6101a97f000000000000000000000000000000000000000000000000000000000000000081565b6000546001600160a01b031633146103535760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79206f776e65722063616e20637265617465206e65772061756374696f6044820152603760f91b60648201526084015b60405180910390fd5b600083815260036020908152604080832080546001600160a01b03191690556004808352818420849055815180830183529081526337b832b760e11b818401528684526005909252909120906103a9908261164a565b5060008381526006602052604090206103c2838261164a565b5060009283526007602052604090922080546001600160a01b0319166001600160a01b039093169290921790915550565b60405165656e64696e6760d01b602082015260260160408051601f19818403018152828252805160209182012060008681526005835292909220919261043a92910161170a565b604051602081830303815290604052805190602001201461049d5760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e4720737461747573000000604482015260640161034a565b6000828152600360205260409020546001600160a01b031633146104d35760405162461bcd60e51b815260040161034a90611780565b60405180604001604052806007815260200166636c6f73696e6760c81b815250600560008481526020019081526020016000209081610512919061164a565b506000828152600360209081526040808320546004835281842054600690935281842091517f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7946105779488946001600160a01b03909416939092909190889061182d565b60405180910390a160008281526004602090815260408083205460038352818420546001600160a01b03168452600190925282208054919290916105bc908490611897565b909155505050600090815260036020908152604080832080546001600160a01b03191690556004909152812055565b6040516337b832b760e11b602082015260009060240160408051601f19818403018152828252805160209182012060008681526005835292909220919261063392910161170a565b60405160208183030381529060405280519060200120036106965760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420696e204f50454e20737461747573000000000000000000604482015260640161034a565b336000908152600160205260409020548015610774573360008181526001602052604080822091909155516323b872dd60e01b81523060048201526024810191909152604481018290526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303816000875af1158015610733573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061075791906118b0565b610774573360009081526001602052604081209190915592915050565b6000838152600660205260409081902090517f9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c916107b891869190339086906118cd565b60405180910390a150600192915050565b600560205260009081526040902080546107e2906115c1565b80601f016020809104026020016040519081016040528092919081815260200182805461080e906115c1565b801561085b5780601f106108305761010080835404028352916020019161085b565b820191906000526020600020905b81548152906001019060200180831161083e57829003601f168201915b505050505081565b6000546001600160a01b0316331461088d5760405162461bcd60e51b815260040161034a90611902565b6040516337b832b760e11b602082015260240160408051601f1981840301815282825280516020918201206000858152600583529290922091926108d292910161170a565b60405160208183030381529060405280519060200120146109055760405162461bcd60e51b815260040161034a90611949565b604051806040016040528060068152602001651c995d99585b60d21b815250600560008381526020019081526020016000209081610943919061164a565b506040518181527f3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba9060200160405180910390a150565b60405166636c6f73696e6760c81b602082015260270160408051601f1981840301815282825280516020918201206000878152600583529290922091926109c292910161170a565b6040516020818303038152906040528051906020012014610a255760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e20434c4f53494e47207374617475730000604482015260640161034a565b6000838152600360205260409020546001600160a01b03163314610a5b5760405162461bcd60e51b815260040161034a90611780565b600083815260076020908152604080832080546001600160a01b0390811685526009845282852080546001818101835591875285872001889055915416845260088352908320805491820181558352912001610ab7828261164a565b506000838152600660205260409081902090517fcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa791610afc9186919086908690611980565b60405180910390a16040518060400160405280600681526020016518db1bdcd95960d21b815250600560008581526020019081526020016000209081610b42919061164a565b50505050565b600660205260009081526040902080546107e2906115c1565b604051651c995d99585b60d21b602082015260260160408051601f198184030181528282528051602091820120600086815260058352929092209192610ba892910161170a565b6040516020818303038152906040528051906020012014610bdb5760405162461bcd60e51b815260040161034a90611949565b6000828152600460205260409020548111610c385760405162461bcd60e51b815260206004820152601e60248201527f546865726520616c7265616479206973206120686967686572206269642e0000604482015260640161034a565b600082815260026020908152604080832033845282529182902054825191820184905291016040516020818303038152906040528051906020012014610c7d57600080fd5b6040516323b872dd60e01b8152336004820152306024820152604481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610cf3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1791906118b0565b905080610d5f5760405162461bcd60e51b81526020600482015260166024820152752a37b5b2b7103a3930b739b332b9103330b4b632b21760511b604482015260640161034a565b60008381526004602052604090205415610db65760008381526004602090815260408083205460038352818420546001600160a01b0316845260019092528220805491929091610db0908490611897565b90915550505b600083815260036020908152604080832080546001600160a01b031916339081179091556004835281842086905560069092529182902091517f463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde0892610e1e92879287906118cd565b60405180910390a1505050565b6000546001600160a01b03163314610e555760405162461bcd60e51b815260040161034a90611902565b6040516337b832b760e11b602082015260240160408051601f198184030181528282528051602091820120600086815260058352929092209192610e9a92910161170a565b6040516020818303038152906040528051906020012014610ecd5760405162461bcd60e51b815260040161034a90611949565b60405180604001604052806006815260200165656e64696e6760d01b815250600560008481526020019081526020016000209081610f0b919061164a565b508080610f245750600082815260046020526040902054155b15610fa5576040518060400160405280600681526020016518db1bdcd95960d21b815250600560008481526020019081526020016000209081610f67919061164a565b5060008281526004602090815260408083205460038352818420546001600160a01b03168452600190925282208054919290916105bc908490611897565b6000828152600360209081526040918290205482518581526001600160a01b03909116918101919091527fa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b484746850027391015b60405180910390a15050565b60405165656e64696e6760d01b602082015260260160408051601f19818403018152828252805160209182012060008681526005835292909220919261104692910161170a565b60405160208183030381529060405280519060200120146110a95760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e4720737461747573000000604482015260640161034a565b6000828152600360205260409020546001600160a01b031633146110df5760405162461bcd60e51b815260040161034a90611780565b60405180604001604052806007815260200166636c6f73696e6760c81b81525060056000848152602001908152602001600020908161111e919061164a565b5060008281526004602090815260408083205483546001600160a01b0316845260019092528220805491929091611156908490611897565b9091555050600082815260036020908152604080832054600483528184205460069093529281902090517f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c793610ff39387936001600160a01b0390921692909190600190889061182d565b6000818152600760209081526040808320546001600160a01b0316835260099091528120548190815b81811015611252576000858152600760209081526040808320546001600160a01b0316835260099091529020805482908110611228576112286119bc565b90600052602060002001548361123e91906119d2565b92508061124a816119fa565b9150506111ea565b508061125f836064611a13565b6112699190611a43565b949350505050565b6040516337b832b760e11b602082015260240160408051601f1981840301815282825280516020918201206000868152600583529290922091926112b692910161170a565b60405160208183030381529060405280519060200120146112e95760405162461bcd60e51b815260040161034a90611949565b60008281526002602090815260408083203380855290835281842085905585845260069092529182902091517f6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f92610ff392869286906118cd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261136b57600080fd5b813567ffffffffffffffff8082111561138657611386611344565b604051601f8301601f19908116603f011681019082821181831017156113ae576113ae611344565b816040528381528660208588010111156113c757600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156113fc57600080fd5b83359250602084013567ffffffffffffffff81111561141a57600080fd5b6114268682870161135a565b92505060408401356001600160a01b038116811461144357600080fd5b809150509250925092565b6000806040838503121561146157600080fd5b82359150602083013567ffffffffffffffff81111561147f57600080fd5b61148b8582860161135a565b9150509250929050565b6000602082840312156114a757600080fd5b5035919050565b6000815180845260005b818110156114d4576020818501810151868301820152016114b8565b506000602082860101526020601f19601f83011685010191505092915050565b60208152600061150760208301846114ae565b9392505050565b60008060006060848603121561152357600080fd5b8335925060208401359150604084013567ffffffffffffffff81111561154857600080fd5b6115548682870161135a565b9150509250925092565b6000806040838503121561157157600080fd5b50508035926020909101359150565b801515811461158e57600080fd5b50565b600080604083850312156115a457600080fd5b8235915060208301356115b681611580565b809150509250929050565b600181811c908216806115d557607f821691505b6020821081036115f557634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561164557600081815260208120601f850160051c810160208610156116225750805b601f850160051c820191505b818110156116415782815560010161162e565b5050505b505050565b815167ffffffffffffffff81111561166457611664611344565b6116788161167284546115c1565b846115fb565b602080601f8311600181146116ad57600084156116955750858301515b600019600386901b1c1916600185901b178555611641565b600085815260208120601f198616915b828110156116dc578886015182559484019460019091019084016116bd565b50858210156116fa5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000808354611718816115c1565b60018281168015611730576001811461174557611774565b60ff1984168752821515830287019450611774565b8760005260208060002060005b8581101561176b5781548a820152908401908201611752565b50505082870194505b50929695505050505050565b6020808252601690820152754e6f7420617574686f72697a6564206163636573732160501b604082015260600190565b600081546117bd816115c1565b8085526020600183811680156117da57600181146117f457611822565b60ff1985168884015283151560051b880183019550611822565b866000528260002060005b8581101561181a5781548a82018601529083019084016117ff565b890184019650505b505050505092915050565b86815260018060a01b038616602082015284604082015260c06060820152600061185a60c08301866117b0565b841515608084015282810360a084015261187481856114ae565b9998505050505050505050565b634e487b7160e01b600052601160045260246000fd5b808201808211156118aa576118aa611881565b92915050565b6000602082840312156118c257600080fd5b815161150781611580565b8481526080602082015260006118e660808301866117b0565b6001600160a01b03949094166040830152506060015292915050565b60208082526027908201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736040820152662073746174757360c81b606082015260800190565b6020808252601b908201527f436f6e7472616374206e6f7420696e204f50454e207374617475730000000000604082015260600190565b84815260806020820152600061199960808301866117b0565b84604084015282810360608401526119b181856114ae565b979650505050505050565b634e487b7160e01b600052603260045260246000fd5b80820182811260008312801582168215821617156119f2576119f2611881565b505092915050565b600060018201611a0c57611a0c611881565b5060010190565b80820260008212600160ff1b84141615611a2f57611a2f611881565b81810583148215176118aa576118aa611881565b600082611a6057634e487b7160e01b600052601260045260246000fd5b600160ff1b821460001984141615611a7a57611a7a611881565b50059056fea264697066735822122079066d911ce5f3fefdd88716e52aa6a62c745fec162b5a3dfcc108a2edb99e1364736f6c63430008120033",
}

// EthClosedBidAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use EthClosedBidAuctionMetaData.ABI instead.
var EthClosedBidAuctionABI = EthClosedBidAuctionMetaData.ABI

// EthClosedBidAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthClosedBidAuctionMetaData.Bin instead.
var EthClosedBidAuctionBin = EthClosedBidAuctionMetaData.Bin

// DeployEthClosedBidAuction deploys a new Ethereum contract, binding an instance of EthClosedBidAuction to it.
func DeployEthClosedBidAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *EthClosedBidAuction, error) {
	parsed, err := EthClosedBidAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthClosedBidAuctionBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthClosedBidAuction{EthClosedBidAuctionCaller: EthClosedBidAuctionCaller{contract: contract}, EthClosedBidAuctionTransactor: EthClosedBidAuctionTransactor{contract: contract}, EthClosedBidAuctionFilterer: EthClosedBidAuctionFilterer{contract: contract}}, nil
}

// EthClosedBidAuction is an auto generated Go binding around an Ethereum contract.
type EthClosedBidAuction struct {
	EthClosedBidAuctionCaller     // Read-only binding to the contract
	EthClosedBidAuctionTransactor // Write-only binding to the contract
	EthClosedBidAuctionFilterer   // Log filterer for contract events
}

// EthClosedBidAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthClosedBidAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClosedBidAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthClosedBidAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClosedBidAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthClosedBidAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClosedBidAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthClosedBidAuctionSession struct {
	Contract     *EthClosedBidAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthClosedBidAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthClosedBidAuctionCallerSession struct {
	Contract *EthClosedBidAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// EthClosedBidAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthClosedBidAuctionTransactorSession struct {
	Contract     *EthClosedBidAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// EthClosedBidAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthClosedBidAuctionRaw struct {
	Contract *EthClosedBidAuction // Generic contract binding to access the raw methods on
}

// EthClosedBidAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthClosedBidAuctionCallerRaw struct {
	Contract *EthClosedBidAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// EthClosedBidAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthClosedBidAuctionTransactorRaw struct {
	Contract *EthClosedBidAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthClosedBidAuction creates a new instance of EthClosedBidAuction, bound to a specific deployed contract.
func NewEthClosedBidAuction(address common.Address, backend bind.ContractBackend) (*EthClosedBidAuction, error) {
	contract, err := bindEthClosedBidAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthClosedBidAuction{EthClosedBidAuctionCaller: EthClosedBidAuctionCaller{contract: contract}, EthClosedBidAuctionTransactor: EthClosedBidAuctionTransactor{contract: contract}, EthClosedBidAuctionFilterer: EthClosedBidAuctionFilterer{contract: contract}}, nil
}

// NewEthClosedBidAuctionCaller creates a new read-only instance of EthClosedBidAuction, bound to a specific deployed contract.
func NewEthClosedBidAuctionCaller(address common.Address, caller bind.ContractCaller) (*EthClosedBidAuctionCaller, error) {
	contract, err := bindEthClosedBidAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthClosedBidAuctionCaller{contract: contract}, nil
}

// NewEthClosedBidAuctionTransactor creates a new write-only instance of EthClosedBidAuction, bound to a specific deployed contract.
func NewEthClosedBidAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*EthClosedBidAuctionTransactor, error) {
	contract, err := bindEthClosedBidAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthClosedBidAuctionTransactor{contract: contract}, nil
}

// NewEthClosedBidAuctionFilterer creates a new log filterer instance of EthClosedBidAuction, bound to a specific deployed contract.
func NewEthClosedBidAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*EthClosedBidAuctionFilterer, error) {
	contract, err := bindEthClosedBidAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthClosedBidAuctionFilterer{contract: contract}, nil
}

// bindEthClosedBidAuction binds a generic wrapper to an already deployed contract.
func bindEthClosedBidAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthClosedBidAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthClosedBidAuction *EthClosedBidAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthClosedBidAuction.Contract.EthClosedBidAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthClosedBidAuction *EthClosedBidAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.EthClosedBidAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthClosedBidAuction *EthClosedBidAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.EthClosedBidAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthClosedBidAuction *EthClosedBidAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthClosedBidAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthClosedBidAuction *EthClosedBidAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthClosedBidAuction *EthClosedBidAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.contract.Transact(opts, method, params...)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EthClosedBidAuction *EthClosedBidAuctionCaller) AssetId(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _EthClosedBidAuction.contract.Call(opts, &out, "asset_id", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EthClosedBidAuction *EthClosedBidAuctionSession) AssetId(arg0 *big.Int) (string, error) {
	return _EthClosedBidAuction.Contract.AssetId(&_EthClosedBidAuction.CallOpts, arg0)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EthClosedBidAuction *EthClosedBidAuctionCallerSession) AssetId(arg0 *big.Int) (string, error) {
	return _EthClosedBidAuction.Contract.AssetId(&_EthClosedBidAuction.CallOpts, arg0)
}

// AuctionOwner is a free data retrieval call binding the contract method 0xb8ea6bf6.
//
// Solidity: function auction_owner(uint256 ) view returns(address)
func (_EthClosedBidAuction *EthClosedBidAuctionCaller) AuctionOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EthClosedBidAuction.contract.Call(opts, &out, "auction_owner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuctionOwner is a free data retrieval call binding the contract method 0xb8ea6bf6.
//
// Solidity: function auction_owner(uint256 ) view returns(address)
func (_EthClosedBidAuction *EthClosedBidAuctionSession) AuctionOwner(arg0 *big.Int) (common.Address, error) {
	return _EthClosedBidAuction.Contract.AuctionOwner(&_EthClosedBidAuction.CallOpts, arg0)
}

// AuctionOwner is a free data retrieval call binding the contract method 0xb8ea6bf6.
//
// Solidity: function auction_owner(uint256 ) view returns(address)
func (_EthClosedBidAuction *EthClosedBidAuctionCallerSession) AuctionOwner(arg0 *big.Int) (common.Address, error) {
	return _EthClosedBidAuction.Contract.AuctionOwner(&_EthClosedBidAuction.CallOpts, arg0)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_EthClosedBidAuction *EthClosedBidAuctionCaller) CheckAverageScore(opts *bind.CallOpts, auctionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EthClosedBidAuction.contract.Call(opts, &out, "checkAverageScore", auctionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_EthClosedBidAuction *EthClosedBidAuctionSession) CheckAverageScore(auctionId *big.Int) (*big.Int, error) {
	return _EthClosedBidAuction.Contract.CheckAverageScore(&_EthClosedBidAuction.CallOpts, auctionId)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_EthClosedBidAuction *EthClosedBidAuctionCallerSession) CheckAverageScore(auctionId *big.Int) (*big.Int, error) {
	return _EthClosedBidAuction.Contract.CheckAverageScore(&_EthClosedBidAuction.CallOpts, auctionId)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EthClosedBidAuction *EthClosedBidAuctionCaller) HighestBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EthClosedBidAuction.contract.Call(opts, &out, "highestBid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EthClosedBidAuction *EthClosedBidAuctionSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _EthClosedBidAuction.Contract.HighestBid(&_EthClosedBidAuction.CallOpts, arg0)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EthClosedBidAuction *EthClosedBidAuctionCallerSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _EthClosedBidAuction.Contract.HighestBid(&_EthClosedBidAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EthClosedBidAuction *EthClosedBidAuctionCaller) HighestBidder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EthClosedBidAuction.contract.Call(opts, &out, "highestBidder", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EthClosedBidAuction *EthClosedBidAuctionSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _EthClosedBidAuction.Contract.HighestBidder(&_EthClosedBidAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EthClosedBidAuction *EthClosedBidAuctionCallerSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _EthClosedBidAuction.Contract.HighestBidder(&_EthClosedBidAuction.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthClosedBidAuction *EthClosedBidAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthClosedBidAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthClosedBidAuction *EthClosedBidAuctionSession) Owner() (common.Address, error) {
	return _EthClosedBidAuction.Contract.Owner(&_EthClosedBidAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthClosedBidAuction *EthClosedBidAuctionCallerSession) Owner() (common.Address, error) {
	return _EthClosedBidAuction.Contract.Owner(&_EthClosedBidAuction.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EthClosedBidAuction *EthClosedBidAuctionCaller) Status(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _EthClosedBidAuction.contract.Call(opts, &out, "status", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EthClosedBidAuction *EthClosedBidAuctionSession) Status(arg0 *big.Int) (string, error) {
	return _EthClosedBidAuction.Contract.Status(&_EthClosedBidAuction.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EthClosedBidAuction *EthClosedBidAuctionCallerSession) Status(arg0 *big.Int) (string, error) {
	return _EthClosedBidAuction.Contract.Status(&_EthClosedBidAuction.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthClosedBidAuction *EthClosedBidAuctionCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthClosedBidAuction.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthClosedBidAuction *EthClosedBidAuctionSession) Token() (common.Address, error) {
	return _EthClosedBidAuction.Contract.Token(&_EthClosedBidAuction.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthClosedBidAuction *EthClosedBidAuctionCallerSession) Token() (common.Address, error) {
	return _EthClosedBidAuction.Contract.Token(&_EthClosedBidAuction.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactor) Abort(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidAuction.contract.Transact(opts, "abort", auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.Abort(&_EthClosedBidAuction.TransactOpts, auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactorSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.Abort(&_EthClosedBidAuction.TransactOpts, auctionId, jsonString)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactor) Bid(opts *bind.TransactOpts, auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _EthClosedBidAuction.contract.Transact(opts, "bid", auctionId, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionSession) Bid(auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.Bid(&_EthClosedBidAuction.TransactOpts, auctionId, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactorSession) Bid(auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.Bid(&_EthClosedBidAuction.TransactOpts, auctionId, bidHash)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactor) CloseAuction(opts *bind.TransactOpts, auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _EthClosedBidAuction.contract.Transact(opts, "closeAuction", auctionId, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionSession) CloseAuction(auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.CloseAuction(&_EthClosedBidAuction.TransactOpts, auctionId, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactorSession) CloseAuction(auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.CloseAuction(&_EthClosedBidAuction.TransactOpts, auctionId, not_winner_platform)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactor) Commit(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidAuction.contract.Transact(opts, "commit", auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.Commit(&_EthClosedBidAuction.TransactOpts, auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactorSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.Commit(&_EthClosedBidAuction.TransactOpts, auctionId, jsonString)
}

// Create is a paid mutator transaction binding the contract method 0x0d058fe3.
//
// Solidity: function create(uint256 _auction_id, string _asset_id, address _owner) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactor) Create(opts *bind.TransactOpts, _auction_id *big.Int, _asset_id string, _owner common.Address) (*types.Transaction, error) {
	return _EthClosedBidAuction.contract.Transact(opts, "create", _auction_id, _asset_id, _owner)
}

// Create is a paid mutator transaction binding the contract method 0x0d058fe3.
//
// Solidity: function create(uint256 _auction_id, string _asset_id, address _owner) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionSession) Create(_auction_id *big.Int, _asset_id string, _owner common.Address) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.Create(&_EthClosedBidAuction.TransactOpts, _auction_id, _asset_id, _owner)
}

// Create is a paid mutator transaction binding the contract method 0x0d058fe3.
//
// Solidity: function create(uint256 _auction_id, string _asset_id, address _owner) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactorSession) Create(_auction_id *big.Int, _asset_id string, _owner common.Address) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.Create(&_EthClosedBidAuction.TransactOpts, _auction_id, _asset_id, _owner)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactor) ProvideFeedback(opts *bind.TransactOpts, auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _EthClosedBidAuction.contract.Transact(opts, "provide_feedback", auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.ProvideFeedback(&_EthClosedBidAuction.TransactOpts, auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactorSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.ProvideFeedback(&_EthClosedBidAuction.TransactOpts, auctionId, _score, _feedback)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactor) Reveal(opts *bind.TransactOpts, auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EthClosedBidAuction.contract.Transact(opts, "reveal", auctionId, bidAmount)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionSession) Reveal(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.Reveal(&_EthClosedBidAuction.TransactOpts, auctionId, bidAmount)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactorSession) Reveal(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.Reveal(&_EthClosedBidAuction.TransactOpts, auctionId, bidAmount)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactor) RevealAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidAuction.contract.Transact(opts, "revealAuction", auctionId)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionSession) RevealAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.RevealAuction(&_EthClosedBidAuction.TransactOpts, auctionId)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_EthClosedBidAuction *EthClosedBidAuctionTransactorSession) RevealAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.RevealAuction(&_EthClosedBidAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EthClosedBidAuction *EthClosedBidAuctionTransactor) Withdraw(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidAuction.contract.Transact(opts, "withdraw", auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EthClosedBidAuction *EthClosedBidAuctionSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.Withdraw(&_EthClosedBidAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EthClosedBidAuction *EthClosedBidAuctionTransactorSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidAuction.Contract.Withdraw(&_EthClosedBidAuction.TransactOpts, auctionId)
}

// EthClosedBidAuctionAwaitResponseIterator is returned from FilterAwaitResponse and is used to iterate over the raw logs and unpacked data for AwaitResponse events raised by the EthClosedBidAuction contract.
type EthClosedBidAuctionAwaitResponseIterator struct {
	Event *EthClosedBidAuctionAwaitResponse // Event containing the contract specifics and raw log

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
func (it *EthClosedBidAuctionAwaitResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidAuctionAwaitResponse)
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
		it.Event = new(EthClosedBidAuctionAwaitResponse)
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
func (it *EthClosedBidAuctionAwaitResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidAuctionAwaitResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidAuctionAwaitResponse represents a AwaitResponse event raised by the EthClosedBidAuction contract.
type EthClosedBidAuctionAwaitResponse struct {
	Auction *big.Int
	Winner  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAwaitResponse is a free log retrieval operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) FilterAwaitResponse(opts *bind.FilterOpts) (*EthClosedBidAuctionAwaitResponseIterator, error) {

	logs, sub, err := _EthClosedBidAuction.contract.FilterLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidAuctionAwaitResponseIterator{contract: _EthClosedBidAuction.contract, event: "AwaitResponse", logs: logs, sub: sub}, nil
}

// WatchAwaitResponse is a free log subscription operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) WatchAwaitResponse(opts *bind.WatchOpts, sink chan<- *EthClosedBidAuctionAwaitResponse) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidAuction.contract.WatchLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidAuctionAwaitResponse)
				if err := _EthClosedBidAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
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
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) ParseAwaitResponse(log types.Log) (*EthClosedBidAuctionAwaitResponse, error) {
	event := new(EthClosedBidAuctionAwaitResponse)
	if err := _EthClosedBidAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidAuctionDecisionMadeIterator is returned from FilterDecisionMade and is used to iterate over the raw logs and unpacked data for DecisionMade events raised by the EthClosedBidAuction contract.
type EthClosedBidAuctionDecisionMadeIterator struct {
	Event *EthClosedBidAuctionDecisionMade // Event containing the contract specifics and raw log

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
func (it *EthClosedBidAuctionDecisionMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidAuctionDecisionMade)
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
		it.Event = new(EthClosedBidAuctionDecisionMade)
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
func (it *EthClosedBidAuctionDecisionMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidAuctionDecisionMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidAuctionDecisionMade represents a DecisionMade event raised by the EthClosedBidAuction contract.
type EthClosedBidAuctionDecisionMade struct {
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
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) FilterDecisionMade(opts *bind.FilterOpts) (*EthClosedBidAuctionDecisionMadeIterator, error) {

	logs, sub, err := _EthClosedBidAuction.contract.FilterLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidAuctionDecisionMadeIterator{contract: _EthClosedBidAuction.contract, event: "DecisionMade", logs: logs, sub: sub}, nil
}

// WatchDecisionMade is a free log subscription operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auction, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) WatchDecisionMade(opts *bind.WatchOpts, sink chan<- *EthClosedBidAuctionDecisionMade) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidAuction.contract.WatchLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidAuctionDecisionMade)
				if err := _EthClosedBidAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
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
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) ParseDecisionMade(log types.Log) (*EthClosedBidAuctionDecisionMade, error) {
	event := new(EthClosedBidAuctionDecisionMade)
	if err := _EthClosedBidAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidAuctionHighestBidIncreasedIterator is returned from FilterHighestBidIncreased and is used to iterate over the raw logs and unpacked data for HighestBidIncreased events raised by the EthClosedBidAuction contract.
type EthClosedBidAuctionHighestBidIncreasedIterator struct {
	Event *EthClosedBidAuctionHighestBidIncreased // Event containing the contract specifics and raw log

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
func (it *EthClosedBidAuctionHighestBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidAuctionHighestBidIncreased)
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
		it.Event = new(EthClosedBidAuctionHighestBidIncreased)
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
func (it *EthClosedBidAuctionHighestBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidAuctionHighestBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidAuctionHighestBidIncreased represents a HighestBidIncreased event raised by the EthClosedBidAuction contract.
type EthClosedBidAuctionHighestBidIncreased struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*EthClosedBidAuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _EthClosedBidAuction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidAuctionHighestBidIncreasedIterator{contract: _EthClosedBidAuction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) WatchHighestBidIncreased(opts *bind.WatchOpts, sink chan<- *EthClosedBidAuctionHighestBidIncreased) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidAuction.contract.WatchLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidAuctionHighestBidIncreased)
				if err := _EthClosedBidAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
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
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) ParseHighestBidIncreased(log types.Log) (*EthClosedBidAuctionHighestBidIncreased, error) {
	event := new(EthClosedBidAuctionHighestBidIncreased)
	if err := _EthClosedBidAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidAuctionNewBidHashIterator is returned from FilterNewBidHash and is used to iterate over the raw logs and unpacked data for NewBidHash events raised by the EthClosedBidAuction contract.
type EthClosedBidAuctionNewBidHashIterator struct {
	Event *EthClosedBidAuctionNewBidHash // Event containing the contract specifics and raw log

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
func (it *EthClosedBidAuctionNewBidHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidAuctionNewBidHash)
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
		it.Event = new(EthClosedBidAuctionNewBidHash)
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
func (it *EthClosedBidAuctionNewBidHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidAuctionNewBidHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidAuctionNewBidHash represents a NewBidHash event raised by the EthClosedBidAuction contract.
type EthClosedBidAuctionNewBidHash struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	BidHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBidHash is a free log retrieval operation binding the contract event 0x6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f.
//
// Solidity: event NewBidHash(uint256 auction, string id, address bidder, bytes32 bidHash)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) FilterNewBidHash(opts *bind.FilterOpts) (*EthClosedBidAuctionNewBidHashIterator, error) {

	logs, sub, err := _EthClosedBidAuction.contract.FilterLogs(opts, "NewBidHash")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidAuctionNewBidHashIterator{contract: _EthClosedBidAuction.contract, event: "NewBidHash", logs: logs, sub: sub}, nil
}

// WatchNewBidHash is a free log subscription operation binding the contract event 0x6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f.
//
// Solidity: event NewBidHash(uint256 auction, string id, address bidder, bytes32 bidHash)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) WatchNewBidHash(opts *bind.WatchOpts, sink chan<- *EthClosedBidAuctionNewBidHash) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidAuction.contract.WatchLogs(opts, "NewBidHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidAuctionNewBidHash)
				if err := _EthClosedBidAuction.contract.UnpackLog(event, "NewBidHash", log); err != nil {
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

// ParseNewBidHash is a log parse operation binding the contract event 0x6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f.
//
// Solidity: event NewBidHash(uint256 auction, string id, address bidder, bytes32 bidHash)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) ParseNewBidHash(log types.Log) (*EthClosedBidAuctionNewBidHash, error) {
	event := new(EthClosedBidAuctionNewBidHash)
	if err := _EthClosedBidAuction.contract.UnpackLog(event, "NewBidHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidAuctionRateAuctionIterator is returned from FilterRateAuction and is used to iterate over the raw logs and unpacked data for RateAuction events raised by the EthClosedBidAuction contract.
type EthClosedBidAuctionRateAuctionIterator struct {
	Event *EthClosedBidAuctionRateAuction // Event containing the contract specifics and raw log

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
func (it *EthClosedBidAuctionRateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidAuctionRateAuction)
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
		it.Event = new(EthClosedBidAuctionRateAuction)
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
func (it *EthClosedBidAuctionRateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidAuctionRateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidAuctionRateAuction represents a RateAuction event raised by the EthClosedBidAuction contract.
type EthClosedBidAuctionRateAuction struct {
	Auction *big.Int
	Id      string
	Rating  *big.Int
	Review  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRateAuction is a free log retrieval operation binding the contract event 0xcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa7.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, string review)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) FilterRateAuction(opts *bind.FilterOpts) (*EthClosedBidAuctionRateAuctionIterator, error) {

	logs, sub, err := _EthClosedBidAuction.contract.FilterLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidAuctionRateAuctionIterator{contract: _EthClosedBidAuction.contract, event: "RateAuction", logs: logs, sub: sub}, nil
}

// WatchRateAuction is a free log subscription operation binding the contract event 0xcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa7.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, string review)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) WatchRateAuction(opts *bind.WatchOpts, sink chan<- *EthClosedBidAuctionRateAuction) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidAuction.contract.WatchLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidAuctionRateAuction)
				if err := _EthClosedBidAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
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
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) ParseRateAuction(log types.Log) (*EthClosedBidAuctionRateAuction, error) {
	event := new(EthClosedBidAuctionRateAuction)
	if err := _EthClosedBidAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidAuctionRevealAuctionIterator is returned from FilterRevealAuction and is used to iterate over the raw logs and unpacked data for RevealAuction events raised by the EthClosedBidAuction contract.
type EthClosedBidAuctionRevealAuctionIterator struct {
	Event *EthClosedBidAuctionRevealAuction // Event containing the contract specifics and raw log

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
func (it *EthClosedBidAuctionRevealAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidAuctionRevealAuction)
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
		it.Event = new(EthClosedBidAuctionRevealAuction)
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
func (it *EthClosedBidAuctionRevealAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidAuctionRevealAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidAuctionRevealAuction represents a RevealAuction event raised by the EthClosedBidAuction contract.
type EthClosedBidAuctionRevealAuction struct {
	Auction *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevealAuction is a free log retrieval operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auction)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) FilterRevealAuction(opts *bind.FilterOpts) (*EthClosedBidAuctionRevealAuctionIterator, error) {

	logs, sub, err := _EthClosedBidAuction.contract.FilterLogs(opts, "RevealAuction")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidAuctionRevealAuctionIterator{contract: _EthClosedBidAuction.contract, event: "RevealAuction", logs: logs, sub: sub}, nil
}

// WatchRevealAuction is a free log subscription operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auction)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) WatchRevealAuction(opts *bind.WatchOpts, sink chan<- *EthClosedBidAuctionRevealAuction) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidAuction.contract.WatchLogs(opts, "RevealAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidAuctionRevealAuction)
				if err := _EthClosedBidAuction.contract.UnpackLog(event, "RevealAuction", log); err != nil {
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

// ParseRevealAuction is a log parse operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auction)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) ParseRevealAuction(log types.Log) (*EthClosedBidAuctionRevealAuction, error) {
	event := new(EthClosedBidAuctionRevealAuction)
	if err := _EthClosedBidAuction.contract.UnpackLog(event, "RevealAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidAuctionWithdrawBidIterator is returned from FilterWithdrawBid and is used to iterate over the raw logs and unpacked data for WithdrawBid events raised by the EthClosedBidAuction contract.
type EthClosedBidAuctionWithdrawBidIterator struct {
	Event *EthClosedBidAuctionWithdrawBid // Event containing the contract specifics and raw log

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
func (it *EthClosedBidAuctionWithdrawBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidAuctionWithdrawBid)
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
		it.Event = new(EthClosedBidAuctionWithdrawBid)
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
func (it *EthClosedBidAuctionWithdrawBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidAuctionWithdrawBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidAuctionWithdrawBid represents a WithdrawBid event raised by the EthClosedBidAuction contract.
type EthClosedBidAuctionWithdrawBid struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBid is a free log retrieval operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) FilterWithdrawBid(opts *bind.FilterOpts) (*EthClosedBidAuctionWithdrawBidIterator, error) {

	logs, sub, err := _EthClosedBidAuction.contract.FilterLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidAuctionWithdrawBidIterator{contract: _EthClosedBidAuction.contract, event: "WithdrawBid", logs: logs, sub: sub}, nil
}

// WatchWithdrawBid is a free log subscription operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) WatchWithdrawBid(opts *bind.WatchOpts, sink chan<- *EthClosedBidAuctionWithdrawBid) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidAuction.contract.WatchLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidAuctionWithdrawBid)
				if err := _EthClosedBidAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
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
func (_EthClosedBidAuction *EthClosedBidAuctionFilterer) ParseWithdrawBid(log types.Log) (*EthClosedBidAuctionWithdrawBid, error) {
	event := new(EthClosedBidAuctionWithdrawBid)
	if err := _EthClosedBidAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
