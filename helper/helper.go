package helper

import (
	"github.com/joho/godotenv"
	"os"
)

func GetDotENVVariable(key string, defaultVar string) string {
	if err := godotenv.Load(); err != nil {
		return defaultVar
	}

	variable := os.Getenv(key)

	if variable == "" {
		return defaultVar
	}
	return variable
}
