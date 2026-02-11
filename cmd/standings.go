package cmd

import (
	"errors"
	"fmt"
	"nbarecap/internal/nba"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var season string

var standingsCmd = &cobra.Command{
	Use:   "standings",
	Short: "View NBA standings for a season (defaults to current season).",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(season)
		err := validateSeason(season)
		if err != nil {
			return err
		}
		//TODO: standings.startUi(season)
		standings, err := nba.GetStandingsForSeason(season)
		fmt.Println("team: " + standings[1].TeamName + "\nrecord: " + standings[1].Record)
		return nil
	},
}

func validateSeason(season string) error {
	if len(season) != 7 {
		return errors.New("invalid season, argument too short. Valid format: \"2025-26\"")
	}
	firstYearTail, err := strconv.Atoi(season[2:4])
	secondYearTail, err1 := strconv.Atoi(season[5:7])
	if err != nil || err1 != nil || firstYearTail != secondYearTail-1 {
		return errors.New("invalid season year. Valid format: \"2025-26\"")
	}
	return nil
}

func buildDefaultSeason() string {
	currentYear := time.Now().Year()
	currentMonth := time.Now().Month()

	if currentMonth < 10 {
		return strconv.Itoa(currentYear-1) + "-" + (strconv.Itoa(currentYear)[2:])
	}
	return strconv.Itoa(currentYear) + "-" + (strconv.Itoa(currentYear + 1)[2:])
}
