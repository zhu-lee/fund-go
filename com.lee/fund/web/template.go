package web

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	_Functions template.FuncMap
)

type ViewEngine struct {
	suffix    string
	dir       string
	once      sync.Once
	templates *template.Template
}

func (v *ViewEngine) Get(name string) *template.Template {
	v.once.Do(v.compile)
	t := v.templates.Lookup(name)
	if t == nil {
		panic(fmt.Errorf("can't parse template：%s", name))
	}
	return t
}

func (v *ViewEngine) compile() {
	v.templates = template.New(v.dir)

	_ = filepath.Walk(v.dir, func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() { //空或路径则返回
			return nil
		}

		rel, err := filepath.Rel(v.dir, path) //返回访问path的相对路径
		if err != nil {
			return err
		}

		ext := ""
		if strings.Contains(rel, ".") {
			ext = filepath.Ext(rel) //获取后缀名
		}

		if ext == v.suffix {
			buf, err := ioutil.ReadFile(path) //读取html文件
			if err != nil {
				panic(err)
			}
			name := rel[0 : len(rel)-len(ext)] //相对路径名（去掉后缀）

			tmpl := v.templates.New(filepath.ToSlash(name)) //创建的template会保留原有的值
			tmpl.Funcs(_Functions)

			template.Must(tmpl.Parse(string(buf))) //读取的html会缓存到namespace的set里
		}

		return nil
	})
}
