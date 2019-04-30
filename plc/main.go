package main

func main() {
	var app = App{}
	app.Initialize("root", "", "places")
	app.Run("localhost:8087")
}
