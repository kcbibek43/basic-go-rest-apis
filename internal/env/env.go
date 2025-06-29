package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetString(key, fallback string) string {
	LoadEnv()

	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}
	return val
}

func getInt(key string, fallback int) int {
	LoadEnv()

	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)

	if err != nil {
		return fallback
	}

	return valAsInt
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
