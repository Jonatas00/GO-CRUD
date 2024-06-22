package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	SslMode  string
}

var AppCfg AppConfig
var DBCfg DatabaseConfig

func getEnvValue(key, Defaultvalue string) string {
	value := os.Getenv(key)
	if value == "" {
		return Defaultvalue
	}
	return value
}

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	AppCfg = AppConfig{
		Port: getEnvValue("APP_PORT", "8080"),
	}

	DBCfg = DatabaseConfig{
		Host:     getEnvValue("DB_HOST", "localhost"),
		User:     getEnvValue("DB_USER", "postgres"),
		Password: getEnvValue("DB_PASSWORD", "postgres"),
		Name:     getEnvValue("DB_NAME", "postgres"),
		Port:     getEnvValue("DB_PORT", "5432"),
		SslMode:  getEnvValue("DB_SSLMODE", "postgres"),
	}

	return nil
}
