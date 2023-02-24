package renderer

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Renderer struct {
	responseWriter *http.ResponseWriter
}

func NewRenderer(responseWriter *http.ResponseWriter) Renderer {
	return Renderer{
		responseWriter: responseWriter,
	}
}

func (r *Renderer) RenderView(viewPath string) {
	defaultViewPath := "/home/galv/goprj/myserv/views"

	resultString := ""
	if _, err := os.Stat(defaultViewPath + viewPath); err == nil {
		fileBytes, err := ioutil.ReadFile(defaultViewPath + viewPath)
		if err != nil {
			panic(err)
		}

		resultString = string(fileBytes)
	}

	io.WriteString(*r.responseWriter, resultString)
}
