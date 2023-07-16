package cmd

import (
	"fmt"
	"os/exec"
	"urlo/util"

	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select and open a URL",
	Run: func(cmd *cobra.Command, args []string) {
		query, err := cmd.Flags().GetString("query")
		if err != nil {
			fmt.Println(err)
			return
		}

		records, err := util.GetRecordsFromOpenCscFile()
		if err != nil {
			return
		}
		if records.IsEmpty() {
			fmt.Println("No records found")
			fmt.Println("Try adding a URL with the 'add' command")
			fmt.Println("urlo add <title> <url>")
			return
		}

		result, _, err := util.PromptGetSelect(records.FilterByRegex(query))
		if err != nil {
			return
		}

		if err = exec.Command("open", result.URL.String()).Start(); err != nil {
			return
		}
	},
}

func init() {
	selectCmd.Flags().StringP("query", "q", "", "Query string for searching through URL titles")
	rootCmd.AddCommand(selectCmd)
}
