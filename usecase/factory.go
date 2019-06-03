package usecase

import (
	"github.com/smith-30/acc/domain"
)

type Usecase interface {
	Exec(c Cmd) ([]domain.TestCase, error)
}

type Cmd interface {
	Key() domain.Key
}
