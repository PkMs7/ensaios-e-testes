package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env not found, using system environment variables")
	}
	fmt.Println("Env Loaded")
	return nil
}
