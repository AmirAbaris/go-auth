package config

import "os"

type Config struct {
	Port      string
	DBUrl     string
	JWTSecret string
}

func LoadConfig() *Config {
	cfg := &Config{
		Port:      getEnv("PORT", "8080"),
		DBUrl:     getEnv("DATABASE_URL", "postgres://amirabaris@localhost:5432/goauth?sslmode=disable"),
		JWTSecret: getEnv("JWT_SECRET", "CHANGE-ME"),
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
