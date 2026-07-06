package config

import "os"

type Config struct {
	Port string
}

func LoadConfig() *Config {
	cfg := &Config{
		Port: getEnv("PORT", "8080"),
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
