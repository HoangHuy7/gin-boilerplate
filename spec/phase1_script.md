# 🗄️ 1. `products`

```sql
CREATE TABLE products (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    category TEXT,            -- pesticide / fertilizer / feed
    unit TEXT,                -- chai, bao, kg
    price NUMERIC(12,2) NOT NULL,
    cost_price NUMERIC(12,2),
    stock_quantity INTEGER DEFAULT 0,
    barcode TEXT UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);
```

---

# 📦 2. `inventory_logs`

```sql
CREATE TABLE inventory_logs (
    id UUID PRIMARY KEY,
    product_id UUID,
    type TEXT,               -- import / sale / adjust
    quantity INTEGER,
    note TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
```

👉 Không cần FK, bạn join bằng view sau

---

# 🧾 3. `orders`

```sql
CREATE TABLE orders (
    id UUID PRIMARY KEY,
    code TEXT,               -- mã bill (in ra)
    customer_id UUID,        -- nullable
    total_amount NUMERIC(12,2),
    paid_amount NUMERIC(12,2),
    debt_amount NUMERIC(12,2),
    status TEXT,             -- paid / debt / delivery
    note TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

# 🛒 4. `order_items`

```sql
CREATE TABLE order_items (
    id UUID PRIMARY KEY,
    order_id UUID,
    product_id UUID,
    product_name TEXT,       -- snapshot (tránh bị đổi tên)
    quantity INTEGER,
    price NUMERIC(12,2),
    total NUMERIC(12,2)
);
```

👉 **Snapshot product_name là cực quan trọng** (sau đổi tên vẫn giữ lịch sử)

---

# 👤 5. `customers`

```sql
CREATE TABLE customers (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    phone TEXT,
    address TEXT,
    total_debt NUMERIC(12,2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

# 💰 6. `debt_transactions` (Phase 2)

```sql
CREATE TABLE debt_transactions (
    id UUID PRIMARY KEY,
    customer_id UUID,
    order_id UUID,
    amount NUMERIC(12,2),
    type TEXT,              -- borrow / pay
    note TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
```

👉 Giúp track:

* Ghi nợ (borrow)
* Trả tiền (pay)

---

# 🚚 7. `deliveries`

```sql
CREATE TABLE deliveries (
    id UUID PRIMARY KEY,
    order_id UUID,
    delivery_date DATE,
    status TEXT,            -- pending / done
    note TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

# 🧾 8. `order_counters` (optional nhưng nên có)

```sql
CREATE TABLE order_counters (
    id TEXT PRIMARY KEY,    -- ví dụ: "20260317"
    current_number INTEGER
);
```

👉 Dùng generate code kiểu:

```
HD-20260317-001
```

---

# 🔥 9. Index gợi ý (rất nên có)

```sql
CREATE INDEX idx_products_barcode ON products(barcode);
CREATE INDEX idx_orders_created_at ON orders(created_at);
CREATE INDEX idx_orders_customer_id ON orders(customer_id);
CREATE INDEX idx_order_items_order_id ON order_items(order_id);
CREATE INDEX idx_inventory_logs_product_id ON inventory_logs(product_id);
```

---

# ⚡ Flow mapping nhanh (để bạn code luôn)

## Khi bán hàng:

* insert `orders`
* insert `order_items`
* update `products.stock_quantity -=`
* insert `inventory_logs (type = sale)`
* nếu nợ:

    * insert `debt_transactions (borrow)`
    * update `customers.total_debt`

---

## Khi khách trả tiền:

* insert `debt_transactions (pay)`
* update `customers.total_debt -=`

---

## Khi nhập hàng:

* update `products.stock_quantity +=`
* insert `inventory_logs (type = import)`

---

# 💡 Tips thực chiến (rất đáng giá)

* Không tính stock từ logs → **luôn lưu stock_quantity**
* Order không update → chỉ insert (tránh sai dữ liệu)
* Không xóa order → chỉ “void” nếu cần
* Barcode nên unique (tránh scan nhầm)
