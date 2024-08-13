package main

import "net/http"

func handlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status"`
	}

	respondWithJSON(w, http.StatusOK, response{
		Status: "ok",
	})
}

func handlerError(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, http.StatusInternalServerError, response{
		Error: "Internal Server Error",
	})
}
