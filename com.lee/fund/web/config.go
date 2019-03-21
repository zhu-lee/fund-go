package web

import "com.lee/fund/config"

type Config struct {
	HttpAddr       string
	HttpPort       int
	ViewFolder     string // 视图文件目录
}

func NewConfig(w *config.WebSetting) *Config {
	return &Config{
		HttpAddr:w.HttpAddr,
		HttpPort:w.HttpPort,
	}
}
