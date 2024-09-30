package cryptogate

import (
	"log/slog"
	"time"

	"github.com/salesforceanton/goquorum-parser/internal/protocol/rpc"
	"github.com/salesforceanton/goquorum-parser/internal/transport"
)

type MsgTransport interface {
	Start(u *Cryptogate) error
	Stop()
}

var (
	_ transport.Basis = &Transport{}
	_ MsgTransport    = &Transport{}
)

type Transport struct {
	client  *transport.Client
	logger  *slog.Logger
	timeout time.Duration
	exit    chan struct{}
}

func NewTransport(
	natsClient *transport.Client,
	logger *slog.Logger,
	timeout time.Duration,
) *Transport {
	return &Transport{
		client:  natsClient,
		logger:  logger.With("module", "cryptogate_transport"),
		timeout: timeout,
		exit:    make(chan struct{}),
	}
}

func (t *Transport) Start(s *Cryptogate) error {
	var ec rpc.ErrorsCatcher

	// responders
	// ec.Catch(rpc.CryptogateSendTransactionResponder(t, true, s.SendTransaction))

	return ec.FirstError()
}

func (t *Transport) Stop() {
	close(t.exit)
}

func (t *Transport) Client() *transport.Client {
	return t.client
}

func (t *Transport) Exit() <-chan struct{} {
	return t.exit
}

func (t *Transport) Timeout() time.Duration {
	return t.timeout
}

func (t *Transport) Logger() *slog.Logger {
	return t.logger
}
