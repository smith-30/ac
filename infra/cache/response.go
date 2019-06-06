package cache

import (
	"go/build"

	"github.com/recoilme/pudge"
	"github.com/smith-30/acc/domain"
)

type httpResponseRepo struct {
	db *pudge.Db
}

func NewCacheRepo() (domain.HttpResponseRepository, error) {
	cfg := &pudge.Config{
		SyncInterval: 0, //disable every second fsync
	}
	db, err := pudge.Open(build.Default.GOPATH+"/bin/acc.db", cfg)
	if err != nil {
		return nil, err
	}
	return &httpResponseRepo{
		db: db,
	}, nil
}

func (a *httpResponseRepo) Get(k domain.Key, value interface{}) (*domain.HttpResponse, error) {
	err := a.db.Get(k.String(), value)
	return value.(*domain.HttpResponse), err
}

func (a *httpResponseRepo) Save(k domain.Key, value interface{}) error {
	err := a.db.Set(k.String(), value)
	return err
}
