package web

import "net/http"

type Response struct {
	app *App
	rw  http.ResponseWriter
}