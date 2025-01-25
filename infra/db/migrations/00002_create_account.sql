-- +goose Up
-- +goose StatementBegin
CREATE TABLE account (
    id INTEGER PRIMARY KEY AUTOINCREMENT, 
    id_customer INTEGER NOT NULL, 
    type_account TEXT NOT NULL, 
    balance REAL DEFAULT 0, 
    status INTEGER DEFAULT 1 NOT NULL,
    FOREIGN KEY (id_customer) REFERENCES customer(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE account;
-- +goose StatementEnd
