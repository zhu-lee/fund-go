package web

import (
	"com.lee/fund/log"
	"com.lee/fund/util/crypto"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"
)

const cookieRememberDays = 30

type cookieAuthUser struct {
	UserID   string      `json:"id"`
	UserName string      `json:"name"`
	UserData interface{} `json:"data,omitempty"`
	Life     int64       `json:"life"`
	Remember bool        `json:"remember"`
}

func (u *cookieAuthUser) ID() string {
	return u.UserID
}
func (u *cookieAuthUser) Name() string {
	return u.UserName
}
func (u *cookieAuthUser) Anonymous() bool {
	return u.UserID == ""
}
func (u *cookieAuthUser) Data() interface{} {
	return u.UserData
}

type CookieAuth struct {
	cfg        *Config
	cookieTime time.Duration
	aesKey     []byte
}

func newCookieAuth(cfg *Config) (Auth, error) {
	cookieTime := cfg.CookieTime
	if cookieTime <= 0 {
		cookieTime = time.Minute * 30
	}

	key := "5*369&2f-3adc_47" //todo 无构建信息，自定一个串
	return &CookieAuth{
		cfg:        cfg,
		cookieTime: cookieTime,
		aesKey:     []byte(key),
	}, nil
}

func (a *CookieAuth) CookieUser(ctx *Context) User {
	if cookie := ctx.GetCookie(a.cfg.CookieName); cookie != nil {
		if user := a.decodeUser(cookie.Value); user != nil {
			if !user.Remember {
				a.updateCookie(ctx, user)
			}
			return user
		}
	}
	return nil
}

func (a *CookieAuth) SignIn(ctx *Context, id, name string, remeber bool) {

}

func (a *CookieAuth) SignOut(ctx *Context) {

}

func (a *CookieAuth) decodeUser(value string) (user *cookieAuthUser) {
	defer func() {
		if err := recover(); err != nil {
			log.Log.Error("auth", "aes decrypt crashed:%v", err)
		}
	}()

	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		log.Log.Error("auth", "base64 decode failed :%v", err)
		return
	}

	data, err = crypto.AesDecrypt(data, a.aesKey)
	if err != nil {
		log.Log.Error("auth", "aes decrypt failed :%v", err)
		return
	}

	u := &cookieAuthUser{}
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Log.Error("auth", "json decode failed :%v", err)
		return
	}

	if u.Life > time.Now().Unix() {
		user = u
	}
	return
}

func (a *CookieAuth) updateCookie(ctx *Context, user *cookieAuthUser) {
	var life time.Time
	if user.Remember {
		life = time.Now().AddDate(0, 0, cookieRememberDays)
	} else {
		life = time.Now().Add(a.cookieTime)
	}

	user.Life = life.Unix()
	value, err := a.encodeUser(user)
	if err != nil {
		panic(err)
	}

	cookie := &http.Cookie{
		Domain:  a.cfg.CookieDomain,
		Name:    a.cfg.CookieName,
		Value:   value,
		Expires: life,
		Path:    "/",
	}
	http.SetCookie(ctx.Response, cookie)
}

func (a *CookieAuth) encodeUser(u *cookieAuthUser) (string, error) {
	data, err := json.Marshal(u)
	if err != nil {
		log.Log.Error("auth", "json encode failed:%v", err)
		return "", err
	}

	data, err = crypto.AesDecrypt(data, a.aesKey)
	if err != nil {
		log.Log.Error("auth", "aes encrypt failed:%v", err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}
