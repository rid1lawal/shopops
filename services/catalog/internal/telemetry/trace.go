package telemetry

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

func LoggerWithTrace(
	logger *slog.Logger,
	ctx context.Context,
) *slog.Logger {
	span := trace.SpanFromContext(ctx)
	spanContext := span.SpanContext()

	if !spanContext.IsValid() {
		return logger
	}

	return logger.With(
		slog.String("trace_id", spanContext.TraceID().String()),
		slog.String("span_id", spanContext.SpanID().String()),
	)
}
