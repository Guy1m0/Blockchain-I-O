// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dutch_auction

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

// DutchAuctionMetaData contains all meta data concerning the DutchAuction contract.
var DutchAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AwaitResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BidReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prcd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"DecisionMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"rating\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"review\",\"type\":\"bytes32\"}],\"name\":\"RateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkAverageScore\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"not_winner_platform\",\"type\":\"bool\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_asset_id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_discountRate\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"discountRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_score\",\"type\":\"int256\"},{\"internalType\":\"bytes32\",\"name\":\"_feedback\",\"type\":\"bytes32\"}],\"name\":\"provide_feedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"startAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"startingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"winningBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"winningBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051611da6380380611da683398101604081905261002f91610052565b6001600160a01b0316608052600080546001600160a01b03191633179055610082565b60006020828403121561006457600080fd5b81516001600160a01b038116811461007b57600080fd5b9392505050565b608051611cfb6100ab600039600081816102a2015281816108be0152610c040152611cfb6000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80638da5cb5b116100a2578063d4ac9b8c11610071578063d4ac9b8c1461025c578063dad239361461026f578063e757223014610277578063f691877e1461028a578063fc0c546a1461029d57600080fd5b80638da5cb5b14610210578063c84d2f6a14610223578063cda3c49914610236578063d0a1414a1461024957600080fd5b806342d21ef7116100e957806342d21ef71461019657806351e6ca40146101b6578063598647f8146101d75780637552c9d2146101ea5780638d10b0a6146101fd57600080fd5b8063176321e91461011b578063274377c0146101305780632dd7f648146101435780632e1a7d4d14610173575b600080fd5b61012e6101293660046116ce565b6102c4565b005b61012e61013e366004611715565b610592565b610156610151366004611741565b61077f565b6040516001600160a01b0390911681526020015b60405180910390f35b610186610181366004611741565b6107a9565b604051901515815260200161016a565b6101a96101a4366004611741565b6109aa565b60405161016a91906117a0565b6101c96101c4366004611741565b610a56565b60405190815260200161016a565b61012e6101e53660046117ba565b610a77565b6101c96101f8366004611741565b610d84565b6101a961020b366004611741565b610d94565b600054610156906001600160a01b031681565b61012e6102313660046117ed565b610da4565b61012e61024436600461181d565b611010565b61012e6102573660046116ce565b611268565b6101c961026a366004611741565b6114a0565b6101c96114b0565b6101c9610285366004611741565b61158f565b6101c9610298366004611741565b61161b565b6101567f000000000000000000000000000000000000000000000000000000000000000081565b60405165656e64696e6760d01b602082015260260160405160208183030381529060405280519060200120600783815481106103025761030261186b565b9060005260206000200160405160200161031c91906118bb565b60405160208183030381529060405280519060200120146103845760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064015b60405180910390fd5b600582815481106103975761039761186b565b6000918252602090912001546001600160a01b031633146103ca5760405162461bcd60e51b815260040161037b90611931565b60405180604001604052806007815260200166636c6f73696e6760c81b815250600783815481106103fd576103fd61186b565b90600052602060002001908161041391906119b0565b507f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c782600584815481106104495761044961186b565b600091825260209091200154600680546001600160a01b0390921691869081106104755761047561186b565b9060005260206000200154600886815481106104935761049361186b565b906000526020600020016000866040516104b296959493929190611aed565b60405180910390a1600682815481106104cd576104cd61186b565b906000526020600020015460046000600585815481106104ef576104ef61186b565b60009182526020808320909101546001600160a01b0316835282019290925260400181208054909190610523908490611b57565b9250508190555060006005838154811061053f5761053f61186b565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506000600683815481106105825761058261186b565b6000918252602090912001555050565b60405166636c6f73696e6760c81b602082015260270160405160208183030381529060405280519060200120600784815481106105d1576105d161186b565b906000526020600020016040516020016105eb91906118bb565b604051602081830303815290604052805190602001201461064e5760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e20434c4f53494e47207374617475730000604482015260640161037b565b600583815481106106615761066161186b565b6000918252602090912001546001600160a01b031633146106945760405162461bcd60e51b815260040161037b90611931565b81600a84815481106106a8576106a861186b565b906000526020600020018190555080600984815481106106ca576106ca61186b565b90600052602060002001819055507f10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e79836008858154811061070d5761070d61186b565b9060005260206000200184846040516107299493929190611b70565b60405180910390a16040518060400160405280600681526020016518db1bdcd95960d21b815250600784815481106107635761076361186b565b90600052602060002001908161077991906119b0565b50505050565b6005818154811061078f57600080fd5b6000918252602090912001546001600160a01b0316905081565b6040516337b832b760e11b602082015260009060240160405160208183030381529060405280519060200120600783815481106107e8576107e861186b565b9060005260206000200160405160200161080291906118bb565b60405160208183030381529060405280519060200120036108655760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420696e204f50454e20737461747573000000000000000000604482015260640161037b565b336000908152600460205260409020548015610948573360008181526004602081905260408083209290925590516323b872dd60e01b815230918101919091526024810191909152604481018290526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303816000875af1158015610907573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061092b9190611b9c565b610948573360009081526004602052604081209190915592915050565b7f9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c836008858154811061097d5761097d61186b565b9060005260206000200133846040516109999493929190611bb9565b60405180910390a150600192915050565b600781815481106109ba57600080fd5b9060005260206000200160009150905080546109d590611881565b80601f0160208091040260200160405190810160405280929190818152602001828054610a0190611881565b8015610a4e5780601f10610a2357610100808354040283529160200191610a4e565b820191906000526020600020905b815481529060010190602001808311610a3157829003601f168201915b505050505081565b60018181548110610a6657600080fd5b600091825260209091200154905081565b6040516337b832b760e11b60208201526024016040516020818303038152906040528051906020012060078381548110610ab357610ab361186b565b90600052602060002001604051602001610acd91906118bb565b6040516020818303038152906040528051906020012014610b305760405162461bcd60e51b815260206004820152601b60248201527f436f6e7472616374206e6f7420696e204f50454e207374617475730000000000604482015260640161037b565b60068281548110610b4357610b4361186b565b9060005260206000200154600014610b915760405162461bcd60e51b8152602060048201526011602482015270139bdd081d1a1948199a5c9cdd08189a59607a1b604482015260640161037b565b6000610b9c8361158f565b905080821015610bdf5760405162461bcd60e51b815260206004820152600e60248201526d42696420697320746f6f206c6f7760901b604482015260640161037b565b6040516323b872dd60e01b8152336004820152306024820152604481018390526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610c55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c799190611b9c565b905080610cc15760405162461bcd60e51b81526020600482015260166024820152752a37b5b2b7103a3930b739b332b9103330b4b632b21760511b604482015260640161037b565b3360058581548110610cd557610cd561186b565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508260068581548110610d1757610d1761186b565b90600052602060002001819055507f491989e56ba54130286a8a051732d8f228353c97d40b74d4584352242b3dba998460088681548110610d5a57610d5a61186b565b906000526020600020013386604051610d769493929190611bb9565b60405180910390a150505050565b60038181548110610a6657600080fd5b600881815481106109ba57600080fd5b6000546001600160a01b03163314610e0e5760405162461bcd60e51b815260206004820152602760248201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736044820152662073746174757360c81b606482015260840161037b565b6040516337b832b760e11b60208201526024016040516020818303038152906040528051906020012060078381548110610e4a57610e4a61186b565b90600052602060002001604051602001610e6491906118bb565b6040516020818303038152906040528051906020012014610ec75760405162461bcd60e51b815260206004820152601b60248201527f436f6e7472616374206e6f7420696e204f50454e207374617475730000000000604482015260640161037b565b60405180604001604052806006815260200165656e64696e6760d01b81525060078381548110610ef957610ef961186b565b906000526020600020019081610f0f91906119b0565b508080610f39575060068281548110610f2a57610f2a61186b565b90600052602060002001546000145b15610f9a576040518060400160405280600681526020016518db1bdcd95960d21b81525060078381548110610f7057610f7061186b565b906000526020600020019081610f8691906119b0565b50600682815481106104cd576104cd61186b565b7fa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b48474685002738260058481548110610fcf57610fcf61186b565b60009182526020909120015460405161100492916001600160a01b0316909182526001600160a01b0316602082015260400190565b60405180910390a15050565b6000546001600160a01b031633146110745760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79206f776e65722063616e20637265617465206e65772061756374696f6044820152603760f91b606482015260840161037b565b6005805460018181019092557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180546001600160a01b03191690556006805480830190915560007ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f909101819055600780549283018155905260408051808201909152600481526337b832b760e11b60208201527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6889091019061113890826119b0565b50600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30161117584826119b0565b506001805480820182557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6019290925560028054808401909155427f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace90910155600380548084019091557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01556009805480830190915560007f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af909101819055600a8054928301815581527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a89091015550565b60405165656e64696e6760d01b602082015260260160405160208183030381529060405280519060200120600783815481106112a6576112a661186b565b906000526020600020016040516020016112c091906118bb565b60405160208183030381529060405280519060200120146113235760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e4720737461747573000000604482015260640161037b565b600582815481106113365761133661186b565b6000918252602090912001546001600160a01b031633146113695760405162461bcd60e51b815260040161037b90611931565b60405180604001604052806007815260200166636c6f73696e6760c81b8152506007838154811061139c5761139c61186b565b9060005260206000200190816113b291906119b0565b50600682815481106113c6576113c661186b565b600091825260208083209091015482546001600160a01b031683526004909152604082208054919290916113fb908490611b57565b925050819055507f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c782600584815481106114375761143761186b565b600091825260209091200154600680546001600160a01b0390921691869081106114635761146361186b565b9060005260206000200154600886815481106114815761148161186b565b9060005260206000200160018660405161100496959493929190611aed565b60068181548110610a6657600080fd5b600080805b600a54811015611570576040516518db1bdcd95960d21b602082015260260160405160208183030381529060405280519060200120600782815481106114fd576114fd61186b565b9060005260206000200160405160200161151791906118bb565b604051602081830303815290604052805190602001200361155e57600a81815481106115455761154561186b565b90600052602060002001548261155b9190611bee565b91505b8061156881611c16565b9150506114b5565b50600a5461157f826064611c2f565b6115899190611c5f565b91505090565b600080600283815481106115a5576115a561186b565b9060005260206000200154426115bb9190611c9b565b9050600081600385815481106115d3576115d361186b565b90600052602060002001546115e89190611cae565b905080600185815481106115fe576115fe61186b565b90600052602060002001546116139190611c9b565b949350505050565b60028181548110610a6657600080fd5b634e487b7160e01b600052604160045260246000fd5b600082601f83011261165257600080fd5b813567ffffffffffffffff8082111561166d5761166d61162b565b604051601f8301601f19908116603f011681019082821181831017156116955761169561162b565b816040528381528660208588010111156116ae57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080604083850312156116e157600080fd5b82359150602083013567ffffffffffffffff8111156116ff57600080fd5b61170b85828601611641565b9150509250929050565b60008060006060848603121561172a57600080fd5b505081359360208301359350604090920135919050565b60006020828403121561175357600080fd5b5035919050565b6000815180845260005b8181101561178057602081850181015186830182015201611764565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006117b3602083018461175a565b9392505050565b600080604083850312156117cd57600080fd5b50508035926020909101359150565b80151581146117ea57600080fd5b50565b6000806040838503121561180057600080fd5b823591506020830135611812816117dc565b809150509250929050565b60008060006060848603121561183257600080fd5b833567ffffffffffffffff81111561184957600080fd5b61185586828701611641565b9660208601359650604090950135949350505050565b634e487b7160e01b600052603260045260246000fd5b600181811c9082168061189557607f821691505b6020821081036118b557634e487b7160e01b600052602260045260246000fd5b50919050565b60008083546118c981611881565b600182811680156118e157600181146118f657611925565b60ff1984168752821515830287019450611925565b8760005260208060002060005b8581101561191c5781548a820152908401908201611903565b50505082870194505b50929695505050505050565b6020808252601690820152754e6f7420617574686f72697a6564206163636573732160501b604082015260600190565b601f8211156119ab57600081815260208120601f850160051c810160208610156119885750805b601f850160051c820191505b818110156119a757828155600101611994565b5050505b505050565b815167ffffffffffffffff8111156119ca576119ca61162b565b6119de816119d88454611881565b84611961565b602080601f831160018114611a1357600084156119fb5750858301515b600019600386901b1c1916600185901b1785556119a7565b600085815260208120601f198616915b82811015611a4257888601518255948401946001909101908401611a23565b5085821015611a605787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008154611a7d81611881565b808552602060018381168015611a9a5760018114611ab457611ae2565b60ff1985168884015283151560051b880183019550611ae2565b866000528260002060005b85811015611ada5781548a8201860152908301908401611abf565b890184019650505b505050505092915050565b86815260018060a01b038616602082015284604082015260c060608201526000611b1a60c0830186611a70565b841515608084015282810360a0840152611b34818561175a565b9998505050505050505050565b634e487b7160e01b600052601160045260246000fd5b80820180821115611b6a57611b6a611b41565b92915050565b848152608060208201526000611b896080830186611a70565b6040830194909452506060015292915050565b600060208284031215611bae57600080fd5b81516117b3816117dc565b848152608060208201526000611bd26080830186611a70565b6001600160a01b03949094166040830152506060015292915050565b8082018281126000831280158216821582161715611c0e57611c0e611b41565b505092915050565b600060018201611c2857611c28611b41565b5060010190565b80820260008212600160ff1b84141615611c4b57611c4b611b41565b8181058314821517611b6a57611b6a611b41565b600082611c7c57634e487b7160e01b600052601260045260246000fd5b600160ff1b821460001984141615611c9657611c96611b41565b500590565b81810381811115611b6a57611b6a611b41565b8082028115828204841417611b6a57611b6a611b4156fea26469706673582212205f46cc8b96d069fedbeb35a689c622597b953b12b15b168c75c2717bdaa9424764736f6c63430008120033",
}

// DutchAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use DutchAuctionMetaData.ABI instead.
var DutchAuctionABI = DutchAuctionMetaData.ABI

// DutchAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DutchAuctionMetaData.Bin instead.
var DutchAuctionBin = DutchAuctionMetaData.Bin

// DeployDutchAuction deploys a new Ethereum contract, binding an instance of DutchAuction to it.
func DeployDutchAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *DutchAuction, error) {
	parsed, err := DutchAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DutchAuctionBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DutchAuction{DutchAuctionCaller: DutchAuctionCaller{contract: contract}, DutchAuctionTransactor: DutchAuctionTransactor{contract: contract}, DutchAuctionFilterer: DutchAuctionFilterer{contract: contract}}, nil
}

// DutchAuction is an auto generated Go binding around an Ethereum contract.
type DutchAuction struct {
	DutchAuctionCaller     // Read-only binding to the contract
	DutchAuctionTransactor // Write-only binding to the contract
	DutchAuctionFilterer   // Log filterer for contract events
}

// DutchAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type DutchAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DutchAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DutchAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DutchAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DutchAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DutchAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DutchAuctionSession struct {
	Contract     *DutchAuction     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DutchAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DutchAuctionCallerSession struct {
	Contract *DutchAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DutchAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DutchAuctionTransactorSession struct {
	Contract     *DutchAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DutchAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type DutchAuctionRaw struct {
	Contract *DutchAuction // Generic contract binding to access the raw methods on
}

// DutchAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DutchAuctionCallerRaw struct {
	Contract *DutchAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// DutchAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DutchAuctionTransactorRaw struct {
	Contract *DutchAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDutchAuction creates a new instance of DutchAuction, bound to a specific deployed contract.
func NewDutchAuction(address common.Address, backend bind.ContractBackend) (*DutchAuction, error) {
	contract, err := bindDutchAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DutchAuction{DutchAuctionCaller: DutchAuctionCaller{contract: contract}, DutchAuctionTransactor: DutchAuctionTransactor{contract: contract}, DutchAuctionFilterer: DutchAuctionFilterer{contract: contract}}, nil
}

// NewDutchAuctionCaller creates a new read-only instance of DutchAuction, bound to a specific deployed contract.
func NewDutchAuctionCaller(address common.Address, caller bind.ContractCaller) (*DutchAuctionCaller, error) {
	contract, err := bindDutchAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DutchAuctionCaller{contract: contract}, nil
}

// NewDutchAuctionTransactor creates a new write-only instance of DutchAuction, bound to a specific deployed contract.
func NewDutchAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*DutchAuctionTransactor, error) {
	contract, err := bindDutchAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DutchAuctionTransactor{contract: contract}, nil
}

// NewDutchAuctionFilterer creates a new log filterer instance of DutchAuction, bound to a specific deployed contract.
func NewDutchAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*DutchAuctionFilterer, error) {
	contract, err := bindDutchAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DutchAuctionFilterer{contract: contract}, nil
}

// bindDutchAuction binds a generic wrapper to an already deployed contract.
func bindDutchAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DutchAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DutchAuction *DutchAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DutchAuction.Contract.DutchAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DutchAuction *DutchAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DutchAuction.Contract.DutchAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DutchAuction *DutchAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DutchAuction.Contract.DutchAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DutchAuction *DutchAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DutchAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DutchAuction *DutchAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DutchAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DutchAuction *DutchAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DutchAuction.Contract.contract.Transact(opts, method, params...)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_DutchAuction *DutchAuctionCaller) AssetId(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "asset_id", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_DutchAuction *DutchAuctionSession) AssetId(arg0 *big.Int) (string, error) {
	return _DutchAuction.Contract.AssetId(&_DutchAuction.CallOpts, arg0)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_DutchAuction *DutchAuctionCallerSession) AssetId(arg0 *big.Int) (string, error) {
	return _DutchAuction.Contract.AssetId(&_DutchAuction.CallOpts, arg0)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_DutchAuction *DutchAuctionCaller) CheckAverageScore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "checkAverageScore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_DutchAuction *DutchAuctionSession) CheckAverageScore() (*big.Int, error) {
	return _DutchAuction.Contract.CheckAverageScore(&_DutchAuction.CallOpts)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_DutchAuction *DutchAuctionCallerSession) CheckAverageScore() (*big.Int, error) {
	return _DutchAuction.Contract.CheckAverageScore(&_DutchAuction.CallOpts)
}

// DiscountRate is a free data retrieval call binding the contract method 0x7552c9d2.
//
// Solidity: function discountRate(uint256 ) view returns(uint256)
func (_DutchAuction *DutchAuctionCaller) DiscountRate(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "discountRate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DiscountRate is a free data retrieval call binding the contract method 0x7552c9d2.
//
// Solidity: function discountRate(uint256 ) view returns(uint256)
func (_DutchAuction *DutchAuctionSession) DiscountRate(arg0 *big.Int) (*big.Int, error) {
	return _DutchAuction.Contract.DiscountRate(&_DutchAuction.CallOpts, arg0)
}

// DiscountRate is a free data retrieval call binding the contract method 0x7552c9d2.
//
// Solidity: function discountRate(uint256 ) view returns(uint256)
func (_DutchAuction *DutchAuctionCallerSession) DiscountRate(arg0 *big.Int) (*big.Int, error) {
	return _DutchAuction.Contract.DiscountRate(&_DutchAuction.CallOpts, arg0)
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 auctionId) view returns(uint256)
func (_DutchAuction *DutchAuctionCaller) GetPrice(opts *bind.CallOpts, auctionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "getPrice", auctionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 auctionId) view returns(uint256)
func (_DutchAuction *DutchAuctionSession) GetPrice(auctionId *big.Int) (*big.Int, error) {
	return _DutchAuction.Contract.GetPrice(&_DutchAuction.CallOpts, auctionId)
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 auctionId) view returns(uint256)
func (_DutchAuction *DutchAuctionCallerSession) GetPrice(auctionId *big.Int) (*big.Int, error) {
	return _DutchAuction.Contract.GetPrice(&_DutchAuction.CallOpts, auctionId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DutchAuction *DutchAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DutchAuction *DutchAuctionSession) Owner() (common.Address, error) {
	return _DutchAuction.Contract.Owner(&_DutchAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DutchAuction *DutchAuctionCallerSession) Owner() (common.Address, error) {
	return _DutchAuction.Contract.Owner(&_DutchAuction.CallOpts)
}

// StartAt is a free data retrieval call binding the contract method 0xf691877e.
//
// Solidity: function startAt(uint256 ) view returns(uint256)
func (_DutchAuction *DutchAuctionCaller) StartAt(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "startAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartAt is a free data retrieval call binding the contract method 0xf691877e.
//
// Solidity: function startAt(uint256 ) view returns(uint256)
func (_DutchAuction *DutchAuctionSession) StartAt(arg0 *big.Int) (*big.Int, error) {
	return _DutchAuction.Contract.StartAt(&_DutchAuction.CallOpts, arg0)
}

// StartAt is a free data retrieval call binding the contract method 0xf691877e.
//
// Solidity: function startAt(uint256 ) view returns(uint256)
func (_DutchAuction *DutchAuctionCallerSession) StartAt(arg0 *big.Int) (*big.Int, error) {
	return _DutchAuction.Contract.StartAt(&_DutchAuction.CallOpts, arg0)
}

// StartingPrice is a free data retrieval call binding the contract method 0x51e6ca40.
//
// Solidity: function startingPrice(uint256 ) view returns(uint256)
func (_DutchAuction *DutchAuctionCaller) StartingPrice(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "startingPrice", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingPrice is a free data retrieval call binding the contract method 0x51e6ca40.
//
// Solidity: function startingPrice(uint256 ) view returns(uint256)
func (_DutchAuction *DutchAuctionSession) StartingPrice(arg0 *big.Int) (*big.Int, error) {
	return _DutchAuction.Contract.StartingPrice(&_DutchAuction.CallOpts, arg0)
}

// StartingPrice is a free data retrieval call binding the contract method 0x51e6ca40.
//
// Solidity: function startingPrice(uint256 ) view returns(uint256)
func (_DutchAuction *DutchAuctionCallerSession) StartingPrice(arg0 *big.Int) (*big.Int, error) {
	return _DutchAuction.Contract.StartingPrice(&_DutchAuction.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_DutchAuction *DutchAuctionCaller) Status(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "status", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_DutchAuction *DutchAuctionSession) Status(arg0 *big.Int) (string, error) {
	return _DutchAuction.Contract.Status(&_DutchAuction.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_DutchAuction *DutchAuctionCallerSession) Status(arg0 *big.Int) (string, error) {
	return _DutchAuction.Contract.Status(&_DutchAuction.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_DutchAuction *DutchAuctionCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_DutchAuction *DutchAuctionSession) Token() (common.Address, error) {
	return _DutchAuction.Contract.Token(&_DutchAuction.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_DutchAuction *DutchAuctionCallerSession) Token() (common.Address, error) {
	return _DutchAuction.Contract.Token(&_DutchAuction.CallOpts)
}

// WinningBid is a free data retrieval call binding the contract method 0xd4ac9b8c.
//
// Solidity: function winningBid(uint256 ) view returns(uint256)
func (_DutchAuction *DutchAuctionCaller) WinningBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "winningBid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WinningBid is a free data retrieval call binding the contract method 0xd4ac9b8c.
//
// Solidity: function winningBid(uint256 ) view returns(uint256)
func (_DutchAuction *DutchAuctionSession) WinningBid(arg0 *big.Int) (*big.Int, error) {
	return _DutchAuction.Contract.WinningBid(&_DutchAuction.CallOpts, arg0)
}

// WinningBid is a free data retrieval call binding the contract method 0xd4ac9b8c.
//
// Solidity: function winningBid(uint256 ) view returns(uint256)
func (_DutchAuction *DutchAuctionCallerSession) WinningBid(arg0 *big.Int) (*big.Int, error) {
	return _DutchAuction.Contract.WinningBid(&_DutchAuction.CallOpts, arg0)
}

// WinningBidder is a free data retrieval call binding the contract method 0x2dd7f648.
//
// Solidity: function winningBidder(uint256 ) view returns(address)
func (_DutchAuction *DutchAuctionCaller) WinningBidder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DutchAuction.contract.Call(opts, &out, "winningBidder", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WinningBidder is a free data retrieval call binding the contract method 0x2dd7f648.
//
// Solidity: function winningBidder(uint256 ) view returns(address)
func (_DutchAuction *DutchAuctionSession) WinningBidder(arg0 *big.Int) (common.Address, error) {
	return _DutchAuction.Contract.WinningBidder(&_DutchAuction.CallOpts, arg0)
}

// WinningBidder is a free data retrieval call binding the contract method 0x2dd7f648.
//
// Solidity: function winningBidder(uint256 ) view returns(address)
func (_DutchAuction *DutchAuctionCallerSession) WinningBidder(arg0 *big.Int) (common.Address, error) {
	return _DutchAuction.Contract.WinningBidder(&_DutchAuction.CallOpts, arg0)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_DutchAuction *DutchAuctionTransactor) Abort(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _DutchAuction.contract.Transact(opts, "abort", auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_DutchAuction *DutchAuctionSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _DutchAuction.Contract.Abort(&_DutchAuction.TransactOpts, auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_DutchAuction *DutchAuctionTransactorSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _DutchAuction.Contract.Abort(&_DutchAuction.TransactOpts, auctionId, jsonString)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 auctionId, uint256 bidAmount) returns()
func (_DutchAuction *DutchAuctionTransactor) Bid(opts *bind.TransactOpts, auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _DutchAuction.contract.Transact(opts, "bid", auctionId, bidAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 auctionId, uint256 bidAmount) returns()
func (_DutchAuction *DutchAuctionSession) Bid(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _DutchAuction.Contract.Bid(&_DutchAuction.TransactOpts, auctionId, bidAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 auctionId, uint256 bidAmount) returns()
func (_DutchAuction *DutchAuctionTransactorSession) Bid(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _DutchAuction.Contract.Bid(&_DutchAuction.TransactOpts, auctionId, bidAmount)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_DutchAuction *DutchAuctionTransactor) CloseAuction(opts *bind.TransactOpts, auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _DutchAuction.contract.Transact(opts, "closeAuction", auctionId, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_DutchAuction *DutchAuctionSession) CloseAuction(auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _DutchAuction.Contract.CloseAuction(&_DutchAuction.TransactOpts, auctionId, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_DutchAuction *DutchAuctionTransactorSession) CloseAuction(auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _DutchAuction.Contract.CloseAuction(&_DutchAuction.TransactOpts, auctionId, not_winner_platform)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_DutchAuction *DutchAuctionTransactor) Commit(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _DutchAuction.contract.Transact(opts, "commit", auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_DutchAuction *DutchAuctionSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _DutchAuction.Contract.Commit(&_DutchAuction.TransactOpts, auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_DutchAuction *DutchAuctionTransactorSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _DutchAuction.Contract.Commit(&_DutchAuction.TransactOpts, auctionId, jsonString)
}

// Create is a paid mutator transaction binding the contract method 0xcda3c499.
//
// Solidity: function create(string _asset_id, uint256 _startingPrice, uint256 _discountRate) returns()
func (_DutchAuction *DutchAuctionTransactor) Create(opts *bind.TransactOpts, _asset_id string, _startingPrice *big.Int, _discountRate *big.Int) (*types.Transaction, error) {
	return _DutchAuction.contract.Transact(opts, "create", _asset_id, _startingPrice, _discountRate)
}

// Create is a paid mutator transaction binding the contract method 0xcda3c499.
//
// Solidity: function create(string _asset_id, uint256 _startingPrice, uint256 _discountRate) returns()
func (_DutchAuction *DutchAuctionSession) Create(_asset_id string, _startingPrice *big.Int, _discountRate *big.Int) (*types.Transaction, error) {
	return _DutchAuction.Contract.Create(&_DutchAuction.TransactOpts, _asset_id, _startingPrice, _discountRate)
}

// Create is a paid mutator transaction binding the contract method 0xcda3c499.
//
// Solidity: function create(string _asset_id, uint256 _startingPrice, uint256 _discountRate) returns()
func (_DutchAuction *DutchAuctionTransactorSession) Create(_asset_id string, _startingPrice *big.Int, _discountRate *big.Int) (*types.Transaction, error) {
	return _DutchAuction.Contract.Create(&_DutchAuction.TransactOpts, _asset_id, _startingPrice, _discountRate)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_DutchAuction *DutchAuctionTransactor) ProvideFeedback(opts *bind.TransactOpts, auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _DutchAuction.contract.Transact(opts, "provide_feedback", auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_DutchAuction *DutchAuctionSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _DutchAuction.Contract.ProvideFeedback(&_DutchAuction.TransactOpts, auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_DutchAuction *DutchAuctionTransactorSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _DutchAuction.Contract.ProvideFeedback(&_DutchAuction.TransactOpts, auctionId, _score, _feedback)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_DutchAuction *DutchAuctionTransactor) Withdraw(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _DutchAuction.contract.Transact(opts, "withdraw", auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_DutchAuction *DutchAuctionSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _DutchAuction.Contract.Withdraw(&_DutchAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_DutchAuction *DutchAuctionTransactorSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _DutchAuction.Contract.Withdraw(&_DutchAuction.TransactOpts, auctionId)
}

// DutchAuctionAwaitResponseIterator is returned from FilterAwaitResponse and is used to iterate over the raw logs and unpacked data for AwaitResponse events raised by the DutchAuction contract.
type DutchAuctionAwaitResponseIterator struct {
	Event *DutchAuctionAwaitResponse // Event containing the contract specifics and raw log

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
func (it *DutchAuctionAwaitResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchAuctionAwaitResponse)
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
		it.Event = new(DutchAuctionAwaitResponse)
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
func (it *DutchAuctionAwaitResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchAuctionAwaitResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchAuctionAwaitResponse represents a AwaitResponse event raised by the DutchAuction contract.
type DutchAuctionAwaitResponse struct {
	Auction *big.Int
	Winner  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAwaitResponse is a free log retrieval operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_DutchAuction *DutchAuctionFilterer) FilterAwaitResponse(opts *bind.FilterOpts) (*DutchAuctionAwaitResponseIterator, error) {

	logs, sub, err := _DutchAuction.contract.FilterLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return &DutchAuctionAwaitResponseIterator{contract: _DutchAuction.contract, event: "AwaitResponse", logs: logs, sub: sub}, nil
}

// WatchAwaitResponse is a free log subscription operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_DutchAuction *DutchAuctionFilterer) WatchAwaitResponse(opts *bind.WatchOpts, sink chan<- *DutchAuctionAwaitResponse) (event.Subscription, error) {

	logs, sub, err := _DutchAuction.contract.WatchLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchAuctionAwaitResponse)
				if err := _DutchAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
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
func (_DutchAuction *DutchAuctionFilterer) ParseAwaitResponse(log types.Log) (*DutchAuctionAwaitResponse, error) {
	event := new(DutchAuctionAwaitResponse)
	if err := _DutchAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DutchAuctionBidReceivedIterator is returned from FilterBidReceived and is used to iterate over the raw logs and unpacked data for BidReceived events raised by the DutchAuction contract.
type DutchAuctionBidReceivedIterator struct {
	Event *DutchAuctionBidReceived // Event containing the contract specifics and raw log

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
func (it *DutchAuctionBidReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchAuctionBidReceived)
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
		it.Event = new(DutchAuctionBidReceived)
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
func (it *DutchAuctionBidReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchAuctionBidReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchAuctionBidReceived represents a BidReceived event raised by the DutchAuction contract.
type DutchAuctionBidReceived struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBidReceived is a free log retrieval operation binding the contract event 0x491989e56ba54130286a8a051732d8f228353c97d40b74d4584352242b3dba99.
//
// Solidity: event BidReceived(uint256 auction, string id, address bidder, uint256 amount)
func (_DutchAuction *DutchAuctionFilterer) FilterBidReceived(opts *bind.FilterOpts) (*DutchAuctionBidReceivedIterator, error) {

	logs, sub, err := _DutchAuction.contract.FilterLogs(opts, "BidReceived")
	if err != nil {
		return nil, err
	}
	return &DutchAuctionBidReceivedIterator{contract: _DutchAuction.contract, event: "BidReceived", logs: logs, sub: sub}, nil
}

// WatchBidReceived is a free log subscription operation binding the contract event 0x491989e56ba54130286a8a051732d8f228353c97d40b74d4584352242b3dba99.
//
// Solidity: event BidReceived(uint256 auction, string id, address bidder, uint256 amount)
func (_DutchAuction *DutchAuctionFilterer) WatchBidReceived(opts *bind.WatchOpts, sink chan<- *DutchAuctionBidReceived) (event.Subscription, error) {

	logs, sub, err := _DutchAuction.contract.WatchLogs(opts, "BidReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchAuctionBidReceived)
				if err := _DutchAuction.contract.UnpackLog(event, "BidReceived", log); err != nil {
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

// ParseBidReceived is a log parse operation binding the contract event 0x491989e56ba54130286a8a051732d8f228353c97d40b74d4584352242b3dba99.
//
// Solidity: event BidReceived(uint256 auction, string id, address bidder, uint256 amount)
func (_DutchAuction *DutchAuctionFilterer) ParseBidReceived(log types.Log) (*DutchAuctionBidReceived, error) {
	event := new(DutchAuctionBidReceived)
	if err := _DutchAuction.contract.UnpackLog(event, "BidReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DutchAuctionDecisionMadeIterator is returned from FilterDecisionMade and is used to iterate over the raw logs and unpacked data for DecisionMade events raised by the DutchAuction contract.
type DutchAuctionDecisionMadeIterator struct {
	Event *DutchAuctionDecisionMade // Event containing the contract specifics and raw log

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
func (it *DutchAuctionDecisionMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchAuctionDecisionMade)
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
		it.Event = new(DutchAuctionDecisionMade)
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
func (it *DutchAuctionDecisionMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchAuctionDecisionMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchAuctionDecisionMade represents a DecisionMade event raised by the DutchAuction contract.
type DutchAuctionDecisionMade struct {
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
func (_DutchAuction *DutchAuctionFilterer) FilterDecisionMade(opts *bind.FilterOpts) (*DutchAuctionDecisionMadeIterator, error) {

	logs, sub, err := _DutchAuction.contract.FilterLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return &DutchAuctionDecisionMadeIterator{contract: _DutchAuction.contract, event: "DecisionMade", logs: logs, sub: sub}, nil
}

// WatchDecisionMade is a free log subscription operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auction, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_DutchAuction *DutchAuctionFilterer) WatchDecisionMade(opts *bind.WatchOpts, sink chan<- *DutchAuctionDecisionMade) (event.Subscription, error) {

	logs, sub, err := _DutchAuction.contract.WatchLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchAuctionDecisionMade)
				if err := _DutchAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
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
func (_DutchAuction *DutchAuctionFilterer) ParseDecisionMade(log types.Log) (*DutchAuctionDecisionMade, error) {
	event := new(DutchAuctionDecisionMade)
	if err := _DutchAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DutchAuctionRateAuctionIterator is returned from FilterRateAuction and is used to iterate over the raw logs and unpacked data for RateAuction events raised by the DutchAuction contract.
type DutchAuctionRateAuctionIterator struct {
	Event *DutchAuctionRateAuction // Event containing the contract specifics and raw log

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
func (it *DutchAuctionRateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchAuctionRateAuction)
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
		it.Event = new(DutchAuctionRateAuction)
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
func (it *DutchAuctionRateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchAuctionRateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchAuctionRateAuction represents a RateAuction event raised by the DutchAuction contract.
type DutchAuctionRateAuction struct {
	Auction *big.Int
	Id      string
	Rating  *big.Int
	Review  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRateAuction is a free log retrieval operation binding the contract event 0x10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e79.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, bytes32 review)
func (_DutchAuction *DutchAuctionFilterer) FilterRateAuction(opts *bind.FilterOpts) (*DutchAuctionRateAuctionIterator, error) {

	logs, sub, err := _DutchAuction.contract.FilterLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return &DutchAuctionRateAuctionIterator{contract: _DutchAuction.contract, event: "RateAuction", logs: logs, sub: sub}, nil
}

// WatchRateAuction is a free log subscription operation binding the contract event 0x10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e79.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, bytes32 review)
func (_DutchAuction *DutchAuctionFilterer) WatchRateAuction(opts *bind.WatchOpts, sink chan<- *DutchAuctionRateAuction) (event.Subscription, error) {

	logs, sub, err := _DutchAuction.contract.WatchLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchAuctionRateAuction)
				if err := _DutchAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
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
func (_DutchAuction *DutchAuctionFilterer) ParseRateAuction(log types.Log) (*DutchAuctionRateAuction, error) {
	event := new(DutchAuctionRateAuction)
	if err := _DutchAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DutchAuctionWithdrawBidIterator is returned from FilterWithdrawBid and is used to iterate over the raw logs and unpacked data for WithdrawBid events raised by the DutchAuction contract.
type DutchAuctionWithdrawBidIterator struct {
	Event *DutchAuctionWithdrawBid // Event containing the contract specifics and raw log

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
func (it *DutchAuctionWithdrawBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchAuctionWithdrawBid)
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
		it.Event = new(DutchAuctionWithdrawBid)
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
func (it *DutchAuctionWithdrawBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchAuctionWithdrawBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchAuctionWithdrawBid represents a WithdrawBid event raised by the DutchAuction contract.
type DutchAuctionWithdrawBid struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBid is a free log retrieval operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_DutchAuction *DutchAuctionFilterer) FilterWithdrawBid(opts *bind.FilterOpts) (*DutchAuctionWithdrawBidIterator, error) {

	logs, sub, err := _DutchAuction.contract.FilterLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return &DutchAuctionWithdrawBidIterator{contract: _DutchAuction.contract, event: "WithdrawBid", logs: logs, sub: sub}, nil
}

// WatchWithdrawBid is a free log subscription operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_DutchAuction *DutchAuctionFilterer) WatchWithdrawBid(opts *bind.WatchOpts, sink chan<- *DutchAuctionWithdrawBid) (event.Subscription, error) {

	logs, sub, err := _DutchAuction.contract.WatchLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchAuctionWithdrawBid)
				if err := _DutchAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
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
func (_DutchAuction *DutchAuctionFilterer) ParseWithdrawBid(log types.Log) (*DutchAuctionWithdrawBid, error) {
	event := new(DutchAuctionWithdrawBid)
	if err := _DutchAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
