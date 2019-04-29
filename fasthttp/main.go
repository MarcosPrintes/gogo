package main

func main() {
	var app = App{}

	app.Initialize("root", "", "rest_api_example")

	app.Run("localhost:3005")

}
