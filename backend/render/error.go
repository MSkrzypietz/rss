package render

import (
	"github.com/MSkrzypietz/rss/views"
	"net/http"
)

func getError(w http.ResponseWriter, r *http.Request) {
	errorMessage := r.URL.Query().Get("message")
	if errorMessage == "" {
		errorMessage = "Unknown error"
	}
	views.Error(errorMessage).Render(r.Context(), w)
}
