package domain

type Repository interface {
	Find(t Title) (UrlMap, error)
	FindAll() (*UrlMaps, error)
	Search(query string) (*UrlMaps, error)
	Add(v UrlMap) error
	Update(vs *UrlMaps) error
	Clear() error
	AskOne(prompt string) (bool, error)
	OpenUrl(u StrictUrl) error
	SelectOne(vs *UrlMaps) (Title, error)
}
