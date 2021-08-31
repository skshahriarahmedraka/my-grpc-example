package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"google.golang.org/grpc"
	"server/messagepb"

)
func CheckError(err error){
	if err !=nil {
		panic(err)
	}
}

type server struct {
}

func (s *server) Speaking(srv messagepb.Conversation_SpeakingServer) error {
	fmt.Println("server is reciving  stream of data from client ")
	result := ""

	for {
		req,err := srv.Recv()
		if err == io.EOF{
			return srv.SendAndClose(&messagepb.ShortResponse{
				ResponseString: result,
			})
		}
		if err != nil {
			log.Fatalln("error : ",err)
		}
		firstName := req.GetRequestString()
		result += "salam "+firstName+ " !!! "
	}



	return nil
}


func main (){
	lis,err := net.Listen("tcp",":50051")
	CheckError(err)

	myServer := grpc.NewServer()


	messagepb.RegisterConversationServer(myServer,&server{})
	fmt.Println("server listening 50051")
	err = myServer.Serve(lis)
	CheckError(err)
}

