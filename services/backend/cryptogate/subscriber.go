package cryptogate

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
)

func (c *Cryptogate) parseEvent(
	log types.Log,
) {
	contractAddress := strings.ToLower(log.Address.Hex())
	if contract, ok := c.contracts.Contracts[contractAddress]; !ok {
		c.logger.Debug(fmt.Sprintf("Contract record is not initialized: %s", contractAddress))
		return
	} else {
		switch contract.Type {
		case cryptogate.SmartContractTypePermissionImpl,
			cryptogate.SmartContractTypeOrgManager,
			cryptogate.SmartContractTypeAccountManager,
			cryptogate.SmartContractTypeNodeManager,
			cryptogate.SmartContractTypeRoleManager,
			cryptogate.SmartContractTypeVoterManager:

			preData, err := c.PrepareParseEvents(c.httpProvider, log, contract)
			if err != nil {
				c.logger.Error(fmt.Sprintf("PrepareParseEvents: %s", err.Error()))
			}

			_, err = c.db.UpsertEvent(preData.DataEvent)
			if err != nil {
				c.logger.Error(fmt.Sprintf("db.UpsertEvent: %s", err.Error()))
			}
		default:
			c.logger.Error("Unknown contract type", "type", contract.Type)
		}
	}
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
		c.parseEvent(log)
	}
}
