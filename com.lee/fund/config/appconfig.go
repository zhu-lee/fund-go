package config

import (
	"com.lee/fund/xml"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type appConfig struct {
	WebSetting *WebSetting
	AppSetting *AppSetting
	Custom     SettingMap
}

type WebSetting struct {
	HttpAddr     string
	HttpPort     int
	LogEnable    bool
	Anonymous    bool
	CookieName   string
	CookieDomain string
	CookieTime   time.Duration
	UrlDefault   string
	UrlSignIn    string
	UrlNotFound  string
	UrlNoRight   string
}

type AppSetting struct {
	AppName   string
	Debug     bool
	LogLevel  int
	GlobalEnv string
}

func (a *appConfig) init() {
	a.Custom = SettingMap{}
	a.AppSetting = newAppSetting(nil)
	a.WebSetting = newWebSetting(nil)
}

func (c *appConfig) load(doc *xml.Document) error {
	sections := make(map[string]SettingMap)
	configNode := doc.SelectNode("", "config")
	for _, sectionNode := range configNode.Children {
		settings := SettingMap{}
		for _, settingNode := range sectionNode.Children {
			if os := settingNode.As("", "os"); os == "" || os == runtime.GOOS {
				settings[settingNode.As("", "key")] = settingNode.As("", "value")
			}
		}
		sections[sectionNode.Name.Local] = settings
	}

	c.AppSetting = newAppSetting(sections["app"])
	c.WebSetting = newWebSetting(sections["web"])
	c.Custom = sections["custom"]
	if c.Custom == nil {
		c.Custom = SettingMap{}
	}
	return nil
}

func newAppSetting(m SettingMap) *AppSetting {
	s := &AppSetting{}
	if m != nil {
		s.AppName = m.String("app_name", "")
		s.Debug = m.Bool("debug", false)
		s.LogLevel = m.Int("log_level", 0)
		s.GlobalEnv = m.String("global_env", "")
	}
	if s.AppName == "" {
		s.AppName = filepath.Base(os.Args[0])
	}
	return s
}

func newWebSetting(settings SettingMap) *WebSetting {
	w := &WebSetting{}
	if settings != nil {
		w.HttpAddr = settings.String("http_addr", "")
		w.HttpPort = settings.Int("http_port", 80)
		w.LogEnable = settings.Bool("log_enable", true)
		w.Anonymous = settings.Bool("anonymous", true)
		w.CookieName = settings.String("cookie_name", "_u")
		w.CookieDomain = settings.String("cookie_domain", "")
		w.CookieTime = time.Duration(settings.Int("cookie_time", 30)) * time.Minute
		w.UrlDefault = settings.String("url_default", "/")
		w.UrlSignIn = settings.String("url_signin", "/signin")
		w.UrlNotFound = settings.String("url_notfound", "/404.html")
		w.UrlNoRight = settings.String("url_noright", "/noright.html")
	}
	return w
}
