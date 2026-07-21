package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/rid1lawal/shopops/services/catalog/internal/product"
)

type Handler struct {
	repository ProductRepository
}

type ProductRepository interface {
	Create(ctx context.Context, p product.Product) error
	GetByID(ctx context.Context, id uuid.UUID) (product.Product, error)
}

func New(repository ProductRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

type createProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	PriceCents  int64  `json:"price_cents"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req createProductRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	p := product.Product{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		PriceCents:  req.PriceCents,
	}

	if err := h.repository.Create(r.Context(), p); err != nil {
		http.Error(w, "failed to create product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(p)
}
