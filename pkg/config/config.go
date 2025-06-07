package config

import (
	"os"

	t "github.com/findsam/auth-micro/pkg/util"
	_ "github.com/joho/godotenv/autoload"
)

var Envs = config()

func config() *t.Config{
	return &t.Config{
		DB_USER: getEnv("DB_USER", ""),
		DB_PWD: getEnv("DB_PWD", ""),
		DB_NAME: getEnv("DB_NAME", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}