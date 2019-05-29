package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com.br/MarcosPrintes/consigment-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	fmt.Println("create")
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

type service struct {
	repo repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	fmt.Println("CreateConsignment")
	c, err := s.repo.Create(req)
	if err != nil {
		log.Fatal("error =>" + err.Error())
		return nil, err
	}
	return &pb.Response{Created: true, Consignment: c}, nil
}

func main() {
	repo := &Repository{}

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("listen error" + err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterShippingServiceServer(s, &service{repo})

	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
