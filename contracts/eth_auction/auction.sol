// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

interface IERC20 {
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    function burn(address usr, uint wad) external;
}

contract Auction {

    // Current state of the auction.
    address public highestBidder;
    uint public highestBid;

    // Allowed withdrawals of previous bids
    mapping(address => uint) pendingReturns;

    // Set to true at the end, disallows any change.
    // By default initialized to `false`.
    bool public ended;
    string public status;
    address public owner;

    string public asset_id;

    // Events that will be emitted on changes.
    event HighestBidIncreased(address bidder, uint amount);
    event WithdarwBid(address bidder, uint amount);
    event DecisionMade(address winner, uint amount, string id);
    event WaitResponse(address winner);

    IERC20 public token;

    constructor(IERC20 _token, string memory _asset_id) {
        //beneficiary = payable(msg.sender);
        token = _token;
        status = "running";
        owner = msg.sender;
        asset_id = _asset_id;
    }

    function bid(uint bidAmount) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("running")), "Contract not in RUNNING status");

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
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("ended")), "Contract not in ENDED status");

        uint amount = pendingReturns[msg.sender];
        if (amount > 0) {
            pendingReturns[msg.sender] = 0;

            if (!token.transferFrom(address(this), msg.sender, amount)) {
                pendingReturns[msg.sender] = amount;
                return false;
            }
        }
        emit WithdarwBid(msg.sender, amount);
        return true;
    }

    function endAuction(bool not_winner_platform) public {
        require(msg.sender == owner, "Only owner can change contract's status");
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("running")), "Contract not in RUNNING status");

        status = "ending";
        if (not_winner_platform || highestBid == 0){
            abort();
            return;
        }

        emit WaitResponse(highestBidder);
    }

    function abort() public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("ending")), "Contract not in ENDING status");

        require(msg.sender == owner || msg.sender == highestBidder, "Not authorized access!");

        status = "ended";

        pendingReturns[highestBidder] = highestBid;
        highestBidder = address(0);
        highestBid = 0;

        emit DecisionMade(highestBidder, highestBid, asset_id);

    }

    function proceed() public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("ending")), "Contract not in ENDING status");
        // For testing only
        require(msg.sender == highestBidder, "Not authorized access!");
        
        status = "ended";

        token.burn(address(this), highestBid);
        emit DecisionMade(highestBidder, highestBid, asset_id);

    }
}
