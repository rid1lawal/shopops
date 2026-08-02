package server

import (
	"net/http"

	"github.com/rid1lawal/shopops/services/order/internal/config"
	"github.com/rid1lawal/shopops/services/order/internal/order/handler"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func New(
	cfg config.Config,
	orderHandler *handler.Handler,
	metricsHandler http.Handler,
) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("POST /orders", orderHandler.Create)
	mux.HandleFunc("GET /orders/{id}", orderHandler.GetByID)
	mux.HandleFunc("PATCH /orders/{id}/cancel", orderHandler.Cancel)

	mux.Handle("GET /metrics", metricsHandler)

	return &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: otelhttp.NewHandler(mux, "order-http"),
	}
}
