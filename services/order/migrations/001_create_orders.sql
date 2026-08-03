CREATE TABLE IF NOT EXISTS orders (
    id          UUID        PRIMARY KEY,
    customer_id UUID        NOT NULL,
    status      TEXT        NOT NULL DEFAULT 'pending',
    total_cents BIGINT      NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL,
    updated_at  TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS order_items (
    id          UUID    PRIMARY KEY,
    order_id    UUID    NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id  UUID    NOT NULL,
    quantity    INT     NOT NULL CHECK (quantity > 0),
    price_cents BIGINT  NOT NULL CHECK (price_cents >= 0)
);

CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON order_items(order_id);
CREATE INDEX IF NOT EXISTS idx_orders_customer_id   ON orders(customer_id);