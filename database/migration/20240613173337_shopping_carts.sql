-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS shopping_carts(
    id serial primary key,
    customer_id bigint not null,
    constraint fk_customers_shopping_carts foreign key (customer_id) references customers(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shopping_carts;
-- +goose StatementEnd
