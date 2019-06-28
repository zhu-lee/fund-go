package main

import (
	"com.lee/config-web/controller"
	"com.lee/fund/web"
)

func main() {
	app := web.NewWebApp()
	app.SetAuth()

	app.Router("/", &controller.DefaultController{}, map[string]*web.RouterOption{"SignIn": {Anonymous: true}})
	app.Router("/rpc", &controller.RpcServerController{})
	app.Router("/task", &controller.TaskController{})
	app.Start()
}
