-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS passport_type
(
    id  SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS passport_type;
-- +goose StatementEnd
