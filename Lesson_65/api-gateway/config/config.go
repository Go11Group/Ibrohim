package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT            string
	AUTH_SERVICE_PORT    string
	PRODUCT_SERVICE_PORT string
	DB_HOST              string
	DB_PORT              int
	DB_USER              string
	DB_PASSWORD          string
	DB_NAME              string
}

func Load() *Config {
	err := godotenv.Load("/home/ibrohim/go/src/gitlab.com/golangN11/e-commerce/api-gateway/.env")
	if err != nil {
		err := godotenv.Load("/home/jons/go/src/github.com/projects/e-commerce/api-gateway/.env")
		if err != nil {
			log.Fatalf("error loading .env: %v", err)
		}
	}
	cfg := Config{}

	cfg.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":8080"))
	cfg.AUTH_SERVICE_PORT = cast.ToString(coalesce("AUTH_SERVICE_PORT", ":8081"))
	cfg.PRODUCT_SERVICE_PORT = cast.ToString(coalesce("PRODUCT_SERVICE_PORT", ":8082"))

	cfg.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	cfg.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	cfg.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	cfg.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "password"))
	cfg.DB_NAME = cast.ToString(coalesce("DB_NAME", "e-commerce"))

	return &cfg
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}
