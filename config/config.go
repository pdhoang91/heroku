package config

import (
	"os"
)

var (
	BaseAPIURL    = GetString("BaseAPIURL", "https://mfx-recruit-dev.herokuapp.com")
	SwaggerDomain = GetString("SwaggerDomain", "127.0.0.1")
)

// Config represents the application configuration.
type Config struct {
	HTTPPort      string
	SwaggerDomain string
	BaseAPIURL    string
}

func NewConfig() *Config {

	cfg := &Config{
		HTTPPort: GetString("PORT", ":80"),
	}

	return cfg
}

func GetString(key string, defaultValue string) string {
	result := os.Getenv(key)
	if len(result) > 0 {

		return result
	}
	return defaultValue
}
