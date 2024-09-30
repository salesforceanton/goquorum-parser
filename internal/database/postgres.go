package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PSQLConfig struct {
	Host     string
	User     string
	Pass     string
	Database string
	Port     uint32
	// MaxConnections defines database/sql SetMaxOpenConns, SetMaxIdleConns
	MaxConnections int
}

const (
	insertBatchSize = 1000
	defaultMaxConn  = 20
)

func NewPSQL(config PSQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		config.Host,
		config.User,
		config.Pass,
		config.Database,
		config.Port,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	// PostgreSQL has limit: "extended protocol limited to 65535 parameters"
	// That mean: (insert tuple columnt count) * (amount of tuples) must be <= 65535
	// 1000 is enough for tables with 65 columns
	c, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:          newLogger,
		CreateBatchSize: insertBatchSize,
	})
	if err != nil {
		return nil, err
	}

	db, err := c.DB()
	if err != nil {
		return nil, err
	}

	maxConn := config.MaxConnections
	if maxConn <= 0 {
		maxConn = defaultMaxConn
	}

	db.SetMaxIdleConns(maxConn)
	db.SetMaxOpenConns(maxConn)
	db.SetConnMaxIdleTime(time.Hour)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return c, nil
}
