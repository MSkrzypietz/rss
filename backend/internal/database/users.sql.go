// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name, apikey)
VALUES (?, ?, ?, ?)
    RETURNING id, created_at, updated_at, name, apikey, telegram_chat_id
`

type CreateUserParams struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Apikey    string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Apikey,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Apikey,
		&i.TelegramChatID,
	)
	return i, err
}

const getUserByApiKey = `-- name: GetUserByApiKey :one
SELECT id, created_at, updated_at, name, apikey, telegram_chat_id
FROM users
WHERE apiKey=?
`

func (q *Queries) GetUserByApiKey(ctx context.Context, apikey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByApiKey, apikey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Apikey,
		&i.TelegramChatID,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, created_at, updated_at, name, apikey, telegram_chat_id
FROM users
WHERE id=?
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Apikey,
		&i.TelegramChatID,
	)
	return i, err
}

const updateUserTelegramChatID = `-- name: UpdateUserTelegramChatID :exec
UPDATE users SET telegram_chat_id=?, updated_at=CURRENT_TIMESTAMP WHERE id=?
`

type UpdateUserTelegramChatIDParams struct {
	TelegramChatID sql.NullInt64
	ID             int64
}

func (q *Queries) UpdateUserTelegramChatID(ctx context.Context, arg UpdateUserTelegramChatIDParams) error {
	_, err := q.db.ExecContext(ctx, updateUserTelegramChatID, arg.TelegramChatID, arg.ID)
	return err
}
