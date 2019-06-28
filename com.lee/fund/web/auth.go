package web

type User interface {
	ID() string
	Name() string
	Anonymous() bool
	Data() interface{}
}

type Auth interface {
	CookieUser(ctx *Context) User
	SignIn(ctx *Context, id, name string, remember bool)
	SignOut(ctx *Context)
}