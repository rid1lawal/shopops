package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/rid1lawal/shopops/services/order/internal/httpresponse"
	"github.com/rid1lawal/shopops/services/order/internal/order"
	"github.com/rid1lawal/shopops/services/order/internal/order/catalog"
	"github.com/rid1lawal/shopops/services/order/internal/order/repository"
)

type OrderRepository interface {
	Create(ctx context.Context, o order.Order) error
	GetByID(ctx context.Context, id uuid.UUID) (order.Order, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status order.Status) error
}

type CatalogClient interface {
	GetProduct(ctx context.Context, id uuid.UUID) (catalog.Product, error)
}

type Handler struct {
	repository    OrderRepository
	catalogClient CatalogClient
	logger        *slog.Logger
}

type createOrderItem struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

type createOrderRequest struct {
	CustomerID uuid.UUID         `json:"customer_id"`
	Items      []createOrderItem `json:"items"`
}

func NewHandler(repo OrderRepository, catalogClient CatalogClient, logger *slog.Logger) *Handler {
	return &Handler{
		repository:    repo,
		catalogClient: catalogClient,
		logger:        logger,
	}
}

func (r createOrderRequest) validate() error {
	if r.CustomerID == uuid.Nil {
		return errors.New("customer_id is required")
	}

	if len(r.Items) == 0 {
		return errors.New("items must not be empty")
	}

	for _, item := range r.Items {
		if item.ProductID == uuid.Nil {
			return errors.New("product_id is required for each item")
		}
		if item.Quantity <= 0 {
			return errors.New("quantity must be greater than zero")
		}
	}

	return nil
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req createOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Warn("invalid create order request", slog.String("error", err.Error()))
		httpresponse.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.validate(); err != nil {
		h.logger.Warn("order validation failed", slog.String("error", err.Error()))
		httpresponse.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	var items []order.Item
	var totalCents int64

	for _, reqItem := range req.Items {
		product, err := h.catalogClient.GetProduct(r.Context(), reqItem.ProductID)
		if err != nil {
			if errors.Is(err, catalog.ErrProductNotFound) {
				httpresponse.Error(
					w,
					http.StatusUnprocessableEntity,
					"product "+reqItem.ProductID.String()+" not found",
				)
				return
			}

			h.logger.Error(
				"failed to fetch product from catalog",
				slog.String("product_id", reqItem.ProductID.String()),
				slog.String("error", err.Error()),
			)

			httpresponse.Error(w, http.StatusInternalServerError, "failed to validate products")
			return
		}

		totalCents += product.PriceCents * int64(reqItem.Quantity)

		items = append(items, order.Item{
			ID:         uuid.New(),
			ProductID:  reqItem.ProductID,
			Quantity:   reqItem.Quantity,
			PriceCents: product.PriceCents,
		})
	}

	now := time.Now()
	o := order.Order{
		ID:         uuid.New(),
		CustomerID: req.CustomerID,
		Items:      items,
		Status:     order.StatusPending,
		TotalCents: totalCents,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	if err := h.repository.Create(r.Context(), o); err != nil {
		h.logger.Error("failed to create order", slog.String("error", err.Error()))
		httpresponse.Error(w, http.StatusInternalServerError, "failed to create order")
		return
	}

	h.logger.Info("order created", slog.String("order_id", o.ID.String()))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(o)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "invalid order id")
		return
	}

	o, err := h.repository.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			httpresponse.Error(w, http.StatusNotFound, "order not found")
			return
		}

		httpresponse.Error(w, http.StatusInternalServerError, "failed to retrieve order")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(o)
}

func (h *Handler) Cancel(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "invalid order id")
		return
	}

	o, err := h.repository.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			httpresponse.Error(w, http.StatusNotFound, "order not found")
			return
		}

		httpresponse.Error(w, http.StatusInternalServerError, "failed to retrieve order")
		return
	}

	if o.Status != order.StatusPending && o.Status != order.StatusConfirmed {
		httpresponse.Error(
			w,
			http.StatusUnprocessableEntity,
			"only pending or confirmed orders can be cancelled",
		)
		return
	}

	if err := h.repository.UpdateStatus(r.Context(), id, order.StatusCancelled); err != nil {
		h.logger.Error(
			"failed to cancel order",
			slog.String("order_id", id.String()),
			slog.String("error", err.Error()),
		)

		httpresponse.Error(w, http.StatusInternalServerError, "failed to cancel order")
		return
	}

	h.logger.Info("order cancelled", slog.String("order_id", id.String()))

	w.WriteHeader(http.StatusNoContent)
}
