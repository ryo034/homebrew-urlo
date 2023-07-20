package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	version = "development"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of urlo",
	Long:  `All software has versions. This is urlo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("urlo v" + version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
