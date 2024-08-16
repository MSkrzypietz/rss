package render

import (
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/MSkrzypietz/rss/views"
	"github.com/MSkrzypietz/rss/views/components"
	"net/http"
	"strconv"
)

func (cfg *Config) EditPage(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := cfg.db.GetFeeds(r.Context())
	if err != nil {
		views.Error("Unable to fetch feeds")
		return
	}

	var feedItems []components.SelectInputItem
	for _, feed := range feeds {
		feedItems = append(feedItems, components.SelectInputItem{
			ID:   strconv.Itoa(int(feed.ID)),
			Name: feed.Name,
		})
	}

	state := views.EditPageState{
		ActivePage: components.ActivePage{
			IsFeedActive: false,
			IsEditActive: true,
		},
		FeedSelectInputConfig: components.SelectInputConfig{
			Name:  "Feed",
			Items: feedItems,
		},
	}
	views.EditPage(state).Render(r.Context(), w)
}
