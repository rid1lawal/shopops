package config

import "os"

type Config struct {
	Environment  string
	DatabaseURL  string
	HTTPPort     string
	OTLPEndpoint string
	ServiceName  string
	CatalogURL   string
}

func Load() Config {
	return Config{
		Environment:  getEnv("ENVIRONMENT", "development"),
		DatabaseURL:  getEnv("DATABASE_URL", ""),
		HTTPPort:     getEnv("HTTP_PORT", "8081"),
		ServiceName:  getEnv("SERVICE_NAME", "order"),
		OTLPEndpoint: getEnv("OTLP_ENDPOINT", "localhost:4317"),
		CatalogURL:   getEnv("CATALOG_SERVICE_URL", "http://localhost:8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
