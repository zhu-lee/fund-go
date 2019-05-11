package model

import (
	"com.lee/config-web/entity"
	"com.lee/fund/web"
)

type BaseModel struct {
	*web.SM
	*web.PM
	entity.MenuInfo
}

func NewBaseModel(ctx *web.Context) *BaseModel {
	bm := BaseModel{
		SM:web.NewSM(ctx),
		PM:new(web.PM),
	}

	bm.AddCss(ctx, "")
	bm.AddJs(ctx,"")
	bm.MenuInfo =

	return &bm
}