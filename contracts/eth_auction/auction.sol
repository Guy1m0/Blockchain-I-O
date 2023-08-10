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

    string public status;
    address public owner;

    string public asset_id;

    string private feedback;
    int private score;

    // Events that will be emitted on changes.
    event HighestBidIncreased(address bidder, uint amount);
    event WithdarwBid(address bidder, uint amount);
    event DecisionMade(address winner, uint amount, string id, bool prcd, string jsonString);
    event WaitResponse(address winner);
    event RateAuction(string id, int rating, string review);

    IERC20 public token;

    constructor(IERC20 _token, string memory _asset_id) {
        //beneficiary = payable(msg.sender);
        token = _token;
        status = "open";
        owner = msg.sender;
        asset_id = _asset_id;
    }

    function bid(uint bidAmount) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("open")), "Contract not in OPEN status");

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
        require(keccak256(abi.encodePacked(status)) != keccak256(abi.encodePacked("open")), "Contract in OPEN status");

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

    function closeAuction(bool not_winner_platform) public {
        require(msg.sender == owner, "Only owner can change contract's status");
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("open")), "Contract not in OPEN status");

        status = "ending";
        if (not_winner_platform || highestBid == 0){
            status = "closed";

            pendingReturns[highestBidder] = highestBid;
            highestBidder = address(0);
            highestBid = 0;

            return;
        }

        emit WaitResponse(highestBidder);
    }

    function abort(string memory jsonString) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("ending")), "Contract not in ENDING status");

        require(msg.sender == highestBidder, "Not authorized access!");

        status = "closing";
        emit DecisionMade(highestBidder, highestBid, asset_id, false, jsonString);

        pendingReturns[highestBidder] = highestBid;
        highestBidder = address(0);
        highestBid = 0;

        
    }

    function proceed(string memory jsonString) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("ending")), "Contract not in ENDING status");
        // For testing only
        require(msg.sender == highestBidder, "Not authorized access!");
        
        status = "closing";

        token.burn(address(this), highestBid);
        emit DecisionMade(highestBidder, highestBid, asset_id, true, jsonString);

    }

    function provide_feedback(int _score, string memory _feedback) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("closing")), "Contract not in CLOSING status");
        // For testing only
        require(msg.sender == highestBidder, "Not authorized access!");

        feedback = _feedback;
        score = _score;

        emit RateAuction(asset_id, _score, _feedback);
        status = "closed";

    }

    function checkFeedback() public view returns (int, string memory){
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("closed")), "Contract not in CLOSED status");

        return (score, feedback);
    }
}
