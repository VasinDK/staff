-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS department
(
    id  SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    phone VARCHAR(50)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS department;
-- +goose StatementEnd
