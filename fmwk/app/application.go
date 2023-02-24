package app

import (
	"net/http"
	"web/serv/fmwk/renderer"
	"web/serv/fmwk/router"
)

type Application struct {
	renderer       renderer.Renderer
	router         router.Router
	ResponseWriter *http.ResponseWriter
	Request        *http.Request
}

func NewApplication(request *http.Request, responseWriter *http.ResponseWriter) Application {
	return Application{
		renderer:       renderer.NewRenderer(responseWriter),
		router:         router.NewRouter(),
		ResponseWriter: responseWriter,
		Request:        request,
	}
}

func (app *Application) GetRenderer() *renderer.Renderer {
	return &app.renderer
}

func (app *Application) GetRouter() *router.Router {
	return &app.router
}
