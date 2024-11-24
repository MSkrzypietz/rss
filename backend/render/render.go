package render

import (
	"database/sql"
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/MSkrzypietz/rss/render/session"
	"net/http"
	"sync"
)

type Config struct {
	db           *database.Queries
	userSessions map[string]database.User
	mu           sync.RWMutex
}

func NewConfig(db *sql.DB) *Config {
	return &Config{
		db:           database.New(db),
		userSessions: map[string]database.User{},
		mu:           sync.RWMutex{},
	}
}

func (cfg *Config) Handlers() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", cfg.IndexPage)
	mux.HandleFunc("POST /login", cfg.PostLogin)

	mux.HandleFunc("GET /feed", cfg.authenticate(cfg.FeedPage))
	mux.HandleFunc("POST /posts/{id}/read", cfg.authenticate(cfg.MarkPostAsRead))

	mux.HandleFunc("GET /edit", cfg.authenticate(cfg.EditPage))

	mux.HandleFunc("GET /error", getError)

	return session.NewMiddleware(mux)
}
