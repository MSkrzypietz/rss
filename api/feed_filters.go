package api

import (
	"encoding/json"
	"github.com/MSkrzypietz/rss/internal/database"
	"net/http"
	"strconv"
)

func (cfg *Config) getFeedFilters(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFilters, err := cfg.db.GetUserFeedFilters(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, feedFilters)
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

	feed, err := cfg.db.CreateFeedFilter(r.Context(), database.CreateFeedFilterParams{
		UserID:     user.ID,
		FeedID:     params.FeedID,
		FilterText: params.FilterText,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, feed)
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
