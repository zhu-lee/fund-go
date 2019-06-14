package controller

import (
	"com.lee/config-web/model"
	"com.lee/config-web/resp"
	"com.lee/fund/web"
)

type TaskController struct {
}

func (c *TaskController) List(ctx *web.Context) {
	m := struct {
		*model.BaseModel
	}{
		BaseModel: model.NewBaseModel(ctx),
	}
	resp.WritePage(ctx, "task/list", m)
}

func (c *TaskController) Logs(ctx *web.Context) {
	m := struct {
		*model.BaseModel
	}{
		BaseModel: model.NewBaseModel(ctx),
	}
	resp.WritePage(ctx,"task/log",m)
}
