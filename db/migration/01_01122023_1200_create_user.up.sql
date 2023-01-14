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
)
