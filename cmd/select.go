package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"urlo/core/infrastructure/injector"
)

func newSelectCmd(i injector.Injector) *cobra.Command {
	selectCmd := &cobra.Command{
		Use:   "select",
		Short: "Select and open a URL",
		Run: func(cmd *cobra.Command, args []string) {
			query, err := cmd.Flags().GetString("query")
			if err != nil {
				fmt.Println(err)
				return
			}
			if err := i.Controller().Select(query); err != nil {
				color.Red("Error: %s\n", err)
				return
			}
		},
	}
	selectCmd.Flags().StringP("query", "q", "", "Query string for searching through URL titles")
	return selectCmd
}
