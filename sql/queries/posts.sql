-- name: GetUnreadPostsForUser :many
SELECT p.* FROM posts p
    INNER JOIN feed_follows f ON f.feed_id=p.feed_id
    LEFT JOIN post_reads pr ON pr.user_id=f.user_id AND pr.post_id=p.id
    WHERE f.user_id=? AND pr.id IS NULL
    ORDER BY p.published_at DESC LIMIT ?;

-- name: CreatePost :one
INSERT INTO posts (title, url, description, published_at, feed_id)
    VALUES (?, ?, ?, ?, ?)
    RETURNING *;
