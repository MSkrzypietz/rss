package main

import (
	"net/http"
)

func (app *application) getReadinessHandler(w http.ResponseWriter, r *http.Request) {
	err := app.writeJSON(w, http.StatusOK, envelope{"status": "ok"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
