package endpoint

import (
	"fmt"
	"os"
)

const (
	PORT     = "PORT"
	HOST     = "HOST"
	PROTOCOL = "PROTOCOL"
	USER     = "USER"
	DATABASE = "DATABASE"
)

// getenv Gets an enviornment variable value, returns a supplied default value
// if the result was empty
func getenv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

// makeKey Creates an enviornment variable compliant
// with the endpoint naming scheme
func makeKey(name, keyType string) string {
	return fmt.Sprintf("%s_%s", name, keyType)
}
