package main

import (
	"bootdev_projects/chirpy_server/internal/database"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlervalidateChirp(res http.ResponseWriter, req *http.Request) {
	type chirps struct {
		Body 		string 		`json:"body"`
		UserId 	uuid.UUID `json:"user_id"`
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

	params := database.CreateChirpParams{
		Body: 		cleaned,
		UserID: 	chirp.UserId,
	}

	savedChirp, err := cfg.databaseQ.CreateChirp(req.Context(), params)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Could not get chirp", err)
		return
	}

	respondWithJSON(res, http.StatusCreated, Chirp{
		ID: 					savedChirp.ID,
		Created_at: 	savedChirp.CreatedAt,
		Updated_at: 	savedChirp.UpdatedAt,
		Body: 				cleaned,
		User_ID: 			savedChirp.UserID,
	})
} 