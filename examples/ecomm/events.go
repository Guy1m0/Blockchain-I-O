package ecomm

const (
	OnInitializedLending = "ecomm.initialize"
	OnLoanSuccessful     = "ecomm.successful"
	OnLoanFail           = "ecomm.fail"
)

var (
	SignedAuctionResultEvent = "auction.signed_result"
	AuctionEndingEvent       = "auction.auction_ending"
)

type SignedAuctionResult struct {
	AuctionResult
	Signature []byte
}
