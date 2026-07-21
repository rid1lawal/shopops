package server

import (
	"net/http"

	"github.com/rid1lawal/shopops/services/catalog/internal/config"
	"github.com/rid1lawal/shopops/services/catalog/internal/product/handler"
)

func New(
	cfg config.Config,
	productHandler *handler.Handler,
) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	mux.HandleFunc(
		"POST /products",
		productHandler.Create,
	)

	return &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: mux,
	}
}
