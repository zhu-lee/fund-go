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
	Anonymous      bool
	ResourceFolder string
	StaticFolder   string
	ViewFolder     string
	CookieName     string
	CookieDomain   string
	CookieTime     time.Duration
	UrlSignIn      string
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
		Anonymous:      settings.Bool("anonymous", true),
		CookieName:     settings.String("cookie_name", "_u"),
		CookieDomain:   settings.String("cookie_domain", ""),
		CookieTime:     time.Duration(settings.Int("cookie_time", 30)) * time.Minute,
		ResourceFolder: filepath.Join(folder, "resources"),
		UrlSignIn:      settings.String("url_signIn", "/signIn"),
	}

	cfg.StaticFolder = filepath.Join(cfg.ResourceFolder, "static")
	cfg.ViewFolder = filepath.Join(cfg.ResourceFolder, "view")
	if cfg.Version == "" {
		cfg.Version = time.Now().Format("20190511.152601") //todo 无构建信息，先给定当前时间version
	}

	//TODO 设置日志文件

	return cfg
}
