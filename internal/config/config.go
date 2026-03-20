package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Address      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file not found, using system environment.")
	} else {
		log.Print(".env file found.")
	}

	config := Config{
		Address:      getEnv("ADDRESS", ":8080"),
		ReadTimeout:  getEnvDuration("READ_TIMEOUT", 15*time.Second),
		WriteTimeout: getEnvDuration("WRITE_TIMEOUT", 15*time.Second),
		IdleTimeout:  getEnvDuration("IDLE_TIMEOUT", 60*time.Second),
	}

	return config
}

func getEnv(key, fallback string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("Environment variable %s not found, using fallback.", key)
	}
	if !ok || v == "" {
		log.Printf("Environment variable %s is empty, using fallback.", key)
		return fallback
	}

	return v
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	v := getEnv(key, fallback.String())

	d, err := time.ParseDuration(v)
	if err != nil {
		log.Printf("Environment variable %s is an invalid duration, using fallback.", key)
		return fallback
	}

	return d
}
