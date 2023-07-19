package cmd

import (
	"encoding/json"
	"fmt"
	"urlo/adapter"
	"urlo/util"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a new URL map from JSON string",
	Long:  `This command sets a new URL map from the provided JSON string.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var newUrlMap []util.UrlMapJson
		//json string出ない場合にエラーメッセージを表示
		if args[0] == "" {
			return fmt.Errorf("JSON string is empty")
		}
		if err := json.Unmarshal([]byte(args[0]), &newUrlMap); err != nil {
			return fmt.Errorf("failed to parse JSON string: %w", err)
		}

		records, err := util.GetRecordsFromFile()
		if err != nil {
			fmt.Println(err)
			return err
		}
		if records.IsEmpty() {
			fmt.Println("No records found")
			return nil
		}

		res, err := adapter.AdaptUrlMapJsonToUrlMaps(newUrlMap)
		if err != nil {
			return err
		}
		rs, err := records.AddAll(res)
		if err != nil {
			return err
		}

		err = util.WriteValuesToFile(rs)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		fmt.Println("Successfully set the new URL map.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
