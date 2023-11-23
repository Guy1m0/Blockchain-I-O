// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_closed_bid_second_auction

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

// EthClosedBidSecondAuctionMetaData contains all meta data concerning the EthClosedBidSecondAuction contract.
var EthClosedBidSecondAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AwaitResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prcd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"DecisionMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"NewBidHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"rating\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"review\",\"type\":\"bytes32\"}],\"name\":\"RateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"}],\"name\":\"RevealAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkAverageScore\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"not_winner_platform\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"secondHighestPrice\",\"type\":\"uint256\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_asset_id\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_score\",\"type\":\"int256\"},{\"internalType\":\"bytes32\",\"name\":\"_feedback\",\"type\":\"bytes32\"}],\"name\":\"provide_feedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"revealAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"secondHighestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b5060405161204238038061204283398101604081905261002e91610050565b6001600160a01b03166080525f80546001600160a01b0319163317905561007d565b5f60208284031215610060575f80fd5b81516001600160a01b0381168114610076575f80fd5b9392505050565b608051611f9f6100a35f395f818161027d015281816108450152610f800152611f9f5ff3fe608060405234801561000f575f80fd5b5060043610610106575f3560e01c80638da5cb5b1161009e578063b6a46b3b1161006e578063b6a46b3b14610237578063d0a1414a1461024a578063dad239361461025d578063ea1591bb14610265578063fc0c546a14610278575f80fd5b80638da5cb5b146101de5780639348cef7146101f0578063966ffcff14610203578063b14c63c514610224575f80fd5b8063451df52e116100d9578063451df52e1461017a57806355f78c8d146101a557806388d3c98e146101b85780638d10b0a6146101cb575f80fd5b8063176321e91461010a578063274377c01461011f5780632e1a7d4d1461013257806342d21ef71461015a575b5f80fd5b61011d610118366004611948565b61029f565b005b61011d61012d36600461198c565b610558565b6101456101403660046119b5565b61073a565b60405190151581526020015b60405180910390f35b61016d6101683660046119b5565b61092b565b6040516101519190611a0f565b61018d6101883660046119b5565b6109d1565b6040516001600160a01b039091168152602001610151565b61011d6101b33660046119b5565b6109f9565b61011d6101c6366004611a38565b610b26565b61016d6101d93660046119b5565b610e17565b5f5461018d906001600160a01b031681565b61011d6101fe366004611a6d565b610e26565b6102166102113660046119b5565b6111b0565b604051908152602001610151565b6102166102323660046119b5565b6111cf565b61011d610245366004611a8d565b6111de565b61011d610258366004611948565b6113d5565b6102166116dd565b61011d610273366004611a6d565b6117b7565b61018d7f000000000000000000000000000000000000000000000000000000000000000081565b60405165656e64696e6760d01b602082015260260160405160208183030381529060405280519060200120600683815481106102dd576102dd611ac7565b905f5260205f20016040516020016102f59190611b13565b604051602081830303815290604052805190602001201461035d5760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064015b60405180910390fd5b6003828154811061037057610370611ac7565b5f918252602090912001546001600160a01b031633146103a25760405162461bcd60e51b815260040161035490611b85565b60405180604001604052806007815260200166636c6f73696e6760c81b815250600683815481106103d5576103d5611ac7565b905f5260205f200190816103e99190611c03565b507f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7826003848154811061041f5761041f611ac7565b5f91825260209091200154600480546001600160a01b03909216918690811061044a5761044a611ac7565b905f5260205f2001546007868154811061046657610466611ac7565b905f5260205f20015f8660405161048296959493929190611d38565b60405180910390a16004828154811061049d5761049d611ac7565b905f5260205f20015460015f600385815481106104bc576104bc611ac7565b5f9182526020808320909101546001600160a01b03168352820192909252604001812080549091906104ef908490611d9f565b925050819055505f6003838154811061050a5761050a611ac7565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055505f6004838154811061054957610549611ac7565b5f918252602090912001555050565b60405166636c6f73696e6760c81b6020820152602701604051602081830303815290604052805190602001206006848154811061059757610597611ac7565b905f5260205f20016040516020016105af9190611b13565b60405160208183030381529060405280519060200120146106125760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e20434c4f53494e472073746174757300006044820152606401610354565b6003838154811061062557610625611ac7565b5f918252602090912001546001600160a01b031633146106575760405162461bcd60e51b815260040161035490611b85565b816009848154811061066b5761066b611ac7565b905f5260205f200181905550806008848154811061068b5761068b611ac7565b905f5260205f2001819055507f10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e7983600785815481106106cc576106cc611ac7565b905f5260205f200184846040516106e69493929190611db8565b60405180910390a16040518060400160405280600681526020016518db1bdcd95960d21b8152506006848154811061072057610720611ac7565b905f5260205f200190816107349190611c03565b50505050565b6040516337b832b760e11b60208201525f90602401604051602081830303815290604052805190602001206006838154811061077857610778611ac7565b905f5260205f20016040516020016107909190611b13565b60405160208183030381529060405280519060200120036107f35760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420696e204f50454e207374617475730000000000000000006044820152606401610354565b335f9081526001602052604090205480156108cb57335f8181526001602052604080822091909155516323b872dd60e01b81523060048201526024810191909152604481018290526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303815f875af115801561088b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108af9190611de3565b6108cb57335f9081526001602052604081209190915592915050565b7f9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c836007858154811061090057610900611ac7565b905f5260205f2001338460405161091a9493929190611dfe565b60405180910390a150600192915050565b6006818154811061093a575f80fd5b905f5260205f20015f91509050805461095290611adb565b80601f016020809104026020016040519081016040528092919081815260200182805461097e90611adb565b80156109c95780601f106109a0576101008083540402835291602001916109c9565b820191905f5260205f20905b8154815290600101906020018083116109ac57829003601f168201915b505050505081565b600381815481106109e0575f80fd5b5f918252602090912001546001600160a01b0316905081565b5f546001600160a01b03163314610a225760405162461bcd60e51b815260040161035490611e32565b6040516337b832b760e11b60208201526024016040516020818303038152906040528051906020012060068281548110610a5e57610a5e611ac7565b905f5260205f2001604051602001610a769190611b13565b6040516020818303038152906040528051906020012014610aa95760405162461bcd60e51b815260040161035490611e79565b604051806040016040528060068152602001651c995d99585b60d21b81525060068281548110610adb57610adb611ac7565b905f5260205f20019081610aef9190611c03565b506040518181527f3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba9060200160405180910390a150565b5f546001600160a01b03163314610b4f5760405162461bcd60e51b815260040161035490611e32565b6040516337b832b760e11b60208201526024016040516020818303038152906040528051906020012060068481548110610b8b57610b8b611ac7565b905f5260205f2001604051602001610ba39190611b13565b6040516020818303038152906040528051906020012014610bd65760405162461bcd60e51b815260040161035490611e79565b60405180604001604052806006815260200165656e64696e6760d01b81525060068481548110610c0857610c08611ac7565b905f5260205f20019081610c1c9190611c03565b508180610c43575060048381548110610c3757610c37611ac7565b905f5260205f2001545f145b15610d5e576040518060400160405280600681526020016518db1bdcd95960d21b81525060068481548110610c7a57610c7a611ac7565b905f5260205f20019081610c8e9190611c03565b5060048381548110610ca257610ca2611ac7565b905f5260205f20015460015f60038681548110610cc157610cc1611ac7565b5f9182526020808320909101546001600160a01b0316835282019290925260400181208054909190610cf4908490611d9f565b925050819055505f60038481548110610d0f57610d0f611ac7565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055505f60048481548110610d4e57610d4e611ac7565b5f91825260209091200155505050565b60058381548110610d7157610d71611ac7565b905f5260205f200154811115610da1578060058481548110610d9557610d95611ac7565b5f918252602090912001555b7fa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b48474685002738360038581548110610dd657610dd6611ac7565b5f91825260209091200154604051610e0a92916001600160a01b0316909182526001600160a01b0316602082015260400190565b60405180910390a1505050565b6007818154811061093a575f80fd5b604051651c995d99585b60d21b60208201526026016040516020818303038152906040528051906020012060068381548110610e6457610e64611ac7565b905f5260205f2001604051602001610e7c9190611b13565b6040516020818303038152906040528051906020012014610eaf5760405162461bcd60e51b815260040161035490611e79565b60048281548110610ec257610ec2611ac7565b905f5260205f2001548111610f195760405162461bcd60e51b815260206004820152601e60248201527f546865726520616c7265616479206973206120686967686572206269642e00006044820152606401610354565b5f82815260026020908152604080832033845282529182902054825191820184905291016040516020818303038152906040528051906020012014610f5c575f80fd5b6040516323b872dd60e01b8152336004820152306024820152604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303815f875af1158015610fce573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ff29190611de3565b90508061103a5760405162461bcd60e51b81526020600482015260166024820152752a37b5b2b7103a3930b739b332b9103330b4b632b21760511b6044820152606401610354565b6004838154811061104d5761104d611ac7565b905f5260205f2001545f146110c7576004838154811061106f5761106f611ac7565b905f5260205f20015460015f6003868154811061108e5761108e611ac7565b5f9182526020808320909101546001600160a01b03168352820192909252604001812080549091906110c1908490611d9f565b90915550505b33600384815481106110db576110db611ac7565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506004838154811061111957611119611ac7565b905f5260205f2001546005848154811061113557611135611ac7565b905f5260205f200181905550816004848154811061115557611155611ac7565b905f5260205f2001819055507f463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08836007858154811061119657611196611ac7565b905f5260205f20013385604051610e0a9493929190611dfe565b600581815481106111bf575f80fd5b5f91825260209091200154905081565b600481815481106111bf575f80fd5b5f546001600160a01b031633146112415760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79206f776e65722063616e20637265617465206e65772061756374696f6044820152603760f91b6064820152608401610354565b6003805460018181019092557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b03191690556004805480830182555f7f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909101819055600580548085019091557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0018190556006805493840181559052604080518082019091529081526337b832b760e11b60208201527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f909101906113319082611c03565b50600780546001810182555f919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880161136d8282611c03565b50506008805460018181019092555f7ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390910181905560098054928301815581527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af90910155565b60405165656e64696e6760d01b6020820152602601604051602081830303815290604052805190602001206006838154811061141357611413611ac7565b905f5260205f200160405160200161142b9190611b13565b604051602081830303815290604052805190602001201461148e5760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e47207374617475730000006044820152606401610354565b600382815481106114a1576114a1611ac7565b5f918252602090912001546001600160a01b031633146114d35760405162461bcd60e51b815260040161035490611b85565b60405180604001604052806007815260200166636c6f73696e6760c81b8152506006838154811061150657611506611ac7565b905f5260205f2001908161151a9190611c03565b506005828154811061152e5761152e611ac7565b905f5260205f2001546004838154811061154a5761154a611ac7565b905f5260205f20015411156115ea576005828154811061156c5761156c611ac7565b905f5260205f2001546004838154811061158857611588611ac7565b905f5260205f20015461159b9190611eb0565b60015f600385815481106115b1576115b1611ac7565b5f9182526020808320909101546001600160a01b03168352820192909252604001812080549091906115e4908490611d9f565b90915550505b600582815481106115fd576115fd611ac7565b5f91825260208083209091015482546001600160a01b03168352600190915260408220805491929091611631908490611d9f565b925050819055507f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7826003848154811061166d5761166d611ac7565b5f91825260209091200154600480546001600160a01b03909216918690811061169857611698611ac7565b905f5260205f200154600786815481106116b4576116b4611ac7565b905f5260205f20016001866040516116d196959493929190611d38565b60405180910390a15050565b5f80805b600954811015611798576040516518db1bdcd95960d21b6020820152602601604051602081830303815290604052805190602001206006828154811061172957611729611ac7565b905f5260205f20016040516020016117419190611b13565b6040516020818303038152906040528051906020012003611786576009818154811061176f5761176f611ac7565b905f5260205f200154826117839190611ec3565b91505b8061179081611eea565b9150506116e1565b506009546117a7826064611f02565b6117b19190611f31565b91505090565b6040516337b832b760e11b602082015260240160405160208183030381529060405280519060200120600683815481106117f3576117f3611ac7565b905f5260205f200160405160200161180b9190611b13565b604051602081830303815290604052805190602001201461183e5760405162461bcd60e51b815260040161035490611e79565b5f8281526002602090815260408083203384529091529020819055600780547f6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f9184918290811061189157611891611ac7565b905f5260205f200133846040516116d19493929190611dfe565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126118ce575f80fd5b813567ffffffffffffffff808211156118e9576118e96118ab565b604051601f8301601f19908116603f01168101908282118183101715611911576119116118ab565b81604052838152866020858801011115611929575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f8060408385031215611959575f80fd5b82359150602083013567ffffffffffffffff811115611976575f80fd5b611982858286016118bf565b9150509250929050565b5f805f6060848603121561199e575f80fd5b505081359360208301359350604090920135919050565b5f602082840312156119c5575f80fd5b5035919050565b5f81518084525f5b818110156119f0576020818501810151868301820152016119d4565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f611a2160208301846119cc565b9392505050565b8015158114611a35575f80fd5b50565b5f805f60608486031215611a4a575f80fd5b833592506020840135611a5c81611a28565b929592945050506040919091013590565b5f8060408385031215611a7e575f80fd5b50508035926020909101359150565b5f60208284031215611a9d575f80fd5b813567ffffffffffffffff811115611ab3575f80fd5b611abf848285016118bf565b949350505050565b634e487b7160e01b5f52603260045260245ffd5b600181811c90821680611aef57607f821691505b602082108103611b0d57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f808354611b2081611adb565b60018281168015611b385760018114611b4d57611b79565b60ff1984168752821515830287019450611b79565b875f526020805f205f5b85811015611b705781548a820152908401908201611b57565b50505082870194505b50929695505050505050565b6020808252601690820152754e6f7420617574686f72697a6564206163636573732160501b604082015260600190565b601f821115611bfe575f81815260208120601f850160051c81016020861015611bdb5750805b601f850160051c820191505b81811015611bfa57828155600101611be7565b5050505b505050565b815167ffffffffffffffff811115611c1d57611c1d6118ab565b611c3181611c2b8454611adb565b84611bb5565b602080601f831160018114611c64575f8415611c4d5750858301515b5f19600386901b1c1916600185901b178555611bfa565b5f85815260208120601f198616915b82811015611c9257888601518255948401946001909101908401611c73565b5085821015611caf57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f8154611ccb81611adb565b808552602060018381168015611ce85760018114611d0257611d2d565b60ff1985168884015283151560051b880183019550611d2d565b865f52825f205f5b85811015611d255781548a8201860152908301908401611d0a565b890184019650505b505050505092915050565b86815260018060a01b038616602082015284604082015260c060608201525f611d6460c0830186611cbf565b841515608084015282810360a0840152611d7e81856119cc565b9998505050505050505050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115611db257611db2611d8b565b92915050565b848152608060208201525f611dd06080830186611cbf565b6040830194909452506060015292915050565b5f60208284031215611df3575f80fd5b8151611a2181611a28565b848152608060208201525f611e166080830186611cbf565b6001600160a01b03949094166040830152506060015292915050565b60208082526027908201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736040820152662073746174757360c81b606082015260800190565b6020808252601b908201527f436f6e7472616374206e6f7420696e204f50454e207374617475730000000000604082015260600190565b81810381811115611db257611db2611d8b565b8082018281125f831280158216821582161715611ee257611ee2611d8b565b505092915050565b5f60018201611efb57611efb611d8b565b5060010190565b8082025f8212600160ff1b84141615611f1d57611f1d611d8b565b8181058314821517611db257611db2611d8b565b5f82611f4b57634e487b7160e01b5f52601260045260245ffd5b600160ff1b82145f1984141615611f6457611f64611d8b565b50059056fea26469706673582212207febd709a51b24d612fa909a5195f1c56f19ca7861276ffb0582eb9e1ed6d88664736f6c63430008140033",
}

// EthClosedBidSecondAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use EthClosedBidSecondAuctionMetaData.ABI instead.
var EthClosedBidSecondAuctionABI = EthClosedBidSecondAuctionMetaData.ABI

// EthClosedBidSecondAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthClosedBidSecondAuctionMetaData.Bin instead.
var EthClosedBidSecondAuctionBin = EthClosedBidSecondAuctionMetaData.Bin

// DeployEthClosedBidSecondAuction deploys a new Ethereum contract, binding an instance of EthClosedBidSecondAuction to it.
func DeployEthClosedBidSecondAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *EthClosedBidSecondAuction, error) {
	parsed, err := EthClosedBidSecondAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthClosedBidSecondAuctionBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthClosedBidSecondAuction{EthClosedBidSecondAuctionCaller: EthClosedBidSecondAuctionCaller{contract: contract}, EthClosedBidSecondAuctionTransactor: EthClosedBidSecondAuctionTransactor{contract: contract}, EthClosedBidSecondAuctionFilterer: EthClosedBidSecondAuctionFilterer{contract: contract}}, nil
}

// EthClosedBidSecondAuction is an auto generated Go binding around an Ethereum contract.
type EthClosedBidSecondAuction struct {
	EthClosedBidSecondAuctionCaller     // Read-only binding to the contract
	EthClosedBidSecondAuctionTransactor // Write-only binding to the contract
	EthClosedBidSecondAuctionFilterer   // Log filterer for contract events
}

// EthClosedBidSecondAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthClosedBidSecondAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClosedBidSecondAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthClosedBidSecondAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClosedBidSecondAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthClosedBidSecondAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClosedBidSecondAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthClosedBidSecondAuctionSession struct {
	Contract     *EthClosedBidSecondAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EthClosedBidSecondAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthClosedBidSecondAuctionCallerSession struct {
	Contract *EthClosedBidSecondAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// EthClosedBidSecondAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthClosedBidSecondAuctionTransactorSession struct {
	Contract     *EthClosedBidSecondAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// EthClosedBidSecondAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthClosedBidSecondAuctionRaw struct {
	Contract *EthClosedBidSecondAuction // Generic contract binding to access the raw methods on
}

// EthClosedBidSecondAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthClosedBidSecondAuctionCallerRaw struct {
	Contract *EthClosedBidSecondAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// EthClosedBidSecondAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthClosedBidSecondAuctionTransactorRaw struct {
	Contract *EthClosedBidSecondAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthClosedBidSecondAuction creates a new instance of EthClosedBidSecondAuction, bound to a specific deployed contract.
func NewEthClosedBidSecondAuction(address common.Address, backend bind.ContractBackend) (*EthClosedBidSecondAuction, error) {
	contract, err := bindEthClosedBidSecondAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthClosedBidSecondAuction{EthClosedBidSecondAuctionCaller: EthClosedBidSecondAuctionCaller{contract: contract}, EthClosedBidSecondAuctionTransactor: EthClosedBidSecondAuctionTransactor{contract: contract}, EthClosedBidSecondAuctionFilterer: EthClosedBidSecondAuctionFilterer{contract: contract}}, nil
}

// NewEthClosedBidSecondAuctionCaller creates a new read-only instance of EthClosedBidSecondAuction, bound to a specific deployed contract.
func NewEthClosedBidSecondAuctionCaller(address common.Address, caller bind.ContractCaller) (*EthClosedBidSecondAuctionCaller, error) {
	contract, err := bindEthClosedBidSecondAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthClosedBidSecondAuctionCaller{contract: contract}, nil
}

// NewEthClosedBidSecondAuctionTransactor creates a new write-only instance of EthClosedBidSecondAuction, bound to a specific deployed contract.
func NewEthClosedBidSecondAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*EthClosedBidSecondAuctionTransactor, error) {
	contract, err := bindEthClosedBidSecondAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthClosedBidSecondAuctionTransactor{contract: contract}, nil
}

// NewEthClosedBidSecondAuctionFilterer creates a new log filterer instance of EthClosedBidSecondAuction, bound to a specific deployed contract.
func NewEthClosedBidSecondAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*EthClosedBidSecondAuctionFilterer, error) {
	contract, err := bindEthClosedBidSecondAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthClosedBidSecondAuctionFilterer{contract: contract}, nil
}

// bindEthClosedBidSecondAuction binds a generic wrapper to an already deployed contract.
func bindEthClosedBidSecondAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthClosedBidSecondAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthClosedBidSecondAuction.Contract.EthClosedBidSecondAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.EthClosedBidSecondAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.EthClosedBidSecondAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthClosedBidSecondAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.contract.Transact(opts, method, params...)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCaller) AssetId(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _EthClosedBidSecondAuction.contract.Call(opts, &out, "asset_id", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) AssetId(arg0 *big.Int) (string, error) {
	return _EthClosedBidSecondAuction.Contract.AssetId(&_EthClosedBidSecondAuction.CallOpts, arg0)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCallerSession) AssetId(arg0 *big.Int) (string, error) {
	return _EthClosedBidSecondAuction.Contract.AssetId(&_EthClosedBidSecondAuction.CallOpts, arg0)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCaller) CheckAverageScore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthClosedBidSecondAuction.contract.Call(opts, &out, "checkAverageScore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) CheckAverageScore() (*big.Int, error) {
	return _EthClosedBidSecondAuction.Contract.CheckAverageScore(&_EthClosedBidSecondAuction.CallOpts)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xdad23936.
//
// Solidity: function checkAverageScore() view returns(int256)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCallerSession) CheckAverageScore() (*big.Int, error) {
	return _EthClosedBidSecondAuction.Contract.CheckAverageScore(&_EthClosedBidSecondAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCaller) HighestBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EthClosedBidSecondAuction.contract.Call(opts, &out, "highestBid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _EthClosedBidSecondAuction.Contract.HighestBid(&_EthClosedBidSecondAuction.CallOpts, arg0)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCallerSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _EthClosedBidSecondAuction.Contract.HighestBid(&_EthClosedBidSecondAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCaller) HighestBidder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EthClosedBidSecondAuction.contract.Call(opts, &out, "highestBidder", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _EthClosedBidSecondAuction.Contract.HighestBidder(&_EthClosedBidSecondAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCallerSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _EthClosedBidSecondAuction.Contract.HighestBidder(&_EthClosedBidSecondAuction.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthClosedBidSecondAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) Owner() (common.Address, error) {
	return _EthClosedBidSecondAuction.Contract.Owner(&_EthClosedBidSecondAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCallerSession) Owner() (common.Address, error) {
	return _EthClosedBidSecondAuction.Contract.Owner(&_EthClosedBidSecondAuction.CallOpts)
}

// SecondHighestBid is a free data retrieval call binding the contract method 0x966ffcff.
//
// Solidity: function secondHighestBid(uint256 ) view returns(uint256)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCaller) SecondHighestBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EthClosedBidSecondAuction.contract.Call(opts, &out, "secondHighestBid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondHighestBid is a free data retrieval call binding the contract method 0x966ffcff.
//
// Solidity: function secondHighestBid(uint256 ) view returns(uint256)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) SecondHighestBid(arg0 *big.Int) (*big.Int, error) {
	return _EthClosedBidSecondAuction.Contract.SecondHighestBid(&_EthClosedBidSecondAuction.CallOpts, arg0)
}

// SecondHighestBid is a free data retrieval call binding the contract method 0x966ffcff.
//
// Solidity: function secondHighestBid(uint256 ) view returns(uint256)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCallerSession) SecondHighestBid(arg0 *big.Int) (*big.Int, error) {
	return _EthClosedBidSecondAuction.Contract.SecondHighestBid(&_EthClosedBidSecondAuction.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCaller) Status(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _EthClosedBidSecondAuction.contract.Call(opts, &out, "status", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) Status(arg0 *big.Int) (string, error) {
	return _EthClosedBidSecondAuction.Contract.Status(&_EthClosedBidSecondAuction.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCallerSession) Status(arg0 *big.Int) (string, error) {
	return _EthClosedBidSecondAuction.Contract.Status(&_EthClosedBidSecondAuction.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthClosedBidSecondAuction.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) Token() (common.Address, error) {
	return _EthClosedBidSecondAuction.Contract.Token(&_EthClosedBidSecondAuction.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionCallerSession) Token() (common.Address, error) {
	return _EthClosedBidSecondAuction.Contract.Token(&_EthClosedBidSecondAuction.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactor) Abort(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.contract.Transact(opts, "abort", auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.Abort(&_EthClosedBidSecondAuction.TransactOpts, auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactorSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.Abort(&_EthClosedBidSecondAuction.TransactOpts, auctionId, jsonString)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactor) Bid(opts *bind.TransactOpts, auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.contract.Transact(opts, "bid", auctionId, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) Bid(auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.Bid(&_EthClosedBidSecondAuction.TransactOpts, auctionId, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactorSession) Bid(auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.Bid(&_EthClosedBidSecondAuction.TransactOpts, auctionId, bidHash)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x88d3c98e.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform, uint256 secondHighestPrice) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactor) CloseAuction(opts *bind.TransactOpts, auctionId *big.Int, not_winner_platform bool, secondHighestPrice *big.Int) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.contract.Transact(opts, "closeAuction", auctionId, not_winner_platform, secondHighestPrice)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x88d3c98e.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform, uint256 secondHighestPrice) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) CloseAuction(auctionId *big.Int, not_winner_platform bool, secondHighestPrice *big.Int) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.CloseAuction(&_EthClosedBidSecondAuction.TransactOpts, auctionId, not_winner_platform, secondHighestPrice)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x88d3c98e.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform, uint256 secondHighestPrice) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactorSession) CloseAuction(auctionId *big.Int, not_winner_platform bool, secondHighestPrice *big.Int) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.CloseAuction(&_EthClosedBidSecondAuction.TransactOpts, auctionId, not_winner_platform, secondHighestPrice)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactor) Commit(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.contract.Transact(opts, "commit", auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.Commit(&_EthClosedBidSecondAuction.TransactOpts, auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactorSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.Commit(&_EthClosedBidSecondAuction.TransactOpts, auctionId, jsonString)
}

// Create is a paid mutator transaction binding the contract method 0xb6a46b3b.
//
// Solidity: function create(string _asset_id) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactor) Create(opts *bind.TransactOpts, _asset_id string) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.contract.Transact(opts, "create", _asset_id)
}

// Create is a paid mutator transaction binding the contract method 0xb6a46b3b.
//
// Solidity: function create(string _asset_id) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) Create(_asset_id string) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.Create(&_EthClosedBidSecondAuction.TransactOpts, _asset_id)
}

// Create is a paid mutator transaction binding the contract method 0xb6a46b3b.
//
// Solidity: function create(string _asset_id) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactorSession) Create(_asset_id string) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.Create(&_EthClosedBidSecondAuction.TransactOpts, _asset_id)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactor) ProvideFeedback(opts *bind.TransactOpts, auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.contract.Transact(opts, "provide_feedback", auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.ProvideFeedback(&_EthClosedBidSecondAuction.TransactOpts, auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x274377c0.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, bytes32 _feedback) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactorSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback [32]byte) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.ProvideFeedback(&_EthClosedBidSecondAuction.TransactOpts, auctionId, _score, _feedback)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactor) Reveal(opts *bind.TransactOpts, auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.contract.Transact(opts, "reveal", auctionId, bidAmount)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) Reveal(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.Reveal(&_EthClosedBidSecondAuction.TransactOpts, auctionId, bidAmount)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactorSession) Reveal(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.Reveal(&_EthClosedBidSecondAuction.TransactOpts, auctionId, bidAmount)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactor) RevealAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.contract.Transact(opts, "revealAuction", auctionId)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) RevealAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.RevealAuction(&_EthClosedBidSecondAuction.TransactOpts, auctionId)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactorSession) RevealAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.RevealAuction(&_EthClosedBidSecondAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactor) Withdraw(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.contract.Transact(opts, "withdraw", auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.Withdraw(&_EthClosedBidSecondAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionTransactorSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _EthClosedBidSecondAuction.Contract.Withdraw(&_EthClosedBidSecondAuction.TransactOpts, auctionId)
}

// EthClosedBidSecondAuctionAwaitResponseIterator is returned from FilterAwaitResponse and is used to iterate over the raw logs and unpacked data for AwaitResponse events raised by the EthClosedBidSecondAuction contract.
type EthClosedBidSecondAuctionAwaitResponseIterator struct {
	Event *EthClosedBidSecondAuctionAwaitResponse // Event containing the contract specifics and raw log

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
func (it *EthClosedBidSecondAuctionAwaitResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidSecondAuctionAwaitResponse)
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
		it.Event = new(EthClosedBidSecondAuctionAwaitResponse)
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
func (it *EthClosedBidSecondAuctionAwaitResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidSecondAuctionAwaitResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidSecondAuctionAwaitResponse represents a AwaitResponse event raised by the EthClosedBidSecondAuction contract.
type EthClosedBidSecondAuctionAwaitResponse struct {
	Auction *big.Int
	Winner  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAwaitResponse is a free log retrieval operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) FilterAwaitResponse(opts *bind.FilterOpts) (*EthClosedBidSecondAuctionAwaitResponseIterator, error) {

	logs, sub, err := _EthClosedBidSecondAuction.contract.FilterLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidSecondAuctionAwaitResponseIterator{contract: _EthClosedBidSecondAuction.contract, event: "AwaitResponse", logs: logs, sub: sub}, nil
}

// WatchAwaitResponse is a free log subscription operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) WatchAwaitResponse(opts *bind.WatchOpts, sink chan<- *EthClosedBidSecondAuctionAwaitResponse) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidSecondAuction.contract.WatchLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidSecondAuctionAwaitResponse)
				if err := _EthClosedBidSecondAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
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
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) ParseAwaitResponse(log types.Log) (*EthClosedBidSecondAuctionAwaitResponse, error) {
	event := new(EthClosedBidSecondAuctionAwaitResponse)
	if err := _EthClosedBidSecondAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidSecondAuctionDecisionMadeIterator is returned from FilterDecisionMade and is used to iterate over the raw logs and unpacked data for DecisionMade events raised by the EthClosedBidSecondAuction contract.
type EthClosedBidSecondAuctionDecisionMadeIterator struct {
	Event *EthClosedBidSecondAuctionDecisionMade // Event containing the contract specifics and raw log

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
func (it *EthClosedBidSecondAuctionDecisionMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidSecondAuctionDecisionMade)
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
		it.Event = new(EthClosedBidSecondAuctionDecisionMade)
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
func (it *EthClosedBidSecondAuctionDecisionMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidSecondAuctionDecisionMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidSecondAuctionDecisionMade represents a DecisionMade event raised by the EthClosedBidSecondAuction contract.
type EthClosedBidSecondAuctionDecisionMade struct {
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
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) FilterDecisionMade(opts *bind.FilterOpts) (*EthClosedBidSecondAuctionDecisionMadeIterator, error) {

	logs, sub, err := _EthClosedBidSecondAuction.contract.FilterLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidSecondAuctionDecisionMadeIterator{contract: _EthClosedBidSecondAuction.contract, event: "DecisionMade", logs: logs, sub: sub}, nil
}

// WatchDecisionMade is a free log subscription operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auction, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) WatchDecisionMade(opts *bind.WatchOpts, sink chan<- *EthClosedBidSecondAuctionDecisionMade) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidSecondAuction.contract.WatchLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidSecondAuctionDecisionMade)
				if err := _EthClosedBidSecondAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
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
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) ParseDecisionMade(log types.Log) (*EthClosedBidSecondAuctionDecisionMade, error) {
	event := new(EthClosedBidSecondAuctionDecisionMade)
	if err := _EthClosedBidSecondAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidSecondAuctionHighestBidIncreasedIterator is returned from FilterHighestBidIncreased and is used to iterate over the raw logs and unpacked data for HighestBidIncreased events raised by the EthClosedBidSecondAuction contract.
type EthClosedBidSecondAuctionHighestBidIncreasedIterator struct {
	Event *EthClosedBidSecondAuctionHighestBidIncreased // Event containing the contract specifics and raw log

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
func (it *EthClosedBidSecondAuctionHighestBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidSecondAuctionHighestBidIncreased)
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
		it.Event = new(EthClosedBidSecondAuctionHighestBidIncreased)
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
func (it *EthClosedBidSecondAuctionHighestBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidSecondAuctionHighestBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidSecondAuctionHighestBidIncreased represents a HighestBidIncreased event raised by the EthClosedBidSecondAuction contract.
type EthClosedBidSecondAuctionHighestBidIncreased struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*EthClosedBidSecondAuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _EthClosedBidSecondAuction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidSecondAuctionHighestBidIncreasedIterator{contract: _EthClosedBidSecondAuction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) WatchHighestBidIncreased(opts *bind.WatchOpts, sink chan<- *EthClosedBidSecondAuctionHighestBidIncreased) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidSecondAuction.contract.WatchLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidSecondAuctionHighestBidIncreased)
				if err := _EthClosedBidSecondAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
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
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) ParseHighestBidIncreased(log types.Log) (*EthClosedBidSecondAuctionHighestBidIncreased, error) {
	event := new(EthClosedBidSecondAuctionHighestBidIncreased)
	if err := _EthClosedBidSecondAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidSecondAuctionNewBidHashIterator is returned from FilterNewBidHash and is used to iterate over the raw logs and unpacked data for NewBidHash events raised by the EthClosedBidSecondAuction contract.
type EthClosedBidSecondAuctionNewBidHashIterator struct {
	Event *EthClosedBidSecondAuctionNewBidHash // Event containing the contract specifics and raw log

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
func (it *EthClosedBidSecondAuctionNewBidHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidSecondAuctionNewBidHash)
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
		it.Event = new(EthClosedBidSecondAuctionNewBidHash)
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
func (it *EthClosedBidSecondAuctionNewBidHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidSecondAuctionNewBidHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidSecondAuctionNewBidHash represents a NewBidHash event raised by the EthClosedBidSecondAuction contract.
type EthClosedBidSecondAuctionNewBidHash struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	BidHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBidHash is a free log retrieval operation binding the contract event 0x6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f.
//
// Solidity: event NewBidHash(uint256 auction, string id, address bidder, bytes32 bidHash)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) FilterNewBidHash(opts *bind.FilterOpts) (*EthClosedBidSecondAuctionNewBidHashIterator, error) {

	logs, sub, err := _EthClosedBidSecondAuction.contract.FilterLogs(opts, "NewBidHash")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidSecondAuctionNewBidHashIterator{contract: _EthClosedBidSecondAuction.contract, event: "NewBidHash", logs: logs, sub: sub}, nil
}

// WatchNewBidHash is a free log subscription operation binding the contract event 0x6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f.
//
// Solidity: event NewBidHash(uint256 auction, string id, address bidder, bytes32 bidHash)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) WatchNewBidHash(opts *bind.WatchOpts, sink chan<- *EthClosedBidSecondAuctionNewBidHash) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidSecondAuction.contract.WatchLogs(opts, "NewBidHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidSecondAuctionNewBidHash)
				if err := _EthClosedBidSecondAuction.contract.UnpackLog(event, "NewBidHash", log); err != nil {
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
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) ParseNewBidHash(log types.Log) (*EthClosedBidSecondAuctionNewBidHash, error) {
	event := new(EthClosedBidSecondAuctionNewBidHash)
	if err := _EthClosedBidSecondAuction.contract.UnpackLog(event, "NewBidHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidSecondAuctionRateAuctionIterator is returned from FilterRateAuction and is used to iterate over the raw logs and unpacked data for RateAuction events raised by the EthClosedBidSecondAuction contract.
type EthClosedBidSecondAuctionRateAuctionIterator struct {
	Event *EthClosedBidSecondAuctionRateAuction // Event containing the contract specifics and raw log

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
func (it *EthClosedBidSecondAuctionRateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidSecondAuctionRateAuction)
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
		it.Event = new(EthClosedBidSecondAuctionRateAuction)
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
func (it *EthClosedBidSecondAuctionRateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidSecondAuctionRateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidSecondAuctionRateAuction represents a RateAuction event raised by the EthClosedBidSecondAuction contract.
type EthClosedBidSecondAuctionRateAuction struct {
	Auction *big.Int
	Id      string
	Rating  *big.Int
	Review  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRateAuction is a free log retrieval operation binding the contract event 0x10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e79.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, bytes32 review)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) FilterRateAuction(opts *bind.FilterOpts) (*EthClosedBidSecondAuctionRateAuctionIterator, error) {

	logs, sub, err := _EthClosedBidSecondAuction.contract.FilterLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidSecondAuctionRateAuctionIterator{contract: _EthClosedBidSecondAuction.contract, event: "RateAuction", logs: logs, sub: sub}, nil
}

// WatchRateAuction is a free log subscription operation binding the contract event 0x10cb752cc311a330a3517759515129336e43ce15717e3cb477ad20ac87c85e79.
//
// Solidity: event RateAuction(uint256 auction, string id, int256 rating, bytes32 review)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) WatchRateAuction(opts *bind.WatchOpts, sink chan<- *EthClosedBidSecondAuctionRateAuction) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidSecondAuction.contract.WatchLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidSecondAuctionRateAuction)
				if err := _EthClosedBidSecondAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
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
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) ParseRateAuction(log types.Log) (*EthClosedBidSecondAuctionRateAuction, error) {
	event := new(EthClosedBidSecondAuctionRateAuction)
	if err := _EthClosedBidSecondAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidSecondAuctionRevealAuctionIterator is returned from FilterRevealAuction and is used to iterate over the raw logs and unpacked data for RevealAuction events raised by the EthClosedBidSecondAuction contract.
type EthClosedBidSecondAuctionRevealAuctionIterator struct {
	Event *EthClosedBidSecondAuctionRevealAuction // Event containing the contract specifics and raw log

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
func (it *EthClosedBidSecondAuctionRevealAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidSecondAuctionRevealAuction)
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
		it.Event = new(EthClosedBidSecondAuctionRevealAuction)
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
func (it *EthClosedBidSecondAuctionRevealAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidSecondAuctionRevealAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidSecondAuctionRevealAuction represents a RevealAuction event raised by the EthClosedBidSecondAuction contract.
type EthClosedBidSecondAuctionRevealAuction struct {
	Auction *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevealAuction is a free log retrieval operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auction)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) FilterRevealAuction(opts *bind.FilterOpts) (*EthClosedBidSecondAuctionRevealAuctionIterator, error) {

	logs, sub, err := _EthClosedBidSecondAuction.contract.FilterLogs(opts, "RevealAuction")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidSecondAuctionRevealAuctionIterator{contract: _EthClosedBidSecondAuction.contract, event: "RevealAuction", logs: logs, sub: sub}, nil
}

// WatchRevealAuction is a free log subscription operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auction)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) WatchRevealAuction(opts *bind.WatchOpts, sink chan<- *EthClosedBidSecondAuctionRevealAuction) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidSecondAuction.contract.WatchLogs(opts, "RevealAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidSecondAuctionRevealAuction)
				if err := _EthClosedBidSecondAuction.contract.UnpackLog(event, "RevealAuction", log); err != nil {
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
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) ParseRevealAuction(log types.Log) (*EthClosedBidSecondAuctionRevealAuction, error) {
	event := new(EthClosedBidSecondAuctionRevealAuction)
	if err := _EthClosedBidSecondAuction.contract.UnpackLog(event, "RevealAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClosedBidSecondAuctionWithdrawBidIterator is returned from FilterWithdrawBid and is used to iterate over the raw logs and unpacked data for WithdrawBid events raised by the EthClosedBidSecondAuction contract.
type EthClosedBidSecondAuctionWithdrawBidIterator struct {
	Event *EthClosedBidSecondAuctionWithdrawBid // Event containing the contract specifics and raw log

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
func (it *EthClosedBidSecondAuctionWithdrawBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClosedBidSecondAuctionWithdrawBid)
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
		it.Event = new(EthClosedBidSecondAuctionWithdrawBid)
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
func (it *EthClosedBidSecondAuctionWithdrawBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClosedBidSecondAuctionWithdrawBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClosedBidSecondAuctionWithdrawBid represents a WithdrawBid event raised by the EthClosedBidSecondAuction contract.
type EthClosedBidSecondAuctionWithdrawBid struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBid is a free log retrieval operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) FilterWithdrawBid(opts *bind.FilterOpts) (*EthClosedBidSecondAuctionWithdrawBidIterator, error) {

	logs, sub, err := _EthClosedBidSecondAuction.contract.FilterLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return &EthClosedBidSecondAuctionWithdrawBidIterator{contract: _EthClosedBidSecondAuction.contract, event: "WithdrawBid", logs: logs, sub: sub}, nil
}

// WatchWithdrawBid is a free log subscription operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) WatchWithdrawBid(opts *bind.WatchOpts, sink chan<- *EthClosedBidSecondAuctionWithdrawBid) (event.Subscription, error) {

	logs, sub, err := _EthClosedBidSecondAuction.contract.WatchLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClosedBidSecondAuctionWithdrawBid)
				if err := _EthClosedBidSecondAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
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
func (_EthClosedBidSecondAuction *EthClosedBidSecondAuctionFilterer) ParseWithdrawBid(log types.Log) (*EthClosedBidSecondAuctionWithdrawBid, error) {
	event := new(EthClosedBidSecondAuctionWithdrawBid)
	if err := _EthClosedBidSecondAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
