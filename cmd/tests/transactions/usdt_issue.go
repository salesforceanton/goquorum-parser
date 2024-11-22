package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/salesforceanton/goquorum-parser/internal/converters"
	"github.com/salesforceanton/goquorum-parser/services/backend/cryptogate/usdt_abi"
)

const (
	testAmount                 = "100000000"
	member1tesseraPublicKey    = "ZTYLEqk5gzZ60uvLJR0zIwPCzzRoRxdyxauV+QxMMRI="
	member2tesseraPublicKey    = "xOt7tRhG5lDBljQfAW7DC0ZDAJlBYbcFZVx4LJEo7xg="
	member3tesseraPublicKey    = "1453wQJ/btTDFlK3DrfVdGglnuPg+RAUXBZMJVzA9Ao="
	gasLimit                   = 47000000
	usdtContractAddressPrivate = "0x3a6050baa11572a72e3eef11f1c7f392d2eba285"
	usdtContractAddressPublic  = "0x4a82Fb9743c03C5fE8395F896CAeB438Ac6F7e95"
	ptmUrl                     = "http://localhost:9081"
	ptm2Url                    = "http://localhost:9082"
	ptm3Url                    = "http://localhost:9083"
	rpcUrl                     = "http://localhost:20000"
	rpcUrl2                    = "http://localhost:20002"
	rpcUrl3                    = "http://localhost:20004"
	ethKey                     = "{\"version\":3,\"id\":\"ee96211f-182c-4ddb-ac9d-9ae69a82bfc3\",\"address\":\"111b0e42105a3412eee350a06a46edf986ce9172\",\"crypto\":{\"ciphertext\":\"0d1b72f1281d8730ada1b4035b611be40e0f74ea9a8788679701c13379e03dc0\",\"cipherparams\":{\"iv\":\"0cccf28477c3f6c7750c2bf60876c58c\"},\"cipher\":\"aes-128-ctr\",\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"salt\":\"72992a5b3e1d18e5fc064da2d5f978dffdd9cff2ab6d1baec98cb949c1fc05a0\",\"n\":262144,\"r\":8,\"p\":1},\"mac\":\"6c581f3ecafa73a2ba60d08827845d9ecd15925cddd000e96d77b1eaf68a1c58\"}}"
	ethKey2                    = "{\"version\":3,\"id\":\"0e1a381a-6578-435c-8047-a797ab89f3cf\",\"address\":\"9eb168ff6ef4369fff996506a1e584e8c0e84a8c\",\"crypto\":{\"ciphertext\":\"275abf6e5123ae2cd5db4ced5d07f9d7d5c56c5ace3cc4ba018c25a8f3028afc\",\"cipherparams\":{\"iv\":\"f8c38eb85109dec7b5fb0753adbdb1e7\"},\"cipher\":\"aes-128-ctr\",\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"salt\":\"8a370e7dea85c109bf7d4e10682ffd201a93869298175a636a932bbf804e7113\",\"n\":262144,\"r\":8,\"p\":1},\"mac\":\"89ded77c79652d36a1faadf03801d2a1f4c7e7d39de138aa00f790f6618fa77d\"}}"
	ethKey3                    = "{\"version\":3,\"id\":\"41b3edd4-7ac5-480f-bcfc-b71a2edc10be\",\"address\":\"5fd43bdb977cfeadb8b3e3a7a205b1bffa954ab1\",\"crypto\":{\"ciphertext\":\"3b00dc67b652cf11fb56e3504dfdfa751fb8c65112f77a95960a7c14ee1c4cc0\",\"cipherparams\":{\"iv\":\"d34b955f75e1f458df6c854abab365e2\"},\"cipher\":\"aes-128-ctr\",\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"salt\":\"304e64fb96d4d01f0f9cfeffc6a8365f1050486c0785c1b31a664ed6bb4e481a\",\"n\":262144,\"r\":8,\"p\":1},\"mac\":\"fee9858acf1d79b52b422c982a7180ae987a0bff5ec97ef934c3df9fbe14c5db\"}}"
)

type TxSetup struct {
	// signer
	ethKey string
	rpcUrl string
	ptmUrl string

	// privacy
	// privateFrom string
	privateFor []string

	// contract
	contractAddress string
}

func sendUsdtIssue(setup TxSetup) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	rpcClient, err := rpc.DialHTTPWithClient(
		setup.rpcUrl,
		&http.Client{
			Timeout: time.Second * 15,
		},
	)
	if err != nil {
		logger.Debug("cryptogateUtils.InitTools", "err", err)
		return
	}

	provider, err := ethclient.NewClient(rpcClient).WithPrivateTransactionManager(setup.ptmUrl)
	if err != nil {
		logger.Error(fmt.Sprintf("ethclient.NewClient(rpcClient).WithPrivateTransactionManager(tessera1Url): %s", err.Error()))
		return
	}

	addressContract := common.HexToAddress(setup.contractAddress)

	gasPrice, err := provider.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Error("Provider.SuggestGasPrice", "err", err)
		return
	}

	chainID, err := provider.ChainID(context.Background())
	if err != nil {
		logger.Debug("tools.Provider.ChainID", "err", err)
		return
	}

	trOpts, err := bind.NewTransactorWithChainID(strings.NewReader(setup.ethKey), "", chainID)
	if err != nil {
		logger.Debug("bind.NewTransactor", "err", err)
		return
	}

	//trOpts.PrivateFrom = setup.privateFrom
	trOpts.PrivateFor = setup.privateFor
	trOpts.GasLimit = gasLimit
	trOpts.IsUsingPrivacyPrecompile = true

	var transaction *types.Transaction

	contractTransactor, err := usdt_abi.NewUsdtAbiTransactor(
		addressContract,
		provider,
	)
	if err != nil {
		logger.Error("usdt_abi.NewUsdtAbiTransactor", "err", err)
		return
	}

	amount, err := converters.StringToBigInt(testAmount)
	if err != nil {
		logger.Error("converters.StringToBigInt", "err", err)
		return
	}

	transaction, err = contractTransactor.Issue(trOpts, amount)
	if err != nil {
		logger.Error("contractTransactor.Issue", "err", err)
		return
	}

	logger.Debug("Transaction sent successfully",
		"function", "issue",
		"to", addressContract.Hex(),
		"gasPrice", gasPrice.String(),
		"txHash", transaction.Hash().Hex(),
	)
}

func main() {
	fmt.Println("-------------------------------PUBLIC CONTRACT-----------------------------------------------------\n")

	// public contract public transaction from 1 member (contract deployed on member 1 node)
	fmt.Println("-------------------------------PUBLIC FROM 1-------------------------------------------------------\n")

	sendUsdtIssue(TxSetup{
		rpcUrl:          rpcUrl,
		ethKey:          ethKey,
		contractAddress: usdtContractAddressPublic,
	})

	fmt.Println("--------------------------------------------------------------------------------------------------\n")

	// public contract public transaction from 2 member (contract deployed on member 1 node)
	fmt.Println("-------------------------------PUBLIC FROM 2------------------------------------------------------\n")

	sendUsdtIssue(TxSetup{
		rpcUrl:          rpcUrl2,
		ethKey:          ethKey2,
		contractAddress: usdtContractAddressPublic,
	})

	fmt.Println("--------------------------------------------------------------------------------------------------\n")

	// public contract public transaction from 2 member (contract deployed on member 1 node)
	fmt.Println("-------------------------------PUBLIC FROM 3------------------------------------------------------\n")

	sendUsdtIssue(TxSetup{
		rpcUrl:          rpcUrl3,
		ethKey:          ethKey3,
		contractAddress: usdtContractAddressPublic,
	})

	fmt.Println("--------------------------------------------------------------------------------------------------\n\n")

	fmt.Println("-------------------------------PRIVATE CONTRACT-----------------------------------------------------\n")

	// private contract public transaction from 1 member (contract deployed on member node privateFor 3 node)
	fmt.Println("-------------------------------PUBLIC FROM 1-------------------------------------------------------\n")

	sendUsdtIssue(TxSetup{
		rpcUrl:          rpcUrl,
		ethKey:          ethKey,
		contractAddress: usdtContractAddressPrivate,
	})

	fmt.Println("--------------------------------------------------------------------------------------------------\n")

	// private contract public transaction from 2 member (contract deployed on member node privateFor 3 node)
	fmt.Println("-------------------------------PUBLIC FROM 2------------------------------------------------------\n")

	sendUsdtIssue(TxSetup{
		rpcUrl:          rpcUrl2,
		ethKey:          ethKey2,
		contractAddress: usdtContractAddressPrivate,
	})

	fmt.Println("--------------------------------------------------------------------------------------------------\n")

	// private contract public transaction from 3 member (contract deployed on member node privateFor 3 node)
	fmt.Println("-------------------------------PUBLIC FROM 3------------------------------------------------------\n")

	sendUsdtIssue(TxSetup{
		rpcUrl:          rpcUrl3,
		ethKey:          ethKey3,
		contractAddress: usdtContractAddressPrivate,
	})

	fmt.Println("--------------------------------------------------------------------------------------------------\n")

	// private contract private transaction from 1 to 3 (contract deployed on 1 member node privateFor 3 node)
	fmt.Println("-------------------------------PRIVATE FROM 1 to 3------------------------------------------------\n")

	sendUsdtIssue(TxSetup{
		//privateFrom:     member1tesseraPublicKey,
		privateFor:      []string{member3tesseraPublicKey},
		ptmUrl:          ptmUrl,
		rpcUrl:          rpcUrl,
		ethKey:          ethKey,
		contractAddress: usdtContractAddressPrivate,
	})

	time.Sleep(7 * time.Second)
	fmt.Println("--------------------------------------------------------------------------------------------------\n")

	// private contract private transaction from 3 to 1 (contract deployed on 1 member node privateFor 3 node)
	fmt.Println("-------------------------------PRIVATE FROM 3 to 1------------------------------------------------\n")

	sendUsdtIssue(TxSetup{
		//privateFrom:     member3tesseraPublicKey,
		privateFor:      []string{member1tesseraPublicKey},
		ptmUrl:          ptm3Url,
		rpcUrl:          rpcUrl3,
		ethKey:          ethKey3,
		contractAddress: usdtContractAddressPrivate,
	})

	time.Sleep(7 * time.Second)
	fmt.Println("--------------------------------------------------------------------------------------------------\n")

	// private contract private transaction from 2 to 1 (contract deployed on 1 member node privateFor 3 node)
	fmt.Println("-------------------------------PRIVATE FROM 2 to 1------------------------------------------------\n")

	sendUsdtIssue(TxSetup{
		//privateFrom:     member2tesseraPublicKey,
		privateFor:      []string{member1tesseraPublicKey},
		ptmUrl:          ptm2Url,
		rpcUrl:          rpcUrl2,
		ethKey:          ethKey2,
		contractAddress: usdtContractAddressPrivate,
	})

	fmt.Println("--------------------------------------------------------------------------------------------------\n")
	fmt.Println("----------------------------------------FINISH------------------------------------------------------")
}
