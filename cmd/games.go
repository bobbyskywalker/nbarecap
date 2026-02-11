package cmd

import (
	"errors"
	"nbarecap/internal/ui/games"
	"time"

	"github.com/spf13/cobra"
)

var date string

var gamesCmd = &cobra.Command{
	Use:   "games",
	Short: "View NBA games for a date (defaults to today).",
	RunE: func(cmd *cobra.Command, args []string) error {
		date, err := time.Parse("2006-01-02", date)
		if err != nil {
			return errors.New("invalid date: correct format is YYYY-MM-DD")
		}
		games.StartUi(date)
		return nil
	},
}
