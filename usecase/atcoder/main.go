package atcoder

import (
	"github.com/smith-30/acc/domain"
	"github.com/smith-30/acc/usecase"
)

type uc struct{}

func NewUsecase() usecase.Usecase {
	return &uc{}
}

func (a *uc) Exec(c usecase.Cmd) ([]domain.TestCase, error) {
	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
	  log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
	  log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return nil, nil
}
