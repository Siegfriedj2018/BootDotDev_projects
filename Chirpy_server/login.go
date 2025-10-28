package main

import (
	"bootdev_projects/chirpy_server/internal/auth"
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerLogin(res http.ResponseWriter, req *http.Request) {
	type userData struct {
			Password 	string `json:"password"`
			Email 		string `json:"email"`
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

	userRes := &User{
		ID: 				user.ID,
		CreatedAt: 	user.CreatedAt,
		UpdatedAt: 	user.UpdatedAt,
		Email: 			user.Email,
	}

	respondWithJSON(res, http.StatusCreated, userRes)
}