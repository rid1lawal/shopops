package server

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/rid1lawal/shopops/services/catalog/internal/config"
	"github.com/rid1lawal/shopops/services/catalog/internal/product/handler"
)

func New(
	cfg config.Config,
	productHandler *handler.Handler,
	metricsHandler http.Handler,
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

	mux.HandleFunc(
		"GET /products/{id}",
		productHandler.GetByID,
	)

	mux.Handle(
		"GET /metrics",
		metricsHandler,
	)

	return &http.Server{
		Addr: ":" + cfg.HTTPPort,
		Handler: otelhttp.NewHandler(
			mux,
			"catalog-http",
		),
	}
}
