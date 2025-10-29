package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"
)

func CreateJWTSecret() (int, error) {
	const jwtString = "JWT_SECRET="
	byteLength := 64

	randomBytes := make([]byte, byteLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return 0, err
	}

	keyString := base64.StdEncoding.EncodeToString(randomBytes)

	filepath := "../../.env"

	file, err := os.OpenFile(filepath, os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}

	defer file.Close()
	readBytes := []byte{}
	_, err = file.Read(readBytes)
	if err != nil {
		return 0, err	
	}

	if !bytes.Contains(readBytes, []byte(jwtString)) {
		return file.Write([]byte(jwtString + keyString))
	}

	
	return file.Write([]byte(keyString))
}