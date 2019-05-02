package main

import . "github.com.br/MarcosPrintes/plc/controller"

func main() {
	var app = App{}
	app.Initialize("root", "", "places")
	app.Run(":8089")
}
