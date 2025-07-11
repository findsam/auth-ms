package config

import (
	"os"

	"github.com/findsam/auth-micro/pkg/util"
	_ "github.com/joho/godotenv/autoload"
)

var Envs = config()

func config() *util.Config {
	return &util.Config{
		DB_USER:                getEnv("DB_USER", ""),
		DB_PWD:                 getEnv("DB_PWD", ""),
		DB_NAME:                getEnv("DB_NAME", ""),
		JWT_SECRET:             getEnv("JWT_SECRET", ""),
		MONGO_URI:              getEnv("MONGO_URI", ""),
		STRIPE_PWD:             getEnv("STRIPE_PWD", ""),
		STRIPE_PUBLISHABLE_KEY: getEnv("STRIPE_PUBLISHABLE_KEY", ""),
		PORT: 				 	getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
