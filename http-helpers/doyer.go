package http_helpers

import "net/http"

type HTTPDoyer interface {
	Do(req *http.Request) (*http.Response, error)
}
