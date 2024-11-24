-- name: CreatePostRead :one
INSERT INTO post_reads (user_id, post_id)
VALUES (?, ?)
    RETURNING *;

