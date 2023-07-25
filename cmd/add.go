package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"urlo/core/infrastructure/injector"
)

func newAddCmd(i injector.Injector) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a URL to the json",
		Long:  "Add a URL to the json",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			override, err := cmd.Flags().GetBool("override")
			if err != nil {
				color.Red("Error: %s\n", err)
				return
			}
			if err := i.Controller().Add(args, override); err != nil {
				color.Red("Error: %s\n", err)
				return
			}
		},
	}
	addCmd.Flags().BoolP("override", "o", false, "If the item exists, override it")
	return addCmd
}
