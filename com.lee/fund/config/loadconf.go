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

var ldc = newLoadConf()

type loadConf struct {
	once        sync.Once
	appConf     *appConf
	profileConf *profileConf
}

func newLoadConf() *loadConf {
	return &loadConf{
		appConf:     new(appConf),
		profileConf: new(profileConf),
	}
}

func GetAppConf() *appConf {
	ldc.once.Do(ldc.init)
	return ldc.appConf
}

func getProfileConf() *profileConf {
	ldc.once.Do(ldc.init)
	return ldc.profileConf
}

func (ld *loadConf) init() {
	ld.loadProfile()
	ld.loadAppConf()
}

func (ld *loadConf) loadProfile() {
	activeProfiles := *app.GetProfileFlag()
	if activeProfiles == "" {
		activeProfiles = os.Getenv(ActiveProfileEnv)
	}
	if activeProfiles != "" {
		ld.profileConf.active = strings.Split(activeProfiles, ",")
	}
	fmt.Println("config > activeProfile：", activeProfiles)

	path := GetConfigPath("profile.conf")
	if path == "" {
		fmt.Println("can't find profile.conf")
		return
	}

	fmt.Printf("load profile.conf：%s\n", path)

	content := ld.readFile(path)
	err := ld.profileConf.load(content)
	if err != nil {
		panic(err)
	}
}

func (ld *loadConf) loadAppConf() {
	path := GetConfigPath("app.conf")
	if path == "" {
		fmt.Println("can't find app.conf")
		ld.appConf.init()
		return
	}

	doc, err := loadXML(path)
	err = ld.appConf.setAppCof(doc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("load app.conf：%s\n", path)
}

func (ld *loadConf) readFile(path string) [] byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("read [%s] failed：%v", path, err))
	}
	return bytes
}
