-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders( 
    id serial primary key,
    transaction_id varchar not null unique,
    customer_id bigint not null,
    status varchar(100) not null,
    payment_method varchar(100) not null,
    order_date timestamp DEFAULT CURRENT_TIMESTAMP,
    constraint fk_orders_customers foreign key (customer_id) references customers(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
