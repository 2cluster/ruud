//SPDX-License-Identifier: Unlicense
pragma solidity ^0.6.0;
import "github.com/smartcontractkit/chainlink/evm-contracts/src/v0.6/VRFConsumerBase.sol";

/**
 * Network: Kovan
 */

contract RandomNumberConsumer is VRFConsumerBase {
    bytes32 internal keyHash;
    uint256 internal fee;
    uint256 public result;
    
    constructor() 
        VRFConsumerBase(
            0xdD3782915140c8f3b190B5D67eAc6dc5760C46E9, // VRF Coordinator
            0xa36085F69e2889c224210F603D836748e7dC0088  // LINK Token 
        ) public
    {
        keyHash = 0x6c3699283bda56ad74f6b855546325b68d482e983852a7a82979cc4807b641f4; // public key against which randomness is generated.
        fee = 0.1 * 10 ** 18; // 0.1 LINK fee, the contract has 18 decimals
    }

    function CheckBalance() public view returns (uint) {
      return LINK.balanceOf(address(this));
    }

    function fulfillRandomness(bytes32 requestId, uint256 randomness) internal override {
        requestId;
        result = randomness;
    }

    function getRandomNumber(uint256 userProvidedSeed) public returns (bytes32 requestId) {
        require (CheckBalance() >= fee,"Not enough LINK Tokens in contract");
        return requestRandomness(keyHash, fee, userProvidedSeed);
    }
}