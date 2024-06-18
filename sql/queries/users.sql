-- name: GetUser :one
SELECT id, created_at, updated_at, name, apikey
FROM users
WHERE apiKey=?;

-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name, apikey)
VALUES (?, ?, ?, ?)
    RETURNING *;
