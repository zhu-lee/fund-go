package web

import (
	"path"
	"reflect"
	"strings"
)

type routeType int32

const (
	ROUTE_CONTROLLER routeType = iota
	ROUTE_PAGE
	ROUTE_ACTION
)

type RouteItem struct {
	handler handler
	Name    string
}

type routeEngine struct {
	routes map[string]*RouteItem
}

type HandlerOption struct {
	Suffix    string
	Anonymous bool
}

func newRoute() *routeEngine {
	return &routeEngine{
		routes: make(map[string]*RouteItem),
	}
}

func (r *routeEngine) RegisterController(url string, controller interface{}) {
	r.RegisterControllerH(url, controller, nil)
}

func (r *routeEngine) RegisterControllerH(url string, controller interface{}, options map[string]*HandlerOption) {
	vt := reflect.Indirect(reflect.ValueOf(controller)).Type()
	tt := reflect.TypeOf(controller)
	controllerName := strings.TrimSuffix(vt.Name(), "Controller")

	for i := 0; i < tt.NumMethod(); i++ {
		m := tt.Method(i)
		h := &controllerHandler{
			controllerType: vt,
			methodName:     m.Name,
		}

		var option *HandlerOption
		if options != nil {
			option = options[m.Name]
		}

		u := path.Join(url, m.Name)
		r.registerRoute(ROUTE_CONTROLLER, controllerName+"/"+m.Name, u, h, option)
	}
}

func (r *routeEngine) registerRoute(rt routeType, name, url string, h handler, option *HandlerOption) {
	routeItem := &RouteItem{
		handler:h,
		Name: name,
		}
	r.routes[strings.ToLower(url)]= routeItem
}

func (r *routeEngine) GetRoute(urlPath string) *RouteItem {
	urlPath = strings.ToLower(urlPath)
	if r, ok := r.routes[urlPath]; ok {
		return r
	}
	return nil
}
