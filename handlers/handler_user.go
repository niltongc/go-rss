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

type ApiConfig struct {
	DB *database.Queries
}

func (cfg *ApiConfig) HandlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	RespondWithJSON(w, http.StatusOK, models.DatabaseUserToUser(user))
}

func (cfg *ApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	RespondWithJSON(w, 200, models.DatabaseUserToUser(user))
}
