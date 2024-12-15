package api

import (
	"encoding/json"
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

const sessionCookieName = "sessionID"

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *Config) authenticate(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, found := cfg.getUserByAuthHeader(r)
		if found {
			handler(w, r, user)
			return
		}

		user, found = cfg.getUserBySessionCookie(r)
		if found {
			handler(w, r, user)
			return
		}

		respondWithError(w, http.StatusUnauthorized)
	}
}

func (cfg *Config) getUserByAuthHeader(r *http.Request) (database.User, bool) {
	authHeader := r.Header.Get("Authorization")
	apiKey, found := strings.CutPrefix(authHeader, "ApiKey ")
	if !found {
		return database.User{}, false
	}

	user, err := cfg.db.GetUser(r.Context(), apiKey)
	if err != nil {
		return database.User{}, false
	}
	return user, true
}

func (cfg *Config) getUserBySessionCookie(r *http.Request) (database.User, bool) {
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		return database.User{}, false
	}

	user, err := cfg.db.GetUserBySession(r.Context(), cookie.Value)
	if err != nil {
		return database.User{}, false
	}
	return user, true
}

func (cfg *Config) login(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ApiKey string `json:"apiKey"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest)
		return
	}

	user, err := cfg.db.GetUser(r.Context(), params.ApiKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError)
		return
	}

	token := uuid.New().String()
	expiryInSeconds := 2592000
	_, err = cfg.db.CreateSession(r.Context(), database.CreateSessionParams{
		Token:     token,
		ExpiresAt: time.Now().Add(time.Duration(expiryInSeconds) * time.Second),
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    token,
		Path:     "/",
		MaxAge:   expiryInSeconds,
		Secure:   true,
		HttpOnly: true,
	})

	respondWithJSON(w, http.StatusOK, mapGetUserResponse(user))
}
