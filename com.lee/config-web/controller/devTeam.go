package controller

import "fmt"

type DevTeamController struct {
}

func (c *DevTeamController) List () {
	fmt.Println("DevTeamController-List")
}

func (c *DevTeamController) Add () {
	fmt.Println("DevTeamController-Add")
}

func (c *DevTeamController) Del () {
	fmt.Println("DevTeamController-Del")
}