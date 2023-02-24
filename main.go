package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"web/serv/fmwk/app"
)

func main() {
	err := http.ListenAndServe("127.0.0.1:8080", &Handler{})

	if err != nil {
		fmt.Println(err.Error())
	}
}

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	app := app.NewApplication(req, &w)
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

	controllerAction, ok := app.GetRouter().GetRouteMap()[req.URL.Path]

	if ok {
		in := []reflect.Value{reflect.ValueOf(&app)}
		reflect.ValueOf(controllerAction).Call(in)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
