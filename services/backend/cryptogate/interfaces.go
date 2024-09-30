package cryptogate

import "github.com/salesforceanton/goquorum-parser/domain/cryptogate"

type DB interface {
	GetBlock() (cryptogate.Block, error)
	UpdateHeight(uint64) error

	UpsertEvent(u *cryptogate.Event) (string, error)
}
