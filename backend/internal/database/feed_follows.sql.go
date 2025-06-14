// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feed_follows.sql

package database

import (
	"context"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
INSERT INTO feed_follows (user_id, feed_id)
VALUES (?, ?)
    RETURNING id, created_at, updated_at, user_id, feed_id
`

type CreateFeedFollowParams struct {
	UserID int64
	FeedID int64
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow, arg.UserID, arg.FeedID)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id=? and user_id=?
`

type DeleteFeedFollowParams struct {
	ID     int64
	UserID int64
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, arg.ID, arg.UserID)
	return err
}

const getFeedFollowers = `-- name: GetFeedFollowers :many
SELECT user_id from feed_follows WHERE feed_id=?
`

func (q *Queries) GetFeedFollowers(ctx context.Context, feedID int64) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowers, feedID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var user_id int64
		if err := rows.Scan(&user_id); err != nil {
			return nil, err
		}
		items = append(items, user_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFeedFollows = `-- name: GetFeedFollows :many
SELECT id, created_at, updated_at, user_id, feed_id from feed_follows WHERE user_id=?
`

func (q *Queries) GetFeedFollows(ctx context.Context, userID int64) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
