package config

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

func Config(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
