package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
	"github.com/salesforceanton/goquorum-parser/internal/protocol/messages/cryptogatemessages"
	"github.com/salesforceanton/goquorum-parser/internal/protocol/subjects"
	"github.com/salesforceanton/goquorum-parser/internal/transport"
	cryptogateservice "github.com/salesforceanton/goquorum-parser/services/backend/cryptogate"
	"github.com/spf13/viper"
)

func testSendTransaction(contract, nameFunction string) {
	logger := slog.Default()

	transportClient, err := transport.New(transport.Config{
		Host:     viper.GetString("NATS_HOST"),
		Port:     viper.GetString("NATS_PORT"),
		User:     viper.GetString("NATS_USER"),
		Password: viper.GetString("NATS_PASS"),
	}, logger)
	if err != nil {
		fmt.Printf("transport.New: %s\n", err.Error())
		return
	}

	req := cryptogatemessages.SendTransactionRequest{
		TypeSmartContract: cryptogate.SmartContractType(contract),
		NameFunction:      nameFunction,
	}

	msg, err := transport.Marshal(&req)
	if err != nil {
		fmt.Printf("transport.Marshal %s\n", err.Error())
		return
	}

	err = transportClient.Publish(subjects.CryptogateSendTransaction, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
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
		return
	}

	fmt.Printf("Balance: %v\n", res)
}
