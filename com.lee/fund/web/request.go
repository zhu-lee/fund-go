package web

import (
	"net/http"
	"net/url"
)

type Request struct {
	req        *http.Request
	Url        *url.URL
	rewrited   bool
	parameters url.Values
}
