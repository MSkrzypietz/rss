package main

import (
	"github.com/MSkrzypietz/rss/internal/vcs"
	"net/http"
)

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"version": vcs.Version(),
		},
	}

	respondWithJSON(w, http.StatusOK, env)
}
