package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if necessary
	if os.Getenv("DB_CONNECTION_STRING") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Unable to load environment files - Err: %s", err)
		}
	}
}
