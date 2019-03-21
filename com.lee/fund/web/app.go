package web

import "com.lee/fund/config"

type App struct {
	Config *Config
	route  *routeEngine
	//viewEngin *ViewEngin
}

func NewDefaultApp() *App {
	w := config.GetAppConfig().WebSetting
	return &App{
		route:  newRoute(),
		Config: NewConfig(w),
	}
}

func (a *App) RegisterController(url string, controller interface{}) {
	a.route.RegisterController(url, controller)
}

func (a *App) Start() {

}