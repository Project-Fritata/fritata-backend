package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnvVar(key string) string {
	env_var, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal("Environment variable not found for key: ", key)
	}
	return env_var
}
