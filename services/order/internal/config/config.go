package config

import "os"

type Config struct {
	Environment string
	DatabaseURL string
	HTTPPort    string
	CatalogURL  string
}

func Load() Config {
	return Config{
		Environment: getEnv("ENVIRONMENT", "development"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		HTTPPort:    getEnv("HTTP_PORT", "8081"),
		CatalogURL:  getEnv("CATALOG_SERVICE_URL", "http://localhost:8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
