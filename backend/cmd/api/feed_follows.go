package main

import (
	"encoding/json"
	"github.com/MSkrzypietz/rss/internal/database"
	"net/http"
	"strconv"
	"time"
)

type GetFeedFollowResponse struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	FeedID    int64     `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func mapGetFeedFollowResponses(dbFeedFollows []database.FeedFollow) []GetFeedFollowResponse {
	var responses []GetFeedFollowResponse
	for _, dbFeedFollow := range dbFeedFollows {
		responses = append(responses, mapGetFeedFollowResponse(dbFeedFollow))
	}
	return responses
}

func mapGetFeedFollowResponse(dbFeedFollow database.FeedFollow) GetFeedFollowResponse {
	return GetFeedFollowResponse{
		ID:        dbFeedFollow.ID,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
	}
}

func (app *application) getFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := app.db.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, mapGetFeedFollowResponses(feedFollows), nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID int64 `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	feedFollow, err := app.db.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, mapGetFeedFollowResponse(feedFollow), nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID, err := strconv.ParseInt(r.PathValue("feedFollowID"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.db.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
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
