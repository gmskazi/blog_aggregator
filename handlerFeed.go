package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gmskazi/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't decode parameters: %v", err))
		return
	}

	userIDNullUUID := uuid.NullUUID{
		UUID:  user.ID,
		Valid: true,
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedName:  params.Name,
		Url:       params.URL,
		UserID:    userIDNullUUID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserFeeds(feed))
}
