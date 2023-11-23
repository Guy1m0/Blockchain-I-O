// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_english_auction

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

// EthEnglishAuctionMetaData contains all meta data concerning the EthEnglishAuction contract.
var EthEnglishAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AwaitResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prcd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"DecisionMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"rating\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"review\",\"type\":\"bytes32\"}],\"name\":\"RateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkAverageScore\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"not_winner_platform\",\"type\":\"bool\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_asset_id\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_score\",\"type\":\"int256\"},{\"internalType\":\"bytes32\",\"name\":\"_feedback\",\"type\":\"bytes32\"}],\"name\":\"provide_feedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051611b46380380611b4683398101604081905261002e91610050565b6001600160a01b03166080525f80546001600160a01b0319163317905561007d565b5f60208284031215610060575f80fd5b81516001600160a01b0381168114610076575f80fd5b9392505050565b608051611aa36100a35f395f8181610223015281816107eb0152610ae40152611aa35ff3fe608060405234801561000f575f80fd5b50600436106100e5575f3560e01c80638da5cb5b11610088578063c84d2f6a11610063578063c84d2f6a146101f0578063d0a1414a14610203578063dad2393614610216578063fc0c546a1461021e575f80fd5b80638da5cb5b146101aa578063b14c63c5146101bc578063b6a46b3b146101dd575f80fd5b806342d21ef7116100c357806342d21ef714610139578063451df52e14610159578063598647f8146101845780638d10b0a614610197575f80fd5b8063176321e9146100e9578063274377c0146100fe5780632e1a7d4d14610111575b5f80fd5b6100fc6100f73660046114e4565b610245565b005b6100fc61010c366004611528565b6104fe565b61012461011f366004611551565b6106e0565b60405190151581526020015b60405180910390f35b61014c610147366004611551565b6108d1565b60405161013091906115ab565b61016c610167366004611551565b610977565b6040516001600160a01b039091168152602001610130565b6100fc6101923660046115c4565b61099f565b61014c6101a5366004611551565b610ce6565b5f5461016c906001600160a01b031681565b6101cf6101ca366004611551565b610cf5565b604051908152602001610130565b6100fc6101eb3660046115e4565b610d14565b6100fc6101fe36600461162e565b610edf565b6100fc6102113660046114e4565b611140565b6101cf61136d565b61016c7f000000000000000000000000000000000000000000000000000000000000000081565b60405165656e64696e6760d01b602082015260260160405160208183030381529060405280519060200120600483815481106102835761028361165c565b905f5260205f200160405160200161029b91906116a8565b60405160208183030381529060405280519060200120146103035760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064015b60405180910390fd5b600282815481106103165761031661165c565b5f918252602090912001546001600160a01b031633146103485760405162461bcd60e51b81526004016102fa9061171a565b60405180604001604052806007815260200166636c6f73696e6760c81b8152506004838154811061037b5761037b61165c565b905f5260205f2001908161038f9190611798565b507f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c782600284815481106103c5576103c561165c565b5f91825260209091200154600380546001600160a01b0390921691869081106103f0576103f061165c565b905f5260205f2001546005868154811061040c5761040c61165c565b905f5260205f20015f86604051610428969594939291906118cd565b60405180910390a1600382815481106104435761044361165c565b905f5260205f20015460015f600285815481106104625761046261165c565b5f9182526020808320909101546001600160a01b0316835282019290925260400181208054909190610495908490611934565b925050819055505f600283815481106104b0576104b061165c565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055505f600383815481106104ef576104ef61165c565b5f918252602090912001555050565b60405166636c6f73696e6760c81b6020820152602701604051602081830303815290604052805190602001206004848154811061053d5761053d61165c565b905f5260205f200160405160200161055591906116a8565b60405160208183030381529060405280519060200120146105b85760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e20434c4f53494e4720737461747573000060448201526064016102fa565b600283815481106105cb576105cb61165c565b5f918252602090912001546001600160a01b031633146105fd5760405162461bcd60e51b81526004016102fa9061171a565b81600784815481106106115761061161165c565b905f5260205f20018190555080600684815481106106315761063161165c565b905f5260205f2001819055507f10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e7983600585815481106106725761067261165c565b905f5260205f2001848460405161068c949392919061194d565b60405180910390a16040518060400160405280600681526020016518db1bdcd95960d21b815250600484815481106106c6576106c661165c565b905f5260205f200190816106da9190611798565b50505050565b6040516337b832b760e11b60208201525f90602401604051602081830303815290604052805190602001206004838154811061071e5761071e61165c565b905f5260205f200160405160200161073691906116a8565b60405160208183030381529060405280519060200120036107995760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420696e204f50454e2073746174757300000000000000000060448201526064016102fa565b335f90815260016020526040902054801561087157335f8181526001602052604080822091909155516323b872dd60e01b81523060048201526024810191909152604481018290526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303815f875af1158015610831573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108559190611978565b61087157335f9081526001602052604081209190915592915050565b7f9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c83600585815481106108a6576108a661165c565b905f5260205f200133846040516108c09493929190611993565b60405180910390a150600192915050565b600481815481106108e0575f80fd5b905f5260205f20015f9150905080546108f890611670565b80601f016020809104026020016040519081016040528092919081815260200182805461092490611670565b801561096f5780601f106109465761010080835404028352916020019161096f565b820191905f5260205f20905b81548152906001019060200180831161095257829003601f168201915b505050505081565b60028181548110610986575f80fd5b5f918252602090912001546001600160a01b0316905081565b6040516337b832b760e11b602082015260240160405160208183030381529060405280519060200120600483815481106109db576109db61165c565b905f5260205f20016040516020016109f391906116a8565b6040516020818303038152906040528051906020012014610a565760405162461bcd60e51b815260206004820152601b60248201527f436f6e7472616374206e6f7420696e204f50454e20737461747573000000000060448201526064016102fa565b60038281548110610a6957610a6961165c565b905f5260205f2001548111610ac05760405162461bcd60e51b815260206004820152601e60248201527f546865726520616c7265616479206973206120686967686572206269642e000060448201526064016102fa565b6040516323b872dd60e01b8152336004820152306024820152604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303815f875af1158015610b32573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b569190611978565b905080610b9e5760405162461bcd60e51b81526020600482015260166024820152752a37b5b2b7103a3930b739b332b9103330b4b632b21760511b60448201526064016102fa565b60038381548110610bb157610bb161165c565b905f5260205f2001545f14610c2b5760038381548110610bd357610bd361165c565b905f5260205f20015460015f60028681548110610bf257610bf261165c565b5f9182526020808320909101546001600160a01b0316835282019290925260400181208054909190610c25908490611934565b90915550505b3360028481548110610c3f57610c3f61165c565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055508160038481548110610c7e57610c7e61165c565b905f5260205f2001819055507f463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde088360058581548110610cbf57610cbf61165c565b905f5260205f20013385604051610cd99493929190611993565b60405180910390a1505050565b600581815481106108e0575f80fd5b60038181548110610d04575f80fd5b5f91825260209091200154905081565b5f546001600160a01b03163314610d775760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79206f776e65722063616e20637265617465206e65772061756374696f6044820152603760f91b60648201526084016102fa565b6002805460018181019092557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180546001600160a01b0319169055600380548083019091555f7fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b90910181905560048054928301815590819052604080518082019091529081526337b832b760e11b60208201527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b90910190610e3b9082611798565b50600580546001810182555f919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001610e778282611798565b50506006805460018181019092555f7ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f90910181905560078054928301815581527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68890910155565b5f546001600160a01b03163314610f485760405162461bcd60e51b815260206004820152602760248201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736044820152662073746174757360c81b60648201526084016102fa565b6040516337b832b760e11b60208201526024016040516020818303038152906040528051906020012060048381548110610f8457610f8461165c565b905f5260205f2001604051602001610f9c91906116a8565b6040516020818303038152906040528051906020012014610fff5760405162461bcd60e51b815260206004820152601b60248201527f436f6e7472616374206e6f7420696e204f50454e20737461747573000000000060448201526064016102fa565b60405180604001604052806006815260200165656e64696e6760d01b815250600483815481106110315761103161165c565b905f5260205f200190816110459190611798565b50808061106c5750600382815481106110605761106061165c565b905f5260205f2001545f145b156110cb576040518060400160405280600681526020016518db1bdcd95960d21b815250600483815481106110a3576110a361165c565b905f5260205f200190816110b79190611798565b50600382815481106104435761044361165c565b7fa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b484746850027382600284815481106111005761110061165c565b5f9182526020909120015460405161113492916001600160a01b0316909182526001600160a01b0316602082015260400190565b60405180910390a15050565b60405165656e64696e6760d01b6020820152602601604051602081830303815290604052805190602001206004838154811061117e5761117e61165c565b905f5260205f200160405160200161119691906116a8565b60405160208183030381529060405280519060200120146111f95760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064016102fa565b6002828154811061120c5761120c61165c565b5f918252602090912001546001600160a01b0316331461123e5760405162461bcd60e51b81526004016102fa9061171a565b60405180604001604052806007815260200166636c6f73696e6760c81b815250600483815481106112715761127161165c565b905f5260205f200190816112859190611798565b50600382815481106112995761129961165c565b5f91825260208083209091015482546001600160a01b031683526001909152604082208054919290916112cd908490611934565b925050819055507f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c782600284815481106113095761130961165c565b5f91825260209091200154600380546001600160a01b0390921691869081106113345761133461165c565b905f5260205f200154600586815481106113505761135061165c565b905f5260205f2001600186604051611134969594939291906118cd565b5f80805b600754811015611428576040516518db1bdcd95960d21b602082015260260160405160208183030381529060405280519060200120600482815481106113b9576113b961165c565b905f5260205f20016040516020016113d191906116a8565b604051602081830303815290604052805190602001200361141657600781815481106113ff576113ff61165c565b905f5260205f2001548261141391906119c7565b91505b80611420816119ee565b915050611371565b50600754611437826064611a06565b6114419190611a35565b91505090565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261146a575f80fd5b813567ffffffffffffffff8082111561148557611485611447565b604051601f8301601f19908116603f011681019082821181831017156114ad576114ad611447565b816040528381528660208588010111156114c5575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f80604083850312156114f5575f80fd5b82359150602083013567ffffffffffffffff811115611512575f80fd5b61151e8582860161145b565b9150509250929050565b5f805f6060848603121561153a575f80fd5b505081359360208301359350604090920135919050565b5f60208284031215611561575f80fd5b5035919050565b5f81518084525f5b8181101561158c57602081850181015186830182015201611570565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f6115bd6020830184611568565b9392505050565b5f80604083850312156115d5575f80fd5b50508035926020909101359150565b5f602082840312156115f4575f80fd5b813567ffffffffffffffff81111561160a575f80fd5b6116168482850161145b565b949350505050565b801515811461162b575f80fd5b50565b5f806040838503121561163f575f80fd5b8235915060208301356116518161161e565b809150509250929050565b634e487b7160e01b5f52603260045260245ffd5b600181811c9082168061168457607f821691505b6020821081036116a257634e487b7160e01b5f52602260045260245ffd5b50919050565b5f8083546116b581611670565b600182811680156116cd57600181146116e25761170e565b60ff198416875282151583028701945061170e565b875f526020805f205f5b858110156117055781548a8201529084019082016116ec565b50505082870194505b50929695505050505050565b6020808252601690820152754e6f7420617574686f72697a6564206163636573732160501b604082015260600190565b601f821115611793575f81815260208120601f850160051c810160208610156117705750805b601f850160051c820191505b8181101561178f5782815560010161177c565b5050505b505050565b815167ffffffffffffffff8111156117b2576117b2611447565b6117c6816117c08454611670565b8461174a565b602080601f8311600181146117f9575f84156117e25750858301515b5f19600386901b1c1916600185901b17855561178f565b5f85815260208120601f198616915b8281101561182757888601518255948401946001909101908401611808565b508582101561184457878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f815461186081611670565b80855260206001838116801561187d5760018114611897576118c2565b60ff1985168884015283151560051b8801830195506118c2565b865f52825f205f5b858110156118ba5781548a820186015290830190840161189f565b890184019650505b505050505092915050565b86815260018060a01b038616602082015284604082015260c060608201525f6118f960c0830186611854565b841515608084015282810360a08401526119138185611568565b9998505050505050505050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561194757611947611920565b92915050565b848152608060208201525f6119656080830186611854565b6040830194909452506060015292915050565b5f60208284031215611988575f80fd5b81516115bd8161161e565b848152608060208201525f6119ab6080830186611854565b6001600160a01b03949094166040830152506060015292915050565b8082018281125f8312801582168215821617156119e6576119e6611920565b505092915050565b5f600182016119ff576119ff611920565b5060010190565b8082025f8212600160ff1b84141615611a2157611a21611920565b818105831482151761194757611947611920565b5f82611a4f57634e487b7160e01b5f52601260045260245ffd5b600160ff1b82145f1984141615611a6857611a68611920565b50059056fea264697066735822122031bbb9dc79a75ad03ae2afc1118175df343e8509af54074fbb36dec01a3a9f1864736f6c63430008140033",
}

// EthEnglishAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use EthEnglishAuctionMetaData.ABI instead.
var EthEnglishAuctionABI = EthEnglishAuctionMetaData.ABI

// EthEnglishAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthEnglishAuctionMetaData.Bin instead.
var EthEnglishAuctionBin = EthEnglishAuctionMetaData.Bin

// DeployEthEnglishAuction deploys a new Ethereum contract, binding an instance of EthEnglishAuction to it.
func DeployEthEnglishAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *EthEnglishAuction, error) {
	parsed, err := EthEnglishAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthEnglishAuctionBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthEnglishAuction{EthEnglishAuctionCaller: EthEnglishAuctionCaller{contract: contract}, EthEnglishAuctionTransactor: EthEnglishAuctionTransactor{contract: contract}, EthEnglishAuctionFilterer: EthEnglishAuctionFilterer{contract: contract}}, nil
}

// EthEnglishAuction is an auto generated Go binding around an Ethereum contract.
type EthEnglishAuction struct {
	EthEnglishAuctionCaller     // Read-only binding to the contract
	EthEnglishAuctionTransactor // Write-only binding to the contract
	EthEnglishAuctionFilterer   // Log filterer for contract events
}

// EthEnglishAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthEnglishAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthEnglishAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthEnglishAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthEnglishAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthEnglishAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthEnglishAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthEnglishAuctionSession struct {
	Contract     *EthEnglishAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EthEnglishAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthEnglishAuctionCallerSession struct {
	Contract *EthEnglishAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// EthEnglishAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthEnglishAuctionTransactorSession struct {
	Contract     *EthEnglishAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// EthEnglishAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthEnglishAuctionRaw struct {
	Contract *EthEnglishAuction // Generic contract binding to access the raw methods on
}

// EthEnglishAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthEnglishAuctionCallerRaw struct {
	Contract *EthEnglishAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// EthEnglishAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthEnglishAuctionTransactorRaw struct {
	Contract *EthEnglishAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthEnglishAuction creates a new instance of EthEnglishAuction, bound to a specific deployed contract.
func NewEthEnglishAuction(address common.Address, backend bind.ContractBackend) (*EthEnglishAuction, error) {
	contract, err := bindEthEnglishAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthEnglishAuction{EthEnglishAuctionCaller: EthEnglishAuctionCaller{contract: contract}, EthEnglishAuctionTransactor: EthEnglishAuctionTransactor{contract: contract}, EthEnglishAuctionFilterer: EthEnglishAuctionFilterer{contract: contract}}, nil
}

// NewEthEnglishAuctionCaller creates a new read-only instance of EthEnglishAuction, bound to a specific deployed contract.
func NewEthEnglishAuctionCaller(address common.Address, caller bind.ContractCaller) (*EthEnglishAuctionCaller, error) {
	contract, err := bindEthEnglishAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthEnglishAuctionCaller{contract: contract}, nil
}

// NewEthEnglishAuctionTransactor creates a new write-only instance of EthEnglishAuction, bound to a specific deployed contract.
func NewEthEnglishAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*EthEnglishAuctionTransactor, error) {
	contract, err := bindEthEnglishAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthEnglishAuctionTransactor{contract: contract}, nil
}

// NewEthEnglishAuctionFilterer creates a new log filterer instance of EthEnglishAuction, bound to a specific deployed contract.
func NewEthEnglishAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*EthEnglishAuctionFilterer, error) {
	contract, err := bindEthEnglishAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthEnglishAuctionFilterer{contract: contract}, nil
}

// bindEthEnglishAuction binds a generic wrapper to an already deployed contract.
func bindEthEnglishAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthEnglishAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthEnglishAuction *EthEnglishAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthEnglishAuction.Contract.EthEnglishAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthEnglishAuction *EthEnglishAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.EthEnglishAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthEnglishAuction *EthEnglishAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.EthEnglishAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthEnglishAuction *EthEnglishAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthEnglishAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthEnglishAuction *EthEnglishAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthEnglishAuction *EthEnglishAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.contract.Transact(opts, method, params...)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EthEnglishAuction *EthEnglishAuctionCaller) AssetId(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _EthEnglishAuction.contract.Call(opts, &out, "asset_id", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EthEnglishAuction *EthEnglishAuctionSession) AssetId(arg0 *big.Int) (string, error) {
	return _EthEnglishAuction.Contract.AssetId(&_EthEnglishAuction.CallOpts, arg0)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EthEnglishAuction *EthEnglishAuctionCallerSession) AssetId(arg0 *big.Int) (string, error) {
	return _EthEnglishAuction.Contract.AssetId(&_EthEnglishAuction.CallOpts, arg0)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_EthEnglishAuction *EthEnglishAuctionCaller) CheckAverageScore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthEnglishAuction.contract.Call(opts, &out, "checkAverageScore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_EthEnglishAuction *EthEnglishAuctionSession) CheckAverageScore() (*big.Int, error) {
	return _EthEnglishAuction.Contract.CheckAverageScore(&_EthEnglishAuction.CallOpts)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_EthEnglishAuction *EthEnglishAuctionCallerSession) CheckAverageScore() (*big.Int, error) {
	return _EthEnglishAuction.Contract.CheckAverageScore(&_EthEnglishAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EthEnglishAuction *EthEnglishAuctionCaller) HighestBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EthEnglishAuction.contract.Call(opts, &out, "highestBid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EthEnglishAuction *EthEnglishAuctionSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _EthEnglishAuction.Contract.HighestBid(&_EthEnglishAuction.CallOpts, arg0)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EthEnglishAuction *EthEnglishAuctionCallerSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _EthEnglishAuction.Contract.HighestBid(&_EthEnglishAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EthEnglishAuction *EthEnglishAuctionCaller) HighestBidder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EthEnglishAuction.contract.Call(opts, &out, "highestBidder", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EthEnglishAuction *EthEnglishAuctionSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _EthEnglishAuction.Contract.HighestBidder(&_EthEnglishAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EthEnglishAuction *EthEnglishAuctionCallerSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _EthEnglishAuction.Contract.HighestBidder(&_EthEnglishAuction.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthEnglishAuction *EthEnglishAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthEnglishAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthEnglishAuction *EthEnglishAuctionSession) Owner() (common.Address, error) {
	return _EthEnglishAuction.Contract.Owner(&_EthEnglishAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthEnglishAuction *EthEnglishAuctionCallerSession) Owner() (common.Address, error) {
	return _EthEnglishAuction.Contract.Owner(&_EthEnglishAuction.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EthEnglishAuction *EthEnglishAuctionCaller) Status(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _EthEnglishAuction.contract.Call(opts, &out, "status", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EthEnglishAuction *EthEnglishAuctionSession) Status(arg0 *big.Int) (string, error) {
	return _EthEnglishAuction.Contract.Status(&_EthEnglishAuction.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EthEnglishAuction *EthEnglishAuctionCallerSession) Status(arg0 *big.Int) (string, error) {
	return _EthEnglishAuction.Contract.Status(&_EthEnglishAuction.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthEnglishAuction *EthEnglishAuctionCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthEnglishAuction.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthEnglishAuction *EthEnglishAuctionSession) Token() (common.Address, error) {
	return _EthEnglishAuction.Contract.Token(&_EthEnglishAuction.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthEnglishAuction *EthEnglishAuctionCallerSession) Token() (common.Address, error) {
	return _EthEnglishAuction.Contract.Token(&_EthEnglishAuction.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EthEnglishAuction *EthEnglishAuctionTransactor) Abort(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthEnglishAuction.contract.Transact(opts, "abort", auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EthEnglishAuction *EthEnglishAuctionSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.Abort(&_EthEnglishAuction.TransactOpts, auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EthEnglishAuction *EthEnglishAuctionTransactorSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.Abort(&_EthEnglishAuction.TransactOpts, auctionId, jsonString)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 auctionId, uint256 bidAmount) returns()
func (_EthEnglishAuction *EthEnglishAuctionTransactor) Bid(opts *bind.TransactOpts, auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EthEnglishAuction.contract.Transact(opts, "bid", auctionId, bidAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 auctionId, uint256 bidAmount) returns()
func (_EthEnglishAuction *EthEnglishAuctionSession) Bid(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.Bid(&_EthEnglishAuction.TransactOpts, auctionId, bidAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 auctionId, uint256 bidAmount) returns()
func (_EthEnglishAuction *EthEnglishAuctionTransactorSession) Bid(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.Bid(&_EthEnglishAuction.TransactOpts, auctionId, bidAmount)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_EthEnglishAuction *EthEnglishAuctionTransactor) CloseAuction(opts *bind.TransactOpts, auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _EthEnglishAuction.contract.Transact(opts, "closeAuction", auctionId, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_EthEnglishAuction *EthEnglishAuctionSession) CloseAuction(auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.CloseAuction(&_EthEnglishAuction.TransactOpts, auctionId, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_EthEnglishAuction *EthEnglishAuctionTransactorSession) CloseAuction(auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.CloseAuction(&_EthEnglishAuction.TransactOpts, auctionId, not_winner_platform)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EthEnglishAuction *EthEnglishAuctionTransactor) Commit(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthEnglishAuction.contract.Transact(opts, "commit", auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EthEnglishAuction *EthEnglishAuctionSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.Commit(&_EthEnglishAuction.TransactOpts, auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EthEnglishAuction *EthEnglishAuctionTransactorSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.Commit(&_EthEnglishAuction.TransactOpts, auctionId, jsonString)
}

// Create is a paid mutator transaction binding the contract method 0xb6a46b3b.
//
// Solidity: function create(string _asset_id) returns()
func (_EthEnglishAuction *EthEnglishAuctionTransactor) Create(opts *bind.TransactOpts, _asset_id string) (*types.Transaction, error) {
	return _EthEnglishAuction.contract.Transact(opts, "create", _asset_id)
}

// Create is a paid mutator transaction binding the contract method 0xb6a46b3b.
//
// Solidity: function create(string _asset_id) returns()
func (_EthEnglishAuction *EthEnglishAuctionSession) Create(_asset_id string) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.Create(&_EthEnglishAuction.TransactOpts, _asset_id)
}

// Create is a paid mutator transaction binding the contract method 0xb6a46b3b.
//
// Solidity: function create(string _asset_id) returns()
func (_EthEnglishAuction *EthEnglishAuctionTransactorSession) Create(_asset_id string) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.Create(&_EthEnglishAuction.TransactOpts, _asset_id)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_EthEnglishAuction *EthEnglishAuctionTransactor) ProvideFeedback(opts *bind.TransactOpts, auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _EthEnglishAuction.contract.Transact(opts, "provide_feedback", auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_EthEnglishAuction *EthEnglishAuctionSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.ProvideFeedback(&_EthEnglishAuction.TransactOpts, auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_EthEnglishAuction *EthEnglishAuctionTransactorSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.ProvideFeedback(&_EthEnglishAuction.TransactOpts, auctionId, _score, _feedback)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EthEnglishAuction *EthEnglishAuctionTransactor) Withdraw(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _EthEnglishAuction.contract.Transact(opts, "withdraw", auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EthEnglishAuction *EthEnglishAuctionSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.Withdraw(&_EthEnglishAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EthEnglishAuction *EthEnglishAuctionTransactorSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _EthEnglishAuction.Contract.Withdraw(&_EthEnglishAuction.TransactOpts, auctionId)
}

// EthEnglishAuctionAwaitResponseIterator is returned from FilterAwaitResponse and is used to iterate over the raw logs and unpacked data for AwaitResponse events raised by the EthEnglishAuction contract.
type EthEnglishAuctionAwaitResponseIterator struct {
	Event *EthEnglishAuctionAwaitResponse // Event containing the contract specifics and raw log

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
func (it *EthEnglishAuctionAwaitResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthEnglishAuctionAwaitResponse)
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
		it.Event = new(EthEnglishAuctionAwaitResponse)
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
func (it *EthEnglishAuctionAwaitResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthEnglishAuctionAwaitResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthEnglishAuctionAwaitResponse represents a AwaitResponse event raised by the EthEnglishAuction contract.
type EthEnglishAuctionAwaitResponse struct {
	Auction *big.Int
	Winner  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAwaitResponse is a free log retrieval operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_EthEnglishAuction *EthEnglishAuctionFilterer) FilterAwaitResponse(opts *bind.FilterOpts) (*EthEnglishAuctionAwaitResponseIterator, error) {

	logs, sub, err := _EthEnglishAuction.contract.FilterLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return &EthEnglishAuctionAwaitResponseIterator{contract: _EthEnglishAuction.contract, event: "AwaitResponse", logs: logs, sub: sub}, nil
}

// WatchAwaitResponse is a free log subscription operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_EthEnglishAuction *EthEnglishAuctionFilterer) WatchAwaitResponse(opts *bind.WatchOpts, sink chan<- *EthEnglishAuctionAwaitResponse) (event.Subscription, error) {

	logs, sub, err := _EthEnglishAuction.contract.WatchLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthEnglishAuctionAwaitResponse)
				if err := _EthEnglishAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
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
func (_EthEnglishAuction *EthEnglishAuctionFilterer) ParseAwaitResponse(log types.Log) (*EthEnglishAuctionAwaitResponse, error) {
	event := new(EthEnglishAuctionAwaitResponse)
	if err := _EthEnglishAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthEnglishAuctionDecisionMadeIterator is returned from FilterDecisionMade and is used to iterate over the raw logs and unpacked data for DecisionMade events raised by the EthEnglishAuction contract.
type EthEnglishAuctionDecisionMadeIterator struct {
	Event *EthEnglishAuctionDecisionMade // Event containing the contract specifics and raw log

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
func (it *EthEnglishAuctionDecisionMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthEnglishAuctionDecisionMade)
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
		it.Event = new(EthEnglishAuctionDecisionMade)
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
func (it *EthEnglishAuctionDecisionMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthEnglishAuctionDecisionMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthEnglishAuctionDecisionMade represents a DecisionMade event raised by the EthEnglishAuction contract.
type EthEnglishAuctionDecisionMade struct {
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
func (_EthEnglishAuction *EthEnglishAuctionFilterer) FilterDecisionMade(opts *bind.FilterOpts) (*EthEnglishAuctionDecisionMadeIterator, error) {

	logs, sub, err := _EthEnglishAuction.contract.FilterLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return &EthEnglishAuctionDecisionMadeIterator{contract: _EthEnglishAuction.contract, event: "DecisionMade", logs: logs, sub: sub}, nil
}

// WatchDecisionMade is a free log subscription operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auction, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_EthEnglishAuction *EthEnglishAuctionFilterer) WatchDecisionMade(opts *bind.WatchOpts, sink chan<- *EthEnglishAuctionDecisionMade) (event.Subscription, error) {

	logs, sub, err := _EthEnglishAuction.contract.WatchLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthEnglishAuctionDecisionMade)
				if err := _EthEnglishAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
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
func (_EthEnglishAuction *EthEnglishAuctionFilterer) ParseDecisionMade(log types.Log) (*EthEnglishAuctionDecisionMade, error) {
	event := new(EthEnglishAuctionDecisionMade)
	if err := _EthEnglishAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthEnglishAuctionHighestBidIncreasedIterator is returned from FilterHighestBidIncreased and is used to iterate over the raw logs and unpacked data for HighestBidIncreased events raised by the EthEnglishAuction contract.
type EthEnglishAuctionHighestBidIncreasedIterator struct {
	Event *EthEnglishAuctionHighestBidIncreased // Event containing the contract specifics and raw log

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
func (it *EthEnglishAuctionHighestBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthEnglishAuctionHighestBidIncreased)
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
		it.Event = new(EthEnglishAuctionHighestBidIncreased)
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
func (it *EthEnglishAuctionHighestBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthEnglishAuctionHighestBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthEnglishAuctionHighestBidIncreased represents a HighestBidIncreased event raised by the EthEnglishAuction contract.
type EthEnglishAuctionHighestBidIncreased struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_EthEnglishAuction *EthEnglishAuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*EthEnglishAuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _EthEnglishAuction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &EthEnglishAuctionHighestBidIncreasedIterator{contract: _EthEnglishAuction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_EthEnglishAuction *EthEnglishAuctionFilterer) WatchHighestBidIncreased(opts *bind.WatchOpts, sink chan<- *EthEnglishAuctionHighestBidIncreased) (event.Subscription, error) {

	logs, sub, err := _EthEnglishAuction.contract.WatchLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthEnglishAuctionHighestBidIncreased)
				if err := _EthEnglishAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
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
func (_EthEnglishAuction *EthEnglishAuctionFilterer) ParseHighestBidIncreased(log types.Log) (*EthEnglishAuctionHighestBidIncreased, error) {
	event := new(EthEnglishAuctionHighestBidIncreased)
	if err := _EthEnglishAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthEnglishAuctionRateAuctionIterator is returned from FilterRateAuction and is used to iterate over the raw logs and unpacked data for RateAuction events raised by the EthEnglishAuction contract.
type EthEnglishAuctionRateAuctionIterator struct {
	Event *EthEnglishAuctionRateAuction // Event containing the contract specifics and raw log

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
func (it *EthEnglishAuctionRateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthEnglishAuctionRateAuction)
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
		it.Event = new(EthEnglishAuctionRateAuction)
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
func (it *EthEnglishAuctionRateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthEnglishAuctionRateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthEnglishAuctionRateAuction represents a RateAuction event raised by the EthEnglishAuction contract.
type EthEnglishAuctionRateAuction struct {
	Auction *big.Int
	Id      string
	Rating  *big.Int
	Review  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRateAuction is a free log retrieval operation binding the contract event 0x10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e79.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, bytes32 review)
func (_EthEnglishAuction *EthEnglishAuctionFilterer) FilterRateAuction(opts *bind.FilterOpts) (*EthEnglishAuctionRateAuctionIterator, error) {

	logs, sub, err := _EthEnglishAuction.contract.FilterLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return &EthEnglishAuctionRateAuctionIterator{contract: _EthEnglishAuction.contract, event: "RateAuction", logs: logs, sub: sub}, nil
}

// WatchRateAuction is a free log subscription operation binding the contract event 0x10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e79.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, bytes32 review)
func (_EthEnglishAuction *EthEnglishAuctionFilterer) WatchRateAuction(opts *bind.WatchOpts, sink chan<- *EthEnglishAuctionRateAuction) (event.Subscription, error) {

	logs, sub, err := _EthEnglishAuction.contract.WatchLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthEnglishAuctionRateAuction)
				if err := _EthEnglishAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
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

// ParseRateAuction is a log parse operation binding the contract event 0x10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e79.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, bytes32 review)
func (_EthEnglishAuction *EthEnglishAuctionFilterer) ParseRateAuction(log types.Log) (*EthEnglishAuctionRateAuction, error) {
	event := new(EthEnglishAuctionRateAuction)
	if err := _EthEnglishAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthEnglishAuctionWithdrawBidIterator is returned from FilterWithdrawBid and is used to iterate over the raw logs and unpacked data for WithdrawBid events raised by the EthEnglishAuction contract.
type EthEnglishAuctionWithdrawBidIterator struct {
	Event *EthEnglishAuctionWithdrawBid // Event containing the contract specifics and raw log

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
func (it *EthEnglishAuctionWithdrawBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthEnglishAuctionWithdrawBid)
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
		it.Event = new(EthEnglishAuctionWithdrawBid)
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
func (it *EthEnglishAuctionWithdrawBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthEnglishAuctionWithdrawBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthEnglishAuctionWithdrawBid represents a WithdrawBid event raised by the EthEnglishAuction contract.
type EthEnglishAuctionWithdrawBid struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBid is a free log retrieval operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_EthEnglishAuction *EthEnglishAuctionFilterer) FilterWithdrawBid(opts *bind.FilterOpts) (*EthEnglishAuctionWithdrawBidIterator, error) {

	logs, sub, err := _EthEnglishAuction.contract.FilterLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return &EthEnglishAuctionWithdrawBidIterator{contract: _EthEnglishAuction.contract, event: "WithdrawBid", logs: logs, sub: sub}, nil
}

// WatchWithdrawBid is a free log subscription operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_EthEnglishAuction *EthEnglishAuctionFilterer) WatchWithdrawBid(opts *bind.WatchOpts, sink chan<- *EthEnglishAuctionWithdrawBid) (event.Subscription, error) {

	logs, sub, err := _EthEnglishAuction.contract.WatchLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthEnglishAuctionWithdrawBid)
				if err := _EthEnglishAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
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
func (_EthEnglishAuction *EthEnglishAuctionFilterer) ParseWithdrawBid(log types.Log) (*EthEnglishAuctionWithdrawBid, error) {
	event := new(EthEnglishAuctionWithdrawBid)
	if err := _EthEnglishAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
