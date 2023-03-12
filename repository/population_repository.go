package repository

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

const apiPopulationPath = "/api/v1/population/composition/perYear"

type PopulationResponse struct {
	Message     string          `json:"message"`
	Result      *PopulationData `json:"result"`
	StatusCode  int             `json:"statusCode"`
	Description string          `json:"description"`
}

type PopulationData struct {
	BoundaryYear int               `json:"boundaryYear"`
	Data         []*PopulationKind `json:"data"`
}

type PopulationKind struct {
	Label string             `json:"label"`
	Data  []*PopulationValue `json:"data"`
}

type PopulationValue struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

func (repo *RESARepository) GetPopulation(ctx context.Context, cityCode, prefCode string, yearFrom, yearTo int) (*PopulationResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", repo.apiEndpoint+apiPopulationPath, nil)
	if err != nil {
		log.Printf("failed in creating request: %v", err)
		return nil, err
	}

	values := req.URL.Query()
	values.Add("cityCode", cityCode)
	values.Add("prefCode", prefCode)
	req.URL.RawQuery = values.Encode()
	log.Printf("requesting %v\n", req.URL)

	req.Header.Add("X-API-KEY", repo.apiKey)
	resp, err := repo.client.Do(req)
	if err != nil {
		log.Printf("failed in making request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error returned from api server")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed in reading response body: %v", err)
		return nil, err
	}
	if len(body) == 5 {
		// returned "500"
		return nil, errors.New("server error occurred")
	}

	var popResp PopulationResponse
	err = json.Unmarshal(body, &popResp)
	if err != nil {
		log.Printf("failed in decoding response: %v", err)
		return nil, err
	}

	return &popResp, nil
}
