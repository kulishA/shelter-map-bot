package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Api *ApiConfig
	Bot *BotConfig
}

type ApiConfig struct {
	Network string
	Port    string
	Host    string
}

type BotConfig struct {
	Key string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		&ApiConfig{
			Network: getEnv("API_NETWORK", ""),
			Host:    getEnv("API_HOST", ""),
			Port:    getEnv("API_PORT", ""),
		},
		&BotConfig{
			Key: getEnv("BOT_KEY", ""),
		},
	}, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
