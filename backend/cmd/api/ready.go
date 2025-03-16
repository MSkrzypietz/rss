package main

import (
	"net/http"
)

type ReadinessResponse struct {
	Status string `json:"status"`
}

func (app *application) getReadiness(w http.ResponseWriter, r *http.Request) {
	err := app.writeJSON(w, http.StatusOK, ReadinessResponse{Status: "ok"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
