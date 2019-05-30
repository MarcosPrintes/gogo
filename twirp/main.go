package main

import (
	"net/http"

	harberdashserver "github.com.br/MarcosPrintes/twirp/pb"
	"github.com.br/MarcosPrintes/twirp/server"
)

func main() {

	server := &server.Server{}

	twirpServer := harberdashserver.NewHarberDashServer(server, nil)

	http.ListenAndServe(":8080", twirpServer)
}
