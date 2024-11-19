package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/wanomir/e"
)

func (db *PostgresDB) InsertUser(ctx context.Context, tgId, tgChatId int) (err error) {
	query := `INSERT INTO users (telegram_id, telegram_chat_id)
	VALUES ($1, $2)
	ON CONFLICT (telegram_id)
	DO
	UPDATE
	SET telegram_chat_id = $2;`

	if err = db.conn.QueryRow(ctx, query, tgId, tgChatId).Scan(); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return e.Wrap("failed to execute 'insert user'", err)
	}

	return nil
}

func (db *PostgresDB) GetChatIdByTelegramId(ctx context.Context, tgId int) (tgChatId int, err error) {
	query := `SELECT telegram_chat_id FROM users WHERE telegram_id = $1;`

	if err = db.conn.QueryRow(ctx, query, tgId).Scan(&tgChatId); errors.Is(err, pgx.ErrNoRows) {
		return 0, errNoRows
	}

	if err != nil {
		return 0, e.Wrap("failed to execute 'select chat id by telegram id'", err)
	}

	return tgChatId, nil
}

func (db *PostgresDB) GetTelegramIdByChatId(ctx context.Context, tgChatId int) (tgId int, err error) {
	query := `SELECT telegram_id FROM users WHERE telegram_chat_id = $1`

	if err = db.conn.QueryRow(ctx, query, tgChatId).Scan(&tgId); errors.Is(err, pgx.ErrNoRows) {
		return 0, errNoRows
	}

	if err != nil {
		return 0, e.Wrap("failed to execute 'select telegram id by chat id'", err)
	}

	return tgId, nil
}
