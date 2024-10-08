-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE IF NOT EXISTS users.user(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    login VARCHAR(100) NOT NULL,
    hashed_password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP EXTENSION IF EXISTS "uuid-ossp";
    DROP TABLE users.user;
    DROP SCHEMA IF EXISTS users;
-- +goose StatementEnd
