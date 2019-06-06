package http

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/smith-30/acc/domain"
)

type httpClient struct{}

func NewHttpClient() domain.HttpClient {
	return &httpClient{}
}

func (a *httpClient) Get(u string) (*domain.HttpResponse, error) {
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &domain.HttpResponse{
		Body: b,
	}, nil
}
