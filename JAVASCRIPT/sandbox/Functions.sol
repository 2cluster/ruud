// SPDX-License-Identifier: MIT
pragma solidity >=0.5.0 <0.8.0;

contract Functions {
    
    function setAge(uint age) public {
        
    }

    function getNameById(uint id) view external returns (bytes32 name) {
        // do something

        return name;
    }

    // function name(type name) view public returns (type name) {
        
    // }




// function visibilities

private // only calls from inside the contract
internal // only from inside the contract and inheratance contracts
external // only from outside of the contract
public // available from inside and outside of the contract

}