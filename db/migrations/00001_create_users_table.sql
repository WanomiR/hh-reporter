-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    telegram_id INTEGER PRIMARY KEY,
    telegram_chat_id INTEGER
);

CREATE TABLE IF NOT EXISTS vacancies (
    vacancy_id INTEGER PRIMARY KEY,
    vacancy_name VARCHAR(255),
    role_id INTEGER,
    experience VARCHAR(63),
    url TEXT,
    published_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS roles (
    role_id INTEGER PRIMARY KEY,
    role_name VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;

DROP TABLE vacancies;

DROP TABLE roles;
-- +goose StatementEnd
