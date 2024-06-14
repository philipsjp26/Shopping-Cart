-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS carts (
    id serial primary key,
    product_id bigint not null,
    shopping_cart_id bigint not null,
    constraint fk_carts_products foreign key (product_id) references products(id),
    constraint fk_shopping_carts_carts foreign key (shopping_cart_id) references shopping_carts(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS carts;
-- +goose StatementEnd
