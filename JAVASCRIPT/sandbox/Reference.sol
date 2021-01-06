// SPDX-License-Identifier: MIT

pragma solidity >=0.5.0 <0.8.0;

contract Types {
    
    function testValue() public pure returns (uint) {
        uint x;
        uint y;

        y = x;
        x = 5;

        return y;
    }

    function testReference() public pure returns (uint) {
        uint[1]  x;
        uint[1]  y;

        x = y;
        x[0] = 5;

        return y[0];
    }

}