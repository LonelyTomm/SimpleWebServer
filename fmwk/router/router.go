package router

import "web/serv/fmwk/controllers"

type Router struct {
	routeMap map[string]interface{}
}

func NewRouter() Router {
	return Router{
		routeMap: map[string]interface{}{"/": controllers.HelloWorld},
	}
}

func (r *Router) GetRouteMap() map[string]interface{} {
	return (*r).routeMap
}
