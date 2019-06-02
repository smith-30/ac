package atcoder

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/smith-30/acc/domain"
)

type client struct {
}

func NewClient() domain.Client {
	return &client{}
}

func (a *client) GetTestCase(u string) ([]domain.TestCase, error) {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocument(u)
	if err != nil {
		return nil, err
	}

	// tcs := make([]domain.TestCase, 0, 10)
	cache := map[string]string{}
	contents := make([]string, 0, 5)
	exps := make([]string, 0, 5)
	hit := 0

	doc.Find(".part > section > pre").Each(func(i int, s *goquery.Selection) {
		if v := s.Has("var").Text(); v == "" {
			if _, ok := cache[s.Text()]; !ok {
				val := s.Text()
				cache[val] = val
				switch hit % 2 {
				case 0:
					contents = append(contents, val)
				case 1:
					exps = append(exps, val)
				}
				hit++
			}
		}
	})

	tcs := make([]domain.TestCase, 0, len(contents))

	for index := 0; index < len(contents); index++ {
		tcs = append(tcs, domain.TestCase{Content: contents[index], Exp: exps[index]})
	}

	return tcs, nil
}
