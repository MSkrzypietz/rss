package main

import (
	"encoding/json"
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type GetUserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func mapGetUserResponse(dbUser database.User) GetUserResponse {
	return GetUserResponse{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		ApiKey:    dbUser.Apikey,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}

func (app *application) getAuthenticatedUser(w http.ResponseWriter, r *http.Request, user database.User) {
	err := app.writeJSON(w, http.StatusOK, mapGetUserResponse(user), nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user, err := app.db.CreateUser(r.Context(), database.CreateUserParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Apikey:    uuid.New().String(),
	})
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, mapGetUserResponse(user), nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
