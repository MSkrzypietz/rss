package main

import (
	"encoding/json"
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest)
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Apikey:    uuid.New().String(),
	})
	if err != nil {
		respondWithErrorText(w, http.StatusInternalServerError, "could not create user")
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (cfg *apiConfig) getUsers(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, user)
}
