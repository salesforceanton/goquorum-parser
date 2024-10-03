package cryptogateutils

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
)

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

	rpcClient, err := rpc.DialOptions(
		context.Background(),
		url,
		rpc.WithHTTPClient(&http.Client{
			Timeout: c.rpcRequestTimeout,
		}),
	)
	if err != nil {
		return &ethclient.Client{}, err
	}

	connection := ethclient.NewClient(rpcClient)

	return connection, nil
}

func (c *CryptogateUtils) InitTools(
	ctx context.Context,
	contractType cryptogate.SmartContractType,
) (Tools, error) {
	contract, ok := c.smartContracts[contractType]
	if !ok {
		return Tools{}, errors.New(fmt.Sprintf("smart contract not defined - %s", contractType))
	}

	rpcClient, err := rpc.DialOptions(
		ctx,
		c.network.Settings.RPCUrlHttp,
		rpc.WithHTTPClient(&http.Client{
			Timeout: c.rpcRequestTimeout,
		}),
	)
	if err != nil {
		return Tools{}, fmt.Errorf("rpc.DialOptions: %w", err)
	}

	provider := ethclient.NewClient(rpcClient)

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
