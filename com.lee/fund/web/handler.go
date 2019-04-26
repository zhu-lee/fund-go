package web

import (
	"com.lee/fund/log"
	"fmt"
	"net/http"
	"reflect"
	"runtime/debug"
)

type handler interface {
	Process(ctx *Context)
}

type ErrorHandler func(err interface{}, w http.ResponseWriter, r *http.Request)

type controllerHandler struct {
	controllerType reflect.Type
	methodName     string
}

func (h *controllerHandler) Process(ctx *Context) {
	controller := reflect.New(h.controllerType)
	method := controller.MethodByName(h.methodName)
	args := make([]reflect.Value, 1)
	args[0] = reflect.ValueOf(ctx)
	method.Call(args)
}

func errHandler(err interface{}, w http.ResponseWriter, r *http.Request) {
	errInfo := fmt.Sprint(err)
	log.Log.Critical("web", errInfo)

	errInfo += "\n"+string(debug.Stack())
	http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
}
