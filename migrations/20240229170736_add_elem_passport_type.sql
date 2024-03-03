-- +goose Up
-- +goose StatementBegin
INSERT INTO passport_type(name) VALUES 
    ('Гражданский паспорт'), 
    ('Заграничный паспорт'), 
    ('Дипломатический паспорт');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM passport_type
    WHERE name IN ('Гражданский паспорт', 'Заграничный паспорт', 'Дипломатический паспорт');
-- +goose StatementEnd
