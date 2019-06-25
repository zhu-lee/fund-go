package biz

import (
	"com.lee/config-web/dao"
	"com.lee/config-web/entity"
)

var TaskBiz = new(taskBiz)

type taskBiz struct {
}

func (biz *taskBiz) GetTasks(name, executor string, pageIndex, pageSize int) (tasks []*entity.Task, totalCount int, err error) {
	return dao.TaskDao.GetTasks(name, executor, pageIndex, pageSize)
}