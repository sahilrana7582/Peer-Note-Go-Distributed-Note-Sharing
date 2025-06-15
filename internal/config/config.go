package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL      string
	ServerPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env variables.")
	}

	return &Config{
		DBURL:      os.Getenv("DB_URL"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
