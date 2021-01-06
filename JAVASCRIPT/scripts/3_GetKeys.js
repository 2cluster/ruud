const HDKey = require('hdkey')
const bip39 = require("bip39");
const ethUtil = require('ethereumjs-util');

const mnemonic = 'injury arch try label empty sniff destroy candy bid talk above remember';


const getKeys = async (mnemonic, count) => {
    const seed = await bip39.mnemonicToSeed(mnemonic); 
    const root = HDKey.fromMasterSeed(seed);
    
    let accounts = [];
    for (let i = 0; i < count; i++) {
        const derivedObj = root.derive("m/44'/60'/0'/0/"+i);
        const prk = derivedObj._privateKey;
        const pbk = ethUtil.privateToPublic(prk);
        const adr = "0x"+ethUtil.publicToAddress(pbk).toString('hex');

        await accounts.push({ address: adr, privateKey: prk.toString('hex')});
    }

    return accounts
}

getKeys(mnemonic, 2).then(console.log)