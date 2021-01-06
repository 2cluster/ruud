package main

import (
	"encoding/json"
	"log"
	"net/http"
    "fmt"
    "math/big"
    "crypto/ecdsa"
    "context"
    "io/ioutil"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
    "github.com/gorilla/mux"
    "github.com/ethereum/go-ethereum/ethclient"

	// looking for this folder in GOPATH => $GOPATH/contracts/SUSD
    "ruud/GO/2_simpleStorage/contract/SUSD"
)

const CONTRACT_ADDRESS = "0x566160489E6548120385Fd5397BEf0E2b608C602" 
const infuraKEY		= "634d2ee71c3e44a4ab4990f90f561398"
const network 		= "rinkeby"

type Msg struct {
    Amount int64
}

type Account struct {
	Holder string `json:"holder"`
	PrivateKey *ecdsa.PrivateKey `json:"privateKey"`
	Address common.Address `json:"address"`
	Allowance *big.Int `json:"allowance"`
	Balance *big.Int `json:"balance"`
}

var accounts []Account

func genAccounts(){
	var account Account
	
	pk1, err := crypto.HexToECDSA("f238a37e42b7062bdbc062a1833a6361f9a6d0e324a95ca2f7c4c3034e67ee5c")
	if err != nil {
		fmt.Println("error")
	}

	account 				= Account{Holder: "party1"}
	account.PrivateKey 	 = pk1
	account.Address 		 = common.HexToAddress("0x559BC07434C89c5496d790DFD2885dC966F9113a")
	
	accounts = append(accounts, account)

	pk2, err := crypto.HexToECDSA("6c0081a5b9511910a6cec018a99d3031197f079cde51c1a78124750a990cdd08")
	if err != nil {
		fmt.Println("error")
	}

	a2 				:= Account{Holder: "party2"}
	a2.PrivateKey 	 = pk2
	a2.Address 		 = common.HexToAddress("0x54806DD512b21814aa560D627432a75720ed6bB3")

	accounts = append(accounts, a2)


	pk3, err := crypto.HexToECDSA("a3968111221303d38214eb7b2db9b04cefb2300b72771d65d258e08322dc573d")
	if err != nil {
		fmt.Println("error")
	}

	a3 				:= Account{Holder: "master"}
	a3.PrivateKey 	 = pk3
	a3.Address 		 = common.HexToAddress("0x6dc89393fa30b64c56deff31daacf10cedcd852d")

    accounts = append(accounts, a3)
    

}



func deployContract() {

	dealblock 	:= accounts[2]


	client, err := ethclient.Dial("https://"+network+".infura.io/v3/"+infuraKEY)
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), dealblock.Address)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(dealblock.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice


	address, tx, instance, err := SUSD.DeploySUSD(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}


func getAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(accounts)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	for _, account := range accounts {
		if account.Holder == params["acc"] {

            account.Balance = queryBalance(account)
            account.Allowance = queryAllowance(account, accounts[2].Address)

			json.NewEncoder(w).Encode(account)
			return
		}
	}
	json.NewEncoder(w).Encode(&Account{})
}

func queryBalance(owner Account)  *big.Int {
	client, err := ethclient.Dial("https://"+network+".infura.io/v3/"+infuraKEY)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}

	contract, err := SUSD.NewSUSD(common.HexToAddress(CONTRACT_ADDRESS), client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	value, _ := contract.BalanceOf(&bind.CallOpts{}, owner.Address)

	return value

}

func queryAllowance(owner Account, spender common.Address)  *big.Int {
	client, err := ethclient.Dial("https://"+network+".infura.io/v3/"+infuraKEY)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}

	contract, err := SUSD.NewSUSD(common.HexToAddress(CONTRACT_ADDRESS), client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	value, _ := contract.Allowance(&bind.CallOpts{}, owner.Address, spender)

	return value

}


func aproveAllowance(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    var p Msg

    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }


	w.Header().Set("Content-Type", "application/json")
	// Loop through books and find one with the id from the params
	for _, account := range accounts {
		if account.Holder == params["acc"] {

            client, err := ethclient.Dial("https://"+network+".infura.io/v3/"+infuraKEY)
            if err != nil {
                log.Fatalf("Failed to connect to the Ethereum network: %v", err)
            }
        
            contract, err := SUSD.NewSUSD(common.HexToAddress(CONTRACT_ADDRESS), client)
            if err != nil {
                log.Fatalf("Failed to instantiate contract: %v", err)
            }
        
            nonce, err := client.PendingNonceAt(context.Background(), account.Address)
            if err != nil {
                log.Fatal(err)
            }
        
            gasPrice, err := client.SuggestGasPrice(context.Background())
            if err != nil {
                log.Fatal(err)
            }
        
            auth := bind.NewKeyedTransactor(account.PrivateKey)
            auth.Nonce = big.NewInt(int64(nonce))
            auth.Value = big.NewInt(0) 
            auth.GasLimit = uint64(3000000) 
            auth.GasPrice = gasPrice

            contract.Approve(auth, accounts[2].Address, big.NewInt(p.Amount))
			fmt.Printf("%s approved master to take %v from your account ", account.Holder, p.Amount)
			fmt.Fprintf(w,"Approved")
			return
		}
	}
	json.NewEncoder(w).Encode(&Account{})
}



func transfer(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Printf("Error reading body: %v", err)
        http.Error(w, "can't read body", http.StatusBadRequest)
        return
    }
    fmt.Println(body)

	w.Header().Set("Content-Type", "application/json")
	// Loop through books and find one with the id from the params
	for _, account := range accounts {
		if account.Holder == params["acc"] {
            client, err := ethclient.Dial("https://"+network+".infura.io/v3/"+infuraKEY)
            if err != nil {
                log.Fatalf("Failed to connect to the Ethereum network: %v", err)
            }

            contract, err := SUSD.NewSUSD(common.HexToAddress(CONTRACT_ADDRESS), client)
            if err != nil {
                log.Fatalf("Failed to instantiate contract: %v", err)
            }

            nonce, err := client.PendingNonceAt(context.Background(), account.Address)
            if err != nil {
                log.Fatal(err)
            }

            gasPrice, err := client.SuggestGasPrice(context.Background())
            if err != nil {
                log.Fatal(err)
            }

            auth := bind.NewKeyedTransactor(account.PrivateKey)
            auth.Nonce = big.NewInt(int64(nonce))
            auth.Value = big.NewInt(0) 
            auth.GasLimit = uint64(3000000) 
            auth.GasPrice = gasPrice
            
            // i, _ := strconv.ParseInt(body["amount"], 10, 64)

            contract.Transfer(auth, common.HexToAddress(params["to"]), big.NewInt(100))

            return
        }
	}
	json.NewEncoder(w).Encode(&Account{})
}


func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func main() {
    r := mux.NewRouter()

    genAccounts()
    // deployContract()

    r.HandleFunc("/account/{acc}", getAccount).Methods("GET")
    r.HandleFunc("/account/{acc}/approve", aproveAllowance).Methods("POST")
    r.HandleFunc("/account/{acc}/tranfer", transfer).Methods("POST")
    r.HandleFunc("/", homePage).Methods("GET")

	fmt.Println("Listening on port 9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}
