-- +goose Up
-- +goose StatementBegin
INSERT INTO staff(name, surname, phone, company_id, passport_type, passport_number, department_id) VALUES 
    ('Михаил', 'Ломоносов', '8-926-111-11-11', 1, 1, '9983232', 1),
    ('Алексей', 'Ломоносов', '8-926-222-22-22', 1, 1, '11112222', 1),
    ('Валерий', 'Гирин', '8-926-333-33-33', 2, 2, '131313', 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DELETE FROM staff
    -- WHERE passport_number IN ('9983232', '11112222', '131313'); 
-- +goose StatementEnd
