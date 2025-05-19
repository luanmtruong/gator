package config

import (
	"os"
	"encoding/json"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	return write(*c)
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
    if err != nil {
		return err
    }

	file, err := os.Create(fullPath)
    if err != nil {
		return err
    }
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
    if err != nil {
		return err
    }

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homeDir, ".gatorconfig.json")
	return fullPath, nil
}
