#!/bin/sh

# install basics
install_base(){
    sudo apt install git -y
    sudo apt install curl -y
    sudo apt install build-essential -y
}

# install docker
install_docker(){
    sudo apt install \
        apt-transport-https \
        ca-certificates \
        gnupg-agent \
        software-properties-common -y
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
    sudo apt-key fingerprint 0EBFCD88
    sudo add-apt-repository \
        "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
        $(lsb_release -cs) \
        stable"
    sudo apt-get update
    sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose -y
    sudo usermod -aG docker $USER
    newgrp docker
}

# install golang
install_golang(){
    curl -o "go1.15.tar.gz" https://storage.googleapis.com/golang/go1.15.linux-amd64.tar.gz
    sudo chmod +x go1.15.tar.gz
    sudo tar -C /opt -xzf "go1.15.tar.gz"

    mkdir /home/vagrant/project/go
    mkdir /home/vagrant/project/go/src
    mkdir /home/vagrant/project/go/pkg
    mkdir /home/vagrant/project/go/lib

    export GOROOT=/opt/go
    export GOPATH=/home/vagrant/project/go
    export PATH=$PATH:/$GOROOT/bin:$GOPATH/bin

    sudo touch "$HOME/.bashrc"
    {
        echo "export GOROOT=/opt/go"
        echo "export GOPATH=/home/vagrant/project/go"
        echo "export PATH=$PATH:$GOROOT/bin:$GOPATH/bin"
    } >> "$HOME/.bashrc"
}

# install node
install_node(){
    curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -
    sudo apt install nodejs -y
}

# install zsh
install_zsh(){
    sudo apt install zsh -y
    sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
    
}

install_ganache(){
    mkdir ~/.ganache
    cd ~/ganache
    mkdir bin
    curl -o "ganache" https://github.com/trufflesuite/ganache/releases/download/v2.1.2/ganache-2.1.2-linux-x86_64.AppImage
    sudo chmod +x ganache
    mv ganache ./bin 

    sudo touch "$HOME/.bashrc"
    {
        echo "export PATH=$PATH:$HOME/.ganache/bin"
    } >> "$HOME/.bashrc"
    
}

# install ipfs-desktop
install_ipfs(){

    git clone https://github.com/ipfs-shipyard/ipfs-desktop.git
    cd ipfs-desktop
    npm ci
    npm run build
    npm start
}

# install solc-select
install_solc(){
    cd /opt
    git clone https://github.com/crytic/solc-select.git
    ./solc-select/scripts/install.sh

    export PATH=/home/vagrant/.solc-select:$PATH

    sudo touch "$HOME/.bashrc"
    {
        echo "export PATH=/home/vagrant/.solc-select:$PATH"
    } >> "$HOME/.bashrc"
}

# install protoc
install_protoc(){
    sudo apt install protobuf-compiler -y
}

# install go-ethereum
install_goeth(){
    sudo add-apt-repository -y ppa:ethereum/ethereum
    sudo apt update -y
    sudo apt install ethereum -y

    go get -u github.com/ethereum/go-ethereum/...
    cd $GOPATH/src/github.com/ethereum/go-ethereum
    make && make devtools
    go install ./cmd/geth
    cp ./build/geth $GOPATH/bin
    
}

# install finish
install_finish(){
    sudo touch "$HOME/.bashrc"
    {
        echo "zsh"
    } >> "$HOME/.bashrc"
}

install_base
install_docker
install_golang
install_node
install_zsh
install_ganache
install_ipfs
install_solc
install_protoc
install_goeth
install_finish
