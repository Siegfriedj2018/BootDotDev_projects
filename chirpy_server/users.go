package main

import (
	"bootdev_projects/chirpy_server/internal/auth"
	"bootdev_projects/chirpy_server/internal/database"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
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
		ID: 				 newUser.ID,
		CreatedAt: 	 newUser.CreatedAt,
		UpdatedAt: 	 newUser.UpdatedAt,
		Email: 			 newUser.Email,
		IsChirpyRed: newUser.IsChirpyRed,
	}

	respondWithJSON(res, http.StatusCreated, userRes)
}

func (cfg *apiConfig) handlerUpgradeUser(res http.ResponseWriter, req *http.Request) {
	type UserData struct {
		UserID string `json:"user_id"`
	}
	type UpgradeRequest struct {
		Event string 		`json:"event"`
		Data	UserData 	`json:"data"`
	}
	
	reqHeader, err := auth.GetAPIKey(req.Header)
  if err != nil {
		respondWithError(res, http.StatusUnauthorized, "No auth header found", err)
		return
	}

	if cfg.polkaKey == reqHeader {
		
		upgradeReq := UpgradeRequest{}
		decoder := json.NewDecoder(req.Body)
		err = decoder.Decode(&upgradeReq)
		if err != nil {
			log.Println("Failed to decode request")
			respondWithJSON(res, http.StatusNoContent, nil)
			return
		}

		if upgradeReq.Event == "user.upgraded" {
			currentID, err := uuid.Parse(upgradeReq.Data.UserID)
			if err != nil {
				log.Println("Invalid uuid")
				respondWithJSON(res, http.StatusNoContent, nil)
				return
			}
			
			currentUser, err := cfg.databaseQ.GetUserByID(req.Context(), currentID)
			if err != nil {
				respondWithError(res, http.StatusNotFound, "User not found", err)
				return
			}
			
			err = cfg.databaseQ.UpgradeUserByID(req.Context(), currentUser.ID)
			if err != nil {
				respondWithError(res, http.StatusNotFound, "Could not upgrade", err)
				return
			}
		}
	
		respondWithJSON(res, http.StatusNoContent, nil)
		return
	}
	respondWithError(res, http.StatusUnauthorized, "No api key or wrong key", err)
}