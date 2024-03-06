--- TABLES
CREATE TABLE "order" (
	id VARCHAR PRIMARY KEY,
	customer_id VARCHAR NOT NULL,
	total_price INTEGER NOT NULL,
	status VARCHAR NOT NULL,
	payment_method VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE order_product (
	id SERIAL PRIMARY KEY,
	order_id VARCHAR NOT NULL,
    product_id INTEGER NOT NULL,
	price INTEGER NOT NULL,
	quantity INTEGER NOT NULL,
	CONSTRAINT fk_order_order_product_id
		FOREIGN KEY (order_id) REFERENCES "order"(id)
);

--- INDEX
CREATE INDEX idx_order_id ON "order" (id);
CREATE INDEX idx_order_customer_id ON "order" (customer_id);
CREATE INDEX idx_order_product_order_id ON order_product (order_id);

