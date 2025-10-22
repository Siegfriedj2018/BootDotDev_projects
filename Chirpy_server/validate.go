package main

import (
	"encoding/json"
	"net/http"
)

func handlervalidateChirp(res http.ResponseWriter, req *http.Request) {
	type chirps struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		Cleaned_body string `json:"cleaned_body"`
	}

	profaneWords := map[string]bool{
		"kerfuffle": 	false,
		"sharbert": 	false,
		"fornax": 		false,
	}


	chirp := chirps{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&chirp)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError,
											"Error decoding json", err)	
		return
	}

	if len(chirp.Body) > 140 {
		respondWithError(res, http.StatusBadRequest,
											"Chirp is too long", err)
		return
	}

	cleaned, err := findReplaceBadWords(profaneWords, chirp.Body)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, "Error replacing word", err)
		return
	}

	respondWithJSON(res, http.StatusOK, returnVals{
		Cleaned_body: cleaned,
	})
} 