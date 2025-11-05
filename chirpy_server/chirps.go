package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirps(res http.ResponseWriter, req *http.Request) {
	chirps, err := cfg.databaseQ.GetChirps(req.Context())
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

	newChirps := []Chirp{}
	for _, sqlChirp := range chirps {
		newChirp := Chirp{
			ID: 				sqlChirp.ID,
			Created_at: sqlChirp.CreatedAt,
			Updated_at: sqlChirp.UpdatedAt,
			Body: 			sqlChirp.Body,
			User_ID: 		sqlChirp.UserID,
		}
		newChirps = append(newChirps, newChirp)
	}
	
	respondWithJSON(res, http.StatusOK, newChirps)
}

func (cfg *apiConfig) handlerGetChirp(res http.ResponseWriter, req *http.Request) {
	chirpIDString := req.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid uuid", err)
		return
	}

	savedChirp, err := cfg.databaseQ.GetChirpByID(req.Context(), chirpID)
	if err != nil {
		respondWithError(res, http.StatusNotFound, "Could not get chirp", err)
		return
	}

	newChirp := Chirp{
		ID: 				savedChirp.ID,
		Created_at: savedChirp.CreatedAt,
		Updated_at: savedChirp.UpdatedAt,
		Body: 			savedChirp.Body,
		User_ID: 		savedChirp.UserID,
	}

	respondWithJSON(res, http.StatusOK, newChirp)
}