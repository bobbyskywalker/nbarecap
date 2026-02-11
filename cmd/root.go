package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nbarecap",
	Short: "Get access to NBA summaries from your terminal.",
	Long: `nbarecap provides several subcommands and flags to access NBA summaries for teams, players and dates.
Try: nbarecap games`,
}

func init() {
	gamesCmd.Flags().StringVarP(
		&date,
		"date",
		"d",
		time.Now().Format("2006-01-02"),
		"Game date (YYYY-MM-DD)",
	)

	standingsCmd.Flags().StringVarP(
		&season,
		"season",
		"s",
		buildDefaultSeason(),
		"NBA Season (YYYY-YY)",
	)

	rootCmd.AddCommand(gamesCmd)
	rootCmd.AddCommand(standingsCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
