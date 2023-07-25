package gateway

import (
	"urlo/core/domain"
)

type GatewayAdapter interface {
	Adapt(v domain.UrlMapJson) (domain.UrlMap, error)
	AdaptAll(vs []domain.UrlMapJson) (*domain.UrlMaps, error)
}

type gatewayAdapter struct{}

func NewGatewayAdapter() GatewayAdapter {
	return &gatewayAdapter{}
}

func (g *gatewayAdapter) Adapt(v domain.UrlMapJson) (domain.UrlMap, error) {
	u, err := domain.NewStrictUrlFromString(v.URL)
	if err != nil {
		return domain.UrlMap{}, err
	}
	t, err := domain.NewTitle(v.Title)
	if err != nil {
		return domain.UrlMap{}, err
	}
	return domain.NewUrlMap(t, u), nil
}

func (g *gatewayAdapter) AdaptAll(vs []domain.UrlMapJson) (*domain.UrlMaps, error) {
	urlMaps := make([]domain.UrlMap, len(vs))
	for i, v := range vs {
		res, err := g.Adapt(v)
		if err != nil {
			return nil, err
		}
		urlMaps[i] = res
	}
	return domain.NewUrlMaps(urlMaps)
}
