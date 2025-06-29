package config

import (
	"errors"
	"os"
)

type Config struct {
	Key string
}

func LoadConfig() (*Config, error) {
	key := os.Getenv("KEY")
	if key == "" {
		return nil, errors.New("ключ KEY не найден в .env")
	}
	return &Config{Key: key}, nil
}
