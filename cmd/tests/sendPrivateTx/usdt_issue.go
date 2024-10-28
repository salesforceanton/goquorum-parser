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
	member3tesseraPublicKey    = "1453wQJ/btTDFlK3DrfVdGglnuPg+RAUXBZMJVzA9Ao="
	gasLimit                   = 47000000
	usdtContractAddressPrivate = "0xebaf93b936e40dcd4f5fc78c15d01e91d09e085d"
	usdtContractAddressPublic  = "0xe2889ac4b39a4d4ff9c81d372d6a682c34bda1e9"
	ptmUrl                     = "http://localhost:9081"
	rpcUrl                     = "http://localhost:20000"
	ethKey                     = "{\"version\":3,\"id\":\"ee96211f-182c-4ddb-ac9d-9ae69a82bfc3\",\"address\":\"111b0e42105a3412eee350a06a46edf986ce9172\",\"crypto\":{\"ciphertext\":\"0d1b72f1281d8730ada1b4035b611be40e0f74ea9a8788679701c13379e03dc0\",\"cipherparams\":{\"iv\":\"0cccf28477c3f6c7750c2bf60876c58c\"},\"cipher\":\"aes-128-ctr\",\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"salt\":\"72992a5b3e1d18e5fc064da2d5f978dffdd9cff2ab6d1baec98cb949c1fc05a0\",\"n\":262144,\"r\":8,\"p\":1},\"mac\":\"6c581f3ecafa73a2ba60d08827845d9ecd15925cddd000e96d77b1eaf68a1c58\"}}"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	rpcClient, err := rpc.DialHTTPWithClient(
		rpcUrl,
		&http.Client{
			Timeout: time.Second * 15,
		},
	)
	if err != nil {
		logger.Debug("cryptogateUtils.InitTools", "err", err)
		return
	}

	provider, err := ethclient.NewClient(rpcClient).WithPrivateTransactionManager(ptmUrl)
	if err != nil {
		logger.Error(fmt.Sprintf("ethclient.NewClient(rpcClient).WithPrivateTransactionManager(tessera1Url): %s", err.Error()))
		return
	}

	addressContract := common.HexToAddress(usdtContractAddressPrivate)

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

	trOpts, err := bind.NewTransactorWithChainID(strings.NewReader(ethKey), "", chainID)
	if err != nil {
		logger.Debug("bind.NewTransactor", "err", err)
		return
	}

	trOpts.PrivateFrom = member1tesseraPublicKey
	trOpts.PrivateFor = []string{member3tesseraPublicKey}
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

// 0x44f3d8794bb12c9273d29efc83c64fbc77b5397465764d65351646ed4a705da9
