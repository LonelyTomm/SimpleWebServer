package controllers

import "web/serv/fmwk/renderer"

type Application interface {
	GetRenderer() *renderer.Renderer
}
