-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS access_tokens(
    id serial primary key,
    customer_id bigint not null,
    access_token text not null default '',
    refresh_token text not null default '',
    revoked_at timestamp null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    
    constraint fk_customers_access_tokens foreign key (customer_id) references customers(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS access_tokens;
-- +goose StatementEnd
