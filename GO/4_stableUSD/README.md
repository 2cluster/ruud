In this folder i created a API server that is pegged to a already deployed ERC20 contract on the rinkeby network
Also the accounts are hardcoded in the application, no need to use your own credentials.

the contract is provided inside this project folder but the need to be relocated relative to the gopath

available accounts are:
- party1
- party2
- master


the API has routes to:

- get balance of accounts
- approve the master account to take a given amount from account
- transfer from own account to another


#    r.HandleFunc("/account/{acc}", getAccount).Methods("GET")
#    r.HandleFunc("/account/{acc}/approve", aproveAllowance).Methods("POST")
#    r.HandleFunc("/account/{acc}/tranfer", transfer).Methods("POST")
#    r.HandleFunc("/", homePage).Methods("GET")