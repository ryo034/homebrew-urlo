package util

import (
	"github.com/AlecAivazis/survey/v2"
)

func PromptGetSelect(items UrlMaps) (UrlMap, int, error) {
	var result string
	for result == "" {
		prompt := &survey.Select{
			Message: "Select a Website",
			Options: items.GetTitles(),
		}
		if err := survey.AskOne(prompt, &result); err != nil {
			return UrlMap{}, 0, err
		}
	}
	return items.GetItemFromLabel(result)
}

func Confirm(prompt string) bool {
	confirmOk := false
	p := &survey.Confirm{
		Message: prompt,
	}
	if err := survey.AskOne(p, &confirmOk); err != nil {
		return false
	}
	return confirmOk
}
