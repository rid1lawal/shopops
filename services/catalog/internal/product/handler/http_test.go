package handler

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/rid1lawal/shopops/services/catalog/internal/product"
)

type fakeRepository struct {
	createFunc  func(ctx context.Context, p product.Product) error
	getByIDFunc func(ctx context.Context, id uuid.UUID) (product.Product, error)
}

func (f *fakeRepository) Create(
	ctx context.Context,
	p product.Product,
) error {
	return f.createFunc(ctx, p)
}

func (f *fakeRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (product.Product, error) {
	return f.getByIDFunc(ctx, id)
}

func TestCreateProduct(t *testing.T) {
	repo := &fakeRepository{
		createFunc: func(
			ctx context.Context,
			p product.Product,
		) error {
			return nil
		},
	}

	h := NewHandler(repo, slog.Default())

	body := strings.NewReader(`{
		"name": "Mechanical Keyboard",
		"description": "A mechanical keyboard",
		"price_cents": 12999
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/products",
		body,
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	h.Create(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusCreated,
			rec.Code,
		)
	}
}

func TestCreateProductValidation(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{
			name: "missing name",
			body: `{
				"name": "",
				"price_cents": 1000
			}`,
		},
		{
			name: "negative price",
			body: `{
				"name": "Keyboard",
				"price_cents": -100
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fakeRepository{
				createFunc: func(
					ctx context.Context,
					p product.Product,
				) error {
					t.Fatal("repository should not be called")
					return nil
				},
			}

			h := NewHandler(repo, slog.Default())

			req := httptest.NewRequest(
				http.MethodPost,
				"/products",
				strings.NewReader(tt.body),
			)

			rec := httptest.NewRecorder()

			h.Create(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf(
					"expected status %d, got %d",
					http.StatusBadRequest,
					rec.Code,
				)
			}
		})
	}
}
func TestGetProduct(t *testing.T) {
	id := uuid.New()

	expected := product.Product{
		ID:         id,
		Name:       "Mechanical Keyboard",
		PriceCents: 12999,
	}

	repo := &fakeRepository{
		getByIDFunc: func(
			ctx context.Context,
			productID uuid.UUID,
		) (product.Product, error) {
			return expected, nil
		},
	}

	h := NewHandler(repo, slog.Default())

	req := httptest.NewRequest(
		http.MethodGet,
		"/products/"+id.String(),
		nil,
	)

	req.SetPathValue("id", id.String())

	rec := httptest.NewRecorder()

	h.GetByID(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}
}
