package repository

import (
	"context"

	"github.com/wanomir/hh-reporter/internal/dto"
)

// TODO: finish writing basic CRUD's
// https://arc.net/l/quote/vylfjpww
// INSERT INTO vacancies (vacancy_id, vacancy_name, role_id,  experience, url, published_at)
// VALUES
//     (
//       1, 'vacancy', 1, 'from0To3', 'https://some-website', NOW()
//     )
// ON CONFLICT ()
// DO
// UPDATE
// SET content=%(content)s, last_analyzed = NOW();

func (db *PostgresDB) InsertVacancy(ctx context.Context, vacancy dto.Vacancy) {}
