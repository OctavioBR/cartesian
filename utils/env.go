package utils

import "os"

// Getenv is very simmilar to standard os.Getenv, but accepts
// a fallback default value if environment variable is not defined
func Getenv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return fallback
}
