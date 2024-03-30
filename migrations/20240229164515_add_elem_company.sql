-- +goose Up
-- +goose StatementBegin
INSERT INTO company(name) VALUES 
    ('ООО Дружок'), 
    ('АО Ракета'), 
    ('ЗАО Радуга');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DELETE FROM company
--     WHERE name IN ('ООО Дружок', 'АО Ракета', 'ЗАО Радуга');
-- +goose StatementEnd
