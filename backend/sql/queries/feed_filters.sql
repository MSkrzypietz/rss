-- name: GetFeedFilters :many
SELECT * FROM feed_filters
WHERE feed_id=?;

-- name: GetUserFeedFilters :many
SELECT * FROM feed_filters
WHERE user_id=?;

-- name: CreateFeedFilter :one
INSERT INTO feed_filters (user_id, feed_id, filter_text)
VALUES (?, ?, ?)
    RETURNING *;

-- name: DeleteFeedFilter :exec
DELETE FROM feed_filters WHERE id=?;
