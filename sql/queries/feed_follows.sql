-- name: GetFeedFollows :many
SELECT * from feed_follows WHERE user_id=$1;

-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, user_id, feed_id)
VALUES ($1, $2, $3)
    RETURNING *;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id=$1 and user_id=$2;
