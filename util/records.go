package util

import (
	"encoding/json"
	"fmt"
	"github.com/mattn/go-runewidth"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

type UrlMap struct {
	Title string
	URL   *url.URL
}

type UrlMapJson struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type UrlMaps struct {
	values []UrlMap
}

func (us UrlMaps) Values() []UrlMap {
	return us.values
}

func (us UrlMaps) Shift() UrlMap {
	return us.values[0]
}

func (us UrlMaps) ToJson() []UrlMapJson {
	jsonValues := make([]UrlMapJson, len(us.values))
	for i, u := range us.values {
		jsonValues[i] = UrlMapJson{Title: u.Title, URL: u.URL.String()}
	}
	return jsonValues
}

func NewUrlMaps(values []UrlMap) (UrlMaps, error) {
	titleToUrls := make(map[string][]UrlMap)
	for _, u := range values {
		if u.Title == "" {
			return UrlMaps{}, fmt.Errorf("title is required")
		}
		if u.URL == nil {
			return UrlMaps{}, fmt.Errorf("url is required")
		}
		titleToUrls[u.Title] = append(titleToUrls[u.Title], u)
	}
	var duplicates []DuplicatedUrlMaps
	for title, urls := range titleToUrls {
		if len(urls) > 1 {
			duplicates = append(duplicates, DuplicatedUrlMaps{Title: title, Values: UrlMaps{values: urls}})
		}
	}
	if len(duplicates) > 0 {
		return UrlMaps{}, NewDuplicationError(duplicates)
	}
	return UrlMaps{values: values}, nil
}

func (us UrlMaps) Delete(idx int) UrlMaps {
	return UrlMaps{values: append(us.values[:idx], us.values[idx+1:]...)}
}

func (us UrlMaps) Update(value UrlMap) UrlMaps {
	for i, u := range us.values {
		if u.Title == value.Title {
			us.values[i] = value
		}
	}
	return us
}

func (us UrlMaps) IsAlreadyExist(title string) bool {
	for _, item := range us.values {
		if item.Title == title {
			return true
		}
	}
	return false
}

func (us UrlMaps) FilterByRegex(query string) (UrlMaps, error) {
	result := make([]UrlMap, 0)
	regex, err := regexp.Compile(query)
	if err != nil {
		fmt.Println("Invalid regex:", query)
		return UrlMaps{}, err
	}
	for _, item := range us.values {
		if regex.MatchString(item.Title) {
			result = append(result, item)
		}
	}
	return UrlMaps{values: result}, nil
}

func (us UrlMaps) TitleMaxLen() int {
	var maxLen int
	for _, u := range us.values {
		if runewidth.StringWidth(u.Title) > maxLen {
			maxLen = runewidth.StringWidth(u.Title)
		}
	}
	return maxLen
}

func (us UrlMaps) Add(value UrlMap) (UrlMaps, error) {
	if us.IsAlreadyExist(value.Title) {
		return UrlMaps{}, fmt.Errorf("already exist: %s", value.Title)
	}
	return UrlMaps{values: append(us.values, value)}, nil
}

func (us UrlMaps) AddAll(values UrlMaps) (UrlMaps, error) {
	for _, item := range values.values {
		if us.IsAlreadyExist(item.Title) {
			return UrlMaps{}, fmt.Errorf("already exist: %s", item.Title)
		}
	}
	return UrlMaps{values: append(us.values, values.values...)}, nil
}

func (us UrlMaps) Size() int {
	return len(us.values)
}

func (us UrlMaps) IsEmpty() bool {
	return us.Size() == 0
}

func (us UrlMaps) IsNotEmpty() bool {
	return !us.IsEmpty()
}

func (us UrlMaps) GetTitles() []string {
	titles := make([]string, 0, us.Size())
	for _, item := range us.values {
		titles = append(titles, item.Title)
	}
	return titles
}

func (us UrlMaps) GetItemFromLabel(label string) (UrlMap, int, error) {
	for idx, item := range us.values {
		if item.Title == label {
			return item, idx, nil
		}
	}
	return UrlMap{}, 0, fmt.Errorf("no item found for label %s", label)
}

func ClearRecords() error {
	f, err := os.Create(FileRelativePath)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			if err = fmt.Errorf("defer close error: %v", closeErr); err != nil {
				fmt.Println(err)
			}
		}
	}()
	b, _ := json.Marshal([]UrlMapJson{})
	if _, err = f.Write(b); err != nil {
		return err
	}
	return nil
}

func WriteValuesToFile(values UrlMaps) error {
	f, err := os.Create(FileRelativePath)
	if err != nil {
		return fmt.Errorf("create error: %v", err)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			if err = fmt.Errorf("defer close error: %v", closeErr); err != nil {
				fmt.Println(err)
			}
		}
	}()

	urlMapJsons := make([]UrlMapJson, 0, values.Size())
	for _, item := range values.values {
		urlMapJsons = append(urlMapJsons, UrlMapJson{Title: item.Title, URL: item.URL.String()})
	}
	b, err := json.Marshal(urlMapJsons)
	if err != nil {
		return fmt.Errorf("json marshal error: %v", err)
	}
	if _, err = f.Write(b); err != nil {
		return fmt.Errorf("write error: %v", err)
	}
	return nil
}

func GetRecordsFromFile() (UrlMaps, error) {
	var urlMaps []UrlMapJson

	_, err := os.Stat(FileRelativePath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(filepath.Dir(FileRelativePath), 0755); err != nil {
			return UrlMaps{}, fmt.Errorf("mkdir error: %v", err)
		}

		f, err := os.Create(FileRelativePath)
		if err != nil {
			return UrlMaps{}, fmt.Errorf("create error: %v", err)
		}
		defer func() {
			if closeErr := f.Close(); closeErr != nil {
				if err = fmt.Errorf("defer close error: %v", closeErr); err != nil {
					fmt.Println(err)
				}
			}
		}()
		return UrlMaps{}, nil
	} else {
		// If the file exists, read it
		f, err := os.Open(FileRelativePath)
		if err != nil {
			return UrlMaps{}, fmt.Errorf("open error: %v", err)
		}
		defer func() {
			if closeErr := f.Close(); closeErr != nil {
				if err = fmt.Errorf("defer close error: %v", closeErr); err != nil {
					fmt.Println(err)
				}
			}
		}()

		if err = json.NewDecoder(f).Decode(&urlMaps); err != nil {
			// If the file is empty, return empty
			if err == io.EOF {
				return UrlMaps{}, nil
			}
			return UrlMaps{}, err
		}
		result := make([]UrlMap, 0)
		for _, item := range urlMaps {
			u, err := url.Parse(item.URL)
			if err != nil {
				return UrlMaps{}, err
			}
			result = append(result, UrlMap{Title: item.Title, URL: u})
		}
		return NewUrlMaps(result)
	}
}
