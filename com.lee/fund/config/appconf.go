package config

import (
	"com.lee/fund/xml"
	"os"
	"path/filepath"
	"runtime"
)

type appConf struct {
	App    *App
	Custom SettingMap
	Web    SettingMap
}

type App struct {
	AppName     string
	Debug       bool //TODO ?
	Version     string
	LogPath     string
	LogLevel    string
	LogProvider string
	LogFormat   string
	GlobalEnv   string //TODO ?
}

func (c *appConf) init() {
	c.App = newApp(nil)
	c.Custom = SettingMap{}
	c.Web = SettingMap{}
}

func (c *appConf) setAppCof(doc *xml.Document) error {
	sectionMap := make(map[string]SettingMap)
	configNode := doc.SelectNode("", "config")
	for _, sectionNode := range configNode.Children {
		settings := SettingMap{}
		for _, settingNode := range sectionNode.Children {
			if os := settingNode.As("", "os"); os == "" || os == runtime.GOOS {
				settings[settingNode.As("", "key")] = settingNode.As("", "value")
			}
		}
		sectionMap[sectionNode.Name.Local] = settings
	}
	c.App = newApp(sectionMap["app"]) //load App
	c.Web = sectionMap["web"]         //load Web
	c.Custom = sectionMap["custom"]   //load Custom
	if c.Web == nil {
		c.Web = SettingMap{}
	}
	if c.Custom == nil {
		c.Custom = SettingMap{}
	}
	return nil
}

func newApp(m SettingMap) *App {
	app := &App{}
	if m != nil {
		app.AppName = m.String("app_name", "")
		app.Debug = m.Bool("debug", false)
		app.Version = m.String("version", "")
		app.LogPath = m.String("log_path", "")
		app.LogLevel = m.String("log_level", "debug")
		app.LogProvider = m.String("log_provider", "console")
		app.LogFormat = m.String("log_format", "")
		app.GlobalEnv = m.String("global_env", "")
	}
	if app.AppName == "" {
		app.AppName = filepath.Base(os.Args[0])
	}
	return app
}