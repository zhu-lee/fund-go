package mgox

import (
	"com.lee/fund/config"
	"com.lee/fund/mgo"
	"com.lee/fund/xml"
	"fmt"
	"sync"
)

var (
	_Infos  map[string]*dbInfo
	_Locker sync.Mutex
)

// 数据库配置信息
type DBConfig struct {
	Name     string
	Settings config.SettingMap
}

type dbInfo struct {
	cfg     *DBConfig
	session *mgo.Session
}

// 获取指定名称的数据库配置
func GetConfig(name string) (*DBConfig, error) {
	di, err := getInfo(name)
	if err != nil {
		return nil, err
	}

	return di.cfg, nil
}

func getInfo(name string) (*dbInfo, error) {
	var err error

	if _Infos == nil {
		_Locker.Lock()
		defer _Locker.Unlock()

		if _Infos == nil {
			configPath := config.GetConfigPath("database.mongo.conf")
			_Infos, err = loadConfig(configPath)
			if err != nil {
				return nil, err
			}
		}
	}

	info, ok := _Infos[name]
	if !ok {
		return nil, fmt.Errorf("cannot find database: %s", name)
	}

	return info, nil
}

func loadConfig(configPath string) (map[string]*dbInfo, error) {
	content, err := config.Filter(configPath)
	if err != nil {
		return nil, err
	}

	doc := xml.New()
	err = doc.LoadString(content, nil)
	if err != nil {
		return nil, err
	}

	dbInfos := make(map[string]*dbInfo)
	dbNodes := doc.SelectNodes("", "database")
	for _, dbNode := range dbNodes {
		var cfg = &DBConfig{
			Name:     dbNode.As("", "name"),
			Settings: config.SettingMap{},
		}

		for _, settingNode := range dbNode.Children {
			cfg.Settings[settingNode.As("", "name")] = settingNode.As("", "value")
		}

		dbInfos[cfg.Name] = &dbInfo{cfg: cfg}
	}

	return dbInfos, nil
}
