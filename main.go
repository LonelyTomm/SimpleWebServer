package main

import (
	"fmt"
	"web/serv/fmwk/app"
)

func main() {
	app := app.NewApplication()
	app.StartListening()
	fmt.Printf("%v Hello, world.", app)
}
