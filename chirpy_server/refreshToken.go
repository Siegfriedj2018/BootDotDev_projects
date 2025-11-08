package main

import (
	"bootdev_projects/chirpy_server/internal/auth"
	"fmt"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerRefresh(res http.ResponseWriter, req *http.Request) {
	type tokenResponse struct {
		Token	string `json:"token"` 	
	}

	token, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, "No token header found", err)
		return
	}

	foundToken, err := cfg.databaseQ.GetToken(req.Context(), token)
	if err != nil {
		respondWithError(res, http.StatusUnauthorized, "Could not find token", err)
	}

	if foundToken.ExpiresAt.Compare(time.Now()) == -1 || foundToken.RevokedAt.Valid {
		respondWithError(res, http.StatusUnauthorized, "Token is expired or already revoked", fmt.Errorf("token exipred or revoked"))
		return
	}

	newToken, err := auth.MakeJWT(foundToken.UserID, cfg.jwtSecret)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Error making new jwt string", err)
		return
	}

	respondWithJSON(res, http.StatusOK, tokenResponse{
		Token: newToken,
	})
}

func (cfg *apiConfig) handlerRevoke(res http.ResponseWriter, req *http.Request) {
	token, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "No token header found", err)
		return
	}

	foundToken, err := cfg.databaseQ.GetToken(req.Context(), token)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Could not find token", err)
		return
	}

	if foundToken.RevokedAt.Valid {
		respondWithJSON(res, http.StatusNoContent, nil)
		return
	}

	err = cfg.databaseQ.RevokeToken(req.Context(), foundToken.Token)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

	respondWithJSON(res, http.StatusNoContent, nil)
}