package asset

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

const (
	KeyAssets        = "assets"
	KeyAuctions      = "auctions"
	KeyLastAuctionID = "lastAuction"
)

func (cc *SmartContract) AddAsset(
	ctx contractapi.TransactionContextInterface, id, owner string,
) error {
	existing, err := ctx.GetStub().GetState(cc.makeAssetKey(id))
	if err != nil {
		return fmt.Errorf("unable to interact with worldstate: %v", err)
	}

	if existing != nil {
		return fmt.Errorf("asset with ID %s already exists", id)
	}

	asset := Asset{
		ID:    id,
		Owner: owner,
	}

	err = cc.setAsset(ctx, &asset)
	if err != nil {
		return err
	}

	// Emit an event when an asset is added
	eventPayload := "Asset added: " + id
	err = ctx.GetStub().SetEvent("AddAsset", []byte(eventPayload))
	if err != nil {
		return fmt.Errorf("error setting event: %v", err)
	}
	return nil
}

func (cc *SmartContract) StartAuction(
	ctx contractapi.TransactionContextInterface, argjson string,
) error {
	var args StartAuctionArgs
	err := json.Unmarshal([]byte(argjson), &args)
	if err != nil {
		return err
	}

	asset, err := cc.GetAsset(ctx, args.AssetID)
	if err != nil {
		return err
	}
	if asset.PendingAuctionID > 0 {
		return fmt.Errorf("pending auction on asset")
	}

	lastID, err := cc.GetLastAuctionID(ctx)
	if err != nil {
		return err
	}
	auction := Auction{
		ID:         lastID + 1,
		AssetID:    args.AssetID,
		EthAddr:    args.EthAddr,
		QuorumAddr: args.QuorumAddr,
		Status:     "Started",
	}
	err = cc.setAuction(ctx, &auction)
	if err != nil {
		return err
	}
	err = cc.setLastAuctionID(ctx, auction.ID)
	if err != nil {
		return err
	}

	asset.PendingAuctionID = auction.ID

	// Emit an event when an auction is started
	eventPayload := fmt.Sprintf("Auction start: %d", auction.ID)
	err = ctx.GetStub().SetEvent("StartAuction", []byte(eventPayload))
	if err != nil {
		return fmt.Errorf("error setting event: %v", err)
	}

	return cc.setAsset(ctx, asset)
}

func (cc *SmartContract) EndAuction(
	ctx contractapi.TransactionContextInterface, assetID string,
) error {
	asset, err := cc.GetAsset(ctx, assetID)
	if err != nil {
		return err
	}
	auction, err := cc.GetAuction(ctx, asset.PendingAuctionID)
	if err != nil {
		return err
	}
	auction.Status = "Ending"

	// Emit an event when an auction is started
	eventPayload := fmt.Sprintf("Auction end: %d", auction.ID)
	err = ctx.GetStub().SetEvent("EndAuction", []byte(eventPayload))
	if err != nil {
		return fmt.Errorf("error setting event: %v", err)
	}

	return cc.setAuction(ctx, auction)
}

func (cc *SmartContract) FinalizeAuction(
	ctx contractapi.TransactionContextInterface, argjson string,
) error {
	var args FinalizeAuctionArgs
	err := json.Unmarshal([]byte(argjson), &args)
	if err != nil {
		return err
	}

	auction, err := cc.GetAuction(ctx, args.AuctionID)
	if err != nil {
		return err
	}

	addrs, min := cc.ethAddrs()
	if !cc.verifyAuctionResult(addrs, min, args.EthResult) {
		return fmt.Errorf("invalid ethereum result")
	}
	addrs, min = cc.quorumAddrs()
	if !cc.verifyAuctionResult(addrs, min, args.EthResult) {
		return fmt.Errorf("invalid quorum result")
	}

	if args.EthResult.AuctionAddr != auction.EthAddr {
		return fmt.Errorf("invalid ethereum address")
	}
	if args.QuorumResult.AuctionAddr != auction.QuorumAddr {
		return fmt.Errorf("invalid quorum address")
	}

	if args.EthResult.HighestBid >= args.QuorumResult.HighestBid {
		auction.HighestBid = args.EthResult.HighestBid
		auction.HighestBidder = args.EthResult.HighestBidder
		auction.HighestBidPlatform = "ethereum"
	} else {
		auction.HighestBid = args.QuorumResult.HighestBid
		auction.HighestBidder = args.QuorumResult.HighestBidder
		auction.HighestBidPlatform = "quorum"
	}

	auction.Status = "Ended"
	err = cc.setAuction(ctx, auction)
	if err != nil {
		return err
	}

	asset, err := cc.GetAsset(ctx, auction.AssetID)
	if err != nil {
		return err
	}

	asset.Owner = auction.HighestBidder
	asset.PendingAuctionID = 0
	err = cc.setAsset(ctx, asset)
	if err != nil {
		return err
	}
	return nil
}

// in other words, it only accepts bidder who uses following keys
// So if want to improve scalability, need to modify this
func (cc *SmartContract) ethAddrs() (addrs []string, min int) {
	return []string{
		"17dc6ca2e1c84ae4107975a48dfd05831b8addff", // key 0
		"ac5580ad28a3c0e044a52541785bfd34c753d3bf", // key 1
		"d0a73fe9d44184e9f1264ce2097064212e67ebfe", // key 2
	}, 2
}

func (cc *SmartContract) quorumAddrs() (addrs []string, min int) {
	return []string{
		"17dc6ca2e1c84ae4107975a48dfd05831b8addff",
		"ac5580ad28a3c0e044a52541785bfd34c753d3bf",
		"d0a73fe9d44184e9f1264ce2097064212e67ebfe",
	}, 2
}

// can add some mech to check if bidder has DID creditional
func (cc *SmartContract) verifyAuctionResult(
	trustedAddrs []string, majority int, result CrossChainAuctionResult,
) bool {
	if len(result.Signatures) < majority {
		return false
	}
	addrs := make([]common.Address, 0, len(trustedAddrs))
	for _, addr := range trustedAddrs {
		addrs = append(addrs, common.HexToAddress(addr))
	}
	return VerifyKnownSignatures(result.Hash(), result.Signatures, addrs)
}

func (cc *SmartContract) GetAsset(
	ctx contractapi.TransactionContextInterface, assetID string,
) (*Asset, error) {
	var asset Asset
	b, err := ctx.GetStub().GetState(cc.makeAssetKey(assetID))
	if err != nil {
		return nil, err
	}
	if b == nil {
		return nil, fmt.Errorf("asset not found")
	}
	err = json.Unmarshal(b, &asset)
	return &asset, err
}

func (cc *SmartContract) GetAuction(
	ctx contractapi.TransactionContextInterface, auctionID int,
) (*Auction, error) {
	b, err := ctx.GetStub().GetState(cc.makeAuctionKey(auctionID))
	if err != nil {
		return nil, err
	}
	if b == nil {
		return nil, fmt.Errorf("auction not found")
	}
	var auction Auction
	err = json.Unmarshal(b, &auction)
	return &auction, err
}

func (cc *SmartContract) GetLastAuctionID(
	ctx contractapi.TransactionContextInterface,
) (int, error) {
	b, err := ctx.GetStub().GetState(KeyLastAuctionID)
	if err != nil {
		return 0, err
	}
	var count int
	json.Unmarshal(b, &count)
	return count, nil
}

func (cc *SmartContract) setAsset(
	ctx contractapi.TransactionContextInterface, asset *Asset,
) error {
	b, _ := json.Marshal(asset)
	err := ctx.GetStub().PutState(cc.makeAssetKey(asset.ID), b)
	if err != nil {
		return fmt.Errorf("set asset error: %v", err)
	}
	return nil
}

func (cc *SmartContract) setAuction(
	ctx contractapi.TransactionContextInterface, auction *Auction,
) error {
	b, _ := json.Marshal(auction)
	err := ctx.GetStub().PutState(cc.makeAuctionKey(auction.ID), b)
	if err != nil {
		return fmt.Errorf("set auction error: %v", err)
	}
	return nil
}

func (cc *SmartContract) setLastAuctionID(
	ctx contractapi.TransactionContextInterface, id int,
) error {
	b, _ := json.Marshal(id)
	return ctx.GetStub().PutState(KeyLastAuctionID, b)
}

func (cc *SmartContract) makeAssetKey(assetID string) string {
	return fmt.Sprintf("%s_%s", KeyAssets, assetID)
}

func (cc *SmartContract) makeAuctionKey(auctionID int) string {
	return fmt.Sprintf("%s_%d", KeyAuctions, auctionID)
}
