-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS staff
(
    id  SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    phone VARCHAR(50),
    company_id INTEGER,
    passport_type INTEGER,
    passport_number VARCHAR(50),
    department_id INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS staff;
-- +goose StatementEnd
