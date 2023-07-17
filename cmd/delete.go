package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"urlo/util"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Select to delete a URL",
	Long:  "Select to delete a URL",
	Run: func(cmd *cobra.Command, args []string) {
		records, err := util.GetRecordsFromOpenCscFile()
		if err != nil {
			return
		}
		if records.IsEmpty() {
			fmt.Println("No records found")
			return
		}

		_, deleteTargetIdx, err := util.PromptGetSelect(records)
		if err != nil {
			return
		}

		err = util.WriteValuesToFile(append(records[:deleteTargetIdx], records[deleteTargetIdx+1:]...))
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Deleted successfully %s - %s\n", records[deleteTargetIdx].Title, records[deleteTargetIdx].URL.String())
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}