package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Note to future self openfile does not understand ~
// thus you have to open this way, see below
func CreateJWTSecret() (int, error) {
	const jwtString = "JWT_SECRET="
	byteLength := 64

	randomBytes := make([]byte, byteLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return 0, err
	}

	keyString := base64.StdEncoding.EncodeToString(randomBytes)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return 0, fmt.Errorf("could not get user home directory: %w", err)
	}

	fullPath := filepath.Join(homeDir, "my_bootdotdev_projects/BootDotDev_projects/Chirpy_server/.env")

	// file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	// if err != nil {
	// 	return 0, fmt.Errorf("file failed to open: %w", err)
	// }

	// defer file.Close()
	file, err := os.ReadFile(fullPath)
	if err != nil {
		return 0, err
	}

	convertedBytes := string(file)
	replaceString := jwtString + "\"" + keyString + "\""

	// TODO: Fix the bug where it find jwtstring and just appends the key on previous key instead of replacing it
	jwtSecret := strings.ReplaceAll(convertedBytes, jwtString, replaceString)
	err = os.WriteFile(fullPath, []byte(jwtSecret), 0644)
	if err != nil {
		return 0, err
	}

	return len(convertedBytes), nil
}