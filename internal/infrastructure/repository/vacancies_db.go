package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/wanomir/e"
	"github.com/wanomir/hh-reporter/internal/dto"
)

func (db *PostgresDB) InsertVacancy(ctx context.Context, v dto.Vacancy) (err error) {
	query := `INSERT INTO vacancies (vacancy_id, vacancy_name, role_id, experience, url, published_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (vacancy_id)
	DO
	UPDATE
	SET vacancy_name = $2, role_id = $3, experience = $4, url = $5, published_at = $6;`

	if err = db.conn.QueryRow(ctx, query,
		v.VacancyId,
		v.VacancyName,
		v.RoleId,
		v.Experience,
		v.Url,
		v.PublisthedAt,
	).Scan(); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return e.Wrap("failed to execute 'insert vacancy'", err)
	}

	return nil
}

func (db *PostgresDB) GetVacancyById(ctx context.Context, id int) (v *dto.Vacancy, err error) {
	v = new(dto.Vacancy)

	query := `SELECT vacancy_id, vacancy_name, role_id, experience, url, published_at
	FROM vacancies
	WHERE vacancy_id = $1;`

	if err = db.conn.QueryRow(ctx, query, id).Scan(
		&v.VacancyId, &v.VacancyName, &v.RoleId, &v.Experience, &v.Url, &v.PublisthedAt,
	); errors.Is(err, pgx.ErrNoRows) {
		return nil, errNoRows
	}

	if err != nil {
		return nil, e.Wrap("failed to execute 'select vacany by vacancy id'", err)
	}

	return v, nil
}
