package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", handlerHttp)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("http connection error: ", err.Error())
		return
	}
}

func handlerHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("http server is 'workando' ")
}
