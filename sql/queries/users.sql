-- name: GetUser :one
SELECT id, created_at, updated_at, name, apikey
FROM users
WHERE apiKey=$1;

-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, apikey)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
    RETURNING *;
