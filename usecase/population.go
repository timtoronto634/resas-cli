// Package usecase handles business logic
package usecase

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/timtoronto634/resas-cli/repository"
)

// PrintPopulation prints population for specified kind, city, year
func PrintPopulation(ctx context.Context, pref, yearFrom, yearTo int) {
	cityCode := "-"

	repo, err := repository.NewRESASRepository()
	if err != nil {
		log.Printf("failed in creating repository: %v", err)
		return
	}
	popResp, err := repo.GetPopulation(cityCode, "13", yearFrom, yearTo)
	if err != nil {
		log.Printf("failed in getting population: %v", err)
		return
	}

	for _, kind := range popResp.Result.Data {
		for _, val := range kind.Data {
			io.WriteString(os.Stdout, fmt.Sprintf("東京都,%v, %v\n", val.Year, val.Value))
		}
	}
}
