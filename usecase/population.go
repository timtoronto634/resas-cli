// Package usecase handles business logic
package usecase

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/timtoronto634/resas-cli/repository"
)

var cityCode = "-"

// PrintPopulation prints population for specified kind, city, year
func PrintPopulation(ctx context.Context, prefCode string, yearFrom, yearTo int) {

	repo, err := repository.NewRESASRepository()
	if err != nil {
		log.Printf("failed in creating repository: %v", err)
		return
	}
	prefResp, err := repo.GetPrefectures(ctx)
	if err != nil {
		log.Printf("failed in getting prefectures: %v", err)
		return
	}
	var targetPref string
	for _, pref := range prefResp.Result {
		if strconv.Itoa(pref.PrefCode) == prefCode {
			targetPref = pref.PrefName
		}
	}

	popResp, err := repo.GetPopulation(ctx, cityCode, prefCode, yearFrom, yearTo)
	if err != nil {
		log.Printf("failed in getting population: %v", err)
		return
	}

	for _, kind := range popResp.Result.Data {
		for _, val := range kind.Data {
			io.WriteString(os.Stdout, fmt.Sprintf("%v,%v, %v\n", targetPref, val.Year, val.Value))
		}
	}
}
