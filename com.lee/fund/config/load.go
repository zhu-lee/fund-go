package config

import (
	"fmt"
	"path/filepath"
	"sync"
)

const (
	ACTIVE_PROFILE_ENV = "FUND-PROFILE-ACTIVE"
)

var ld = newLoad()

type load struct {
	once    sync.Once
	appConf *appConf
}

func newLoad() *load {
	return &load{
		appConf: new(appConf),
	}
}

func GetAppConf() *appConf {
	ld.once.Do(ld.init)
	return ld.appConf
}

func (m *load) init() {
	m.loadAppConf()
	//TODO loadGlobalConf
}

func (m *load) loadAppConf() {
	path := filepath.Join(FindConfigPath(), "app.conf")
	if path == "" {
		fmt.Println("can't find app.conf")
		m.appConf.init()
		return
	}

	doc, err := loadXML(path)
	err = m.appConf.setAppCof(doc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("load app.conf：%s\n", path)
}
