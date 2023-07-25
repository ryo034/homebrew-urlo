package error

import (
	"fmt"
	"github.com/fatih/color"
)

type UrlMap struct {
	Title string
	URL   string
}

type DuplicatedUrlMaps struct {
	Values []UrlMap
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
			for _, dup := range dups.Values {
				color.Yellow(fmt.Sprintf("- %s\n", dup.URL))
			}
		}
	}
	return ""
}
