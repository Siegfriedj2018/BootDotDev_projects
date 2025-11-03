package main

import (
	"bootdev_projects/chirpy_server/internal/auth"
	"encoding/json"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerLogin(res http.ResponseWriter, req *http.Request) {
	type userData struct {
		Password 		string `json:"password"`
		Email 			string `json:"email"`
		Expiration	int		 `json:"expires_in_seconds"`
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

	if userReq.Expiration == 0 || userReq.Expiration > 60 {
		userReq.Expiration = 60
	}
	
	jwtString, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Duration(userReq.Expiration) * time.Second)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Error making jwt string", err)
		return
	}
	
	userRes := &User{
		ID: 				user.ID,
		CreatedAt: 	user.CreatedAt,
		UpdatedAt: 	user.UpdatedAt,
		Email: 			user.Email,
		Token: 			jwtString,
	}

	respondWithJSON(res, http.StatusOK, userRes)
}