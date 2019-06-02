package atcoder

type AtcoderClient interface {
	GetTestCase() ([]TestCase, error)
}

type TestCase struct {
	Exp string
}
