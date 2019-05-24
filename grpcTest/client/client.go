package main

import (
	"context"
	"fmt"
	"log"

	"github.com.br/MarcosPrintes/grpcTest/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8088"
)

func main() {
	connection, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal()
	}
	defer connection.Close()

	mClient := proto.NewPingClient(connection)
	msg, err := mClient.PingMethod(context.Background(), &proto.PingMessage{PingMsg: "hmmmmmmmm!!"})
	if err != nil {
		log.Fatal("error => ", err.Error())
	}

	fmt.Println("message sent : ", creds)
	fmt.Println("message sent : ", msg)
}
