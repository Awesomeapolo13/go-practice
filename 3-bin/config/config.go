package config

import "os"

type Config struct {
	key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	if key == "" {
		panic("KEY environment variable is not set")
	}

	return &Config{
		key: key,
	}
}

func (c *Config) GetKey() string {
	return c.key
}
