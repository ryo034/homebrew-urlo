package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "urlo",
	Short: "A simple CLI for opening URLs from a csv file",
	Long:  "A simple CLI for opening URLs from a csv file",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
