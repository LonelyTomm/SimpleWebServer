package controllers

import (
	"fmt"
	"net/http"
	"reflect"
)

func Test(w http.ResponseWriter, r *http.Request, app Application) {
	fmt.Printf("%v", reflect.TypeOf(app))
	//io.WriteString(w, (*(*app).GetRenderer()).RenderView("/index.gotpl"))
}
