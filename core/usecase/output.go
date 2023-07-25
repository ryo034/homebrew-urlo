package usecase

import "urlo/core/domain"

type OutputPort interface {
	Cancelled()
	NotURLFound()
	NotFoundRecords()
	AlreadyExist(v domain.UrlMap)
	SetSuccessfully()
	DeleteSuccessfully(v domain.UrlMap)
	AddSuccessfully(v domain.UrlMap)
	UpdateSuccessfully()
	List(vs *domain.UrlMaps)
	ListAsJson(vs *domain.UrlMaps) error
	ListAsJsonString(vs *domain.UrlMaps) error
}
