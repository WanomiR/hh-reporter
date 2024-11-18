-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    telegram_id INTEGER,
    telegram_chat_id INTEGER
);

CREATE TABLE IF NOT EXISTS vacancies (
    id SERIAL PRIMARY KEY,
    vacancy_id INTEGER,
    vacancy_name VARCHAR(255),
    role_id INTEGER,
    experience VARCHAR(63),
    url TEXT,
    published_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    role_id INTEGER,
    role_name VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;

DROP TABLE vacancies;

DROP TABLE roles;
-- +goose StatementEnd
