package cryptogate

import (
	"context"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
	"github.com/salesforceanton/goquorum-parser/internal/converters"
	cryptogateutils "github.com/salesforceanton/goquorum-parser/internal/cryptogate_utils"
	"github.com/salesforceanton/goquorum-parser/internal/protocol/messages/cryptogatemessages"
	"github.com/salesforceanton/goquorum-parser/services/backend/cryptogate/usdt_abi"
)

const (
	gasLimit = 47000000
)

func (c *Cryptogate) SendTransaction( //nolint:funlen,maintidx,gocyclo
	ctx context.Context,
	req cryptogatemessages.SendTransactionRequest,
) (cryptogatemessages.SendTransactionResponse, error) {
	tools, err := c.cryptogateUtils.InitTools(ctx, req.TypeSmartContract)
	if err != nil {
		c.logger.Debug("cryptogateUtils.InitTools", "err", err)
		return cryptogatemessages.SendTransactionResponse{}, err
	}

	addressContract := common.HexToAddress(tools.Contract.Address)

	gasPrice, err := tools.Provider.SuggestGasPrice(ctx)
	if err != nil {
		c.logger.Debug("Provider.SuggestGasPrice", "err", err)
		return cryptogatemessages.SendTransactionResponse{}, err
	}

	chainID, err := tools.Provider.ChainID(context.Background())
	if err != nil {
		c.logger.Debug("tools.Provider.ChainID", "err", err)
		return cryptogatemessages.SendTransactionResponse{}, err
	}

	trOpts, err := bind.NewTransactorWithChainID(strings.NewReader(c.ethKey), "", chainID)
	if err != nil {
		c.logger.Debug("bind.NewTransactor", "err", err)
		return cryptogatemessages.SendTransactionResponse{}, err
	}

	trOpts.GasLimit = gasLimit
	trOpts.PrivateFor = req.PrivateFor
	trOpts.PrivateFrom = req.PrivateFrom
	trOpts.IsUsingPrivacyPrecompile = true

	var transaction *types.Transaction

	switch req.TypeSmartContract {
	case cryptogate.SmartContractTypeUSDT:
		contractTransactor, err := usdt_abi.NewUsdtAbiTransactor(
			addressContract,
			tools.Provider,
		)
		if err != nil {
			c.logger.Error("usdt_abi.NewUsdtAbiTransactor", "err", err)
			return cryptogatemessages.SendTransactionResponse{}, err
		}

		switch req.NameFunction {
		case "issue":
			amount, err := converters.StringToBigInt(req.Amount)
			if err != nil {
				c.logger.Error("converters.StringToBigInt", "err", err)
				return cryptogatemessages.SendTransactionResponse{}, err
			}

			transaction, err = contractTransactor.Issue(trOpts, amount)
			if err != nil {
				c.logger.Error("contractTransactor.Issue", "err", err)
				return cryptogatemessages.SendTransactionResponse{}, err
			}
		case "transfer":
			amount, err := converters.StringToBigInt(req.Amount)
			if err != nil {
				c.logger.Error("converters.StringToBigInt", "err", err)
				return cryptogatemessages.SendTransactionResponse{}, err
			}

			transaction, err = contractTransactor.Transfer(trOpts, common.HexToAddress(req.UserAddress), amount)
			if err != nil {
				c.logger.Error("contractTransactor.Transfer", "err", err)
				return cryptogatemessages.SendTransactionResponse{}, err
			}
		}
	}

	c.logger.Debug("Transaction sent successfully",
		"function", req.NameFunction,
		"to", addressContract.Hex(),
		"gasPrice", gasPrice.String(),
		"txHash", transaction.Hash().Hex(),
	)

	return cryptogatemessages.SendTransactionResponse{}, nil //nolint:exhaustruct
}

func (c *Cryptogate) GetGasLimit(req cryptogatemessages.SendTransactionRequest,
	tools cryptogateutils.Tools,
	msg ethereum.CallMsg,
) (uint64, error) {
	gasLimit, err := tools.Provider.EstimateGas(context.Background(), msg)
	if err != nil {
		c.logger.Error("tools.Provider.EstimateGas", "err", err)

		return 0, errors.New("execution reverted")
	}

	return gasLimit, nil
}
