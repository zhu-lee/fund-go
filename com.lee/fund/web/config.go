package web

import (
	"com.lee/fund/config"
	"path/filepath"
	"time"
)

type Config struct {
	AppName        string
	Version        string
	HttpAddr       string
	HttpPort       int
	ResourceFolder string
	StaticFolder   string
	ViewFolder     string
}

func NewConfig(settings config.SettingMap) *Config {
	folder, err := config.GetAppFolder()
	if err != nil {
		panic(err)
	}

	cfg := &Config{
		AppName:        config.GetAppConf().App.AppName,
		Version:        config.GetAppConf().App.Version,
		HttpAddr:       settings.String("http_addr", ""),
		HttpPort:       settings.Int("http_port", 80),
		ResourceFolder: filepath.Join(folder, "resources"),
	}
	cfg.StaticFolder = filepath.Join(cfg.ResourceFolder, "static")
	cfg.ViewFolder = filepath.Join(cfg.ResourceFolder, "view")
	if cfg.Version == "" {
		//todo 构建信息的version
		cfg.Version = time.Now().Format("20190511.152601")
	}

	//TODO 设置日志文件
	
	return cfg
}
