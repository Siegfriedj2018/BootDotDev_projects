package main

import (
	"bootdev_projects/chirpy_server/internal/auth"
	"bootdev_projects/chirpy_server/internal/database"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerLogin(res http.ResponseWriter, req *http.Request) {
	type userData struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	userReq := userData{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&userReq)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, "Error decoding user", err)
		return
	}

	user, err := cfg.databaseQ.GetUserbyEmail(req.Context(), userReq.Email)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Error getting user", err)
		return
	}

	valid, err := auth.CheckPasswordHash(userReq.Password, user.HashedPassword)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Error checking password", err)
		return
	}

	if !valid {
		respondWithError(res, http.StatusUnauthorized, "Invalid Password", err)
		return
	}

	jwtString, err := auth.MakeJWT(user.ID, cfg.jwtSecret)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Error making jwt string", err)
		return
	}

	refreshToken, _ := auth.MakeRefreshToken()

	// creating params to pass to the create token
	// the expiresAt is supposed to be 60 days
	params := database.CreateTokenParams{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 60),
		RevokedAt: sql.NullTime{},
	}

	reftokn, err := cfg.databaseQ.CreateToken(req.Context(), params)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Error creating refresh token", err)
		return
	}

	userRes := &User{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Email:        user.Email,
		Token:        jwtString,
		RefreshToken: reftokn.Token,
	}

	respondWithJSON(res, http.StatusOK, userRes)
}
