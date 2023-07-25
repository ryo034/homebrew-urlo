package controller

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"urlo/core/domain"
	"urlo/core/usecase"
)

type Controller struct {
	interactor usecase.InputPort
}

func NewController(interactor usecase.InputPort) *Controller {
	return &Controller{interactor}
}

func (c *Controller) Add(args []string, override bool) error {
	title := args[0]
	if title == "" {
		color.Red("Error: Title is empty\n")
		return nil
	}
	ur := args[1]
	if ur == "" {
		color.Red("Error: URL is empty\n")
		return nil
	}

	t, err := domain.NewTitle(title)
	if err != nil {
		return err
	}

	u, err := domain.NewStrictUrlFromString(ur)
	if err != nil {
		return err
	}
	if err = c.interactor.Add(domain.NewUrlMap(t, u), override); err != nil {
		return err
	}
	return nil
}

func (c *Controller) BulkAdd(args []string) error {
	jsonString := args[0]
	if jsonString == "" {
		color.Red("Error: input list is empty\n")
		return nil
	}

	var newUrlMap []domain.UrlMapJson
	if err := json.Unmarshal([]byte(args[0]), &newUrlMap); err != nil {
		if err := fmt.Errorf("failed to parse JSON string: %w", err); err != nil {
			fmt.Println(err)
		}
		return nil
	}

	result := make([]domain.UrlMap, len(newUrlMap))
	for i, v := range newUrlMap {
		t, err := domain.NewTitle(v.Title)
		if err != nil {
			return err
		}
		u, err := domain.NewStrictUrlFromString(v.URL)
		if err != nil {
			return err
		}
		result[i] = domain.NewUrlMap(t, u)
	}

	res, err := domain.NewUrlMaps(result)
	if err != nil {
		return err
	}

	if err := c.interactor.BulkAdd(res); err != nil {
		return err
	}
	return nil
}

func (c *Controller) List(jsonOutput bool, jsonStringOutput bool) error {
	if jsonOutput && jsonStringOutput {
		fmt.Println("Can't use both -j and -s")
		return nil
	}
	if err := c.interactor.List(jsonOutput, jsonStringOutput); err != nil {
		return err
	}
	return nil
}

func (c *Controller) Set(args []string) error {
	var newUrlMap []domain.UrlMapJson
	if err := json.Unmarshal([]byte(args[0]), &newUrlMap); err != nil {
		if err := fmt.Errorf("failed to parse JSON string: %w", err); err != nil {
			fmt.Println(err)
		}
		return nil
	}

	vs := make([]domain.UrlMap, len(newUrlMap))
	for i, v := range newUrlMap {
		t, err := domain.NewTitle(v.Title)
		if err != nil {
			return err
		}
		u, err := domain.NewStrictUrlFromString(v.URL)
		if err != nil {
			return err
		}
		vs[i] = domain.NewUrlMap(t, u)
	}
	result, err := domain.NewUrlMaps(vs)
	if err != nil {
		return err
	}

	if err := c.interactor.Set(result); err != nil {
		return err
	}
	return nil
}

func (c *Controller) Open(args []string) error {
	title := args[0]
	t, err := domain.NewTitle(title)
	if err != nil {
		return err
	}
	if err := c.interactor.Open(t); err != nil {
		return err
	}
	return nil
}

func (c *Controller) Delete() error {
	return c.interactor.Delete()
}

func (c *Controller) Select(query string) error {
	return c.interactor.Select(query)
}
