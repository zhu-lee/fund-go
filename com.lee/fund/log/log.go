package log

import (
	"com.lee/fund/config"
	"github.com/astaxie/beego/logs"
	"strings"
)

var logLevelMap = map[string]int{
	"emergency": logs.LevelEmergency, //紧急级别
	"alert":     logs.LevelAlert,     //报警级别
	"critical":  logs.LevelCritical,  //严重错误级别
	"error":     logs.LevelError,     //错误级别
	"warning":   logs.LevelWarning,   //警告级别
	"notice":    logs.LevelNotice,    //注意级别
	"info":      logs.LevelInfo,      //报告级别
	"debug":     logs.LevelDebug,     //出错级别
}

var Log *logs.BeeLogger

/**
 * 初始化日志工具
 */
func init() {
	appConf := config.GetAppConf()
	level := logLevelMap[strings.ToLower(appConf.App.LogLevel)]
	logger := appConf.App.LogProvider

	Log = logs.NewLogger()        // 创建一个日志记录器
	_ = Log.SetLogger(logger, "") // 输出到控制台
	Log.SetLevel(level)           // 设置日志级别
	Log.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）
	Log.SetLogFuncCallDepth(2)
	//TODO 扩展 json 日志
}
