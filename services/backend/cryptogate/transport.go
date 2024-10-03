package cryptogate

import (
	"context"
	"log/slog"
	"time"

	"github.com/salesforceanton/goquorum-parser/internal/protocol/messages/cryptogatemessages"
	"github.com/salesforceanton/goquorum-parser/internal/protocol/rpc"
	"github.com/salesforceanton/goquorum-parser/internal/transport"
)

type MsgTransport interface {
	Start(u *Cryptogate) error
	Stop()

	GetBalanceNative(context.Context, cryptogatemessages.BalanceNativeRequest) (cryptogatemessages.BalanceNativeResponse, error)
	GetBalanceToken(context.Context, cryptogatemessages.BalanceTokenRequest) (cryptogatemessages.BalanceTokenResponse, error)
	SendTransaction(context.Context, cryptogatemessages.SendTransactionRequest) (cryptogatemessages.SendTransactionResponse, error)
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
	ec.Catch(rpc.CryptogateSendTransactionResponder(t, true, s.SendTransaction))
	ec.Catch(rpc.CryptogateGetBalanceNativeResponder(t, true, s.GetBalanceNative))
	ec.Catch(rpc.CryptogateGetBalanceTokenResponder(t, true, s.GetBalanceToken))

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

func (t *Transport) GetBalanceNative(
	ctx context.Context,
	req cryptogatemessages.BalanceNativeRequest,
) (cryptogatemessages.BalanceNativeResponse, error) {
	return rpc.CryptogateGetBalanceNativeRequester(t, ctx, req)
}

func (t *Transport) GetBalanceToken(
	ctx context.Context,
	req cryptogatemessages.BalanceTokenRequest,
) (cryptogatemessages.BalanceTokenResponse, error) {
	return rpc.CryptogateGetBalanceTokenRequester(t, ctx, req)
}

func (t *Transport) SendTransaction(
	ctx context.Context,
	req cryptogatemessages.SendTransactionRequest,
) (cryptogatemessages.SendTransactionResponse, error) {
	return rpc.CryptogateSendTransactionRequester(t, ctx, req)
}
