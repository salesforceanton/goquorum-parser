package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/salesforceanton/goquorum-parser/services/backend/cryptogate/usdt_abi"
)

const (
	userAddress                = "0x111B0E42105A3412eEE350A06a46EdF986Ce9172"
	member1tesseraPublicKey    = "ZTYLEqk5gzZ60uvLJR0zIwPCzzRoRxdyxauV+QxMMRI="
	member3tesseraPublicKey    = "1453wQJ/btTDFlK3DrfVdGglnuPg+RAUXBZMJVzA9Ao="
	gasLimit                   = 47000000
	usdtContractAddressPrivate = "0x2cd3fDD0e7c91c656edaF094A1671729A8ce48c8"
	usdtContractAddressPublic  = "0x4a82Fb9743c03C5fE8395F896CAeB438Ac6F7e95"
	ptmUrl                     = "http://localhost:9081"
	ptm2Url                    = "http://localhost:9082"
	ptm3Url                    = "http://localhost:9083"
	rpcUrl                     = "http://localhost:20000"
	rpcUrl2                    = "http://localhost:20002"
	rpcUrl3                    = "http://localhost:20004"
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

	contractCaller, err := usdt_abi.NewUsdtAbiCaller(
		addressContract,
		provider,
	)
	if err != nil {
		logger.Error("usdt_abi.NewUsdtAbiCaller", "err", err)
		return
	}

	// view hex data
	// parsedABI, err := abi.JSON(strings.NewReader(cryptogate.GetABI(cryptogate.SmartContractTypeUSDT)))
	// if err != nil {
	// 	return
	// }

	// balanceData, err := parsedABI.Pack("balanceOf", common.HexToAddress(userAddress))
	// if err != nil {
	// 	return
	// }

	// fmt.Println("0x" + hex.EncodeToString(balanceData))
	//

	balance, err := contractCaller.BalanceOf(nil, common.HexToAddress(userAddress))
	if err != nil {
		logger.Error("contractCaller.BalanceOf", "err", err)
		return
	}

	logger.Debug("Usdt token balance",
		"address", addressContract.Hex(),
		"balance", balance.String(),
	)
}
