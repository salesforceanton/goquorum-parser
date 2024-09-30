package rpc

import (
	"github.com/salesforceanton/goquorum-parser/internal/protocol/messages/cryptogatemessages"
	"github.com/salesforceanton/goquorum-parser/internal/protocol/subjects"
	"github.com/salesforceanton/goquorum-parser/internal/transport"
)

var (
	CryptogateSendTransactionRequester = transport.NewRequester[cryptogatemessages.SendTransactionRequest, cryptogatemessages.SendTransactionResponse](subjects.CryptogateSendTransaction, "CryptogateSendTransaction")
	CryptogateSendTransactionResponder = transport.NewResponder[cryptogatemessages.SendTransactionRequest, cryptogatemessages.SendTransactionResponse](subjects.CryptogateSendTransaction, "CryptogateSendTransaction")
)
