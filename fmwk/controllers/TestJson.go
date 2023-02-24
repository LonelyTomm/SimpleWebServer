package controllers

func TestJson(app Application) {
	app.GetRenderer().RenderJSON(map[string]string{"name": "Seva"})
}
