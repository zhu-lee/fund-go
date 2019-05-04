package controller

import (
	"com.lee/fund/web"
	"fmt"
)

type DefaultController struct {
	prefix string "aa"
}

func (c *DefaultController) Index (ctx *web.Context) {
	fmt.Println("DefaultController-Index")
	ctx.Response.Write([]byte("hello,index！"))
}
