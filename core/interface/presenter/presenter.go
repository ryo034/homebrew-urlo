package presenter

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
	"urlo/core/domain"
	"urlo/core/usecase"
)

type presenter struct {
}

func NewPresenter() usecase.OutputPort {
	return &presenter{}
}

func (p *presenter) Cancelled() {
	fmt.Println("Cancelled")
}

func (p *presenter) NotURLFound() {
	color.Red("Error: No URL found\n")
}

func (p *presenter) NotFoundRecords() {
	color.Red("Records not found\n")
}

func (p *presenter) AlreadyExist(v domain.UrlMap) {
	color.Red("Error: Already exist title: '%s'\n", v.Title().String())
}

func (p *presenter) SetSuccessfully() {
	color.Green("Set successfully\n")
}

func (p *presenter) DeleteSuccessfully(v domain.UrlMap) {
	color.Green("Delete successfully %s\n", v.Title().String())
}

func (p *presenter) UpdateSuccessfully() {
	color.Green("Update successfully\n")
}

func (p *presenter) AddSuccessfully(v domain.UrlMap) {
	color.Green("Add successfully %s - %s\n", v.Title().String(), v.URL().String())
}

func (p *presenter) List(vs *domain.UrlMaps) {
	for _, r := range vs.Values() {
		fmt.Printf("%s - %s\n", runewidth.FillRight(r.Title().String(), vs.TitleMaxLen()), r.URL().String())
	}
}

func (p *presenter) ListAsJson(vs *domain.UrlMaps) error {
	jsonData, err := json.MarshalIndent(vs.ToJson(), "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))
	return nil
}

func (p *presenter) ListAsJsonString(vs *domain.UrlMaps) error {
	jsonData, err := json.Marshal(vs.ToJson())
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))
	return nil
}
