package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rid1lawal/shopops/services/order/internal/order"
)

var ErrNotFound = errors.New("order not found")

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r *PostgresRepository) Create(ctx context.Context, ord order.Order) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}
	defer tx.Rollback(ctx) // no-op if Commit succeeds

	_, err = tx.Exec(ctx, `
		INSERT INTO orders (id, customer_id, status, total_cents, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, ord.ID, ord.CustomerID, string(ord.Status), ord.TotalCents, ord.CreatedAt, ord.UpdatedAt)
	if err != nil {
		return fmt.Errorf("inserting order: %w", err)
	}

	for _, item := range ord.Items {
		_, err = tx.Exec(ctx, `
			INSERT INTO order_items (id, order_id, product_id, quantity, price_cents)
			VALUES ($1, $2, $3, $4, $5)
		`, item.ID, ord.ID, item.ProductID, item.Quantity, item.PriceCents)
		if err != nil {
			return fmt.Errorf("inserting order item %s: %w", item.ProductID, err)
		}
	}

	return tx.Commit(ctx)
}

func (r *PostgresRepository) GetByID(ctx context.Context, id uuid.UUID) (order.Order, error) {
	row := r.pool.QueryRow(ctx, `
		SELECT id, customer_id, status, total_cents, created_at, updated_at
		FROM orders
		WHERE id = $1
	`, id)

	var ord order.Order
	var status string

	if err := row.Scan(
		&ord.ID,
		&ord.CustomerID,
		&status,
		&ord.TotalCents,
		&ord.CreatedAt,
		&ord.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return order.Order{}, ErrNotFound
		}
		return order.Order{}, fmt.Errorf("scanning order: %w", err)
	}

	ord.Status = order.Status(status)

	rows, err := r.pool.Query(ctx, `
		SELECT id, product_id, quantity, price_cents
		FROM order_items
		WHERE order_id = $1
	`, id)
	if err != nil {
		return order.Order{}, fmt.Errorf("querying order items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item order.Item
		if err := rows.Scan(&item.ID, &item.ProductID, &item.Quantity, &item.PriceCents); err != nil {
			return order.Order{}, fmt.Errorf("scanning order item: %w", err)
		}
		ord.Items = append(ord.Items, item)
	}

	if err := rows.Err(); err != nil {
		return order.Order{}, fmt.Errorf("iterating order items: %w", err)
	}

	return ord, nil
}

func (r *PostgresRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status order.Status) error {
	result, err := r.pool.Exec(ctx, `
		UPDATE orders
		SET status = $1, updated_at = NOW()
		WHERE id = $2
	`, string(status), id)
	if err != nil {
		return fmt.Errorf("updating order status: %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
