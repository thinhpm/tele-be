package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBURI string
	Database   string
}

func LoadConfig() (*Config, error) {
	if error := godotenv.Load(); error != nil {
		log.Println("No .env file found")
	}

	config := &Config{
		MongoDBURI: getEnv("MONGODB_URI", ""),
		Database:   getEnv("MONGO_DATABASE", ""),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)

	if exists {
		return value
	}

	return defaultValue
}
