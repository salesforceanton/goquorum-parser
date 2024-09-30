package repository

import (
	"errors"

	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
	"github.com/salesforceanton/goquorum-parser/internal/timeutc"
	"gorm.io/gorm"
)

func (r *Repository) GetBlock() (cryptogate.Block, error) {
	var res cryptogate.Block

	err := r.db.Table(`cryptogate.blocks`).First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cryptogate.Block{}, ErrBlockNotFound
		}
		return cryptogate.Block{}, err
	}

	return res, nil
}

func (r *Repository) UpdateHeight(height uint64) error {
	updateFields := map[string]interface{}{
		"height": timeutc.Now(),
	}
	return r.db.Table("cryptogate.blocks").
		Where("true").
		Updates(updateFields).Error
}
