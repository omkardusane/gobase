package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/omkardusane/gobase/grpctest/protos"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Define the message flag
	message := flag.String("message", "Gopher", "The message to send to the gRPC server")
	flag.Parse()

	// Check if a message is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run client.go --message=<your-message>")
		os.Exit(1)
	}

	// Create a connection to the server
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client instance
	client := pb.NewGeneralMessageServiceClient(conn)

	// Make the gRPC call with the provided message
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SayHello(ctx, &pb.HelloRequest{Name: *message})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	log.Printf("Response from server: %s", response.Message)
}
