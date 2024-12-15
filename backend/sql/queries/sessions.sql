-- name: GetUserBySession :one
SELECT u.* FROM sessions s
    INNER JOIN users u ON u.id=s.user_id
    WHERE token=? AND expires_at>CURRENT_TIMESTAMP;

-- name: CreateSession :one
INSERT INTO sessions (token, expires_at, user_id)
VALUES (?, ?, ?)
    RETURNING *;
