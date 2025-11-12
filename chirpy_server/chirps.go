package main

import (
	"bootdev_projects/chirpy_server/internal/auth"
	"bootdev_projects/chirpy_server/internal/database"
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirps(res http.ResponseWriter, req *http.Request) {
	authorId := req.URL.Query().Get("author_id")
	sortChirps := req.URL.Query().Get("sort")

	var chirps []database.Chirp
	var err error
	if authorId != "" {
		userId, err := uuid.Parse(authorId)
		if err != nil {
			respondWithError(res, http.StatusBadRequest, "Could not parse author id", err)
			return
		}
		chirps, err = cfg.databaseQ.GetChirpsByUserID(req.Context(), userId)
		if err != nil {
			respondWithError(res, http.StatusInternalServerError, "Could not get chirps for user", err)
			return
		}
	} else {
		chirps, err = cfg.databaseQ.GetChirps(req.Context())
		if err != nil {
			respondWithError(res, http.StatusInternalServerError, "Something went wrong", err)
			return
		}
	}

	if sortChirps == "desc" {
		sort.Slice(chirps, func (a, b int) bool {
			return chirps[b].CreatedAt.Before(chirps[a].CreatedAt)
		})
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

func (cfg *apiConfig) handlerDeleteChirp(res http.ResponseWriter, req *http.Request) {
	tokenHeader, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(res, http.StatusUnauthorized, "No token header found", err)
		return
	}

	userId, err := auth.ValidateJWT(tokenHeader, cfg.jwtSecret)
	if err != nil {
		respondWithError(res, http.StatusUnauthorized, "Token is malformed or does not exist", err)
		return
	}

	validUser, err := cfg.databaseQ.GetUserByID(req.Context(), userId)
	if err != nil {
		respondWithError(res, http.StatusUnauthorized, "Invalid User", err)
		return
	}

	
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

	if validUser.ID != savedChirp.UserID {
		respondWithError(res, http.StatusForbidden, "You do not have acces to this chirp", err)
		return
	}

	err = cfg.databaseQ.DeleteChirpById(req.Context(), savedChirp.ID)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Could not delete chirp", err)
		return
	}

	respondWithJSON(res, http.StatusNoContent, nil)
}