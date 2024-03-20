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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AwaitResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prcd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"DecisionMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"auctionType\",\"type\":\"string\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"rating\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"review\",\"type\":\"string\"}],\"name\":\"RateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_owner\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auction_type\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"not_winner_platform\",\"type\":\"bool\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auction_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_asset_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_asset_owner\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_score\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_feedback\",\"type\":\"string\"}],\"name\":\"provide_feedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620019ae380380620019ae833981016040819052620000349162000090565b6001600160a01b038116608052600080546001600160a01b0319163317905560408051808201909152600f81526e22b733b634b9b41020bab1ba34b7b760891b602082015260019062000088908262000167565b505062000233565b600060208284031215620000a357600080fd5b81516001600160a01b0381168114620000bb57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620000ed57607f821691505b6020821081036200010e57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200016257600081815260208120601f850160051c810160208610156200013d5750805b601f850160051c820191505b818110156200015e5782815560010162000149565b5050505b505050565b81516001600160401b03811115620001835762000183620000c2565b6200019b81620001948454620000d8565b8462000114565b602080601f831160018114620001d35760008415620001ba5750858301515b600019600386901b1c1916600185901b1785556200015e565b600085815260208120601f198616915b828110156200020457888601518255948401946001909101908401620001e3565b5085821015620002235787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6080516117516200025d6000396000818161026b01528181610589015261084501526117516000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80637efbf8ac11610097578063c84d2f6a11610066578063c84d2f6a14610238578063d0a1414a1461024b578063e1081bb31461025e578063fc0c546a1461026657600080fd5b80637efbf8ac146101d15780638d10b0a6146101e45780638da5cb5b146101f7578063b14c63c51461020a57600080fd5b806344a770bf116100d357806344a770bf14610157578063451df52e1461016a578063598647f8146101ab5780635a4463a8146101be57600080fd5b8063176321e9146100fa5780632e1a7d4d1461010f57806342d21ef714610137575b600080fd5b61010d610108366004611183565b61028d565b005b61012261011d3660046111ca565b61048a565b60405190151581526020015b60405180910390f35b61014a6101453660046111ca565b610668565b60405161012e9190611229565b61014a6101653660046111ca565b610702565b6101936101783660046111ca565b6003602052600090815260409020546001600160a01b031681565b6040516001600160a01b03909116815260200161012e565b61010d6101b9366004611243565b61071b565b61010d6101cc366004611265565b6109d1565b61010d6101df3660046112d2565b610ac3565b61014a6101f23660046111ca565b610cb4565b600054610193906001600160a01b031681565b61022a6102183660046111ca565b60046020526000908152604090205481565b60405190815260200161012e565b61010d610246366004611329565b610ccd565b61010d610259366004611183565b610f11565b61014a6110d3565b6101937f000000000000000000000000000000000000000000000000000000000000000081565b60405165656e64696e6760d01b602082015260260160408051601f1981840301815282825280516020918201206000868152600583529290922091926102d4929101611393565b604051602081830303815290604052805190602001201461033c5760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064015b60405180910390fd5b6000828152600360205260409020546001600160a01b031633146103725760405162461bcd60e51b815260040161033390611409565b60405180604001604052806007815260200166636c6f73696e6760c81b8152506005600084815260200190815260200160002090816103b19190611488565b506000828152600360209081526040808320546004835281842054600690935281842091517f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7946104169488946001600160a01b0390941693909290919088906115c5565b60405180910390a160008281526004602090815260408083205460038352818420546001600160a01b031684526002909252822080549192909161045b908490611619565b909155505050600090815260036020908152604080832080546001600160a01b03191690556004909152812055565b6040516337b832b760e11b602082015260009060240160408051601f1981840301815282825280516020918201206000868152600583529290922091926104d2929101611393565b60405160208183030381529060405280519060200120036105355760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420696e204f50454e207374617475730000000000000000006044820152606401610333565b336000908152600260205260409020548015610613573360008181526002602052604080822091909155516323b872dd60e01b81523060048201526024810191909152604481018290526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303816000875af11580156105d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105f69190611640565b610613573360009081526002602052604081209190915592915050565b6000838152600660205260409081902090517f9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c91610657918691903390869061165d565b60405180910390a150600192915050565b6005602052600090815260409020805461068190611359565b80601f01602080910402602001604051908101604052809291908181526020018280546106ad90611359565b80156106fa5780601f106106cf576101008083540402835291602001916106fa565b820191906000526020600020905b8154815290600101906020018083116106dd57829003601f168201915b505050505081565b6007602052600090815260409020805461068190611359565b6040516337b832b760e11b602082015260240160408051601f198184030181528282528051602091820120600086815260058352929092209192610760929101611393565b60405160208183030381529060405280519060200120146107c35760405162461bcd60e51b815260206004820152601b60248201527f436f6e7472616374206e6f7420696e204f50454e2073746174757300000000006044820152606401610333565b60008281526004602052604090205481116108205760405162461bcd60e51b815260206004820152601e60248201527f546865726520616c7265616479206973206120686967686572206269642e00006044820152606401610333565b6040516323b872dd60e01b8152336004820152306024820152604481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610896573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ba9190611640565b9050806109025760405162461bcd60e51b81526020600482015260166024820152752a37b5b2b7103a3930b739b332b9103330b4b632b21760511b6044820152606401610333565b600083815260046020526040902054156109595760008381526004602090815260408083205460038352818420546001600160a01b0316845260029092528220805491929091610953908490611619565b90915550505b600083815260036020908152604080832080546001600160a01b031916339081179091556004835281842086905560069092529182902091517fb95c1199e4385d3e33ad9cffdc96a2f61491b3798965fd93633c02ff7ade77f0926109c49287928790600190611692565b60405180910390a1505050565b6000546001600160a01b03163314610a355760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79206f776e65722063616e20637265617465206e65772061756374696f6044820152603760f91b6064820152608401610333565b600083815260036020908152604080832080546001600160a01b03191690556004808352818420849055815180830183529081526337b832b760e11b81840152868452600590925290912090610a8b9082611488565b506000838152600660205260409020610aa48382611488565b506000838152600760205260409020610abd8282611488565b50505050565b60405166636c6f73696e6760c81b602082015260270160408051601f198184030181528282528051602091820120600087815260058352929092209192610b0b929101611393565b6040516020818303038152906040528051906020012014610b6e5760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e20434c4f53494e472073746174757300006044820152606401610333565b6000838152600360205260409020546001600160a01b03163314610ba45760405162461bcd60e51b815260040161033390611409565b600083815260076020526040908190209051600991610bc291611393565b9081526040805191829003602090810183208054600181018255600091825282822001869055868152600790915220600891610bfe9190611393565b908152604051602091819003820190208054600181018255600091825291902001610c298282611488565b506000838152600660205260409081902090517fcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa791610c6e91869190869086906116df565b60405180910390a16040518060400160405280600681526020016518db1bdcd95960d21b815250600560008581526020019081526020016000209081610abd9190611488565b6006602052600090815260409020805461068190611359565b6000546001600160a01b03163314610d375760405162461bcd60e51b815260206004820152602760248201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736044820152662073746174757360c81b6064820152608401610333565b6040516337b832b760e11b602082015260240160408051601f198184030181528282528051602091820120600086815260058352929092209192610d7c929101611393565b6040516020818303038152906040528051906020012014610ddf5760405162461bcd60e51b815260206004820152601b60248201527f436f6e7472616374206e6f7420696e204f50454e2073746174757300000000006044820152606401610333565b60405180604001604052806006815260200165656e64696e6760d01b815250600560008481526020019081526020016000209081610e1d9190611488565b508080610e365750600082815260046020526040902054155b15610eb7576040518060400160405280600681526020016518db1bdcd95960d21b815250600560008481526020019081526020016000209081610e799190611488565b5060008281526004602090815260408083205460038352818420546001600160a01b031684526002909252822080549192909161045b908490611619565b6000828152600360209081526040918290205482518581526001600160a01b03909116918101919091527fa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b484746850027391015b60405180910390a15050565b60405165656e64696e6760d01b602082015260260160408051601f198184030181528282528051602091820120600086815260058352929092209192610f58929101611393565b6040516020818303038152906040528051906020012014610fbb5760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e47207374617475730000006044820152606401610333565b6000828152600360205260409020546001600160a01b03163314610ff15760405162461bcd60e51b815260040161033390611409565b60405180604001604052806007815260200166636c6f73696e6760c81b8152506005600084815260200190815260200160002090816110309190611488565b5060008281526004602090815260408083205483546001600160a01b0316845260029092528220805491929091611068908490611619565b9091555050600082815260036020908152604080832054600483528184205460069093529281902090517f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c793610f059387936001600160a01b039092169290919060019088906115c5565b6001805461068190611359565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261110757600080fd5b813567ffffffffffffffff80821115611122576111226110e0565b604051601f8301601f19908116603f0116810190828211818310171561114a5761114a6110e0565b8160405283815286602085880101111561116357600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561119657600080fd5b82359150602083013567ffffffffffffffff8111156111b457600080fd5b6111c0858286016110f6565b9150509250929050565b6000602082840312156111dc57600080fd5b5035919050565b6000815180845260005b81811015611209576020818501810151868301820152016111ed565b506000602082860101526020601f19601f83011685010191505092915050565b60208152600061123c60208301846111e3565b9392505050565b6000806040838503121561125657600080fd5b50508035926020909101359150565b60008060006060848603121561127a57600080fd5b83359250602084013567ffffffffffffffff8082111561129957600080fd5b6112a5878388016110f6565b935060408601359150808211156112bb57600080fd5b506112c8868287016110f6565b9150509250925092565b6000806000606084860312156112e757600080fd5b8335925060208401359150604084013567ffffffffffffffff81111561130c57600080fd5b6112c8868287016110f6565b801515811461132657600080fd5b50565b6000806040838503121561133c57600080fd5b82359150602083013561134e81611318565b809150509250929050565b600181811c9082168061136d57607f821691505b60208210810361138d57634e487b7160e01b600052602260045260246000fd5b50919050565b60008083546113a181611359565b600182811680156113b957600181146113ce576113fd565b60ff19841687528215158302870194506113fd565b8760005260208060002060005b858110156113f45781548a8201529084019082016113db565b50505082870194505b50929695505050505050565b6020808252601690820152754e6f7420617574686f72697a6564206163636573732160501b604082015260600190565b601f82111561148357600081815260208120601f850160051c810160208610156114605750805b601f850160051c820191505b8181101561147f5782815560010161146c565b5050505b505050565b815167ffffffffffffffff8111156114a2576114a26110e0565b6114b6816114b08454611359565b84611439565b602080601f8311600181146114eb57600084156114d35750858301515b600019600386901b1c1916600185901b17855561147f565b600085815260208120601f198616915b8281101561151a578886015182559484019460019091019084016114fb565b50858210156115385787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000815461155581611359565b808552602060018381168015611572576001811461158c576115ba565b60ff1985168884015283151560051b8801830195506115ba565b866000528260002060005b858110156115b25781548a8201860152908301908401611597565b890184019650505b505050505092915050565b86815260018060a01b038616602082015284604082015260c0606082015260006115f260c0830186611548565b841515608084015282810360a084015261160c81856111e3565b9998505050505050505050565b8082018082111561163a57634e487b7160e01b600052601160045260246000fd5b92915050565b60006020828403121561165257600080fd5b815161123c81611318565b8481526080602082015260006116766080830186611548565b6001600160a01b03949094166040830152506060015292915050565b85815260a0602082015260006116ab60a0830187611548565b6001600160a01b03861660408401526060830185905282810360808401526116d38185611548565b98975050505050505050565b8481526080602082015260006116f86080830186611548565b846040840152828103606084015261171081856111e3565b97965050505050505056fea26469706673582212209d040559a7a24cca328058356707c7347f5e40fbb668b63cdf67a4afe4589ef164736f6c63430008120033",
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
	AuctionId   *big.Int
	AssetId     string
	Bidder      common.Address
	BidAmount   *big.Int
	AuctionType string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0xb95c1199e4385d3e33ad9cffdc96a2f61491b3798965fd93633c02ff7ade77f0.
//
// Solidity: event HighestBidIncreased(uint256 auctionId, string assetId, address bidder, uint256 bidAmount, string auctionType)
func (_EnglishAuction *EnglishAuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*EnglishAuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _EnglishAuction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &EnglishAuctionHighestBidIncreasedIterator{contract: _EnglishAuction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0xb95c1199e4385d3e33ad9cffdc96a2f61491b3798965fd93633c02ff7ade77f0.
//
// Solidity: event HighestBidIncreased(uint256 auctionId, string assetId, address bidder, uint256 bidAmount, string auctionType)
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
// Solidity: event HighestBidIncreased(uint256 auctionId, string assetId, address bidder, uint256 bidAmount, string auctionType)
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
