package main

import (
	"net/http"

	"gitlab.com/kw3a/go-rss/internal/auth"
	"gitlab.com/kw3a/go-rss/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		dbUser, err := apiCfg.DB.GetUser(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "User not found")
			return
		}
		handler(w, r, dbUser)
	}
}
