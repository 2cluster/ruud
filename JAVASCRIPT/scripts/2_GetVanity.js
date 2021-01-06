Web3 = require('web3');

const createVanity = async (vanity) => {
    web3 = new Web3();
    let i = 0;
    let findlength = vanity.length+2;
    let prefix;
    do {
        newAddress=web3.eth.accounts.create();
        prefix=newAddress.address.slice(2,findlength);    
        if (++i % 1000 == 0) console.log(i);
    } while (prefix !=vanity );

    console.log(`Found an address with prefix ${prefix}`);
    console.log(`address=${newAddress.address}`);
    console.log(`privatekey=${newAddress.privateKey}`);
}
createVanity('111')