package web

type App struct {
	Config *Config
	route *routeEngine
	//viewEngin *ViewEngin
}

func NewDefaultApp() *App {

	return &App{
		route: newRoute(),
	}
}

func (app *App) RegisterController(url string, controller interface{}) {
	app.route.RegisterController(url, controller)
}
