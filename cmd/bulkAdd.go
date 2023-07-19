package cmd

import (
	"encoding/json"
	"fmt"
	"urlo/adapter"
	"urlo/util"

	"github.com/spf13/cobra"
)

var bulkAddCmd = &cobra.Command{
	Use:   "bulk-add",
	Short: "Bulk add the new URL map",
	Long:  `Bdd add the new URL map`,
	Run: func(cmd *cobra.Command, args []string) {
		var newUrlMap []util.UrlMapJson
		if args[0] == "" {
			if err := fmt.Errorf("JSON string is empty"); err != nil {
				fmt.Println(err)
			}
			return
		}
		if err := json.Unmarshal([]byte(args[0]), &newUrlMap); err != nil {
			if err := fmt.Errorf("failed to parse JSON string: %w", err); err != nil {
				fmt.Println(err)
			}
			return
		}

		records, err := util.GetRecordsFromFile()
		if err != nil {
			fmt.Println(err)
			return
		}

		res, err := adapter.AdaptUrlMapJsonToUrlMaps(newUrlMap)
		if err != nil {
			fmt.Println(err)
			return
		}
		rs, err := records.AddAll(res)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = util.WriteValuesToFile(rs)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Successfully add all the new URL map.")
		return
	},
}

func init() {
	rootCmd.AddCommand(bulkAddCmd)
}
