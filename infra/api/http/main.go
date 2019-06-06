package http

import "github.com/smith-30/acc/domain"

type httpClient struct{}

func (a *httpClient) Get(u string) (*domain.HttpResponse, error) {
	return &domain.HttpResponse{}, nil
}
