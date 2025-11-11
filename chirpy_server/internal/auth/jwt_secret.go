package auth

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Note to future self openfile does not understand ~
// thus you have to open file this way, see below
// TODO: Maybe ask if you want to generate a new jwt secret
// if not then return previous secret else generate new
func CreateJWTSecret() (int, error) {

	const jwtString = "JWT_SECRET="
	byteLength := 64
	is_Reset := false
	inputReader := bufio.NewReader(os.Stdin)

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

	fullPath := filepath.Join(homeDir, "my_bootdotdev_projects/BootDotDev_projects/chirpy_server/.env")

	file, err := os.ReadFile(fullPath)
	if err != nil {
		return 0, err
	}

	log.Print("Do you want to reset the JWT secret (type y for yes or n for no): ")
	userReset, err := inputReader.ReadByte()
	if err != nil {
		log.Println("Something unexpected happend")
		return 0, err
	}
	switch userReset {
	case 'y':
		is_Reset = true
	case 'n':
		is_Reset = false
	default:
		log.Println("Invalid response")
		return 0, errors.New("invalid response, Goodbye")
	}
	if is_Reset {
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
	
	return 0, nil
}