package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"urlo/core/infrastructure/injector"
)

func newListCmd(i injector.Injector) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all URLs from the json",
		Run: func(cmd *cobra.Command, args []string) {
			jsonOutput, err := cmd.Flags().GetBool("json")
			if err != nil {
				color.Red("Error: %s\n", err)
				return
			}
			jsonStringOutput, err := cmd.Flags().GetBool("string")
			if err != nil {
				color.Red("Error: %s\n", err)
				return
			}
			if err := i.Controller().List(jsonOutput, jsonStringOutput); err != nil {
				color.Red("Error: %s\n", err)
				return
			}
		},
	}
	listCmd.Flags().BoolP("json", "j", false, "Output in JSON format")
	listCmd.Flags().BoolP("string", "s", false, "Output in JSON String format")
	return listCmd
}
