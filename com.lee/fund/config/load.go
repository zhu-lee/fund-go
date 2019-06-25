package config

import (
	"com.lee/fund/util/app"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

const (
	ActiveProfileEnv = "FUND_PROFILE_ACTIVE"
)

//初始化load
var ld = newLoad()

type load struct {
	once    sync.Once
	appConf *appConf
	profile *profileConfig
}

func newLoad() *load {
	return &load{
		appConf: new(appConf),
		profile: new(profileConfig),
	}
}

func GetAppConf() *appConf {
	ld.once.Do(ld.init)
	return ld.appConf
}

func getProfileConfig() *profileConfig {
	ld.once.Do(ld.init)
	return ld.profile
}

func (m *load) init() {
	m.loadProfile()
	m.loadAppConf()
}

//加载profile.conf配置文件
func (m *load) loadProfile() {
	activeProfiles := *app.GetProfileFlag()
	if activeProfiles == "" {
		activeProfiles = os.Getenv(ActiveProfileEnv)
	}
	if activeProfiles != "" {
		m.profile.active = strings.Split(activeProfiles, ",")
	}
	fmt.Println("config > activeProfile：", activeProfiles)

	path := GetConfigPath("profile.conf")
	if path == "" {
		fmt.Println("can't find profile.conf")
		return
	}

	fmt.Printf("load profile.conf：%s\n", path)

	content := m.readFile(path)
	err := m.profile.load(content)
	if err != nil {
		panic(err)
	}
}

//加载app.conf配置文件
func (m *load) loadAppConf() {
	path := GetConfigPath("app.conf")
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

func (m *load) readFile(path string) [] byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("read [%s] failed：%v", path, err))
	}
	return bytes
}
