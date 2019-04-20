package main

/**
EXAMPLE GO CRUD: Creation of a REST API that allow us to read, delete, insert and upadte articles on our site

	- a simple server to handle requests http
	- 3 distincts functions
		- homePage()       - handle all request to root URL
		- handleRequests() - match the URL path hit with a defined function
		- main()           - will start api
	PKG's
	- http, package to provide http client and server implementations

**/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
	- ResponseWriter
		- interaface used by a http handler to construct http response
		- may not be used after Handler.ServerHTTP has returned
*/

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage")
	fmt.Println("Endpoint hit: HomePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Artigo 1", Description: "Description Artigo 1", Content: "Content Artigo 1"},
		Article{Title: "Artigo 2", Description: "Description Artigo 2", Content: "Content Artigo 2"},
	}

	fmt.Println("Endpoint hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

func eita(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Eita man!")
	fmt.Println("d")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", returnAllArticles)
	http.HandleFunc("/eita", eita)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

type Article struct {
	Title       string `json: "Title"`
	Description string ` json: "Description"`
	Content     string `json: "Content"`
}

type Articles []Article

func main() {
	handleRequests()
}
