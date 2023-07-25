package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"urlo/core/infrastructure"
	"urlo/core/infrastructure/injector"
)

var rootCmd = &cobra.Command{
	Use:   "urlo",
	Short: "A simple CLI for opening URLs from a json file",
	Long:  "A simple CLI for opening URLs from a json file",
}

func Execute() {
	container := injector.NewInjector(
		infrastructure.FileRelativePath,
		infrastructure.NewCommandExecutor(),
		infrastructure.NewPromptExecutor(),
	)

	rootCmd.AddCommand(newAddCmd(container))
	rootCmd.AddCommand(newListCmd(container))
	rootCmd.AddCommand(newBulkAddCmd(container))
	rootCmd.AddCommand(newSetCmd(container))
	rootCmd.AddCommand(newOpenCmd(container))
	rootCmd.AddCommand(newDeleteCmd(container))
	rootCmd.AddCommand(newSelectCmd(container))
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
