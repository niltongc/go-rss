package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	database "github.com/niltongc/rssagg/internal/database"
	"github.com/niltongc/rssagg/models"
)

// type ApiConfig struct {
// 	DB *database.Queries
// }

func (cfg *ApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feedfollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	RespondWithJSON(w, http.StatusCreated, models.DatabaseFeedFollowToFeedFollow(feedfollow))
}

func (cfg *ApiConfig) HandlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feedfollow, err := cfg.DB.GetFeedFollow(r.Context(), user.ID)
	if err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusInternalServerError, "Couldn't get feed follow")
		return
	}

	RespondWithJSON(w, http.StatusCreated, models.DatabaseFeedFollowsToFeedFollows(feedfollow))
}

func (cfg *ApiConfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusInternalServerError, "Couldn't parse feed follow")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}

	RespondWithJSON(w, http.StatusOK, struct{}{})
}
