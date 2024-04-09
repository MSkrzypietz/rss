package main

import (
	"github.com/MSkrzypietz/rss/internal/database"
	"net/http"
)

const postGetterLimit = 10

func (cfg *apiConfig) getPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := cfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  postGetterLimit,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
