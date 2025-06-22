package config

import "os"

type Config struct {
	accessKey string
	masterKey string
}

func NewConfig() *Config {
	accessKey := os.Getenv("ACCESS_KEY")
	if accessKey == "" {
		panic("ACCESS_KEY environment variable is not set")
	}

	masterKey := os.Getenv("MASTER_KEY")
	if masterKey == "" {
		panic("MASTER_KEY environment variable is not set")
	}

	return &Config{
		accessKey: accessKey,
		masterKey: masterKey,
	}
}

func (c *Config) GetAccessKey() string {
	return c.accessKey
}

func (c *Config) GetMasterKey() string {
	return c.masterKey
}
