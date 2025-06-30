package config

import (
	"errors"
	"os"
)

type Config struct {
	HOST            string
	PORT            string
	SECRET_KEY_PATH string
	API_KEY         string
	BASE_URL        string
	DB_HOST         string
	DB_USER         string
	DB_PASSWORD     string
	DB_NAME         string
	DB_PORT         string
}

func NewConfig() (Config, error) {
	host := os.Getenv("HOST")
	if host == "" {
		return Config{}, errors.New("environment variable HOST is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		return Config{}, errors.New("environment variable PORT is required")
	}

	secretKeyPath := os.Getenv("SECRET_KEY_PATH")
	if secretKeyPath == "" {
		return Config{}, errors.New("environment variable SECRET_KEY_PATH is required")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return Config{}, errors.New("environment variable API_KEY is required")
	}

	baseURL := os.Getenv("BASE_URL")

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return Config{}, errors.New("environment variable DB_HOST is required")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return Config{}, errors.New("environment variable DB_USER is required")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return Config{}, errors.New("environment variable DB_PASSWORD is required")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return Config{}, errors.New("environment variable DB_NAME is required")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return Config{}, errors.New("environment variable DB_PORT is required")
	}

	return Config{
		HOST:            host + ":" + port,
		PORT:            port,
		SECRET_KEY_PATH: secretKeyPath,
		API_KEY:         apiKey,
		BASE_URL:        baseURL,
		DB_HOST:         dbHost,
		DB_USER:         dbUser,
		DB_PASSWORD:     dbPassword,
		DB_NAME:         dbName,
		DB_PORT:         dbPort,
	}, nil
}
