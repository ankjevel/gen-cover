package utils

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func init() {
	var (
		host = getEnv("HOST", "localhost")
		port = getEnv("PORT", "8091")
	)

	Config = config{
		Host:  host,
		Port:  port,
		Addr:  getEnv("ADDR", fmt.Sprintf("%s:%s", host, port)),
		Title: getEnv("TITLE", ""),
	}
}

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value != "" {
		return value
	}
	return fallback
}
