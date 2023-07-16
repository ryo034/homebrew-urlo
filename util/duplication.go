package util

import (
	"fmt"
	"github.com/fatih/color"
)

type DuplicatedUrlMaps struct {
	Values UrlMaps
	Title  string
}

func (us UrlMaps) GetDuplications() []DuplicatedUrlMaps {
	titleToUrls := make(map[string]UrlMaps)
	for _, u := range us {
		titleToUrls[u.Title] = append(titleToUrls[u.Title], u)
	}
	var duplicates []DuplicatedUrlMaps
	for title, urls := range titleToUrls {
		if len(urls) > 1 {
			duplicates = append(duplicates, DuplicatedUrlMaps{Title: title, Values: urls})
		}
	}
	return duplicates
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
				color.Yellow(fmt.Sprintf("- %s\n", dup.URL.String()))
			}
		}
	}
	return ""
}
