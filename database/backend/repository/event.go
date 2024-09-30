package repository

import (
	"github.com/salesforceanton/goquorum-parser/domain/cryptogate"
	"gorm.io/gorm/clause"
)

func (r *Repository) UpsertEvent(u *cryptogate.Event) (string, error) {
	result := r.db.Table(`cryptogate.events`).
		Clauses(clause.OnConflict{
			DoNothing: true,
		}).
		Create(u)

	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "", ErrEventHasAlreadyBeenProcessed
	}

	return u.ID, nil
}

func (r *Repository) GetEventsByStatusList(statuses []cryptogate.EventStatus) ([]cryptogate.Event, error) {
	var res []cryptogate.Event

	err := r.db.Table("cryptogate.events").
		Where("status IN (?)", statuses).
		Order("height ASC").
		Order("log_index ASC").
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Repository) UpdateEventStatus(id string, status cryptogate.EventStatus) error {
	return r.db.Table("cryptogate.events").
		Where("id = ?", id).
		Update("status", status).
		Error
}
