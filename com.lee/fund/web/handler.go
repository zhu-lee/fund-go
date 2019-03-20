package web

import "reflect"

type handler interface {
	Process(ctx *Context)
}

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
