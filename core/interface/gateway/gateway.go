package gateway

import (
	"fmt"
	"urlo/core/domain"
	"urlo/core/driver"
)

type gateway struct {
	driver  driver.Driver
	adapter GatewayAdapter
}

func NewGateway(driver driver.Driver, adapter GatewayAdapter) domain.Repository {
	return &gateway{driver, adapter}
}

func (g *gateway) Find(t domain.Title) (domain.UrlMap, error) {
	rs, err := g.driver.ReadRecords()
	if err != nil {
		return domain.UrlMap{}, err
	}
	for idx, r := range rs {
		if r.Title == t.String() {
			return g.adapter.Adapt(rs[idx])
		}
	}
	return domain.UrlMap{}, fmt.Errorf("not found title: %s", t.String())
}

func (g *gateway) FindAll() (*domain.UrlMaps, error) {
	rs, err := g.driver.ReadRecords()
	if err != nil {
		return nil, err
	}
	return g.adapter.AdaptAll(rs)
}

func (g *gateway) Search(query string) (*domain.UrlMaps, error) {
	rs, err := g.driver.ReadRecords()
	if err != nil {
		return nil, err
	}
	res, err := g.adapter.AdaptAll(rs)
	if err != nil {
		return nil, err
	}
	return res.FilterByRegex(query)
}

func (g *gateway) Add(v domain.UrlMap) error {
	res, err := g.FindAll()
	if err != nil {
		return err
	}
	re, err := res.Add(v)
	if err != nil {
		return err
	}
	return g.driver.WriteAll(re)
}

func (g *gateway) Update(values *domain.UrlMaps) error {
	return g.driver.WriteAll(values)
}

func (g *gateway) Clear() error {
	return g.driver.ClearRecords()
}

func (g *gateway) AskOne(prompt string) (bool, error) {
	return g.driver.AskOne(prompt)
}

func (g *gateway) OpenUrl(u domain.StrictUrl) error {
	return g.driver.OpenUrl(u)
}

func (g *gateway) SelectOne(vs *domain.UrlMaps) (domain.Title, error) {
	s, err := g.driver.SelectOne(vs)
	if err != nil {
		return "", err
	}
	return domain.NewTitle(s)
}
