package app

import (
	"io/ioutil"
	"os"
)

type Renderer struct {
}

func (r *Renderer) RenderView(viewPath string) string {
	defaultViewPath := "/home/galv/goprj/myserv/views"

	if _, err := os.Stat(defaultViewPath + viewPath); err == nil {
		fileBytes, err := ioutil.ReadFile(defaultViewPath + viewPath)
		if err != nil {
			panic(err)
		}

		return string(fileBytes)
	}
	return ""
}
