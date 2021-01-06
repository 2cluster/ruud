// SPDX-License-Identifier: MIT
pragma solidity >=0.5.0 <0.8.0;

contract Arrays {
    
    // fixed size
    bytes1 arr = "a";
    bytes32 arr2 = "this is a 32 bytes array"; 
    bytes32 arr3; // default value is ""


    // dynamic sized, is a reference type! 

    bytes1[2] arr4;
    uint8[] arr5;

    arr4[0] = arr;

}