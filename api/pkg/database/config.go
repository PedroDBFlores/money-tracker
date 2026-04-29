package database

import "os"

type Config struct {
	URI string
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if uri := os.Getenv("DATABASE_URI"); uri != "" {
		cfg.URI = uri
	} else {
		cfg.URI = "mongodb://localhost:27017/"
	}
	return cfg, nil
}
