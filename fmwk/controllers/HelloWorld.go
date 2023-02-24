package controllers

func HelloWorld(app Application) {
	app.GetRenderer().RenderView("/index.gotpl")
}
