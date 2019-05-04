package controller

import (
	"com.lee/fund/web"
	"fmt"
)

type DevTeamController struct {
}

func (c *DevTeamController) List (ctx *web.Context) {
	fmt.Println("DevTeamController-List")
}

func (c *DevTeamController) Add (ctx *web.Context) {
	fmt.Println("DevTeamController-Add")
}

func (c *DevTeamController) Del (ctx *web.Context) {
	fmt.Println("DevTeamController-Del")
}