package domain

type UrlMap struct {
	title Title
	url   StrictUrl
}

func (u UrlMap) Title() Title {
	return u.title
}

func (u UrlMap) URL() StrictUrl {
	return u.url
}

func NewUrlMap(title Title, url StrictUrl) UrlMap {
	return UrlMap{title, url}
}
