// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cb1p_auction

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

// Cb1pAuctionMetaData contains all meta data concerning the Cb1pAuction contract.
var Cb1pAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AwaitResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prcd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"DecisionMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"NewBidHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"rating\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"review\",\"type\":\"string\"}],\"name\":\"RateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"}],\"name\":\"RevealAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auction_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"checkAverageScore\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"not_winner_platform\",\"type\":\"bool\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auction_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_asset_id\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_score\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_feedback\",\"type\":\"string\"}],\"name\":\"provide_feedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"revealAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051611acc380380611acc83398101604081905261002f91610052565b6001600160a01b0316608052600080546001600160a01b03191633179055610082565b60006020828403121561006457600080fd5b81516001600160a01b038116811461007b57600080fd5b9392505050565b608051611a216100ab600039600081816102c8015281816106be0152610c760152611a216000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80638da5cb5b116100a2578063c84d2f6a11610071578063c84d2f6a14610277578063d0a1414a1461028a578063d960d5731461029d578063ea1591bb146102b0578063fc0c546a146102c357600080fd5b80638da5cb5b146101fa5780639348cef71461020d578063b14c63c514610220578063b8ea6bf61461024e57600080fd5b8063451df52e116100de578063451df52e1461018057806355f78c8d146101c15780637efbf8ac146101d45780638d10b0a6146101e757600080fd5b80630118fa4914610110578063176321e9146101255780632e1a7d4d1461013857806342d21ef714610160575b600080fd5b61012361011e3660046113bb565b6102ea565b005b6101236101333660046113bb565b6103c7565b61014b610146366004611402565b6105bf565b60405190151581526020015b60405180910390f35b61017361016e366004611402565b61079d565b6040516101579190611461565b6101a961018e366004611402565b6003602052600090815260409020546001600160a01b031681565b6040516001600160a01b039091168152602001610157565b6101236101cf366004611402565b610837565b6101236101e236600461147b565b61094e565b6101736101f5366004611402565b610b1c565b6000546101a9906001600160a01b031681565b61012361021b3660046114cb565b610b35565b61024061022e366004611402565b60046020526000908152604090205481565b604051908152602001610157565b6101a961025c366004611402565b6007602052600090815260409020546001600160a01b031681565b6101236102853660046114fe565b610dff565b6101236102983660046113bb565b610fd3565b6102406102ab366004611402565b611195565b6101236102be3660046114cb565b611245565b6101a97f000000000000000000000000000000000000000000000000000000000000000081565b6000546001600160a01b031633146103535760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79206f776e65722063616e20637265617465206e65772061756374696f6044820152603760f91b60648201526084015b60405180910390fd5b600082815260036020908152604080832080546001600160a01b03191690556004808352818420849055815180830183529081526337b832b760e11b818401528584526005909252909120906103a990826115b6565b5060008281526006602052604090206103c282826115b6565b505050565b60405165656e64696e6760d01b602082015260260160408051601f19818403018152828252805160209182012060008681526005835292909220919261040e929101611676565b60405160208183030381529060405280519060200120146104715760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e4720737461747573000000604482015260640161034a565b6000828152600360205260409020546001600160a01b031633146104a75760405162461bcd60e51b815260040161034a906116ec565b60405180604001604052806007815260200166636c6f73696e6760c81b8152506005600084815260200190815260200160002090816104e691906115b6565b506000828152600360209081526040808320546004835281842054600690935281842091517f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c79461054b9488946001600160a01b039094169390929091908890611799565b60405180910390a160008281526004602090815260408083205460038352818420546001600160a01b0316845260019092528220805491929091610590908490611803565b909155505050600090815260036020908152604080832080546001600160a01b03191690556004909152812055565b6040516337b832b760e11b602082015260009060240160408051601f198184030181528282528051602091820120600086815260058352929092209192610607929101611676565b604051602081830303815290604052805190602001200361066a5760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420696e204f50454e20737461747573000000000000000000604482015260640161034a565b336000908152600160205260409020548015610748573360008181526001602052604080822091909155516323b872dd60e01b81523060048201526024810191909152604481018290526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303816000875af1158015610707573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072b919061181c565b610748573360009081526001602052604081209190915592915050565b6000838152600660205260409081902090517f9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c9161078c9186919033908690611839565b60405180910390a150600192915050565b600560205260009081526040902080546107b69061152e565b80601f01602080910402602001604051908101604052809291908181526020018280546107e29061152e565b801561082f5780601f106108045761010080835404028352916020019161082f565b820191906000526020600020905b81548152906001019060200180831161081257829003601f168201915b505050505081565b6000546001600160a01b031633146108615760405162461bcd60e51b815260040161034a9061186e565b6040516337b832b760e11b602082015260240160408051601f1981840301815282825280516020918201206000858152600583529290922091926108a6929101611676565b60405160208183030381529060405280519060200120146108d95760405162461bcd60e51b815260040161034a906118b5565b604051806040016040528060068152602001651c995d99585b60d21b81525060056000838152602001908152602001600020908161091791906115b6565b506040518181527f3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba9060200160405180910390a150565b60405166636c6f73696e6760c81b602082015260270160408051601f198184030181528282528051602091820120600087815260058352929092209192610996929101611676565b60405160208183030381529060405280519060200120146109f95760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e20434c4f53494e47207374617475730000604482015260640161034a565b6000838152600360205260409020546001600160a01b03163314610a2f5760405162461bcd60e51b815260040161034a906116ec565b600083815260076020908152604080832080546001600160a01b0390811685526009845282852080546001818101835591875285872001889055915416845260088352908320805491820181558352912001610a8b82826115b6565b506000838152600660205260409081902090517fcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa791610ad091869190869086906118ec565b60405180910390a16040518060400160405280600681526020016518db1bdcd95960d21b815250600560008581526020019081526020016000209081610b1691906115b6565b50505050565b600660205260009081526040902080546107b69061152e565b604051651c995d99585b60d21b602082015260260160408051601f198184030181528282528051602091820120600086815260058352929092209192610b7c929101611676565b6040516020818303038152906040528051906020012014610baf5760405162461bcd60e51b815260040161034a906118b5565b6000828152600460205260409020548111610c0c5760405162461bcd60e51b815260206004820152601e60248201527f546865726520616c7265616479206973206120686967686572206269642e0000604482015260640161034a565b600082815260026020908152604080832033845282529182902054825191820184905291016040516020818303038152906040528051906020012014610c5157600080fd5b6040516323b872dd60e01b8152336004820152306024820152604481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610cc7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ceb919061181c565b905080610d335760405162461bcd60e51b81526020600482015260166024820152752a37b5b2b7103a3930b739b332b9103330b4b632b21760511b604482015260640161034a565b60008381526004602052604090205415610d8a5760008381526004602090815260408083205460038352818420546001600160a01b0316845260019092528220805491929091610d84908490611803565b90915550505b600083815260036020908152604080832080546001600160a01b031916339081179091556004835281842086905560069092529182902091517f463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde0892610df29287928790611839565b60405180910390a1505050565b6000546001600160a01b03163314610e295760405162461bcd60e51b815260040161034a9061186e565b6040516337b832b760e11b602082015260240160408051601f198184030181528282528051602091820120600086815260058352929092209192610e6e929101611676565b6040516020818303038152906040528051906020012014610ea15760405162461bcd60e51b815260040161034a906118b5565b60405180604001604052806006815260200165656e64696e6760d01b815250600560008481526020019081526020016000209081610edf91906115b6565b508080610ef85750600082815260046020526040902054155b15610f79576040518060400160405280600681526020016518db1bdcd95960d21b815250600560008481526020019081526020016000209081610f3b91906115b6565b5060008281526004602090815260408083205460038352818420546001600160a01b0316845260019092528220805491929091610590908490611803565b6000828152600360209081526040918290205482518581526001600160a01b03909116918101919091527fa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b484746850027391015b60405180910390a15050565b60405165656e64696e6760d01b602082015260260160408051601f19818403018152828252805160209182012060008681526005835292909220919261101a929101611676565b604051602081830303815290604052805190602001201461107d5760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e4720737461747573000000604482015260640161034a565b6000828152600360205260409020546001600160a01b031633146110b35760405162461bcd60e51b815260040161034a906116ec565b60405180604001604052806007815260200166636c6f73696e6760c81b8152506005600084815260200190815260200160002090816110f291906115b6565b5060008281526004602090815260408083205483546001600160a01b031684526001909252822080549192909161112a908490611803565b9091555050600082815260036020908152604080832054600483528184205460069093529281902090517f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c793610fc79387936001600160a01b03909216929091906001908890611799565b6000818152600760209081526040808320546001600160a01b0316835260099091528120548190815b81811015611226576000858152600760209081526040808320546001600160a01b03168352600990915290208054829081106111fc576111fc611928565b906000526020600020015483611212919061193e565b92508061121e81611966565b9150506111be565b508061123383606461197f565b61123d91906119af565b949350505050565b6040516337b832b760e11b602082015260240160408051601f19818403018152828252805160209182012060008681526005835292909220919261128a929101611676565b60405160208183030381529060405280519060200120146112bd5760405162461bcd60e51b815260040161034a906118b5565b60008281526002602090815260408083203380855290835281842085905585845260069092529182902091517f6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f92610fc79286928690611839565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261133f57600080fd5b813567ffffffffffffffff8082111561135a5761135a611318565b604051601f8301601f19908116603f0116810190828211818310171561138257611382611318565b8160405283815286602085880101111561139b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080604083850312156113ce57600080fd5b82359150602083013567ffffffffffffffff8111156113ec57600080fd5b6113f88582860161132e565b9150509250929050565b60006020828403121561141457600080fd5b5035919050565b6000815180845260005b8181101561144157602081850181015186830182015201611425565b506000602082860101526020601f19601f83011685010191505092915050565b602081526000611474602083018461141b565b9392505050565b60008060006060848603121561149057600080fd5b8335925060208401359150604084013567ffffffffffffffff8111156114b557600080fd5b6114c18682870161132e565b9150509250925092565b600080604083850312156114de57600080fd5b50508035926020909101359150565b80151581146114fb57600080fd5b50565b6000806040838503121561151157600080fd5b823591506020830135611523816114ed565b809150509250929050565b600181811c9082168061154257607f821691505b60208210810361156257634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156103c257600081815260208120601f850160051c8101602086101561158f5750805b601f850160051c820191505b818110156115ae5782815560010161159b565b505050505050565b815167ffffffffffffffff8111156115d0576115d0611318565b6115e4816115de845461152e565b84611568565b602080601f83116001811461161957600084156116015750858301515b600019600386901b1c1916600185901b1785556115ae565b600085815260208120601f198616915b8281101561164857888601518255948401946001909101908401611629565b50858210156116665787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008083546116848161152e565b6001828116801561169c57600181146116b1576116e0565b60ff19841687528215158302870194506116e0565b8760005260208060002060005b858110156116d75781548a8201529084019082016116be565b50505082870194505b50929695505050505050565b6020808252601690820152754e6f7420617574686f72697a6564206163636573732160501b604082015260600190565b600081546117298161152e565b80855260206001838116801561174657600181146117605761178e565b60ff1985168884015283151560051b88018301955061178e565b866000528260002060005b858110156117865781548a820186015290830190840161176b565b890184019650505b505050505092915050565b86815260018060a01b038616602082015284604082015260c0606082015260006117c660c083018661171c565b841515608084015282810360a08401526117e0818561141b565b9998505050505050505050565b634e487b7160e01b600052601160045260246000fd5b80820180821115611816576118166117ed565b92915050565b60006020828403121561182e57600080fd5b8151611474816114ed565b848152608060208201526000611852608083018661171c565b6001600160a01b03949094166040830152506060015292915050565b60208082526027908201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736040820152662073746174757360c81b606082015260800190565b6020808252601b908201527f436f6e7472616374206e6f7420696e204f50454e207374617475730000000000604082015260600190565b848152608060208201526000611905608083018661171c565b846040840152828103606084015261191d818561141b565b979650505050505050565b634e487b7160e01b600052603260045260246000fd5b808201828112600083128015821682158216171561195e5761195e6117ed565b505092915050565b600060018201611978576119786117ed565b5060010190565b80820260008212600160ff1b8414161561199b5761199b6117ed565b8181058314821517611816576118166117ed565b6000826119cc57634e487b7160e01b600052601260045260246000fd5b600160ff1b8214600019841416156119e6576119e66117ed565b50059056fea2646970667358221220ce6aeb792ebceef58d0d22d958dc971ad36286419b8ff4feae4d15269879e15264736f6c63430008120033",
}

// Cb1pAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use Cb1pAuctionMetaData.ABI instead.
var Cb1pAuctionABI = Cb1pAuctionMetaData.ABI

// Cb1pAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Cb1pAuctionMetaData.Bin instead.
var Cb1pAuctionBin = Cb1pAuctionMetaData.Bin

// DeployCb1pAuction deploys a new Ethereum contract, binding an instance of Cb1pAuction to it.
func DeployCb1pAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *Cb1pAuction, error) {
	parsed, err := Cb1pAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Cb1pAuctionBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cb1pAuction{Cb1pAuctionCaller: Cb1pAuctionCaller{contract: contract}, Cb1pAuctionTransactor: Cb1pAuctionTransactor{contract: contract}, Cb1pAuctionFilterer: Cb1pAuctionFilterer{contract: contract}}, nil
}

// Cb1pAuction is an auto generated Go binding around an Ethereum contract.
type Cb1pAuction struct {
	Cb1pAuctionCaller     // Read-only binding to the contract
	Cb1pAuctionTransactor // Write-only binding to the contract
	Cb1pAuctionFilterer   // Log filterer for contract events
}

// Cb1pAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type Cb1pAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cb1pAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Cb1pAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cb1pAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Cb1pAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cb1pAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Cb1pAuctionSession struct {
	Contract     *Cb1pAuction      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Cb1pAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Cb1pAuctionCallerSession struct {
	Contract *Cb1pAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Cb1pAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Cb1pAuctionTransactorSession struct {
	Contract     *Cb1pAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Cb1pAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type Cb1pAuctionRaw struct {
	Contract *Cb1pAuction // Generic contract binding to access the raw methods on
}

// Cb1pAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Cb1pAuctionCallerRaw struct {
	Contract *Cb1pAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// Cb1pAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Cb1pAuctionTransactorRaw struct {
	Contract *Cb1pAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCb1pAuction creates a new instance of Cb1pAuction, bound to a specific deployed contract.
func NewCb1pAuction(address common.Address, backend bind.ContractBackend) (*Cb1pAuction, error) {
	contract, err := bindCb1pAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cb1pAuction{Cb1pAuctionCaller: Cb1pAuctionCaller{contract: contract}, Cb1pAuctionTransactor: Cb1pAuctionTransactor{contract: contract}, Cb1pAuctionFilterer: Cb1pAuctionFilterer{contract: contract}}, nil
}

// NewCb1pAuctionCaller creates a new read-only instance of Cb1pAuction, bound to a specific deployed contract.
func NewCb1pAuctionCaller(address common.Address, caller bind.ContractCaller) (*Cb1pAuctionCaller, error) {
	contract, err := bindCb1pAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionCaller{contract: contract}, nil
}

// NewCb1pAuctionTransactor creates a new write-only instance of Cb1pAuction, bound to a specific deployed contract.
func NewCb1pAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*Cb1pAuctionTransactor, error) {
	contract, err := bindCb1pAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionTransactor{contract: contract}, nil
}

// NewCb1pAuctionFilterer creates a new log filterer instance of Cb1pAuction, bound to a specific deployed contract.
func NewCb1pAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*Cb1pAuctionFilterer, error) {
	contract, err := bindCb1pAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionFilterer{contract: contract}, nil
}

// bindCb1pAuction binds a generic wrapper to an already deployed contract.
func bindCb1pAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Cb1pAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cb1pAuction *Cb1pAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cb1pAuction.Contract.Cb1pAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cb1pAuction *Cb1pAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Cb1pAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cb1pAuction *Cb1pAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Cb1pAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cb1pAuction *Cb1pAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cb1pAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cb1pAuction *Cb1pAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cb1pAuction *Cb1pAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.contract.Transact(opts, method, params...)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionCaller) AssetId(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "asset_id", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionSession) AssetId(arg0 *big.Int) (string, error) {
	return _Cb1pAuction.Contract.AssetId(&_Cb1pAuction.CallOpts, arg0)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionCallerSession) AssetId(arg0 *big.Int) (string, error) {
	return _Cb1pAuction.Contract.AssetId(&_Cb1pAuction.CallOpts, arg0)
}

// AuctionOwner is a free data retrieval call binding the contract method 0xb8ea6bf6.
//
// Solidity: function auction_owner(uint256 ) view returns(address)
func (_Cb1pAuction *Cb1pAuctionCaller) AuctionOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "auction_owner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuctionOwner is a free data retrieval call binding the contract method 0xb8ea6bf6.
//
// Solidity: function auction_owner(uint256 ) view returns(address)
func (_Cb1pAuction *Cb1pAuctionSession) AuctionOwner(arg0 *big.Int) (common.Address, error) {
	return _Cb1pAuction.Contract.AuctionOwner(&_Cb1pAuction.CallOpts, arg0)
}

// AuctionOwner is a free data retrieval call binding the contract method 0xb8ea6bf6.
//
// Solidity: function auction_owner(uint256 ) view returns(address)
func (_Cb1pAuction *Cb1pAuctionCallerSession) AuctionOwner(arg0 *big.Int) (common.Address, error) {
	return _Cb1pAuction.Contract.AuctionOwner(&_Cb1pAuction.CallOpts, arg0)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_Cb1pAuction *Cb1pAuctionCaller) CheckAverageScore(opts *bind.CallOpts, auctionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "checkAverageScore", auctionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_Cb1pAuction *Cb1pAuctionSession) CheckAverageScore(auctionId *big.Int) (*big.Int, error) {
	return _Cb1pAuction.Contract.CheckAverageScore(&_Cb1pAuction.CallOpts, auctionId)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_Cb1pAuction *Cb1pAuctionCallerSession) CheckAverageScore(auctionId *big.Int) (*big.Int, error) {
	return _Cb1pAuction.Contract.CheckAverageScore(&_Cb1pAuction.CallOpts, auctionId)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_Cb1pAuction *Cb1pAuctionCaller) HighestBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "highestBid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_Cb1pAuction *Cb1pAuctionSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _Cb1pAuction.Contract.HighestBid(&_Cb1pAuction.CallOpts, arg0)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_Cb1pAuction *Cb1pAuctionCallerSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _Cb1pAuction.Contract.HighestBid(&_Cb1pAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_Cb1pAuction *Cb1pAuctionCaller) HighestBidder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "highestBidder", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_Cb1pAuction *Cb1pAuctionSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _Cb1pAuction.Contract.HighestBidder(&_Cb1pAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_Cb1pAuction *Cb1pAuctionCallerSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _Cb1pAuction.Contract.HighestBidder(&_Cb1pAuction.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cb1pAuction *Cb1pAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cb1pAuction *Cb1pAuctionSession) Owner() (common.Address, error) {
	return _Cb1pAuction.Contract.Owner(&_Cb1pAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cb1pAuction *Cb1pAuctionCallerSession) Owner() (common.Address, error) {
	return _Cb1pAuction.Contract.Owner(&_Cb1pAuction.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionCaller) Status(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "status", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionSession) Status(arg0 *big.Int) (string, error) {
	return _Cb1pAuction.Contract.Status(&_Cb1pAuction.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionCallerSession) Status(arg0 *big.Int) (string, error) {
	return _Cb1pAuction.Contract.Status(&_Cb1pAuction.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cb1pAuction *Cb1pAuctionCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cb1pAuction *Cb1pAuctionSession) Token() (common.Address, error) {
	return _Cb1pAuction.Contract.Token(&_Cb1pAuction.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cb1pAuction *Cb1pAuctionCallerSession) Token() (common.Address, error) {
	return _Cb1pAuction.Contract.Token(&_Cb1pAuction.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) Abort(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "abort", auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_Cb1pAuction *Cb1pAuctionSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Abort(&_Cb1pAuction.TransactOpts, auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Abort(&_Cb1pAuction.TransactOpts, auctionId, jsonString)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) Bid(opts *bind.TransactOpts, auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "bid", auctionId, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_Cb1pAuction *Cb1pAuctionSession) Bid(auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Bid(&_Cb1pAuction.TransactOpts, auctionId, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) Bid(auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Bid(&_Cb1pAuction.TransactOpts, auctionId, bidHash)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) CloseAuction(opts *bind.TransactOpts, auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "closeAuction", auctionId, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_Cb1pAuction *Cb1pAuctionSession) CloseAuction(auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.CloseAuction(&_Cb1pAuction.TransactOpts, auctionId, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) CloseAuction(auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.CloseAuction(&_Cb1pAuction.TransactOpts, auctionId, not_winner_platform)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) Commit(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "commit", auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_Cb1pAuction *Cb1pAuctionSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Commit(&_Cb1pAuction.TransactOpts, auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Commit(&_Cb1pAuction.TransactOpts, auctionId, jsonString)
}

// Create is a paid mutator transaction binding the contract method 0x0118fa49.
//
// Solidity: function create(uint256 _auction_id, string _asset_id) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) Create(opts *bind.TransactOpts, _auction_id *big.Int, _asset_id string) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "create", _auction_id, _asset_id)
}

// Create is a paid mutator transaction binding the contract method 0x0118fa49.
//
// Solidity: function create(uint256 _auction_id, string _asset_id) returns()
func (_Cb1pAuction *Cb1pAuctionSession) Create(_auction_id *big.Int, _asset_id string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Create(&_Cb1pAuction.TransactOpts, _auction_id, _asset_id)
}

// Create is a paid mutator transaction binding the contract method 0x0118fa49.
//
// Solidity: function create(uint256 _auction_id, string _asset_id) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) Create(_auction_id *big.Int, _asset_id string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Create(&_Cb1pAuction.TransactOpts, _auction_id, _asset_id)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) ProvideFeedback(opts *bind.TransactOpts, auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "provide_feedback", auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_Cb1pAuction *Cb1pAuctionSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.ProvideFeedback(&_Cb1pAuction.TransactOpts, auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.ProvideFeedback(&_Cb1pAuction.TransactOpts, auctionId, _score, _feedback)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) Reveal(opts *bind.TransactOpts, auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "reveal", auctionId, bidAmount)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_Cb1pAuction *Cb1pAuctionSession) Reveal(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Reveal(&_Cb1pAuction.TransactOpts, auctionId, bidAmount)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) Reveal(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Reveal(&_Cb1pAuction.TransactOpts, auctionId, bidAmount)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) RevealAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "revealAuction", auctionId)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_Cb1pAuction *Cb1pAuctionSession) RevealAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.RevealAuction(&_Cb1pAuction.TransactOpts, auctionId)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) RevealAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.RevealAuction(&_Cb1pAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_Cb1pAuction *Cb1pAuctionTransactor) Withdraw(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "withdraw", auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_Cb1pAuction *Cb1pAuctionSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Withdraw(&_Cb1pAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_Cb1pAuction *Cb1pAuctionTransactorSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Withdraw(&_Cb1pAuction.TransactOpts, auctionId)
}

// Cb1pAuctionAwaitResponseIterator is returned from FilterAwaitResponse and is used to iterate over the raw logs and unpacked data for AwaitResponse events raised by the Cb1pAuction contract.
type Cb1pAuctionAwaitResponseIterator struct {
	Event *Cb1pAuctionAwaitResponse // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionAwaitResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionAwaitResponse)
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
		it.Event = new(Cb1pAuctionAwaitResponse)
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
func (it *Cb1pAuctionAwaitResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionAwaitResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionAwaitResponse represents a AwaitResponse event raised by the Cb1pAuction contract.
type Cb1pAuctionAwaitResponse struct {
	Auction *big.Int
	Winner  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAwaitResponse is a free log retrieval operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterAwaitResponse(opts *bind.FilterOpts) (*Cb1pAuctionAwaitResponseIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionAwaitResponseIterator{contract: _Cb1pAuction.contract, event: "AwaitResponse", logs: logs, sub: sub}, nil
}

// WatchAwaitResponse is a free log subscription operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchAwaitResponse(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionAwaitResponse) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionAwaitResponse)
				if err := _Cb1pAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
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
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseAwaitResponse(log types.Log) (*Cb1pAuctionAwaitResponse, error) {
	event := new(Cb1pAuctionAwaitResponse)
	if err := _Cb1pAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb1pAuctionDecisionMadeIterator is returned from FilterDecisionMade and is used to iterate over the raw logs and unpacked data for DecisionMade events raised by the Cb1pAuction contract.
type Cb1pAuctionDecisionMadeIterator struct {
	Event *Cb1pAuctionDecisionMade // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionDecisionMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionDecisionMade)
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
		it.Event = new(Cb1pAuctionDecisionMade)
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
func (it *Cb1pAuctionDecisionMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionDecisionMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionDecisionMade represents a DecisionMade event raised by the Cb1pAuction contract.
type Cb1pAuctionDecisionMade struct {
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
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterDecisionMade(opts *bind.FilterOpts) (*Cb1pAuctionDecisionMadeIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionDecisionMadeIterator{contract: _Cb1pAuction.contract, event: "DecisionMade", logs: logs, sub: sub}, nil
}

// WatchDecisionMade is a free log subscription operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auction, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchDecisionMade(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionDecisionMade) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionDecisionMade)
				if err := _Cb1pAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
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
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseDecisionMade(log types.Log) (*Cb1pAuctionDecisionMade, error) {
	event := new(Cb1pAuctionDecisionMade)
	if err := _Cb1pAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb1pAuctionHighestBidIncreasedIterator is returned from FilterHighestBidIncreased and is used to iterate over the raw logs and unpacked data for HighestBidIncreased events raised by the Cb1pAuction contract.
type Cb1pAuctionHighestBidIncreasedIterator struct {
	Event *Cb1pAuctionHighestBidIncreased // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionHighestBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionHighestBidIncreased)
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
		it.Event = new(Cb1pAuctionHighestBidIncreased)
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
func (it *Cb1pAuctionHighestBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionHighestBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionHighestBidIncreased represents a HighestBidIncreased event raised by the Cb1pAuction contract.
type Cb1pAuctionHighestBidIncreased struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*Cb1pAuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionHighestBidIncreasedIterator{contract: _Cb1pAuction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchHighestBidIncreased(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionHighestBidIncreased) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionHighestBidIncreased)
				if err := _Cb1pAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
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
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseHighestBidIncreased(log types.Log) (*Cb1pAuctionHighestBidIncreased, error) {
	event := new(Cb1pAuctionHighestBidIncreased)
	if err := _Cb1pAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb1pAuctionNewBidHashIterator is returned from FilterNewBidHash and is used to iterate over the raw logs and unpacked data for NewBidHash events raised by the Cb1pAuction contract.
type Cb1pAuctionNewBidHashIterator struct {
	Event *Cb1pAuctionNewBidHash // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionNewBidHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionNewBidHash)
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
		it.Event = new(Cb1pAuctionNewBidHash)
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
func (it *Cb1pAuctionNewBidHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionNewBidHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionNewBidHash represents a NewBidHash event raised by the Cb1pAuction contract.
type Cb1pAuctionNewBidHash struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	BidHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBidHash is a free log retrieval operation binding the contract event 0x6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f.
//
// Solidity: event NewBidHash(uint256 auction, string id, address bidder, bytes32 bidHash)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterNewBidHash(opts *bind.FilterOpts) (*Cb1pAuctionNewBidHashIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "NewBidHash")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionNewBidHashIterator{contract: _Cb1pAuction.contract, event: "NewBidHash", logs: logs, sub: sub}, nil
}

// WatchNewBidHash is a free log subscription operation binding the contract event 0x6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f.
//
// Solidity: event NewBidHash(uint256 auction, string id, address bidder, bytes32 bidHash)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchNewBidHash(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionNewBidHash) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "NewBidHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionNewBidHash)
				if err := _Cb1pAuction.contract.UnpackLog(event, "NewBidHash", log); err != nil {
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
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseNewBidHash(log types.Log) (*Cb1pAuctionNewBidHash, error) {
	event := new(Cb1pAuctionNewBidHash)
	if err := _Cb1pAuction.contract.UnpackLog(event, "NewBidHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb1pAuctionRateAuctionIterator is returned from FilterRateAuction and is used to iterate over the raw logs and unpacked data for RateAuction events raised by the Cb1pAuction contract.
type Cb1pAuctionRateAuctionIterator struct {
	Event *Cb1pAuctionRateAuction // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionRateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionRateAuction)
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
		it.Event = new(Cb1pAuctionRateAuction)
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
func (it *Cb1pAuctionRateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionRateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionRateAuction represents a RateAuction event raised by the Cb1pAuction contract.
type Cb1pAuctionRateAuction struct {
	Auction *big.Int
	Id      string
	Rating  *big.Int
	Review  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRateAuction is a free log retrieval operation binding the contract event 0xcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa7.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, string review)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterRateAuction(opts *bind.FilterOpts) (*Cb1pAuctionRateAuctionIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionRateAuctionIterator{contract: _Cb1pAuction.contract, event: "RateAuction", logs: logs, sub: sub}, nil
}

// WatchRateAuction is a free log subscription operation binding the contract event 0xcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa7.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, string review)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchRateAuction(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionRateAuction) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionRateAuction)
				if err := _Cb1pAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
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
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseRateAuction(log types.Log) (*Cb1pAuctionRateAuction, error) {
	event := new(Cb1pAuctionRateAuction)
	if err := _Cb1pAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb1pAuctionRevealAuctionIterator is returned from FilterRevealAuction and is used to iterate over the raw logs and unpacked data for RevealAuction events raised by the Cb1pAuction contract.
type Cb1pAuctionRevealAuctionIterator struct {
	Event *Cb1pAuctionRevealAuction // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionRevealAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionRevealAuction)
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
		it.Event = new(Cb1pAuctionRevealAuction)
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
func (it *Cb1pAuctionRevealAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionRevealAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionRevealAuction represents a RevealAuction event raised by the Cb1pAuction contract.
type Cb1pAuctionRevealAuction struct {
	Auction *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevealAuction is a free log retrieval operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auction)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterRevealAuction(opts *bind.FilterOpts) (*Cb1pAuctionRevealAuctionIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "RevealAuction")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionRevealAuctionIterator{contract: _Cb1pAuction.contract, event: "RevealAuction", logs: logs, sub: sub}, nil
}

// WatchRevealAuction is a free log subscription operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auction)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchRevealAuction(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionRevealAuction) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "RevealAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionRevealAuction)
				if err := _Cb1pAuction.contract.UnpackLog(event, "RevealAuction", log); err != nil {
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
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseRevealAuction(log types.Log) (*Cb1pAuctionRevealAuction, error) {
	event := new(Cb1pAuctionRevealAuction)
	if err := _Cb1pAuction.contract.UnpackLog(event, "RevealAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb1pAuctionWithdrawBidIterator is returned from FilterWithdrawBid and is used to iterate over the raw logs and unpacked data for WithdrawBid events raised by the Cb1pAuction contract.
type Cb1pAuctionWithdrawBidIterator struct {
	Event *Cb1pAuctionWithdrawBid // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionWithdrawBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionWithdrawBid)
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
		it.Event = new(Cb1pAuctionWithdrawBid)
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
func (it *Cb1pAuctionWithdrawBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionWithdrawBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionWithdrawBid represents a WithdrawBid event raised by the Cb1pAuction contract.
type Cb1pAuctionWithdrawBid struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBid is a free log retrieval operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterWithdrawBid(opts *bind.FilterOpts) (*Cb1pAuctionWithdrawBidIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionWithdrawBidIterator{contract: _Cb1pAuction.contract, event: "WithdrawBid", logs: logs, sub: sub}, nil
}

// WatchWithdrawBid is a free log subscription operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchWithdrawBid(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionWithdrawBid) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionWithdrawBid)
				if err := _Cb1pAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
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
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseWithdrawBid(log types.Log) (*Cb1pAuctionWithdrawBid, error) {
	event := new(Cb1pAuctionWithdrawBid)
	if err := _Cb1pAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
