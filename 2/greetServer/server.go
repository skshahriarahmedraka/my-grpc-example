package main

import (
	// "fmt"
	// "log"
	// "net"
	"log"
	"net"
	"r2/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func main () {
	println("sk shahriar ahmed raka \n Server is running !!! \n ")

	lis ,err := net.Listen("tcp",":50051")

	if err != nil {
		log.Fatalf("failed to listen %v ", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s,&server{})

	if err := s.Serve(lis) ; err != nil {
		log.Fatalf("fail to serve %v ", err)
	}
}