package transport

import (
	"log/slog"
	"testing"

	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/require"

	"github.com/salesforceanton/goquorum-parser/internal/custerror"
)

func TestErrorPass(t *testing.T) {
	cerr := custerror.NewGeneralError(101, "AAA", "bbb")
	msg := nats.Msg{}
	client := Client{
		logger: slog.Default(),
		nc:     &nats.Conn{},
	}

	client.RespondWithError(&msg, cerr)
	err := ErrorFromMsg(&msg)
	custErr, ok := err.(custerror.CustomError)
	require.True(t, ok)
	require.Equal(t, cerr.HTTP(), custErr.HTTP())
	require.Equal(t, cerr.CustomCode(), custErr.CustomCode())
	require.Equal(t, cerr.CodeExpr(), custErr.CodeExpr())
	require.Equal(t, cerr.Error(), custErr.Error())
}
