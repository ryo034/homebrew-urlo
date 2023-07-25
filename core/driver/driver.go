package driver

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"urlo/core/domain"
	"urlo/core/infrastructure"
)

type Driver interface {
	ReadRecords() ([]domain.UrlMapJson, error)
	Write(v domain.UrlMap) error
	WriteAll(vs *domain.UrlMaps) error
	ClearRecords() error
	OpenUrl(u domain.StrictUrl) error
	AskOne(prompt string) (bool, error)
	SelectOne(vs *domain.UrlMaps) (string, error)
}

type driver struct {
	filePath       string
	cmdExecutor    infrastructure.CommandExecutor
	promptExecutor infrastructure.PromptExecutor
}

func NewDriver(filePath string, cmdExecutor infrastructure.CommandExecutor, promptExecutor infrastructure.PromptExecutor) Driver {
	return &driver{filePath, cmdExecutor, promptExecutor}
}

func (d *driver) getFile() (*os.File, error) {
	_, err := os.Stat(d.filePath)
	if os.IsNotExist(err) {
		f, err := os.Create(d.filePath)
		if err != nil {
			return nil, fmt.Errorf("create error: %v", err)
		}
		return f, nil
	}
	f, err := os.OpenFile(d.filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("open error: %v", err)
	}
	return f, nil
}

func (d *driver) ReadRecords() ([]domain.UrlMapJson, error) {
	var urlMaps []domain.UrlMapJson
	f, err := d.getFile()
	if err != nil {
		return nil, fmt.Errorf("get file error: %v", err)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			if err = fmt.Errorf("defer close error: %v", closeErr); err != nil {
				fmt.Println(err)
			}
		}
	}()

	if err := json.NewDecoder(f).Decode(&urlMaps); err != nil {
		// If the file is empty, return empty
		if err == io.EOF {
			return nil, nil
		}
		return nil, fmt.Errorf("json decode error: %v", err)
	}
	return urlMaps, nil
}

func (d *driver) Write(v domain.UrlMap) error {
	vs, _ := domain.NewUrlMaps([]domain.UrlMap{v})
	return d.WriteAll(vs)
}

func (d *driver) WriteAll(vs *domain.UrlMaps) error {
	f, err := d.getFile()
	if err != nil {
		return fmt.Errorf("get file error: %v", err)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			if err = fmt.Errorf("defer close error: %v", closeErr); err != nil {
				fmt.Println(err)
			}
		}
	}()

	urlMapJsons := make([]domain.UrlMapJson, 0, vs.Size())
	for _, item := range vs.Values() {
		urlMapJsons = append(urlMapJsons, domain.UrlMapJson{Title: item.Title().String(), URL: item.URL().String()})
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

func (d *driver) ClearRecords() error {
	f, err := d.getFile()
	if err != nil {
		return fmt.Errorf("get file error: %v", err)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			if err = fmt.Errorf("defer close error: %v", closeErr); err != nil {
				fmt.Println(err)
			}
		}
	}()

	b, _ := json.Marshal([]domain.UrlMapJson{})
	if _, err := f.Write(b); err != nil {
		return err
	}
	return nil
}

func (d *driver) OpenUrl(u domain.StrictUrl) error {
	return d.cmdExecutor.Open(u)
}

func (d *driver) AskOne(prompt string) (bool, error) {
	return d.promptExecutor.AskOne(prompt)
}

func (d *driver) SelectOne(vs *domain.UrlMaps) (string, error) {
	return d.promptExecutor.SelectOne("Select title", vs.GetTitles())
}
