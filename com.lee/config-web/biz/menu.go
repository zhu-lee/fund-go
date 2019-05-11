package biz

import (
	"com.lee/config-web/entity"
	"com.lee/fund/config"
	"com.lee/fund/log"
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
	"sync"
)

type menuBiz struct {
	MenuList [] *entity.MenuNode
	MenuMap  map[string]*entity.MenuNode
	sync.Once
}

func (m *menuBiz) GetMenuInfo(url string) (mi *entity.MenuInfo) {
	m.Once.Do(loadmenu)
}

func (m *menuBiz) loadMenuConf() {
	path := filepath.Join(config.FindConfigPath(), "menu.conf")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Log.Error("failed to read [menu.conf] file：%v", err)
		return
	}

	menu := struct {
		Items []*entity.MenuNode `xml:"item"`
	}{}

	err = xml.Unmarshal([]byte(bytes),&menu)
	if err != nil {
		log.Log.Error("failed to parse menu.conf：%v",err)
		return
	}
	m.MenuList=menu.Items
	m.MenuMap = make(map[string]*entity.MenuNode)
	for i,m := range m.MenuList {
		m.Level=1
		m
	}
}
