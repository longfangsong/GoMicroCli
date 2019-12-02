package tools

import (
	"log"
	"os"
)

func OrElse(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func OrPanic(value string, panic string) string {
	if value == "" {
		log.Fatal(panic)
	}
	return value
}

func EnvOrElse(envName string, defaultValue string) string {
	return OrElse(os.Getenv(envName), defaultValue)
}

func EnvOrPanic(envName string, panic string) string {
	return OrPanic(os.Getenv(envName), panic)
}
