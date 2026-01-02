package cmd

import (
	"nbarecap/internal/scores"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nbarecap",
	Short: "Get access to NBA summaries from your terminal.",
	Long: `nbarecap provides several subcommands and flags to access NBA summaries for teams, players and dates.
Try: nbarecap games`,
}

var date string

var gamesCmd = &cobra.Command{
	Use:   "games",
	Short: "View NBA games for a date (defaults to today).",
	RunE: func(cmd *cobra.Command, args []string) error {
		return scores.GetAllGamesForDate(&date)
	},
}

func init() {
	gamesCmd.Flags().StringVarP(
		&date,
		"date",
		"d",
		time.Now().Format("2006-01-02"),
		"Game date (YYYY-MM-DD)",
	)
	rootCmd.AddCommand(gamesCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
