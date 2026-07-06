package config

import "os"

type Config struct {
	Port  string
	DBUrl string
}

func LoadConfig() *Config {
	cfg := &Config{
		Port:  getEnv("PORT", "8080"),
		DBUrl: getEnv("DATABASE_URL", "postgres://amirabaris@localhost:5432/goauth?sslmode=disable"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
