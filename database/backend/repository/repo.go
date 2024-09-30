package repository

import (
	"github.com/salesforceanton/goquorum-parser/internal/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(cfg database.PSQLConfig) (*Repository, error) {
	db, err := database.NewPSQL(cfg)
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

func (r *Repository) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
