--- TABLES
CREATE TABLE customer (
	id VARCHAR PRIMARY KEY,
	name VARCHAR NOT NULL,
	email VARCHAR NOT NULL,
	password VARCHAR NOT NULL
);

--- INDEX
CREATE INDEX idx_customer_id ON customer (id);
CREATE INDEX idx_customer_email ON customer (email);
