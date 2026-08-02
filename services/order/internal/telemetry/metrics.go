package telemetry

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

type Metrics struct {
	Meter   metric.Meter
	Handler http.Handler
}

func NewMetrics() (*Metrics, error) {
	exporter, err := prometheus.New()
	if err != nil {
		return nil, fmt.Errorf("creating prometheus exporter: %w", err)
	}

	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(exporter),
	)

	meter := provider.Meter("order")

	return &Metrics{
		Meter:   meter,
		Handler: promhttp.Handler(),
	}, nil
}
