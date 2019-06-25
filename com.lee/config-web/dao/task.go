package dao

import (
	"com.lee/config-web/entity"
	"com.lee/fund/mgo/bson"
	"com.lee/fund/mgo/mgox"
)

// 数据库
const (
	dbMongoConfig   = "config"
	dbMongoSchedule = "schedule"
	dbMongoAppLog   = "app_log"
)

var TaskDao = new(taskDao)

type taskDao struct {
}

func (dao *taskDao) GetTasks(name, executor string, pageIndex, pageSize int) (tasks []*entity.Task, totalCount int, err error) {
	db := mgox.MustOpen(dbMongoSchedule)
	defer db.Close()

	query := bson.M{}
	if name != "" {
		query["_id"] = bson.M{"$regex": name, "$options": "i"}
	}
	if executor != "" {
		query["executor.name"] = executor
	}
	q := db.Coll("Task").Find(query)

	totalCount, err = q.Count()
	if err != nil {
		return
	}

	tasks = []*entity.Task{}
	err = q.Skip(pageSize * (pageIndex - 1)).Limit(pageSize).All(&tasks)
	return
}
