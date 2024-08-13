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

	mux.HandleFunc("GET /", cfg.GetLogin)
	mux.HandleFunc("POST /login", cfg.PostLogin)

	mux.HandleFunc("GET /posts", cfg.authenticate(cfg.GetPosts))

	mux.HandleFunc("GET /error", getError)

	return session.NewMiddleware(mux)
}
