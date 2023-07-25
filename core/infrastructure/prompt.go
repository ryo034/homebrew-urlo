package infrastructure

import (
	"github.com/AlecAivazis/survey/v2"
)

type PromptExecutor interface {
	SelectOne(msg string, options []string) (string, error)
	AskOne(prompt string) (bool, error)
}

type promptExecutor struct{}

func NewPromptExecutor() PromptExecutor {
	return &promptExecutor{}
}

func (p *promptExecutor) SelectOne(msg string, options []string) (string, error) {
	var selected string
	if err := survey.AskOne(&survey.Select{
		Message: msg,
		Options: options,
	}, &selected); err != nil {
		return "", err
	}
	return selected, nil
}

func (p *promptExecutor) AskOne(prompt string) (bool, error) {
	confirmOk := false
	if err := survey.AskOne(&survey.Confirm{Message: prompt}, &confirmOk); err != nil {
		return false, err
	}
	return confirmOk, nil
}
