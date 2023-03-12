package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
		for y := 1980; y <= 2020; y++ {
			fmt.Printf("%v,%v,%d\n", "東京都", y, 1000000)
		}
	},
}
