package config

import "os"

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

func NewConfig() Config {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.fireblocks.io"
	}

	return Config{
		HOST:            host + ":" + port,
		PORT:            port,
		SECRET_KEY_PATH: os.Getenv("SECRET_KEY_PATH"),
		API_KEY:         os.Getenv("API_KEY"),
		BASE_URL:        baseURL,
		DB_HOST:         os.Getenv("DB_HOST"),
		DB_USER:         os.Getenv("DB_USER"),
		DB_PASSWORD:     os.Getenv("DB_PASSWORD"),
		DB_NAME:         os.Getenv("DB_NAME"),
		DB_PORT:         os.Getenv("DB_PORT"),
	}
}
