package telemetry

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel/metric"
)

type DBMetrics struct{}

func NewDBMetrics(meter metric.Meter, pool *pgxpool.Pool) (*DBMetrics, error) {
	_, err := meter.Int64ObservableGauge(
		"db.pool.acquired_connections",
		metric.WithDescription("Number of acquired connections in the pool"),
		metric.WithInt64Callback(func(_ context.Context, o metric.Int64Observer) error {
			o.Observe(int64(pool.Stat().AcquiredConns()))
			return nil
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("creating acquired connections gauge: %w", err)
	}

	_, err = meter.Int64ObservableGauge(
		"db.pool.idle_connections",
		metric.WithDescription("Number of idle connections in the pool"),
		metric.WithInt64Callback(func(_ context.Context, o metric.Int64Observer) error {
			o.Observe(int64(pool.Stat().IdleConns()))
			return nil
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("creating idle connections gauge: %w", err)
	}

	return &DBMetrics{}, nil
}
