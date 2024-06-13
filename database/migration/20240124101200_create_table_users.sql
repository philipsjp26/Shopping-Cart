-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `users` (
    `id` BIGINT NOT NULL auto_increment,
    `name` VARCHAR(150) NOT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
