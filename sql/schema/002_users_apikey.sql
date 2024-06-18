-- +goose Up
ALTER TABLE users ADD COLUMN apikey VARCHAR(64) NOT NULL;
CREATE UNIQUE INDEX users_apikey_uniq_idx ON users(apikey);

-- +goose Down
DROP INDEX users_apikey_uniq_idx;
ALTER TABLE users DROP COLUMN apikey;
