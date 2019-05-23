package main

import (
	"log"
	"net"

	"github.com.br/MarcosPrintes/grpcTest/api"
	"github.com.br/MarcosPrintes/grpcTest/proto"
	"google.golang.org/grpc"
)

func main() {

	list, err := net.Listen("tcp", "localhost:8088")

	if err != nil {
		log.Fatal("error start grpc server => ", err.Error())
	}
	server := api.Server{}
	grpcServer := grpc.NewServer()
	proto.RegisterPingServer(grpcServer, &server)
	grpcServer.Serve(list)
}
