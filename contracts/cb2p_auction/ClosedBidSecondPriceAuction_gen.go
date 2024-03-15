// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cb2p_auction

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

// Cb2pAuctionMetaData contains all meta data concerning the Cb2pAuction contract.
var Cb2pAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AwaitResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prcd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"DecisionMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"NewBidHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"rating\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"review\",\"type\":\"bytes32\"}],\"name\":\"RateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"}],\"name\":\"RevealAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkAverageScore\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"not_winner_platform\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"secondHighestPrice\",\"type\":\"uint256\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_asset_id\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_score\",\"type\":\"int256\"},{\"internalType\":\"bytes32\",\"name\":\"_feedback\",\"type\":\"bytes32\"}],\"name\":\"provide_feedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"revealAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"secondHighestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051620021213803806200212183398101604081905261003191610054565b6001600160a01b0316608052600080546001600160a01b03191633179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b608051612073620000ae60003960008181610284015281816108710152610fdf01526120736000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80638da5cb5b116100a2578063b6a46b3b11610071578063b6a46b3b1461023e578063d0a1414a14610251578063dad2393614610264578063ea1591bb1461026c578063fc0c546a1461027f57600080fd5b80638da5cb5b146101e45780639348cef7146101f7578063966ffcff1461020a578063b14c63c51461022b57600080fd5b8063451df52e116100de578063451df52e1461018057806355f78c8d146101ab57806388d3c98e146101be5780638d10b0a6146101d157600080fd5b8063176321e914610110578063274377c0146101255780632e1a7d4d1461013857806342d21ef714610160575b600080fd5b61012361011e3660046119e8565b6102a6565b005b610123610133366004611a2f565b610574565b61014b610146366004611a5b565b610761565b60405190151581526020015b60405180910390f35b61017361016e366004611a5b565b61095d565b6040516101579190611aba565b61019361018e366004611a5b565b610a09565b6040516001600160a01b039091168152602001610157565b6101236101b9366004611a5b565b610a33565b6101236101cc366004611ae5565b610b65565b6101736101df366004611a5b565b610e6e565b600054610193906001600160a01b031681565b610123610205366004611b1d565b610e7e565b61021d610218366004611a5b565b611224565b604051908152602001610157565b61021d610239366004611a5b565b611245565b61012361024c366004611b3f565b611255565b61012361025f3660046119e8565b611450565b61021d61176d565b61012361027a366004611b1d565b61184c565b6101937f000000000000000000000000000000000000000000000000000000000000000081565b60405165656e64696e6760d01b602082015260260160405160208183030381529060405280519060200120600683815481106102e4576102e4611b7c565b906000526020600020016040516020016102fe9190611bcc565b60405160208183030381529060405280519060200120146103665760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064015b60405180910390fd5b6003828154811061037957610379611b7c565b6000918252602090912001546001600160a01b031633146103ac5760405162461bcd60e51b815260040161035d90611c42565b60405180604001604052806007815260200166636c6f73696e6760c81b815250600683815481106103df576103df611b7c565b9060005260206000200190816103f59190611cc1565b507f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7826003848154811061042b5761042b611b7c565b600091825260209091200154600480546001600160a01b03909216918690811061045757610457611b7c565b90600052602060002001546007868154811061047557610475611b7c565b9060005260206000200160008660405161049496959493929190611dfe565b60405180910390a1600482815481106104af576104af611b7c565b906000526020600020015460016000600385815481106104d1576104d1611b7c565b60009182526020808320909101546001600160a01b0316835282019290925260400181208054909190610505908490611e68565b9250508190555060006003838154811061052157610521611b7c565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060006004838154811061056457610564611b7c565b6000918252602090912001555050565b60405166636c6f73696e6760c81b602082015260270160405160208183030381529060405280519060200120600684815481106105b3576105b3611b7c565b906000526020600020016040516020016105cd9190611bcc565b60405160208183030381529060405280519060200120146106305760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e20434c4f53494e47207374617475730000604482015260640161035d565b6003838154811061064357610643611b7c565b6000918252602090912001546001600160a01b031633146106765760405162461bcd60e51b815260040161035d90611c42565b816009848154811061068a5761068a611b7c565b906000526020600020018190555080600884815481106106ac576106ac611b7c565b90600052602060002001819055507f10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e7983600785815481106106ef576106ef611b7c565b90600052602060002001848460405161070b9493929190611e81565b60405180910390a16040518060400160405280600681526020016518db1bdcd95960d21b8152506006848154811061074557610745611b7c565b90600052602060002001908161075b9190611cc1565b50505050565b6040516337b832b760e11b602082015260009060240160405160208183030381529060405280519060200120600683815481106107a0576107a0611b7c565b906000526020600020016040516020016107ba9190611bcc565b604051602081830303815290604052805190602001200361081d5760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420696e204f50454e20737461747573000000000000000000604482015260640161035d565b3360009081526001602052604090205480156108fb573360008181526001602052604080822091909155516323b872dd60e01b81523060048201526024810191909152604481018290526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303816000875af11580156108ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108de9190611ead565b6108fb573360009081526001602052604081209190915592915050565b7f9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c836007858154811061093057610930611b7c565b90600052602060002001338460405161094c9493929190611eca565b60405180910390a150600192915050565b6006818154811061096d57600080fd5b90600052602060002001600091509050805461098890611b92565b80601f01602080910402602001604051908101604052809291908181526020018280546109b490611b92565b8015610a015780601f106109d657610100808354040283529160200191610a01565b820191906000526020600020905b8154815290600101906020018083116109e457829003601f168201915b505050505081565b60038181548110610a1957600080fd5b6000918252602090912001546001600160a01b0316905081565b6000546001600160a01b03163314610a5d5760405162461bcd60e51b815260040161035d90611eff565b6040516337b832b760e11b60208201526024016040516020818303038152906040528051906020012060068281548110610a9957610a99611b7c565b90600052602060002001604051602001610ab39190611bcc565b6040516020818303038152906040528051906020012014610ae65760405162461bcd60e51b815260040161035d90611f46565b604051806040016040528060068152602001651c995d99585b60d21b81525060068281548110610b1857610b18611b7c565b906000526020600020019081610b2e9190611cc1565b506040518181527f3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba9060200160405180910390a150565b6000546001600160a01b03163314610b8f5760405162461bcd60e51b815260040161035d90611eff565b6040516337b832b760e11b60208201526024016040516020818303038152906040528051906020012060068481548110610bcb57610bcb611b7c565b90600052602060002001604051602001610be59190611bcc565b6040516020818303038152906040528051906020012014610c185760405162461bcd60e51b815260040161035d90611f46565b60405180604001604052806006815260200165656e64696e6760d01b81525060068481548110610c4a57610c4a611b7c565b906000526020600020019081610c609190611cc1565b508180610c8a575060048381548110610c7b57610c7b611b7c565b90600052602060002001546000145b15610db1576040518060400160405280600681526020016518db1bdcd95960d21b81525060068481548110610cc157610cc1611b7c565b906000526020600020019081610cd79190611cc1565b5060048381548110610ceb57610ceb611b7c565b90600052602060002001546001600060038681548110610d0d57610d0d611b7c565b60009182526020808320909101546001600160a01b0316835282019290925260400181208054909190610d41908490611e68565b92505081905550600060038481548110610d5d57610d5d611b7c565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600060048481548110610da057610da0611b7c565b600091825260209091200155505050565b60058381548110610dc457610dc4611b7c565b9060005260206000200154811115610df7578060058481548110610dea57610dea611b7c565b6000918252602090912001555b7fa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b48474685002738360038581548110610e2c57610e2c611b7c565b600091825260209091200154604051610e6192916001600160a01b0316909182526001600160a01b0316602082015260400190565b60405180910390a1505050565b6007818154811061096d57600080fd5b604051651c995d99585b60d21b60208201526026016040516020818303038152906040528051906020012060068381548110610ebc57610ebc611b7c565b90600052602060002001604051602001610ed69190611bcc565b6040516020818303038152906040528051906020012014610f095760405162461bcd60e51b815260040161035d90611f46565b60048281548110610f1c57610f1c611b7c565b90600052602060002001548111610f755760405162461bcd60e51b815260206004820152601e60248201527f546865726520616c7265616479206973206120686967686572206269642e0000604482015260640161035d565b600082815260026020908152604080832033845282529182902054825191820184905291016040516020818303038152906040528051906020012014610fba57600080fd5b6040516323b872dd60e01b8152336004820152306024820152604481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015611030573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110549190611ead565b90508061109c5760405162461bcd60e51b81526020600482015260166024820152752a37b5b2b7103a3930b739b332b9103330b4b632b21760511b604482015260640161035d565b600483815481106110af576110af611b7c565b906000526020600020015460001461113057600483815481106110d4576110d4611b7c565b906000526020600020015460016000600386815481106110f6576110f6611b7c565b60009182526020808320909101546001600160a01b031683528201929092526040018120805490919061112a908490611e68565b90915550505b336003848154811061114457611144611b7c565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506004838154811061118557611185611b7c565b9060005260206000200154600584815481106111a3576111a3611b7c565b906000526020600020018190555081600484815481106111c5576111c5611b7c565b90600052602060002001819055507f463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08836007858154811061120857611208611b7c565b906000526020600020013385604051610e619493929190611eca565b6005818154811061123457600080fd5b600091825260209091200154905081565b6004818154811061123457600080fd5b6000546001600160a01b031633146112b95760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79206f776e65722063616e20637265617465206e65772061756374696f6044820152603760f91b606482015260840161035d565b6003805460018181019092557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b031916905560048054808301825560007f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909101819055600580548085019091557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0018190556006805493840181559052604080518082019091529081526337b832b760e11b60208201527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f909101906113aa9082611cc1565b50600780546001810182556000919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688016113e78282611cc1565b505060088054600181810190925560007ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390910181905560098054928301815581527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af90910155565b60405165656e64696e6760d01b6020820152602601604051602081830303815290604052805190602001206006838154811061148e5761148e611b7c565b906000526020600020016040516020016114a89190611bcc565b604051602081830303815290604052805190602001201461150b5760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e4720737461747573000000604482015260640161035d565b6003828154811061151e5761151e611b7c565b6000918252602090912001546001600160a01b031633146115515760405162461bcd60e51b815260040161035d90611c42565b60405180604001604052806007815260200166636c6f73696e6760c81b8152506006838154811061158457611584611b7c565b90600052602060002001908161159a9190611cc1565b50600582815481106115ae576115ae611b7c565b9060005260206000200154600483815481106115cc576115cc611b7c565b9060005260206000200154111561167457600582815481106115f0576115f0611b7c565b90600052602060002001546004838154811061160e5761160e611b7c565b90600052602060002001546116239190611f7d565b600160006003858154811061163a5761163a611b7c565b60009182526020808320909101546001600160a01b031683528201929092526040018120805490919061166e908490611e68565b90915550505b6005828154811061168757611687611b7c565b600091825260208083209091015482546001600160a01b031683526001909152604082208054919290916116bc908490611e68565b925050819055507f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c782600384815481106116f8576116f8611b7c565b600091825260209091200154600480546001600160a01b03909216918690811061172457611724611b7c565b90600052602060002001546007868154811061174257611742611b7c565b9060005260206000200160018660405161176196959493929190611dfe565b60405180910390a15050565b600080805b60095481101561182d576040516518db1bdcd95960d21b602082015260260160405160208183030381529060405280519060200120600682815481106117ba576117ba611b7c565b906000526020600020016040516020016117d49190611bcc565b604051602081830303815290604052805190602001200361181b576009818154811061180257611802611b7c565b9060005260206000200154826118189190611f90565b91505b8061182581611fb8565b915050611772565b5060095461183c826064611fd1565b6118469190612001565b91505090565b6040516337b832b760e11b6020820152602401604051602081830303815290604052805190602001206006838154811061188857611888611b7c565b906000526020600020016040516020016118a29190611bcc565b60405160208183030381529060405280519060200120146118d55760405162461bcd60e51b815260040161035d90611f46565b60008281526002602090815260408083203384529091529020819055600780547f6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f9184918290811061192957611929611b7c565b9060005260206000200133846040516117619493929190611eca565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261196c57600080fd5b813567ffffffffffffffff8082111561198757611987611945565b604051601f8301601f19908116603f011681019082821181831017156119af576119af611945565b816040528381528660208588010111156119c857600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080604083850312156119fb57600080fd5b82359150602083013567ffffffffffffffff811115611a1957600080fd5b611a258582860161195b565b9150509250929050565b600080600060608486031215611a4457600080fd5b505081359360208301359350604090920135919050565b600060208284031215611a6d57600080fd5b5035919050565b6000815180845260005b81811015611a9a57602081850181015186830182015201611a7e565b506000602082860101526020601f19601f83011685010191505092915050565b602081526000611acd6020830184611a74565b9392505050565b8015158114611ae257600080fd5b50565b600080600060608486031215611afa57600080fd5b833592506020840135611b0c81611ad4565b929592945050506040919091013590565b60008060408385031215611b3057600080fd5b50508035926020909101359150565b600060208284031215611b5157600080fd5b813567ffffffffffffffff811115611b6857600080fd5b611b748482850161195b565b949350505050565b634e487b7160e01b600052603260045260246000fd5b600181811c90821680611ba657607f821691505b602082108103611bc657634e487b7160e01b600052602260045260246000fd5b50919050565b6000808354611bda81611b92565b60018281168015611bf25760018114611c0757611c36565b60ff1984168752821515830287019450611c36565b8760005260208060002060005b85811015611c2d5781548a820152908401908201611c14565b50505082870194505b50929695505050505050565b6020808252601690820152754e6f7420617574686f72697a6564206163636573732160501b604082015260600190565b601f821115611cbc57600081815260208120601f850160051c81016020861015611c995750805b601f850160051c820191505b81811015611cb857828155600101611ca5565b5050505b505050565b815167ffffffffffffffff811115611cdb57611cdb611945565b611cef81611ce98454611b92565b84611c72565b602080601f831160018114611d245760008415611d0c5750858301515b600019600386901b1c1916600185901b178555611cb8565b600085815260208120601f198616915b82811015611d5357888601518255948401946001909101908401611d34565b5085821015611d715787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008154611d8e81611b92565b808552602060018381168015611dab5760018114611dc557611df3565b60ff1985168884015283151560051b880183019550611df3565b866000528260002060005b85811015611deb5781548a8201860152908301908401611dd0565b890184019650505b505050505092915050565b86815260018060a01b038616602082015284604082015260c060608201526000611e2b60c0830186611d81565b841515608084015282810360a0840152611e458185611a74565b9998505050505050505050565b634e487b7160e01b600052601160045260246000fd5b80820180821115611e7b57611e7b611e52565b92915050565b848152608060208201526000611e9a6080830186611d81565b6040830194909452506060015292915050565b600060208284031215611ebf57600080fd5b8151611acd81611ad4565b848152608060208201526000611ee36080830186611d81565b6001600160a01b03949094166040830152506060015292915050565b60208082526027908201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736040820152662073746174757360c81b606082015260800190565b6020808252601b908201527f436f6e7472616374206e6f7420696e204f50454e207374617475730000000000604082015260600190565b81810381811115611e7b57611e7b611e52565b8082018281126000831280158216821582161715611fb057611fb0611e52565b505092915050565b600060018201611fca57611fca611e52565b5060010190565b80820260008212600160ff1b84141615611fed57611fed611e52565b8181058314821517611e7b57611e7b611e52565b60008261201e57634e487b7160e01b600052601260045260246000fd5b600160ff1b82146000198414161561203857612038611e52565b50059056fea2646970667358221220cf3f932c2c674a88598a054a92e1e1a7b597b22d6996a0cc16be58c08eed669764736f6c63430008120033",
}

// Cb2pAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use Cb2pAuctionMetaData.ABI instead.
var Cb2pAuctionABI = Cb2pAuctionMetaData.ABI

// Cb2pAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Cb2pAuctionMetaData.Bin instead.
var Cb2pAuctionBin = Cb2pAuctionMetaData.Bin

// DeployCb2pAuction deploys a new Ethereum contract, binding an instance of Cb2pAuction to it.
func DeployCb2pAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *Cb2pAuction, error) {
	parsed, err := Cb2pAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Cb2pAuctionBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cb2pAuction{Cb2pAuctionCaller: Cb2pAuctionCaller{contract: contract}, Cb2pAuctionTransactor: Cb2pAuctionTransactor{contract: contract}, Cb2pAuctionFilterer: Cb2pAuctionFilterer{contract: contract}}, nil
}

// Cb2pAuction is an auto generated Go binding around an Ethereum contract.
type Cb2pAuction struct {
	Cb2pAuctionCaller     // Read-only binding to the contract
	Cb2pAuctionTransactor // Write-only binding to the contract
	Cb2pAuctionFilterer   // Log filterer for contract events
}

// Cb2pAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type Cb2pAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cb2pAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Cb2pAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cb2pAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Cb2pAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cb2pAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Cb2pAuctionSession struct {
	Contract     *Cb2pAuction      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Cb2pAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Cb2pAuctionCallerSession struct {
	Contract *Cb2pAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Cb2pAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Cb2pAuctionTransactorSession struct {
	Contract     *Cb2pAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Cb2pAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type Cb2pAuctionRaw struct {
	Contract *Cb2pAuction // Generic contract binding to access the raw methods on
}

// Cb2pAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Cb2pAuctionCallerRaw struct {
	Contract *Cb2pAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// Cb2pAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Cb2pAuctionTransactorRaw struct {
	Contract *Cb2pAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCb2pAuction creates a new instance of Cb2pAuction, bound to a specific deployed contract.
func NewCb2pAuction(address common.Address, backend bind.ContractBackend) (*Cb2pAuction, error) {
	contract, err := bindCb2pAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cb2pAuction{Cb2pAuctionCaller: Cb2pAuctionCaller{contract: contract}, Cb2pAuctionTransactor: Cb2pAuctionTransactor{contract: contract}, Cb2pAuctionFilterer: Cb2pAuctionFilterer{contract: contract}}, nil
}

// NewCb2pAuctionCaller creates a new read-only instance of Cb2pAuction, bound to a specific deployed contract.
func NewCb2pAuctionCaller(address common.Address, caller bind.ContractCaller) (*Cb2pAuctionCaller, error) {
	contract, err := bindCb2pAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionCaller{contract: contract}, nil
}

// NewCb2pAuctionTransactor creates a new write-only instance of Cb2pAuction, bound to a specific deployed contract.
func NewCb2pAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*Cb2pAuctionTransactor, error) {
	contract, err := bindCb2pAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionTransactor{contract: contract}, nil
}

// NewCb2pAuctionFilterer creates a new log filterer instance of Cb2pAuction, bound to a specific deployed contract.
func NewCb2pAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*Cb2pAuctionFilterer, error) {
	contract, err := bindCb2pAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionFilterer{contract: contract}, nil
}

// bindCb2pAuction binds a generic wrapper to an already deployed contract.
func bindCb2pAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Cb2pAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cb2pAuction *Cb2pAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cb2pAuction.Contract.Cb2pAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cb2pAuction *Cb2pAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Cb2pAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cb2pAuction *Cb2pAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Cb2pAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cb2pAuction *Cb2pAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cb2pAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cb2pAuction *Cb2pAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cb2pAuction *Cb2pAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.contract.Transact(opts, method, params...)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionCaller) AssetId(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "asset_id", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionSession) AssetId(arg0 *big.Int) (string, error) {
	return _Cb2pAuction.Contract.AssetId(&_Cb2pAuction.CallOpts, arg0)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionCallerSession) AssetId(arg0 *big.Int) (string, error) {
	return _Cb2pAuction.Contract.AssetId(&_Cb2pAuction.CallOpts, arg0)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_Cb2pAuction *Cb2pAuctionCaller) CheckAverageScore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "checkAverageScore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_Cb2pAuction *Cb2pAuctionSession) CheckAverageScore() (*big.Int, error) {
	return _Cb2pAuction.Contract.CheckAverageScore(&_Cb2pAuction.CallOpts)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_Cb2pAuction *Cb2pAuctionCallerSession) CheckAverageScore() (*big.Int, error) {
	return _Cb2pAuction.Contract.CheckAverageScore(&_Cb2pAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_Cb2pAuction *Cb2pAuctionCaller) HighestBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "highestBid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_Cb2pAuction *Cb2pAuctionSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _Cb2pAuction.Contract.HighestBid(&_Cb2pAuction.CallOpts, arg0)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_Cb2pAuction *Cb2pAuctionCallerSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _Cb2pAuction.Contract.HighestBid(&_Cb2pAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_Cb2pAuction *Cb2pAuctionCaller) HighestBidder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "highestBidder", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_Cb2pAuction *Cb2pAuctionSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _Cb2pAuction.Contract.HighestBidder(&_Cb2pAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_Cb2pAuction *Cb2pAuctionCallerSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _Cb2pAuction.Contract.HighestBidder(&_Cb2pAuction.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cb2pAuction *Cb2pAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cb2pAuction *Cb2pAuctionSession) Owner() (common.Address, error) {
	return _Cb2pAuction.Contract.Owner(&_Cb2pAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cb2pAuction *Cb2pAuctionCallerSession) Owner() (common.Address, error) {
	return _Cb2pAuction.Contract.Owner(&_Cb2pAuction.CallOpts)
}

// SecondHighestBid is a free data retrieval call binding the contract method 0x966ffcff.
//
// Solidity: function secondHighestBid(uint256 ) view returns(uint256)
func (_Cb2pAuction *Cb2pAuctionCaller) SecondHighestBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "secondHighestBid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondHighestBid is a free data retrieval call binding the contract method 0x966ffcff.
//
// Solidity: function secondHighestBid(uint256 ) view returns(uint256)
func (_Cb2pAuction *Cb2pAuctionSession) SecondHighestBid(arg0 *big.Int) (*big.Int, error) {
	return _Cb2pAuction.Contract.SecondHighestBid(&_Cb2pAuction.CallOpts, arg0)
}

// SecondHighestBid is a free data retrieval call binding the contract method 0x966ffcff.
//
// Solidity: function secondHighestBid(uint256 ) view returns(uint256)
func (_Cb2pAuction *Cb2pAuctionCallerSession) SecondHighestBid(arg0 *big.Int) (*big.Int, error) {
	return _Cb2pAuction.Contract.SecondHighestBid(&_Cb2pAuction.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionCaller) Status(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "status", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionSession) Status(arg0 *big.Int) (string, error) {
	return _Cb2pAuction.Contract.Status(&_Cb2pAuction.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionCallerSession) Status(arg0 *big.Int) (string, error) {
	return _Cb2pAuction.Contract.Status(&_Cb2pAuction.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cb2pAuction *Cb2pAuctionCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cb2pAuction *Cb2pAuctionSession) Token() (common.Address, error) {
	return _Cb2pAuction.Contract.Token(&_Cb2pAuction.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cb2pAuction *Cb2pAuctionCallerSession) Token() (common.Address, error) {
	return _Cb2pAuction.Contract.Token(&_Cb2pAuction.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) Abort(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "abort", auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_Cb2pAuction *Cb2pAuctionSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Abort(&_Cb2pAuction.TransactOpts, auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Abort(&_Cb2pAuction.TransactOpts, auctionId, jsonString)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) Bid(opts *bind.TransactOpts, auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "bid", auctionId, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_Cb2pAuction *Cb2pAuctionSession) Bid(auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Bid(&_Cb2pAuction.TransactOpts, auctionId, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) Bid(auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Bid(&_Cb2pAuction.TransactOpts, auctionId, bidHash)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x88d3c98e.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform, uint256 secondHighestPrice) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) CloseAuction(opts *bind.TransactOpts, auctionId *big.Int, not_winner_platform bool, secondHighestPrice *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "closeAuction", auctionId, not_winner_platform, secondHighestPrice)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x88d3c98e.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform, uint256 secondHighestPrice) returns()
func (_Cb2pAuction *Cb2pAuctionSession) CloseAuction(auctionId *big.Int, not_winner_platform bool, secondHighestPrice *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.CloseAuction(&_Cb2pAuction.TransactOpts, auctionId, not_winner_platform, secondHighestPrice)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x88d3c98e.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform, uint256 secondHighestPrice) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) CloseAuction(auctionId *big.Int, not_winner_platform bool, secondHighestPrice *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.CloseAuction(&_Cb2pAuction.TransactOpts, auctionId, not_winner_platform, secondHighestPrice)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) Commit(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "commit", auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_Cb2pAuction *Cb2pAuctionSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Commit(&_Cb2pAuction.TransactOpts, auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Commit(&_Cb2pAuction.TransactOpts, auctionId, jsonString)
}

// Create is a paid mutator transaction binding the contract method 0xb6a46b3b.
//
// Solidity: function create(string _asset_id) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) Create(opts *bind.TransactOpts, _asset_id string) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "create", _asset_id)
}

// Create is a paid mutator transaction binding the contract method 0xb6a46b3b.
//
// Solidity: function create(string _asset_id) returns()
func (_Cb2pAuction *Cb2pAuctionSession) Create(_asset_id string) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Create(&_Cb2pAuction.TransactOpts, _asset_id)
}

// Create is a paid mutator transaction binding the contract method 0xb6a46b3b.
//
// Solidity: function create(string _asset_id) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) Create(_asset_id string) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Create(&_Cb2pAuction.TransactOpts, _asset_id)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) ProvideFeedback(opts *bind.TransactOpts, auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "provide_feedback", auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_Cb2pAuction *Cb2pAuctionSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.ProvideFeedback(&_Cb2pAuction.TransactOpts, auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.ProvideFeedback(&_Cb2pAuction.TransactOpts, auctionId, _score, _feedback)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) Reveal(opts *bind.TransactOpts, auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "reveal", auctionId, bidAmount)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_Cb2pAuction *Cb2pAuctionSession) Reveal(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Reveal(&_Cb2pAuction.TransactOpts, auctionId, bidAmount)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) Reveal(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Reveal(&_Cb2pAuction.TransactOpts, auctionId, bidAmount)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) RevealAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "revealAuction", auctionId)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_Cb2pAuction *Cb2pAuctionSession) RevealAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.RevealAuction(&_Cb2pAuction.TransactOpts, auctionId)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) RevealAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.RevealAuction(&_Cb2pAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_Cb2pAuction *Cb2pAuctionTransactor) Withdraw(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "withdraw", auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_Cb2pAuction *Cb2pAuctionSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Withdraw(&_Cb2pAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_Cb2pAuction *Cb2pAuctionTransactorSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Withdraw(&_Cb2pAuction.TransactOpts, auctionId)
}

// Cb2pAuctionAwaitResponseIterator is returned from FilterAwaitResponse and is used to iterate over the raw logs and unpacked data for AwaitResponse events raised by the Cb2pAuction contract.
type Cb2pAuctionAwaitResponseIterator struct {
	Event *Cb2pAuctionAwaitResponse // Event containing the contract specifics and raw log

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
func (it *Cb2pAuctionAwaitResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb2pAuctionAwaitResponse)
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
		it.Event = new(Cb2pAuctionAwaitResponse)
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
func (it *Cb2pAuctionAwaitResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb2pAuctionAwaitResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb2pAuctionAwaitResponse represents a AwaitResponse event raised by the Cb2pAuction contract.
type Cb2pAuctionAwaitResponse struct {
	Auction *big.Int
	Winner  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAwaitResponse is a free log retrieval operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_Cb2pAuction *Cb2pAuctionFilterer) FilterAwaitResponse(opts *bind.FilterOpts) (*Cb2pAuctionAwaitResponseIterator, error) {

	logs, sub, err := _Cb2pAuction.contract.FilterLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionAwaitResponseIterator{contract: _Cb2pAuction.contract, event: "AwaitResponse", logs: logs, sub: sub}, nil
}

// WatchAwaitResponse is a free log subscription operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_Cb2pAuction *Cb2pAuctionFilterer) WatchAwaitResponse(opts *bind.WatchOpts, sink chan<- *Cb2pAuctionAwaitResponse) (event.Subscription, error) {

	logs, sub, err := _Cb2pAuction.contract.WatchLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb2pAuctionAwaitResponse)
				if err := _Cb2pAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
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
func (_Cb2pAuction *Cb2pAuctionFilterer) ParseAwaitResponse(log types.Log) (*Cb2pAuctionAwaitResponse, error) {
	event := new(Cb2pAuctionAwaitResponse)
	if err := _Cb2pAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb2pAuctionDecisionMadeIterator is returned from FilterDecisionMade and is used to iterate over the raw logs and unpacked data for DecisionMade events raised by the Cb2pAuction contract.
type Cb2pAuctionDecisionMadeIterator struct {
	Event *Cb2pAuctionDecisionMade // Event containing the contract specifics and raw log

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
func (it *Cb2pAuctionDecisionMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb2pAuctionDecisionMade)
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
		it.Event = new(Cb2pAuctionDecisionMade)
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
func (it *Cb2pAuctionDecisionMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb2pAuctionDecisionMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb2pAuctionDecisionMade represents a DecisionMade event raised by the Cb2pAuction contract.
type Cb2pAuctionDecisionMade struct {
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
func (_Cb2pAuction *Cb2pAuctionFilterer) FilterDecisionMade(opts *bind.FilterOpts) (*Cb2pAuctionDecisionMadeIterator, error) {

	logs, sub, err := _Cb2pAuction.contract.FilterLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionDecisionMadeIterator{contract: _Cb2pAuction.contract, event: "DecisionMade", logs: logs, sub: sub}, nil
}

// WatchDecisionMade is a free log subscription operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auction, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_Cb2pAuction *Cb2pAuctionFilterer) WatchDecisionMade(opts *bind.WatchOpts, sink chan<- *Cb2pAuctionDecisionMade) (event.Subscription, error) {

	logs, sub, err := _Cb2pAuction.contract.WatchLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb2pAuctionDecisionMade)
				if err := _Cb2pAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
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
func (_Cb2pAuction *Cb2pAuctionFilterer) ParseDecisionMade(log types.Log) (*Cb2pAuctionDecisionMade, error) {
	event := new(Cb2pAuctionDecisionMade)
	if err := _Cb2pAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb2pAuctionHighestBidIncreasedIterator is returned from FilterHighestBidIncreased and is used to iterate over the raw logs and unpacked data for HighestBidIncreased events raised by the Cb2pAuction contract.
type Cb2pAuctionHighestBidIncreasedIterator struct {
	Event *Cb2pAuctionHighestBidIncreased // Event containing the contract specifics and raw log

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
func (it *Cb2pAuctionHighestBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb2pAuctionHighestBidIncreased)
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
		it.Event = new(Cb2pAuctionHighestBidIncreased)
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
func (it *Cb2pAuctionHighestBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb2pAuctionHighestBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb2pAuctionHighestBidIncreased represents a HighestBidIncreased event raised by the Cb2pAuction contract.
type Cb2pAuctionHighestBidIncreased struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_Cb2pAuction *Cb2pAuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*Cb2pAuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _Cb2pAuction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionHighestBidIncreasedIterator{contract: _Cb2pAuction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_Cb2pAuction *Cb2pAuctionFilterer) WatchHighestBidIncreased(opts *bind.WatchOpts, sink chan<- *Cb2pAuctionHighestBidIncreased) (event.Subscription, error) {

	logs, sub, err := _Cb2pAuction.contract.WatchLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb2pAuctionHighestBidIncreased)
				if err := _Cb2pAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
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
func (_Cb2pAuction *Cb2pAuctionFilterer) ParseHighestBidIncreased(log types.Log) (*Cb2pAuctionHighestBidIncreased, error) {
	event := new(Cb2pAuctionHighestBidIncreased)
	if err := _Cb2pAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb2pAuctionNewBidHashIterator is returned from FilterNewBidHash and is used to iterate over the raw logs and unpacked data for NewBidHash events raised by the Cb2pAuction contract.
type Cb2pAuctionNewBidHashIterator struct {
	Event *Cb2pAuctionNewBidHash // Event containing the contract specifics and raw log

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
func (it *Cb2pAuctionNewBidHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb2pAuctionNewBidHash)
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
		it.Event = new(Cb2pAuctionNewBidHash)
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
func (it *Cb2pAuctionNewBidHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb2pAuctionNewBidHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb2pAuctionNewBidHash represents a NewBidHash event raised by the Cb2pAuction contract.
type Cb2pAuctionNewBidHash struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	BidHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBidHash is a free log retrieval operation binding the contract event 0x6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f.
//
// Solidity: event NewBidHash(uint256 auction, string id, address bidder, bytes32 bidHash)
func (_Cb2pAuction *Cb2pAuctionFilterer) FilterNewBidHash(opts *bind.FilterOpts) (*Cb2pAuctionNewBidHashIterator, error) {

	logs, sub, err := _Cb2pAuction.contract.FilterLogs(opts, "NewBidHash")
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionNewBidHashIterator{contract: _Cb2pAuction.contract, event: "NewBidHash", logs: logs, sub: sub}, nil
}

// WatchNewBidHash is a free log subscription operation binding the contract event 0x6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f.
//
// Solidity: event NewBidHash(uint256 auction, string id, address bidder, bytes32 bidHash)
func (_Cb2pAuction *Cb2pAuctionFilterer) WatchNewBidHash(opts *bind.WatchOpts, sink chan<- *Cb2pAuctionNewBidHash) (event.Subscription, error) {

	logs, sub, err := _Cb2pAuction.contract.WatchLogs(opts, "NewBidHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb2pAuctionNewBidHash)
				if err := _Cb2pAuction.contract.UnpackLog(event, "NewBidHash", log); err != nil {
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
func (_Cb2pAuction *Cb2pAuctionFilterer) ParseNewBidHash(log types.Log) (*Cb2pAuctionNewBidHash, error) {
	event := new(Cb2pAuctionNewBidHash)
	if err := _Cb2pAuction.contract.UnpackLog(event, "NewBidHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb2pAuctionRateAuctionIterator is returned from FilterRateAuction and is used to iterate over the raw logs and unpacked data for RateAuction events raised by the Cb2pAuction contract.
type Cb2pAuctionRateAuctionIterator struct {
	Event *Cb2pAuctionRateAuction // Event containing the contract specifics and raw log

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
func (it *Cb2pAuctionRateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb2pAuctionRateAuction)
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
		it.Event = new(Cb2pAuctionRateAuction)
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
func (it *Cb2pAuctionRateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb2pAuctionRateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb2pAuctionRateAuction represents a RateAuction event raised by the Cb2pAuction contract.
type Cb2pAuctionRateAuction struct {
	Auction *big.Int
	Id      string
	Rating  *big.Int
	Review  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRateAuction is a free log retrieval operation binding the contract event 0x10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e79.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, bytes32 review)
func (_Cb2pAuction *Cb2pAuctionFilterer) FilterRateAuction(opts *bind.FilterOpts) (*Cb2pAuctionRateAuctionIterator, error) {

	logs, sub, err := _Cb2pAuction.contract.FilterLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionRateAuctionIterator{contract: _Cb2pAuction.contract, event: "RateAuction", logs: logs, sub: sub}, nil
}

// WatchRateAuction is a free log subscription operation binding the contract event 0x10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e79.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, bytes32 review)
func (_Cb2pAuction *Cb2pAuctionFilterer) WatchRateAuction(opts *bind.WatchOpts, sink chan<- *Cb2pAuctionRateAuction) (event.Subscription, error) {

	logs, sub, err := _Cb2pAuction.contract.WatchLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb2pAuctionRateAuction)
				if err := _Cb2pAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
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
func (_Cb2pAuction *Cb2pAuctionFilterer) ParseRateAuction(log types.Log) (*Cb2pAuctionRateAuction, error) {
	event := new(Cb2pAuctionRateAuction)
	if err := _Cb2pAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb2pAuctionRevealAuctionIterator is returned from FilterRevealAuction and is used to iterate over the raw logs and unpacked data for RevealAuction events raised by the Cb2pAuction contract.
type Cb2pAuctionRevealAuctionIterator struct {
	Event *Cb2pAuctionRevealAuction // Event containing the contract specifics and raw log

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
func (it *Cb2pAuctionRevealAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb2pAuctionRevealAuction)
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
		it.Event = new(Cb2pAuctionRevealAuction)
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
func (it *Cb2pAuctionRevealAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb2pAuctionRevealAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb2pAuctionRevealAuction represents a RevealAuction event raised by the Cb2pAuction contract.
type Cb2pAuctionRevealAuction struct {
	Auction *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevealAuction is a free log retrieval operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auction)
func (_Cb2pAuction *Cb2pAuctionFilterer) FilterRevealAuction(opts *bind.FilterOpts) (*Cb2pAuctionRevealAuctionIterator, error) {

	logs, sub, err := _Cb2pAuction.contract.FilterLogs(opts, "RevealAuction")
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionRevealAuctionIterator{contract: _Cb2pAuction.contract, event: "RevealAuction", logs: logs, sub: sub}, nil
}

// WatchRevealAuction is a free log subscription operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auction)
func (_Cb2pAuction *Cb2pAuctionFilterer) WatchRevealAuction(opts *bind.WatchOpts, sink chan<- *Cb2pAuctionRevealAuction) (event.Subscription, error) {

	logs, sub, err := _Cb2pAuction.contract.WatchLogs(opts, "RevealAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb2pAuctionRevealAuction)
				if err := _Cb2pAuction.contract.UnpackLog(event, "RevealAuction", log); err != nil {
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
func (_Cb2pAuction *Cb2pAuctionFilterer) ParseRevealAuction(log types.Log) (*Cb2pAuctionRevealAuction, error) {
	event := new(Cb2pAuctionRevealAuction)
	if err := _Cb2pAuction.contract.UnpackLog(event, "RevealAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb2pAuctionWithdrawBidIterator is returned from FilterWithdrawBid and is used to iterate over the raw logs and unpacked data for WithdrawBid events raised by the Cb2pAuction contract.
type Cb2pAuctionWithdrawBidIterator struct {
	Event *Cb2pAuctionWithdrawBid // Event containing the contract specifics and raw log

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
func (it *Cb2pAuctionWithdrawBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb2pAuctionWithdrawBid)
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
		it.Event = new(Cb2pAuctionWithdrawBid)
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
func (it *Cb2pAuctionWithdrawBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb2pAuctionWithdrawBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb2pAuctionWithdrawBid represents a WithdrawBid event raised by the Cb2pAuction contract.
type Cb2pAuctionWithdrawBid struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBid is a free log retrieval operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_Cb2pAuction *Cb2pAuctionFilterer) FilterWithdrawBid(opts *bind.FilterOpts) (*Cb2pAuctionWithdrawBidIterator, error) {

	logs, sub, err := _Cb2pAuction.contract.FilterLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionWithdrawBidIterator{contract: _Cb2pAuction.contract, event: "WithdrawBid", logs: logs, sub: sub}, nil
}

// WatchWithdrawBid is a free log subscription operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_Cb2pAuction *Cb2pAuctionFilterer) WatchWithdrawBid(opts *bind.WatchOpts, sink chan<- *Cb2pAuctionWithdrawBid) (event.Subscription, error) {

	logs, sub, err := _Cb2pAuction.contract.WatchLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb2pAuctionWithdrawBid)
				if err := _Cb2pAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
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
func (_Cb2pAuction *Cb2pAuctionFilterer) ParseWithdrawBid(log types.Log) (*Cb2pAuctionWithdrawBid, error) {
	event := new(Cb2pAuctionWithdrawBid)
	if err := _Cb2pAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
