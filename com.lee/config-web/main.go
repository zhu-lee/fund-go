package main

import (
	"com.lee/config-web/controller"
	"com.lee/fund/web"
)

func main() {
	app := web.NewWebApp()
	app.RegisterController("/rpcserver", &controller.RpcServerController{})
	app.RegisterController("/", &controller.DefaultController{})
	app.Start()
}
