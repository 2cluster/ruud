// SPDX-License-Identifier: MIT
pragma solidity >=0.5.0 <0.8.0;

contract Enums {
    
    // DECLARATION
    enum SizeChoice {Small, Medium, Large} //default value is the first one
    SizeChoice size = SizeChoice.Large;

    // OPERATION

    size = SizeChoice.Medium;


}