package api

import (
	"github.com/MSkrzypietz/rss/internal/database"
	"net/http"
	"time"
)

const postGetterLimit = 10

type GetUnreadPostResponse struct {
	ID          int64      `json:"id"`
	FeedID      int64      `json:"feed_id"`
	FeedName    string     `json:"feed_name"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func mapGetUnreadPostResponses(dbUnreadPosts []database.GetUnreadPostsForUserRow) []GetUnreadPostResponse {
	var responses []GetUnreadPostResponse
	for _, dbUnreadPost := range dbUnreadPosts {
		responses = append(responses, mapGetUnreadPostResponse(dbUnreadPost))
	}
	return responses
}

func mapGetUnreadPostResponse(dbUnreadPost database.GetUnreadPostsForUserRow) GetUnreadPostResponse {
	var description *string
	if dbUnreadPost.Description.Valid {
		description = &dbUnreadPost.Description.String
	}
	var publishedAt *time.Time
	if dbUnreadPost.PublishedAt.Valid {
		publishedAt = &dbUnreadPost.PublishedAt.Time
	}
	return GetUnreadPostResponse{
		ID:          dbUnreadPost.ID,
		FeedID:      dbUnreadPost.FeedID,
		FeedName:    dbUnreadPost.FeedName,
		Title:       dbUnreadPost.Title,
		Url:         dbUnreadPost.Url,
		Description: description,
		PublishedAt: publishedAt,
		CreatedAt:   dbUnreadPost.CreatedAt,
		UpdatedAt:   dbUnreadPost.UpdatedAt,
	}
}

func (cfg *Config) getUnreadPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := cfg.db.GetUnreadPostsForUser(r.Context(), database.GetUnreadPostsForUserParams{
		UserID: user.ID,
		Limit:  postGetterLimit,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, mapGetUnreadPostResponses(posts))
}