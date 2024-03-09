--- TABLES
CREATE TABLE "order" (
	id VARCHAR PRIMARY KEY,
	amount INTEGER NOT NULL,
	status VARCHAR NOT NULL,
	payment_method VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE payment (
	id VARCHAR PRIMARY KEY,
	order_id VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	CONSTRAINT fk_order_payment_id
		FOREIGN KEY (order_id) REFERENCES "order"(id)
);

--- INDEX
CREATE INDEX idx_order_id ON "order" (id);
CREATE INDEX idx_payment_id ON payment (id);
CREATE INDEX idx_payment_order_id ON payment (order_id);

