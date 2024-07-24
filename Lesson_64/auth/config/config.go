package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	AUTH_SERVICE_PORT string
	AUTH_ROUTER_PORT  string
	DB_HOST           string
	DB_PORT           int
	DB_USER           string
	DB_NAME           string
	DB_PASSWORD       string
	ACCESS_TOKEN      string
	REFRESH_TOKEN     string
}

func coalesce(env string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(env)
	if !exists {
		return defaultValue
	}
	return value
}

func Load() *Config {
	if err := godotenv.Load("/home/jons/go/src/github.com/projects/e-commerce/E-commerce/.env"); err != nil {
		if err = godotenv.Load("/home/ibrohim/go/src/gitlab.com/golangN11/e-commerce/E-commerce/.env"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	cfg := Config{}

	cfg.AUTH_ROUTER_PORT = cast.ToString(coalesce("AUTH_ROUTER_PORT", ":8081"))
	cfg.AUTH_SERVICE_PORT = cast.ToString(coalesce("AUTH_SERVICE_PORT", ":50051"))
	cfg.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	cfg.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	cfg.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	cfg.DB_NAME = cast.ToString(coalesce("DB_NAME", "postgres"))
	cfg.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "password"))
	cfg.ACCESS_TOKEN = cast.ToString(coalesce("ACCESS_TOKEN", "my_secret_key"))
	cfg.REFRESH_TOKEN = cast.ToString(coalesce("REFRESH_TOKEN", "my_secret_key"))

	return &cfg
}
