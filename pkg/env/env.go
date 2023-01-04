package env

import (
	"os"
)

func Get(key string) string {
	if lookup, ok := os.LookupEnv(key); ok {
		return lookup
	}
	return ""
}

func Set(key, value string) error {
	return os.Setenv(key, value)
}
