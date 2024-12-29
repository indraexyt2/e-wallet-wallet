package helpers

import (
	"github.com/joho/godotenv"
	"log"
)

var EnvMap = map[string]string{}

func SetupConfig() {
	var err error
	EnvMap, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("Failed to read .env file", err)
	}
}

func GetEnv(key string, val string) string {
	result := EnvMap[key]
	if result == "" {
		result = val
	}
	return result
}
