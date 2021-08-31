package main

import (
	//"context"
	"fmt"
	"io"
	"log"
	"net"

	"server/messagepb"

	"google.golang.org/grpc"
)

func CheckError(err error) {
    if err != nil {
        log.Fatalln(err)
    }
}

type server struct{

}

func (s *server) Speaking(srv messagepb.Conversation_SpeakingServer) error {
	fmt.Println("server reciving streaming ...")

	for {
		req,err := srv.Recv()
		if err ==io.EOF {
			return nil
		}
		CheckError(err)

		firstName := req.GetClientString()
		result := "salam bro "+ firstName

		sendErr := srv.Send(&messagepb.LongResponse{
			ServerString: result,
		})

		CheckError(sendErr)
	}

}


func main(){

	lis ,err := net.Listen("tcp",":50052")
	CheckError( err)

	MyServer := grpc.NewServer()

	messagepb.RegisterConversationServer(MyServer,&server{})

	fmt.Println("server listening 50052")
	err = MyServer.Serve(lis) 

	CheckError(err)


}