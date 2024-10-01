package cryptogatemessages

import "github.com/salesforceanton/goquorum-parser/domain/cryptogate"

type BalanceTokenRequest struct {
	ContractAddress   string                       `json:"contractAddress"`
	Address           string                       `json:"address"`
	TypeSmartContract cryptogate.SmartContractType `json:"typeSmartContract"`
}

type BalanceTokenResponse struct {
	Balance string `json:"balance"`
}

type BalanceNativeRequest struct {
	Address string `json:"address"`
}

type BalanceNativeResponse struct {
	Balance string `json:"balance"`
}
