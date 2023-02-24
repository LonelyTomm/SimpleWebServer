package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"web/serv/fmwk/controllers"
)

type Application struct {
	renderer Renderer
}

var controllerMap = map[string]interface{}{
	"/": controllers.Test}

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defaultWebFolderPath := "/home/galv/goprj/myserv/web"
	fmt.Println(req.URL.Path)

	if _, err := os.Stat(defaultWebFolderPath + req.URL.Path); err == nil && req.URL.Path != "/" {
		fileBytes, err := ioutil.ReadFile(defaultWebFolderPath + req.URL.Path)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(fileBytes)
		return
	}

	controllerAction, ok := controllerMap[req.URL.Path]

	if ok {
		app := NewApplication()
		in := []reflect.Value{reflect.ValueOf(w), reflect.ValueOf(req), reflect.ValueOf(&app)}
		reflect.ValueOf(controllerAction).Call(in)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func NewApplication() Application {
	return Application{
		renderer: Renderer{},
	}
}

func (app Application) StartListening() {

	err := http.ListenAndServe("127.0.0.1:8080", &Handler{})

	if err != nil {
		fmt.Println(err.Error())
	}
}

func (app *Application) GetRenderer() *Renderer {
	return &app.renderer
}
