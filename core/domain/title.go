package domain

import "fmt"

type Title string

func NewTitle(v string) (Title, error) {
	if v == "" {
		return "", fmt.Errorf("title is empty")
	}
	return Title(v), nil
}

func (t Title) String() string {
	return string(t)
}

func (t Title) IsEmpty() bool {
	return t == ""
}

func (t Title) Eq(v Title) bool {
	return t.String() == v.String()
}
