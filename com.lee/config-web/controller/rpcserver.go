package controller

import (
	"com.lee/config-web/model"
	"com.lee/config-web/resp"
	"com.lee/fund/web"
)

type RpcServerController struct {
}

func (c *RpcServerController) List(ctx *web.Context) {
	m := struct {
		*model.BaseModel
	}{
		BaseModel: model.NewBaseModel(ctx),
	}

	m.AddCss(ctx, "")
	m.AddJs(ctx, "")
	resp.WritePage(ctx,"rpcserver/list", m)
}
