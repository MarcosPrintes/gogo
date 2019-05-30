package main

import (
	"context"
	"fmt"
	"log"

	"github.com.br/MarcosPrintes/grpcTest/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = "localhost:8088"
)

func main() {

	var conn *grpc.ClientConn

	//create the client tsl credentials
	creds, err := credentials.NewClientTLSFromFile("../cert/server.crt", "")

	if err != nil {
		log.Fatalf("culd not load tls cert: %s", err)
	}

	// connection, err := grpc.Dial(port, grpc.WithInsecure())
	conn, err = grpc.Dial(port, grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	client := proto.NewPingClient(conn)

	// Ping(client, &proto.PingMessage{PingMsg: "a message from client"})
	Login(client, &proto.LoginRequest{Name: "marcos", Password: "1234567"})
}

func Ping(client proto.PingClient, request *proto.PingMessage) {
	response, err := client.PingMethod(context.Background(), request)

	if err != nil {
		log.Fatal("ping error: ", err.Error())
	}

	fmt.Println("ping response: ", response)
}

func Login(client proto.PingClient, request *proto.LoginRequest) {
	response, err := client.LoginPing(context.Background(), request)

	if err != nil {
		log.Fatal("Login error")
	}

	fmt.Println("Login response => ", response)
}
