package render

import (
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/MSkrzypietz/rss/render/session"
	"github.com/MSkrzypietz/rss/views"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *Config) authenticate(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := cfg.getUserSession(r)
		if !ok {
			views.Error("Unauthorized").Render(r.Context(), w)
			return
		}

		handler(w, r, user)
	}
}

func (cfg *Config) getUserSession(r *http.Request) (database.User, bool) {
	sessionID := session.ID(r)
	cfg.mu.RLock()
	defer cfg.mu.RUnlock()
	user, ok := cfg.userSessions[sessionID]
	return user, ok
}
