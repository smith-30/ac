package domain

type HttpClient interface {
	Get(u string) (*HttpResponse, error)
}

type HttpResponseRepository interface {
	Get(k Key, value interface{}) (*HttpResponse, error)
	Save(k Key, value interface{}) error 
}

type HttpResponse struct {
	Body []byte
}
