package cryptogate

import (
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Cryptogate) parseEvent(
	log types.Log,
) {
	contractAddress := strings.ToLower(log.Address.Hex())
	if contract, ok := c.contracts.Contracts[contractAddress]; !ok {
		c.logger.Debug("Contract record is not initialized")
		return
	} else {
		switch contract.Type {
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
