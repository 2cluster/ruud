//SPDX-License-Identifier: Unlicense
pragma solidity ^0.6.0;

import "./SafeMath.sol";

contract Counter {

    using SafeMath for uint256;
    
    uint256 public counter = 0;
    
    event update(int amount, uint256 counter);
    
     function increment() public {
        counter = counter.add(1);
        emit update(1, counter);
    }
    
    function decrement() public {
        counter = counter.sub(1);
        emit update(-1, counter);
    }
    
}