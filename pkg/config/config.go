package config

import (
	"github.com/caarlos0/env/v10"
)

type Config struct {
	FEDAPIKey string `env:"FED_API_KEY,required"`
}

func New() (*Config, error) {
	config := Config{}

	if err := env.Parse(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Config) FEDAPIKEY() string {
	return c.FEDAPIKey
}
