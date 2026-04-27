CREATE EXTENSION IF NOT EXISTS pgcrypto;


CREATE TABLE IF NOT EXISTS gas_price (
    id BIGSERIAL PRIMARY KEY,
    price NUMERIC(10, 2) NOT NULL,
    "timestamp" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS mkrtb_customers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    phone TEXT,
    address TEXT,
    total_debt NUMERIC(12, 2) NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS mkrtb_products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    category TEXT,
    unit TEXT,
    price NUMERIC(12, 2) NOT NULL,
    cost_price NUMERIC(12, 2),
    stock_quantity INTEGER NOT NULL DEFAULT 0,
    barcode TEXT UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    img_url VARCHAR(1024)
);

CREATE TABLE IF NOT EXISTS mkrtb_orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(64),
    customer_id UUID,
    total_amount NUMERIC(12, 2) NOT NULL DEFAULT 0,
    paid_amount NUMERIC(12, 2) NOT NULL DEFAULT 0,
    debt_amount NUMERIC(12, 2) NOT NULL DEFAULT 0,
    status VARCHAR(20),
    note VARCHAR(512),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS mkrtb_order_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL,
    product_id UUID NOT NULL,
    product_name VARCHAR(255),
    quantity INTEGER NOT NULL,
    price NUMERIC(12, 2) NOT NULL DEFAULT 0,
    total NUMERIC(12, 2) NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS mkrtb_deliveries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL,
    delivery_date DATE,
    status TEXT,
    note TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS mkrtb_debt_transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id UUID,
    order_id UUID,
    amount NUMERIC(12, 2) NOT NULL DEFAULT 0,
    type TEXT,
    note TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS mkrtb_inventory_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL,
    type TEXT,
    quantity INTEGER NOT NULL DEFAULT 0,
    note TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS mkrtb_logs (
    id BIGSERIAL PRIMARY KEY,
    action VARCHAR(100),
    status VARCHAR(100),
    feature VARCHAR(255),
    old_data JSONB,
    new_data JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(100),
    ip_address VARCHAR(45),
    user_agent TEXT,
    request_id VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS mkrtb_order_counters (
    id TEXT PRIMARY KEY,
    current_number INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_mkrtb_orders_customer_id ON mkrtb_orders(customer_id);
CREATE INDEX IF NOT EXISTS idx_mkrtb_orders_created_at ON mkrtb_orders(created_at);
CREATE INDEX IF NOT EXISTS idx_mkrtb_order_items_order_id ON mkrtb_order_items(order_id);
CREATE INDEX IF NOT EXISTS idx_mkrtb_order_items_product_id ON mkrtb_order_items(product_id);
CREATE INDEX IF NOT EXISTS idx_mkrtb_deliveries_order_id ON mkrtb_deliveries(order_id);
CREATE INDEX IF NOT EXISTS idx_mkrtb_debt_tx_customer_id ON mkrtb_debt_transactions(customer_id);
CREATE INDEX IF NOT EXISTS idx_mkrtb_debt_tx_order_id ON mkrtb_debt_transactions(order_id);
CREATE INDEX IF NOT EXISTS idx_mkrtb_inventory_logs_product_id ON mkrtb_inventory_logs(product_id);
CREATE INDEX IF NOT EXISTS idx_mkrtb_logs_created_at ON mkrtb_logs(created_at);
CREATE INDEX IF NOT EXISTS idx_mkrtb_logs_request_id ON mkrtb_logs(request_id);

CREATE OR REPLACE VIEW vw_sales_summary AS
SELECT
    DATE_TRUNC('day', o.created_at) AS created_at,
    DATE(o.created_at) AS sale_date,
    DATE_TRUNC('month', o.created_at)::DATE AS sale_month,
    EXTRACT(YEAR FROM o.created_at)::INT AS sale_year,
    COALESCE(SUM(oi.quantity), 0)::INT AS quantity,
    COALESCE(SUM(oi.total), 0)::NUMERIC(12, 2) AS total
FROM mkrtb_orders o
LEFT JOIN mkrtb_order_items oi ON oi.order_id = o.id
GROUP BY DATE(o.created_at), DATE_TRUNC('day', o.created_at), DATE_TRUNC('month', o.created_at), EXTRACT(YEAR FROM o.created_at);
