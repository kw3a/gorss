package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gitlab.com/kw3a/go-rss/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing json:%v", err))
		return
	}
	dbUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create user:%v", err))
		return
	}
	respondWithJSON(w, http.StatusCreated, databaseUserToUser(dbUser))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(dbUser))
}

func (apiCfg *apiConfig) handlerGetUserPosts(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	posts, err := apiCfg.DB.GetUserPosts(r.Context(), database.GetUserPostsParams{
		UserID: dbUser.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get user posts:%v", err))
		return
	}
	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
