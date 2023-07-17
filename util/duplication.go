package util

import (
	"fmt"
	"github.com/fatih/color"
)

type DuplicatedUrlMaps struct {
	Values UrlMaps
	Title  string
}

type DuplicationError struct {
	Values []DuplicatedUrlMaps
}

func NewDuplicationError(values []DuplicatedUrlMaps) error {
	return &DuplicationError{
		Values: values,
	}
}

func (e *DuplicationError) Error() string {
	color.Red("Command can not be executed because of the following errors:\n\n")
	if len(e.Values) > 0 {
		for _, dups := range e.Values {
			color.Red("Error: The title '%s' is already used for the following URLs:\n", dups.Title)
			for _, dup := range dups.Values.values {
				color.Yellow(fmt.Sprintf("- %s\n", dup.URL.String()))
			}
		}
	}
	return ""
}
