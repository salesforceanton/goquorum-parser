// Package with all root logic.
package cryptogate

import (
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
	utils "github.com/salesforceanton/goquorum-parser/internal/cryptogate_utils"
	ethereumextractor "github.com/salesforceanton/goquorum-parser/internal/ethereum_extractor"
	"github.com/spf13/viper"
)

const (
	ServiceName           = "cryptogate"
	parserEventsTick      = 3 * time.Second
	defaultRequestTimeout = 15 * time.Second
)

type Cryptogate struct {
	db              DB
	client          MsgTransport
	logger          *slog.Logger
	cryptogateUtils *utils.CryptogateUtils

	network cryptogate.Network

	privateKeyServiceBackend string
	contractAddresses        map[string]string
	rpcURLs                  map[cryptogate.TypeConnection]string

	deployBlock uint64

	httpRequestClient *http.Client

	// extraction
	extractor         ethereumextractor.Extractor
	contracts         *ethereumextractor.ContractMap
	ethLogs           chan types.Log
	lastBlock         uint64
	httpProvider      *ethclient.Client
	rpcRequestTimeout time.Duration

	wg   sync.WaitGroup
	exit chan struct{}
}

func New(
	db DB,
	client MsgTransport,
) *Cryptogate {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
	})).With("service", ServiceName)

	privateKeyServiceBackend := viper.GetString("PRIVATE_KEY_SERVICE_BACKEND")

	// Network parameters
	deployBlock := viper.GetUint64("DEPLOY_BLOCK")

	rpcURLs := make(map[cryptogate.TypeConnection]string)

	rpcURLs[cryptogate.TypeConnectionHTTP] = viper.GetString("RPC_URL_HTTP")
	rpcURLs[cryptogate.TypeConnectionWS] = viper.GetString("RPC_URL_WS")
	rpcRequestTimeout := viper.GetDuration("RPC_REQUEST_TIMEOUT")

	if rpcRequestTimeout == 0 {
		rpcRequestTimeout = defaultRequestTimeout
	}

	ex := ethereumextractor.NewExtractor(logger, rpcRequestTimeout)

	// Contract addresses
	contractAddresses := make(map[string]string)
	//contractAddresses[cryptogate.SmartContractTypeSTBL] = viper.GetString("STBL")

	// Set network settings here
	network := cryptogate.Network{
		Slug:   cryptogate.NetworkSlugGoQuorum,
		Status: cryptogate.NetworkStatusActive,
		Settings: cryptogate.NetworkSettings{
			RPCUrlHttp: "",
			RPCUrlWs:   "",
		},
	}

	// Set contracts here
	smartContracts := []cryptogate.SmartContract{}

	contractMap := &ethereumextractor.ContractMap{
		Contracts: make(map[string]cryptogate.SmartContract),
		Addresses: make([]common.Address, 0),
	}
	contractTypeMap := make(map[cryptogate.SmartContractType]cryptogate.SmartContract)

	for _, contract := range smartContracts {
		contractMap.Contracts[strings.ToLower(contract.Address)] = contract
		contractMap.Addresses = append(contractMap.Addresses, common.HexToAddress(contract.Address))
		contractTypeMap[contract.Type] = contract
	}

	// cryptogate utils
	cryptogateUtils := utils.NewCryptogateUtils(
		network,
		contractTypeMap,
		defaultRequestTimeout,
	)

	return &Cryptogate{
		db:                       db,
		client:                   client,
		network:                  network,
		privateKeyServiceBackend: privateKeyServiceBackend,
		httpRequestClient:        &http.Client{Timeout: defaultRequestTimeout},
		extractor:                ex,
		contracts:                contractMap,
		ethLogs:                  make(chan types.Log),
		rpcRequestTimeout:        rpcRequestTimeout,
		contractAddresses:        contractAddresses,
		rpcURLs:                  rpcURLs,
		deployBlock:              deployBlock,
		logger:                   logger,
		cryptogateUtils:          cryptogateUtils,
	}
}

func (c *Cryptogate) Start() error {
	c.exit = make(chan struct{})

	err := c.initDB()
	if err != nil {
		c.logger.Error("initDB", "err", err)
		return err
	}

	//////////
	// Extractor section
	blockDB, err := c.db.GetBlock()
	if err != nil {
		c.logger.Error("db.GetBlock", "err", err)
		c.lastBlock = 0
	} else {
		c.lastBlock = blockDB.Height
	}

	networkSettings := c.network.Settings

	c.extractor.SetConnectionOptions(networkSettings.RPCUrlHttp, networkSettings.RPCUrlWs)
	c.extractor.SetConnectionPolicy(
		ethereumextractor.ConnectionPolicy{
			HTTPTimeout:  time.Second * 10,
			RetryTimeout: time.Second,
		},
	)
	c.extractor.SetBlockchainOptions(
		blockDB.Height,
		*c.contracts,
	)
	c.extractor.SetChannel(c.ethLogs)

	go c.processLogs()
	go c.extractor.Start()

	// others
	c.httpProvider, err = c.cryptogateUtils.InitProvider(cryptogate.TypeConnectionHTTP)
	if err != nil {
		c.logger.Error("cryptogateUtils.InitProvider", "err", err)
		return err
	}

	err = c.client.Start(c)
	if err != nil {
		c.logger.Error("client.Start", "err", err)
		return err
	}

	c.logger.Debug("Start module")

	return nil
}

func (c *Cryptogate) Stop() error {
	close(c.exit)
	c.extractor.Stop()
	close(c.ethLogs)
	c.client.Stop()
	c.wg.Wait()

	c.logger.Debug("Stop module")
	return nil
}

func (c *Cryptogate) initDB() error {
	// Set default deployment block
	block, err := c.db.GetBlock()
	if err != nil {
		c.logger.Debug("db.GetBlock", "err", err)
		return err
	}
	if block.Height == 0 {
		err = c.db.UpdateHeight(c.deployBlock)
		if err != nil {
			c.logger.Debug("db.UpdateHeight", "err", err)
			return err
		}
	}

	return nil
}
