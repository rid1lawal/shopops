package config

import "os"

type Config struct {
	HTTPPort    string
	Environment string
	DatabaseURL string
}

func Load() Config {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "development"
	}

	return Config{
		HTTPPort:    port,
		Environment: environment,
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
