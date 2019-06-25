package controller

import (
	"com.lee/config-web/biz"
	"com.lee/config-web/entity"
	"com.lee/config-web/model"
	"com.lee/config-web/resp"
	"com.lee/fund/web"
)

type TaskController struct {
}

func (c *TaskController) List(ctx *web.Context) {
	name := ctx.GetParamString("name")
	executor := ctx.GetParamString("executor")
	pageIndex := ctx.GetParamInt("pi", 1)
	pageSize := model.DefaultPageSize

	tasks, totalCount, err := biz.TaskBiz.GetTasks(name, executor, pageIndex, pageSize)
	if err != nil {
		resp.WriteString(ctx, "查询任务列表出错", err)
		return
	}

	m := struct {
		*model.PageModel
		*model.BaseModel
		Tasks []*entity.Task
	}{
		PageModel: model.NewPageModel(ctx, pageIndex, pageSize, totalCount),
		BaseModel: model.NewBaseModel(ctx),
		Tasks:     tasks,
	}
	resp.WritePage(ctx, "task/list", m)
}

func (c *TaskController) Logs(ctx *web.Context) {
	m := struct {
		*model.BaseModel
	}{
		BaseModel: model.NewBaseModel(ctx),
	}
	resp.WritePage(ctx, "task/log", m)
}
