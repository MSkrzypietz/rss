-- name: CreatePostRead :one
INSERT INTO post_reads (user_id, post_id)
VALUES (?, ?)
ON CONFLICT(user_id, post_id) DO UPDATE SET user_id=post_reads.user_id, post_id=post_reads.post_id
    RETURNING *;

