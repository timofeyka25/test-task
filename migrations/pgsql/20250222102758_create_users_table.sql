-- +goose Up
CREATE TABLE IF NOT EXISTS users
(
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at    TIMESTAMPTZ DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS users;
