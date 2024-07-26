package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PRODUCT_SERVICE_PORT string
	MongoDB_NAME         string
	MongoURI             string
	RedisAddress         string
	RedisPassword        string
	RedisDB              int
}

func coalesce(env string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(env)
	if exists {
		return value
	}
	return defaultValue
}

func Load() *Config {
	if err := godotenv.Load("/home/jons/go/src/github.com/projects/e-commerce/auth_service/.env"); err != nil {
		if err = godotenv.Load("/home/ibrohim/go/src/gitlab.com/golangN11/e-commerce/products/.env"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	cfg := Config{}

	cfg.PRODUCT_SERVICE_PORT = cast.ToString(coalesce("PRODUCT_SERVICE_PORT", ":50052"))

	cfg.MongoDB_NAME = cast.ToString(coalesce("MongoDB_NAME", "test"))
	cfg.MongoURI = cast.ToString(coalesce("MongoURI", "mongodb://localhost:27017"))

	cfg.RedisAddress = cast.ToString(coalesce("RedisAddress", "localhost:6379"))
	cfg.RedisPassword = cast.ToString(coalesce("RedisPassword", ""))
	cfg.RedisDB = cast.ToInt(coalesce("RedisDB", 0))

	return &cfg
}
