// Package repository provides data access
package repository

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/timtoronto634/resas-cli/entity"
)

const apiPopulationPath = "/api/v1/population/composition/perYear"

// GetPopulation get population data from RESAS api
func (repo *RESASRepository) GetPopulation(ctx context.Context, cityCode, prefCode string) ([]*entity.PopulationGroup, error) {
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

	var popResp entity.PopulationResponse
	err = json.Unmarshal(body, &popResp)
	if err != nil {
		log.Printf("failed in decoding response: %v", err)
		return nil, err
	}

	return popResp.Result.Data, nil
}
