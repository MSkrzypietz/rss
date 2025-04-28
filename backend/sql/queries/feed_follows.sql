-- name: GetFeedFollows :many
SELECT * from feed_follows WHERE user_id=?;

-- name: GetFeedFollowers :many
SELECT user_id from feed_follows WHERE feed_id=?;

-- name: CreateFeedFollow :one
INSERT INTO feed_follows (user_id, feed_id)
VALUES (?, ?)
    RETURNING *;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id=? and user_id=?;
