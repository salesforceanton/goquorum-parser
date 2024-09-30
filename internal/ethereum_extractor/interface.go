package ethereumextractor

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type Extractor interface {
	SetConnectionOptions(httpUrl, wsUrl string)
	SetConnectionPolicy(policy ConnectionPolicy)
	SetBlockchainOptions(currentHeight uint64, contracts ContractMap)
	SetChannel(outerLogs chan types.Log)
	Start()
	Stop()
}
