package cmd

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"urlo/util"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all URLs from the csv",
	Run: func(cmd *cobra.Command, args []string) {
		// get the flags
		showURLs, err := cmd.Flags().GetBool("urls")
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
			return
		}

		ml := records.TitleMaxLen()
		for _, r := range records {
			if showURLs {
				fmt.Printf("%s - %s\n", runewidth.FillRight(r.Title, ml), r.URL.String())
			} else {
				fmt.Println(r.Title)
			}
		}
	},
}

func init() {
	listCmd.Flags().BoolP("urls", "u", false, "Show the URLs as well as the titles")
	rootCmd.AddCommand(listCmd)
}
