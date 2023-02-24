package controllers

type Renderer interface {
	RenderView(viewPath string) string
}

type Application interface {
	GetRenderer() interface{}
}
