package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nexlight101/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	// Create a server type
	type server struct{}

	fmt.Println("Hello from grpc")

	//Create a listener
	lis, err := net.Listen("tcp", "0.0.0.0:500051")
	// handle error
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	// check if the server is serving the listener
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
