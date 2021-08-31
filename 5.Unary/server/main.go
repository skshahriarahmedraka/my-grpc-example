package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"server/messagepb"
	"google.golang.org/grpc"
)

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

type server struct{

}

func (s *server) Speaking(ctx context.Context, in *messagepb.SpeakRequest)(*messagepb.SpeakResponse,error){
	log.Printf("Received Client Request : %v \n",in.GetClient_Request())
	log.Printf("server LIstenning 50051 ...\n")

	return &messagepb.SpeakResponse{Server_Response : "hello world whats are you doing "+in.GetClient_Request()},nil
}


func main(){

	lis ,err := net.Listen("tcp",":50051")
	CheckError(err)

	MyServer := grpc.NewServer()

	messagepb.RegisterConversationServer(MyServer,&server{})
	
	fmt.Println("listening on 50051 ... ")
	err = MyServer.Serve(lis) 

	CheckError(err)


}