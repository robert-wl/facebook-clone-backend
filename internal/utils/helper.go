package utils

import (
	"os"
)

func GetDotENVVariable(key string, defaultVar string) string {
	variable := os.Getenv(key)

	if variable == "" {
		return defaultVar
	}
	return variable
}
