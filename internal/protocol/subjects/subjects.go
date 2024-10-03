package subjects

type Subject = string

const (
	CryptogateSendTransaction  Subject = "CRYPTOGATE.transaction.send"
	CryptogateGetBalanceNative Subject = "CRYPTOGATE.balance.native.get"
	CryptogateGetBalanceToken  Subject = "CRYPTOGATE.balance.token.get"
)
