-- +goose Up
-- +goose StatementBegin
-- +goose StatementEnd
CREATE TABLE users
(
    id      INTEGER      NOT NULL,
    name    VARCHAR(250) NOT NULL,
    balance INTEGER,
    PRIMARY KEY (id),
    CONSTRAINT balance CHECK (balance >= 0)

);
-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
