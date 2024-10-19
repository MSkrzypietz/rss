// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name, apikey)
VALUES (?, ?, ?, ?)
    RETURNING id, created_at, updated_at, name, apikey
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
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, created_at, updated_at, name, apikey
FROM users
WHERE apiKey=?
`

func (q *Queries) GetUser(ctx context.Context, apikey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, apikey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Apikey,
	)
	return i, err
}
