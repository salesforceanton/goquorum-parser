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

func testSendTransaction(req cryptogatemessages.SendTransactionRequest) {
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

	_, err = transport.SendTransaction(context.Background(), req)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
}

func testMintToken(contract, amount string) {
	testSendTransaction(cryptogatemessages.SendTransactionRequest{
		TypeSmartContract: cryptogate.SmartContractType(strings.ToUpper(contract)),
		NameFunction:      "issue",
		Amount:            amount,
	})
}

func testTransferToken(contract, amount, userAddress string) {
	testSendTransaction(cryptogatemessages.SendTransactionRequest{
		TypeSmartContract: cryptogate.SmartContractType(strings.ToUpper(contract)),
		NameFunction:      "transfer",
		UserAddress:       userAddress,
		Amount:            amount,
	})
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
