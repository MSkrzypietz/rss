-- name: GetUserByApiKey :one
SELECT id, created_at, updated_at, name, apikey, telegram_chat_id
FROM users
WHERE apiKey=?;

-- name: GetUserByID :one
SELECT id, created_at, updated_at, name, apikey, telegram_chat_id
FROM users
WHERE id=?;

-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name, apikey)
VALUES (?, ?, ?, ?)
    RETURNING *;

-- name: UpdateUserTelegramChatID :exec
UPDATE users SET telegram_chat_id=?, updated_at=CURRENT_TIMESTAMP WHERE id=?;
