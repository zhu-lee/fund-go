package main

import (
	"com.lee/config-web/controller"
	"com.lee/fund/log"
	"com.lee/fund/web"
)

func main() {
	app := web.NewDefaultApp()
	app.RegisterController("/team", &controller.DevTeamController{})
	app.Start()
	log.Log().Info("starting config...")
}