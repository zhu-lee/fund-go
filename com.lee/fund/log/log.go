package log

import (
	"com.lee/fund/config"
	"github.com/astaxie/beego/logs"
	"strings"
)

var logLevelMap = map[string]int{
	"emergency": logs.LevelEmergency,
	"alert":     logs.LevelAlert,
	"critical":  logs.LevelCritical,
	"error":     logs.LevelError,
	"warning":   logs.LevelWarning,
	"notice":    logs.LevelNotice,
	"info":      logs.LevelInfo,
	"debug":     logs.LevelDebug,
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
