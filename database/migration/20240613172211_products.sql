-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
    id serial primary key,
    category_id integer not null,
    name varchar not null,
    description text null default '',
    quantity bigint not null,
    price decimal(10,2) not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    constraint fk_cateogries_products foreign key (category_id) REFERENCES categories (id)
);

-- index
CREATE INDEX IF NOT EXISTS products_name_idx ON products(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS products_name_idx;
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
