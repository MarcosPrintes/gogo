package server

import (
	"context"

	"github.com/twitchtv/twirp"

	pb "github.com.br/MarcosPrintes/twirp/pb"
)

type Server struct {
}

//MakeHat bla bla
func (s *Server) Makehat(ctx context.Context, size *pb.Size) (hat *pb.Hat, err error) {
	// fmt.Println("server makehat", size)
	if size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("error inches", "hat error")
	}

	return &pb.Hat{
		Inches: 10,
		Color:  "blue",
		Name:   "name",
	}, nil
}
