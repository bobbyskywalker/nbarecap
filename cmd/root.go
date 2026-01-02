package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nbarecap",
	Short: "Get access to NBA summaries from your terminal.",
	Long: `nbarecap provides several subcommands and flags to access NBA summaries for teams, players and dates.
Try: nbarecap today`,
}

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "View today's scores.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("today's scores here")
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
