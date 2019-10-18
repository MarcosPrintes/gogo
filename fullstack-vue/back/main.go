package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {

	http.HandleFunc("/rota", handler)
	http.HandleFunc("/post", mpost)

	err := http.ListenAndServe(":8082", nil)

	if err != nil {
		log.Fatal("erro na conexão", err.Error())
		return
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Println("http is working")
}

func mpost(w http.ResponseWriter, r *http.Request) string {

	return json.NewEncoder(w).Encode()
}

func enableCors(w *http.ResponseWriter) {
	//response headers
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
