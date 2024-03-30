-- +goose Up
-- +goose StatementBegin
INSERT INTO department(name, phone) VALUES 
    ('Административный департамент', '8-495-111-11-11'),
    ('Департамент бюджетной политики', '8-495-222-11-11'),
    ('It департамент', '8-495-333-11-11');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DELETE FROM department
--     WHERE name IN ('Административный департамент', 'Департамент бюджетной политики', 'It департамент');
-- +goose StatementEnd
