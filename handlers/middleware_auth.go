package handlers

import (
	"fmt"
	"net/http"

	"github.com/niltongc/rssagg/internal/auth"
	"github.com/niltongc/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GeteAPIKey(r.Header)
		if err != nil {
			RespondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			RespondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		handler(w, r, user)
	}
}
