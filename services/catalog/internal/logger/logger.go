package logger

import (
	"log/slog"
	"os"
)

func New(environment string) *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return slog.New(handler).With(
		slog.String("service", "catalog"),
		slog.String("environment", environment),
	)
}
