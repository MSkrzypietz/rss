package render

import (
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/MSkrzypietz/rss/views"
	"github.com/MSkrzypietz/rss/views/components"
	"net/http"
	"strconv"
)

const postGetterLimit = 10

func (cfg *Config) FeedPage(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := cfg.db.GetUnreadPostsForUser(r.Context(), database.GetUnreadPostsForUserParams{
		UserID: user.ID,
		Limit:  postGetterLimit,
	})

	if err != nil {
		views.Error("Unable to fetch new posts").Render(r.Context(), w)
		return
	}

	state := views.FeedPageState{
		Posts: posts,
		ActivePage: components.ActivePage{
			IsFeedActive: true,
			IsEditActive: false,
		},
	}
	views.FeedPage(state).Render(r.Context(), w)
}

func (cfg *Config) MarkPostAsRead(w http.ResponseWriter, r *http.Request, user database.User) {
	postID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		views.Error("Bad request").Render(r.Context(), w)
		return
	}

	_, err = cfg.db.CreatePostRead(r.Context(), database.CreatePostReadParams{
		UserID: user.ID,
		PostID: int64(postID),
	})
	if err != nil {
		views.Error("Internal server error").Render(r.Context(), w)
		return
	}
}
