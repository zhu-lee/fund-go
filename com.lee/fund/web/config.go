package web

import (
	"com.lee/fund/config"
)

type Config struct {
	AppName  string
	Version  string
	HttpAddr string
	HttpPort int
}

func NewConfig(settings config.SettingMap) *Config {
	cfg := &Config{
		AppName:  config.GetAppConf().App.AppName,
		Version:  config.GetAppConf().App.Version,
		HttpAddr: settings.String("http_addr", ""),
		HttpPort: settings.Int("http_port", 80),
	}
	return cfg
}
