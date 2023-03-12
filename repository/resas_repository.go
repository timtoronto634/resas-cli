package repository

import (
	"net/http"

	"github.com/timtoronto634/resas-cli/config"
)

const apiPrefecturesPath = "/api/v1/prefectures"

// RESASRepository is a repository for RESAS api
type RESASRepository struct {
	client      *http.Client
	apiEndpoint string
	apiKey      string
}

// NewRESASRepository returns RESASRepository after initilizing configuration
func NewRESASRepository() (*RESASRepository, error) {
	client := http.Client{}
	cfg, err := config.NewRESASConfig()
	if err != nil {
		return nil, err
	}
	repo := RESASRepository{
		client:      &client,
		apiEndpoint: cfg.Endpoint,
		apiKey:      cfg.Key,
	}
	return &repo, nil
}
