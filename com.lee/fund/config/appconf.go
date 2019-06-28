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

func newApp(settingMap SettingMap) *App {
	app := &App{}
	if settingMap != nil {
		app.AppName = settingMap.String("app_name", "")
		app.Debug = settingMap.Bool("debug", false)
		app.Version = settingMap.String("version", "")
		app.LogPath = settingMap.String("log_path", "")
		app.LogLevel = settingMap.String("log_level", "debug")
		app.LogProvider = settingMap.String("log_provider", "console")
		app.LogFormat = settingMap.String("log_format", "")
		app.GlobalEnv = settingMap.String("global_env", "")
	}
	if app.AppName == "" {
		app.AppName = filepath.Base(os.Args[0])
	}
	return app
}
