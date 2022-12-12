package main

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func loadEnv() error {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func validateEnvironments() error {

	if strings.TrimSpace(os.Getenv("SERVER_PORT")) == "" {
		return errors.New("server port is required")
	}
	if strings.TrimSpace(os.Getenv("DB_PORT")) == "" {
		return errors.New("DB PORT is not set")
	}
	if strings.TrimSpace(os.Getenv("DB_USER")) == "" {
		return errors.New("DB USER is not set")
	}
	if strings.TrimSpace(os.Getenv("DB_PASSWORD")) == "" {
		return errors.New("DB PASSWORD is not set")
	}
	if strings.TrimSpace(os.Getenv("DB_NAME")) == "" {
		return errors.New("DB NAME is not set")
	}
	if strings.TrimSpace((os.Getenv("DB_SSL_MODE"))) == "" {
		return errors.New("DB_SSL_MODE is not set")
	}

	return nil
}
