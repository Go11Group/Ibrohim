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
}

func Load() *Config {
	err := godotenv.Load("/home/ibrohim/go/src/gitlab.com/golangN11/e-commerce/api-gateway/.env")
	if err != nil {
		log.Fatalf("error loading .env: %v", err)
	}

	cfg := Config{}

	cfg.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":8080"))
	cfg.AUTH_SERVICE_PORT = cast.ToString(coalesce("AUTH_SERVICE_PORT", ":8081"))
	cfg.PRODUCT_SERVICE_PORT = cast.ToString(coalesce("PRODUCT_SERVICE_PORT", ":8082"))

	return &cfg
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}
