package usecase

import "urlo/core/domain"

type InputPort interface {
	List(jsonOutput bool, jsonStringOutput bool) error
	Add(v domain.UrlMap, override bool) error
	BulkAdd(vs *domain.UrlMaps) error
	Update(vs *domain.UrlMaps) error
	Set(vs *domain.UrlMaps) error
	Open(t domain.Title) error
	Select(query string) error
	Delete() error
}
