package config

import "os"

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	if key == "" {
		panic("KEY environment variable is not set")
	}

	return &Config{
		Key: key,
	}
}
