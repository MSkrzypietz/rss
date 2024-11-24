package api

import (
	"github.com/MSkrzypietz/rss/internal/database"
	"net/http"
	"strings"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *Config) authenticate(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		apiKey, _ := strings.CutPrefix(authHeader, "ApiKey ")

		user, err := cfg.db.GetUser(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized)
			return
		}

		handler(w, r, user)
	}
}
