package main

import (
	"encoding/json"
	"net/http"
)

func handlervalidateChirp(res http.ResponseWriter, req *http.Request) {
	type params struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		Valid bool `json:"valid"`
	}


	parms := params{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&parms)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError,
											"Error decoding json", err)	
		return
	}

	if len(parms.Body) > 140 {
		respondWithError(res, http.StatusBadRequest,
											"Chirp is too long", err)
		return
	}

	respondWithJSON(res, http.StatusOK, returnVals{
		Valid: true,
	})
} 