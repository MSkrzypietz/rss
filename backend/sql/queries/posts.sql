-- name: GetUnreadPostsForUser :many
SELECT p.*, f.name as feed_name FROM posts p
    INNER JOIN feed_follows ff ON ff.feed_id=p.feed_id
    INNER JOIN feeds f ON f.id=p.feed_id
    LEFT JOIN post_reads pr ON pr.user_id=ff.user_id AND pr.post_id=p.id
    WHERE ff.user_id=? AND pr.id IS NULL AND
        p.title LIKE @searchText AND
        (@feedIDsLength=0 OR p.feed_id IN (sqlc.slice('feedIDs')))
    ORDER BY p.published_at DESC LIMIT ?;

-- name: CreatePost :one
INSERT INTO posts (title, url, description, published_at, feed_id)
    VALUES (?, ?, ?, ?, ?)
    ON CONFLICT(url) DO UPDATE SET url=posts.url
    RETURNING *;
