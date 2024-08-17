package main

import (
	"net/http"

	"github.com/gmskazi/blog_aggregator/auth"
	"github.com/gmskazi/blog_aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// NOTE: Checking if authenticated
		apiKey, err := auth.GetApiKeyToken(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't find APIKEY")
			return
		}

		user, err := cfg.DB.GetUserByAPIKEY(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Couldn't find user")
			return
		}

		handler(w, r, user)
	}
}
