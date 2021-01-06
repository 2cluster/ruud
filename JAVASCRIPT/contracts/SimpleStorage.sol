// SPDX-License-Identifier: UNLICENSED
pragma solidity >= 0.4.16 <0.8.0;

contract SimpleStorage {
    uint data;

    function set(uint x) public {
        data = x;
    }

    function get() public view returns (uint) {
        return data;
    }


}