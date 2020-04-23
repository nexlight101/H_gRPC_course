package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nexlight101/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am a gRPC client")

	// Create connection to the server
	options := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	// CLose the connection at exit
	defer cc.Close()

	// Establish a new client
	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Client activated: %v\n", c)

	// Call doUnary Greet RPC
	doUnary(c)
}
func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			First_Name: "Hendrink",
			Last_Name:  "Pienaar",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Cannot get response from server: %v\n", err)
	}
	fmt.Println(res.GetResult())
}
