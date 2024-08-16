package render

import (
	"github.com/MSkrzypietz/rss/render/session"
	"github.com/MSkrzypietz/rss/views"
	"net/http"
)

func (cfg *Config) IndexPage(w http.ResponseWriter, r *http.Request) {
	user, ok := cfg.getUserSession(r)
	if !ok {
		views.Login().Render(r.Context(), w)
		return
	}
	cfg.FeedPage(w, r, user)
}

func (cfg *Config) PostLogin(w http.ResponseWriter, r *http.Request) {
	apiKey := r.FormValue("apiKey")
	user, err := cfg.db.GetUser(r.Context(), apiKey)
	if err != nil {
		views.Error("Incorrect API key").Render(r.Context(), w)
		return
	}

	w.Header().Set("HX-Redirect", "/feed")

	sessionID := session.ID(r)
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.userSessions[sessionID] = user
}
