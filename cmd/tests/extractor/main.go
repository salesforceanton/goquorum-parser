package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethereumextractor "github.com/salesforceanton/goquorum-parser/internal/ethereum_extractor"
)

// test Ethereum extractor in real environment.
func main() {
	testExtractor()
}

func testExtractor() {
	ex := ethereumextractor.NewExtractor(
		slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})),
		time.Minute,
	)

	logs := make(chan types.Log)

	ex.SetChannel(logs)
	ex.SetConnectionOptions(
		"",
		"",
	)
	ex.SetConnectionPolicy(
		ethereumextractor.ConnectionPolicy{
			HTTPTimeout:  time.Second * 2, // 1 second is too fast for alchemy
			RetryTimeout: time.Second,
		},
	)
	ex.SetBlockchainOptions(
		0,
		ethereumextractor.ContractMap{
			Addresses: []common.Address{
				common.HexToAddress(""),
			},
		},
	)

	go func() {
		for range logs {
			i := 0
			for log := range logs {
				i++
				fmt.Printf("msg %d, height=%d\n", i, log.BlockNumber)
			}
		}
	}()

	go ex.Start()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	ex.Stop()
	close(logs)
}
