package api

import (
	"net/http"
)

func getReadiness(w http.ResponseWriter, r *http.Request) {
	type returnVals struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, returnVals{Status: "ok"})
}

func getError(w http.ResponseWriter, r *http.Request) {
	respondWithErrorText(w, http.StatusInternalServerError, "Something went wrong")
}
