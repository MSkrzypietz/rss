package render

import (
	"github.com/MSkrzypietz/rss/render/session"
	"github.com/MSkrzypietz/rss/views"
	"net/http"
)

func (cfg *Config) GetLogin(w http.ResponseWriter, r *http.Request) {
	views.Login().Render(r.Context(), w)
}

func (cfg *Config) PostLogin(w http.ResponseWriter, r *http.Request) {
	apiKey := r.FormValue("apiKey")
	user, err := cfg.db.GetUser(r.Context(), apiKey)
	if err != nil {
		views.Error("Incorrect API key").Render(r.Context(), w)
		return
	}

	w.Header().Set("HX-Redirect", "/posts")

	sessionID := session.ID(r)
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.userSessions[sessionID] = user
}
