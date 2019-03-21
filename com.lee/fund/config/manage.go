package config

import "sync"

const (
	ACTIVE_PROFILE_ENV = "FUND-PROFILE-ACTIVE"
)

var mg = newManage()

type manage struct {
	once sync.Once
	app *appConfig
}

func newManage() *manage {
	return &manage{
		app:new(appConfig),
	}
}

func GetAppConfig() *appConfig {
	mg.once.Do(mg.init)
	return mg.app
}

func (m *manage) init() {
	m.loadAppConfig()
}

func (m *manage) loadAppConfig() {
	path := FindConfigPath("app.conf")
	if path == "" {
		m.app.init()
		return
	}

	doc, err := loadXML(path)
	err = m.app.load(doc)
	if err != nil {
		panic(err)
	}
}