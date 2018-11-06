package helpers

import "os"

// GetEnv returns a default value if not exist
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
