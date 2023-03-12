// Package usecase handles business logic
package usecase

import (
	"context"
	"fmt"
	"io"
	"os"
)

var code2Pref = map[int]string{13: "東京都"}

// PrintPopulation prints population for specified kind, city, year
func PrintPopulation(ctx context.Context, pref, yearFrom, yearTo int) {
	population := 1000000
	prefecture := code2Pref[pref]

	for y := yearFrom; y <= yearTo; y++ {
		str := fmt.Sprintf("%v,%v,%d\n", prefecture, y, population)
		io.WriteString(os.Stdout, str)
	}
}
