package tracing

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type ContextKey string

const (
	TraceID ContextKey = "traceID"
)

func Start(ctx context.Context, logger *slog.Logger) (context.Context, *slog.Logger) {
	traceID := GetTraceID(ctx)
	if traceID == "" {
		traceID = uuid.NewString()
	}

	return SetTraceID(ctx, traceID), logger.With("traceID", traceID)
}

func SetTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceID, traceID)
}

func GetTraceID(ctx context.Context) string {
	raw := ctx.Value(TraceID)
	if raw == nil {
		return ""
	}

	traceID, ok := raw.(string)
	if !ok {
		return ""
	}

	return traceID
}
