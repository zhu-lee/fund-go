package main

import (
	"com.lee/config-web/controller"
	"com.lee/fund/web"
)

func main() {
	app := web.NewWebApp()
	app.RegisterController("/team", &controller.DevTeamController{})
	app.RegisterController("/", &controller.DefaultController{})
	app.Start()
}