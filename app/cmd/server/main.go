package main

import (
	pb "github.com/k-vanio/gRPC-with-docker/internal/pb/auth"
	"github.com/k-vanio/gRPC-with-docker/internal/services/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	service := auth.NewAuthService()

	gServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(gServer, service)
	// Register reflection service on gRPC server.
	reflection.Register(gServer)

	// open tcp connection
	list, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	// start grpc server
	log.Printf("Starting gRPC server on %s", ":50051")
	if err := gServer.Serve(list); err != nil {
		panic(err)
	}
}
