package entity

import "time"

type Task struct {
	Name           string     `bson:"_id" json:"name"`
	Alias          string     `bson:"alias" json:"alias"`
	Note           string     `bson:"note" json:"note"`
	Executor       *Executor  `bson:"executor" json:"executor"`
	Triggers       []*Trigger `bson:"triggers" json:"triggers"`
	Args           []*Arg     `bson:"args" json:"args"`
	Warning        *Warning   `bson:"warning" json:"warning"`
	Status         int        `bson:"status" json:"status"`
	CreateUser     int32      `bson:"createUser" json:"createUser"`
	CreateTime     time.Time  `bson:"createTime" json:"createTime"`
	UpdateUser     int32      `bson:"updateUser" json:"updateUser"`
	UpdateTime     time.Time  `bson:"updateTime" json:"updateTime"`
	Type           int        `bson:"type" json:"type"`
	UpdateUserName string     `bson:"-" json:"-"`
}

type Executor struct {
	Name string `bson:"name" json:"name"`
	Type string `bson:"type" json:"type"`
}

type Trigger struct {
	Value string    `bson:"value" json:"value"`
	Start time.Time `bson:"start" json:"start"`
	End   time.Time `bson:"end" json:"end"`
}

type Arg struct {
	Name  string `bson:"name" json:"name"`
	Value string `bson:"value" json:"value"`
}

type Warning struct {
	Emails  []string `bson:"emails" json:"emails"`
	Mobiles []string `bson:"mobiles" json:"mobiles"`
}
