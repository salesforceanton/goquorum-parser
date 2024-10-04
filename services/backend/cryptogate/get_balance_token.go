package cryptogate

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
	"github.com/salesforceanton/goquorum-parser/internal/protocol/messages/cryptogatemessages"
)

func (c *Cryptogate) GetBalanceToken(
	ctx context.Context,
	req cryptogatemessages.BalanceTokenRequest,
) (cryptogatemessages.BalanceTokenResponse, error) {
	var result cryptogatemessages.BalanceTokenResponse

	tools, err := c.cryptogateUtils.InitTools(ctx, req.TypeSmartContract)
	if err != nil {
		return cryptogatemessages.BalanceTokenResponse{}, err
	}

	contractAddr := common.HexToAddress(c.contractAddresses[req.TypeSmartContract])

	parsedABI, err := abi.JSON(strings.NewReader(cryptogate.GetABI(tools.Contract.Type)))
	if err != nil {
		c.logger.Error("abi.JSON", "err", err)
		return cryptogatemessages.BalanceTokenResponse{}, err
	}

	balanceData, err := parsedABI.Pack("balanceOf", common.HexToAddress(req.Address))
	if err != nil {
		c.logger.Error("parsedABI.Pack", "err", err)
		return cryptogatemessages.BalanceTokenResponse{}, err
	}

	callMsg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: balanceData,
	}

	balance, err := tools.Provider.CallContract(ctx, callMsg, big.NewInt(int64(tools.BlockNumber)))
	if err != nil {
		c.logger.Error("tools.Provider.CallContract", "err", err)
		return cryptogatemessages.BalanceTokenResponse{}, err
	}

	balanceStr, err := parsedABI.Unpack("balanceOf", balance)
	if err != nil {
		c.logger.Error("parsedABI.Unpack", "err", err)
		return cryptogatemessages.BalanceTokenResponse{}, err
	}
	result.Balance = balanceStr[0].(*big.Int).String()

	return result, nil
}
