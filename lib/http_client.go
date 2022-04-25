package lib

import "net/http"

//counterfeiter:generate . HTTPClient
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}
