package main

import (
	"context"
	"log"
	"net"

	pb "github.com/omkardusane/gobase/grpctest/protos"

	grpc "google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGeneralMessageServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %s", req.Name)
	return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGeneralMessageServiceServer(grpcServer, &server{})

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

//go run client/client.go --message="Hello from the command line!"
