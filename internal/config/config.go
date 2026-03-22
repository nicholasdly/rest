package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Address      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	ApiKey       string
	DatabaseUrl  string
}

func Load() (Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Warn(".env file not found, using system environment.")
	} else {
		slog.Warn(".env file found.")
	}

	envApiKey, err := getRequiredEnv("API_KEY")
	if err != nil {
		return Config{}, err
	}

	envDatabaseUrl, err := getRequiredEnv("DATABASE_URL")
	if err != nil {
		return Config{}, err
	}

	config := Config{
		Address:      getEnv("ADDRESS", ":8080"),
		ReadTimeout:  getEnvDuration("READ_TIMEOUT", 15*time.Second),
		WriteTimeout: getEnvDuration("WRITE_TIMEOUT", 15*time.Second),
		IdleTimeout:  getEnvDuration("IDLE_TIMEOUT", 60*time.Second),
		ApiKey:       envApiKey,
		DatabaseUrl:  envDatabaseUrl,
	}

	return config, nil
}

func getRequiredEnv(key string) (string, error) {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return "", fmt.Errorf("Environment variable %s is required.", key)
	}

	return v, nil
}

func getEnv(key, fallback string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		slog.Warn("Environment variable not found, using fallback.", "key", key)
		return fallback
	}
	if v == "" {
		slog.Warn("Environment variable is empty, using fallback.", "key", key)
		return fallback
	}

	return v
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	v := getEnv(key, fallback.String())

	d, err := time.ParseDuration(v)
	if err != nil {
		slog.Warn("Environment variable is an invalid duration, using fallback.", "key", key)
		return fallback
	}

	return d
}
