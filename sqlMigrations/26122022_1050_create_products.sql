CREATE TABLE IF NOT EXISTS products(
    id UUID NOT NULL,
    name VARCHAR(128) NOT NULL,
    price NUMERIC(10) NOT NULL,
    images JSONB NOT NULL,
    description TEXT NOT NULL,
    features JSONB NOT NULL,
    created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INTEGER,
    updated_at INTEGER,
    CONSTRAINT products_id_pk PRIMARY KEY (id)
)
