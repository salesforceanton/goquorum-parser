package cryptogate

type NetworkSlug string

const (
	NetworkSlugGoQuorum NetworkSlug = "goquorum"
)

type Block struct {
	ID          uint64      `gorm:"column:id"`
	Height      uint64      `gorm:"column:height"`
	NetworkSlug NetworkSlug `gorm:"column:network_slug"`
}
