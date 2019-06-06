package atcoder

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
	"github.com/smith-30/acc/domain"
)

type AtcoderUC interface {
	Exec(c Command) ([]domain.TestCase, error)
}

type Command struct {
	*domain.CacheKey
}

type uc struct {
	repo    domain.HttpResponseRepository
	httpCli domain.HttpClient
}

func NewUsecase(r domain.HttpResponseRepository, hc domain.HttpClient) AtcoderUC {
	return &uc{
		repo:    r,
		httpCli: hc,
	}
}

func (a *uc) Exec(c Command) ([]domain.TestCase, error) {
	var r *domain.HttpResponse
	// get cache by key
	r = a.repo.Get(c.Key())
	// no cache
	if r == nil {
		res, err := a.httpCli.Get(c.RequestDest())
		if err != nil {
			return nil, err
		}
		err = a.repo.Save(res)
		if err != nil {
			return nil, err
		}
		r = res
	}

	// Request the HTML page.
	// res, err := http.Get("http://metalsucks.net")
	// if err != nil {
	//   log.Fatal(err)
	// }
	// defer res.Body.Close()
	// if res.StatusCode != 200 {
	//   log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	// }
	reader := bytes.NewBuffer(r.Body)
	doc, err := goquery.NewDocumentFromReader(reader)
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
