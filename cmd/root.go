package cmd

import (
	"errors"
	"log"
	"nbarecap/internal/ui"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nbarecap",
	Short: "Get access to NBA summaries from your terminal.",
	Long: `nbarecap provides several subcommands and flags to access NBA summaries for teams, players and dates.
Try: nbarecap games`,
}

var gamesCmd = &cobra.Command{
	Use:   "games",
	Short: "View NBA games for a date (defaults to today).",
	RunE: func(cmd *cobra.Command, args []string) error {
		date, err := time.Parse("2006-01-02", date)
		if err != nil {
			return errors.New("invalid date: correct format is YYYY-MM-DD")
		}
		ui.StartUi(date)
		return nil
	},
}

var date string

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
		log.Fatal(err)
	}
}
