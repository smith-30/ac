package domain

import "net/url"

type Key string

type CacheKey struct {
	URL *url.URL
}

func NewCacheKey(ru string) (*CacheKey, error) {
	u, err := url.ParseRequestURI(ru)
	if err != nil {
		return nil, err
	}
	return &CacheKey{
		URL: u,
	}, nil
}

func (a *CacheKey) Key() Key {
	return Key(a.URL.Hostname() + a.URL.Path)
}

func (a *CacheKey) RequestDest() string {
	return a.URL.String()
}
