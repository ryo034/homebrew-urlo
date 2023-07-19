package adapter

import (
	"fmt"
	"net/url"
	"urlo/util"
)

func CheckUrlStrictly(v string) (*url.URL, error) {
	u, err := url.Parse(v)
	if err != nil {
		return nil, err
	}
	// check url validation
	if !u.IsAbs() {
		return nil, fmt.Errorf("URL must be absolute: %s", u.String())
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, fmt.Errorf("URL scheme must be http or https: %s", u.Scheme)
	}
	if u.Host == "" {
		return nil, fmt.Errorf("URL host must not be empty")
	}
	return u, nil
}

func AdaptUrlMapJsonToUrlMaps(values []util.UrlMapJson) (util.UrlMaps, error) {
	urlMaps := make([]util.UrlMap, len(values))
	for i, v := range values {
		u, err := CheckUrlStrictly(v.URL)
		if err != nil {
			return util.UrlMaps{}, err
		}
		urlMaps[i] = util.UrlMap{Title: v.Title, URL: u}
	}
	return util.NewUrlMaps(urlMaps)
}
