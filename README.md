DAPP
 - here i created a application that should work without having to install dependencies
    - git clone https://github.com/2cluster/token-application.git
    - git clone https://github.com/2cluster/upgrade-test.git

DEV_VM
 - this is just a template structure that i use to create new development environments using a vagrant VM and a script to install software

GO 
 - here i put all the go scripts and projects. i tried to make a vendor folder for all the go files so that you dont have to download dependencies but due to some error this is not possible, one needed dependency is of written in C+ and when u run 'GO MOD VENDOR' this dependency is missing

 for the installation of the dependencies, the 'ruud' folder needs to be placed in $GOPATH/src, so next to where the github.com folder lives. then at the location of the go script run
 
- go get .

 to get the dependencies
 then to run the script

- go run <filename.go>

JAVASCRIPT 

 contains all the scripts written in .js, node_modules included


\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\

Added the following

- events folder
    ./events

- oracles folder
    ./oracles

- upgrade-test project
    - git clone https://github.com/2cluster/upgrade-test.git

    this project can deplot a simpleStorage contract with a proxy

    then a simpleStorageV2 contract can replace the contract that is now reachable from the same address
    In the V2 contract i have added a new function 

    All commands can be found in the readme



 