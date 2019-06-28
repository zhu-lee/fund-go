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

type RouterItem struct {
	handler   handler
	Name      string
	Anonymous bool
}

type routerEngine struct {
	routes map[string]*RouterItem
	cfg    *Config
}

type RouterOption struct {
	Anonymous bool
}

func newRoute(cfg *Config) *routerEngine {
	return &routerEngine{
		routes: make(map[string]*RouterItem),
		cfg:    cfg,
	}
}

func (r *routerEngine) Router(url string, controller interface{}, options ...map[string]*RouterOption) {
	vt := reflect.Indirect(reflect.ValueOf(controller)).Type()
	tt := reflect.TypeOf(controller)
	controllerName := strings.TrimSuffix(vt.Name(), "Controller")

	for i := 0; i < tt.NumMethod(); i++ {
		m := tt.Method(i)
		h := &controllerHandler{
			controllerType: vt,
			methodName:     m.Name,
		}

		var option *RouterOption
		if len(options) > 0 {
			option = options[0][m.Name]
		}

		u := path.Join(url, m.Name)
		r.registerRoute(ROUTE_CONTROLLER, controllerName+"/"+m.Name, u, h, option)
	}
}

func (r *routerEngine) registerRoute(rt routeType, name, url string, h handler, option *RouterOption) {
	item := &RouterItem{
		handler: h,
		Name:    name,
	}

	if option == nil {
		item.Anonymous = r.cfg.Anonymous
	} else {
		item.Anonymous = option.Anonymous
	}

	r.routes[strings.ToLower(url)] = item
}

func (r *routerEngine) GetRouter(urlPath string) *RouterItem {
	urlPath = strings.ToLower(urlPath)
	if r, ok := r.routes[urlPath]; ok {
		return r
	}
	return nil
}
