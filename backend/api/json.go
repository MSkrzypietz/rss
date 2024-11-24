package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int) {
	respondWithErrorText(w, code, http.StatusText(code))
}

func respondWithErrorText(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5xx error: %v\n", msg)
	}
	type returnVals struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, returnVals{Error: msg})
}
