package auth

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtCustomClaim struct {
	Name  		string `json:"name"`
	jwt.RegisteredClaims
}

func MakeJWT(userID uuid.UUID, tokenSecret string) (string, error) {
	claim := jwtCustomClaim{
		"jwtChirpyClaim",
		jwt.RegisteredClaims{
			Issuer: 		"chirpy",
			IssuedAt: 	jwt.NewNumericDate(time.Now()),
			ExpiresAt: 	jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			Subject: 		userID.String(),
			Audience: 	jwt.ClaimStrings{"chirpy-user-api"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(tokenSecret))
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claims := jwtCustomClaim{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	uuidString, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, err
	}
	
	return uuid.Parse(uuidString)
}

func MakeRefreshToken() (string, error) {
	key := make([]byte, 32)
	rand.Read(key)
	encodedToken := hex.EncodeToString(key)
	return encodedToken, nil
}