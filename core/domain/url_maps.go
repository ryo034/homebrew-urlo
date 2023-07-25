package domain

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"regexp"
	infraErr "urlo/core/infrastructure/error"
)

type UrlMapJson struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type UrlMaps struct {
	values []UrlMap
}

func (us *UrlMaps) Values() []UrlMap {
	return us.values
}

func (us *UrlMaps) Shift() UrlMap {
	return us.values[0]
}

func (us *UrlMaps) ToJson() []UrlMapJson {
	jsonValues := make([]UrlMapJson, len(us.values))
	for i, u := range us.values {
		jsonValues[i] = UrlMapJson{Title: u.Title().String(), URL: u.URL().String()}
	}
	return jsonValues
}

func NewUrlMaps(values []UrlMap) (*UrlMaps, error) {
	titleToUrls := make(map[string][]infraErr.UrlMap)
	for _, u := range values {
		titleToUrls[u.Title().String()] = append(titleToUrls[u.Title().String()], infraErr.UrlMap{
			Title: u.Title().String(),
			URL:   u.URL().String(),
		})
	}
	var duplicates []infraErr.DuplicatedUrlMaps
	for title, urls := range titleToUrls {
		if len(urls) > 1 {
			duplicates = append(duplicates, infraErr.DuplicatedUrlMaps{Title: title, Values: urls})
		}
	}
	if len(duplicates) > 0 {
		return nil, infraErr.NewDuplicationError(duplicates)
	}
	return &UrlMaps{values: values}, nil
}

func (us *UrlMaps) Delete(idx int) *UrlMaps {
	return &UrlMaps{values: append(us.values[:idx], us.values[idx+1:]...)}
}

func (us *UrlMaps) Update(value UrlMap) *UrlMaps {
	for i, item := range us.values {
		if item.Title().Eq(value.Title()) {
			us.values[i] = value
		}
	}
	return us
}

func (us *UrlMaps) IsAlreadyExist(title Title) bool {
	for _, item := range us.values {
		if item.Title().Eq(title) {
			return true
		}
	}
	return false
}

func (us *UrlMaps) FilterByRegex(query string) (*UrlMaps, error) {
	result := make([]UrlMap, 0)
	regex, err := regexp.Compile(query)
	if err != nil {
		fmt.Println("Invalid regex:", query)
		return nil, err
	}
	for _, item := range us.values {
		if regex.MatchString(item.Title().String()) {
			result = append(result, item)
		}
	}
	return &UrlMaps{values: result}, nil
}

func (us *UrlMaps) TitleMaxLen() int {
	var maxLen int
	for _, u := range us.values {
		rw := runewidth.StringWidth(u.Title().String())
		if rw > maxLen {
			maxLen = rw
		}
	}
	return maxLen
}

func (us *UrlMaps) Add(value UrlMap) (*UrlMaps, error) {
	if us.IsAlreadyExist(value.Title()) {
		return nil, fmt.Errorf("already exist: %s", value.Title().String())
	}
	return &UrlMaps{values: append(us.values, value)}, nil
}

func (us *UrlMaps) AddAll(values *UrlMaps) (*UrlMaps, error) {
	for _, item := range values.values {
		if us.IsAlreadyExist(item.Title()) {
			return nil, fmt.Errorf("already exist: %s", item.Title().String())
		}
	}
	return &UrlMaps{values: append(us.values, values.values...)}, nil
}

func (us *UrlMaps) Size() int {
	return len(us.values)
}

func (us *UrlMaps) IsEmpty() bool {
	return us.Size() == 0
}

func (us *UrlMaps) IsNotEmpty() bool {
	return !us.IsEmpty()
}

func (us *UrlMaps) GetTitles() []string {
	titles := make([]string, 0, us.Size())
	for _, item := range us.values {
		titles = append(titles, item.Title().String())
	}
	return titles
}

func (us *UrlMaps) GetItemFromLabel(label Title) (UrlMap, int, error) {
	for idx, item := range us.values {
		if item.Title().Eq(label) {
			return item, idx, nil
		}
	}
	return UrlMap{}, 0, fmt.Errorf("no item found for label %s", label)
}
