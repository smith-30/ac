package domain

type Client interface {
	GetTestCase(url string) ([]TestCase, error)
}

type TestCase struct {
	Content string
	Exp     string
}
