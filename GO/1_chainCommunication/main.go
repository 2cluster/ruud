package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

)


const infuraURL		= "https://rinkeby.infura.io/v3/634d2ee71c3e44a4ab4990f90f561398"

const hash			= "0x9f8e73263cdad38f1a9aee8207712afb3d69442d626f0fecfcf1ebfb4d2a62e4"


var accounts = []struct {
	Pk  string
	Adr common.Address
}{
	{"f238a37e42b7062bdbc062a1833a6361f9a6d0e324a95ca2f7c4c3034e67ee5c", common.HexToAddress("0x559BC07434C89c5496d790DFD2885dC966F9113a")},
	{"6c0081a5b9511910a6cec018a99d3031197f079cde51c1a78124750a990cdd08", common.HexToAddress("0x54806DD512b21814aa560D627432a75720ed6bB3")},
	{"a3968111221303d38214eb7b2db9b04cefb2300b72771d65d258e08322dc573d", common.HexToAddress("0x6dc89393fa30b64c56deff31daacf10cedcd852d")},
}

func queryBlockByID() {
	// Create Client
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatal(err)
	}

	// Get Latest Block Number
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := header.Number

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block nr:         ", blockNumber)
	fmt.Println("Block time:       ", block.Time())
	fmt.Println("Block difeculty:  ", block.Difficulty())
	fmt.Println("Block hash:       ", block.Hash())
	fmt.Println("Block tx count:   ", len(block.Transactions()))
}

func queryTxByHash() {
	// Create Client
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	tx, pending, _ := client.TransactionByHash(ctx, common.HexToHash(hash))
	if !pending {
		fmt.Println("TX Hash:           ", tx.Hash().Hex())
		fmt.Println("TX Value:          ", tx.Value().String())
		fmt.Println("TX Gas:            ", tx.Gas())
		fmt.Println("TX Gas Price:      ", tx.GasPrice().Uint64())
		fmt.Println("TX Nonce:          ", tx.Nonce())
		fmt.Println("TX Data:           ", tx.Data())
		fmt.Println("TX To:             ", tx.To().Hex())

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("TX Status:    ", receipt.Status)
	}
}

func ethTransfer() {
	// Create Client
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(accounts[2].Pk)
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), accounts[0].Adr)
	if err != nil {
		log.Fatal(err)
	}

	// Setup TX values
	value := big.NewInt(500000000000000000) // in wei (.5 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	to := accounts[1].Adr
	var data []byte
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sent %s wei to %s: %s\n", value.String(), to.Hex(), signedTx.Hash().Hex())
}




func main() {

	fmt.Println("\n\n-----------------------------------\n\n")
	fmt.Println("Get block info\n")
	queryBlockByID()

	fmt.Println("\n\n-----------------------------------\n\n")
	fmt.Println("Get Tx info\n")
	queryTxByHash()

	fmt.Println("\n\n-----------------------------------\n\n")
	fmt.Println("Transfer Eth\n")
	ethTransfer()

}