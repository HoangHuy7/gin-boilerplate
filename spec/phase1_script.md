-- =========================
-- DROP OLD TABLES
-- =========================
-- =========================
-- PRODUCTS
-- =========================
CREATE TABLE  mkrtb_products (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
name TEXT NOT NULL,
category TEXT,
unit TEXT,
price NUMERIC(12,2) NOT NULL,
cost_price NUMERIC(12,2),
stock_quantity INTEGER DEFAULT 0,
barcode TEXT UNIQUE,
created_at TIMESTAMP DEFAULT NOW(),
updated_at TIMESTAMP
);

-- =========================
-- INVENTORY LOGS
-- =========================
CREATE TABLE mkrtb_inventory_logs (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
product_id UUID,
type TEXT, -- import / sale / adjust
quantity INTEGER,
note TEXT,
created_at TIMESTAMP DEFAULT NOW()
);

-- =========================
-- CUSTOMERS
-- =========================
CREATE TABLE mkrtb_customers (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
name TEXT NOT NULL,
phone TEXT,
address TEXT,
total_debt NUMERIC(12,2) DEFAULT 0,
created_at TIMESTAMP DEFAULT NOW()
);

-- =========================
-- ORDERS
-- =========================
CREATE TABLE mkrtb_orders (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
code TEXT,
customer_id UUID,
total_amount NUMERIC(12,2),
paid_amount NUMERIC(12,2),
debt_amount NUMERIC(12,2),
status TEXT, -- paid / debt / delivery
note TEXT,
created_at TIMESTAMP DEFAULT NOW()
);

-- =========================
-- ORDER ITEMS
-- =========================
CREATE TABLE mkrtb_order_items (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
order_id UUID,
product_id UUID,
product_name TEXT,
quantity INTEGER,
price NUMERIC(12,2),
total NUMERIC(12,2)
);

-- =========================
-- DEBT TRANSACTIONS
-- =========================
CREATE TABLE mkrtb_debt_transactions (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
customer_id UUID,
order_id UUID,
amount NUMERIC(12,2),
type TEXT, -- borrow / pay
note TEXT,
created_at TIMESTAMP DEFAULT NOW()
);

-- =========================
-- DELIVERIES
-- =========================
CREATE TABLE mkrtb_deliveries (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
order_id UUID,
delivery_date DATE,
status TEXT, -- pending / done
note TEXT,
created_at TIMESTAMP DEFAULT NOW()
);

-- =========================
-- ORDER COUNTERS
-- =========================
CREATE TABLE mkrtb_order_counters (
id TEXT PRIMARY KEY, -- ví dụ: 20260317
current_number INTEGER
);

-- =========================
-- INDEXES
-- =========================
CREATE INDEX idx_products_barcode ON  mkrtb_products(barcode);
CREATE INDEX idx_orders_created_at ON  mkrtb_orders(created_at);
CREATE INDEX idx_orders_customer_id ON  mkrtb_orders(customer_id);
CREATE INDEX idx_order_items_order_id ON  mkrtb_order_items(order_id);
CREATE INDEX idx_inventory_logs_product_id ON  mkrtb_inventory_logs(product_id);

INSERT INTO mkrtb_products (
name, category, unit, price, cost_price, stock_quantity, barcode
) VALUES
-- Thuốc trừ sâu
('Thuốc trừ sâu Regent 800WG', 'pesticide', 'gói', 35000, 28000, 120, '893850597001'),
('Thuốc trừ sâu Confidor 100SL', 'pesticide', 'chai', 45000, 37000, 80, '893850597002'),
('Thuốc trừ sâu Radiant 60SC', 'pesticide', 'chai', 95000, 80000, 60, '893850597003'),

-- Phân bón
('Phân NPK 16-16-8', 'fertilizer', 'bao', 320000, 290000, 50, '893850597004'),
('Phân Ure Phú Mỹ', 'fertilizer', 'bao', 280000, 250000, 70, '893850597005'),
('Phân Kali đỏ', 'fertilizer', 'bao', 400000, 360000, 40, '893850597006'),

-- Thuốc bệnh cây
('Thuốc trị nấm Ridomil Gold', 'pesticide', 'gói', 60000, 50000, 90, '893850597007'),
('Thuốc trừ bệnh Antracol 70WP', 'pesticide', 'gói', 55000, 47000, 100, '893850597008'),

-- Thức ăn chăn nuôi
('Cám heo CP 551', 'feed', 'bao', 310000, 295000, 30, '893850597009'),
('Cám gà CP 201', 'feed', 'bao', 290000, 270000, 35, '893850597010');