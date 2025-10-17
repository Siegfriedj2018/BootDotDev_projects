package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"
// This has to be changed on other os
const configFilePath = "/mnt/c/Users/siegf/SynologyDrive/Drive/Go_Tutorial/bootdotdev/blog_aggregator"

type Config struct {
	DbUrl 				string  `json:"db_url"`
	CurrentUser		string  `json:"current_user_name"`
}

func Read() *Config {
	path, err := getConfigFilePath()
	if err != nil {
		log.Fatalf("there was an invalid filepath: %v", err)
	}
	
	
	data, err := os.Open(path[0])
	if err != nil {
		log.Fatalf("there was an error opening the file: %v", err)
	}
	defer data.Close()
	
	configData := &Config{}
	decoder := json.NewDecoder(data)
	err = decoder.Decode(&configData)
	if err != nil {
		log.Fatalf("there was an error unmarshalling the data: %v", err)
	}
	
	return configData
}

func getConfigFilePath() ([]string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return []string{}, fmt.Errorf("there was an err getting home dir: %w", err)
	}
	fullPath := filepath.Join(configFilePath, configFileName)
	testPath := filepath.Join(userHome, configFileName)
	allpath := []string{fullPath, testPath}
	return allpath, nil
}

func write(conf Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("there was an error getting config path: %w", err)
	}

	dataRun, err := os.Create(path[0])
	if err != nil {
		return fmt.Errorf("err creating file: %w",err)
	}
	dataTest, err := os.Create(path[1])
	if err != nil {
		return fmt.Errorf("err creating file: %w",err)
	}
	defer dataRun.Close()
	defer dataTest.Close()

	userData := json.NewEncoder(dataRun)
	userData.SetIndent("", "  ")
	err = userData.Encode(conf)
	if err != nil {
		return fmt.Errorf("there was an error encoding the data: %w", err)
	}

	serData := json.NewEncoder(dataTest)
	serData.SetIndent("", "  ")
	err = serData.Encode(conf)
	if err != nil {
		return fmt.Errorf("there was an error encoding the data: %w", err)
	}

	return nil
}

func (confData *Config) SetUser(username string) error {
	confData.CurrentUser = username
	return write(*confData)
}
