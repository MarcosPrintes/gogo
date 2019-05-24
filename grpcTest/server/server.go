package main

import (
	"fmt"
	"log"
	"net"

	"github.com.br/MarcosPrintes/grpcTest/api"
	"github.com.br/MarcosPrintes/grpcTest/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func init() {
}

func main() {

	list, err := net.Listen("tcp", "localhost:8088")

	if err != nil {
		log.Fatal("error start grpc server => ", err.Error())
	}

	server := api.Server{}

	creds, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Fatal("credentials error => ", err.Error())
	}

	//array options  with credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	fmt.Println("credentials options", opts)

	// grpcServer := grpc.NewServer()
	grpcServer := grpc.NewServer(opts...) // pass options to new server grpc
	proto.RegisterPingServer(grpcServer, &server)
	if err := grpcServer.Serve(list); err != nil {
		log.Fatal("grpc server error credentials:", err.Error())
	}
}
