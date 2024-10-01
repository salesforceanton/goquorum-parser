package rpc

import (
	"github.com/salesforceanton/goquorum-parser/internal/protocol/messages/cryptogatemessages"
	"github.com/salesforceanton/goquorum-parser/internal/protocol/subjects"
	"github.com/salesforceanton/goquorum-parser/internal/transport"
)

var (
	CryptogateGetBalanceNativeRequester = transport.NewRequester[cryptogatemessages.BalanceNativeRequest, cryptogatemessages.BalanceNativeResponse](subjects.CryptogateSendTransaction, "CryptogateSendTransaction")
	CryptogateGetBalanceNativeResponder = transport.NewResponder[cryptogatemessages.BalanceNativeRequest, cryptogatemessages.BalanceNativeResponse](subjects.CryptogateSendTransaction, "CryptogateSendTransaction")

	CryptogateSendTransactionRequester = transport.NewRequester[cryptogatemessages.SendTransactionRequest, cryptogatemessages.SendTransactionResponse](subjects.CryptogateGetBalanceNative, "CryptogateSendTransaction")
	CryptogateSendTransactionResponder = transport.NewResponder[cryptogatemessages.SendTransactionRequest, cryptogatemessages.SendTransactionResponse](subjects.CryptogateGetBalanceNative, "CryptogateSendTransaction")
)
