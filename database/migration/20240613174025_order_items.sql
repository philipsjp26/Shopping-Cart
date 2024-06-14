-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_items(
    order_id bigint not null,
    product_id bigint not null,
    quantity bigint not null,
    price decimal(10,2) not null,
    constraint fk_order_items_orders foreign key (order_id) references orders(id),
    constraint fk_products_orders foreign key (product_id) references products(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;
-- +goose StatementEnd
