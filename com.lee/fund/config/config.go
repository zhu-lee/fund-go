package config

import (
	x "com.lee/fund/xml"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var ConfDir string

func FindConfigPath(fileName string) string {
	p, err := getCurrentPath()
	if err != nil {
		return ""
	}
	p = filepath.Join(p, fileName)
	if _, err := os.Stat(p); err == nil {
		return p
	}
	return ""
}

func getCurrentPath() (string, error) {
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

func getConfigDir() string {
	if ConfDir == "" {
		folder, err := getCurrentPath()
		if err != nil {
			panic(err)
		}
		ConfDir = filepath.Join(folder, "config")
	}
	return ConfDir
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