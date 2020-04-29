package env

import (
	"os"
	"strconv"
)

// Get string from environment variables
func Get(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetBool from environment variables
func GetBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.ParseBool(value); err == nil {
			return result
		}
	}
	return fallback
}
