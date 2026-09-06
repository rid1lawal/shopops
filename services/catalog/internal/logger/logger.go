package logger

import (
	"log/slog"
	"os"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

func New(environment string, provider *sdklog.LoggerProvider) *slog.Logger {
	stdoutHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	otelHandler := otelslog.NewHandler("catalog", otelslog.WithLoggerProvider(provider))

	handler := NewMultiHandler(stdoutHandler, otelHandler)

	return slog.New(handler).With(
		slog.String("service", "catalog"),
		slog.String("environment", environment),
	)
}
