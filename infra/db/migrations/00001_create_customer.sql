-- +goose Up
-- +goose StatementBegin
CREATE TABLE customer (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    city TEXT NOT NULL,
    street TEXT NOT NULL,
    zipcode TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE customer;
-- +goose StatementEnd
