package config

import (
	"errors"
	"os"
)

type RESASConfig struct {
	Endpoint string
	Key      string
}

func NewRESASConfig() (*RESASConfig, error) {
	apiEndpoint := os.Getenv("RESAS_API_ENDPOINT")
	if apiEndpoint == "" {
		return nil, errors.New("failed in retrieving RESAS_API_ENDPOINT. please set RESAS_API_ENDPOINT in environment variable")
	}
	apiKey := os.Getenv("RESAS_API_KEY")
	if apiKey == "" {
		return nil, errors.New("failed in retrieving API_KEY. please set RESAS_API_KEY in environment variable")
	}
	cfg := RESASConfig{
		Endpoint: apiEndpoint,
		Key:      apiKey,
	}
	return &cfg, nil
}
