package util

import (
	"github.com/manifoldco/promptui"
)

func PromptGetSelect(items UrlMaps) (UrlMap, int, error) {
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.Select{
			Label: "Select a Website",
			Items: items.GetTitles(),
		}

		index, result, err = prompt.Run()
	}

	if err != nil {
		return UrlMap{}, index, err
	}

	res, err := items.GetItemFromLabel(result)
	return res, index, err
}
