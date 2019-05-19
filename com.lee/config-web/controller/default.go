package controller

import (
	"com.lee/config-web/model"
	"com.lee/config-web/resp"
	"com.lee/fund/web"
)

type DefaultController struct {
}

func (c *DefaultController) Index (ctx *web.Context) {
	m:= struct {
		*model.BaseModel
		Version  string   //todo
		Profiles []string //todo
	}{
		BaseModel: model.NewBaseModel(ctx),
	}
	m.AddCss(ctx, "")
	m.AddJs(ctx, "")
	resp.WritePage(ctx,"default/index",m)
}
