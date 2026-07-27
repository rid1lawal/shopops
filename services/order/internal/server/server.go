package server

import (
	"net/http"

	"github.com/rid1lawal/shopops/services/order/internal/config"
	"github.com/rid1lawal/shopops/services/order/internal/order/handler"
)

func New(
	cfg config.Config,
	orderHandler *handler.Handler,
) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("POST /orders", orderHandler.Create)
	mux.HandleFunc("GET /orders/{id}", orderHandler.GetByID)
	mux.HandleFunc("PATCH /orders/{id}/cancel", orderHandler.Cancel)

	return &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: mux,
	}
}
