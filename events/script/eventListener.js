Web3 = require("web3");

const web3 = new Web3('wss://rinkeby.infura.io/ws/v3/634d2ee71c3e44a4ab4990f90f561398');

// this is the abi of the Counter contract
const abi = [
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "int256",
                "name": "amount",
                "type": "int256"
            }
        ],
        "name": "update",
        "type": "event"
    },
    {
        "inputs": [],
        "name": "counter",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "decrement",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "increment",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    }
];

// this contract is deployed on the rinkeby network
const address = "0x4aB9A307820a61f12c129BAC17c3d078905A87Ed";


async function getPastEvents() {
    const contract = new web3.eth.Contract(abi, address);

    contract.getPastEvents("update",{fromBlock: 0, toBlock: 'latest'})                              
        .then(events => console.log(events))
        .catch((err) => console.error(err));
  }


async function eventQuery() {
  const contract = new web3.eth.Contract(abi, address);
  contract.events.update()
  .on('data', (event) => {
      console.log("\nnew event: ")
      console.log("blockNr: ", event.blockNumber);
      console.log("values:  ", event.returnValues);
  })
  .on('error', console.error);
}

getPastEvents();
// eventQuery();
