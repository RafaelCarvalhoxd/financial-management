package config

import (
	"os"
)

type Config struct {
	Port             string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	JWTSecret        string
}

func NewConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresPort == "" {
		postgresPort = "5432"
	}
	postgresUser := os.Getenv("POSTGRES_USER")
	if postgresUser == "" {
		postgresUser = "postgres"
	}
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		postgresPassword = "postgres"
	}
	postgresDB := os.Getenv("POSTGRES_DB")
	if postgresDB == "" {
		postgresDB = "postgres"
	}
	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		postgresHost = "localhost"
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key-change-in-production"
	}

	return &Config{
		Port:             port,
		PostgresPort:     postgresPort,
		PostgresUser:     postgresUser,
		PostgresPassword: postgresPassword,
		PostgresDB:       postgresDB,
		PostgresHost:     postgresHost,
		JWTSecret:        jwtSecret,
	}
}