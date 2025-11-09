package main

import (
	"bootdev_projects/chirpy_server/internal/auth"
	"bootdev_projects/chirpy_server/internal/database"
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerUpdatePassword(res http.ResponseWriter, req *http.Request) {
	type NewEmailPass struct {
		Email 	 string `json:"email"`
		Password string `json:"password"`
	}

	tokenHeader, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(res, http.StatusUnauthorized, "No token header found", err)
		return
	}

	validToken, err := cfg.databaseQ.GetToken(req.Context(), tokenHeader)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Error checking refresh token", err)
		return
	}

	userReq := NewEmailPass{}
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&userReq)
	if err != nil {
		respondWithError(res, http.StatusUnauthorized, "Error decoding user", err)
		return
	}

	hashed, err := auth.HashPassword(userReq.Password)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Error hashing password", err)
		return
	}
	params := database.UpdateEmailPasswordParams{
		Email: 					userReq.Email,
		HashedPassword: hashed,
		ID:							validToken.UserID,
	}
	updatedUser, err := cfg.databaseQ.UpdateEmailPassword(req.Context(), params)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Error updating email and password", err)
		return
	}

	userRes := &User{
		ID:	 				updatedUser.ID,
		CreatedAt: 	updatedUser.CreatedAt,
		UpdatedAt: 	updatedUser.UpdatedAt,
		Email: 			updatedUser.Email,
	}
	respondWithJSON(res, http.StatusOK, userRes)
}