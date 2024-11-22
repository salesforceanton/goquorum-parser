package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/salesforceanton/goquorum-parser/services/backend/cryptogate/usdt_abi"
)

const (
	member1tesseraPublicKey = "ZTYLEqk5gzZ60uvLJR0zIwPCzzRoRxdyxauV+QxMMRI="
	member3tesseraPublicKey = "1453wQJ/btTDFlK3DrfVdGglnuPg+RAUXBZMJVzA9Ao="
	gasLimit                = 47000000
	ptmUrl                  = "http://localhost:9081"
	rpcUrl                  = "http://localhost:20000"
	ethKey                  = "{\"version\":3,\"id\":\"ee96211f-182c-4ddb-ac9d-9ae69a82bfc3\",\"address\":\"111b0e42105a3412eee350a06a46edf986ce9172\",\"crypto\":{\"ciphertext\":\"0d1b72f1281d8730ada1b4035b611be40e0f74ea9a8788679701c13379e03dc0\",\"cipherparams\":{\"iv\":\"0cccf28477c3f6c7750c2bf60876c58c\"},\"cipher\":\"aes-128-ctr\",\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"salt\":\"72992a5b3e1d18e5fc064da2d5f978dffdd9cff2ab6d1baec98cb949c1fc05a0\",\"n\":262144,\"r\":8,\"p\":1},\"mac\":\"6c581f3ecafa73a2ba60d08827845d9ecd15925cddd000e96d77b1eaf68a1c58\"}}"
)

type DeploySetup struct {
	// deployer config
	rpcUrl string
	ethKey string
	ptmUrl string

	// privacy
	privateFor []string
	// privateFrom string TODO: check is this param unnecessary

	// contract
	name       string
	abi        string
	initialize []interface{}
}

func deployContract(setup DeploySetup) {
	isPrivate := len(setup.privateFor) > 0

	// connect to member1
	rpcClient, err := rpc.DialHTTP(setup.rpcUrl)
	if err != nil {
		println(err.Error())
		return
	}

	provider, err := ethclient.NewClient(rpcClient).WithPrivateTransactionManager(setup.ptmUrl)
	if err != nil {
		println(err.Error())
		return
	}

	chainID, err := provider.ChainID(context.Background())
	if err != nil {
		println(err.Error())
		return
	}

	trOpts, err := bind.NewTransactorWithChainID(strings.NewReader(setup.ethKey), "", chainID)
	if err != nil {
		println(err.Error())
		return
	}
	trOpts.GasLimit = gasLimit
	if isPrivate {
		trOpts.PrivateFor = setup.privateFor
		//trOpts.PrivateFrom = setup.privateFrom
	}

	fmt.Println("trOpts", *trOpts)

	parsedAbi, err := abi.JSON(strings.NewReader(setup.abi))
	if err != nil {
		println(err.Error())
		return
	}

	contractBytecode, err := os.ReadFile("./TetherToken.bin")
	if err != nil {
		println(err.Error())
		return
	}

	address, _, _, err := bind.DeployContract(trOpts, parsedAbi, common.FromHex(string(contractBytecode)), provider, setup.initialize...)
	if err != nil {
		println(err.Error())
		return
	}

	println("contractAddress: " + strings.ToLower(address.Hex()))
}

func main() {
	setup := DeploySetup{
		name:   "USDT",
		rpcUrl: rpcUrl,
		ethKey: ethKey,
		ptmUrl: ptmUrl,
		abi:    usdt_abi.UsdtABI,
	}

	// contract init params
	var initialSupply = big.NewInt(100000000000000)
	const name = "TetherToken"
	const symbol = "USDT"
	var decimals = big.NewInt(6)

	setup.initialize = []interface{}{initialSupply, name, symbol, decimals}

	// fmt.Println("-------------------------------PUBLIC CONTRACT DEPLOY----------------------------------------------\n")

	// deployContract(setup)

	// fmt.Println("-----------------------------------------END-----------------------------------------------------\n\n")

	fmt.Printf("-------------------------------CONTRACT DEPLOY: %s ---------------------------------------------\n", setup.name)

	setup.privateFor = []string{member1tesseraPublicKey}
	//setup.privateFrom = member1tesseraPublicKey
	deployContract(setup)

	fmt.Println("-----------------------------------------END-----------------------------------------------------")
}
