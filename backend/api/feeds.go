package api

import (
	"encoding/json"
	"github.com/MSkrzypietz/rss/internal/database"
	"net/http"
	"time"
)

type GetFeedResponse struct {
	ID            int64      `json:"id"`
	UserID        int64      `json:"user_id"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	LastFetchedAt *time.Time `json:"-"`
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

func (cfg *Config) getFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := cfg.db.GetUserFeeds(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, mapGetFeedResponses(feeds))
}

func (cfg *Config) createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest)
		return
	}

	// TODO: Should probably add a transaction...
	feed, err := cfg.db.CreateFeed(r.Context(), database.CreateFeedParams{
		Name:   params.Name,
		Url:    params.Url,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError)
		return
	}

	feedFollow, err := cfg.db.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError)
		return
	}

	type returnVals struct {
		Feed       GetFeedResponse       `json:"feed"`
		FeedFollow GetFeedFollowResponse `json:"feed_follow"`
	}
	respondWithJSON(w, http.StatusOK, returnVals{
		Feed:       mapGetFeedResponse(feed),
		FeedFollow: mapGetFeedFollowResponse(feedFollow),
	})
}
