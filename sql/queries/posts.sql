-- name: GetPostsForUser :many
SELECT p.* FROM posts p
    INNER JOIN feed_follows f ON f.feed_id=p.feed_id
    WHERE f.user_id=$1
    ORDER BY p.published_at DESC LIMIT $2;

-- name: CreatePost :one
INSERT INTO posts (id, title, url, description, published_at, feed_id)
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING *;
