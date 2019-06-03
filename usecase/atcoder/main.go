package atcoder

import (
	"github.com/smith-30/acc/domain"
	"github.com/smith-30/acc/usecase"
)

type usecase struct{}

func NewUsecase() usecase.Usecase {
	return &usecase{}
}

func (a *usecase) Exec(c usecase.Cmd) ([]domain.TestCase, error) {
	return nil, nil
}
