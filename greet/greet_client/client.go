package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rksmannem/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I 'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v\n", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created client: %f\n", c)
	invokeUnaryRPC(c)
}

func invokeUnaryRPC(c greetpb.GreetServiceClient) {
	fmt.Println("Starting invokeUnaryRPC ...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ramakrishna",
			LastName:  "Mannem",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("RPC method/api failed: %v\n", err)
	}
	log.Printf("RPC Method Returns: %v\n", res.GetResult())
}
