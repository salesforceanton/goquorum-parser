package main

import (
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

func main() {
	// connect to member1
	rpcClient, err := rpc.DialHTTP(rpcUrl)
	if err != nil {
		println(err.Error())
		return
	}

	trOpts, err := bind.NewTransactor(strings.NewReader(ethKey), "")
	if err != nil {
		println(err.Error())
		return
	}
	trOpts.GasLimit = gasLimit
	trOpts.PrivateFor = []string{member3tesseraPublicKey}
	trOpts.PrivateFrom = member1tesseraPublicKey

	ethClient, err := ethclient.NewClient(rpcClient).WithPrivateTransactionManager(ptmUrl)
	if err != nil {
		println(err.Error())
		return
	}

	parsedAbi, err := abi.JSON(strings.NewReader(usdt_abi.UsdtABI))
	if err != nil {
		println(err.Error())
		return
	}

	contractBytecode, err := os.ReadFile("./TetherToken.bin")
	if err != nil {
		println(err.Error())
		return
	}

	// params
	var initialSupply = big.NewInt(100000000000000)
	const name = "TetherToken"
	const symbol = "USDT"
	var decimals = big.NewInt(6)

	initialize := []interface{}{initialSupply, name, symbol, decimals}

	address, _, _, err := bind.DeployContract(trOpts, parsedAbi, common.FromHex(string(contractBytecode)), ethClient, initialize...)
	if err != nil {
		println(err.Error())
		return
	}

	println("contractAddress: " + address.Hex())
}

// 0xEbAf93B936e40dCd4F5Fc78C15D01E91D09e085D
