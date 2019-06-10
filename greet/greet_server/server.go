package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/rksmannem/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
}

// Greet ...
func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("SERVER: Greet function was invoked with: %v\n", req)
	firstName := req.GetGreeting().GetFirstName()

	reslt := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: reslt,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello World!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to Listen: %v\n", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
	}
}
