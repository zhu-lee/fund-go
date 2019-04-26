package controller

import "fmt"

type DefaultController struct {
}

func (c *DefaultController) Index () {
	fmt.Println("DefaultController-Index")
}

