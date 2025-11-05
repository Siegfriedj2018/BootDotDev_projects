package main

import (
	"bootdev_projects/chirpy_server/internal/auth"
	"bootdev_projects/chirpy_server/internal/database"
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerCreateUser(res http.ResponseWriter, req *http.Request) {
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
	
	hashed, err := auth.HashPassword(userReq.Password)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Error hashing password", err)
		return
	}

	params := database.CreateUserParams{
		Email: userReq.Email,
		HashedPassword: hashed,
	}
	
	newUser, err := cfg.databaseQ.CreateUser(req.Context(), params)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Error with creating user", err)
		return
	}

	userRes := &User{
		ID: 				newUser.ID,
		CreatedAt: 	newUser.CreatedAt,
		UpdatedAt: 	newUser.UpdatedAt,
		Email: 			newUser.Email,
	}

	respondWithJSON(res, http.StatusCreated, userRes)
}