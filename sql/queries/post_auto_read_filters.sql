-- name: GetPostAutoReadFilters :many
SELECT * FROM post_auto_read_filters
WHERE user_id=? AND post_id=?;

