package cryptogate

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
	"github.com/salesforceanton/goquorum-parser/internal/converters"
	cryptogateutils "github.com/salesforceanton/goquorum-parser/internal/cryptogate_utils"
	"github.com/salesforceanton/goquorum-parser/internal/protocol/messages/cryptogatemessages"
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

	privateKey, err := crypto.HexToECDSA(c.privateKeyServiceBackend)
	if err != nil {
		c.logger.Debug("Error with encoding private service key", "err", err)
		return cryptogatemessages.SendTransactionResponse{}, err
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return cryptogatemessages.SendTransactionResponse{}, errors.New("invalid public key")
	}

	var sender common.Address

	if req.Sender == "" {
		sender = crypto.PubkeyToAddress(*publicKeyECDSA)
	} else {
		sender = common.HexToAddress(req.Sender)
	}

	addressContract := common.HexToAddress(tools.Contract.Address)

	value := big.NewInt(0) // Set the value to 0 if you're not sending ETH

	gasPrice, err := tools.Provider.SuggestGasPrice(ctx)
	if err != nil {
		c.logger.Debug("Provider.SuggestGasPrice", "err", err)
		return cryptogatemessages.SendTransactionResponse{}, err
	}

	nonce, err := tools.Provider.PendingNonceAt(ctx, sender)
	if err != nil {
		c.logger.Debug("Provider.PendingNonceAt", "err", err)
		return cryptogatemessages.SendTransactionResponse{}, err
	}

	var tx *types.Transaction

	// Set the method and parameters
	contractAbi, err := abi.JSON(strings.NewReader(cryptogate.GetABI(tools.Contract.Type)))
	if err != nil {
		c.logger.Debug("abi.JSON", "err", err)
		return cryptogatemessages.SendTransactionResponse{}, err
	}

	var data []byte

	switch req.TypeSmartContract {
	case cryptogate.SmartContractTypeUSDT:
		switch req.NameFunction {
		case "issue":
			amount, err := converters.StringToBigInt(req.Amount)
			if err != nil {
				c.logger.Debug("converters.StringToBigInt", "err", err)
				return cryptogatemessages.SendTransactionResponse{}, err
			}

			data, err = contractAbi.Pack(
				req.NameFunction,
				amount,
			)
			if err != nil {
				c.logger.Debug("contractAbi.Pack", "err", err)
				return cryptogatemessages.SendTransactionResponse{}, err
			}
		case "transfer":
			amount, err := converters.StringToBigInt(req.Amount)
			if err != nil {
				c.logger.Debug("converters.StringToBigInt", "err", err)
				return cryptogatemessages.SendTransactionResponse{}, err
			}
			data, err = contractAbi.Pack(
				req.NameFunction,
				common.HexToAddress(req.UserAddress),
				amount,
			)
			if err != nil {
				c.logger.Debug("contractAbi.Pack", "err", err)
				return cryptogatemessages.SendTransactionResponse{}, err
			}
		}
	}

	msg := ethereum.CallMsg{
		From:     sender,
		To:       &addressContract,
		GasPrice: gasPrice,
		Value:    value,
		Data:     data,
	}

	gasLimit, err := c.GetGasLimit(req, tools, msg)
	if err != nil {
		c.logger.Debug("GetGasLimit", "err", err)
		return cryptogatemessages.SendTransactionResponse{}, err
	}

	tx = types.NewTransaction(
		nonce,
		addressContract,
		value,
		gasLimit,
		gasPrice,
		data,
	)

	chainID, err := tools.Provider.ChainID(context.Background())
	if err != nil {
		c.logger.Debug("tools.Provider.ChainID", "err", err)
		return cryptogatemessages.SendTransactionResponse{}, err
	}

	// Sign the transaction
	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		c.logger.Debug("bind.NewKeyedTransactorWithChainID", "err", err)
		return cryptogatemessages.SendTransactionResponse{}, err
	}

	signedTx, err := signer.Signer(sender, tx)
	if err != nil {
		c.logger.Debug("signer.Signer", "err", err)
		return cryptogatemessages.SendTransactionResponse{}, err
	}

	// Send the transaction
	err = tools.Provider.SendTransaction(context.Background(), signedTx)
	if err != nil {
		c.logger.Debug("tools.Provider.SendTransaction", "err", err)
		return cryptogatemessages.SendTransactionResponse{}, err
	}

	c.logger.Debug("Transaction sent successfully",
		"function", req.NameFunction,
		"to", addressContract.Hex(),
		"nonce", nonce,
		"amount", value.String(),
		"gasPrice", gasPrice.String(),
		"txHash", tx.Hash().Hex(),
		"signedTxHash", signedTx.Hash().Hex(),
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
