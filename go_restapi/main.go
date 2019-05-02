package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com.br/MarcosPrintes/go_restapi/config"

	. "github.com.br/MarcosPrintes/go_restapi/config/dao"
	movierouter "github.com.br/MarcosPrintes/go_restapi/router"
	"github.com/gorilla/mux"
)

var dao = MoviesDAO{}
var config = Config{}

// Inicializando e estabilizando a conexão com db
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

//informando as rotas e a porta
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/movies", movierouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/movies/{id}", movierouter.GetById).Methods("GET")
	r.HandleFunc("/api/v1/movies", movierouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/movies/update/{id}", movierouter.Update).Methods("PUT")
	r.HandleFunc("/api/v1/movies/delete/{id}", movierouter.Delete).Methods("DELETE")

	port := "localhost:3002"
	fmt.Println("server is running on port", port)
	log.Fatal(http.ListenAndServe(port, r))

}
