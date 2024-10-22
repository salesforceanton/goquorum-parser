package cryptogate

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
	"github.com/salesforceanton/goquorum-parser/internal/protocol/messages/cryptogatemessages"
	"github.com/salesforceanton/goquorum-parser/services/backend/cryptogate/usdt_abi"
)

func (c *Cryptogate) GetBalanceToken(
	ctx context.Context,
	req cryptogatemessages.BalanceTokenRequest,
) (cryptogatemessages.BalanceTokenResponse, error) {
	tools, err := c.cryptogateUtils.InitTools(ctx, req.TypeSmartContract)
	if err != nil {
		return cryptogatemessages.BalanceTokenResponse{}, err
	}

	contractAddr := common.HexToAddress(c.contractAddresses[req.TypeSmartContract])

	switch req.TypeSmartContract {
	case cryptogate.SmartContractTypeUSDT:
		contractCaller, err := usdt_abi.NewUsdtAbiCaller(
			contractAddr,
			tools.Provider,
		)
		if err != nil {
			c.logger.Error("usdt_abi.NewUsdtAbiCaller", "err", err)
			return cryptogatemessages.BalanceTokenResponse{}, err
		}

		balance, err := contractCaller.BalanceOf(nil, common.HexToAddress(req.Address))
		if err != nil {
			c.logger.Error("contractCaller.BalanceOf", "err", err)
			return cryptogatemessages.BalanceTokenResponse{}, err
		}

		return cryptogatemessages.BalanceTokenResponse{
			Balance: balance.String(),
		}, nil
	default:
		return cryptogatemessages.BalanceTokenResponse{}, fmt.Errorf("wrong type of token contract")
	}
}
