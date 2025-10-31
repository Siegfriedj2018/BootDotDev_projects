package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"iter"
	"log"
	"os"
	"path/filepath"
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

	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return 0, fmt.Errorf("file failed to open: %w", err)
	}

	defer file.Close()
	readBytes := []byte{}
	_, err = file.Read(readBytes)
	if err != nil {
		return 0, err
	}

	var val1, val2 []byte
	var ok1 bool
	var ok2 bool
	line := bytes.Lines(readBytes)
	nextVal, stop := iter.Pull(line)
	value, _ := nextVal()
	log.Printf("Lines: %v\n", value)
	defer stop()
	for {
		val1, ok1 = nextVal()
		if !ok1 {
			stop()
			log.Printf("don't know what i am doing here: %v", val1)
			break
		}
		val2, ok2 = nextVal()
		if !ok2 {
			stop()
			log.Printf("Dont know where to go from here: %v", val2)
			break
		}

		log.Printf("Val1: %v, Ok: %v", val1, ok1)
		log.Printf("Val2: %v, Ok: %v", val2, ok2)
	}
	log.Printf("Val1: %v, Ok: %v", val1, ok1)
	log.Printf("Val2: %v, Ok: %v", val2, ok2)

	jwtSecret := bytes.ReplaceAll(readBytes, val2, []byte("\n" + jwtString+"\""+keyString+"\""))
	return file.Write(jwtSecret)
}