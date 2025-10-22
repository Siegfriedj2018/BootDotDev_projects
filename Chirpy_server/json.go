package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(res http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}

	if code > 499 {
		log.Printf("Responding with 5XX error: %s\n", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(res, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(res http.ResponseWriter, code int, payload interface{}) {
	res.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling json: %s\n", err)
		res.WriteHeader(http.StatusInternalServerError)
	}

	res.WriteHeader(code)
	res.Write(dat)
}