package cryptogate

type SmartContract struct {
	Address string
	Type    SmartContractType
}

type SmartContractType string

const (
	SmartContractTypeSTBL SmartContractType = "STBL"
	SmartContractTypeUSDT SmartContractType = "USDT"
)

func GetABI(contractType SmartContractType) string {
	switch contractType {
	case SmartContractTypeSTBL:
		return ""
	case SmartContractTypeUSDT:
		return ""
	default:
		return ""
	}
}
