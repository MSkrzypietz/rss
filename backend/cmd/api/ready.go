package main

import (
	"net/http"
)

type ReadinessResponse struct {
	Status string `json:"status"`
}

func getReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, ReadinessResponse{Status: "ok"})
}

func getError(w http.ResponseWriter, r *http.Request) {
	respondWithErrorText(w, http.StatusInternalServerError, "Something went wrong")
}
