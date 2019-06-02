package client

import (
	"github.com/smith-30/acc/domain"
	"github.com/smith-30/acc/infra/client/atcoder"
)

func NewClient(cn domain.ContestName) domain.Client {
	switch cn {
	case domain.Atcoder:
		return atcoder.NewClient()
	}
	return nil
}
