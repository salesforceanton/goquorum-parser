package ethereumextractor

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
	"github.com/salesforceanton/goquorum-parser/internal/timeutc"
)

// Alchemy limits amount of blocks in response. We need respect this limit.
var batchSize = big.NewInt(2000)

type ContractMap struct {
	Contracts map[string]cryptogate.SmartContract
	Addresses []common.Address
}

type ConnectionPolicy struct {
	// For Alchemy we need 2 seconds.
	HTTPTimeout time.Duration
	// Timeout between reconnect retries.
	RetryTimeout time.Duration
}

type concreteExtractor struct {
	httpUrl           string
	wsUrl             string
	policy            ConnectionPolicy
	currentHeight     uint64
	contractMap       ContractMap
	outerLogs         chan types.Log // outer channel
	logger            *slog.Logger
	rpcRequestTimeout time.Duration

	httpClient *ethclient.Client
	wsClient   *ethclient.Client
	chLogs     chan types.Log // internal channel to catch block number
	sub        ethereum.Subscription
	exit       chan struct{}
	lastRetry  time.Time
}

func NewExtractor(
	logger *slog.Logger,
	rpcRequestTimeout time.Duration,
) Extractor {
	return &concreteExtractor{
		chLogs:            make(chan types.Log),
		logger:            logger,
		rpcRequestTimeout: rpcRequestTimeout,
		exit:              make(chan struct{}),
	}
}

func (ex *concreteExtractor) SetConnectionOptions(httpUrl, wsUrl string) {
	ex.httpUrl = httpUrl
	ex.wsUrl = wsUrl
}

func (ex *concreteExtractor) SetConnectionPolicy(policy ConnectionPolicy) {
	ex.policy = policy
}

func (ex *concreteExtractor) SetBlockchainOptions(currentHeight uint64, contracts ContractMap) {
	ex.currentHeight = currentHeight
	ex.contractMap = contracts
}

func (ex *concreteExtractor) SetChannel(outerLogs chan types.Log) {
	ex.outerLogs = outerLogs
}

func (ex *concreteExtractor) Start() {
	go ex.catchHeight()

	for {
		select {
		case <-ex.exit:
			ex.logger.Debug("Extractor is stopping")
			close(ex.chLogs)
		default:
			if time.Since(ex.lastRetry) < ex.policy.RetryTimeout {
				continue
			}

			ex.close()

			err := ex.reconnect()
			if err != nil {
				ex.logger.Error("Failed to reconnect", "err", err)
				continue
			}

			bcHeight, err := ex.httpClient.BlockNumber(context.Background())
			if err != nil {
				ex.logger.Error("Failed to get current height", "err", err)
				continue
			}
			if bcHeight > ex.currentHeight {
				err = ex.parserHistoryEvents(
					ex.currentHeight, // we can have unprocessed logs at current height
					bcHeight,
				)
				if err != nil {
					ex.logger.Error("Failed to parse history events", "err", err)
					continue
				}
			}

			filter := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(ex.currentHeight) + 1),
				Addresses: ex.contractMap.Addresses,
			}

			ex.sub, err = ex.wsClient.SubscribeFilterLogs(context.Background(), filter, ex.chLogs)
			if err != nil {
				ex.logger.Error("Failed to subscribe for logs", "err", err)
				continue
			}

			err = <-ex.sub.Err()
			if err != nil {
				ex.logger.Error("Subscription error", "err", err)
			}
			ex.sub.Unsubscribe()
		}
	}
}

func (ex *concreteExtractor) Stop() {
	close(ex.exit)
	if ex.sub != nil {
		ex.sub.Unsubscribe()
	}
}

func (ex *concreteExtractor) reconnect() error {
	var err error

	ex.lastRetry = timeutc.Now()

	{
		ctx, cancel := context.WithTimeout(context.Background(), ex.policy.HTTPTimeout)
		defer cancel()

		rpcClient, err := rpc.DialOptions(
			ctx,
			ex.httpUrl,
			rpc.WithHTTPClient(&http.Client{
				Timeout: ex.rpcRequestTimeout,
			}),
		)
		if err != nil {
			ex.close()
			ex.logger.Error(fmt.Sprintf("error with http connect %s", ex.httpUrl))
			return err
		}

		ex.httpClient = ethclient.NewClient(rpcClient)
		ex.logger.Debug(fmt.Sprintf("successfully connected by http to %s", ex.httpUrl))
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), ex.policy.HTTPTimeout)
		defer cancel()

		ex.wsClient, err = ethclient.DialContext(ctx, ex.wsUrl)
		if err != nil {
			ex.close()
			ex.logger.Error(fmt.Sprintf("error with websocket connect %s", ex.wsUrl))
			return err
		}
		ex.logger.Debug(fmt.Sprintf("successfully connected by websocket to %s", ex.wsUrl))
	}

	return nil
}

func (ex *concreteExtractor) close() {
	if ex.httpClient != nil {
		ex.httpClient.Close()
	}
	if ex.wsClient != nil {
		ex.wsClient.Close()
	}
}

func (ex *concreteExtractor) parserHistoryEvents(
	fromHeight uint64,
	toHeight uint64,
) error {
	currentBlock := big.NewInt(int64(fromHeight))
	toBlock := big.NewInt(int64(toHeight))

	for currentBlock.Cmp(toBlock) <= 0 {
		endBlock := new(big.Int).Add(currentBlock, batchSize)
		if endBlock.Cmp(toBlock) > 0 {
			endBlock = toBlock
		}

		query := ethereum.FilterQuery{
			Addresses: ex.contractMap.Addresses,
			FromBlock: currentBlock,
			ToBlock:   endBlock,
		}

		ex.logger.Debug("Parse history events",
			slog.Group("range",
				"from", currentBlock.Uint64(),
				"to", endBlock.Uint64(),
			))

		logs, err := ex.httpClient.FilterLogs(context.Background(), query)
		if err != nil {
			ex.logger.Error("Failed to filter logs", "err", err)
			return err
		}

		for _, log := range logs {
			ex.chLogs <- log
		}

		currentBlock = new(big.Int).Add(endBlock, big.NewInt(1))
	}

	return nil
}

func (ex *concreteExtractor) catchHeight() {
	for log := range ex.chLogs {
		ex.outerLogs <- log
		ex.currentHeight = log.BlockNumber
	}
}
