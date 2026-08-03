package config

import "os"

type Config struct {
	HTTPPort     string
	Environment  string
	DatabaseURL  string
	OTLPEndpoint string
	ServiceName  string
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

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		serviceName = "catalog"
	}

	return Config{
		HTTPPort:     port,
		Environment:  environment,
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		OTLPEndpoint: os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"),
		ServiceName:  serviceName,
	}
}
