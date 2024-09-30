package transport

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/salesforceanton/goquorum-parser/internal/tracing"
)

const traceIDHeader = "TraceID"

func propagateTraceID(ctx context.Context, msg *nats.Msg) {
	traceID := tracing.GetTraceID(ctx)
	if traceID == "" {
		return
	}

	if msg.Header == nil {
		msg.Header = make(nats.Header)
	}

	msg.Header.Set(traceIDHeader, traceID)
}

func restoreTraceID(ctx context.Context, msg *nats.Msg) context.Context {
	if msg.Header == nil {
		return ctx
	}

	traceID := msg.Header.Get(traceIDHeader)
	if traceID == "" {
		return ctx
	}

	return tracing.SetTraceID(ctx, traceID)
}
