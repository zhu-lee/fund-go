package model

import (
	"com.lee/config-web/biz"
	"com.lee/config-web/entity"
	"com.lee/fund/web"
)

const (
	DefaultPageSize = 25
)

type BaseModel struct {
	*web.SM
	*web.PM
	*entity.MenuInfo
}

func NewBaseModel(ctx *web.Context) *BaseModel {
	bm := BaseModel{
		SM: web.NewSM(ctx),
		PM: new(web.PM),
	}

	bm.AddCss(ctx, "")
	bm.AddJs(ctx, "")
	bm.MenuInfo = biz.Menu.GetMenuInfo(ctx.Request.URL.Path)

	return &bm
}
