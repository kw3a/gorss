package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	type responseBody struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, responseBody{
		Status: "ok",
	})
}

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
