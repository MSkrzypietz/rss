package api

import (
	"database/sql"
	"github.com/MSkrzypietz/rss/internal/database"
	"net/http"
	"time"
)

type Config struct {
	db         *database.Queries
	httpClient *http.Client
}

func NewConfig(db *sql.DB) *Config {
	return &Config{
		db:         database.New(db),
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}
}

func (cfg *Config) Handlers() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/login", cfg.login)

	mux.HandleFunc("GET /users", cfg.authenticate(cfg.getAuthenticatedUser))
	mux.HandleFunc("POST /users", cfg.createUser)

	mux.HandleFunc("GET /feeds", cfg.getFeeds)
	mux.HandleFunc("POST /feeds", cfg.authenticate(cfg.createFeed))

	mux.HandleFunc("GET /feed_follows", cfg.authenticate(cfg.getFeedFollows))
	mux.HandleFunc("POST /feed_follows", cfg.authenticate(cfg.createFeedFollow))
	mux.HandleFunc("DELETE /feed_follows/{feedFollowID}", cfg.authenticate(cfg.deleteFeedFollow))

	mux.HandleFunc("GET /feed_filters", cfg.authenticate(cfg.getFeedFilters))
	mux.HandleFunc("POST /feed_filters", cfg.authenticate(cfg.createFeedFilter))
	mux.HandleFunc("DELETE /feed_filters/{feedFilterID}", cfg.authenticate(cfg.deleteFeedFilter))

	mux.HandleFunc("GET /posts", cfg.authenticate(cfg.getUnreadPosts))

	mux.HandleFunc("GET /readiness", getReadiness)
	mux.HandleFunc("GET /err", getError)

	return mux
}
