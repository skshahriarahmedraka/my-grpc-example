package main 

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"r1/proto"
)

func Panicking(err error)  {
	if err != nil {
		log.Println("error : ",err)
	}
}

func main (){
	lis ,err := net.Listen("tcp",":9000")
	Panicking(err)

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer,&s)

	err =grpcServer.Serve(lis) 

	Panicking(err)

	log.Println("listening in port 9000")

}