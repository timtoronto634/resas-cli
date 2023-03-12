package repository

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type PrefectureMaps struct {
	Message string        `json:"message"`
	Result  []Prefectures `json:"result"`
}

type Prefectures struct {
	PrefCode int    `json:"prefCode"`
	PrefName string `json:"prefName"`
}

func (repo *RESARepository) GetPrefectures(ctx context.Context) (*PrefectureMaps, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", repo.apiEndpoint+apiPrefecturesPath, nil)
	if err != nil {
		log.Printf("failed in creating request: %v", err)
		return nil, err
	}
	req.Header.Add("X-API-KEY", repo.apiKey)
	resp, err := repo.client.Do(req)
	if err != nil {
		log.Printf("failed in making request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed in reading response body: %v", err)
		return nil, err
	}
	if len(body) == 5 {
		// server responded "500"
		return nil, errors.New("server error occurred")
	}

	var prefResp PrefectureMaps
	err = json.Unmarshal(body, &prefResp)
	if err != nil {
		log.Printf("failed in decoding response: %v", err)
		return nil, err
	}

	return &prefResp, nil
}
