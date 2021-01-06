// SPDX-License-Identifier: MIT
pragma solidity >=0.5.0 <0.8.0;

contract Booleans {
    
// DECLARATION
    bool red = true;
    bool round; // default value is 'false'

// OPERATIONS

    // conditional
    if (red) {
        // conditional code
    }

    // logical negation
    if (!red) {
        // conditional code
    }

    // logical conjunction (AND)
    if (red && round) {
        // conditional code
    }

    // logical conjunction (OR)
    if (red || round) {
        // conditional code
    }

    // equality
    if (red == round) {
        // conditional code
    }

    // inequality
    if (red != round) {
        // conditional code
    }
}