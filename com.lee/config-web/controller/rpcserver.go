package controller

import (
	"com.lee/config-web/biz"
	"com.lee/config-web/model"
	"com.lee/config-web/resp"
	"com.lee/fund/web"
)

type RpcServerController struct {
}

func (c *RpcServerController) List(ctx *web.Context) {
	name := ctx.GetParamString("name")
	pageIndex := ctx.GetParamInt("page", 1)
	pageSize := model.DefaultPageSize

	_, _, _ = biz.RpcServerBiz.GetRpcServices(name, pageIndex, pageSize)

	m := struct {
		*model.BaseModel
	}{
		BaseModel: model.NewBaseModel(ctx),
	}

	//m.AddCss(ctx, "")
	//m.AddJs(ctx, "")
	resp.WritePage(ctx, "rpcserver/list", m)
}
