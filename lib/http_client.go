package lib

import (
	"io"
	"net/http"
)

//counterfeiter:generate . HTTPClient
type HTTPClient interface {
	Get(url string) (*http.Response, error)
	Post(url, contentType string, content io.Reader) (*http.Response, error)
}
