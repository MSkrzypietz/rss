package main

import (
	"encoding/json"
	"github.com/MSkrzypietz/rss/internal/database"
	"net/http"
	"strconv"
	"time"
)

type GetFeedFilterResponse struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	FeedID     int64     `json:"feed_id"`
	FilterText string    `json:"filter_text"`
	Active     bool      `json:"active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func mapGetFeedFilterResponses(dbFeedFilters []database.FeedFilter) []GetFeedFilterResponse {
	var responses []GetFeedFilterResponse
	for _, dbFeedFilter := range dbFeedFilters {
		responses = append(responses, mapGetFeedFilterResponse(dbFeedFilter))
	}
	return responses
}

func mapGetFeedFilterResponse(dbFeedFilter database.FeedFilter) GetFeedFilterResponse {
	return GetFeedFilterResponse{
		ID:         dbFeedFilter.ID,
		UserID:     dbFeedFilter.UserID,
		FeedID:     dbFeedFilter.FeedID,
		FilterText: dbFeedFilter.FilterText,
		Active:     dbFeedFilter.Active,
		CreatedAt:  dbFeedFilter.CreatedAt,
		UpdatedAt:  dbFeedFilter.UpdatedAt,
	}
}

func (cfg *Config) getFeedFilters(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFilters, err := cfg.db.GetUserFeedFilters(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, mapGetFeedFilterResponses(feedFilters))
}

func (cfg *Config) createFeedFilter(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID     int64  `json:"feed_id"`
		FilterText string `json:"filter_text"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest)
		return
	}

	feedFilter, err := cfg.db.CreateFeedFilter(r.Context(), database.CreateFeedFilterParams{
		UserID:     user.ID,
		FeedID:     params.FeedID,
		FilterText: params.FilterText,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, mapGetFeedFilterResponse(feedFilter))
}

func (cfg *Config) deleteFeedFilter(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFilterID, err := strconv.ParseInt(r.PathValue("feedFilterID"), 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest)
		return
	}

	err = cfg.db.DeleteFeedFilter(r.Context(), feedFilterID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
