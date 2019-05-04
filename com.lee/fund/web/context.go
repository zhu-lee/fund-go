package web

import (
	"net/http"
	"net/url"
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	Url      *url.URL
	App      *App
	data     map[interface{}]interface{}
}
