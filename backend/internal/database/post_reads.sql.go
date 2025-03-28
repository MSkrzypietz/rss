// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: post_reads.sql

package database

import (
	"context"
)

const createPostRead = `-- name: CreatePostRead :one
INSERT INTO post_reads (user_id, post_id)
VALUES (?, ?)
ON CONFLICT(user_id, post_id) DO UPDATE SET user_id=post_reads.user_id, post_id=post_reads.post_id
    RETURNING id, created_at, updated_at, user_id, post_id
`

type CreatePostReadParams struct {
	UserID int64
	PostID int64
}

func (q *Queries) CreatePostRead(ctx context.Context, arg CreatePostReadParams) (PostRead, error) {
	row := q.db.QueryRowContext(ctx, createPostRead, arg.UserID, arg.PostID)
	var i PostRead
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.PostID,
	)
	return i, err
}
