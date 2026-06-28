package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	MongoURI      string
	MongoDatabase string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	config := &Config{
		Port:          os.Getenv("PORT"),
		MongoURI:      os.Getenv("MONGO_URI"),
		MongoDatabase: os.Getenv("MONGO_DATABASE"),
	}

	return config, nil
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func getRequiredEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		panic(fmt.Sprintf("Environment variable %s is required but not set", key))
	}

	return value
}
