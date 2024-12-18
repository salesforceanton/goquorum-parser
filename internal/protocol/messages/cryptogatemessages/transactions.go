package cryptogatemessages

import "github.com/salesforceanton/goquorum-parser/domain/cryptogate"

type SendTransactionRequest struct {
	TypeSmartContract cryptogate.SmartContractType `json:"typeSmartContract"`
	NameFunction      string                       `json:"nameFunction"`
	UserAddress       string                       `json:"userAddress"`
	Sender            string                       `json:"sender"`
	Amount            string                       `json:"amount"`
	PrivateFor        []string                     `json:"privateFor"`
	PrivateFrom       string                       `json:"privateFrom"`
}

type SendTransactionResponse struct {
}
