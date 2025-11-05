package main

import (
	"time"

	"github.com/google/uuid"
)

type Chirps struct {
	Chirps []Chirp
}

type Chirp struct {
	ID 						uuid.UUID `json:"id"`
	Created_at		time.Time `json:"created_at"`
	Updated_at		time.Time `json:"updated_at"`
	Body 					string 		`json:"body"`
	User_ID				uuid.UUID `json:"user_id"`
}