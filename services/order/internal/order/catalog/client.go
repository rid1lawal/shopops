package catalog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

var ErrProductNotFound = errors.New("product not found")

type Product struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	PriceCents int64     `json:"price_cents"`
}

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Transport: otelhttp.NewTransport(http.DefaultTransport),
		},
	}
}

func (client *Client) GetProduct(ctx context.Context, id uuid.UUID) (Product, error) {
	url := fmt.Sprintf("%s/products/%s", client.baseURL, id.String())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Product{}, fmt.Errorf("creating catalog request: %w", err)
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return Product{}, fmt.Errorf("calling catalog service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return Product{}, ErrProductNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return Product{}, fmt.Errorf("catalog service returned unexpected status %d", resp.StatusCode)
	}

	var p Product
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return Product{}, fmt.Errorf("decoding catalog response: %w", err)
	}

	return p, nil
}
