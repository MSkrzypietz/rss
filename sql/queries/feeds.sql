-- name: GetFeeds :many
SELECT * from feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at NULLS FIRST LIMIT ?;

-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id)
VALUES (?, ?, ?)
    RETURNING *;

-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at=CURRENT_TIMESTAMP, updated_at=CURRENT_TIMESTAMP WHERE id=?;