In this folder you find:

- the Counter.sol contract
- the eventListener.js script

The contract emits events when endusers add or subtract the stored counter value.
It also uses SafeMath to avoid overflow and instead reverts.  

The script has 2 functions:
- Live function listener called: "eventQuery"
- Past event caller called: "getPastEvents"