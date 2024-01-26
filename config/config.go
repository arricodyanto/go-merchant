package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
}

type Config struct {
	ApiConfig
}

func (c *Config) ConfigConfiguration() error {
	// Cek apakah file env ada
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("missing env file %v", err.Error())
	}

	// simpan api port env ke variable
	c.ApiConfig = ApiConfig{ApiPort: os.Getenv("API_PORT")}

	return nil
}

// buat bridge
func NewConfig() (*Config, error) {
	config := &Config{}
	if err := config.ConfigConfiguration(); err != nil {
		return nil, err
	}
	return config, nil
}
