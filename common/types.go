package common

import (
	"io"
	"net/http"
)

//ClientHTTPGetter Simple interface to net/http Get()
type ClientHTTPGetter interface {
	Get(string) (*http.Response, error)
}

//ClientHTTPPoster Simple interface to net/http Post()
type ClientHTTPPoster interface {
	Post(string, string, io.Reader) (*http.Response, error)
}
