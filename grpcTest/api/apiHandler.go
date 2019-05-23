package api

import (
	"context"
	"log"

	"github.com.br/MarcosPrintes/grpcTest/proto"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) PingMethod(ctx context.Context, msg *proto.PingMessage) (*proto.PingMessage, error) {
	log.Println(msg)
	return &proto.PingMessage{PingMsg: "msg"}, nil
}
