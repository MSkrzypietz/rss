-- +goose Up
ALTER TABLE users ADD COLUMN telegram_chat_id INTEGER;

-- +goose Down
ALTER TABLE users DROP COLUMN telegram_chat_id;
