package config

import (
	"os"
)

type Config struct {
	Port    string
	ApiURL  string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://backend:5007"
	}

	return &Config{
		Port:   port,
		ApiURL: apiURL,
	}
}
