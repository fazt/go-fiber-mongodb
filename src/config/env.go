package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MONGODB_URI string
	PORT        string
}

func GetConfig() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{
		MONGODB_URI: os.Getenv("MONGODB_URI"),
		PORT:        ":4000",
	}

	return config
}
