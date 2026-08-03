package telemetry

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel/metric"
)

type DBMetrics struct {
	TotalConnections    metric.Int64ObservableGauge
	IdleConnections     metric.Int64ObservableGauge
	AcquiredConnections metric.Int64ObservableGauge
	MaxConnections      metric.Int64ObservableGauge
}

func NewDBMetrics(
	meter metric.Meter,
	pool *pgxpool.Pool,
) (*DBMetrics, error) {
	totalConnections, err := meter.Int64ObservableGauge(
		"catalog_db_connections_total",
		metric.WithDescription(
			"Current number of connections in the database pool",
		),
	)
	if err != nil {
		return nil, fmt.Errorf("create total connections metric: %w", err)
	}

	idleConnections, err := meter.Int64ObservableGauge(
		"catalog_db_connections_idle",
		metric.WithDescription(
			"Current number of idle connections in the database pool",
		),
	)
	if err != nil {
		return nil, fmt.Errorf("create idle connections metric: %w", err)
	}

	acquiredConnections, err := meter.Int64ObservableGauge(
		"catalog_db_connections_acquired",
		metric.WithDescription(
			"Current number of acquired database connections",
		),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"create acquired connections metric: %w",
			err,
		)
	}

	maxConnections, err := meter.Int64ObservableGauge(
		"catalog_db_connections_max",
		metric.WithDescription(
			"Maximum number of connections allowed in the database pool",
		),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"create max connections metric: %w",
			err,
		)
	}

	_, err = meter.RegisterCallback(
		func(
			ctx context.Context,
			observer metric.Observer,
		) error {
			stats := pool.Stat()

			observer.ObserveInt64(
				totalConnections,
				int64(stats.TotalConns()),
			)

			observer.ObserveInt64(
				idleConnections,
				int64(stats.IdleConns()),
			)

			observer.ObserveInt64(
				acquiredConnections,
				int64(stats.AcquiredConns()),
			)

			observer.ObserveInt64(
				maxConnections,
				int64(stats.MaxConns()),
			)

			return nil
		},
		totalConnections,
		idleConnections,
		acquiredConnections,
		maxConnections,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"register database metrics callback: %w",
			err,
		)
	}

	return &DBMetrics{
		TotalConnections:    totalConnections,
		IdleConnections:     idleConnections,
		AcquiredConnections: acquiredConnections,
		MaxConnections:      maxConnections,
	}, nil
}
