package main

import (
	"github.com/MSkrzypietz/rss/internal/database"
	"net/http"
	"strconv"
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

func (app *application) getUnreadPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	qs := r.URL.Query()
	searchText := app.readString(qs, "searchText", "")
	feedIDs, err := app.readCSVInt64s(qs, "feedIDs", []int64{})
	if err != nil {
		return
	}

	posts, err := app.db.GetUnreadPostsForUser(r.Context(), database.GetUnreadPostsForUserParams{
		SearchText:    "%" + searchText + "%",
		FeedIDsLength: len(feedIDs),
		FeedIDs:       feedIDs,
		UserID:        user.ID,
		Limit:         postGetterLimit,
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, mapGetUnreadPostResponses(posts), nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) markPostAsRead(w http.ResponseWriter, r *http.Request, user database.User) {
	postID, err := strconv.ParseInt(r.PathValue("postID"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	_, err = app.db.CreatePostRead(r.Context(), database.CreatePostReadParams{
		UserID: user.ID,
		PostID: postID,
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, struct{}{}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
