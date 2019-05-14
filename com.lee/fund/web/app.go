package web

import (
	"com.lee/fund/config"
	"com.lee/fund/log"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	F_NONE = iota
	F_DIR
	F_FILE
)

type App struct {
	Config     *Config
	route      *routeEngine
	errHandler ErrorHandler
	ve         *ViewEngine
}

func NewWebApp() *App {
	webSetting := config.GetAppConf().Web
	cfg := NewConfig(webSetting)
	ve := &ViewEngine{
		suffix: ".html",
		dir:    cfg.ViewFolder,
	}
	return &App{
		Config:     cfg,
		route:      newRoute(),
		errHandler: errHandler,
		ve:         ve,
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

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		reqUrl *url.URL
	)

	defer func() {
		//捕获异常 panic用于抛出异常, recover用于捕获异常
		if err := recover(); err != nil {
			a.errHandler(err, w, r)
		}

		//TODO 如果成功，输出web访问日志到指路径的日志文件
	}()

	reqUrl = r.URL

	//默认路由到首页
	if reqUrl.Path == "/" {
		reqUrl.Path = "/index"
	}

	//获取注册路由
	route := a.route.GetRoute(reqUrl.Path)

	//如果静态页，路由到哪儿
	if route == nil {
		a.handleStatic(w, r, reqUrl)
	} else {
		ctx := a.newContext(w, r, reqUrl)
		a.handleRoute(ctx, w, r, route)
	}
}

func (a *App) handleStatic(w http.ResponseWriter, r *http.Request, url *url.URL) {
	//禁止包令相对路径
	if strings.Contains(url.Path, "/../") {
		http.Error(w, "禁止包令相对路径："+url.Path, http.StatusForbidden)
		return
	}

	//检查合法性
	filePath := filepath.Join(a.Config.StaticFolder, url.Path)
	f := getFileType(filePath)
	if f == F_NONE {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if f == F_DIR {
		filePath = filepath.Join(filePath, "index.html")
		f = getFileType(filePath)
		if f == F_NONE {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
	}

	http.ServeFile(w, r, filePath)
}

func getFileType(filepath string) int32 {
	f, err := os.Stat(filepath)
	if err != nil {
		return F_NONE
	} else if f.IsDir() {
		return F_DIR
	} else {
		return F_FILE
	}
}

func (a *App) newContext(w http.ResponseWriter, r *http.Request, url *url.URL) *Context {
	return &Context{
		App:      a,
		Request:  r,
		Response: w,
		Url:      url,
	}
}

func (a *App) handleRoute(ctx *Context, w http.ResponseWriter, r *http.Request, route *RouteItem) {
	//TODO 设置权限

	//TODO 过滤请求

	route.handler.Process(ctx)
}
