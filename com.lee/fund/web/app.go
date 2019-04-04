package web

import (
	"com.lee/fund/config"
	"com.lee/fund/log"
	"fmt"
	"net/http"
)

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
	addr := fmt.Sprintf("%s:%d",a.Config.HttpAddr,a.Config.HttpPort)
	log.Log().Info("web server ready to run on %s",addr)

	h := &http.Server{
		Addr:addr,
		Handler:a,
	}

	err := h.ListenAndServe()
	if err != nil {
		log.Log().Error("web server startup failed：%v",err)
	}

	log.Log().Info("web server terminated")
}

func (app *App)ServeHTTP(rpw http.ResponseWriter, req *http.Request) {
	fmt.Println("1=========")
}