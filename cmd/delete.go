package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"urlo/core/infrastructure/injector"
)

func newDeleteCmd(i injector.Injector) *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Select to delete a URL",
		Long:  "Select to delete a URL",
		Run: func(cmd *cobra.Command, args []string) {
			if err := i.Controller().Delete(); err != nil {
				color.Red("Error: %s\n", err)
				return
			}
		},
	}
	return deleteCmd
}
