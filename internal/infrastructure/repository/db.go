package repository

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var errNoRows = errors.New("no rows in result set")

type PostgresDB struct {
	conn    *pgxpool.Pool
	timeout time.Duration
}

func NewPostgresDB(pool *pgxpool.Pool) *PostgresDB {
	return &PostgresDB{
		conn:    pool,
		timeout: 1 * time.Second,
	}
}
