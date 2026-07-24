package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/google/uuid"

	"github.com/rid1lawal/shopops/services/catalog/internal/httpresponse"
	"github.com/rid1lawal/shopops/services/catalog/internal/product"
	"github.com/rid1lawal/shopops/services/catalog/internal/product/repository"
	"github.com/rid1lawal/shopops/services/catalog/internal/telemetry"
)

type Handler struct {
	repository ProductRepository
	logger     *slog.Logger
}

type ProductRepository interface {
	Create(ctx context.Context, p product.Product) error
	GetByID(ctx context.Context, id uuid.UUID) (product.Product, error)
}

func NewHandler(
	repository ProductRepository,
	logger *slog.Logger,
) *Handler {
	return &Handler{
		repository: repository,
		logger:     logger,
	}
}

type createProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	PriceCents  int64  `json:"price_cents"`
}

func (r createProductRequest) validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}

	if r.PriceCents < 0 {
		return errors.New("price_cents must be greater than or equal to zero")
	}

	return nil
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	logger := telemetry.LoggerWithTrace(
		h.logger,
		r.Context(),
	)

	var req createProductRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Warn(
			"invalid create product request",
			slog.String("error", err.Error()),
		)

		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		logger.Warn(
			"product validation failed",
			slog.String("error", err.Error()),
		)

		httpresponse.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	p := product.Product{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		PriceCents:  req.PriceCents,
	}

	if err := h.repository.Create(r.Context(), p); err != nil {
		logger.Error(
			"failed to create product",
			slog.String("error", err.Error()),
		)

		http.Error(w, "failed to create product", http.StatusInternalServerError)
		return
	}

	logger.Info(
		"product created",
		slog.String("product_id", p.ID.String()),
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(p)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httpresponse.Error(
			w,
			http.StatusBadRequest,
			"invalid product id",
		)
		return
	}

	p, err := h.repository.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			httpresponse.Error(
				w,
				http.StatusNotFound,
				"product not found",
			)
			return
		}

		httpresponse.Error(
			w,
			http.StatusInternalServerError,
			"failed to retrieve product",
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(p); err != nil {
		return
	}
}
