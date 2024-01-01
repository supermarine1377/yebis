package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	FEDAPIKey string `env:"FED_API_KEY,required"`
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to godotenv.Load: %w", err)
	}

	config := Config{}

	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}
	return &config, nil
}

func (c *Config) FEDAPIKEY() string {
	return c.FEDAPIKey
}
