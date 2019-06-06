package domain

type HttpClient interface {
	Get(u string) (*HttpResponse, error)
}

type HttpResponseRepository interface {
	Get(k Key) *HttpResponse
}

type HttpResponse struct {
	Body []byte
}
