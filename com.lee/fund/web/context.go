package web

import (
	"fmt"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Context struct {
	Request    *http.Request
	Response   http.ResponseWriter
	Url        *url.URL
	parameters url.Values
	App        *App
	data       map[interface{}]interface{}
}

func (ctx *Context) GetParamString(key string) string {
	if ctx.parameters == nil {
		ctx.Request.ParseForm()
		ctx.parameters = ctx.Request.Form
	}
	return ctx.parameters.Get(key)
}

func (ctx *Context) GetParamInt(key string, defaultValue int) int {
	v := ctx.GetParamString(key)
	if v != "" {
		if value, err := strconv.Atoi(v); err == nil {
			return value
		}
	}
	return defaultValue
}

func (ctx *Context) GetParamBool(key string, defaultValue bool) bool {
	v := ctx.GetParamString(key)
	if v != "" {
		if value, err := strconv.ParseBool(v); err == nil {
			return value
		}
	}
	return defaultValue
}

/**
 * 获取上传文件
 */
func (ctx *Context) GetFile(name string) (multipart.File, *multipart.FileHeader, error) {
	return ctx.Request.FormFile(name)
}

func (ctx *Context) GetCookie(name string) (cookie *http.Cookie) {
	if cookie, err := ctx.Request.Cookie(name); err == nil {
		return cookie
	}
	return nil
}

func (ctx *Context) GetHeader(name string) string {
	return ctx.Request.Header.Get(name)
}

func (ctx *Context) GetLocalIp() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("获取本机IP出错：", err, "返回 127.0.0.1")
		return "127.0.0.1"
	}

	for _, addr := range addresses {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if ip == nil {
			continue
		}
		ip = ip.To4()
		if ip == nil {
			continue
		}
		return ip.String()
	}
	return "127.0.0.1"
}

func (ctx *Context) GetRemoteAddr() (ip string, port int) {
	ip = ctx.Request.Header.Get("X-Proxy-Real-IP")
	if len(ip) == 0 {
		ip = ctx.Request.Header.Get("X-Real-IP")
	}

	arr := strings.Split(ctx.Request.RemoteAddr, ":")
	if len(ip) == 0 && len(arr) > 0 {
		ip = arr[0]
	}
	if len(arr) > 1 {
		if intVa, err := strconv.Atoi(arr[1]); err == nil {
			port = intVa
		}
	}
	return
}

func (ctx *Context) GetCssUrl(url string) string {
	return url + "?v=" + ctx.App.Config.Version
}

func (ctx *Context) GetJsUrl(url string) string {
	return url + "?v=" + ctx.App.Config.Version
}
