package session

import (
	"github.com/google/uuid"
	"net/http"
)

func NewMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ID(r) == "" {
			http.SetCookie(w, &http.Cookie{
				Name:     "sessionID",
				Value:    uuid.New().String(),
				MaxAge:   2592000,
				Secure:   true,
				HttpOnly: true,
			})
		}
		next.ServeHTTP(w, r)
	})
}

func ID(r *http.Request) string {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return ""
	}
	return cookie.Value
}
