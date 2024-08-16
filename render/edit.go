package render

import (
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/MSkrzypietz/rss/views"
	"github.com/MSkrzypietz/rss/views/components"
	"net/http"
)

func (cfg *Config) EditPage(w http.ResponseWriter, r *http.Request, user database.User) {
	state := views.EditPageState{
		ActivePage: components.ActivePage{
			IsFeedActive: false,
			IsEditActive: true,
		},
	}
	views.EditPage(state).Render(r.Context(), w)
}
