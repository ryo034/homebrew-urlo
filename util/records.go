package util

import (
	"encoding/csv"
	"fmt"
	"github.com/mattn/go-runewidth"
	"net/url"
	"os"
	"regexp"
)

type UrlMap struct {
	Title string
	URL   *url.URL
}

type UrlMaps []UrlMap

func (us UrlMaps) IsAlreadyExist(title string) bool {
	for _, item := range us {
		if item.Title == title {
			return true
		}
	}
	return false
}

func (us UrlMaps) FilterByRegex(query string) UrlMaps {
	var result UrlMaps
	regex, err := regexp.Compile(query)
	if err != nil {
		fmt.Println("Invalid regex:", query)
		return result
	}
	for _, item := range us {
		if regex.MatchString(item.Title) {
			result = append(result, item)
		}
	}
	return result
}

func (us UrlMaps) TitleMaxLen() int {
	var maxLen int
	for _, u := range us {
		if runewidth.StringWidth(u.Title) > maxLen {
			maxLen = runewidth.StringWidth(u.Title)
		}
	}
	return maxLen
}

func (us UrlMaps) Add(value UrlMap) UrlMaps {
	return append(us, value)
}

func (us UrlMaps) Size() int {
	return len(us)
}

func (us UrlMaps) IsEmpty() bool {
	return us.Size() == 0
}

func (us UrlMaps) IsNotEmpty() bool {
	return !us.IsEmpty()
}

func ConvertToUrlMaps(items [][]string) (UrlMaps, error) {
	var urlMaps UrlMaps
	for _, item := range items {
		nu, err := url.Parse(item[1])
		if err != nil {
			fmt.Printf("parse error: %v", err)
			return nil, fmt.Errorf("parse error: %v", err)
		}
		urlMaps = append(urlMaps, UrlMap{Title: item[0], URL: nu})
	}
	return urlMaps, nil

}

func (us UrlMaps) GetTitles() []string {
	titles := make([]string, 0, us.Size())
	for _, item := range us {
		titles = append(titles, item.Title)
	}
	return titles
}

func (us UrlMaps) GetItemFromLabel(label string) (UrlMap, error) {
	for _, item := range us {
		if item.Title == label {
			return item, nil
		}
	}
	return UrlMap{}, fmt.Errorf("no item found for label %s", label)
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

	// create a new csv writer
	w := csv.NewWriter(f)
	defer w.Flush()

	// write a record
	for _, value := range values {
		if err := w.Write([]string{value.Title, value.URL.String()}); err != nil {
			return fmt.Errorf("write error: %v", err)
		}
	}
	return nil
}

func GetRecordsFromOpenCscFile() (UrlMaps, error) {
	f, err := os.OpenFile(FileRelativePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, fmt.Errorf("open error: %v", err)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			if err = fmt.Errorf("defer close error: %v", closeErr); err != nil {
				fmt.Println(err)
			}
		}
	}()

	// create a new csv reader
	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read error: %v", err)
	}
	result := make([][]string, 0, len(records))
	// print the records
	for _, record := range records {
		if len(record) == 0 || (len(record) > 0 && record[0] == "") {
			// skip empty lines
			continue
		}
		result = append(result, record)
	}
	rs, err := ConvertToUrlMaps(result)
	if err != nil {
		return nil, err
	}

	ds := rs.GetDuplications()
	if len(ds) != 0 {
		err = NewDuplicationError(ds)
		err.Error()
		return nil, err
	}
	return rs, nil
}
