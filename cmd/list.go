package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mattn/go-runewidth"
	"urlo/util"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all URLs from the json",
	Run: func(cmd *cobra.Command, args []string) {
		records, err := util.GetRecordsFromFile()
		if err != nil {
			return
		}
		if records.IsEmpty() {
			fmt.Println("No records found")
			return
		}

		jsonOutput, err := cmd.Flags().GetBool("json")
		if err != nil {
			fmt.Println(err)
			return
		}
		jsonStringOutput, err := cmd.Flags().GetBool("string")
		if err != nil {
			fmt.Println(err)
			return
		}

		// If both j and s are specified at the same time, return an error
		if jsonOutput && jsonStringOutput {
			fmt.Println("Can't use both -j and -s")
			return
		}

		if jsonOutput {
			jsonData, err := json.MarshalIndent(records.ToJson(), "", "  ")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(jsonData))
			return
		}

		if jsonStringOutput {
			jsonData, err := json.Marshal(records.ToJson())
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(jsonData))
			return
		}

		for _, r := range records.Values() {
			fmt.Printf("%s - %s\n", runewidth.FillRight(r.Title, records.TitleMaxLen()), r.URL.String())
		}
	},
}

func init() {
	listCmd.Flags().BoolP("json", "j", false, "Output in JSON format")
	listCmd.Flags().BoolP("string", "s", false, "Output in JSON String format")
	rootCmd.AddCommand(listCmd)
}
