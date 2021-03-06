package config

import (
	x "com.lee/fund/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var ConfDir string

func findConfigPath() string {
	if ConfDir == "" {
		folder, err := GetAppFolder()
		if err != nil {
			panic(err)
		}
		ConfDir = filepath.Join(folder, "config")
		if _, err := os.Stat(ConfDir); err == nil {
			fmt.Printf("find config folder：%s\n", ConfDir)
			return ConfDir
		}
	}
	return ConfDir
}

func GetAppFolder() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}

	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "\\", "/", -1)
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		return "", errors.New("no executable file path found")
	}
	return string(path[0 : i+1]), nil
}

func loadXML(filePath string) (*x.Document, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	text := string(bytes)
	doc := x.New()
	if err := doc.LoadString(text, nil); err != nil {
		return nil, err
	}
	return doc, nil
}

func GetConfigPath(fileName string) string {
	return filepath.Join(findConfigPath(), fileName)
}

// Filter 返回经过 Profile 过滤处理过的配置文件内容
func Filter(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return getProfileConf().filter(string(bytes)), nil
}
