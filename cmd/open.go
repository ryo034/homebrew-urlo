package cmd

import (
	"fmt"
	"os/exec"
	"urlo/util"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a URL from the csv",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		records, err := util.GetRecordsFromOpenCscFile()
		if err != nil {
			return
		}
		if records.IsEmpty() {
			fmt.Println("No records found")
			return
		}

		// find the url
		for _, r := range records {
			if r.Title == title {
				if err = exec.Command("open", r.URL.String()).Start(); err != nil {
					fmt.Println(err)
					return
				}
				return
			}
		}
		fmt.Println("URL not found")
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		records, err := util.GetRecordsFromOpenCscFile()
		if err != nil {
			fmt.Println(err)
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		var titles []string
		for _, r := range records {
			titles = append(titles, r.Title)
		}

		fmt.Println(titles)
		return titles, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
