package cryptogateutils

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
)

const tessera1Url = "http://localhost:9081"

type Tools struct {
	Provider    *ethclient.Client
	Network     cryptogate.Network
	Contract    cryptogate.SmartContract
	BlockNumber uint64
}

func (c *CryptogateUtils) InitProvider(
	typeConnection cryptogate.TypeConnection,
) (*ethclient.Client, error) {
	provider, err := c.GetProvider(c.network.Settings, typeConnection)
	if err != nil {
		return &ethclient.Client{}, err
	}

	return provider, nil
}

func (c *CryptogateUtils) GetProvider(
	networkSettings cryptogate.NetworkSettings,
	typeconnection cryptogate.TypeConnection,
) (*ethclient.Client, error) {
	var url string
	if typeconnection == cryptogate.TypeConnectionHTTP {
		url = networkSettings.RPCUrlHttp
	} else if typeconnection == cryptogate.TypeConnectionWS {
		url = networkSettings.RPCUrlWs
	} else {
		return &ethclient.Client{}, fmt.Errorf("wrong connection type")
	}

	rpcClient, err := rpc.DialHTTPWithClient(
		url,
		&http.Client{
			Timeout: c.rpcRequestTimeout,
		},
	)
	if err != nil {
		return &ethclient.Client{}, err
	}

	return ethclient.NewClient(rpcClient).WithPrivateTransactionManager(tessera1Url)
}

func (c *CryptogateUtils) InitTools(
	ctx context.Context,
	contractType cryptogate.SmartContractType,
) (Tools, error) {
	contract, ok := c.smartContracts[contractType]
	if !ok {
		return Tools{}, fmt.Errorf(fmt.Sprintf("smart contract not defined - %s", contractType))
	}

	rpcClient, err := rpc.DialHTTPWithClient(
		c.network.Settings.RPCUrlHttp,
		&http.Client{
			Timeout: c.rpcRequestTimeout,
		},
	)
	if err != nil {
		return Tools{}, fmt.Errorf("rpc.DialOptions: %w", err)
	}

	provider, err := ethclient.NewClient(rpcClient).WithPrivateTransactionManager(tessera1Url)
	if err != nil {
		return Tools{}, fmt.Errorf("ethclient.NewClient(rpcClient).WithPrivateTransactionManager(tessera1Url): %w", err)
	}

	blockNumber, err := provider.BlockNumber(ctx)
	if err != nil {
		return Tools{}, err
	}

	return Tools{
		Provider:    provider,
		Network:     c.network,
		Contract:    contract,
		BlockNumber: blockNumber,
	}, nil
}
