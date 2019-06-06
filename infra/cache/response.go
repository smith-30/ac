package cache

import (
	"github.com/recoilme/pudge"
	"github.com/smith-30/acc/domain"
)

type httpResponseRepo struct{}

func NewCacheRepo() (domain.HttpResponseRepository, err) {
	cfg := &pudge.Config{
		SyncInterval: 0} //disable every second fsync
	db, err := pudge.Open("./acc.db", cfg)
	if err != nil {
		return nil, err
	}
	return &httpResponseRepo{}, nil
}

func (a *httpResponseRepo) Get(k domain.Key) *domain.HttpResponse {
	return nil
}
