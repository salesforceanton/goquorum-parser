package cryptogate

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/salesforceanton/goquorum-parser/internal/protocol/messages/cryptogatemessages"
)

func (c *Cryptogate) GetBalanceNative(
	ctx context.Context,
	req cryptogatemessages.BalanceNativeRequest,
) (cryptogatemessages.BalanceNativeResponse, error) {
	var result cryptogatemessages.BalanceNativeResponse

	address := common.HexToAddress(req.Address)

	balance, err := c.httpProvider.BalanceAt(ctx, address, nil)
	if err != nil {
		return cryptogatemessages.BalanceNativeResponse{}, err
	}

	result.Balance = balance.String()

	return result, nil
}
