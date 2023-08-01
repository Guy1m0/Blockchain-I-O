// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

interface IERC20 {
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    function burn(uint256 amount) external;
}

contract Auction {
    //address payable public beneficiary;

    // Current state of the auction.
    address public highestBidder;
    uint public highestBid;

    // Allowed withdrawals of previous bids
    mapping(address => uint) pendingReturns;

    // Set to true at the end, disallows any change.
    // By default initialized to `false`.
    bool public ended;

    // Events that will be emitted on changes.
    event HighestBidIncreased(address bidder, uint amount);
    event AuctionEnded(address winner, uint amount);
    event Transfer(address indexed src, address indexed dst, uint wad);

    IERC20 public token;

    constructor(IERC20 _token) {
        //beneficiary = payable(msg.sender);
        token = _token;
    }

    function bid(uint bidAmount) public {
        // Check that the bid is higher than the current highest bid
        require(bidAmount > highestBid, "There already is a higher bid.");

        // Attempt to transfer the tokens from the bidder to the contract
        bool transferSuccessful = token.transferFrom(msg.sender, address(this), bidAmount);
        
        // Check that the token transfer was successful
        require(transferSuccessful, "Token transfer failed.");

        // If there was a previous bid, allow the previous highest bidder to withdraw their bid
        if (highestBid != 0) {
            pendingReturns[highestBidder] += highestBid;
        }
        
        // Update the highest bid and highest bidder
        highestBidder = msg.sender;
        highestBid = bidAmount;
        emit HighestBidIncreased(msg.sender, bidAmount);
    }


    // Withdraw a bid that was overbid.
    function withdraw() public returns (bool) {
        uint amount = pendingReturns[msg.sender];
        if (amount > 0) {
            pendingReturns[msg.sender] = 0;

            if (!token.transferFrom(address(this), msg.sender, amount)) {
                pendingReturns[msg.sender] = amount;
                return false;
            }
        }
        return true;
    }

    function endAuction() public {
        // 1. Conditions
        require(!ended, "auctionEnd has already been called.");

        // 2. Effects
        ended = true;
        emit AuctionEnded(highestBidder, highestBid);

        // 3. Interaction
        token.burn(highestBid);
    }
}
