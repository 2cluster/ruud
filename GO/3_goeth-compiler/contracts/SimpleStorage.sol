// SPDX-License-Identifier: MIT
pragma solidity >=0.4.16 <0.8.0;

contract SimpleStorage {
    
    uint8 data;
    event DataSet(uint256 indexed oldData, uint256 indexed newData);

    function set(uint8 x) public {
        emit DataSet(data, x);
        data = x;

    }

    function get() public view returns (uint8) {
        return data;
    }
}
