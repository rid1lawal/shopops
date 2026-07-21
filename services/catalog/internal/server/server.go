package server

import (
	"net/http"

	"github.com/rid1lawal/shopops/services/catalog/internal/config"
)

func New(cfg config.Config) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	return &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: mux,
	}
}
