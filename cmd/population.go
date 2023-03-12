package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/timtoronto634/resas-cli/usecase"
)

func init() {
	rootCmd.AddCommand(populationCmd)
}

var populationCmd = &cobra.Command{
	Use:   "population",
	Short: "Print the population of specified options",
	Long: `Print the population of the specified argument: kind, prefecture, year.
	default is all, Tokyo, 1980-2020`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		pref := "13"
		yearFrom := 1980
		yearTo := 2020
		label := "総人口"

		usecase.PrintPopulation(ctx, label, pref, yearFrom, yearTo)
	},
}
