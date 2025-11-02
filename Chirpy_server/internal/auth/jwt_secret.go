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

	// defer file.Close()
	file, err := os.ReadFile(fullPath)
	if err != nil {
		return 0, err
	}

	convertedBytes := strings.Split(string(file), "\n")
	replaceString := jwtString + "\"" + keyString + "\""

	var jwtSecret string
	var newLines []string
	isFound := false
	for _, line := range convertedBytes {
		if strings.Contains(line, jwtString) {
			jwtSecret = strings.ReplaceAll(line, line, replaceString)
			newLines = append(newLines, jwtSecret)
			isFound = true
			continue
		}
		newLines = append(newLines, line)
	}
	if !isFound {
		newLines = append(newLines, replaceString)
	}
	newJWTFile := strings.Join(newLines, "\n")

	err = os.WriteFile(fullPath, []byte(newJWTFile), 0644)
	if err != nil {
		return 0, err
	}

	return len(newJWTFile), nil
}