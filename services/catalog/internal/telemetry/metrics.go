package telemetry

import (
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	Handler http.Handler
	Meter   metric.Meter
}

func NewMetrics() (*Metrics, error) {
	exporter, err := prometheus.New()
	if err != nil {
		return nil, err
	}

	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(exporter),
	)

	otel.SetMeterProvider(provider)

	meter := provider.Meter("catalog")

	return &Metrics{
		Handler: promhttp.Handler(),
		Meter:   meter,
	}, nil
}
