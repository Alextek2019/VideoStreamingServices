-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE IF NOT EXISTS users.user(
                                         id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    login VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(60) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users.user;
DROP SCHEMA IF EXISTS users;
-- +goose StatementEnd
