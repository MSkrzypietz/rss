package main

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

func (app *application) authenticate(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, found := app.getUserByAuthHeader(r)
		if found {
			handler(w, r, user)
			return
		}

		user, found = app.getUserBySessionCookie(r)
		if found {
			handler(w, r, user)
			return
		}

		app.authenticationRequiredResponse(w, r)
	}
}

func (app *application) getUserByAuthHeader(r *http.Request) (database.User, bool) {
	authHeader := r.Header.Get("Authorization")
	apiKey, found := strings.CutPrefix(authHeader, "ApiKey ")
	if !found {
		return database.User{}, false
	}

	user, err := app.db.GetUser(r.Context(), apiKey)
	if err != nil {
		return database.User{}, false
	}
	return user, true
}

func (app *application) getUserBySessionCookie(r *http.Request) (database.User, bool) {
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		return database.User{}, false
	}

	user, err := app.db.GetUserBySession(r.Context(), cookie.Value)
	if err != nil {
		return database.User{}, false
	}
	return user, true
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ApiKey string `json:"apiKey"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user, err := app.db.GetUser(r.Context(), params.ApiKey)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	token := uuid.New().String()
	expiryInSeconds := 2592000
	_, err = app.db.CreateSession(r.Context(), database.CreateSessionParams{
		Token:     token,
		ExpiresAt: time.Now().Add(time.Duration(expiryInSeconds) * time.Second),
		UserID:    user.ID,
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
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

	err = app.writeJSON(w, http.StatusOK, mapGetUserResponse(user), nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
