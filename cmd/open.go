package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"urlo/core/infrastructure/injector"
)

func newOpenCmd(i injector.Injector) *cobra.Command {
	openCmd := &cobra.Command{
		Use:   "open",
		Short: "Open a URL from the json",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := i.Controller().Open(args); err != nil {
				color.Red("Error: %s\n", err)
				return
			}
		},
	}
	return openCmd
}
