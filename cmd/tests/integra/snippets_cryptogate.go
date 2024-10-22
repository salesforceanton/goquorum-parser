package main

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
	"github.com/salesforceanton/goquorum-parser/internal/protocol/messages/cryptogatemessages"
	"github.com/salesforceanton/goquorum-parser/internal/transport"
	cryptogateservice "github.com/salesforceanton/goquorum-parser/services/backend/cryptogate"
	"github.com/spf13/viper"
)

const (
	member1tesseraPublicKey = "ZTYLEqk5gzZ60uvLJR0zIwPCzzRoRxdyxauV+QxMMRI="
	member3tesseraPublicKey = "1453wQJ/btTDFlK3DrfVdGglnuPg+RAUXBZMJVzA9Ao="
	gasLimit                = 47000000
)

func testSendTransaction(req cryptogatemessages.SendTransactionRequest, isPrivate bool) {
	logger := slog.Default()

	client, err := transport.New(transport.Config{
		Host:     viper.GetString("NATS_HOST"),
		Port:     viper.GetString("NATS_PORT"),
		User:     viper.GetString("NATS_USER"),
		Password: viper.GetString("NATS_PASS"),
	}, logger)
	if err != nil {
		fmt.Printf("transport.New: %s\n", err.Error())
		return
	}

	transport := cryptogateservice.NewTransport(client, logger, time.Second*5)

	// add only for tests
	if isPrivate {
		req.PrivateFor = []string{member3tesseraPublicKey}
		req.PrivateFrom = member1tesseraPublicKey
	}

	_, err = transport.SendTransaction(context.Background(), req)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
}

func testMintToken(contract, amount string, isPrivate bool) {
	testSendTransaction(cryptogatemessages.SendTransactionRequest{
		TypeSmartContract: cryptogate.SmartContractType(strings.ToUpper(contract)),
		NameFunction:      "issue",
		Amount:            amount,
	}, isPrivate)
}

func testTransferToken(contract, amount, userAddress string) {
	testSendTransaction(cryptogatemessages.SendTransactionRequest{
		TypeSmartContract: cryptogate.SmartContractType(strings.ToUpper(contract)),
		NameFunction:      "transfer",
		UserAddress:       userAddress,
		Amount:            amount,
	}, false)
}

func testGetNativeBalance(address string) {
	logger := slog.Default()

	client, err := transport.New(transport.Config{
		Host:     viper.GetString("NATS_HOST"),
		Port:     viper.GetString("NATS_PORT"),
		User:     viper.GetString("NATS_USER"),
		Password: viper.GetString("NATS_PASS"),
	}, logger)
	if err != nil {
		fmt.Printf("transport.New: %s\n", err.Error())
		return
	}

	transport := cryptogateservice.NewTransport(client, logger, time.Second*5)

	res, err := transport.GetBalanceNative(context.Background(), cryptogatemessages.BalanceNativeRequest{
		Address: address,
	})
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}

	fmt.Printf("Balance: %v\n", res)
}

func testGetBalanceToken(contract, address string) {
	logger := slog.Default()

	client, err := transport.New(transport.Config{
		Host:     viper.GetString("NATS_HOST"),
		Port:     viper.GetString("NATS_PORT"),
		User:     viper.GetString("NATS_USER"),
		Password: viper.GetString("NATS_PASS"),
	}, logger)
	if err != nil {
		fmt.Printf("transport.New: %s\n", err.Error())
		return
	}

	transport := cryptogateservice.NewTransport(client, logger, time.Second*5)
	fmt.Println(contract)
	fmt.Println(cryptogate.SmartContractType(strings.ToUpper(contract)))

	res, err := transport.GetBalanceToken(context.Background(), cryptogatemessages.BalanceTokenRequest{
		Address:           address,
		TypeSmartContract: cryptogate.SmartContractType(strings.ToUpper(contract)),
	})
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}

	fmt.Printf("Balance: %v\n", res)
}
