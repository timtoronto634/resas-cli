// Package usecase handles business logic
package usecase

import (
	"context"
	"fmt"
	"io"
	"log"
	"sort"
	"strconv"
	"sync"

	"github.com/timtoronto634/resas-cli/entity"
	"github.com/timtoronto634/resas-cli/repository"
)

var cityCode = "-"

type writable interface {
	io.Writer
}

// PrintPopulations prints population for specified kind, city, year
func PrintPopulations(ctx context.Context, output writable, label string, prefCodes []int, yearFrom, yearTo int) {
	sort.Ints(prefCodes)

	repo, err := repository.NewRESASRepository()
	if err != nil {
		log.Printf("failed in creating repository: %v", err)
		return
	}
	prefectures, err := repo.GetPrefectures(ctx)
	if err != nil {
		log.Printf("failed in getting prefectures: %v", err)
		return
	}

	codeToName := buildPrefMaps(prefectures)

	wg := sync.WaitGroup{}
	resultCh := make(chan string)

	for _, prefCode := range prefCodes {
		wg.Add(1)
		// NOTE: if the average number of api call exceeds 5 per second, use another goroutine to add sleep
		go func(pcode int) {
			defer wg.Done()

			popData, err := repo.GetPopulation(ctx, cityCode, strconv.Itoa(pcode))
			if err != nil {
				log.Printf("failed in getting population: %v", err)
				return
			}
			targetLabelData := takeByLabel(popData, label)
			if targetLabelData == nil {
				log.Printf("could not find label with: %v", label)
				return
			}
			populations := filterSortWithYear(targetLabelData.Data, yearFrom, yearTo)
			if len(populations) == 0 {
				log.Printf("could not find data within year range of: %v~%v", yearFrom, yearTo)
				return
			}
			dataPerPref := ""
			for _, population := range populations {
				dataPerPref += fmt.Sprintf("%v,%v,%v\n", codeToName[pcode], population.Year, population.Value)
			}
			resultCh <- dataPerPref
		}(prefCode)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for population := range resultCh {
		io.WriteString(output, population)
	}
}

func buildPrefMaps(prefMaps []*repository.Prefectures) map[int]string {
	prefCodeToName := make(map[int]string)
	for _, p := range prefMaps {
		prefCodeToName[p.PrefCode] = p.PrefName
	}
	return prefCodeToName
}

func takeByLabel(kinds []*entity.PopulationGroup, target string) *entity.PopulationGroup {
	for _, kind := range kinds {
		if kind.Label == target {
			return kind
		}
	}
	return nil
}

func filterSortWithYear(data []*entity.Population, from, to int) []*entity.Population {
	filteredData := make([]*entity.Population, 0, len(data))
	for _, popData := range data {
		if popData.Year >= from && popData.Year <= to {
			filteredData = append(filteredData, popData)
		}
	}

	sort.Slice(filteredData, func(i, j int) bool {
		return filteredData[i].Year > filteredData[j].Year
	})

	return filteredData
}
