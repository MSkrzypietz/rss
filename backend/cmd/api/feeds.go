package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MSkrzypietz/rss/internal/database"
)

type GetFeedResponse struct {
	ID            int64      `json:"id"`
	UserID        int64      `json:"user_id"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func mapGetFeedResponses(dbFeeds []database.Feed) []GetFeedResponse {
	var responses []GetFeedResponse
	for _, dbFeed := range dbFeeds {
		responses = append(responses, mapGetFeedResponse(dbFeed))
	}
	return responses
}

func mapGetFeedResponse(dbFeed database.Feed) GetFeedResponse {
	var lastFetchedAt *time.Time
	if dbFeed.LastFetchedAt.Valid {
		lastFetchedAt = &dbFeed.LastFetchedAt.Time
	}
	return GetFeedResponse{
		ID:            dbFeed.ID,
		UserID:        dbFeed.UserID,
		Name:          dbFeed.Name,
		Url:           dbFeed.Url,
		LastFetchedAt: lastFetchedAt,
		CreatedAt:     dbFeed.CreatedAt,
		UpdatedAt:     dbFeed.UpdatedAt,
	}
}

func (app *application) listFeedsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := app.db.GetUserFeeds(r.Context(), user.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"feeds": mapGetFeedResponses(feeds)}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// TODO: Test feed url if it reachable / valid rss feed
func (app *application) createFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// TODO: Should probably add a transaction...
	feed, err := app.db.CreateFeed(r.Context(), database.CreateFeedParams{
		Name:   params.Name,
		Url:    params.Url,
		UserID: user.ID,
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	feedFollow, err := app.db.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{
		"feed":        mapGetFeedResponse(feed),
		"feed_follow": mapGetFeedFollowResponse(feedFollow),
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) fetchFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedID, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	feed, err := app.db.GetFeedByID(r.Context(), feedID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	data, err := json.Marshal(feed)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.nc.Publish(topicFeedFetch, data)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
