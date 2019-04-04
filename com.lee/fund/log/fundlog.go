package log

import (
	"github.com/astaxie/beego/logs"
)

type fundLog struct {
	blg *logs.BeeLogger
}

var flg *fundLog

func init() {
	//根据配置文件设置log
	log := logs.NewLogger()       // 创建一个日志记录器，参数为缓冲区的大小
	log.SetLogger("console", "")  // 设置日志记录方式：控制台记录
	log.SetLevel(logs.LevelDebug) // 设置日志写入缓冲区的等级：Debug级别（最低级别，所以所有log都会输入到缓冲区）
	log.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）
	log.SetLogFuncCallDepth(3)
	flg = &fundLog{blg: log}
	//TODO 扩展 json
}

func Log() *fundLog {
	return flg
}

func (fl *fundLog) Emergency(msg string, args ...interface{}) {
	if args != nil {
		flg.blg.Emergency(msg, args)
	} else {
		flg.blg.Emergency(msg)
	}
}
func (fl *fundLog) Alert(msg string, args ...interface{}) {
	if args != nil {
		flg.blg.Alert(msg, args)
	} else {
		flg.blg.Alert(msg)
	}
}
func (fl *fundLog) Critical(msg string, args ...interface{}) {
	if args != nil {
		flg.blg.Critical(msg, args)
	} else {
		flg.blg.Critical(msg)
	}
}
func (fl *fundLog) Error(msg string, args ...interface{}) {
	if args != nil {
		flg.blg.Error(msg, args)
	} else {
		flg.blg.Error(msg)
	}
}
func (fl *fundLog) Warning(msg string, args ...interface{}) {
	if args != nil {
		flg.blg.Warning(msg, args)
	} else {
		flg.blg.Warning(msg)
	}
}
func (fl *fundLog) Notice(msg string, args ...interface{}) {
	if args != nil {
		flg.blg.Notice(msg, args)
	} else {
		flg.blg.Notice(msg)
	}
}
func (fl *fundLog) Info(msg string, args ...interface{}) {
	if args != nil {
		flg.blg.Informational(msg, args)
	} else {
		flg.blg.Informational(msg)
	}
}
func (fl *fundLog) Debug(msg string, args ...interface{}) {
	if args != nil {
		flg.blg.Debug(msg, args)
	} else {
		flg.blg.Debug(msg)
	}
}
