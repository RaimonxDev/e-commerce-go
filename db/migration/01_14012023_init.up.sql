CREATE DATABASE IF NOT EXISTS "ecommerce" WITH OWNER = "ramon" ENCODING = 'UTF8' CONNECTION LIMIT = -1;


CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_admin BOOLEAN NOT NULL,
    details jsonb NOT NULL ,
    created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
    updated_at integer,
    CONSTRAINT user_id_pk PRIMARY KEY (id),
    CONSTRAINT user_email_unique UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS purchase_orders (
	id UUID NOT NULL,
	userID UUID NOT NULL,
	created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::INT,
	updated_at INTEGER,
	CONSTRAINT purchase_orders_id_pk PRIMARY KEY (id),
	CONSTRAINT purchase_orders_userID_fk FOREIGN KEY (userID) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE IF NOT EXISTS products(
    id UUID NOT NULL,
    name VARCHAR(128) NOT NULL,
    price NUMERIC(10) NOT NULL,
	details JSONB NOT NULL,
    images JSONB NOT NULL,
    description TEXT NOT NULL,
    features JSONB NOT NULL,
    created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
    updated_at INTEGER,
    CONSTRAINT products_id_pk PRIMARY KEY (id)
);

