package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rid1lawal/shopops/services/catalog/internal/product"
)

var ErrNotFound = errors.New("product not found")

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(
	ctx context.Context,
	p product.Product,
) error {
	_, err := r.db.Exec(
		ctx,
		`
		INSERT INTO products (
			id,
			name,
			description,
			price_cents
		)
		VALUES ($1, $2, $3, $4)
		`,
		p.ID,
		p.Name,
		p.Description,
		p.PriceCents,
	)

	if err != nil {
		return fmt.Errorf("create product: %w", err)
	}

	return nil
}

func (r *PostgresRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (product.Product, error) {
	var p product.Product

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			name,
			description,
			price_cents,
			created_at
		FROM products
		WHERE id = $1
		`,
		id,
	).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.PriceCents,
		&p.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return product.Product{}, ErrNotFound
	}

	if err != nil {
		return product.Product{}, fmt.Errorf("get product: %w", err)
	}

	return p, nil
}
