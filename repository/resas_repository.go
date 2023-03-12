package repository

import (
	"net/http"

	"github.com/timtoronto634/resas-cli/config"
)

const apiPrefecturesPath = "/api/v1/prefectures"

type RESARepository struct {
	client      *http.Client
	apiEndpoint string
	apiKey      string
}

func NewRESASRepository() (*RESARepository, error) {
	client := http.Client{}
	cfg, err := config.NewRESASConfig()
	if err != nil {
		return nil, err
	}
	repo := RESARepository{
		client:      &client,
		apiEndpoint: cfg.Endpoint,
		apiKey:      cfg.Key,
	}
	return &repo, nil
}
