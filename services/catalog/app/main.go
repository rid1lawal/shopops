package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rid1lawal/shopops/services/catalog/internal/config"
	"github.com/rid1lawal/shopops/services/catalog/internal/database"
	"github.com/rid1lawal/shopops/services/catalog/internal/logger"
	"github.com/rid1lawal/shopops/services/catalog/internal/product/handler"
	"github.com/rid1lawal/shopops/services/catalog/internal/product/repository"
	"github.com/rid1lawal/shopops/services/catalog/internal/server"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Environment)

	ctx := context.Background()

	dbPool, err := database.NewPostgresPool(
		ctx,
		cfg.DatabaseURL,
	)
	if err != nil {
		log.Error("database connection failed",
			slog.String("error", err.Error()),
		)

		os.Exit(1)
	}

	defer dbPool.Close()

	log.Info("database connection established")

	productRepository := repository.NewPostgresRepository(dbPool)

	productHandler := handler.New(productRepository)

	srv := server.New(
		cfg,
		productHandler,
	)

	serverErrors := make(chan error, 1)

	go func() {
		log.Info("starting HTTP server",
			slog.String("address", srv.Addr),
		)

		serverErrors <- srv.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)

	signal.Notify(
		shutdown,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	select {
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			log.Error("server error",
				slog.String("error", err.Error()),
			)

			os.Exit(1)
		}

	case sig := <-shutdown:
		log.Info("shutdown signal received",
			slog.String("signal", sig.String()),
		)

		ctx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Error("graceful shutdown failed",
				slog.String("error", err.Error()),
			)

			if err := srv.Close(); err != nil {
				log.Error("forced shutdown failed",
					slog.String("error", err.Error()),
				)
			}
		}

		log.Info("server shutdown complete")
	}
}
