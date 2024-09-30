// this package contains helper utils
package cryptogateutils

import (
	"time"

	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
)

type CryptogateUtils struct {
	network           cryptogate.Network
	smartContracts    map[cryptogate.SmartContractType]cryptogate.SmartContract
	rpcRequestTimeout time.Duration
}

func NewCryptogateUtils(
	network cryptogate.Network,
	smartContracts map[cryptogate.SmartContractType]cryptogate.SmartContract,
	rpcRequestTimeout time.Duration,
) *CryptogateUtils {
	return &CryptogateUtils{
		network:           network,
		smartContracts:    smartContracts,
		rpcRequestTimeout: rpcRequestTimeout,
	}
}
