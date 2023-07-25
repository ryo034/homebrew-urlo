package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"urlo/core/infrastructure/injector"
)

func newSetCmd(i injector.Injector) *cobra.Command {
	setCmd := &cobra.Command{
		Use:   "set",
		Short: "Set a new URL map from JSON string",
		Long:  `This command sets a new URL map from the provided JSON string.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := i.Controller().Set(args); err != nil {
				color.Red("Error: %s\n", err)
				return
			}
		},
	}
	return setCmd
}
