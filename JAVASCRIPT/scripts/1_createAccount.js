Web3 = require('web3');

const createAccount = async () => {
    web3 = new Web3();


    let newAddress=web3.eth.accounts.create();
    const account = {
        privateKey: newAddress.privateKey,
        address: newAddress.address,
    }
    return account

}
createAccount().then(console.log)