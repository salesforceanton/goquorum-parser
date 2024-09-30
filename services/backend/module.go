package backend

import (
	"context"
	"log/slog"
	"os"

	repositoryCryptogate "github.com/salesforceanton/goquorum-parser/database/backend/repository"

	"github.com/salesforceanton/goquorum-parser/internal/config"
	"github.com/salesforceanton/goquorum-parser/internal/database"
	"github.com/salesforceanton/goquorum-parser/internal/transport"
	cryptogateservice "github.com/salesforceanton/goquorum-parser/services/backend/cryptogate"

	"github.com/spf13/viper"
)

const (
	ServiceName = "backend"
)

type Module struct {
	logger            *slog.Logger
	cryptogateService *cryptogateservice.Cryptogate
}

func New() (*Module, error) {
	config.LoadConfig()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})).With("service", ServiceName)

	backendDb, err := repositoryCryptogate.New(database.PSQLConfig{
		Host:           viper.GetString("PSQL_BACKEND_HOST"),
		User:           viper.GetString("PSQL_BACKEND_USER"),
		Pass:           viper.GetString("PSQL_BACKEND_PASS"),
		Database:       viper.GetString("PSQL_BACKEND_DB"),
		Port:           viper.GetUint32("PSQL_BACKEND_PORT"),
		MaxConnections: viper.GetInt("PSQL_BACKEND_MAX_CONN"),
	})
	if err != nil {
		return nil, err
	}

	backendNats, err := transport.New(transport.Config{
		Host:     viper.GetString("NATS_HOST"),
		Port:     viper.GetString("NATS_PORT"),
		User:     viper.GetString("NATS_USER"),
		Password: viper.GetString("NATS_PASS"),
	}, logger)
	if err != nil {
		return nil, err
	}

	natsRequestTimeout := viper.GetDuration("NATS_REQUEST_TIMEOUT")
	cryptogateTransport := cryptogateservice.NewTransport(backendNats, logger, natsRequestTimeout)

	cryptogateService := cryptogateservice.New(
		backendDb,
		cryptogateTransport,
	)

	return &Module{
		logger:            logger,
		cryptogateService: cryptogateService,
	}, nil
}

func (m *Module) Start(ctx context.Context) error {
	m.logger.Debug("Start module")

	// start cryptogate module
	err := m.cryptogateService.Start()
	if err != nil {
		return err
	}

	return nil
}

func (m *Module) Stop() error {
	// stop cryptogate module
	m.cryptogateService.Stop()

	return nil
}
