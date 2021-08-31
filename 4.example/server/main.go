package main

import (
	"context"
	"log"
	"net"

	"r4/messagepb"

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
	log.Printf("Received : %v ",in.GetMyRequest())
	return &messagepb.SpeakResponse{MyResponse : "hello "+in.GetMyRequest()},nil
}


func main(){

	lis ,err := net.Listen("tcp",":50051")
	Panicking("error : ", err)

	MyServer := grpc.NewServer()

	messagepb.RegisterConversationServer(MyServer,&server{})

	err = MyServer.Serve(lis) 

	CheckError("error : ",err)


}