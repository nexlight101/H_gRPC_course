package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/nexlight101/greet/greetpb"
	"google.golang.org/grpc"
)

// Create a server type
type server struct{}

// Implemente Greet method
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	// Extract the name from the request
	fullname := req.GetGreeting().GetFirst_Name() + " " + req.GetGreeting().GetLast_Name()

	// Responde with a greeting containing the name
	result := "Hello " + fullname
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {

	fmt.Println("Hello from grpc")

	//Create a listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
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
