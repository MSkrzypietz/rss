-- +goose Up
CREATE TABLE feed_filters (
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    filter_text TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT 1,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id INTEGER NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feed_filters;