package cryptogate

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
)

func (c *Cryptogate) parseEvent(
	log types.Log,
) error {
	contractAddress := strings.ToLower(log.Address.Hex())
	if contract, ok := c.contracts.Contracts[contractAddress]; !ok {
		c.logger.Debug(fmt.Sprintf("Contract record is not initialized: %s", contractAddress))
		return fmt.Errorf("contract record is not initialized: %s", contractAddress)
	} else {
		switch contract.Type {
		case cryptogate.SmartContractTypePermissionImpl,
			cryptogate.SmartContractTypeOrgManager,
			cryptogate.SmartContractTypeAccountManager,
			cryptogate.SmartContractTypeNodeManager,
			cryptogate.SmartContractTypeRoleManager,
			cryptogate.SmartContractTypeVoterManager,
			cryptogate.SmartContractTypeUSDT:

			preData, err := c.PrepareParseEvents(c.httpProvider, log, contract)
			if err != nil {
				c.logger.Error(fmt.Sprintf("PrepareParseEvents: %s", err.Error()))
				return err
			}

			_, err = c.db.UpsertEvent(preData.DataEvent)
			if err != nil {
				c.logger.Error(fmt.Sprintf("db.UpsertEvent: %s", err.Error()))
				return err
			}

			c.logger.Debug("event recieved",
				"type", preData.DataEvent.Type,
				"smart contract type", preData.DataEvent.SmartContractType,
				"contract address", preData.DataEvent.SmartContract,
				"from", preData.DataEvent.From,
				"block", preData.DataEvent.Height,
				"hash", preData.DataEvent.Hash,
				"gasUsed", preData.DataEvent.GasUsed,
			)
		default:
			c.logger.Error("Unknown contract type", "type", contract.Type)
		}
	}

	return nil
}

func (c *Cryptogate) updateHeight(blockNumber uint64) {
	if c.lastBlock >= blockNumber {
		return
	}

	c.logger.Debug("Block height from event", "block", blockNumber)

	err := c.db.UpdateHeight(blockNumber)
	if err != nil {
		c.logger.Error("Update height", "err", err)
		return
	}

	c.lastBlock = blockNumber
}

func (c *Cryptogate) processLogs() {
	for log := range c.ethLogs {
		c.updateHeight(log.BlockNumber)

		err := c.parseEvent(log)
		// get logs from http filter logs here if err with getting private tx receipt appeared
		if err != nil && errors.Is(err, ErrWithGetTransactionReceipt) {
			c.ParseBlockEvents(log)
		}
	}
}

func (c *Cryptogate) ParseBlockEvents(log types.Log) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{log.Address},
		FromBlock: big.NewInt(int64(log.BlockNumber)),
		ToBlock:   big.NewInt(int64(log.BlockNumber)),
	}

	logs, err := c.httpProvider.FilterLogs(context.Background(), query)
	if err != nil {
		c.logger.Error("httpProvider.FilterLogs", "err", err)
	}

	for _, log := range logs {
		c.parseEvent(log)
	}
}
