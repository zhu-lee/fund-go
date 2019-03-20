package config

type appConfig struct {
	Web    *webSetting
	App    *appSetting
	Custom map[string]string
}

type webSetting struct {
	HttpAddr string
	HttpPort string
	LogEnable     bool
	Anonymous     bool
	CookieName    string
	CookieDomain string
	CookieTime  int8
	UrlDefault  string
	UrlSignIn   string
	UrlNotFound string
	UrlNoRight  string
}

type appSetting struct {
	AppName   string
	Debug     bool
	LogLevel  int
	GlobalEnv string
}

func (a *appConfig) init() {
	a.Custom = make(map[string]string)
	a.App =
}
