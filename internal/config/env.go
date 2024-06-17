package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	APP_PORT string
}

type DatabaseConfig struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_SSLMODE  string
}

var AppCfg AppConfig
var DBCfg DatabaseConfig

func GetEnvValue(key, Defaultvalue string) string {
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
		APP_PORT: GetEnvValue("APP_PORT", "8080"),
	}

	DBCfg = DatabaseConfig{
		DB_HOST:     GetEnvValue("DB_HOST", "localhost"),
		DB_USER:     GetEnvValue("DB_USER", "postgres"),
		DB_PASSWORD: GetEnvValue("DB_PASSWORD", "postgres"),
		DB_NAME:     GetEnvValue("DB_NAME", "postgres"),
		DB_PORT:     GetEnvValue("DB_PORT", "5432"),
		DB_SSLMODE:  GetEnvValue("DB_SSLMODE", "postgres"),
	}

	fmt.Println("Loading environment variables")
	fmt.Println(AppCfg)
	fmt.Println(DBCfg)
	return nil
}
