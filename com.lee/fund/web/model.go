package web

//站点Model
type SM struct {//todo usage
	Version   string
	Anonymous string
	UserName  string
}

func NewSM(ctx *Context) *SM {
	sm := SM{
		Version:ctx.App.Config.Version,
		Anonymous:"1",
		UserName:"",
	}
	return &sm
}

//页面Model
type PM struct {
	Title   string
	Metas   map[string]string
	Styles  []string
	Scripts []string
}

func (p *PM) AddCss(ctx *Context, css ...string) {
	for _, c := range css {
		if c!= "" {
			p.Styles = append(p.Styles, ctx.GetCssUrl(c))
		}
	}
}

func (p *PM) AddJs(ctx *Context, js ...string) {
	for _, j := range js {
		if j!= "" {
			p.Scripts = append(p.Scripts, ctx.GetJsUrl(j))
		}
	}
}