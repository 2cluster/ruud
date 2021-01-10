This folder contains 2 contracts that interact with the chainlink network

- aggregator.sol
    This contract sets up a aggregator that automatically retrieves the latest ETH/USD price. 
    It needs to interact with another contract located on the Kovan network that is hardcored in the contract. 
    This is contract requires LINK tokens on the deployed contract's address. It does not require users to invoke
    Without cost users can retrieve the latest price by querying the 'getLatestPrice' function


- random.sol
    This contract uses Chainlink VRF and this is a verifiable source of randomness for smart contracts. 
    This contract also needs a balance to work, when users invoke the 'getRandomNumber' function, 
    the contract will create a callback to the Chainlink oracle
