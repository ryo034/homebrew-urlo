package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Version = "1.0.0"
)

// go build -ldflags "-X main.Version=1.0.0" -o urlo main.go

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of urlo",
	Long:  `All software has versions. This is urlo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("urlo v" + Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
