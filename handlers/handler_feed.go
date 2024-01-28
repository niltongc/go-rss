package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	database "github.com/niltongc/rssagg/internal/database"
	"github.com/niltongc/rssagg/models"
)

// type ApiConfig struct {
// 	DB *database.Queries
// }

func (cfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	RespondWithJSON(w, http.StatusCreated, models.DatabaseFeedToFeed(feed))
}

func (cfg *ApiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}

	RespondWithJSON(w, http.StatusCreated, models.DatabaseFeedsToFeeds(feeds))
}
