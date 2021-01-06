package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	// My own dependencies
	"ruud/GO/2_simpleStorage/contract/simpleStorage"
)


const infuraURL		= "https://rinkeby.infura.io/v3/634d2ee71c3e44a4ab4990f90f561398"

const hash			= "0x9f8e73263cdad38f1a9aee8207712afb3d69442d626f0fecfcf1ebfb4d2a62e4"

const CONTRACT_ADDRESS = "0x43585B6FB9aF1Bd09C671BFaC8a48102e28F9686" 


var accounts = []struct {
	Pk  string
	Adr common.Address
}{
	{"f238a37e42b7062bdbc062a1833a6361f9a6d0e324a95ca2f7c4c3034e67ee5c", common.HexToAddress("0x559BC07434C89c5496d790DFD2885dC966F9113a")},
	{"6c0081a5b9511910a6cec018a99d3031197f079cde51c1a78124750a990cdd08", common.HexToAddress("0x54806DD512b21814aa560D627432a75720ed6bB3")},
	{"a3968111221303d38214eb7b2db9b04cefb2300b72771d65d258e08322dc573d", common.HexToAddress("0x6dc89393fa30b64c56deff31daacf10cedcd852d")},
}


func deploySimpleStorage() {
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(accounts[0].Pk)
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), accounts[0].Adr)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	address, tx, instance, err := simpleStorage.DeploySimpleStorage(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())

	_ = tx 
	_ = instance
}

func queryValue() {
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}

	contract, err := simpleStorage.NewSimpleStorage(common.HexToAddress(CONTRACT_ADDRESS), client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	value, _ := contract.Get(&bind.CallOpts{})

	fmt.Println(value)

}

func setValue() {
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}

	contract, err := simpleStorage.NewSimpleStorage(common.HexToAddress(CONTRACT_ADDRESS), client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(accounts[0].Pk)
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), accounts[0].Adr)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) 
	auth.GasLimit = uint64(3000000) 
	auth.GasPrice = gasPrice
	contract.Set(auth, 100)

}


func main() {

	fmt.Println("\n\n-----------------------------------\n")
	fmt.Println("Deploy SimpleStorage Contract\n")
	deploySimpleStorage()

	time.Sleep(2 * time.Second)

	fmt.Println("\n\n-----------------------------------\n")
	fmt.Println("Invoke Blockchain\n")
	fmt.Println("Setting value....\n")

	setValue()

	time.Sleep(2 * time.Second)

	fmt.Println("\n\n-----------------------------------\n")
	fmt.Println("Query Blockchain\n")
	fmt.Println("Getting Value\n")
	queryValue()

}