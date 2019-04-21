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
}

func NewWebApp() *App {
	web := config.GetAppConf().Web
	return &App{
		route:  newRoute(),
		Config: NewConfig(web),
	}
}

func (a *App) RegisterController(url string, controller interface{}) {
	a.route.RegisterController(url, controller)
}

func (a *App) Start() {
	addr := fmt.Sprintf("%s:%d", a.Config.HttpAddr, a.Config.HttpPort)
	log.Log.Info("web server [%s] startup on %s", a.Config.AppName, addr)

	h := &http.Server{
		Addr:    addr,
		Handler: a,
	}

	err := h.ListenAndServe()
	if err != nil {
		log.Log.Error("web server [%s] startup failed：%v", a.Config.AppName, err)
		panic(err)
	}

	log.Log.Info("web server [%s] terminated", a.Config.AppName)
}

func (a *App) ServeHTTP(rpw http.ResponseWriter, req *http.Request) {
	fmt.Println("1=========")
}