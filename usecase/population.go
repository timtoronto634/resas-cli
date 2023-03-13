// Package usecase handles business logic
package usecase

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/timtoronto634/resas-cli/entity"
	"github.com/timtoronto634/resas-cli/repository"
)

var cityCode = "-"

type writable interface {
	io.Writer
}

// PrintPopulation prints population for specified kind, city, year
func PrintPopulation(ctx context.Context, output writable, label string, prefCode []string, yearFrom, yearTo int) {

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
	for _, pref := range prefResp {
		if strconv.Itoa(pref.PrefCode) == prefCode[0] {
			targetPref = pref.PrefName
		}
	}

	popData, err := repo.GetPopulation(ctx, cityCode, prefCode[0])
	if err != nil {
		log.Printf("failed in getting population: %v", err)
		return
	}
	targetLabelData := takeWithLabel(popData, label)
	if targetLabelData == nil {
		log.Printf("could not find label with: %v", label)
		return
	}
	populations := filterWithYear(targetLabelData.Data, yearFrom, yearTo)
	if len(populations) == 0 {
		log.Printf("could not find data within year range of: %v~%v", yearFrom, yearTo)
		return
	}

	for _, p := range populations {
		io.WriteString(output, fmt.Sprintf("%v,%v,%v\n", targetPref, p.Year, p.Value))
	}
}

func takeWithLabel(kinds []*entity.PopulationGroup, target string) *entity.PopulationGroup {
	for _, kind := range kinds {
		if kind.Label == target {
			return kind
		}
	}
	return nil
}

func filterWithYear(data []*entity.Population, from, to int) []*entity.Population {
	popDatas := make([]*entity.Population, 0, len(data))
	for _, popData := range data {
		if popData.Year >= from && popData.Year <= to {
			popDatas = append(popDatas, popData)
		}
	}
	return popDatas
}
