package cryptogate

type Event struct {
	ID            string       `gorm:"column:id"`
	LogIndex      uint64       `gorm:"column:log_index"`
	Height        uint64       `gorm:"column:height"`
	Hash          string       `gorm:"column:hash"`
	From          string       `gorm:"column:from"`
	SmartContract string       `gorm:"column:smart_contract"`
	Amount        string       `gorm:"column:amount"`
	CoinSlug      string       `gorm:"column:coin_slug"`
	Type          FunctionType `gorm:"column:type"`
	Status        EventStatus  `gorm:"column:status"`
	GasUsed       string       `gorm:"column:gas_used"`
	CreatedAt     uint64       `gorm:"column:created_at"`
}

type FunctionType string

const (
	FunctionTypeTransfer       FunctionType = "transfer"
	FunctionTypeRequestUnknown FunctionType = "unknown"
)

type EventType string

const (
	EventTypeTransfer EventType = "Transfer"
)

type EventStatus string

const (
	EventStatusNew          EventStatus = "new"
	EventStatusInProcessing EventStatus = "inProcessing"
	EventStatusSuccess      EventStatus = "success"
	EventStatusFailed       EventStatus = "failed"
	EventStatusUnknown      EventStatus = "unknown"
)
