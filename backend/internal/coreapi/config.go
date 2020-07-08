package coreapi

import "github.com/caarlos0/env"

type Config struct {
	Environment    string `env:"ENVIRONMENT"`
	DatabaseURL    string `env:"DATABASE_URL"`
	OAuthSecretKey string `env:"OAUTH_SECRET_KEY"`
}

func (c *Config) Load() error {
	return env.Parse(c)
}
