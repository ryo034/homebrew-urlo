package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"urlo/core/infrastructure/injector"
)

func newBulkAddCmd(i injector.Injector) *cobra.Command {
	bulkAddCmd := &cobra.Command{
		Use:   "bulk-add",
		Short: "Bulk add the new URL map",
		Long:  `Bdd add the new URL map`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := i.Controller().BulkAdd(args); err != nil {
				color.Red("Error: %s\n", err)
				return
			}
		},
	}
	return bulkAddCmd
}
