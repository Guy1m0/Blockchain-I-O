// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_closed_bid_first_auction

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

// EthClosedBidFirstAuctionMetaData contains all meta data concerning the EthClosedBidFirstAuction contract.
var EthClosedBidFirstAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AwaitResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prcd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"DecisionMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"NewBidHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"rating\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"review\",\"type\":\"bytes32\"}],\"name\":\"RateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"}],\"name\":\"RevealAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkAverageScore\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"not_winner_platform\",\"type\":\"bool\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_asset_id\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_score\",\"type\":\"int256\"},{\"internalType\":\"bytes32\",\"name\":\"_feedback\",\"type\":\"bytes32\"}],\"name\":\"provide_feedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"revealAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051611dc3380380611dc383398101604081905261002e91610050565b6001600160a01b03166080525f80546001600160a01b0319163317905561007d565b5f60208284031215610060575f80fd5b81516001600160a01b0381168114610076575f80fd5b9392505050565b608051611d206100a35f395f818161025f015281816108270152610c710152611d205ff3fe608060405234801561000f575f80fd5b50600436106100fb575f3560e01c80639348cef711610093578063d0a1414a11610063578063d0a1414a1461022c578063dad239361461023f578063ea1591bb14610247578063fc0c546a1461025a575f80fd5b80639348cef7146101d2578063b14c63c5146101e5578063b6a46b3b14610206578063c84d2f6a14610219575f80fd5b8063451df52e116100ce578063451df52e1461016f57806355f78c8d1461019a5780638d10b0a6146101ad5780638da5cb5b146101c0575f80fd5b8063176321e9146100ff578063274377c0146101145780632e1a7d4d1461012757806342d21ef71461014f575b5f80fd5b61011261010d3660046116e3565b610281565b005b610112610122366004611727565b61053a565b61013a610135366004611750565b61071c565b60405190151581526020015b60405180910390f35b61016261015d366004611750565b61090d565b60405161014691906117aa565b61018261017d366004611750565b6109b3565b6040516001600160a01b039091168152602001610146565b6101126101a8366004611750565b6109db565b6101626101bb366004611750565b610b08565b5f54610182906001600160a01b031681565b6101126101e03660046117c3565b610b17565b6101f86101f3366004611750565b610e73565b604051908152602001610146565b6101126102143660046117e3565b610e92565b61011261022736600461182d565b61105a565b61011261023a3660046116e3565b61124b565b6101f8611478565b6101126102553660046117c3565b611552565b6101827f000000000000000000000000000000000000000000000000000000000000000081565b60405165656e64696e6760d01b602082015260260160405160208183030381529060405280519060200120600583815481106102bf576102bf61185b565b905f5260205f20016040516020016102d791906118a7565b604051602081830303815290604052805190602001201461033f5760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064015b60405180910390fd5b600382815481106103525761035261185b565b5f918252602090912001546001600160a01b031633146103845760405162461bcd60e51b815260040161033690611919565b60405180604001604052806007815260200166636c6f73696e6760c81b815250600583815481106103b7576103b761185b565b905f5260205f200190816103cb9190611997565b507f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c782600384815481106104015761040161185b565b5f91825260209091200154600480546001600160a01b03909216918690811061042c5761042c61185b565b905f5260205f200154600686815481106104485761044861185b565b905f5260205f20015f8660405161046496959493929190611acc565b60405180910390a16004828154811061047f5761047f61185b565b905f5260205f20015460015f6003858154811061049e5761049e61185b565b5f9182526020808320909101546001600160a01b03168352820192909252604001812080549091906104d1908490611b33565b925050819055505f600383815481106104ec576104ec61185b565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055505f6004838154811061052b5761052b61185b565b5f918252602090912001555050565b60405166636c6f73696e6760c81b602082015260270160405160208183030381529060405280519060200120600584815481106105795761057961185b565b905f5260205f200160405160200161059191906118a7565b60405160208183030381529060405280519060200120146105f45760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e20434c4f53494e472073746174757300006044820152606401610336565b600383815481106106075761060761185b565b5f918252602090912001546001600160a01b031633146106395760405162461bcd60e51b815260040161033690611919565b816008848154811061064d5761064d61185b565b905f5260205f200181905550806007848154811061066d5761066d61185b565b905f5260205f2001819055507f10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e7983600685815481106106ae576106ae61185b565b905f5260205f200184846040516106c89493929190611b4c565b60405180910390a16040518060400160405280600681526020016518db1bdcd95960d21b815250600584815481106107025761070261185b565b905f5260205f200190816107169190611997565b50505050565b6040516337b832b760e11b60208201525f90602401604051602081830303815290604052805190602001206005838154811061075a5761075a61185b565b905f5260205f200160405160200161077291906118a7565b60405160208183030381529060405280519060200120036107d55760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420696e204f50454e207374617475730000000000000000006044820152606401610336565b335f9081526001602052604090205480156108ad57335f8181526001602052604080822091909155516323b872dd60e01b81523060048201526024810191909152604481018290526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303815f875af115801561086d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108919190611b77565b6108ad57335f9081526001602052604081209190915592915050565b7f9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c83600685815481106108e2576108e261185b565b905f5260205f200133846040516108fc9493929190611b92565b60405180910390a150600192915050565b6005818154811061091c575f80fd5b905f5260205f20015f9150905080546109349061186f565b80601f01602080910402602001604051908101604052809291908181526020018280546109609061186f565b80156109ab5780601f10610982576101008083540402835291602001916109ab565b820191905f5260205f20905b81548152906001019060200180831161098e57829003601f168201915b505050505081565b600381815481106109c2575f80fd5b5f918252602090912001546001600160a01b0316905081565b5f546001600160a01b03163314610a045760405162461bcd60e51b815260040161033690611bc6565b6040516337b832b760e11b60208201526024016040516020818303038152906040528051906020012060058281548110610a4057610a4061185b565b905f5260205f2001604051602001610a5891906118a7565b6040516020818303038152906040528051906020012014610a8b5760405162461bcd60e51b815260040161033690611c0d565b604051806040016040528060068152602001651c995d99585b60d21b81525060058281548110610abd57610abd61185b565b905f5260205f20019081610ad19190611997565b506040518181527f3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba9060200160405180910390a150565b6006818154811061091c575f80fd5b604051651c995d99585b60d21b60208201526026016040516020818303038152906040528051906020012060058381548110610b5557610b5561185b565b905f5260205f2001604051602001610b6d91906118a7565b6040516020818303038152906040528051906020012014610ba05760405162461bcd60e51b815260040161033690611c0d565b60048281548110610bb357610bb361185b565b905f5260205f2001548111610c0a5760405162461bcd60e51b815260206004820152601e60248201527f546865726520616c7265616479206973206120686967686572206269642e00006044820152606401610336565b5f82815260026020908152604080832033845282529182902054825191820184905291016040516020818303038152906040528051906020012014610c4d575f80fd5b6040516323b872dd60e01b8152336004820152306024820152604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303815f875af1158015610cbf573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ce39190611b77565b905080610d2b5760405162461bcd60e51b81526020600482015260166024820152752a37b5b2b7103a3930b739b332b9103330b4b632b21760511b6044820152606401610336565b60048381548110610d3e57610d3e61185b565b905f5260205f2001545f14610db85760048381548110610d6057610d6061185b565b905f5260205f20015460015f60038681548110610d7f57610d7f61185b565b5f9182526020808320909101546001600160a01b0316835282019290925260400181208054909190610db2908490611b33565b90915550505b3360038481548110610dcc57610dcc61185b565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055508160048481548110610e0b57610e0b61185b565b905f5260205f2001819055507f463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde088360068581548110610e4c57610e4c61185b565b905f5260205f20013385604051610e669493929190611b92565b60405180910390a1505050565b60048181548110610e82575f80fd5b5f91825260209091200154905081565b5f546001600160a01b03163314610ef55760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79206f776e65722063616e20637265617465206e65772061756374696f6044820152603760f91b6064820152608401610336565b6003805460018181019092557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b03191690556004805480830182555f7f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9091018190556005805493840181559052604080518082019091529081526337b832b760e11b60208201527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db090910190610fb69082611997565b50600680546001810182555f919091527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f01610ff28282611997565b50506007805460018181019092555f7fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68890910181905560088054928301815581527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390910155565b5f546001600160a01b031633146110835760405162461bcd60e51b815260040161033690611bc6565b6040516337b832b760e11b602082015260240160405160208183030381529060405280519060200120600583815481106110bf576110bf61185b565b905f5260205f20016040516020016110d791906118a7565b604051602081830303815290604052805190602001201461110a5760405162461bcd60e51b815260040161033690611c0d565b60405180604001604052806006815260200165656e64696e6760d01b8152506005838154811061113c5761113c61185b565b905f5260205f200190816111509190611997565b50808061117757506004828154811061116b5761116b61185b565b905f5260205f2001545f145b156111d6576040518060400160405280600681526020016518db1bdcd95960d21b815250600583815481106111ae576111ae61185b565b905f5260205f200190816111c29190611997565b506004828154811061047f5761047f61185b565b7fa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273826003848154811061120b5761120b61185b565b5f9182526020909120015460405161123f92916001600160a01b0316909182526001600160a01b0316602082015260400190565b60405180910390a15050565b60405165656e64696e6760d01b602082015260260160405160208183030381529060405280519060200120600583815481106112895761128961185b565b905f5260205f20016040516020016112a191906118a7565b60405160208183030381529060405280519060200120146113045760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e47207374617475730000006044820152606401610336565b600382815481106113175761131761185b565b5f918252602090912001546001600160a01b031633146113495760405162461bcd60e51b815260040161033690611919565b60405180604001604052806007815260200166636c6f73696e6760c81b8152506005838154811061137c5761137c61185b565b905f5260205f200190816113909190611997565b50600482815481106113a4576113a461185b565b5f91825260208083209091015482546001600160a01b031683526001909152604082208054919290916113d8908490611b33565b925050819055507f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c782600384815481106114145761141461185b565b5f91825260209091200154600480546001600160a01b03909216918690811061143f5761143f61185b565b905f5260205f2001546006868154811061145b5761145b61185b565b905f5260205f200160018660405161123f96959493929190611acc565b5f80805b600854811015611533576040516518db1bdcd95960d21b602082015260260160405160208183030381529060405280519060200120600582815481106114c4576114c461185b565b905f5260205f20016040516020016114dc91906118a7565b6040516020818303038152906040528051906020012003611521576008818154811061150a5761150a61185b565b905f5260205f2001548261151e9190611c44565b91505b8061152b81611c6b565b91505061147c565b50600854611542826064611c83565b61154c9190611cb2565b91505090565b6040516337b832b760e11b6020820152602401604051602081830303815290604052805190602001206005838154811061158e5761158e61185b565b905f5260205f20016040516020016115a691906118a7565b60405160208183030381529060405280519060200120146115d95760405162461bcd60e51b815260040161033690611c0d565b5f8281526002602090815260408083203384529091529020819055600680547f6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f9184918290811061162c5761162c61185b565b905f5260205f2001338460405161123f9493929190611b92565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112611669575f80fd5b813567ffffffffffffffff8082111561168457611684611646565b604051601f8301601f19908116603f011681019082821181831017156116ac576116ac611646565b816040528381528660208588010111156116c4575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f80604083850312156116f4575f80fd5b82359150602083013567ffffffffffffffff811115611711575f80fd5b61171d8582860161165a565b9150509250929050565b5f805f60608486031215611739575f80fd5b505081359360208301359350604090920135919050565b5f60208284031215611760575f80fd5b5035919050565b5f81518084525f5b8181101561178b5760208185018101518683018201520161176f565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f6117bc6020830184611767565b9392505050565b5f80604083850312156117d4575f80fd5b50508035926020909101359150565b5f602082840312156117f3575f80fd5b813567ffffffffffffffff811115611809575f80fd5b6118158482850161165a565b949350505050565b801515811461182a575f80fd5b50565b5f806040838503121561183e575f80fd5b8235915060208301356118508161181d565b809150509250929050565b634e487b7160e01b5f52603260045260245ffd5b600181811c9082168061188357607f821691505b6020821081036118a157634e487b7160e01b5f52602260045260245ffd5b50919050565b5f8083546118b48161186f565b600182811680156118cc57600181146118e15761190d565b60ff198416875282151583028701945061190d565b875f526020805f205f5b858110156119045781548a8201529084019082016118eb565b50505082870194505b50929695505050505050565b6020808252601690820152754e6f7420617574686f72697a6564206163636573732160501b604082015260600190565b601f821115611992575f81815260208120601f850160051c8101602086101561196f5750805b601f850160051c820191505b8181101561198e5782815560010161197b565b5050505b505050565b815167ffffffffffffffff8111156119b1576119b1611646565b6119c5816119bf845461186f565b84611949565b602080601f8311600181146119f8575f84156119e15750858301515b5f19600386901b1c1916600185901b17855561198e565b5f85815260208120601f198616915b82811015611a2657888601518255948401946001909101908401611a07565b5085821015611a4357878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f8154611a5f8161186f565b808552602060018381168015611a7c5760018114611a9657611ac1565b60ff1985168884015283151560051b880183019550611ac1565b865f52825f205f5b85811015611ab95781548a8201860152908301908401611a9e565b890184019650505b505050505092915050565b86815260018060a01b038616602082015284604082015260c060608201525f611af860c0830186611a53565b841515608084015282810360a0840152611b128185611767565b9998505050505050505050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115611b4657611b46611b1f565b92915050565b848152608060208201525f611b646080830186611a53565b6040830194909452506060015292915050565b5f60208284031215611b87575f80fd5b81516117bc8161181d565b848152608060208201525f611baa6080830186611a53565b6001600160a01b03949094166040830152506060015292915050565b60208082526027908201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736040820152662073746174757360c81b606082015260800190565b6020808252601b908201527f436f6e7472616374206e6f7420696e204f50454e207374617475730000000000604082015260600190565b8082018281125f831280158216821582161715611c6357611c63611b1f565b505092915050565b5f60018201611c7c57611c7c611b1f565b5060010190565b8082025f8212600160ff1b84141615611c9e57611c9e611b1f565b8181058314821517611b4657611b46611b1f565b5f82611ccc57634e487b7160e01b5f52601260045260245ffd5b600160ff1b82145f1984141615611ce557611ce5611b1f565b50059056fea2646970667358221220c263b58818786629a6c7f1ae5581a5449056b158816f2d23265e88c2d2de820364736f6c63430008140033",
}

// EthClosedBidFirstAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use EthClosedBidFirstAuctionMetaData.ABI instead.
var EthClosedBidFirstAuctionABI = EthClosedBidFirstAuctionMetaData.ABI

// EthClosedBidFirstAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthClosedBidFirstAuctionMetaData.Bin instead.
var EthClosedBidFirstAuctionBin = EthClosedBidFirstAuctionMetaData.Bin

// DeployEthClosedBidFirstAuction deploys a new Ethereum contract, binding an instance of EthClosedBidFirstAuction to it.
func DeployEthClosedBidFirstAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *EthClosedBidFirstAuction, error) {
	parsed, err := EthClosedBidFirstAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthClosedBidFirstAuctionBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthClosedBidFirstAuction{EthClosedBidFirstAuctionCaller: EthClosedBidFirstAuctionCaller{contract: contract}, EthClosedBidFirstAuctionTransactor: EthClosedBidFirstAuctionTransactor{contract: contract}, EthClosedBidFirstAuctionFilterer: EthClosedBidFirstAuctionFilterer{contract: contract}}, nil
}

// EthClosedBidFirstAuction is an auto generated Go binding around an Ethereum contract.
type EthClosedBidFirstAuction struct {
	EthClosedBidFirstAuctionCaller     // Read-only binding to the contract
	EthClosedBidFirstAuctionTransactor // Write-only binding to the contract
	EthClosedBidFirstAuctionFilterer   // Log filterer for contract events
}

// EthClosedBidFirstAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthClosedBidFirstAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClosedBidFirstAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthClosedBidFirstAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClosedBidFirstAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthClosedBidFirstAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClosedBidFirstAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthClosedBidFirstAuctionSession struct {
	Contract     *EthClosedBidFirstAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EthClosedBidFirstAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthClosedBidFirstAuctionCallerSession struct {
	Contract *EthClosedBidFirstAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// EthClosedBidFirstAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthClosedBidFirstAuctionTransactorSession struct {
	Contract     *EthClosedBidFirstAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// EthClosedBidFirstAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthClosedBidFirstAuctionRaw struct {
	Contract *EthClosedBidFirstAuction // Generic contract binding to access the raw methods on
}

// EthClosedBidFirstAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthClosedBidFirstAuctionCallerRaw struct {
	Contract *EthClosedBidFirstAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// EthClosedBidFirstAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthClosedBidFirstAuctionTransactorRaw struct {
	Contract *EthClosedBidFirstAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthClosedBidFirstAuction creates a new instance of EthClosedBidFirstAuction, bound to a specific deployed contract.
func NewEthClosedBidFirstAuction(address common.Address, backend bind.ContractBackend) (*EthClosedBidFirstAuction, error) {
	contract, err := bindEthClosedBidFirstAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthClosedBidFirstAuction{EthClosedBidFirstAuctionCaller: EthClosedBidFirstAuctionCaller{contract: contract}, EthClosedBidFirstAuctionTransactor: EthClosedBidFirstAuctionTransactor{contract: contract}, EthClosedBidFirstAuctionFilterer: EthClosedBidFirstAuctionFilterer{contract: contract}}, nil
}

// NewEthClosedBidFirstAuctionCaller creates a new read-only instance of EthClosedBidFirstAuction, bound to a specific deployed contract.
func NewEthClosedBidFirstAuctionCaller(address common.Address, caller bind.ContractCaller) (*EthClosedBidFirstAuctionCaller, error) {
	contract, err := bindEthClosedBidFirstAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthClosedBidFirstAuctionCaller{contract: contract}, nil
}

// NewEthClosedBidFirstAuctionTransactor creates a new write-only instance of EthClosedBidFirstAuction, bound to a specific deployed contract.
func NewEthClosedBidFirstAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*EthClosedBidFirstAuctionTransactor, error) {
	contract, err := bindEthClosedBidFirstAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthClosedBidFirstAuctionTransactor{contract: contract}, nil
}

// NewEthClosedBidFirstAuctionFilterer creates a new log filterer instance of EthClosedBidFirstAuction, bound to a specific deployed contract.
func NewEthClosedBidFirstAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*EthClosedBidFirstAuctionFilterer, error) {
	contract, err := bindEthClosedBidFirstAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthClosedBidFirstAuctionFilterer{contract: contract}, nil
}

// bindEthClosedBidFirstAuction binds a generic wrapper to an already deployed contract.
func bindEthClosedBidFirstAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthClosedBidFirstAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthClosedBidFirstAuction.Contract.EthClosedBidFirstAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.EthClosedBidFirstAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.EthClosedBidFirstAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthClosedBidFirstAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.contract.Transact(opts, method, params...)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCaller) AssetId(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _EthClosedBidFirstAuction.contract.Call(opts, &out, "asset_id", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) AssetId(arg0 *big.Int) (string, error) {
	return _EthClosedBidFirstAuction.Contract.AssetId(&_EthClosedBidFirstAuction.CallOpts, arg0)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCallerSession) AssetId(arg0 *big.Int) (string, error) {
	return _EthClosedBidFirstAuction.Contract.AssetId(&_EthClosedBidFirstAuction.CallOpts, arg0)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCaller) CheckAverageScore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthClosedBidFirstAuction.contract.Call(opts, &out, "checkAverageScore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) CheckAverageScore() (*big.Int, error) {
	return _EthClosedBidFirstAuction.Contract.CheckAverageScore(&_EthClosedBidFirstAuction.CallOpts)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCallerSession) CheckAverageScore() (*big.Int, error) {
	return _EthClosedBidFirstAuction.Contract.CheckAverageScore(&_EthClosedBidFirstAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCaller) HighestBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EthClosedBidFirstAuction.contract.Call(opts, &out, "highestBid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _EthClosedBidFirstAuction.Contract.HighestBid(&_EthClosedBidFirstAuction.CallOpts, arg0)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCallerSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _EthClosedBidFirstAuction.Contract.HighestBid(&_EthClosedBidFirstAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCaller) HighestBidder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EthClosedBidFirstAuction.contract.Call(opts, &out, "highestBidder", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _EthClosedBidFirstAuction.Contract.HighestBidder(&_EthClosedBidFirstAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCallerSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _EthClosedBidFirstAuction.Contract.HighestBidder(&_EthClosedBidFirstAuction.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthClosedBidFirstAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) Owner() (common.Address, error) {
	return _EthClosedBidFirstAuction.Contract.Owner(&_EthClosedBidFirstAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCallerSession) Owner() (common.Address, error) {
	return _EthClosedBidFirstAuction.Contract.Owner(&_EthClosedBidFirstAuction.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCaller) Status(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _EthClosedBidFirstAuction.contract.Call(opts, &out, "status", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) Status(arg0 *big.Int) (string, error) {
	return _EthClosedBidFirstAuction.Contract.Status(&_EthClosedBidFirstAuction.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCallerSession) Status(arg0 *big.Int) (string, error) {
	return _EthClosedBidFirstAuction.Contract.Status(&_EthClosedBidFirstAuction.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthClosedBidFirstAuction.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) Token() (common.Address, error) {
	return _EthClosedBidFirstAuction.Contract.Token(&_EthClosedBidFirstAuction.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionCallerSession) Token() (common.Address, error) {
	return _EthClosedBidFirstAuction.Contract.Token(&_EthClosedBidFirstAuction.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactor) Abort(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.contract.Transact(opts, "abort", auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.Abort(&_EthClosedBidFirstAuction.TransactOpts, auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactorSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.Abort(&_EthClosedBidFirstAuction.TransactOpts, auctionId, jsonString)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactor) Bid(opts *bind.TransactOpts, auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.contract.Transact(opts, "bid", auctionId, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) Bid(auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.Bid(&_EthClosedBidFirstAuction.TransactOpts, auctionId, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactorSession) Bid(auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.Bid(&_EthClosedBidFirstAuction.TransactOpts, auctionId, bidHash)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactor) CloseAuction(opts *bind.TransactOpts, auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.contract.Transact(opts, "closeAuction", auctionId, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) CloseAuction(auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.CloseAuction(&_EthClosedBidFirstAuction.TransactOpts, auctionId, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactorSession) CloseAuction(auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.CloseAuction(&_EthClosedBidFirstAuction.TransactOpts, auctionId, not_winner_platform)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactor) Commit(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.contract.Transact(opts, "commit", auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.Commit(&_EthClosedBidFirstAuction.TransactOpts, auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactorSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.Commit(&_EthClosedBidFirstAuction.TransactOpts, auctionId, jsonString)
}

// Create is a paid mutator transaction binding the contract method 0xb6a46b3b.
//
// Solidity: function create(string _asset_id) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactor) Create(opts *bind.TransactOpts, _asset_id string) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.contract.Transact(opts, "create", _asset_id)
}

// Create is a paid mutator transaction binding the contract method 0xb6a46b3b.
//
// Solidity: function create(string _asset_id) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) Create(_asset_id string) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.Create(&_EthClosedBidFirstAuction.TransactOpts, _asset_id)
}

// Create is a paid mutator transaction binding the contract method 0xb6a46b3b.
//
// Solidity: function create(string _asset_id) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactorSession) Create(_asset_id string) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.Create(&_EthClosedBidFirstAuction.TransactOpts, _asset_id)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactor) ProvideFeedback(opts *bind.TransactOpts, auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.contract.Transact(opts, "provide_feedback", auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.ProvideFeedback(&_EthClosedBidFirstAuction.TransactOpts, auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactorSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.ProvideFeedback(&_EthClosedBidFirstAuction.TransactOpts, auctionId, _score, _feedback)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactor) Reveal(opts *bind.TransactOpts, auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.contract.Transact(opts, "reveal", auctionId, bidAmount)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) Reveal(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.Reveal(&_EthClosedBidFirstAuction.TransactOpts, auctionId, bidAmount)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactorSession) Reveal(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.Reveal(&_EthClosedBidFirstAuction.TransactOpts, auctionId, bidAmount)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactor) RevealAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.contract.Transact(opts, "revealAuction", auctionId)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) RevealAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.RevealAuction(&_EthClosedBidFirstAuction.TransactOpts, auctionId)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactorSession) RevealAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.RevealAuction(&_EthClosedBidFirstAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactor) Withdraw(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.contract.Transact(opts, "withdraw", auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.Withdraw(&_EthClosedBidFirstAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionTransactorSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidFirstAuction.Contract.Withdraw(&_EthClosedBidFirstAuction.TransactOpts, auctionId)
}

// EthClosedBidFirstAuctionAwaitResponseIterator is returned from FilterAwaitResponse and is used to iterate over the raw logs and unpacked data for AwaitResponse events raised by the EthClosedBidFirstAuction contract.
type EthClosedBidFirstAuctionAwaitResponseIterator struct {
	Event *EthClosedBidFirstAuctionAwaitResponse // Event containing the contract specifics and raw log

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
func (it *EthClosedBidFirstAuctionAwaitResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidFirstAuctionAwaitResponse)
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
		it.Event = new(EthClosedBidFirstAuctionAwaitResponse)
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
func (it *EthClosedBidFirstAuctionAwaitResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidFirstAuctionAwaitResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidFirstAuctionAwaitResponse represents a AwaitResponse event raised by the EthClosedBidFirstAuction contract.
type EthClosedBidFirstAuctionAwaitResponse struct {
	Auction *big.Int
	Winner  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAwaitResponse is a free log retrieval operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) FilterAwaitResponse(opts *bind.FilterOpts) (*EthClosedBidFirstAuctionAwaitResponseIterator, error) {

	logs, sub, err := _EthClosedBidFirstAuction.contract.FilterLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidFirstAuctionAwaitResponseIterator{contract: _EthClosedBidFirstAuction.contract, event: "AwaitResponse", logs: logs, sub: sub}, nil
}

// WatchAwaitResponse is a free log subscription operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) WatchAwaitResponse(opts *bind.WatchOpts, sink chan<- *EthClosedBidFirstAuctionAwaitResponse) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidFirstAuction.contract.WatchLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidFirstAuctionAwaitResponse)
				if err := _EthClosedBidFirstAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
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
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) ParseAwaitResponse(log types.Log) (*EthClosedBidFirstAuctionAwaitResponse, error) {
	event := new(EthClosedBidFirstAuctionAwaitResponse)
	if err := _EthClosedBidFirstAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidFirstAuctionDecisionMadeIterator is returned from FilterDecisionMade and is used to iterate over the raw logs and unpacked data for DecisionMade events raised by the EthClosedBidFirstAuction contract.
type EthClosedBidFirstAuctionDecisionMadeIterator struct {
	Event *EthClosedBidFirstAuctionDecisionMade // Event containing the contract specifics and raw log

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
func (it *EthClosedBidFirstAuctionDecisionMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidFirstAuctionDecisionMade)
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
		it.Event = new(EthClosedBidFirstAuctionDecisionMade)
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
func (it *EthClosedBidFirstAuctionDecisionMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidFirstAuctionDecisionMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidFirstAuctionDecisionMade represents a DecisionMade event raised by the EthClosedBidFirstAuction contract.
type EthClosedBidFirstAuctionDecisionMade struct {
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
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) FilterDecisionMade(opts *bind.FilterOpts) (*EthClosedBidFirstAuctionDecisionMadeIterator, error) {

	logs, sub, err := _EthClosedBidFirstAuction.contract.FilterLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidFirstAuctionDecisionMadeIterator{contract: _EthClosedBidFirstAuction.contract, event: "DecisionMade", logs: logs, sub: sub}, nil
}

// WatchDecisionMade is a free log subscription operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auction, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) WatchDecisionMade(opts *bind.WatchOpts, sink chan<- *EthClosedBidFirstAuctionDecisionMade) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidFirstAuction.contract.WatchLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidFirstAuctionDecisionMade)
				if err := _EthClosedBidFirstAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
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
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) ParseDecisionMade(log types.Log) (*EthClosedBidFirstAuctionDecisionMade, error) {
	event := new(EthClosedBidFirstAuctionDecisionMade)
	if err := _EthClosedBidFirstAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidFirstAuctionHighestBidIncreasedIterator is returned from FilterHighestBidIncreased and is used to iterate over the raw logs and unpacked data for HighestBidIncreased events raised by the EthClosedBidFirstAuction contract.
type EthClosedBidFirstAuctionHighestBidIncreasedIterator struct {
	Event *EthClosedBidFirstAuctionHighestBidIncreased // Event containing the contract specifics and raw log

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
func (it *EthClosedBidFirstAuctionHighestBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidFirstAuctionHighestBidIncreased)
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
		it.Event = new(EthClosedBidFirstAuctionHighestBidIncreased)
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
func (it *EthClosedBidFirstAuctionHighestBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidFirstAuctionHighestBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidFirstAuctionHighestBidIncreased represents a HighestBidIncreased event raised by the EthClosedBidFirstAuction contract.
type EthClosedBidFirstAuctionHighestBidIncreased struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*EthClosedBidFirstAuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _EthClosedBidFirstAuction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidFirstAuctionHighestBidIncreasedIterator{contract: _EthClosedBidFirstAuction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) WatchHighestBidIncreased(opts *bind.WatchOpts, sink chan<- *EthClosedBidFirstAuctionHighestBidIncreased) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidFirstAuction.contract.WatchLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidFirstAuctionHighestBidIncreased)
				if err := _EthClosedBidFirstAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
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
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) ParseHighestBidIncreased(log types.Log) (*EthClosedBidFirstAuctionHighestBidIncreased, error) {
	event := new(EthClosedBidFirstAuctionHighestBidIncreased)
	if err := _EthClosedBidFirstAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidFirstAuctionNewBidHashIterator is returned from FilterNewBidHash and is used to iterate over the raw logs and unpacked data for NewBidHash events raised by the EthClosedBidFirstAuction contract.
type EthClosedBidFirstAuctionNewBidHashIterator struct {
	Event *EthClosedBidFirstAuctionNewBidHash // Event containing the contract specifics and raw log

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
func (it *EthClosedBidFirstAuctionNewBidHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidFirstAuctionNewBidHash)
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
		it.Event = new(EthClosedBidFirstAuctionNewBidHash)
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
func (it *EthClosedBidFirstAuctionNewBidHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidFirstAuctionNewBidHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidFirstAuctionNewBidHash represents a NewBidHash event raised by the EthClosedBidFirstAuction contract.
type EthClosedBidFirstAuctionNewBidHash struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	BidHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBidHash is a free log retrieval operation binding the contract event 0x6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f.
//
// Solidity: event NewBidHash(uint256 auction, string id, address bidder, bytes32 bidHash)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) FilterNewBidHash(opts *bind.FilterOpts) (*EthClosedBidFirstAuctionNewBidHashIterator, error) {

	logs, sub, err := _EthClosedBidFirstAuction.contract.FilterLogs(opts, "NewBidHash")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidFirstAuctionNewBidHashIterator{contract: _EthClosedBidFirstAuction.contract, event: "NewBidHash", logs: logs, sub: sub}, nil
}

// WatchNewBidHash is a free log subscription operation binding the contract event 0x6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f.
//
// Solidity: event NewBidHash(uint256 auction, string id, address bidder, bytes32 bidHash)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) WatchNewBidHash(opts *bind.WatchOpts, sink chan<- *EthClosedBidFirstAuctionNewBidHash) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidFirstAuction.contract.WatchLogs(opts, "NewBidHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidFirstAuctionNewBidHash)
				if err := _EthClosedBidFirstAuction.contract.UnpackLog(event, "NewBidHash", log); err != nil {
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
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) ParseNewBidHash(log types.Log) (*EthClosedBidFirstAuctionNewBidHash, error) {
	event := new(EthClosedBidFirstAuctionNewBidHash)
	if err := _EthClosedBidFirstAuction.contract.UnpackLog(event, "NewBidHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidFirstAuctionRateAuctionIterator is returned from FilterRateAuction and is used to iterate over the raw logs and unpacked data for RateAuction events raised by the EthClosedBidFirstAuction contract.
type EthClosedBidFirstAuctionRateAuctionIterator struct {
	Event *EthClosedBidFirstAuctionRateAuction // Event containing the contract specifics and raw log

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
func (it *EthClosedBidFirstAuctionRateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidFirstAuctionRateAuction)
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
		it.Event = new(EthClosedBidFirstAuctionRateAuction)
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
func (it *EthClosedBidFirstAuctionRateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidFirstAuctionRateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidFirstAuctionRateAuction represents a RateAuction event raised by the EthClosedBidFirstAuction contract.
type EthClosedBidFirstAuctionRateAuction struct {
	Auction *big.Int
	Id      string
	Rating  *big.Int
	Review  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRateAuction is a free log retrieval operation binding the contract event 0x10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e79.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, bytes32 review)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) FilterRateAuction(opts *bind.FilterOpts) (*EthClosedBidFirstAuctionRateAuctionIterator, error) {

	logs, sub, err := _EthClosedBidFirstAuction.contract.FilterLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidFirstAuctionRateAuctionIterator{contract: _EthClosedBidFirstAuction.contract, event: "RateAuction", logs: logs, sub: sub}, nil
}

// WatchRateAuction is a free log subscription operation binding the contract event 0x10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e79.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, bytes32 review)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) WatchRateAuction(opts *bind.WatchOpts, sink chan<- *EthClosedBidFirstAuctionRateAuction) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidFirstAuction.contract.WatchLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidFirstAuctionRateAuction)
				if err := _EthClosedBidFirstAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
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
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) ParseRateAuction(log types.Log) (*EthClosedBidFirstAuctionRateAuction, error) {
	event := new(EthClosedBidFirstAuctionRateAuction)
	if err := _EthClosedBidFirstAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidFirstAuctionRevealAuctionIterator is returned from FilterRevealAuction and is used to iterate over the raw logs and unpacked data for RevealAuction events raised by the EthClosedBidFirstAuction contract.
type EthClosedBidFirstAuctionRevealAuctionIterator struct {
	Event *EthClosedBidFirstAuctionRevealAuction // Event containing the contract specifics and raw log

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
func (it *EthClosedBidFirstAuctionRevealAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidFirstAuctionRevealAuction)
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
		it.Event = new(EthClosedBidFirstAuctionRevealAuction)
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
func (it *EthClosedBidFirstAuctionRevealAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidFirstAuctionRevealAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidFirstAuctionRevealAuction represents a RevealAuction event raised by the EthClosedBidFirstAuction contract.
type EthClosedBidFirstAuctionRevealAuction struct {
	Auction *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevealAuction is a free log retrieval operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auction)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) FilterRevealAuction(opts *bind.FilterOpts) (*EthClosedBidFirstAuctionRevealAuctionIterator, error) {

	logs, sub, err := _EthClosedBidFirstAuction.contract.FilterLogs(opts, "RevealAuction")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidFirstAuctionRevealAuctionIterator{contract: _EthClosedBidFirstAuction.contract, event: "RevealAuction", logs: logs, sub: sub}, nil
}

// WatchRevealAuction is a free log subscription operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auction)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) WatchRevealAuction(opts *bind.WatchOpts, sink chan<- *EthClosedBidFirstAuctionRevealAuction) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidFirstAuction.contract.WatchLogs(opts, "RevealAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidFirstAuctionRevealAuction)
				if err := _EthClosedBidFirstAuction.contract.UnpackLog(event, "RevealAuction", log); err != nil {
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
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) ParseRevealAuction(log types.Log) (*EthClosedBidFirstAuctionRevealAuction, error) {
	event := new(EthClosedBidFirstAuctionRevealAuction)
	if err := _EthClosedBidFirstAuction.contract.UnpackLog(event, "RevealAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidFirstAuctionWithdrawBidIterator is returned from FilterWithdrawBid and is used to iterate over the raw logs and unpacked data for WithdrawBid events raised by the EthClosedBidFirstAuction contract.
type EthClosedBidFirstAuctionWithdrawBidIterator struct {
	Event *EthClosedBidFirstAuctionWithdrawBid // Event containing the contract specifics and raw log

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
func (it *EthClosedBidFirstAuctionWithdrawBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidFirstAuctionWithdrawBid)
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
		it.Event = new(EthClosedBidFirstAuctionWithdrawBid)
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
func (it *EthClosedBidFirstAuctionWithdrawBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidFirstAuctionWithdrawBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidFirstAuctionWithdrawBid represents a WithdrawBid event raised by the EthClosedBidFirstAuction contract.
type EthClosedBidFirstAuctionWithdrawBid struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBid is a free log retrieval operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) FilterWithdrawBid(opts *bind.FilterOpts) (*EthClosedBidFirstAuctionWithdrawBidIterator, error) {

	logs, sub, err := _EthClosedBidFirstAuction.contract.FilterLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidFirstAuctionWithdrawBidIterator{contract: _EthClosedBidFirstAuction.contract, event: "WithdrawBid", logs: logs, sub: sub}, nil
}

// WatchWithdrawBid is a free log subscription operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) WatchWithdrawBid(opts *bind.WatchOpts, sink chan<- *EthClosedBidFirstAuctionWithdrawBid) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidFirstAuction.contract.WatchLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidFirstAuctionWithdrawBid)
				if err := _EthClosedBidFirstAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
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
func (_EthClosedBidFirstAuction *EthClosedBidFirstAuctionFilterer) ParseWithdrawBid(log types.Log) (*EthClosedBidFirstAuctionWithdrawBid, error) {
	event := new(EthClosedBidFirstAuctionWithdrawBid)
	if err := _EthClosedBidFirstAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
