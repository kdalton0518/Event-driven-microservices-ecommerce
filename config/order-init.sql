--- TABLES
CREATE TABLE "order" (
	id VARCHAR PRIMARY KEY,
	customer_id VARCHAR NOT NULL,
    product_id_list INTEGER[] NOT NULL,
	total_price INTEGER NOT NULL,
	status VARCHAR NOT NULL,
	payment_method VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

--- INDEX
CREATE INDEX idx_order_id ON "order" (id);
CREATE INDEX idx_order_customer_id ON "order" (customer_id);

