package cryptogate

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
	"github.com/salesforceanton/goquorum-parser/internal/timeutc"
)

type PrepareParseEventsResult struct {
	Event     *abi.Event
	Topics    []common.Hash
	DataEvent *cryptogate.Event
}

func (c *Cryptogate) PrepareParseEvents(
	provider *ethclient.Client,
	log types.Log,
	contract cryptogate.SmartContract,
) (PrepareParseEventsResult, error) {
	receipt, err := provider.TransactionReceipt(context.Background(), log.TxHash)
	if err != nil {
		c.logger.Debug(fmt.Sprintf("provider.TransactionReceipt: %s", err.Error()))

		return PrepareParseEventsResult{}, err
	}

	tx, _, err := provider.TransactionByHash(context.Background(), log.TxHash)
	if err != nil {
		c.logger.Debug(fmt.Sprintf("provider.TransactionByHash: %s", err.Error()))

		return PrepareParseEventsResult{}, err
	}

	headers, err := provider.HeaderByHash(context.Background(), receipt.BlockHash)
	if err != nil {
		c.logger.Debug(fmt.Sprintf("Failed to get header by hash: %s", err.Error()))

		return PrepareParseEventsResult{}, err
	}

	sender, err := provider.TransactionSender(context.Background(), tx, receipt.BlockHash, receipt.TransactionIndex)
	if err != nil {
		c.logger.Debug(fmt.Sprintf("Failed to get transaction sender: %s", err.Error()))

		return PrepareParseEventsResult{}, err
	}

	topics := log.Topics
	if len(topics) < 1 {
		return PrepareParseEventsResult{}, errors.New("incorrect data topics")
	}

	parsedABI, err := abi.JSON(strings.NewReader(cryptogate.GetABI(contract.Type)))
	if err != nil {
		c.logger.Debug(fmt.Sprintf("abi.JSON: %s", err.Error()))
	}

	event, err := parsedABI.EventByID(common.HexToHash(topics[0].Hex()))
	if err != nil {
		c.logger.Debug(
			"parsedABI.EventByID",
			"id", common.HexToHash(topics[0].Hex()),
			"hash", log.TxHash.Hex(),
			"error", err.Error(),
		)
		return PrepareParseEventsResult{}, err
	}

	gasUsed := big.NewInt(0).SetUint64(receipt.GasUsed)
	// gasUsed := new(big.Int)
	// if receipt.EffectiveGasPrice != nil {
	// 	gasUsed = gasUsed.Mul(gas, receipt.EffectiveGasPrice)
	// }

	dataEvent := &cryptogate.Event{
		ID:                uuid.NewString(),
		LogIndex:          uint64(log.Index),
		From:              strings.ToLower(sender.String()),
		SmartContract:     strings.ToLower(contract.Address),
		SmartContractType: contract.Type,
		Type:              event.Name,
		Hash:              log.TxHash.Hex(),
		Height:            log.BlockNumber,
		Status:            cryptogate.EventStatusNew,
		GasUsed:           gasUsed.String(),
		CreatedAt:         timeutc.Unix(int64(headers.Time)),
	}

	return PrepareParseEventsResult{
		Event:     event,
		Topics:    topics,
		DataEvent: dataEvent,
	}, nil
}
