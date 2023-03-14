package cmd

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/timtoronto634/resas-cli/usecase"
)

var codeToLabel = map[string]string{
	"all":        "総人口",
	"youth":      "年少人口",
	"productive": "生産年齢人口",
	"elderly":    "老年人口",
}

var populationCmd = &cobra.Command{
	Use:   "population",
	Short: "Print the population of specified options",
	Long: `Print the population of the specified argument: kind, prefecture, year.
	default is all, Tokyo, 1980-2020`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		prefCodes, err := cmd.Flags().GetIntSlice("prefectures")
		if err != nil {
			log.Printf("[error] failed to get prefecture from argument: %v", err)
			return
		}
		labelFlag, err := cmd.Flags().GetString("label")
		if err != nil {
			log.Printf("[error] failed to get label from argument: %v", err)
			return
		}
		yearFrom, err := cmd.Flags().GetInt("from")
		if err != nil {
			log.Printf("[error] failed to get yearFrom from argument: %v", err)
			return
		}
		yearTo, err := cmd.Flags().GetInt("to")
		if err != nil {
			log.Printf("[error] failed to get yearTo from argument: %v", err)
			return
		}

		validLabels := []string{"all", "youth", "productive", "elderly"}
		if !isValidLabel(labelFlag, validLabels) {
			log.Printf("[error] Invalid label provided. Allowed values %v", validLabels)
			return
		}
		label := codeToLabel[labelFlag]

		if !isValidPrefecture(prefCodes) {
			log.Print("[error] Invalid prefecture code(s) provided.")
			return
		}

		if yearTo < yearFrom {
			log.Print("[error] Invalid year range provided")
			return
		}

		usecase.PrintPopulations(ctx, os.Stdout, label, prefCodes, yearFrom, yearTo)
	},
}

func init() {
	populationCmd.Flags().IntSliceP("prefectures", "p", []int{13}, "Comma-separated list of prefecture codes")
	populationCmd.Flags().StringP("label", "l", "all", "Population label (Allowed values: all, youth, productive, elderly)")
	populationCmd.Flags().IntP("from", "f", 1980, "Year from")
	populationCmd.Flags().IntP("to", "t", 2020, "Year to")
	rootCmd.AddCommand(populationCmd)
}

func isValidPrefecture(providedPrefs []int) bool {
	if len(providedPrefs) == 0 {
		return false
	}
	valid := true
	for _, providedPref := range providedPrefs {
		if !(1 <= providedPref && providedPref <= 47) {
			valid = false
		}
		if !valid {
			return false
		}
	}
	return true
}

func isValidLabel(providedLabel string, validOpts []string) bool {
	for _, opt := range validOpts {
		if providedLabel == opt {
			return true
		}
	}
	return false
}
