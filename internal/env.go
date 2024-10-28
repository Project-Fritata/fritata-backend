package internal

import (
	"log"
	"os"
)

func GetEnvVar(key string) string {
	env_var, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal("Environment variable not found for key: ", key)
	}
	return env_var
}
