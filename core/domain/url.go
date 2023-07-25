package domain

import (
	"fmt"
	"net/url"
)

type StrictUrl struct {
	v *url.URL
}

func NewStrictUrlFromString(v string) (StrictUrl, error) {
	u, err := url.Parse(v)
	if err != nil {
		return StrictUrl{}, err
	}
	if !u.IsAbs() {
		return StrictUrl{}, fmt.Errorf("URL must be absolute: %s", u.String())
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return StrictUrl{}, fmt.Errorf("URL scheme must be http or https: %s", u.Scheme)
	}
	if u.Host == "" {
		return StrictUrl{}, fmt.Errorf("URL host must not be empty")
	}
	return StrictUrl{u}, nil
}

func (u StrictUrl) String() string {
	return u.v.String()
}
