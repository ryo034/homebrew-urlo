package cmd

import (
	"github.com/fatih/color"
	"urlo/adapter"
	"urlo/util"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a URL to the json",
	Long:  "Add a URL to the json",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		override, err := cmd.Flags().GetBool("override")
		if err != nil {
			color.Red("Error: %s\n", err)
			return
		}
		title := args[0]
		if title == "" {
			color.Red("Error: Title is empty\n")
			return
		}
		u := args[1]
		if u == "" {
			color.Red("Error: URL is empty\n")
			return
		}

		res, err := adapter.AdaptUrlMapJsonToUrlMaps([]util.UrlMapJson{{Title: title, URL: u}})
		if err != nil {
			color.Red("Error: %s\n", err)
			return
		}

		addTarget := res.Shift()

		records, err := util.GetRecordsFromFile()
		if err != nil {
			color.Red("Error: %s\n", err)
			return
		}

		if override && records.IsAlreadyExist(title) {
			if err = util.WriteValuesToFile(records.Update(addTarget)); err != nil {
				color.Red("Error: %s\n", err)
				return
			}
			color.Green("Update successfully %s\n", title)
			return
		}

		if records.IsAlreadyExist(title) {
			color.Red("Error: Already exist title: '%s'\n", title)
			return
		}

		nrs, err := records.Add(addTarget)
		if err != nil {
			color.Red("Error: %s\n", err)
			return
		}
		err = util.WriteValuesToFile(nrs)
		if err != nil {
			color.Red("Error: %s\n", err)
			return
		}
		color.Green("Add successfully %s - %s\n", addTarget.Title, addTarget.URL.String())
	},
}

func init() {
	addCmd.Flags().BoolP("override", "o", false, "If the item exists, override it")
	rootCmd.AddCommand(addCmd)
}
