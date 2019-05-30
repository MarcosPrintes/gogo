package api

import (
	"context"
	"fmt"

	"github.com.br/MarcosPrintes/grpcTest/proto"
)

type Server struct {
}

// func NewServer() *Server {
// 	return &Server{}
// }

func (server *Server) PingMethod(ctx context.Context, msg *proto.PingMessage) (*proto.PingResponse, error) {
	fmt.Println("message => ", msg)
	return &proto.PingResponse{Response: "msg"}, nil
}

func (server *Server) LoginPing(ctx context.Context, loginRequest *proto.LoginRequest) (*proto.LoginResponse, error) {
	return &proto.LoginResponse{
		State:   "logged",
		Code:    200,
		Message: "login ok",
	}, nil

}
