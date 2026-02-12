package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	DBName   string
	Port     string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
		DBName:   getEnv("DB_NAME", "blogdb"),
		Port:     getEnv("PORT", "8080"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	if fallback == "" {
		log.Fatalf("Missing required env variable: %s", key)
	}
	return fallback
}
