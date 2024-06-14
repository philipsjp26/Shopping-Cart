-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS customers(
    id serial primary key,
    email varchar unique not null,
    phone varchar null,
    name varchar not null,
    username varchar unique not null,
    password varchar not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- index
CREATE INDEX IF NOT EXISTS customers_username_email_created_at_idx ON customers(username, email, created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS customers_username_email_created_at_idx;
DROP TABLE IF EXISTS customers;
-- +goose StatementEnd
