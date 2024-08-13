package render

import (
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/MSkrzypietz/rss/views"
	"net/http"
)

const postGetterLimit = 10

func (cfg *Config) GetPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := cfg.db.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  postGetterLimit,
	})

	if err != nil {
		views.Error("Unable to fetch new posts").Render(r.Context(), w)
		return
	}

	views.Posts(posts).Render(r.Context(), w)
}
