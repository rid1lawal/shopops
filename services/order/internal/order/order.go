package order

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusConfirmed Status = "confirmed"
	StatusCancelled Status = "cancelled"
)

// Item captures the product price at the moment the order was placed.
// This is intentional: product prices in the catalog can change after
// an order is created, but the order must always reflect what the
// customer was charged.
type Item struct {
	ID         uuid.UUID
	ProductID  uuid.UUID
	Quantity   int
	PriceCents int64
}

type Order struct {
	ID         uuid.UUID
	CustomerID uuid.UUID
	Items      []Item
	Status     Status
	TotalCents int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
