package rpc

import (
	"github.com/salesforceanton/goquorum-parser/internal/protocol/messages/cryptogatemessages"
	"github.com/salesforceanton/goquorum-parser/internal/protocol/subjects"
	"github.com/salesforceanton/goquorum-parser/internal/transport"
)

var (
	CryptogateGetBalanceNativeRequester = transport.NewRequester[cryptogatemessages.BalanceNativeRequest, cryptogatemessages.BalanceNativeResponse](subjects.CryptogateGetBalanceNative, "CryptogateGetBalanceNative")
	CryptogateGetBalanceNativeResponder = transport.NewResponder[cryptogatemessages.BalanceNativeRequest, cryptogatemessages.BalanceNativeResponse](subjects.CryptogateGetBalanceNative, "CryptogateGetBalanceNative")

	CryptogateGetBalanceTokenRequester = transport.NewRequester[cryptogatemessages.BalanceTokenRequest, cryptogatemessages.BalanceTokenResponse](subjects.CryptogateGetBalanceToken, "CryptogateGetBalanceToken")
	CryptogateGetBalanceTokenResponder = transport.NewResponder[cryptogatemessages.BalanceTokenRequest, cryptogatemessages.BalanceTokenResponse](subjects.CryptogateGetBalanceToken, "CryptogateGetBalanceToken")

	CryptogateSendTransactionRequester = transport.NewRequester[cryptogatemessages.SendTransactionRequest, cryptogatemessages.SendTransactionResponse](subjects.CryptogateSendTransaction, "CryptogateSendTransaction")
	CryptogateSendTransactionResponder = transport.NewResponder[cryptogatemessages.SendTransactionRequest, cryptogatemessages.SendTransactionResponse](subjects.CryptogateSendTransaction, "CryptogateSendTransaction")
)
