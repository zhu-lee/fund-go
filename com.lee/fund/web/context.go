package web

import "net/http"

type Context struct {
	RawRequest  *http.Request
	RawResponse http.ResponseWriter
	App         *App
	data        map[interface{}]interface{}
}
