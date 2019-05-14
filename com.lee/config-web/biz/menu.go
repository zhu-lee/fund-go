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

var (
	Menu = new(menuBiz)
)

type menuBiz struct {
	MenuList [] *entity.MenuNode
	MenuMap  map[string]*entity.MenuNode
	sync.Once
}

func (m *menuBiz) GetMenuInfo(url string) (mi *entity.MenuInfo) {
	m.Once.Do(m.loadMenuConf)

	mi = &entity.MenuInfo{
		TopMenus: m.MenuList,
	}

	c := m.MenuMap[url]
	if c == nil {
		return
	}

	mi.CurrentMenu = c
	mi.LeftMenus = c.GetTop().Items
	mi.Breadcrumb = make([]*entity.MenuNode, c.Level)//存放当前menu的到顶层menu
	for i := c.Level - 1; i >= 0; i-- {
		mi.Breadcrumb[i] = c
		if i > 0 {
			c = c.Parent
		}
	}
	return
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

	err = xml.Unmarshal([]byte(bytes), &menu)
	if err != nil {
		log.Log.Error("failed to parse menu.conf：%v", err)
		return
	}

	m.MenuList = menu.Items
	m.MenuMap = make(map[string]*entity.MenuNode)

	for i, n := range m.MenuList {
		n.Level = 1
		n.SetID(i + 1)
		m.MenuMap[n.Url] = n
		m.initMenu(n)
	}
}

func (m *menuBiz) initMenu(n *entity.MenuNode) {
	for i, sm := range n.Items {
		sm.Parent = n
		sm.Level = n.Level + 1
		sm.SetID(i + 1)

		m.MenuMap[sm.Url] = sm
		m.initMenu(sm)

		if !sm.Hidden {
			n.VisibleItems = append(n.VisibleItems, sm)
		}
	}
}
