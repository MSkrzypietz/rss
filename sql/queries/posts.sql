-- name: GetPostsForUser :many
SELECT p.* FROM posts p
    INNER JOIN feed_follows f ON f.feed_id=p.feed_id
    WHERE f.user_id=?
    ORDER BY p.published_at DESC LIMIT ?;

-- name: CreatePost :one
INSERT INTO posts (title, url, description, published_at, feed_id)
    VALUES (?, ?, ?, ?, ?)
    RETURNING *;
