#!/bin/bash
set -e

# Create order_db — connect to catalog_db which already exists
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "catalog_db" <<-EOSQL
    CREATE DATABASE order_db;
EOSQL

# Create catalog tables
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "catalog_db" <<-EOSQL
    CREATE TABLE IF NOT EXISTS products (
        id          UUID        PRIMARY KEY,
        name        TEXT        NOT NULL,
        description TEXT        NOT NULL DEFAULT '',
        price_cents BIGINT      NOT NULL,
        created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );
EOSQL

# Create order tables
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "order_db" <<-EOSQL
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
EOSQL