package main

import (
	"fmt"
	"log"
	"net"

	"github.com/andy_tete/grpc-test-project/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("helulu")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Fail roiii")
	}
	s := grpc.NewServer()
	//	greetpb.R
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("loiii")
	}
}
