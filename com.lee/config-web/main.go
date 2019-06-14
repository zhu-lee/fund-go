package main

import (
	"com.lee/config-web/controller"
	"com.lee/fund/web"
)

func main() {
	app := web.NewWebApp()
	app.RegisterController("/", &controller.DefaultController{})
	app.RegisterController("/rpc", &controller.RpcServerController{})
	app.RegisterController("/task", &controller.TaskController{})
	app.Start()
}
